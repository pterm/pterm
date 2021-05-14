package pterm

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParagraphPrinterNilPrint(t *testing.T) {
	p := ParagraphPrinter{}
	p.Println("Hello, World!")
}

func TestParagraphPrinterPrintMethods(t *testing.T) {
	p := DefaultParagraph

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
		})
	})

	t.Run("PrintWithLongText", func(t *testing.T) {
		proxyToDevNull()
		assert.NotEmpty(t, p.Print("This is a longer text to test the paragraph printer. I don't know when this text will be long enough so I will just write until I get the feeling that it's enough. Maybe about now."))
	})

	t.Run("PrintWithoutText", func(t *testing.T) {
		proxyToDevNull()
		assert.NotEmpty(t, p.Print(""))
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

func TestParagraphPrinter_WithMaxWidth(t *testing.T) {
	p := ParagraphPrinter{}
	p2 := p.WithMaxWidth(1337)

	assert.Equal(t, 1337, p2.MaxWidth)
}
