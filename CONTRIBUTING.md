# Contributing to PTerm

## Template Snippets

### GenericPrinter Template

```go
package pterm

type TemplatePrinter struct{
	// TODO: Add printer settings here
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p TemplatePrinter) Sprint(a ...interface{}) string {
	panic("write printer code here")
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p TemplatePrinter) Sprintln(a ...interface{}) string {
	return Sprintln(p.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p TemplatePrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Print(a ...interface{}) TextPrinter {
	Print(p.Sprint(a...))
	return &p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Println(a ...interface{}) TextPrinter {
	Println(p.Sprint(a...))
	return &p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Printf(format string, a ...interface{}) TextPrinter {
	Print(p.Sprintf(format, a...))
	return &p
}
```