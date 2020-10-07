package pterm

import (
	"os"
	"strings"

	"github.com/pterm/pterm/internal"
)

var (
	// GrayBoxStyle wraps text in a gray box.
	GrayBoxStyle = NewStyle(BgGray, FgLightWhite)
)

var (
	// Info returns a PrefixPrinter, which can be used to print text with an "info" Prefix.
	Info = PrefixPrinter{
		Prefix: Prefix{
			Text:  "INFO",
			Style: NewStyle(FgBlack, BgCyan),
		},
		MessageStyle: NewStyle(FgLightCyan),
	}

	// Warning returns a PrefixPrinter, which can be used to print text with a "warning" Prefix.
	Warning = PrefixPrinter{
		Prefix: Prefix{
			Text:  "WARNING",
			Style: NewStyle(FgBlack, BgYellow),
		},
		MessageStyle: NewStyle(FgYellow),
	}

	// Success returns a PrefixPrinter, which can be used to print text with a "success" Prefix.
	Success = PrefixPrinter{
		Prefix: Prefix{
			Text:  "SUCCESS",
			Style: NewStyle(FgBlack, BgGreen),
		},
		MessageStyle: NewStyle(FgGreen),
	}

	// Error returns a PrefixPrinter, which can be used to print text with an "error" Prefix.
	Error = PrefixPrinter{
		Prefix: Prefix{
			Text:  "ERROR",
			Style: NewStyle(FgBlack, BgLightRed),
		},
		MessageStyle: NewStyle(FgLightRed),
	}

	// Fatal returns a PrefixPrinter, which can be used to print text with an "fatal" Prefix.
	// NOTICE: Fatal terminates the application immediately!
	Fatal = PrefixPrinter{
		Prefix: Prefix{
			Text:  "FATAL",
			Style: NewStyle(FgBlack, BgLightRed),
		},
		MessageStyle: NewStyle(FgLightRed),
		Fatal:        true,
	}

	// Description returns a PrefixPrinter, which can be used to print text with a "description" Prefix.
	Description = PrefixPrinter{
		Prefix: Prefix{
			Text:  "Description",
			Style: Style{FgLightWhite, BgDarkGray},
		},
		MessageStyle: Style{FgLightWhite},
	}
)

// PrefixPrinter is the printer used to print a Prefix.
type PrefixPrinter struct {
	Prefix       Prefix
	Scope        Scope
	MessageStyle Style
	Fatal        bool
}

// WithPrefix adds a custom prefix to the printer.
func (p PrefixPrinter) WithPrefix(prefix Prefix) *PrefixPrinter {
	p.Prefix = prefix
	return &p
}

// WithScope adds a scope to the Prefix.
func (p PrefixPrinter) WithScope(scope Scope) *PrefixPrinter {
	p.Scope = scope
	return &p
}

// WithMessageStyle adds a custom prefix to the printer.
func (p PrefixPrinter) WithMessageStyle(style Style) *PrefixPrinter {
	p.MessageStyle = style
	return &p
}

// WithFatal sets if the printer should panic after printing.
// NOTE:
// The printer will only panic if either PrefixPrinter.Println, PrefixPrinter.Print
// or PrefixPrinter.Printf is called.
func (p PrefixPrinter) WithFatal(b ...bool) *PrefixPrinter {
	p.Fatal = internal.WithBoolean(b)
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p PrefixPrinter) Sprint(a ...interface{}) string {
	var ret string
	messageLines := strings.Split(Sprint(a...), "\n")
	for i, m := range messageLines {
		if i == 0 {
			ret += p.GetFormattedPrefix() + " "
			if p.Scope.Text != "" {
				ret += NewStyle(p.Scope.Style...).Sprint(" (" + p.Scope.Text + ") ")
			}
			ret += p.MessageStyle.Sprint(m)
		} else {
			ret += "\n" + p.Prefix.Style.Sprint(strings.Repeat(" ", len(p.Prefix.Text)+2)) + " " + p.MessageStyle.Sprint(m)
		}
	}

	return Sprint(ret)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p PrefixPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(a...) + "\n"
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p PrefixPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p PrefixPrinter) Print(a ...interface{}) GenericPrinter {
	Print(p.Sprint(a...))
	checkFatal(&p)
	return p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p PrefixPrinter) Println(a ...interface{}) GenericPrinter {
	Print(p.Sprintln(a...))
	checkFatal(&p)
	return p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p PrefixPrinter) Printf(format string, a ...interface{}) GenericPrinter {
	Print(Sprintf(format, a...))
	checkFatal(&p)
	return p
}

// GetFormattedPrefix returns the Prefix as a styled text string.
func (p PrefixPrinter) GetFormattedPrefix() string {
	return NewStyle(p.Prefix.Style...).Sprint(" " + p.Prefix.Text + " ")
}

// Prefix contains the data used as the beginning of a printed text via a PrefixPrinter.
type Prefix struct {
	Text  string
	Style Style
}

// Scope contains the data of the optional scope of a prefix.
// If it has a text, it will be printed after the Prefix in brackets.
type Scope struct {
	Text  string
	Style Style
}

func checkFatal(p *PrefixPrinter) {
	if p.Fatal {
		os.Exit(1)
	}
}
