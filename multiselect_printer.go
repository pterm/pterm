package pterm

import (
	"context"
	"io"
	"strconv"
	"strings"
	"unicode"

	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm/internal/text"
)

// MultiselectPrinter asks the user to pick any number of options:
//
//	? What should be cleaned? 3 of 5
//
//	  PACKAGE MANAGERS
//	  ❯ ◉ npm cache     ~/.npm                1.2 GB
//	    ○ yarn cache    ~/.cache/yarn       640.0 MB
//
//	  [space] toggle · [a] all · [n] none · [↑↓] move · [enter] confirm
//
// Options can have groups, descriptions and right-aligned details.
// Disabled options are muted and keep their state. With Filter on, "/"
// starts a fuzzy filter.
type MultiselectPrinter struct {
	// Text is the question.
	Text string
	// Options are the choices. WithOptions fills it from plain labels.
	Options []Option
	// DefaultOptions are the labels or values of the options that start
	// checked, in addition to options with Selected set. They are returned
	// without asking when standard input is not a terminal.
	DefaultOptions []string
	// MaxHeight is the number of options shown at once; the rest scrolls.
	// Zero or less shows every option.
	MaxHeight int
	// Filter lets the user narrow the list: "/" starts a filter, escape
	// clears it.
	Filter bool
	// FilterPlaceholder is shown muted while the filter is empty.
	FilterPlaceholder string
	// Pointer marks the active option. Empty uses the pointer glyph.
	Pointer string
	// Checked and Unchecked are the checkboxes. Empty uses the glyphs.
	Checked   string
	Unchecked string
	// Styles style the rows.
	Styles OptionStyles
	// KeyHints shows the key hint footer.
	KeyHints bool
	// ShowSelectedOptions lists the chosen options under the answered
	// prompt.
	ShowSelectedOptions bool
	// KeySelect toggles the active option, KeyConfirm finishes.
	KeySelect  keys.KeyCode
	KeyConfirm keys.KeyCode
	// KeyAll checks and KeyNone unchecks every option the filter shows.
	// The right and left arrow keys do the same.
	KeyAll  rune
	KeyNone rune
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Input is read for key presses instead of the terminal.
	Input io.Reader
	// Context cancels the prompt when it ends.
	Context context.Context
	// OnInterrupt runs on Ctrl+C instead of exiting the program.
	OnInterrupt func()
}

// Multiselect is the default multiple choice prompt.
var Multiselect = MultiselectPrinter{
	Text:              "Select options",
	MaxHeight:         8,
	Filter:            true,
	FilterPlaceholder: "type to filter",
	Styles:            DefaultOptionStyles,
	KeyHints:          true,
	KeySelect:         keys.Space,
	KeyConfirm:        keys.Enter,
	KeyAll:            'a',
	KeyNone:           'n',
}

// WithText returns a copy of p that asks t.
func (p MultiselectPrinter) WithText(t string) *MultiselectPrinter {
	p.Text = t
	return &p
}

// WithOptions returns a copy of p that offers the given labels.
func (p MultiselectPrinter) WithOptions(labels []string) *MultiselectPrinter {
	p.Options = optionsOf(labels)
	return &p
}

// WithItems returns a copy of p that offers options with descriptions,
// details, groups or disabled entries.
func (p MultiselectPrinter) WithItems(options []Option) *MultiselectPrinter {
	p.Options = options
	return &p
}

// WithDefaultOptions returns a copy of p whose options with the given
// labels or values start checked.
func (p MultiselectPrinter) WithDefaultOptions(options []string) *MultiselectPrinter {
	p.DefaultOptions = options
	return &p
}

// WithMaxHeight returns a copy of p that shows n options at once.
func (p MultiselectPrinter) WithMaxHeight(n int) *MultiselectPrinter {
	p.MaxHeight = n
	return &p
}

// WithFilter returns a copy of p that does or does not offer a filter.
func (p MultiselectPrinter) WithFilter(b ...bool) *MultiselectPrinter {
	p.Filter = flag(b)
	return &p
}

// WithFilterPlaceholder returns a copy of p that shows s while the filter
// is empty.
func (p MultiselectPrinter) WithFilterPlaceholder(s string) *MultiselectPrinter {
	p.FilterPlaceholder = s
	return &p
}

// WithPointer returns a copy of p that marks the active option with s.
func (p MultiselectPrinter) WithPointer(s string) *MultiselectPrinter {
	p.Pointer = s
	return &p
}

// WithCheckmark returns a copy of p that draws checked options with
// checked and the others with unchecked.
func (p MultiselectPrinter) WithCheckmark(checked, unchecked string) *MultiselectPrinter {
	p.Checked, p.Unchecked = checked, unchecked
	return &p
}

// WithStyles returns a copy of p with the row styles st.
func (p MultiselectPrinter) WithStyles(st OptionStyles) *MultiselectPrinter {
	p.Styles = st
	return &p
}

// WithKeyHints returns a copy of p that does or does not show the key
// hint footer.
func (p MultiselectPrinter) WithKeyHints(b ...bool) *MultiselectPrinter {
	p.KeyHints = flag(b)
	return &p
}

// WithShowSelectedOptions returns a copy of p that lists the chosen
// options under the answered prompt.
func (p MultiselectPrinter) WithShowSelectedOptions(b ...bool) *MultiselectPrinter {
	p.ShowSelectedOptions = flag(b)
	return &p
}

// WithKeySelect returns a copy of p that toggles options with k.
func (p MultiselectPrinter) WithKeySelect(k keys.KeyCode) *MultiselectPrinter {
	p.KeySelect = k
	return &p
}

// WithKeyConfirm returns a copy of p that finishes with k.
func (p MultiselectPrinter) WithKeyConfirm(k keys.KeyCode) *MultiselectPrinter {
	p.KeyConfirm = k
	return &p
}

// WithKeyAll returns a copy of p that checks every shown option with r.
func (p MultiselectPrinter) WithKeyAll(r rune) *MultiselectPrinter {
	p.KeyAll = r
	return &p
}

// WithKeyNone returns a copy of p that unchecks every shown option with r.
func (p MultiselectPrinter) WithKeyNone(r rune) *MultiselectPrinter {
	p.KeyNone = r
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p MultiselectPrinter) WithWriter(w io.Writer) *MultiselectPrinter {
	p.Writer = w
	return &p
}

// WithInput returns a copy of p that reads key presses from r, for tests
// and scripted input.
func (p MultiselectPrinter) WithInput(r io.Reader) *MultiselectPrinter {
	p.Input = r
	return &p
}

// WithContext returns a copy of p that gives up with ErrCanceled when ctx
// ends.
func (p MultiselectPrinter) WithContext(ctx context.Context) *MultiselectPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting. Show then returns ErrInterrupted.
func (p MultiselectPrinter) WithOnInterrupt(fn func()) *MultiselectPrinter {
	p.OnInterrupt = fn
	return &p
}

// multiselectState is what the user changed so far.
type multiselectState struct {
	list      *optionList
	checked   []bool
	filtering bool
	query     []rune
}

func (s *multiselectState) count() int {
	n := 0

	for _, c := range s.checked {
		if c {
			n++
		}
	}

	return n
}

// setShown checks or unchecks every enabled option the filter shows.
func (s *multiselectState) setShown(v bool) {
	for _, m := range s.list.matches {
		if !s.list.items[m.index].Disabled {
			s.checked[m.index] = v
		}
	}
}

// Show asks the question, or text when given, and returns the values of
// the checked options in the order of the options. It returns ErrEmpty
// when there are no options.
func (p MultiselectPrinter) Show(text ...string) ([]string, error) {
	if len(text) > 0 {
		p.Text = strings.Join(text, " ")
	}

	if len(p.Options) == 0 {
		return nil, ErrEmpty
	}

	s := &multiselectState{
		list:    newOptionList(p.Options, p.MaxHeight, false),
		checked: make([]bool, len(p.Options)),
	}

	hasDefault := false

	for i, o := range p.Options {
		if o.Selected || p.isDefault(o) {
			s.checked[i] = true
			hasDefault = true
		}
	}

	pr := prompt[[]string]{
		cfg:  promptConfig{writer: p.Writer, input: p.Input, ctx: p.Context, onInterrupt: p.OnInterrupt},
		view: func(c canvas) string { return p.view(c, s) },
		key:  func(k keys.Key) bool { return p.key(s, k) },
		result: func() []string {
			return p.values(s.checked)
		},
		collapse: func(c canvas, st status) string { return p.collapse(c, st, s) },
	}

	return pr.run(p.values(s.checked), hasDefault)
}

func (p MultiselectPrinter) isDefault(o Option) bool {
	for _, d := range p.DefaultOptions {
		if o.is(d) {
			return true
		}
	}

	return false
}

func (p MultiselectPrinter) values(checked []bool) []string {
	values := []string{}

	for i, c := range checked {
		if c {
			values = append(values, p.Options[i].value())
		}
	}

	return values
}

// key handles one key press and reports whether the prompt is done.
func (p MultiselectPrinter) key(s *multiselectState, k keys.Key) bool {
	switch k.Code {
	case p.KeyConfirm:
		return true
	case p.KeySelect:
		if _, i, ok := s.list.current(); ok && !p.Options[i].Disabled {
			s.checked[i] = !s.checked[i]
		}

		return false
	}

	switch k.Code {
	case keys.Up, keys.CtrlP, keys.ShiftTab:
		s.list.move(-1)
	case keys.Down, keys.CtrlN, keys.Tab:
		s.list.move(1)
	case keys.PgUp:
		s.list.page(-1)
	case keys.PgDown:
		s.list.page(1)
	case keys.Home:
		s.list.edge(false)
	case keys.End:
		s.list.edge(true)
	case keys.Right:
		s.setShown(true)
	case keys.Left:
		s.setShown(false)
	case keys.Escape:
		s.filtering, s.query = false, nil
		s.list.setQuery(nil)
	case keys.Backspace, keys.CtrlH:
		switch {
		case len(s.query) > 0:
			s.query = s.query[:len(s.query)-1]
			s.list.setQuery(s.query)
		case s.filtering:
			s.filtering = false
		}

	case keys.RuneKey, keys.Space:
		if !k.AltPressed {
			p.typed(s, k.Runes)
		}
	}

	return false
}

// typed handles typed letters: shortcuts, or the filter while it is open.
func (p MultiselectPrinter) typed(s *multiselectState, runes []rune) {
	if s.filtering {
		s.query = append(s.query, runes...)
		s.list.setQuery(s.query)

		return
	}

	if len(runes) != 1 {
		return
	}

	switch r := unicode.ToLower(runes[0]); {
	case r == '/' && p.Filter:
		s.filtering = true
	case p.KeyAll != 0 && r == unicode.ToLower(p.KeyAll):
		s.setShown(true)
	case p.KeyNone != 0 && r == unicode.ToLower(p.KeyNone):
		s.setShown(false)
	}
}

func (p MultiselectPrinter) glyphs(c canvas) (pointer, checked, unchecked string) {
	pointer, checked, unchecked = p.Pointer, p.Checked, p.Unchecked
	if pointer == "" {
		pointer = c.g.Pointer
	}

	if checked == "" {
		checked = c.g.Checked
	}

	if unchecked == "" {
		unchecked = c.g.Unchecked
	}

	return pointer, checked, unchecked
}

func (p MultiselectPrinter) progress(s *multiselectState) string {
	return strconv.Itoa(s.count()) + " of " + strconv.Itoa(len(p.Options))
}

func (p MultiselectPrinter) view(c canvas, s *multiselectState) string {
	first := question(c, p.Text, p.progress(s))

	if s.filtering {
		first += " " + c.paint(ToneMuted, "/")

		if len(s.query) > 0 {
			first += c.paint(ToneText, string(s.query))
		} else if p.FilterPlaceholder != "" {
			first += c.paint(ToneMuted, p.FilterPlaceholder)
		}
	}

	pointer, checked, unchecked := p.glyphs(c)

	v := listView{
		styles:    p.Styles,
		pointer:   pointer,
		noMatches: "no matches",
		check: func(i int) string {
			if s.checked[i] {
				return c.style(p.Styles.Checked, checked)
			}

			return c.style(p.Styles.Unchecked, unchecked)
		},
	}

	lines := []string{first}
	lines = append(lines, s.list.render(c, v)...)

	if p.KeyHints {
		lines = append(lines, "", footer(c, p.hints(c, s)...))
	}

	return strings.Join(lines, "\n")
}

func (p MultiselectPrinter) hints(c canvas, s *multiselectState) []keyHint {
	move := keyHint{c.g.ArrowUp + c.g.ArrowDown, "move", 3}
	confirm := keyHint{p.KeyConfirm.String(), "confirm", 0}
	cancel := keyHint{"ctrl+c", "cancel", 4}
	hs := []keyHint{{p.KeySelect.String(), "toggle", 0}}

	if s.filtering {
		return append(hs, move, confirm, keyHint{"esc", "clear filter", 2}, cancel)
	}

	if p.KeyAll != 0 {
		hs = append(hs, keyHint{string(p.KeyAll), "all", 1})
	}

	if p.KeyNone != 0 {
		hs = append(hs, keyHint{string(p.KeyNone), "none", 1})
	}

	hs = append(hs, move, confirm)

	if p.Filter {
		hs = append(hs, keyHint{"/", "filter", 2})
	}

	return append(hs, cancel)
}

func (p MultiselectPrinter) collapse(c canvas, st status, s *multiselectState) string {
	line := answered(c, st, p.Text, p.progress(s))
	if st == statusCanceled || !p.ShowSelectedOptions || s.count() == 0 {
		return line
	}

	lines := []string{line}
	sep := c.paint(ToneMuted, " "+c.g.Dot+" ")
	sepWidth := text.Width(sep)

	var (
		row  []string
		used int
	)

	// Lines break between options, never inside a label.
	for i, checked := range s.checked {
		if !checked {
			continue
		}

		label := p.Options[i].Label
		w := text.Width(label)

		if c.width > 0 && len(row) > 0 && len(promptIndent)+used+sepWidth+w > c.width-1 {
			lines = append(lines, promptIndent+strings.Join(row, sep))
			row, used = nil, 0
		}

		if len(row) > 0 {
			used += sepWidth
		}

		row = append(row, c.paint(ToneText, label))
		used += w
	}

	lines = append(lines, promptIndent+strings.Join(row, sep))

	return strings.Join(lines, "\n")
}
