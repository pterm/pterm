package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestHeaderPrinterNilPrint(t *testing.T) {
	p := pterm.HeaderPrinter{}
	p.Println("Hello, World!")
}

func TestHeaderPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultHeader

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("PrintWithFullWidth", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a any) {
			p2 := p.WithFullWidth()
			p2.Print(a)
		})
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

func TestHeaderPrinter_WithBackgroundStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgGray, pterm.Bold)
	p := pterm.HeaderPrinter{}
	p2 := p.WithBackgroundStyle(s)

	testza.AssertEqual(t, s, p2.BackgroundStyle)
}

func TestHeaderPrinter_WithFullWidth(t *testing.T) {
	p := pterm.HeaderPrinter{}
	p2 := p.WithFullWidth()

	testza.AssertEqual(t, true, p2.FullWidth)
}

func TestHeaderPrinter_WithFullWidthToLongForTerminal(t *testing.T) {
	p := pterm.HeaderPrinter{}
	p2 := p.WithFullWidth().Sprint("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	testza.AssertContains(t, p2, "a")
}

func TestHeaderPrinter_ToLongForTerminal(t *testing.T) {
	p := pterm.HeaderPrinter{}
	p2 := p.Sprint("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	testza.AssertContains(t, p2, "a")
}

func TestHeaderPrinter_WithMargin(t *testing.T) {
	p := pterm.HeaderPrinter{}
	p2 := p.WithMargin(1337)

	testza.AssertEqual(t, 1337, p2.Margin)
}

func TestHeaderPrinter_WithTextStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgGray, pterm.Bold)
	p := pterm.HeaderPrinter{}
	p2 := p.WithTextStyle(s)

	testza.AssertEqual(t, s, p2.TextStyle)
}

func TestHeaderPrinter_WithWriter(t *testing.T) {
	p := pterm.HeaderPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}
