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

func TestSectionPrinterPrintMethods(t *testing.T) {
	p := DefaultSection

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

	t.Run("Printfln", func(t *testing.T) {
		internal.TestPrintflnContains(t, func(w io.Writer, format string, a interface{}) {
			p.Printfln(format, a)
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

	t.Run("Sprintfln", func(t *testing.T) {
		internal.TestSprintflnContains(t, func(format string, a interface{}) string {
			return p.Sprintfln(format, a)
		})
	})

	t.Run("Sprintln", func(t *testing.T) {
		internal.TestSprintlnContains(t, func(a interface{}) string {
			return p.Sprintln(a)
		})
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

func TestSectionPrinter_WithIndentCharacter(t *testing.T) {
	p := SectionPrinter{}
	p2 := p.WithIndentCharacter("#")

	assert.Equal(t, "#", p2.IndentCharacter)
	assert.Empty(t, p.IndentCharacter)
}
