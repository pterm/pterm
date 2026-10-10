package pterm

import (
	"fmt"
	"io"
	"strings"
)

// BasicTextPrinter prints plain text in the Text tone, wrapped to the
// terminal width. Use it instead of fmt when output should follow the
// theme and appear above running live printers.
type BasicTextPrinter struct {
	// Style styles the text.
	Style Style
	// NoWrap keeps long lines as they are instead of wrapping them to the
	// terminal width.
	NoWrap bool
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Text is the default text printer.
var Text = BasicTextPrinter{Style: ToneText.Style()}

// WithStyle returns a copy of p that styles text with s.
func (p BasicTextPrinter) WithStyle(s Style) *BasicTextPrinter {
	p.Style = s
	return &p
}

// WithTone returns a copy of p that colors text with t.
func (p BasicTextPrinter) WithTone(t Tone) *BasicTextPrinter {
	p.Style.Fg = t
	return &p
}

// WithNoWrap returns a copy of p that does not wrap long lines.
func (p BasicTextPrinter) WithNoWrap(b ...bool) *BasicTextPrinter {
	p.NoWrap = flag(b)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p BasicTextPrinter) WithWriter(w io.Writer) *BasicTextPrinter {
	p.Writer = w
	return &p
}

func (p BasicTextPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p BasicTextPrinter) render(c canvas, s string) string {
	if p.NoWrap {
		return c.style(p.Style, s)
	}

	return c.style(p.Style, strings.Join(c.wrap(s, 0), "\n"))
}

// Sprint returns the styled text.
func (p BasicTextPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p BasicTextPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats and returns the styled text.
func (p BasicTextPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p BasicTextPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the text without a newline, like fmt.Print.
func (p BasicTextPrinter) Print(a ...any) { printText(p, fmt.Sprint(a...)) }

// Println prints the text and a newline.
func (p BasicTextPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats and prints the text without a newline.
func (p BasicTextPrinter) Printf(format string, a ...any) { printText(p, fmt.Sprintf(format, a...)) }

// Printfln formats and prints the text and a newline.
func (p BasicTextPrinter) Printfln(format string, a ...any) {
	printLine(p, fmt.Sprintf(format, a...))
}

// PrintOnError prints every non-nil error in a.
func (p BasicTextPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
