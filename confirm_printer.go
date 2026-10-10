package pterm

import (
	"context"
	"io"
	"strings"
	"unicode"

	"atomicgo.dev/keyboard/keys"
)

// ConfirmPrinter asks a yes or no question:
//
//	? Delete 23 caches?  Yes / No
//
// "y" and "n" answer at once, the arrow keys and tab move between the two
// answers, and enter takes the highlighted one.
type ConfirmPrinter struct {
	// Text is the question.
	Text string
	// DefaultValue is highlighted first, and returned without asking when
	// standard input is not a terminal.
	DefaultValue bool
	// ConfirmText and RejectText label the answers. Their first letters are
	// the shortcut keys.
	ConfirmText string
	RejectText  string
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Input is read for key presses instead of the terminal.
	Input io.Reader
	// Context cancels the prompt when it ends.
	Context context.Context
	// OnInterrupt runs on Ctrl+C instead of exiting the program.
	OnInterrupt func()
}

// Confirm is the default yes or no prompt.
var Confirm = ConfirmPrinter{
	Text:        "Continue?",
	ConfirmText: "Yes",
	RejectText:  "No",
}

// WithText returns a copy of p that asks t.
func (p ConfirmPrinter) WithText(t string) *ConfirmPrinter {
	p.Text = t
	return &p
}

// WithDefaultValue returns a copy of p that starts on v.
func (p ConfirmPrinter) WithDefaultValue(v bool) *ConfirmPrinter {
	p.DefaultValue = v
	return &p
}

// WithConfirmText returns a copy of p whose yes answer reads t.
func (p ConfirmPrinter) WithConfirmText(t string) *ConfirmPrinter {
	p.ConfirmText = t
	return &p
}

// WithRejectText returns a copy of p whose no answer reads t.
func (p ConfirmPrinter) WithRejectText(t string) *ConfirmPrinter {
	p.RejectText = t
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p ConfirmPrinter) WithWriter(w io.Writer) *ConfirmPrinter {
	p.Writer = w
	return &p
}

// WithInput returns a copy of p that reads key presses from r, for tests
// and scripted input.
func (p ConfirmPrinter) WithInput(r io.Reader) *ConfirmPrinter {
	p.Input = r
	return &p
}

// WithContext returns a copy of p that gives up with ErrCanceled when ctx
// ends.
func (p ConfirmPrinter) WithContext(ctx context.Context) *ConfirmPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting. Show then returns ErrInterrupted.
func (p ConfirmPrinter) WithOnInterrupt(fn func()) *ConfirmPrinter {
	p.OnInterrupt = fn
	return &p
}

// Show asks the question, or text when given, and returns the answer.
func (p ConfirmPrinter) Show(text ...string) (bool, error) {
	if len(text) > 0 {
		p.Text = strings.Join(text, " ")
	}

	value := p.DefaultValue

	pr := prompt[bool]{
		cfg: promptConfig{writer: p.Writer, input: p.Input, ctx: p.Context, onInterrupt: p.OnInterrupt},
		view: func(c canvas) string {
			yes, no := c.paint(ToneMuted, p.ConfirmText), c.paint(ToneMuted, p.RejectText)
			if value {
				yes = c.style(Style{Fg: TonePrimary, Attrs: boldUnderline}, p.ConfirmText)
			} else {
				no = c.style(Style{Fg: TonePrimary, Attrs: boldUnderline}, p.RejectText)
			}

			return question(c, p.Text, "") + " " + yes + c.paint(ToneMuted, " / ") + no
		},
		key: func(k keys.Key) bool {
			switch k.Code {
			case keys.Enter:
				return true
			case keys.Left, keys.Right, keys.Tab, keys.Up, keys.Down:
				value = !value
			case keys.RuneKey:
				switch r := firstRune(k); {
				case r == firstLetter(p.ConfirmText):
					value = true
					return true
				case r == firstLetter(p.RejectText):
					value = false
					return true
				}
			}

			return false
		},
		result: func() bool { return value },
		collapse: func(c canvas, st status) string {
			answer := p.RejectText
			if value {
				answer = p.ConfirmText
			}

			return answered(c, st, p.Text, answer)
		},
	}

	return pr.run(p.DefaultValue, true)
}

func firstRune(k keys.Key) rune {
	if len(k.Runes) == 0 {
		return 0
	}

	return unicode.ToLower(k.Runes[0])
}

func firstLetter(s string) rune {
	for _, r := range s {
		return unicode.ToLower(r)
	}

	return -1
}
