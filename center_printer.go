package pterm

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
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
	if RawOutput {
		return Sprint(a...)
	}

	lines := strings.Split(Sprint(a...), "\n")

	var ret string

	if p.CenterEachLineSeparately {
		for _, line := range lines {
			margin := (GetTerminalWidth() - runewidth.StringWidth(RemoveColorFromString(line))) / 2
			if margin < 1 {
				ret += line + "\n"
			} else {
				ret += strings.Repeat(" ", margin) + line + "\n"
			}
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

	if indent/2 < 1 {
		for _, line := range lines {
			ret += line + "\n"
		}

		return ret
	}

	for _, line := range lines {
		ret += strings.Repeat(" ", indent/2) + line + "\n"
	}

	return ret
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p CenterPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(Sprintln(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p CenterPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p CenterPrinter) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
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
	Print(p.Sprintln(a...))
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

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p CenterPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p CenterPrinter) PrintOnError(a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				p.Println(err)
			}
		}
	}

	tp := TextPrinter(p)
	return &tp
}

// PrintOnErrorf wraps every error which is not nil and prints it.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p CenterPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				p.Println(fmt.Errorf(format, err))
			}
		}
	}

	tp := TextPrinter(p)
	return &tp
}
