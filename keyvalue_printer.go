package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// KeyValuePair is one line of a KeyValuePrinter: a key, its value and an
// optional muted suffix after the value.
type KeyValuePair struct {
	Key    string
	Value  string
	Suffix string
}

// KV returns a pair with key and value. Values that are not strings are
// formatted with fmt.Sprint.
func KV(key string, value any) KeyValuePair {
	return KeyValuePair{Key: key, Value: fmt.Sprint(value)}
}

// WithSuffix returns a copy of kv with a muted suffix after the value, such
// as a unit or a note.
func (kv KeyValuePair) WithSuffix(s string) KeyValuePair {
	kv.Suffix = s
	return kv
}

// KeyValuePrinter renders key-value pairs with aligned keys:
//
//	Version  0.1.0
//	Path     ~/.local/fleet/dotfiles/src/cli/main.ts
//	Uptime   1m 34s since last boot
//
// Values that do not fit the terminal wrap under themselves.
type KeyValuePrinter struct {
	// Pairs are the lines in order.
	Pairs []KeyValuePair
	// KeyStyle styles the keys.
	KeyStyle Style
	// ValueStyle styles the values.
	ValueStyle Style
	// SuffixStyle styles the suffixes.
	SuffixStyle Style
	// Indent is the number of spaces before every key.
	Indent int
	// Gap is the number of spaces between the widest key and the values.
	Gap int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// KeyValue is the default key-value printer.
var KeyValue = KeyValuePrinter{
	KeyStyle:    ToneMuted.Style(),
	ValueStyle:  ToneText.Style(),
	SuffixStyle: ToneMuted.Style(),
	Indent:      2,
	Gap:         2,
}

// WithPairs returns a copy of p that renders pairs.
func (p KeyValuePrinter) WithPairs(pairs ...KeyValuePair) *KeyValuePrinter {
	p.Pairs = pairs
	return &p
}

// WithKeyStyle returns a copy of p whose keys use s.
func (p KeyValuePrinter) WithKeyStyle(s Style) *KeyValuePrinter {
	p.KeyStyle = s
	return &p
}

// WithValueStyle returns a copy of p whose values use s.
func (p KeyValuePrinter) WithValueStyle(s Style) *KeyValuePrinter {
	p.ValueStyle = s
	return &p
}

// WithAccentValues returns a copy of p that prints values in the Accent
// tone, which suits versions, paths and numbers. False goes back to Text.
func (p KeyValuePrinter) WithAccentValues(b ...bool) *KeyValuePrinter {
	p.ValueStyle = ToneText.Style()
	if flag(b) {
		p.ValueStyle = ToneAccent.Style()
	}

	return &p
}

// WithSuffixStyle returns a copy of p whose suffixes use s.
func (p KeyValuePrinter) WithSuffixStyle(s Style) *KeyValuePrinter {
	p.SuffixStyle = s
	return &p
}

// WithIndent returns a copy of p with n spaces before every key.
func (p KeyValuePrinter) WithIndent(n int) *KeyValuePrinter {
	p.Indent = max(n, 0)
	return &p
}

// WithGap returns a copy of p with n spaces between the keys and values.
func (p KeyValuePrinter) WithGap(n int) *KeyValuePrinter {
	p.Gap = max(n, 1)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p KeyValuePrinter) WithWriter(w io.Writer) *KeyValuePrinter {
	p.Writer = w
	return &p
}

func (p KeyValuePrinter) out() io.Writer { return orDefault(p.Writer) }

// Render prints the pairs.
func (p KeyValuePrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// Srender returns the pairs. It returns ErrEmpty when there are none.
func (p KeyValuePrinter) Srender() (string, error) {
	if len(p.Pairs) == 0 {
		return "", ErrEmpty
	}

	c := canvasFor(p.out())

	keys := 0
	for _, kv := range p.Pairs {
		keys = max(keys, text.Width(kv.Key))
	}

	var lines []string

	for _, kv := range p.Pairs {
		lead := strings.Repeat(" ", p.Indent) + c.style(p.KeyStyle, kv.Key)
		body := []span{{kv.Value, p.ValueStyle}, {kv.Suffix, p.SuffixStyle}}
		lines = append(lines, hanging(c, lead, p.Indent+keys, p.Gap, body)...)
	}

	return strings.Join(lines, "\n"), nil
}

// span is a piece of text with its style.
type span struct {
	text  string
	style Style
}

// flow wraps segments to width columns, keeping each segment's style. Lines
// break at spaces and at newlines; a word wider than a line is broken. A
// width of zero or less only breaks at newlines.
func flow(c canvas, segments []span, width int) []string {
	type run struct {
		style Style
		words []string
	}

	var (
		lines [][]run
		cur   []run
		used  int
	)

	add := func(word string, st Style) {
		w := text.Width(word)
		if len(cur) > 0 && width > 0 && used+1+w > width {
			lines = append(lines, cur)
			cur, used = nil, 0
		}

		if len(cur) > 0 {
			used++
		}

		used += w

		if len(cur) > 0 && cur[len(cur)-1].style == st {
			cur[len(cur)-1].words = append(cur[len(cur)-1].words, word)
			return
		}

		cur = append(cur, run{style: st, words: []string{word}})
	}

	for _, seg := range segments {
		for i, line := range strings.Split(strings.ReplaceAll(seg.text, "\r\n", "\n"), "\n") {
			if i > 0 {
				lines = append(lines, cur)
				cur, used = nil, 0
			}

			for word := range strings.SplitSeq(line, " ") {
				if word == "" {
					continue
				}

				pieces := []string{word}
				if width > 0 && text.Width(word) > width {
					pieces = text.Wrap(word, width)
				}

				for _, piece := range pieces {
					add(piece, seg.style)
				}
			}
		}
	}

	lines = append(lines, cur)

	out := make([]string, len(lines))

	for i, line := range lines {
		parts := make([]string, len(line))
		for j, r := range line {
			parts[j] = c.style(r.style, strings.Join(r.words, " "))
		}

		out[i] = strings.Join(parts, " ")
	}

	return out
}

// hanging returns a two-column entry: lead, which is column columns wide,
// then gap spaces and body. Body wraps to the canvas with every following
// line indented to where the body starts. A lead wider than the column
// puts the body on the next line.
func hanging(c canvas, lead string, column, gap int, body []span) []string {
	start := column + gap
	leadWidth := text.Width(lead)

	var lines []string

	first := lead + strings.Repeat(" ", max(start-leadWidth, gap))
	if leadWidth > column {
		lines = append(lines, lead)
		first = strings.Repeat(" ", start)
	}

	indent := strings.Repeat(" ", start)

	width := 0
	if c.width > 0 {
		width = max(c.width-start, 10)
	}

	for i, line := range flow(c, body, width) {
		switch {
		case line == "" && i == 0 && len(lines) == 0:
			lines = append(lines, lead)
			continue
		case line == "":
			lines = append(lines, "")
			continue
		}

		if i == 0 {
			lines = append(lines, first+line)
			continue
		}

		lines = append(lines, indent+line)
	}

	return lines
}
