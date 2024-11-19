package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestParagraphPrinterNilPrint(t *testing.T) {
	p := pterm.ParagraphPrinter{}
	p.Println("Hello, World!")
}

func TestParagraphPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultParagraph

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("PrintWithLongText", func(t *testing.T) {
		proxyToDevNull()
		testza.AssertNotZero(t, p.Print("This is a longer text to test the paragraph printer. I don't know when this text will be long enough so I will just write until I get the feeling that it's enough. Maybe about now."))
	})

	t.Run("PrintWithoutText", func(t *testing.T) {
		proxyToDevNull()
		testza.AssertNotZero(t, p.Print(""))
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(w io.Writer, format string, a any) {
			p.Printf(format, a)
		})
	})

	t.Run("Printfln", func(t *testing.T) {
		testPrintflnContains(t, func(w io.Writer, format string, a any) {
			p.Printfln(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(w io.Writer, a any) {
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
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(errors.New("hello world"))
		})
		testza.AssertContains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(nil)
		})
		testza.AssertZero(t, result)
	})

	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		testza.AssertContains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnErrorf("", nil)
		})
		testza.AssertZero(t, result)
	})
}

func TestParagraphPrinter_WithMaxWidth(t *testing.T) {
	p := pterm.ParagraphPrinter{}
	p2 := p.WithMaxWidth(1337)

	testza.AssertEqual(t, 1337, p2.MaxWidth)
}

func TestParagraphPrinter_WithWriter(t *testing.T) {
	p := pterm.ParagraphPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}
