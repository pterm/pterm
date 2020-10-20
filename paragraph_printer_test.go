package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestParagraphPrinter_Print(t *testing.T) {
	p := DefaultParagraph
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		p.Print(a)
	})
}

func TestParagraphPrinter_Printf(t *testing.T) {
	p := DefaultParagraph
	internal.TestPrintfContains(t, func(w io.Writer, format string, a interface{}) {
		p.Printf(format, a)
	})
}

func TestParagraphPrinter_Println(t *testing.T) {
	p := DefaultParagraph
	internal.TestPrintlnContains(t, func(w io.Writer, a interface{}) {
		p.Println(a)
	})
}

func TestParagraphPrinter_Sprint(t *testing.T) {
	p := DefaultParagraph
	internal.TestSprintContains(t, func(a interface{}) string {
		return p.Sprint(a)
	})
}

func TestParagraphPrinter_Sprintf(t *testing.T) {
	p := DefaultParagraph
	internal.TestSprintfContains(t, func(format string, a interface{}) string {
		return p.Sprintf(format, a)
	})
}

func TestParagraphPrinter_Sprintln(t *testing.T) {
	p := DefaultParagraph
	internal.TestSprintlnContains(t, func(a interface{}) string {
		return p.Sprintln(a)
	})
}

func TestParagraphPrinter_WithMaxWidth(t *testing.T) {
	p := ParagraphPrinter{}
	p2 := p.WithMaxWidth(1337)

	assert.Equal(t, 1337, p2.MaxWidth)
}
