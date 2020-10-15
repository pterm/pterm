package pterm

import (
	"fmt"
	"os"

	"github.com/pterm/pterm/internal"
)

var (
	// MessageVanilla returns a default MessagePrinter, which can be used to print text as is
	// No default style is present for MessagePrinter
	MessageVanilla = MessagePrinter{
		Style: nil,
	}
)

// MessagePrinter is the printer used to print the input as-is or as specified by user formatting
type MessagePrinter struct {
	Style Style
	Fatal bool
}

// WithStyle adds a style to the printer.
// unlike prefix printer, there is a single style for the message=
func (p MessagePrinter) WithStyle(style Style) *MessagePrinter {
	p.Style = style
	return &p
}

// WithFatal sets if the printer should panic after printing.
// NOTE:
// The printer will only panic if either MessagePrinter.Println, MessagePrinter.Print
// or MessagePrinter.Printf is called.
func (p MessagePrinter) WithFatal(b ...bool) *MessagePrinter {
	p.Fatal = internal.WithBoolean(b)
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p MessagePrinter) Sprint(a ...interface{}) string {
	if p.Style == nil {
		return fmt.Sprint(a...)
	}
	return p.Style.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p MessagePrinter) Sprintf(format string, a ...interface{}) string {
	if p.Style == nil {
		return fmt.Sprintf(format, a...)
	}
	return p.Style.Sprintf(format, a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p MessagePrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p MessagePrinter) Print(a ...interface{}) TextPrinter {
	if p.Style == nil {
		fmt.Print(p.Sprint(a...))
	} else {
		Print(p.Sprint(a...))
	}
	if p.Fatal {
		os.Exit(1)
	}
	return p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p MessagePrinter) Println(a ...interface{}) TextPrinter {
	if p.Style == nil {
		fmt.Print(p.Sprintln(a...))
	} else {
		Print(p.Sprintln(a...))
	}
	if p.Fatal {
		os.Exit(1)
	}
	return p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p MessagePrinter) Printf(format string, a ...interface{}) TextPrinter {
	if p.Style == nil {
		fmt.Print(p.Sprintf(format, a...))
	} else {
		Print(p.Sprintf(format, a...))
	}
	if p.Fatal {
		os.Exit(1)
	}
	return p
}
