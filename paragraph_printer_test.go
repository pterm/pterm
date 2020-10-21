package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestParagraphPrinterNilPrint(t *testing.T) {
	p := ParagraphPrinter{}
	p.Println("Hello, World!")
}

func TestParagraphPrinterPrintMethods(t *testing.T) {
	p := DefaultParagraph

	t.Run("Print", func(t *testing.T) {
		internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
		})
	})

	t.Run("PrintWithLongText", func(t *testing.T) {
		assert.NotEmpty(t, p.Print("This is a longer text to test the paragraph printer. I don't know when this text will be long enough so I will just write until I get the feeling that it's enough. Maybe about now."))
	})

	t.Run("PrintWithoutText", func(t *testing.T) {
		assert.NotEmpty(t, p.Print(""))
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

func TestParagraphPrinter_WithMaxWidth(t *testing.T) {
	p := ParagraphPrinter{}
	p2 := p.WithMaxWidth(1337)

	assert.Equal(t, 1337, p2.MaxWidth)
}
