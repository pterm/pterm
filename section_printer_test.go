package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestSectionPrinterNilPrint(t *testing.T) {
	p := SectionPrinter{}
	p.Println("Hello, World!")
}

func TestSectionPrinter_Print(t *testing.T) {
	p := DefaultSection
	internal.TestPrintContains(t, func(w io.Writer, a string) {
		p.Print(a)
	})
}

func TestSectionPrinter_Printf(t *testing.T) {
	p := DefaultSection
	internal.TestPrintfContains(t, func(w io.Writer, format string, a string) {
		p.Printf(format, a)
	})
}

func TestSectionPrinter_Println(t *testing.T) {
	p := DefaultSection
	internal.TestPrintlnContains(t, func(w io.Writer, a string) {
		p.Println(a)
	})
}

func TestSectionPrinter_Sprint(t *testing.T) {
	p := DefaultSection
	internal.TestSprintContains(t, func(a string) string {
		return p.Sprint(a)
	})
}

func TestSectionPrinter_Sprintf(t *testing.T) {
	p := DefaultSection
	internal.TestSprintfContains(t, func(format string, a string) string {
		return p.Sprintf(format, a)
	})
}

func TestSectionPrinter_Sprintln(t *testing.T) {
	p := DefaultSection
	internal.TestSprintlnContains(t, func(a string) string {
		return p.Sprintln(a)
	})
}

func TestSectionPrinter_WithBottomPadding(t *testing.T) {
	p := SectionPrinter{}
	p2 := p.WithBottomPadding(1337)

	assert.Equal(t, 1337, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestSectionPrinter_WithLevel(t *testing.T) {
	p := SectionPrinter{}
	p2 := p.WithLevel(1337)

	assert.Equal(t, 1337, p2.Level)
	assert.Empty(t, p.Level)
}

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
