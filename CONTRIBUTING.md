# Contributing to PTerm

## Writing Tests

This is the base of every printer test

```go
func TestTemplatePrinterNilPrint(t *testing.T) {
	p := TemplatePrinter{}
	p.Println("Hello, World!")
}

func TestTemplatePrinterPrintMethods(t *testing.T) {
	p := DefaultTemplate

	t.Run("Print", func(t *testing.T) {
		internal.TestPrintContains(t, func(w io.Writer, a string) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		internal.TestPrintfContains(t, func(w io.Writer, format string, a string) {
			p.Printf(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		internal.TestPrintlnContains(t, func(w io.Writer, a string) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		internal.TestSprintContains(t, func(a string) string {
			return p.Sprint(a)
		})
	})

	t.Run("Sprintf", func(t *testing.T) {
		internal.TestSprintfContains(t, func(format string, a string) string {
			return p.Sprintf(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		internal.TestSprintlnContains(t, func(a string) string {
			return p.Sprintln(a)
		})
	})
}
```

## Template Snippets

### TextPrinter Template

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