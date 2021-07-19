package pterm_test

import (
	"errors"
	"io"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestSectionPrinterNilPrint(t *testing.T) {
	p := pterm.SectionPrinter{}
	p.Println("Hello, World!")
}

func TestSectionPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultSection

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

	t.Run("Printfln", func(t *testing.T) {
		testPrintflnContains(t, func(w io.Writer, format string, a interface{}) {
			p.Printfln(format, a)
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

	t.Run("Sprintfln", func(t *testing.T) {
		testSprintflnContains(t, func(format string, a interface{}) string {
			return p.Sprintfln(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		testSprintlnContains(t, func(a interface{}) string {
			return p.Sprintln(a)
		})
	})

	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(nil)
		})
		assert.Empty(t, result)
	})
}

func TestSectionPrinter_WithBottomPadding(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithBottomPadding(1337)

	assert.Equal(t, 1337, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestSectionPrinter_WithLevel(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithLevel(1337)

	assert.Equal(t, 1337, p2.Level)
	assert.Empty(t, p.Level)
}

func TestSectionPrinter_WithStyle(t *testing.T) {
	p := pterm.SectionPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Empty(t, p.Style)
}

func TestSectionPrinter_WithTopPadding(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithTopPadding(1337)

	assert.Equal(t, 1337, p2.TopPadding)
	assert.Empty(t, p.TopPadding)
}

func TestSectionPrinter_WithIndentCharacter(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithIndentCharacter("#")

	assert.Equal(t, "#", p2.IndentCharacter)
	assert.Empty(t, p.IndentCharacter)
}
