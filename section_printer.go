package pterm

import (
	"fmt"
	"io"
	"strings"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/text"
)

// SectionPrinter prints the title of a part of the output, with one blank
// line before it. Level 1 is a bold Primary title, level 2 a bold Text
// title and deeper levels are muted.
type SectionPrinter struct {
	// Level is the depth of the section, starting at 1.
	Level int
	// Styles are the title styles per level: Styles[0] for level 1 and so
	// on. Levels deeper than the list use its last style.
	Styles []Style
	// Prefix is repeated Level times before the title, like "#" in
	// Markdown. Empty prints no prefix.
	Prefix string
	// PrefixTone colors the prefix.
	PrefixTone Tone
	// TopPadding is the number of blank lines before the title.
	TopPadding int
	// BottomPadding is the number of blank lines after the title.
	BottomPadding int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Section is the default section printer.
var Section = SectionPrinter{
	Level: 1,
	Styles: []Style{
		TonePrimary.Bold(),
		ToneText.Bold(),
		ToneMuted.Style(),
	},
	PrefixTone: ToneMuted,
	TopPadding: 1,
}

// WithLevel returns a copy of p at level n, starting at 1.
func (p SectionPrinter) WithLevel(n int) *SectionPrinter {
	p.Level = max(n, 1)
	return &p
}

// WithStyle returns a copy of p that styles titles of every level with s.
func (p SectionPrinter) WithStyle(s Style) *SectionPrinter {
	p.Styles = []Style{s}
	return &p
}

// WithStyles returns a copy of p with one title style per level.
func (p SectionPrinter) WithStyles(s ...Style) *SectionPrinter {
	p.Styles = s
	return &p
}

// WithPrefix returns a copy of p that repeats s Level times before the
// title.
func (p SectionPrinter) WithPrefix(s string) *SectionPrinter {
	p.Prefix = s
	return &p
}

// WithPrefixTone returns a copy of p whose prefix is in tone t.
func (p SectionPrinter) WithPrefixTone(t Tone) *SectionPrinter {
	p.PrefixTone = t
	return &p
}

// WithTopPadding returns a copy of p with n blank lines before the title.
func (p SectionPrinter) WithTopPadding(n int) *SectionPrinter {
	p.TopPadding = max(n, 0)
	return &p
}

// WithBottomPadding returns a copy of p with n blank lines after the
// title.
func (p SectionPrinter) WithBottomPadding(n int) *SectionPrinter {
	p.BottomPadding = max(n, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p SectionPrinter) WithWriter(w io.Writer) *SectionPrinter {
	p.Writer = w
	return &p
}

func (p SectionPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p SectionPrinter) style() Style {
	if len(p.Styles) == 0 {
		return Style{Attrs: color.Bold}
	}

	return p.Styles[min(max(p.Level, 1), len(p.Styles))-1]
}

func (p SectionPrinter) render(c canvas, s string) string {
	s = strings.TrimRight(s, "\n")

	indent := 0
	prefix := ""

	if p.Prefix != "" {
		prefix = strings.Repeat(p.Prefix, max(p.Level, 1)) + " "
		indent = text.Width(prefix)
		prefix = c.paint(p.PrefixTone, prefix)
	}

	lines := c.wrap(s, indent)
	for i, line := range lines {
		lines[i] = c.style(p.style(), line)
		if i == 0 {
			lines[i] = prefix + lines[i]
		} else if indent > 0 {
			lines[i] = strings.Repeat(" ", indent) + lines[i]
		}
	}

	return strings.Repeat("\n", p.TopPadding) + strings.Join(lines, "\n") + strings.Repeat("\n", p.BottomPadding)
}

// Sprint returns the section title, with its padding.
func (p SectionPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p SectionPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the title and returns the section title.
func (p SectionPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p SectionPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the section title. A section is a whole line, so Print and
// Println print the same.
func (p SectionPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the section title.
func (p SectionPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats and prints the section title.
func (p SectionPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats and prints the section title.
func (p SectionPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints a section titled with every non-nil error in a.
func (p SectionPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
