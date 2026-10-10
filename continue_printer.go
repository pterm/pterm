package pterm

import (
	"context"
	"io"
	"strings"
	"unicode"

	"atomicgo.dev/keyboard/keys"
)

// ContinuePrinter asks a question with a few short answers:
//
//	? Overwrite config.yaml? [Y/n/a/c] yes · no · all · cancel
//
// Each answer has a handle, by default its first letter, and typing it
// answers at once. Enter takes the highlighted answer, which starts on the
// default and moves with the arrow keys.
type ContinuePrinter struct {
	// Text is the question.
	Text string
	// Options are the answers.
	Options []string
	// Handles are the keys of the answers, one per option. Only their
	// first letter counts. Nil uses the first letter of every option.
	Handles []string
	// DefaultValueIndex is the index of the default answer. It is
	// returned without asking when standard input is not a terminal.
	DefaultValueIndex int
	// ShowOptions lists the full answers muted after the handles.
	ShowOptions bool
	// Styles style the handles.
	Styles ContinueStyles
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Input is read for key presses instead of the terminal.
	Input io.Reader
	// Context cancels the prompt when it ends.
	Context context.Context
	// OnInterrupt runs on Ctrl+C instead of exiting the program.
	OnInterrupt func()
}

// ContinueStyles style a Continue prompt.
type ContinueStyles struct {
	// Handle styles the handles.
	Handle Style
	// Active styles the highlighted handle.
	Active Style
	// Frame styles the brackets and slashes around the handles.
	Frame Style
	// Hint styles the list of full answers.
	Hint Style
}

// Continue is the default continue prompt.
var Continue = ContinuePrinter{
	Text:        "Continue?",
	Options:     []string{"yes", "no", "all", "cancel"},
	ShowOptions: true,
	Styles: ContinueStyles{
		Handle: Style{Fg: ToneText},
		Active: Style{Fg: TonePrimary, Attrs: boldUnderline},
		Frame:  Style{Fg: ToneMuted},
		Hint:   Style{Fg: ToneMuted},
	},
}

// WithText returns a copy of p that asks t.
func (p ContinuePrinter) WithText(t string) *ContinuePrinter {
	p.Text = t
	return &p
}

// WithOptions returns a copy of p with the answers options.
func (p ContinuePrinter) WithOptions(options []string) *ContinuePrinter {
	p.Options = options
	return &p
}

// WithHandles returns a copy of p whose answers are typed with handles.
// Handles that do not match the options in number are ignored.
func (p ContinuePrinter) WithHandles(handles []string) *ContinuePrinter {
	p.Handles = handles
	return &p
}

// WithDefaultValueIndex returns a copy of p whose default answer is the
// option at index i.
func (p ContinuePrinter) WithDefaultValueIndex(i int) *ContinuePrinter {
	p.DefaultValueIndex = i
	return &p
}

// WithDefaultValue returns a copy of p whose default answer is the option
// v.
func (p ContinuePrinter) WithDefaultValue(v string) *ContinuePrinter {
	for i, o := range p.Options {
		if o == v {
			p.DefaultValueIndex = i
			break
		}
	}

	return &p
}

// WithShowOptions returns a copy of p that does or does not list the full
// answers.
func (p ContinuePrinter) WithShowOptions(b ...bool) *ContinuePrinter {
	p.ShowOptions = flag(b)
	return &p
}

// WithStyles returns a copy of p with the styles st.
func (p ContinuePrinter) WithStyles(st ContinueStyles) *ContinuePrinter {
	p.Styles = st
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p ContinuePrinter) WithWriter(w io.Writer) *ContinuePrinter {
	p.Writer = w
	return &p
}

// WithInput returns a copy of p that reads key presses from r, for tests
// and scripted input.
func (p ContinuePrinter) WithInput(r io.Reader) *ContinuePrinter {
	p.Input = r
	return &p
}

// WithContext returns a copy of p that gives up with ErrCanceled when ctx
// ends.
func (p ContinuePrinter) WithContext(ctx context.Context) *ContinuePrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting. Show then returns ErrInterrupted.
func (p ContinuePrinter) WithOnInterrupt(fn func()) *ContinuePrinter {
	p.OnInterrupt = fn
	return &p
}

// handles returns one handle per option.
func (p ContinuePrinter) handles() []string {
	if len(p.Handles) == len(p.Options) {
		return p.Handles
	}

	hs := make([]string, len(p.Options))
	for i, o := range p.Options {
		for _, r := range o {
			hs[i] = string(unicode.ToLower(r))
			break
		}
	}

	return hs
}

// Show asks the question, or text when given, and returns the chosen
// option. It returns ErrEmpty when there are no options.
func (p ContinuePrinter) Show(text ...string) (string, error) {
	if len(text) > 0 {
		p.Text = strings.Join(text, " ")
	}

	if len(p.Options) == 0 {
		return "", ErrEmpty
	}

	def := min(max(p.DefaultValueIndex, 0), len(p.Options)-1)
	current := def
	handles := p.handles()

	pr := prompt[string]{
		cfg:  promptConfig{writer: p.Writer, input: p.Input, ctx: p.Context, onInterrupt: p.OnInterrupt},
		view: func(c canvas) string { return p.view(c, handles, def, current) },
		key: func(k keys.Key) bool {
			n := len(p.Options)

			switch k.Code {
			case keys.Enter:
				return true
			case keys.Right, keys.Tab, keys.Down:
				current = (current + 1) % n
			case keys.Left, keys.ShiftTab, keys.Up:
				current = (current + n - 1) % n
			case keys.RuneKey:
				r := firstRune(k)
				for i, h := range handles {
					if firstLetter(h) == r {
						current = i
						return true
					}
				}
			}

			return false
		},
		result: func() string { return p.Options[current] },
		collapse: func(c canvas, st status) string {
			return answered(c, st, p.Text, p.Options[current])
		},
	}

	return pr.run(p.Options[def], true)
}

func (p ContinuePrinter) view(c canvas, handles []string, def, current int) string {
	parts := make([]string, len(handles))

	for i, h := range handles {
		if i == def {
			h = strings.ToUpper(h)
		}

		st := p.Styles.Handle
		if i == current {
			st = p.Styles.Active
		}

		parts[i] = c.style(st, h)
	}

	line := question(c, p.Text, "") + " " +
		c.style(p.Styles.Frame, "[") +
		strings.Join(parts, c.style(p.Styles.Frame, "/")) +
		c.style(p.Styles.Frame, "]")

	if p.ShowOptions {
		line += " " + c.style(p.Styles.Hint, strings.Join(p.Options, " "+c.g.Dot+" "))
	}

	return line
}
