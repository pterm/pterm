package pterm

import (
	"fmt"
	"io"
	"strings"
)

// DefaultSection is the default section printer.
var DefaultSection = SectionPrinter{
	Style:           &ThemeDefault.SectionStyle,
	Level:           1,
	TopPadding:      1,
	BottomPadding:   1,
	IndentCharacter: "#",
}

// SectionPrinter prints a new section title.
// It can be used to structure longer text, or different chapters of your program.
type SectionPrinter struct {
	Style           *Style
	Level           int
	IndentCharacter string
	TopPadding      int
	BottomPadding   int
	Writer          io.Writer
}

// WithStyle returns a new SectionPrinter with a specific style.
func (p SectionPrinter) WithStyle(style *Style) *SectionPrinter {
	p.Style = style
	return &p
}

// WithLevel returns a new SectionPrinter with a specific level.
func (p SectionPrinter) WithLevel(level int) *SectionPrinter {
	p.Level = level
	return &p
}

// WithIndentCharacter returns a new SectionPrinter with a specific IndentCharacter.
func (p SectionPrinter) WithIndentCharacter(char string) *SectionPrinter {
	p.IndentCharacter = char
	return &p
}

// WithTopPadding returns a new SectionPrinter with a specific top padding.
func (p SectionPrinter) WithTopPadding(padding int) *SectionPrinter {
	p.TopPadding = padding
	return &p
}

// WithBottomPadding returns a new SectionPrinter with a specific top padding.
func (p SectionPrinter) WithBottomPadding(padding int) *SectionPrinter {
	p.BottomPadding = padding
	return &p
}

// WithWriter sets the custom Writer.
func (p SectionPrinter) WithWriter(writer io.Writer) *SectionPrinter {
	p.Writer = writer
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p SectionPrinter) Sprint(a ...interface{}) string {
	if p.Style == nil {
		p.Style = NewStyle()
	}

	var ret string

	for i := 0; i < p.TopPadding; i++ {
		ret += "\n"
	}

	if p.Level > 0 {
		ret += strings.Repeat(p.IndentCharacter, p.Level) + " "
	}

	ret += p.Style.Sprint(a...)

	for i := 0; i < p.BottomPadding; i++ {
		ret += "\n"
	}

	return ret
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p SectionPrinter) Sprintln(a ...interface{}) string {
	str := fmt.Sprintln(a...)
	return Sprint(p.Sprint(str))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p SectionPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p SectionPrinter) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Print(a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Println(a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintln(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p *SectionPrinter) PrintOnError(a ...interface{}) *TextPrinter {
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
func (p *SectionPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
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
