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
func (p TemplatePrinter) Sprint(a ...interface{}) (string, error) {
	panic("write printer code here")
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p TemplatePrinter) Sprintln(a ...interface{}) (string, error) {
    text, err := p.Sprint(a...)
	return Sprintln(text), err
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p TemplatePrinter) Sprintf(format string, a ...interface{}) (string, error) {
    text, err := Sprintf(format, a...)
	return p.Sprint(text), err
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Print(a ...interface{}) *TemplatePrinter {
	text, err := p.Sprint(a...)
    if err != nil {
            return err
        }
    Print(text)
	tp := TemplatePrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Println(a ...interface{}) *TemplatePrinter {
	text, err := p.Sprint(a...)
    if err != nil {
            return err
        }
    Println(text)
    tp := TemplatePrinter(p)
    return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p TemplatePrinter) Printf(format string, a ...interface{}) *TemplatePrinter {
	text, err := p.Sprintf(format, a...)
    if err != nil {
        return err
    }
    Print(text)
	tp := TemplatePrinter(p)
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
	return "", nil
}

// Render prints the Template to the terminal.
func (p TemplatePrinter) Render() {
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
// Start the spinner.
func (s Spinner) Start(text...interface{}) *Template { // TODO: Replace Template with actual printer.
	// TODO: start logic
	return &s
}

// Stop terminates the Spinner immediately.
// The Spinner will not resolve into anything.
func (s *Spinner) Stop() {
	// TODO: stop logic
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (s *Spinner) GenericStart() *LivePrinter {
	s.Start()
	lp := LivePrinter(s)
	return &lp
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (s *Spinner) GenericStop() *LivePrinter {
	s.Stop()
	lp := LivePrinter(s)
	return &lp
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
		internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		internal.TestPrintfContains(t, func(w io.Writer, format string, a interface{}) {
			p.Printf(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		internal.TestPrintlnContains(t, func(w io.Writer, a interface{}) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		internal.TestSprintContains(t, func(a interface{}) string {
			return p.Sprint(a)
		})
	})

	t.Run("Sprintf", func(t *testing.T) {
		internal.TestSprintfContains(t, func(format string, a interface{}) string {
			return p.Sprintf(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		internal.TestSprintlnContains(t, func(a interface{}) string {
			return p.Sprintln(a)
		})
	})
}
```
