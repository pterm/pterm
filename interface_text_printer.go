package pterm

// TextPrinter contains methods to print formatted text to the console or return it as a string.
type TextPrinter interface {
	// Sprint formats using the default formats for its operands and returns the resulting string.
	// Spaces are added between operands when neither is a string.
	Sprint(a ...any) string

	// Sprintln formats using the default formats for its operands and returns the resulting string.
	// Spaces are always added between operands and a newline is appended.
	Sprintln(a ...any) string

	// Sprintf formats according to a format specifier and returns the resulting string.
	Sprintf(format string, a ...any) string

	// Sprintfln formats according to a format specifier and returns the resulting string.
	// Spaces are always added between operands and a newline is appended.
	Sprintfln(format string, a ...any) string

	// Print formats using the default formats for its operands and writes to standard output.
	// Spaces are added between operands when neither is a string.
	// It returns the number of bytes written and any write error encountered.
	Print(a ...any) *TextPrinter

	// Println formats using the default formats for its operands and writes to standard output.
	// Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.
	Println(a ...any) *TextPrinter

	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered.
	Printf(format string, a ...any) *TextPrinter

	// Printfln formats according to a format specifier and writes to standard output.
	// Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.
	Printfln(format string, a ...any) *TextPrinter

	// PrintOnError prints every error which is not nil.
	// If every error is nil, nothing will be printed.
	// This can be used for simple error checking.
	PrintOnError(a ...any) *TextPrinter

	// PrintOnErrorf wraps every error which is not nil and prints it.
	// If every error is nil, nothing will be printed.
	// This can be used for simple error checking.
	PrintOnErrorf(format string, a ...any) *TextPrinter
}
