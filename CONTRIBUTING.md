# Contributing to PTerm

> This document explains how to participate in the development of PTerm.\
If your goal is to report a bug instead of programming PTerm, you can do so [here](https://github.com/pterm/pterm/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc).

## Creating a new printer

> In this chapter we will show you how to create a new printer.

### `TextPrinter` Template
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
func (p TemplatePrinter) Print(a ...interface{}) *TextPrinter {
	Print(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Println(a ...interface{}) *TextPrinter {
	Println(p.Sprint(a...))
    tp := TextPrinter(p)
    return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}
```

### `RenderablePrinter` Template

```go
package pterm

type TemplatePrinter struct{
	// TODO: Add printer settings here
}

// Srender renders the Template as a string.
func (p TemplatePrinter) Srender() (string, error) {
	var ret string

    return ret, nil
}

// Render prints the Template to the terminal.
func (p TemplatePrinter) Render() error {
	s, err := p.Srender()
    if err != nil {
        return err
    }
    Println(s)

    return nil
}
```

### `LivePrinter` Template

```go
// Start the TemplatePrinter.
package pterm
import "github.com/pterm/pterm"

type TemplatePrinter struct{

}


func (s TemplatePrinter) Start(text...interface{}) (*TemplatePrinter, error) { // TODO: Replace Template with actual printer.
	// TODO: start logic
	return &s, nil
}

// Stop terminates the TemplatePrinter immediately.
// The TemplatePrinter will not resolve into anything.
func (s *TemplatePrinter) Stop() error {
	// TODO: stop logic
    return nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (s *TemplatePrinter) GenericStart() (*LivePrinter, error) {
	_, err := s.Start()
	lp := LivePrinter(s)
	return &lp, err
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (s *TemplatePrinter) GenericStop() (*LivePrinter, error) {
	err := s.Stop()
	lp := LivePrinter(s)
	return &lp, err
}
```

## Writing Tests

> Each method of PTerm must be tested.

### Required tests for every printer

#### Nil Check

> This ensures that a printer without set values will not produce errors.

```go
func TestTemplatePrinterNilPrint(t *testing.T) { // TODO: Replace "Template" with actual printer name.
	p := TemplatePrinter{} // TODO: Replace "Template" with actual printer name.
	p.Println("Hello, World!")
}
```

#### `WithXxx()` Methods

> Each method, which starts with `With` can be tested by checking if it actually creates a new printer and sets the value.

Example from `SectionPrinter`:

```go
func TestSectionPrinter_WithStyle(t *testing.T) {
	p := SectionPrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Empty(t, p.Style)
}

func TestSectionPrinter_WithTopPadding(t *testing.T) {
	p := SectionPrinter{}
	p2 := p.WithTopPadding(1337)

	assert.Equal(t, 1337, p2.TopPadding)
	assert.Empty(t, p.TopPadding)
}
```

### `TextPrinter` Tests Template

```go
func TestTemplatePrinterPrintMethods(t *testing.T) { // TODO: Replace "Template" with actual printer name.
	p := DefaultTemplate // TODO: Replace "Template" with actual printer name.

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(w io.Writer, format string, a interface{}) {
			p.Printf(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(w io.Writer, a interface{}) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		testSprintContains(t, func(a interface{}) string {
			return p.Sprint(a)
		})
	})

	t.Run("Sprintf", func(t *testing.T) {
		testSprintfContains(t, func(format string, a interface{}) string {
			return p.Sprintf(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		testSprintlnContains(t, func(a interface{}) string {
			return p.Sprintln(a)
		})
	})
}
```
