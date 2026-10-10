package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// CenterPrinter centers text in the terminal. By default the text moves as
// one block, so its lines stay aligned with each other; use it to center
// the output of other printers, such as a box. When the terminal width is
// unknown the text is printed as it is.
//
// The text keeps its own styling; the center printer adds no colors.
type CenterPrinter struct {
	// CenterEachLineSeparately centers every line on its own instead of
	// the text as a block.
	CenterEachLineSeparately bool
	// Width is the width to center in. Zero uses the terminal width.
	Width int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Center is the default center printer.
var Center = CenterPrinter{}

// WithCenterEachLineSeparately returns a copy of p that centers every line
// on its own.
func (p CenterPrinter) WithCenterEachLineSeparately(b ...bool) *CenterPrinter {
	p.CenterEachLineSeparately = flag(b)
	return &p
}

// WithWidth returns a copy of p that centers in w columns. Zero uses the
// terminal width.
func (p CenterPrinter) WithWidth(w int) *CenterPrinter {
	p.Width = max(w, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p CenterPrinter) WithWriter(w io.Writer) *CenterPrinter {
	p.Writer = w
	return &p
}

func (p CenterPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p CenterPrinter) render(c canvas, s string) string {
	width := c.width
	if p.Width > 0 {
		width = p.Width
	}

	s = strings.TrimRight(s, "\n")
	if width <= 0 {
		return s
	}

	lines := strings.Split(s, "\n")
	block := text.MaxWidth(s)

	for i, line := range lines {
		w := block
		if p.CenterEachLineSeparately {
			w = text.Width(line)
		}

		if margin := (width - w) / 2; margin > 0 && line != "" {
			lines[i] = strings.Repeat(" ", margin) + line
		}
	}

	return strings.Join(lines, "\n")
}

// Sprint returns the centered text.
func (p CenterPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p CenterPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats and returns the centered text.
func (p CenterPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p CenterPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the centered text. Centered text is a whole block, so
// Print and Println print the same.
func (p CenterPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the centered text.
func (p CenterPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats and prints the centered text.
func (p CenterPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats and prints the centered text.
func (p CenterPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints every non-nil error in a, centered.
func (p CenterPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
