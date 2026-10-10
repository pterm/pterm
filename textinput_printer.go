package pterm

import (
	"context"
	"io"
	"strconv"
	"strings"
	"unicode"

	"atomicgo.dev/color"
	"atomicgo.dev/keyboard/keys"
	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/text"
)

// TextInputPrinter asks the user to type a line of text, or several:
//
//	? Project name my-app█
//
// The input supports the usual editing keys: the arrow keys, home and end,
// backspace and delete, ctrl+w to delete a word, ctrl+u and ctrl+k to
// delete to the start or end of the line. Pasted text is inserted as a
// whole. In multi-line mode enter starts a new line and SubmitKey (ctrl+d)
// submits.
type TextInputPrinter struct {
	// Text is the question.
	Text string
	// DefaultValue is returned when the user submits an empty input, and
	// without asking when standard input is not a terminal. Tab copies it
	// into the input for editing.
	DefaultValue string
	// Placeholder is shown muted while the input is empty. Without one,
	// the default value is shown.
	Placeholder string
	// Mask replaces every typed character, for passwords.
	Mask string
	// MultiLine makes enter start a new line instead of submitting.
	MultiLine bool
	// SubmitKey submits the input. It is how a multi-line input ends.
	SubmitKey keys.KeyCode
	// KeyHints shows the key hint footer in multi-line mode.
	KeyHints bool
	// Styles style the input.
	Styles TextInputStyles
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Input is read for key presses instead of the terminal.
	Input io.Reader
	// Context cancels the prompt when it ends.
	Context context.Context
	// OnInterrupt runs on Ctrl+C instead of exiting the program.
	OnInterrupt func()
}

// TextInputStyles style a TextInput.
type TextInputStyles struct {
	// Input styles the typed text.
	Input Style
	// Placeholder styles the placeholder and the default value.
	Placeholder Style
	// Cursor styles the cell under the cursor.
	Cursor Style
}

// TextInput is the default text prompt.
var TextInput = TextInputPrinter{
	Text:      "Enter text",
	SubmitKey: keys.CtrlD,
	KeyHints:  true,
	Styles: TextInputStyles{
		Input:       Style{Fg: ToneText},
		Placeholder: Style{Fg: ToneMuted},
		Cursor:      Style{Attrs: color.Reverse},
	},
}

// WithText returns a copy of p that asks t.
func (p TextInputPrinter) WithText(t string) *TextInputPrinter {
	p.Text = t
	return &p
}

// WithDefaultValue returns a copy of p that returns v for an empty input.
func (p TextInputPrinter) WithDefaultValue(v string) *TextInputPrinter {
	p.DefaultValue = v
	return &p
}

// WithPlaceholder returns a copy of p that shows s while the input is
// empty.
func (p TextInputPrinter) WithPlaceholder(s string) *TextInputPrinter {
	p.Placeholder = s
	return &p
}

// WithMask returns a copy of p that shows mask in place of every typed
// character.
func (p TextInputPrinter) WithMask(mask string) *TextInputPrinter {
	p.Mask = mask
	return &p
}

// WithMultiLine returns a copy of p that does or does not accept several
// lines.
func (p TextInputPrinter) WithMultiLine(b ...bool) *TextInputPrinter {
	p.MultiLine = flag(b)
	return &p
}

// WithSubmitKey returns a copy of p that submits with k.
func (p TextInputPrinter) WithSubmitKey(k keys.KeyCode) *TextInputPrinter {
	p.SubmitKey = k
	return &p
}

// WithKeyHints returns a copy of p that does or does not show the key
// hint footer.
func (p TextInputPrinter) WithKeyHints(b ...bool) *TextInputPrinter {
	p.KeyHints = flag(b)
	return &p
}

// WithStyles returns a copy of p with the styles st.
func (p TextInputPrinter) WithStyles(st TextInputStyles) *TextInputPrinter {
	p.Styles = st
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p TextInputPrinter) WithWriter(w io.Writer) *TextInputPrinter {
	p.Writer = w
	return &p
}

// WithInput returns a copy of p that reads key presses from r, for tests
// and scripted input.
func (p TextInputPrinter) WithInput(r io.Reader) *TextInputPrinter {
	p.Input = r
	return &p
}

// WithContext returns a copy of p that gives up with ErrCanceled when ctx
// ends.
func (p TextInputPrinter) WithContext(ctx context.Context) *TextInputPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting. Show then returns ErrInterrupted.
func (p TextInputPrinter) WithOnInterrupt(fn func()) *TextInputPrinter {
	p.OnInterrupt = fn
	return &p
}

// Show asks the question, or text when given, and returns what the user
// typed. Lines of a multi-line input are joined with "\n".
func (p TextInputPrinter) Show(text ...string) (string, error) {
	if len(text) > 0 {
		p.Text = strings.Join(text, " ")
	}

	e := &editor{lines: [][]rune{nil}}

	var prev keys.KeyCode

	pr := prompt[string]{
		cfg:  promptConfig{writer: p.Writer, input: p.Input, ctx: p.Context, onInterrupt: p.OnInterrupt},
		view: func(c canvas) string { return p.view(c, e) },
		key: func(k keys.Key) bool {
			// A pasted "\r\n" arrives as enter and ctrl+j: one line break.
			crlf := k.Code == keys.CtrlJ && prev == keys.Enter
			prev = k.Code

			if crlf {
				return false
			}

			return p.key(e, k)
		},
		result:   func() string { return p.value(e) },
		collapse: func(c canvas, st status) string { return p.collapse(c, st, e) },
	}

	return pr.run(p.DefaultValue, p.DefaultValue != "")
}

func (p TextInputPrinter) value(e *editor) string {
	if e.empty() {
		return p.DefaultValue
	}

	return e.String()
}

// key handles one key press and reports whether the input is submitted.
func (p TextInputPrinter) key(e *editor, k keys.Key) bool {
	if p.SubmitKey != 0 && k.Code == p.SubmitKey {
		return true
	}

	switch k.Code {
	case keys.Enter, keys.CtrlJ:
		if !p.MultiLine {
			return true
		}

		e.newline()

	case keys.RuneKey, keys.Space:
		if !k.AltPressed {
			e.insert(k.Runes)
		}

	case keys.Tab:
		if e.empty() && p.DefaultValue != "" && p.Mask == "" {
			e.insert([]rune(p.DefaultValue))
		}

	case keys.Backspace, keys.CtrlH:
		if k.AltPressed {
			e.deleteWord()
		} else {
			e.backspace()
		}

	case keys.CtrlW:
		e.deleteWord()
	case keys.Delete:
		e.del()
	case keys.CtrlU:
		e.killStart()
	case keys.CtrlK:
		e.killEnd()
	default:
		p.motion(e, k)
	}

	return false
}

// motion handles the keys that move the cursor.
func (p TextInputPrinter) motion(e *editor, k keys.Key) {
	switch k.Code {
	case keys.Left, keys.CtrlB:
		if k.AltPressed {
			e.wordLeft()
		} else {
			e.left()
		}

	case keys.Right, keys.CtrlF:
		if k.AltPressed {
			e.wordRight()
		} else {
			e.right()
		}

	case keys.CtrlLeft:
		e.wordLeft()
	case keys.CtrlRight:
		e.wordRight()
	case keys.Up:
		e.vertical(-1)
	case keys.Down:
		e.vertical(1)
	case keys.Home, keys.CtrlA:
		e.col = 0
	case keys.End, keys.CtrlE:
		e.col = len(e.lines[e.row])
	}
}

// cell is one rune of the input as it is drawn.
type cell struct {
	s     string
	w     int
	space bool // a space the row may break after
}

func (p TextInputPrinter) view(c canvas, e *editor) string {
	head := question(c, p.Text, "")
	if p.MultiLine {
		lines := []string{head}

		for row := range e.lines {
			for _, r := range p.layout(c, p.cells(c, e, row), c.width-1-len(promptIndent), c.width-1-len(promptIndent)) {
				if r == "" {
					lines = append(lines, "")
					continue
				}

				lines = append(lines, promptIndent+r)
			}
		}

		if p.KeyHints {
			lines = append(lines, "", footer(c,
				keyHint{"enter", "new line", 1},
				keyHint{p.SubmitKey.String(), "submit", 0},
			))
		}

		return strings.Join(lines, "\n")
	}

	head += " "
	first := c.width - 1 - text.Width(screen.Strip(head))
	rows := p.layout(c, p.cells(c, e, 0), first, c.width-1-len(promptIndent))

	rows[0] = head + rows[0]
	for i := 1; i < len(rows); i++ {
		rows[i] = promptIndent + rows[i]
	}

	return strings.Join(rows, "\n")
}

// cells returns the drawn runes of one line of the input, the cursor
// included.
func (p TextInputPrinter) cells(c canvas, e *editor, row int) []cell {
	line := e.lines[row]

	if e.empty() {
		hint := p.Placeholder
		if hint == "" && p.Mask == "" {
			hint = p.DefaultValue
		}

		cells := []cell{p.cursor(c, " ")}

		for i, r := range []rune(hint) {
			s := c.style(p.Styles.Placeholder, string(r))
			if i == 0 {
				cells[0] = p.cursor(c, string(r))
				continue
			}

			cells = append(cells, cell{s: s, w: runewidth.RuneWidth(r), space: r == ' '})
		}

		return cells
	}

	cells := make([]cell, 0, len(line)+1)

	for i, r := range line {
		s := string(r)
		if p.Mask != "" {
			s = p.Mask
		}

		space := p.Mask == "" && r == ' '

		if row == e.row && i == e.col {
			cl := p.cursor(c, s)
			cl.space = space
			cells = append(cells, cl)

			continue
		}

		cells = append(cells, cell{s: c.style(p.Styles.Input, s), w: text.Width(s), space: space})
	}

	if row == e.row && e.col == len(line) {
		cells = append(cells, p.cursor(c, " "))
	}

	return cells
}

// cursor draws s as the cursor. Without colors the cursor still needs to
// be seen, so it falls back to plain reverse video.
func (p TextInputPrinter) cursor(c canvas, s string) cell {
	styled := c.style(p.Styles.Cursor, s)
	if styled == s {
		styled = "\x1b[7m" + s + "\x1b[27m"
	}

	return cell{s: styled, w: max(text.Width(s), 1)}
}

// layout breaks cells into rows: the first row holds first columns, the
// others rest columns. Rows break after a space where possible, like
// wrapped text. With an unknown width everything stays on one row.
func (p TextInputPrinter) layout(c canvas, cells []cell, first, rest int) []string {
	if c.width <= 0 {
		return []string{joinCells(cells)}
	}

	var (
		rows  []string
		row   []cell
		used  int
		avail = max(first, 1)
	)

	for _, cl := range cells {
		if used+cl.w > avail && len(row) > 0 {
			cut := len(row)

			for i := len(row) - 1; i > 0; i-- {
				if row[i].space {
					cut = i + 1
					break
				}
			}

			rows = append(rows, joinCells(row[:cut]))
			row = append([]cell(nil), row[cut:]...)
			used = 0

			for _, carried := range row {
				used += carried.w
			}

			avail = max(rest, 1)
		}

		row = append(row, cl)
		used += cl.w
	}

	return append(rows, joinCells(row))
}

func joinCells(cells []cell) string {
	var b strings.Builder
	for _, cl := range cells {
		b.WriteString(cl.s)
	}

	return b.String()
}

func (p TextInputPrinter) collapse(c canvas, st status, e *editor) string {
	value := p.value(e)
	first, rest, multi := strings.Cut(value, "\n")

	answer := first
	if p.Mask != "" && value != "" {
		// A fixed number of masks, so the answer does not tell the length.
		answer = strings.Repeat(p.Mask, 8)
	}

	if c.width > 0 {
		answer = c.cut(answer, c.width-1-text.Width(p.Text)-4)
	}

	line := answered(c, st, p.Text, answer)
	if multi && st != statusCanceled && p.Mask == "" {
		line += " " + c.paint(ToneMuted, "+"+strconv.Itoa(strings.Count(rest, "\n")+1)+" lines")
	}

	return line
}

// editor holds the text being typed and the cursor in it.
type editor struct {
	lines    [][]rune
	row, col int
}

func (e *editor) empty() bool {
	return len(e.lines) == 1 && len(e.lines[0]) == 0
}

func (e *editor) String() string {
	parts := make([]string, len(e.lines))
	for i, l := range e.lines {
		parts[i] = string(l)
	}

	return strings.Join(parts, "\n")
}

// insert types rs at the cursor. Control characters are dropped, so a
// paste cannot smuggle escape sequences into the view.
func (e *editor) insert(rs []rune) {
	clean := make([]rune, 0, len(rs))

	for _, r := range rs {
		if r == '\t' {
			r = ' '
		}

		if !unicode.IsControl(r) {
			clean = append(clean, r)
		}
	}

	line := e.lines[e.row]

	next := make([]rune, 0, len(line)+len(clean))
	next = append(next, line[:e.col]...)
	next = append(next, clean...)
	next = append(next, line[e.col:]...)

	e.lines[e.row] = next
	e.col += len(clean)
}

func (e *editor) newline() {
	line := e.lines[e.row]
	head := append([]rune(nil), line[:e.col]...)
	tail := append([]rune(nil), line[e.col:]...)

	lines := make([][]rune, 0, len(e.lines)+1)
	lines = append(lines, e.lines[:e.row]...)
	lines = append(lines, head, tail)
	lines = append(lines, e.lines[e.row+1:]...)

	e.lines = lines
	e.row++
	e.col = 0
}

func (e *editor) backspace() {
	switch {
	case e.col > 0:
		line := e.lines[e.row]
		e.lines[e.row] = append(line[:e.col-1:e.col-1], line[e.col:]...)
		e.col--

	case e.row > 0:
		e.row--
		e.col = len(e.lines[e.row])
		e.join()
	}
}

func (e *editor) del() {
	line := e.lines[e.row]

	switch {
	case e.col < len(line):
		e.lines[e.row] = append(line[:e.col:e.col], line[e.col+1:]...)
	case e.row < len(e.lines)-1:
		e.join()
	}
}

// join appends the next line to the line of the cursor.
func (e *editor) join() {
	line := e.lines[e.row]
	e.lines[e.row] = append(line[:len(line):len(line)], e.lines[e.row+1]...)
	e.lines = append(e.lines[:e.row+1], e.lines[e.row+2:]...)
}

func (e *editor) left() {
	switch {
	case e.col > 0:
		e.col--
	case e.row > 0:
		e.row--
		e.col = len(e.lines[e.row])
	}
}

func (e *editor) right() {
	switch {
	case e.col < len(e.lines[e.row]):
		e.col++
	case e.row < len(e.lines)-1:
		e.row++
		e.col = 0
	}
}

func (e *editor) vertical(d int) {
	row := e.row + d
	if row < 0 || row >= len(e.lines) {
		return
	}

	e.row = row
	e.col = min(e.col, len(e.lines[row]))
}

// wordStart returns the column where the word before the cursor starts.
func (e *editor) wordStart() int {
	line, i := e.lines[e.row], e.col

	for i > 0 && unicode.IsSpace(line[i-1]) {
		i--
	}

	for i > 0 && !unicode.IsSpace(line[i-1]) {
		i--
	}

	return i
}

func (e *editor) deleteWord() {
	if e.col == 0 {
		e.backspace()
		return
	}

	start := e.wordStart()
	line := e.lines[e.row]
	e.lines[e.row] = append(line[:start:start], line[e.col:]...)
	e.col = start
}

func (e *editor) wordLeft() {
	if e.col == 0 {
		e.left()
		return
	}

	e.col = e.wordStart()
}

func (e *editor) wordRight() {
	line := e.lines[e.row]
	if e.col == len(line) {
		e.right()
		return
	}

	i := e.col
	for i < len(line) && unicode.IsSpace(line[i]) {
		i++
	}

	for i < len(line) && !unicode.IsSpace(line[i]) {
		i++
	}

	e.col = i
}

func (e *editor) killStart() {
	e.lines[e.row] = append([]rune(nil), e.lines[e.row][e.col:]...)
	e.col = 0
}

func (e *editor) killEnd() {
	e.lines[e.row] = e.lines[e.row][:e.col:e.col]
}
