package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// ParagraphPrinter prints running text, wrapped at word boundaries to the
// terminal width or MaxWidth, whichever is smaller. Whitespace inside a
// paragraph collapses to single spaces, and a blank line starts a new
// paragraph.
type ParagraphPrinter struct {
	// MaxWidth is the widest a line gets. Zero only wraps to the terminal.
	MaxWidth int
	// Style styles the text.
	Style Style
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Paragraph is the default paragraph printer. Its lines are at most 80
// columns wide, which keeps long text readable on wide terminals.
var Paragraph = ParagraphPrinter{
	MaxWidth: 80,
	Style:    ToneText.Style(),
}

// WithMaxWidth returns a copy of p whose lines are at most w columns wide.
// Zero only wraps to the terminal.
func (p ParagraphPrinter) WithMaxWidth(w int) *ParagraphPrinter {
	p.MaxWidth = max(w, 0)
	return &p
}

// WithStyle returns a copy of p that styles the text with s.
func (p ParagraphPrinter) WithStyle(s Style) *ParagraphPrinter {
	p.Style = s
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p ParagraphPrinter) WithWriter(w io.Writer) *ParagraphPrinter {
	p.Writer = w
	return &p
}

func (p ParagraphPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p ParagraphPrinter) render(c canvas, s string) string {
	width := c.width
	if p.MaxWidth > 0 && (width <= 0 || p.MaxWidth < width) {
		width = p.MaxWidth
	}

	s = strings.ReplaceAll(s, "\r\n", "\n")

	var paragraphs []string

	for part := range strings.SplitSeq(s, "\n\n") {
		if part = strings.Join(strings.Fields(part), " "); part != "" {
			paragraphs = append(paragraphs, strings.Join(text.Wrap(part, width), "\n"))
		}
	}

	return c.style(p.Style, strings.Join(paragraphs, "\n\n"))
}

// Sprint returns the wrapped text.
func (p ParagraphPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p ParagraphPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats and returns the wrapped text.
func (p ParagraphPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p ParagraphPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the wrapped text. A paragraph is a whole block, so Print
// and Println print the same.
func (p ParagraphPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the wrapped text.
func (p ParagraphPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats and prints the wrapped text.
func (p ParagraphPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats and prints the wrapped text.
func (p ParagraphPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints every non-nil error in a as a paragraph.
func (p ParagraphPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
