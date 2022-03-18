package pterm

import (
	"fmt"
	"io"
)

var (
	// DefaultBasicText returns a default BasicTextPrinter, which can be used to print text as is.
	// No default style is present for BasicTextPrinter.
	DefaultBasicText = BasicTextPrinter{}
)

// BasicTextPrinter is the printer used to print the input as-is or as specified by user formatting.
type BasicTextPrinter struct {
	Style  *Style
	Writer io.Writer
}

// WithStyle adds a style to the printer.
func (p BasicTextPrinter) WithStyle(style *Style) *BasicTextPrinter {
	p.Style = style
	return &p
}

func (p BasicTextPrinter) WithWriter(writer io.Writer) *BasicTextPrinter {
	p.Writer = writer
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p BasicTextPrinter) Sprint(a ...interface{}) string {
	if p.Style == nil {
		p.Style = NewStyle()
	}
	return p.Style.Sprint(a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p BasicTextPrinter) Sprintln(a ...interface{}) string {
	str := fmt.Sprintln(a...)
	return Sprintln(p.Sprint(str))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p BasicTextPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p BasicTextPrinter) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to provided writer.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *BasicTextPrinter) Print(a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to provided writer.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *BasicTextPrinter) Println(a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintln(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to provided writer.
// It returns the number of bytes written and any write error encountered.
func (p *BasicTextPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to provided writer.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *BasicTextPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p *BasicTextPrinter) PrintOnError(a ...interface{}) *TextPrinter {
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
func (p *BasicTextPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
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
