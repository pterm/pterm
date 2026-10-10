package pterm

import (
	"io"
	"strings"
)

// KeyHint is one entry of a key hint line: a key and what it does. A hint
// without a key is shown as plain muted information, such as "every 5s".
type KeyHint struct {
	Key  string
	Desc string
}

// KeyHintsPrinter prints the keys a user can press, on one line:
//
//	[space] toggle · [a] all · [enter] confirm
//
// Interactive printers draw their footer with it.
type KeyHintsPrinter struct {
	// Hints are the entries, in order.
	Hints []KeyHint
	// KeyStyle styles the keys.
	KeyStyle Style
	// DescStyle styles the descriptions and hints without a key.
	DescStyle Style
	// BracketTone colors the brackets around keys.
	BracketTone Tone
	// Open and Close surround each key.
	Open, Close string
	// Separator goes between entries. Empty uses the inline separator of
	// the glyph set, with a space on each side.
	Separator string
	// SeparatorTone colors the separator.
	SeparatorTone Tone
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// KeyHints is the default key hint printer.
var KeyHints = KeyHintsPrinter{
	KeyStyle:      ToneText.Style(),
	DescStyle:     ToneMuted.Style(),
	BracketTone:   ToneMuted,
	Open:          "[",
	Close:         "]",
	SeparatorTone: ToneMuted,
}

// WithHints returns a copy of p with the entries h.
func (p KeyHintsPrinter) WithHints(h ...KeyHint) *KeyHintsPrinter {
	p.Hints = h
	return &p
}

// WithHint returns a copy of p with one more entry: key and what it does.
// An empty key adds plain information.
func (p KeyHintsPrinter) WithHint(key, desc string) *KeyHintsPrinter {
	p.Hints = append(p.Hints[:len(p.Hints):len(p.Hints)], KeyHint{Key: key, Desc: desc})
	return &p
}

// WithKeyStyle returns a copy of p that styles keys with s.
func (p KeyHintsPrinter) WithKeyStyle(s Style) *KeyHintsPrinter {
	p.KeyStyle = s
	return &p
}

// WithDescStyle returns a copy of p that styles descriptions with s.
func (p KeyHintsPrinter) WithDescStyle(s Style) *KeyHintsPrinter {
	p.DescStyle = s
	return &p
}

// WithBracketTone returns a copy of p whose brackets are in tone t.
func (p KeyHintsPrinter) WithBracketTone(t Tone) *KeyHintsPrinter {
	p.BracketTone = t
	return &p
}

// WithBrackets returns a copy of p that surrounds keys with open and
// closing. Empty strings draw keys without brackets.
func (p KeyHintsPrinter) WithBrackets(open, closing string) *KeyHintsPrinter {
	p.Open, p.Close = open, closing
	return &p
}

// WithSeparator returns a copy of p that puts s between entries.
func (p KeyHintsPrinter) WithSeparator(s string) *KeyHintsPrinter {
	p.Separator = s
	return &p
}

// WithSeparatorTone returns a copy of p whose separator is in tone t.
func (p KeyHintsPrinter) WithSeparatorTone(t Tone) *KeyHintsPrinter {
	p.SeparatorTone = t
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p KeyHintsPrinter) WithWriter(w io.Writer) *KeyHintsPrinter {
	p.Writer = w
	return &p
}

func (p KeyHintsPrinter) out() io.Writer { return orDefault(p.Writer) }

// render draws the hint line on c, cut to the canvas width. Interactive
// printers call it for their footer. It returns "" without hints.
func (p KeyHintsPrinter) render(c canvas) string {
	sep := p.Separator
	if sep == "" {
		sep = " " + c.g.Dot + " "
	}

	sep = c.paint(p.SeparatorTone, sep)
	parts := make([]string, 0, len(p.Hints))

	for _, h := range p.Hints {
		var part string

		if h.Key != "" {
			part = c.paint(p.BracketTone, p.Open) + c.style(p.KeyStyle, h.Key) + c.paint(p.BracketTone, p.Close)
			if h.Desc != "" {
				part += " "
			}
		}

		if h.Desc != "" {
			part += c.style(p.DescStyle, h.Desc)
		}

		if part != "" {
			parts = append(parts, part)
		}
	}

	return c.cut(strings.Join(parts, sep), c.width)
}

// keyHints draws hints with the default key hint printer, for the footers
// of interactive printers.
//
//nolint:unused // the interactive printers call it.
func keyHints(c canvas, hints ...KeyHint) string {
	return KeyHints.WithHints(hints...).render(c)
}

// Srender returns the hint line. It returns ErrEmpty without hints.
func (p KeyHintsPrinter) Srender() (string, error) {
	s := p.render(canvasFor(p.out()))
	if s == "" {
		return "", ErrEmpty
	}

	return s, nil
}

// Render prints the hint line. It returns ErrEmpty without hints.
func (p KeyHintsPrinter) Render() error {
	s, err := p.Srender()

	return renderTo(p.out(), s, err)
}
