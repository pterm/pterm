package pterm

import (
	"github.com/mattn/go-runewidth"
	"github.com/pterm/pterm/internal"
	"strings"
)

// DefaultCenter is the default CenterPrinter.
var DefaultCenter = CenterPrinter{
	CenterEachLineSeparately: false,
}

// CenterPrinter prints centered text.
type CenterPrinter struct {
	CenterEachLineSeparately bool
}

// WithCenterEachLineSeparately centers each line separately.
func (p CenterPrinter) WithCenterEachLineSeparately(b ...bool) *CenterPrinter {
	bt := internal.WithBoolean(b)
	p.CenterEachLineSeparately = bt
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p CenterPrinter) Sprint(a ...interface{}) string {
	lines := strings.Split(Sprint(a...), "\n")

	var ret string

	if p.CenterEachLineSeparately {
		for _, line := range lines {
			ret += strings.Repeat(" ", (GetTerminalWidth()-runewidth.StringWidth(RemoveColorFromString(line)))/2) + line + "\n"
		}
		return ret
	}

	var maxLineWidth int

	for _, line := range lines {
		lineLength := runewidth.StringWidth(RemoveColorFromString(line))
		if maxLineWidth < lineLength {
			maxLineWidth = lineLength
		}
	}

	indent := GetTerminalWidth() - maxLineWidth

	for _, line := range lines {
		ret += strings.Repeat(" ", indent/2) + line + "\n"
	}

	return ret
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p CenterPrinter) Sprintln(a ...interface{}) string {
	return Sprintln(p.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p CenterPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p CenterPrinter) Print(a ...interface{}) *TextPrinter {
	Print(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p CenterPrinter) Println(a ...interface{}) *TextPrinter {
	Println(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p CenterPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}
