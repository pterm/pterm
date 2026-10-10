// Package screen owns everything pterm writes to a terminal.
//
// Each writer gets one Screen. All output to that writer goes through the
// Screen under one lock, which is what makes pterm race-free: a spinner
// redrawing from its own goroutine and a Println from the main goroutine
// cannot interleave half-written lines.
//
// A Screen has two kinds of output. Write prints text that scrolls up like
// normal terminal output. A Region is a block of lines pinned below that
// output, which a live printer redraws in place. While regions are drawn,
// Write erases them, prints its text, and draws them again underneath, so
// normal output appears above a running spinner instead of tearing it.
//
// On writers that are not terminals, regions are never drawn. Closing a
// region prints its final content once, so logs and CI output get the end
// state of each live printer and none of the frames.
package screen

import (
	"io"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/term"

	"github.com/pterm/pterm/internal/text"
)

// Node marks a line as a node line: a line that starts with a status glyph
// and stands on the rail of a flow instead of being indented next to it.
// Printers prefix such lines with Node. Screens remove the marker before
// writing, and Strip removes it from strings that never reach a screen.
const Node = "\x00"

// Terminal is implemented by writers that are terminals without being an
// *os.File, such as the virtual terminal pterm's tests use.
type Terminal interface {
	IsTerminal() bool
	Size() (width, height int)
}

// Decorator rewrites output before a Screen writes it. A flow uses one to
// draw its rail next to everything printed while it runs. item is true for
// text passed to Write and Region.Close, which is a new item of output, and
// false for the content of a live region.
type Decorator func(s string, item bool) string

// Screen serializes all output to one writer.
type Screen struct {
	mu sync.Mutex
	w  io.Writer

	regions []*Region
	drawn   int  // rows the regions currently take on the terminal
	atBOL   bool // the cursor is at the beginning of a line
	hidden  bool // the cursor is hidden
	partial string

	decorators []decorator
	nextID     int
}

type decorator struct {
	id int
	fn Decorator
}

var (
	registryMu sync.Mutex
	registry   = map[any]*Screen{}
)

// For returns the Screen of w. Every call with the same writer returns the
// same Screen, so every printer writing to os.Stdout shares one lock.
func For(w io.Writer) *Screen {
	if w == nil {
		w = os.Stdout
	}

	key := any(w)
	if t := reflect.TypeOf(w); !t.Comparable() {
		// Writers that cannot be map keys get a Screen of their own.
		return &Screen{w: w, atBOL: true}
	}

	registryMu.Lock()
	defer registryMu.Unlock()

	s, ok := registry[key]
	if !ok {
		s = &Screen{w: w, atBOL: true}
		registry[key] = s
	}

	return s
}

// Writer returns the writer of s.
func (s *Screen) Writer() io.Writer {
	return s.w
}

// IsTerminal reports whether s writes to a terminal, so regions can be drawn.
func (s *Screen) IsTerminal() bool {
	return IsTerminal(s.w)
}

// IsTerminal reports whether w is a terminal.
func IsTerminal(w io.Writer) bool {
	if t, ok := w.(Terminal); ok {
		return t.IsTerminal()
	}

	f, ok := w.(interface{ Fd() uintptr })
	if !ok {
		return false
	}

	return term.IsTerminal(int(f.Fd())) //nolint:gosec // file descriptors fit in an int.
}

// Size returns the width and height of the terminal behind w. Writers that
// are not terminals report the COLUMNS and LINES environment variables, or 0
// when those are unset, which means "unknown": nothing gets wrapped or cut.
func Size(w io.Writer) (width, height int) {
	if t, ok := w.(Terminal); ok {
		return t.Size()
	}

	if f, ok := w.(interface{ Fd() uintptr }); ok {
		if width, height, err := term.GetSize(int(f.Fd())); err == nil && width > 0 { //nolint:gosec // file descriptors fit in an int.
			return width, height
		}
	}

	width, _ = strconv.Atoi(os.Getenv("COLUMNS"))
	height, _ = strconv.Atoi(os.Getenv("LINES"))

	return max(width, 0), max(height, 0)
}

// Write prints str as normal output, above any live region.
func (s *Screen) Write(str string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.writeItem(str)
}

// writeItem prints str above the regions. The caller holds s.mu.
func (s *Screen) writeItem(str string) {
	str = strings.ReplaceAll(s.decorate(str, true), Node, "")

	if s.drawn == 0 && len(s.regions) == 0 {
		s.raw(str)
		return
	}

	if !s.IsTerminal() {
		s.raw(str)
		return
	}

	// While regions are drawn, only whole lines go above them. A partial
	// line waits until it is finished or the last region closes.
	str = s.partial + str
	s.partial = ""

	if len(s.regions) > 0 {
		cut := strings.LastIndexByte(str, '\n')
		s.partial = str[cut+1:]
		str = str[:cut+1]
	}

	var b strings.Builder

	s.erase(&b)
	b.WriteString(str)

	if str != "" {
		s.atBOL = strings.HasSuffix(str, "\n")
	}

	s.draw(&b)
	s.finishIfIdle(&b)
	s.flush(&b)
}

// raw writes str unchanged and keeps track of the cursor column.
func (s *Screen) raw(str string) {
	if str == "" {
		return
	}

	_, _ = io.WriteString(s.w, str)
	s.atBOL = strings.HasSuffix(str, "\n")
}

func (s *Screen) flush(b *strings.Builder) {
	_, _ = io.WriteString(s.w, b.String())
}

// PushDecorator adds d to the decorators of s and returns a function that
// removes it again.
func (s *Screen) PushDecorator(d Decorator) (pop func()) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	id := s.nextID
	s.decorators = append(s.decorators, decorator{id: id, fn: d})

	return func() {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.decorators = slices.DeleteFunc(s.decorators, func(d decorator) bool { return d.id == id })
	}
}

// HasDecorators reports whether a decorator is active, which is the case
// while a flow runs on s.
func (s *Screen) HasDecorators() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.decorators) > 0
}

func (s *Screen) decorate(str string, item bool) string {
	for i := len(s.decorators) - 1; i >= 0; i-- {
		str = s.decorators[i].fn(str, item)
	}

	return str
}

// Strip removes node markers from str. Use it on rendered text that is
// returned to the user instead of being written to a screen.
func Strip(str string) string {
	return strings.ReplaceAll(str, Node, "")
}

// erase moves the cursor to the first row of the regions and clears
// everything below it. The caller holds s.mu.
func (s *Screen) erase(b *strings.Builder) {
	if s.drawn == 0 {
		return
	}

	b.WriteString("\r")

	if s.drawn > 1 {
		b.WriteString("\x1b[" + strconv.Itoa(s.drawn-1) + "A")
	}

	b.WriteString("\x1b[J")

	s.drawn = 0
	s.atBOL = true
}

// draw writes every region below the cursor. Lines are cut to the terminal
// width so each one takes exactly one row and erase can count them, and the
// regions together never take more rows than the terminal has, because rows
// that scrolled off cannot be erased. The caller holds s.mu.
func (s *Screen) draw(b *strings.Builder) {
	if len(s.regions) == 0 {
		return
	}

	if !s.atBOL {
		b.WriteString("\n")
		s.atBOL = true
	}

	if !s.hidden {
		b.WriteString("\x1b[?25l")
		s.hidden = true
	}

	width, height := Size(s.w)

	var lines []string

	for _, r := range s.regions {
		if r.content == "" {
			continue
		}

		content := strings.ReplaceAll(s.decorate(r.content, false), Node, "")
		lines = append(lines, strings.Split(content, "\n")...)
	}

	if height > 1 && len(lines) > height-1 {
		lines = lines[len(lines)-(height-1):]
	}

	for i, line := range lines {
		if width > 0 {
			line = text.Truncate(line, width-1, "…")
		}

		b.WriteString(line)

		if i < len(lines)-1 {
			b.WriteString("\n")
		}
	}

	s.drawn = len(lines)
	s.atBOL = s.drawn == 0
}

// redraw erases and draws the regions. The caller holds s.mu.
func (s *Screen) redraw() {
	if !s.IsTerminal() {
		return
	}

	var b strings.Builder

	s.erase(&b)
	s.draw(&b)
	s.finishIfIdle(&b)
	s.flush(&b)
}

// finishIfIdle ends the line under the last region and shows the cursor
// again once no region is left. The caller holds s.mu.
func (s *Screen) finishIfIdle(b *strings.Builder) {
	if len(s.regions) > 0 {
		return
	}

	if s.partial != "" {
		b.WriteString(s.partial)
		s.atBOL = strings.HasSuffix(s.partial, "\n")
		s.partial = ""
	}

	if s.hidden {
		b.WriteString("\x1b[?25h")
		s.hidden = false
	}
}
