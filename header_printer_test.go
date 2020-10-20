package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestHeaderPrinterNilPrint(t *testing.T) {
	p := HeaderPrinter{}
	p.Println("Hello, World!")
}

func TestHeaderPrinterPrintMethods(t *testing.T) {
	p := DefaultHeader

	t.Run("Print", func(t *testing.T) {
		internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
		})
	})

	t.Run("PrintWithFullWidth", func(t *testing.T) {
		internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
			p2 := p.WithFullWidth()
			p2.Print(a)
		})
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
