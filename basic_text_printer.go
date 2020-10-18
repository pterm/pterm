package pterm

var (
	// DefaultBasic returns a default BasicTextPrinter, which can be used to print text as is
	// No default style is present for BasicTextPrinter
	DefaultBasic = BasicTextPrinter{
		Style: nil,
	}
)

// BasicTextPrinter is the printer used to print the input as-is or as specified by user formatting
type BasicTextPrinter struct {
	Style Style
}

// WithStyle adds a style to the printer.
// unlike prefix printer, there is a single style for the message=
func (p BasicTextPrinter) WithStyle(style Style) *BasicTextPrinter {
	p.Style = style
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p BasicTextPrinter) Sprint(a ...interface{}) string {
	return p.Style.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p BasicTextPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p BasicTextPrinter) Sprintln(a ...interface{}) string {
	return Sprintln(p.Sprint(a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p BasicTextPrinter) Print(a ...interface{}) TextPrinter {
	Print(p.Sprint(a...))
	return &p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p BasicTextPrinter) Println(a ...interface{}) TextPrinter {
	Println(p.Sprint(a...))
	return &p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p BasicTextPrinter) Printf(format string, a ...interface{}) TextPrinter {
	Print(p.Sprintf(format, a...))
	return &p
}
