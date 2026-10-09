package pterm

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"sync"
	"time"

	"atomicgo.dev/cursor"
	"atomicgo.dev/schedule"

	"github.com/pterm/pterm/internal"
)

// ActiveProgressBarPrinters contains all running ProgressbarPrinters.
// Generally, there should only be one active ProgressbarPrinter at a time.
var ActiveProgressBarPrinters []*ProgressbarPrinter

// DefaultProgressbar is the default ProgressbarPrinter.
var DefaultProgressbar = ProgressbarPrinter{
	Total:         100,
	BarCharacter:  "█",
	LastCharacter: "█",
	// Eighth-block glyphs, ordered from least to most filled, give the bar a
	// smooth edge that advances one eighth of a cell at a time.
	BarPartialCharacters:      []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	BarFiller:                 "░",
	BarFillerStyle:            &ThemeDefault.ProgressbarFillerStyle,
	ElapsedTimeRoundingFactor: time.Second,
	BarStyle:                  &ThemeDefault.ProgressbarBarStyle,
	TitleStyle:                &ThemeDefault.ProgressbarTitleStyle,
	ShowTitle:                 true,
	ShowCount:                 true,
	ShowPercentage:            true,
	ShowElapsedTime:           true,
	MaxWidth:                  80,
	Writer:                    os.Stderr,
}

// ProgressbarPrinter shows a progress animation in the terminal.
type ProgressbarPrinter struct {
	Title                     string
	Total                     int
	Current                   int
	BarCharacter              string
	LastCharacter             string
	ElapsedTimeRoundingFactor time.Duration

	// BarFiller is repeated to fill the unfilled portion of the bar. When
	// empty the remaining space is left blank.
	BarFiller string

	// BarFillerStyle styles the unfilled portion of the bar. When nil, the
	// theme's ProgressbarFillerStyle is used.
	BarFillerStyle *Style

	// BarPartialCharacters holds the glyphs used to draw the leading edge of
	// the bar at sub-character resolution, ordered from least to most filled.
	// When empty, the bar falls back to whole-cell rendering using
	// BarCharacter and LastCharacter.
	BarPartialCharacters []string

	MaxWidth int

	ShowElapsedTime bool
	ShowCount       bool
	ShowTitle       bool
	ShowPercentage  bool
	RemoveWhenDone  bool

	TitleStyle *Style
	BarStyle   *Style

	IsActive bool

	startedAt    time.Time
	rerenderTask *schedule.Task

	Writer io.Writer

	// mu serializes access to the printer's mutable state once Start has run.
	// It is a pointer so the value-receiver With* methods can copy the struct
	// without tripping go vet's copylocks check; Start lazily allocates it.
	mu *sync.RWMutex
}

// lock locks the printer's runtime mutex, allocating it on first use so the
// value-receiver builders can be used safely before Start.
func (p *ProgressbarPrinter) lock() {
	if p.mu == nil {
		p.mu = &sync.RWMutex{}
	}

	p.mu.Lock()
}

// unlock releases the printer's runtime mutex.
func (p *ProgressbarPrinter) unlock() {
	p.mu.Unlock()
}

// rlock takes a read lock on the printer's runtime mutex, allocating it on
// first use.
func (p *ProgressbarPrinter) rlock() {
	if p.mu == nil {
		p.mu = &sync.RWMutex{}
	}

	p.mu.RLock()
}

// runlock releases a read lock taken by rlock.
func (p *ProgressbarPrinter) runlock() {
	p.mu.RUnlock()
}

// isActive returns whether the printer is currently active. Used internally by
// print.go to inspect bars from another goroutine without racing with Start/Stop.
func (p *ProgressbarPrinter) isActive() bool {
	p.rlock()
	defer p.runlock()

	return p.IsActive
}

// title returns the current title under the runtime lock.
func (p *ProgressbarPrinter) title() string {
	p.rlock()
	defer p.runlock()

	return p.Title
}

// writer returns the configured writer under the runtime lock.
func (p *ProgressbarPrinter) writer() io.Writer {
	p.rlock()
	defer p.runlock()

	return p.Writer
}

// WithTitle sets the name of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithTitle(name string) *ProgressbarPrinter {
	p.Title = name
	return &p
}

// WithMaxWidth sets the maximum width of the ProgressbarPrinter.
// If the terminal is smaller than the given width, the terminal width will be used instead.
// If the width is set to zero, or below, the terminal width will be used.
func (p ProgressbarPrinter) WithMaxWidth(maxWidth int) *ProgressbarPrinter {
	p.MaxWidth = maxWidth
	return &p
}

// WithTotal sets the total value of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithTotal(total int) *ProgressbarPrinter {
	p.Total = total
	return &p
}

// WithCurrent sets the current value of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithCurrent(current int) *ProgressbarPrinter {
	p.Current = current
	return &p
}

// WithBarCharacter sets the bar character of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithBarCharacter(char string) *ProgressbarPrinter {
	p.BarCharacter = char
	return &p
}

// WithLastCharacter sets the last character of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithLastCharacter(char string) *ProgressbarPrinter {
	p.LastCharacter = char
	return &p
}

// WithElapsedTimeRoundingFactor sets the rounding factor of the elapsed time.
func (p ProgressbarPrinter) WithElapsedTimeRoundingFactor(duration time.Duration) *ProgressbarPrinter {
	p.ElapsedTimeRoundingFactor = duration
	return &p
}

// WithShowElapsedTime sets if the elapsed time should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowElapsedTime(b ...bool) *ProgressbarPrinter {
	p.ShowElapsedTime = internal.WithBoolean(b)
	return &p
}

// WithShowCount sets if the total and current count should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowCount(b ...bool) *ProgressbarPrinter {
	p.ShowCount = internal.WithBoolean(b)
	return &p
}

// WithShowTitle sets if the title should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowTitle(b ...bool) *ProgressbarPrinter {
	p.ShowTitle = internal.WithBoolean(b)
	return &p
}

// WithShowPercentage sets if the completed percentage should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowPercentage(b ...bool) *ProgressbarPrinter {
	p.ShowPercentage = internal.WithBoolean(b)
	return &p
}

// WithStartedAt sets the time when the ProgressbarPrinter started.
func (p ProgressbarPrinter) WithStartedAt(t time.Time) *ProgressbarPrinter {
	p.startedAt = t
	return &p
}

// WithTitleStyle sets the style of the title.
func (p ProgressbarPrinter) WithTitleStyle(style *Style) *ProgressbarPrinter {
	p.TitleStyle = style
	return &p
}

// WithBarStyle sets the style of the bar.
func (p ProgressbarPrinter) WithBarStyle(style *Style) *ProgressbarPrinter {
	p.BarStyle = style
	return &p
}

// WithRemoveWhenDone sets if the ProgressbarPrinter should be removed when it is done.
func (p ProgressbarPrinter) WithRemoveWhenDone(b ...bool) *ProgressbarPrinter {
	p.RemoveWhenDone = internal.WithBoolean(b)
	return &p
}

// WithBarFiller sets the filler character for the ProgressbarPrinter.
func (p ProgressbarPrinter) WithBarFiller(char string) *ProgressbarPrinter {
	p.BarFiller = char
	return &p
}

// WithBarFillerStyle sets the style of the unfilled portion of the bar.
func (p ProgressbarPrinter) WithBarFillerStyle(style *Style) *ProgressbarPrinter {
	p.BarFillerStyle = style
	return &p
}

// WithBarPartialCharacters sets the glyphs used to render the leading edge of
// the bar at sub-character resolution, ordered from least to most filled (e.g.
// "▏" through "▉"). Passing an empty slice disables smooth rendering and falls
// back to whole-cell rendering with BarCharacter and LastCharacter.
func (p ProgressbarPrinter) WithBarPartialCharacters(chars []string) *ProgressbarPrinter {
	p.BarPartialCharacters = chars
	return &p
}

// WithWriter sets the custom Writer.
func (p ProgressbarPrinter) WithWriter(writer io.Writer) *ProgressbarPrinter {
	p.Writer = writer
	return &p
}

// SetWriter sets the custom Writer.
func (p *ProgressbarPrinter) SetWriter(writer io.Writer) {
	p.Writer = writer
}

// SetStartedAt sets the time when the ProgressbarPrinter started.
func (p *ProgressbarPrinter) SetStartedAt(t time.Time) {
	p.lock()
	defer p.unlock()

	p.startedAt = t
}

// ResetTimer resets the timer of the ProgressbarPrinter.
func (p *ProgressbarPrinter) ResetTimer() {
	p.lock()
	defer p.unlock()

	p.startedAt = time.Now()
}

// Increment current value by one.
func (p *ProgressbarPrinter) Increment() *ProgressbarPrinter {
	p.Add(1)
	return p
}

// UpdateTitle updates the title and re-renders the progressbar
func (p *ProgressbarPrinter) UpdateTitle(title string) *ProgressbarPrinter {
	p.lock()
	p.Title = title
	p.unlock()
	p.updateProgress()

	return p
}

// updateProgress renders the progressbar to the configured writer. It locks
// internally via getString and reads the writer under the same lock.
func (p *ProgressbarPrinter) updateProgress() *ProgressbarPrinter {
	rendered, writer := p.renderAndWriter()
	if rendered == "" {
		return p
	}

	Fprinto(writer, rendered)

	return p
}

// renderAndWriter renders the bar and returns it along with the writer to
// emit to. Both reads happen under a single lock acquisition.
func (p *ProgressbarPrinter) renderAndWriter() (string, io.Writer) {
	p.lock()
	defer p.unlock()

	return p.getStringLocked(), p.Writer
}

// getStringLocked renders the progress bar. The caller must hold p.mu.
func (p *ProgressbarPrinter) getStringLocked() string {
	if !p.IsActive || p.Total == 0 {
		return ""
	}

	if p.TitleStyle == nil {
		p.TitleStyle = NewStyle()
	}

	if p.BarStyle == nil {
		p.BarStyle = NewStyle()
	}

	before := p.decoratorsBeforeLocked()
	after := p.decoratorsAfterLocked()

	// The bar fills whatever horizontal space the decorators leave behind.
	barWidth := p.lineWidthLocked() - internal.GetStringMaxWidth(before) - internal.GetStringMaxWidth(after)

	return before + p.renderBar(barWidth) + after
}

// lineWidthLocked returns the total width the rendered line may occupy, clamped
// to the terminal. The caller must hold p.mu.
func (p *ProgressbarPrinter) lineWidthLocked() int {
	terminalWidth := GetTerminalWidth()
	if p.MaxWidth <= 0 || terminalWidth < p.MaxWidth {
		return terminalWidth
	}

	return p.MaxWidth
}

// decoratorsBeforeLocked builds the prefix shown left of the bar (title and
// count). The caller must hold p.mu.
func (p *ProgressbarPrinter) decoratorsBeforeLocked() string {
	var b strings.Builder

	if p.ShowTitle {
		b.WriteString(p.TitleStyle.Sprint(p.Title) + " ")
	}

	if p.ShowCount {
		// Space-pad the current value to the width of the total so the line
		// keeps a stable layout while the numbers grow.
		padding := 1 + int(math.Log10(float64(p.Total)))
		b.WriteString(LightWhite(fmt.Sprintf("%*d", padding, p.Current)) + Gray(fmt.Sprintf("/%d", p.Total)) + " ")
	}

	return b.String()
}

// decoratorsAfterLocked builds the suffix shown right of the bar (percentage
// and elapsed time). The caller must hold p.mu.
func (p *ProgressbarPrinter) decoratorsAfterLocked() string {
	var b strings.Builder

	b.WriteString(" ")

	if p.ShowPercentage {
		percentage := int(internal.PercentageRound(float64(int64(p.Total)), float64(int64(p.Current))))

		text := fmt.Sprintf("%3d%%", percentage)
		if p.Current >= p.Total {
			text = Green(text)
		} else {
			text = LightWhite(text)
		}

		b.WriteString(text + " ")
	}

	if p.ShowElapsedTime {
		b.WriteString(Gray("· " + p.parseElapsedTimeLocked()))
	}

	return b.String()
}

// renderBar draws the filled/unfilled bar to exactly width display cells.
// The default block bar is drawn with sub-character resolution for a smooth
// edge; a custom BarCharacter falls back to whole-cell rendering.
func (p *ProgressbarPrinter) renderBar(width int) string {
	if width <= 0 {
		return ""
	}

	filler := p.BarFiller
	if filler == "" {
		filler = " "
	}

	ratio := float64(p.Current) / float64(p.Total)
	switch {
	case ratio > 1:
		ratio = 1
	case ratio < 0:
		ratio = 0
	}

	if len(p.BarPartialCharacters) > 0 {
		return p.renderSmoothBar(ratio, width, filler)
	}

	full := int(ratio * float64(width))

	var bar string
	if full > 0 {
		bar = p.BarStyle.Sprint(strings.Repeat(p.BarCharacter, full-1) + p.LastCharacter)
	}

	if full < width {
		bar += p.sprintFiller(filler, width-full)
	}

	return bar
}

// sprintFiller renders n cells of the unfilled bar track. Blank fillers are
// returned unstyled to keep the output free of redundant escape codes.
func (p *ProgressbarPrinter) sprintFiller(filler string, n int) string {
	s := strings.Repeat(filler, n)
	if strings.TrimSpace(s) == "" {
		return s
	}

	style := p.BarFillerStyle
	if style == nil {
		style = &ThemeDefault.ProgressbarFillerStyle
	}

	return style.Sprint(s)
}

// renderSmoothBar draws the bar using a partial-block glyph for the leading
// edge so the fill grows a fraction of a cell at a time.
func (p *ProgressbarPrinter) renderSmoothBar(ratio float64, width int, filler string) string {
	exact := ratio * float64(width)
	full := int(exact)

	// Map the fractional remainder to a partial glyph. Slot 0 means the fill
	// lands exactly on a cell boundary, so no partial edge is drawn.
	var partial string
	if slot := int((exact - float64(full)) * float64(len(p.BarPartialCharacters)+1)); slot > 0 {
		partial = p.BarPartialCharacters[slot-1]
	}

	cells := full
	if partial != "" {
		cells++
	}

	bar := p.BarStyle.Sprint(strings.Repeat(p.BarCharacter, full) + partial)
	if cells < width {
		bar += p.sprintFiller(filler, width-cells)
	}

	return bar
}

// Add to current value.
func (p *ProgressbarPrinter) Add(count int) *ProgressbarPrinter {
	p.lock()

	if p.Total == 0 {
		p.unlock()
		return nil
	}

	p.Current += count

	reachedEnd := p.Current >= p.Total
	if reachedEnd {
		p.Total = p.Current
	}

	p.unlock()

	p.updateProgress()

	if reachedEnd {
		p.updateProgress()
		_, _ = p.Stop()
	}

	return p
}

// Start the ProgressbarPrinter.
func (p ProgressbarPrinter) Start(title ...any) (*ProgressbarPrinter, error) {
	cursor.Hide()

	// Allocate the runtime mutex eagerly so the rerender goroutine and any
	// subsequent user calls share the same lock without racing on lazy init.
	p.mu = &sync.RWMutex{}

	if rawOutput() && p.ShowTitle {
		Fprintln(p.Writer, p.Title)
	}

	p.IsActive = true
	if len(title) != 0 {
		p.Title = Sprint(title...)
	}

	registerProgressBar(&p)
	p.startedAt = time.Now()

	p.updateProgress()

	if p.ShowElapsedTime {
		p.rerenderTask = schedule.Every(time.Second, func() bool {
			p.updateProgress()
			return true
		})
	}

	return &p, nil
}

// Stop the ProgressbarPrinter.
func (p *ProgressbarPrinter) Stop() (*ProgressbarPrinter, error) {
	p.lock()
	task := p.rerenderTask
	wasActive := p.IsActive
	p.IsActive = false
	removeWhenDone := p.RemoveWhenDone
	writer := p.Writer
	p.unlock()

	// Stop the rerender task outside the lock: schedule.Task.Stop blocks until
	// the in-flight callback returns, and that callback also takes p.mu.
	if task != nil && task.IsActive() {
		task.Stop()
	}

	cursor.Show()

	if !wasActive {
		return p, nil
	}

	if removeWhenDone {
		fClearLine(writer)
		Fprinto(writer)
	} else {
		Fprintln(writer)
	}

	return p, nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (p *ProgressbarPrinter) GenericStart() (*LivePrinter, error) {
	p2, _ := p.Start()
	lp := LivePrinter(p2)

	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (p *ProgressbarPrinter) GenericStop() (*LivePrinter, error) {
	p2, _ := p.Stop()
	lp := LivePrinter(p2)

	return &lp, nil
}

// GetElapsedTime returns the elapsed time, since the ProgressbarPrinter was started.
func (p *ProgressbarPrinter) GetElapsedTime() time.Duration {
	p.rlock()
	defer p.runlock()

	return time.Since(p.startedAt)
}

// parseElapsedTimeLocked formats the elapsed time. The caller must hold p.mu.
func (p *ProgressbarPrinter) parseElapsedTimeLocked() string {
	elapsed := time.Since(p.startedAt)
	// time.Duration.Round panics if the rounding factor is <= 0.
	// Guard against invalid values by skipping rounding in this case.
	if p.ElapsedTimeRoundingFactor <= 0 {
		return elapsed.String()
	}

	return elapsed.Round(p.ElapsedTimeRoundingFactor).String()
}
