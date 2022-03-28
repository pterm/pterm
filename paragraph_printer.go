package pterm

import (
	"fmt"
	"io"
	"strings"
)

// DefaultParagraph contains the default values for a ParagraphPrinter.
var DefaultParagraph = ParagraphPrinter{
	MaxWidth: GetTerminalWidth(),
}

// ParagraphPrinter can print paragraphs to a fixed line width.
// The text will split between words, so that words will stick together.
// It's like in a book.
type ParagraphPrinter struct {
	MaxWidth int
	Writer   io.Writer
}

// WithMaxWidth returns a new ParagraphPrinter with a specific MaxWidth
func (p ParagraphPrinter) WithMaxWidth(width int) *ParagraphPrinter {
	p.MaxWidth = width
	return &p
}

// WithWriter sets the custom Writer.
func (p ParagraphPrinter) WithWriter(writer io.Writer) *ParagraphPrinter {
	p.Writer = writer
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p ParagraphPrinter) Sprint(a ...interface{}) string {
	if RawOutput {
		return Sprint(a...)
	}

	words := strings.Fields(strings.TrimSpace(Sprint(a...)))
	if len(words) == 0 {
		return ""
	}
	wrapped := words[0]
	spaceLeft := p.MaxWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = p.MaxWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}

	return wrapped
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p ParagraphPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(Sprintln(a...)) + "\n"
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p ParagraphPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p ParagraphPrinter) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *ParagraphPrinter) Print(a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *ParagraphPrinter) Println(a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintln(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *ParagraphPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *ParagraphPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	Fprint(p.Writer, p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p *ParagraphPrinter) PrintOnError(a ...interface{}) *TextPrinter {
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
func (p *ParagraphPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
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
