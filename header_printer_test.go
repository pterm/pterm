package pterm

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeaderPrinterNilPrint(t *testing.T) {
	p := HeaderPrinter{}
	p.Println("Hello, World!")
}

func TestHeaderPrinterPrintMethods(t *testing.T) {
	p := DefaultHeader

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
		})
	})

	t.Run("PrintWithFullWidth", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a interface{}) {
			p2 := p.WithFullWidth()
			p2.Print(a)
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

func TestHeaderPrinter_WithBackgroundStyle(t *testing.T) {
	s := NewStyle(FgRed, BgGray, Bold)
	p := HeaderPrinter{}
	p2 := p.WithBackgroundStyle(s)

	assert.Equal(t, s, p2.BackgroundStyle)
}

func TestHeaderPrinter_WithFullWidth(t *testing.T) {
	p := HeaderPrinter{}
	p2 := p.WithFullWidth()

	assert.Equal(t, true, p2.FullWidth)
}

func TestHeaderPrinter_WithFullWidthToLongForTerminal(t *testing.T) {
	p := HeaderPrinter{}
	p2 := p.WithFullWidth().Sprint("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	assert.Equal(t, "                                                                                \naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa                        \n                                                                                \n", p2)
}

func TestHeaderPrinter_ToLongForTerminal(t *testing.T) {
	p := HeaderPrinter{}
	p2 := p.Sprint("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")

	assert.Equal(t, "                                                                                \naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa                        \n                                                                                \n", p2)
}

func TestHeaderPrinter_WithMargin(t *testing.T) {
	p := HeaderPrinter{}
	p2 := p.WithMargin(1337)

	assert.Equal(t, 1337, p2.Margin)
}

func TestHeaderPrinter_WithTextStyle(t *testing.T) {
	s := NewStyle(FgRed, BgGray, Bold)
	p := HeaderPrinter{}
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
}
