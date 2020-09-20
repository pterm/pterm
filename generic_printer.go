package pterm

type GenericPrinter interface {
	// Sprint formats using the default formats for its operands and returns the resulting string.
	// Spaces are added between operands when neither is a string.
	Sprint(a ...interface{}) string

	// Sprintln formats using the default formats for its operands and returns the resulting string.
	// Spaces are always added between operands and a newline is appended.
	Sprintln(a ...interface{}) string

	// Sprintf formats according to a format specifier and returns the resulting string.
	Sprintf(format string, a ...interface{}) string

	// Print formats using the default formats for its operands and writes to standard output.
	// Spaces are added between operands when neither is a string.
	// It returns the number of bytes written and any write error encountered.
	Print(a ...interface{}) GenericPrinter

	// Println formats using the default formats for its operands and writes to standard output.
	// Spaces are always added between operands and a newline is appended.
	// It returns the number of bytes written and any write error encountered.
	Println(a ...interface{}) GenericPrinter

	// Printf formats according to a format specifier and writes to standard output.
	// It returns the number of bytes written and any write error encountered.
	Printf(format string, a ...interface{}) GenericPrinter
}
