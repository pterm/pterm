package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestSectionPrinterNilPrint(_ *testing.T) {
	p := pterm.SectionPrinter{}
	p.Println("Hello, World!")
}

func TestSectionPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultSection

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(_ io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(_ io.Writer, format string, a any) {
			p.Printf(format, a)
		})
	})

	t.Run("Printfln", func(t *testing.T) {
		testPrintflnContains(t, func(_ io.Writer, format string, a any) {
			p.Printfln(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(_ io.Writer, a any) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		testSprintContains(t, func(a any) string {
			return p.Sprint(a)
		})
	})

	t.Run("Sprintf", func(t *testing.T) {
		testSprintfContains(t, func(format string, a any) string {
			return p.Sprintf(format, a)
		})
	})

	t.Run("Sprintfln", func(t *testing.T) {
		testSprintflnContains(t, func(format string, a any) string {
			return p.Sprintfln(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		testSprintlnContains(t, func(a any) string {
			return p.Sprintln(a)
		})
	})

	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnError(nil)
		})
		assert.Zero(t, result)
	})

	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnErrorf("", nil)
		})
		assert.Zero(t, result)
	})
}

func TestSectionPrinter_WithBottomPadding(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithBottomPadding(1337)

	assert.Equal(t, 1337, p2.BottomPadding)
	assert.Zero(t, p.BottomPadding)
}

func TestSectionPrinter_WithLevel(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithLevel(1337)

	assert.Equal(t, 1337, p2.Level)
	assert.Zero(t, p.Level)
}

func TestSectionPrinter_WithStyle(t *testing.T) {
	p := pterm.SectionPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Zero(t, p.Style)
}

func TestSectionPrinter_WithTopPadding(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithTopPadding(1337)

	assert.Equal(t, 1337, p2.TopPadding)
	assert.Zero(t, p.TopPadding)
}

func TestSectionPrinter_WithIndentCharacter(t *testing.T) {
	p := pterm.SectionPrinter{}
	p2 := p.WithIndentCharacter("#")

	assert.Equal(t, "#", p2.IndentCharacter)
	assert.Zero(t, p.IndentCharacter)
}

func TestSectionPrinter_WithWriter(t *testing.T) {
	p := pterm.SectionPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}
