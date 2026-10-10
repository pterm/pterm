package pterm

import (
	"context"
	"io"
	"strings"

	"atomicgo.dev/keyboard/keys"
)

// SelectPrinter asks the user to pick one option from a list:
//
//	? Which cluster?
//	  ❯ kapp-k8s-at   Vienna     3 nodes
//	    kapp-k8s-de   Frankfurt  5 nodes
//
//	  [↑↓] move · [enter] select · [type] filter
//
// Typing filters the list with a fuzzy match and highlights the matched
// letters. Long lists scroll, options can have groups, descriptions and
// details, and disabled options are skipped.
type SelectPrinter struct {
	// Text is the question.
	Text string
	// Options are the choices. WithOptions fills it from plain labels.
	Options []Option
	// DefaultOption is the label or value of the option the cursor starts
	// on. It is returned without asking when standard input is not a
	// terminal. An option with Selected set works the same way.
	DefaultOption string
	// MaxHeight is the number of options shown at once; the rest scrolls.
	// Zero or less shows every option.
	MaxHeight int
	// Filter lets the user narrow the list by typing.
	Filter bool
	// FilterPlaceholder is shown muted after the question until the user
	// types, for example "type to search".
	FilterPlaceholder string
	// Pointer marks the active option. Empty uses the pointer glyph.
	Pointer string
	// Styles style the rows.
	Styles OptionStyles
	// KeyHints shows the key hint footer.
	KeyHints bool
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Input is read for key presses instead of the terminal.
	Input io.Reader
	// Context cancels the prompt when it ends.
	Context context.Context
	// OnInterrupt runs on Ctrl+C instead of exiting the program.
	OnInterrupt func()
}

// Select is the default single choice prompt.
var Select = SelectPrinter{
	Text:      "Select an option",
	MaxHeight: 8,
	Filter:    true,
	Styles:    DefaultOptionStyles,
	KeyHints:  true,
}

// WithText returns a copy of p that asks t.
func (p SelectPrinter) WithText(t string) *SelectPrinter {
	p.Text = t
	return &p
}

// WithOptions returns a copy of p that offers the given labels.
func (p SelectPrinter) WithOptions(labels []string) *SelectPrinter {
	p.Options = optionsOf(labels)
	return &p
}

// WithItems returns a copy of p that offers options with descriptions,
// details, groups or disabled entries.
func (p SelectPrinter) WithItems(options []Option) *SelectPrinter {
	p.Options = options
	return &p
}

// WithDefaultOption returns a copy of p that starts on the option with
// the label or value s.
func (p SelectPrinter) WithDefaultOption(s string) *SelectPrinter {
	p.DefaultOption = s
	return &p
}

// WithMaxHeight returns a copy of p that shows n options at once.
func (p SelectPrinter) WithMaxHeight(n int) *SelectPrinter {
	p.MaxHeight = n
	return &p
}

// WithFilter returns a copy of p that does or does not filter by typing.
func (p SelectPrinter) WithFilter(b ...bool) *SelectPrinter {
	p.Filter = flag(b)
	return &p
}

// WithFilterPlaceholder returns a copy of p that shows s until the user
// types a filter.
func (p SelectPrinter) WithFilterPlaceholder(s string) *SelectPrinter {
	p.FilterPlaceholder = s
	return &p
}

// WithPointer returns a copy of p that marks the active option with s.
func (p SelectPrinter) WithPointer(s string) *SelectPrinter {
	p.Pointer = s
	return &p
}

// WithStyles returns a copy of p with the row styles st.
func (p SelectPrinter) WithStyles(st OptionStyles) *SelectPrinter {
	p.Styles = st
	return &p
}

// WithKeyHints returns a copy of p that does or does not show the key
// hint footer.
func (p SelectPrinter) WithKeyHints(b ...bool) *SelectPrinter {
	p.KeyHints = flag(b)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p SelectPrinter) WithWriter(w io.Writer) *SelectPrinter {
	p.Writer = w
	return &p
}

// WithInput returns a copy of p that reads key presses from r, for tests
// and scripted input.
func (p SelectPrinter) WithInput(r io.Reader) *SelectPrinter {
	p.Input = r
	return &p
}

// WithContext returns a copy of p that gives up with ErrCanceled when ctx
// ends.
func (p SelectPrinter) WithContext(ctx context.Context) *SelectPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting. Show then returns ErrInterrupted.
func (p SelectPrinter) WithOnInterrupt(fn func()) *SelectPrinter {
	p.OnInterrupt = fn
	return &p
}

// defaultIndex returns the index of the default option, or -1.
func (p SelectPrinter) defaultIndex() int {
	for i, o := range p.Options {
		if o.Disabled {
			continue
		}

		if (p.DefaultOption != "" && o.is(p.DefaultOption)) || (p.DefaultOption == "" && o.Selected) {
			return i
		}
	}

	return -1
}

// Show asks the question, or text when given, and returns the value of
// the chosen option. It returns ErrEmpty when there are no options.
func (p SelectPrinter) Show(text ...string) (string, error) {
	if len(text) > 0 {
		p.Text = strings.Join(text, " ")
	}

	if len(p.Options) == 0 {
		return "", ErrEmpty
	}

	def := p.defaultIndex()

	fallback := ""
	if def >= 0 {
		fallback = p.Options[def].value()
	}

	list := newOptionList(p.Options, p.MaxHeight, true)
	if def >= 0 {
		list.moveTo(def)
	}

	var (
		query  []rune
		chosen Option
	)

	pr := prompt[string]{
		cfg: promptConfig{writer: p.Writer, input: p.Input, ctx: p.Context, onInterrupt: p.OnInterrupt},
		view: func(c canvas) string {
			return p.view(c, list, query)
		},
		key: func(k keys.Key) bool {
			switch k.Code {
			case keys.Enter:
				o, _, ok := list.current()
				if !ok || o.Disabled {
					return false
				}

				chosen = o

				return true

			case keys.Up, keys.CtrlP, keys.ShiftTab:
				list.move(-1)
			case keys.Down, keys.CtrlN, keys.Tab:
				list.move(1)
			case keys.PgUp:
				list.page(-1)
			case keys.PgDown:
				list.page(1)
			case keys.Home:
				list.edge(false)
			case keys.End:
				list.edge(true)
			case keys.Backspace, keys.CtrlH:
				if p.Filter && len(query) > 0 {
					query = query[:len(query)-1]
					list.setQuery(query)
				}

			case keys.CtrlU, keys.Escape:
				if p.Filter && len(query) > 0 {
					query = nil
					list.setQuery(query)
				}

			case keys.RuneKey, keys.Space:
				if p.Filter && !k.AltPressed {
					query = append(query, k.Runes...)
					list.setQuery(query)
				}
			}

			return false
		},
		result: func() string { return chosen.value() }, //nolint:gocritic // chosen is set later, a method value would copy it now.
		collapse: func(c canvas, st status) string {
			return answered(c, st, p.Text, chosen.Label)
		},
	}

	return pr.run(fallback, def >= 0)
}

func (p SelectPrinter) view(c canvas, list *optionList, query []rune) string {
	first := question(c, p.Text, "")

	switch {
	case len(query) > 0:
		first += " " + c.paint(ToneText, string(query))
	case p.Filter && p.FilterPlaceholder != "":
		first += " " + c.paint(ToneMuted, p.FilterPlaceholder)
	}

	pointer := p.Pointer
	if pointer == "" {
		pointer = c.g.Pointer
	}

	lines := []string{first}
	lines = append(lines, list.render(c, listView{styles: p.Styles, pointer: pointer, noMatches: "no matches"})...)

	if p.KeyHints {
		hs := []keyHint{
			{c.g.ArrowUp + c.g.ArrowDown, "move", 1},
			{"enter", "select", 0},
		}

		if p.Filter {
			hs = append(hs, keyHint{"type", "filter", 2})
		}

		lines = append(lines, "", footer(c, hs...))
	}

	return strings.Join(lines, "\n")
}
