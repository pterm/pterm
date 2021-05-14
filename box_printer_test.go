package pterm

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoxPrinterNilPrint(t *testing.T) {
	p := BoxPrinter{}
	p.Println("Hello, World!")
}

func TestBoxPrinterPrintMethods(t *testing.T) {
	p := DefaultBox

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a interface{}) {
			p.Print(a)
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

	t.Run("SprintMultipleLines", func(t *testing.T) {
		testSprintContains(t, func(a interface{}) string {
			return p.Sprint("testing\ntesting2" + Sprint(a))
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

func TestBoxPrinter_WithBottomLeftCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomLeftCornerString("-")

	assert.Equal(t, "-", p2.BottomLeftCornerString)
	assert.Empty(t, p.BottomLeftCornerString)
}

func TestBoxPrinter_WithBottomPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomPadding(5)

	assert.Equal(t, 5, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestBoxPrinter_WithBottomRightCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomRightCornerString("-")

	assert.Equal(t, "-", p2.BottomRightCornerString)
	assert.Empty(t, p.BottomRightCornerString)
}

func TestBoxPrinter_WithBoxStyle(t *testing.T) {
	p := BoxPrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithBoxStyle(s)

	assert.Equal(t, s, p2.BoxStyle)
	assert.Empty(t, p.BoxStyle)
}

func TestBoxPrinter_WithLeftPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithLeftPadding(5)

	assert.Equal(t, 5, p2.LeftPadding)
	assert.Empty(t, p.LeftPadding)
}

func TestBoxPrinter_WithRightPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithRightPadding(5)

	assert.Equal(t, 5, p2.RightPadding)
	assert.Empty(t, p.RightPadding)
}

func TestBoxPrinter_WithTextStyle(t *testing.T) {
	p := BoxPrinter{}
	s := NewStyle(FgRed, BgRed, Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Empty(t, p.TextStyle)
}

func TestBoxPrinter_WithTopLeftCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopLeftCornerString("-")

	assert.Equal(t, "-", p2.TopLeftCornerString)
	assert.Empty(t, p.TopLeftCornerString)
}

func TestBoxPrinter_WithTopPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopPadding(5)

	assert.Equal(t, 5, p2.TopPadding)
	assert.Empty(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidTopPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopPadding(-5)

	assert.Equal(t, 0, p2.TopPadding)
	assert.Empty(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidBottomPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithBottomPadding(-5)

	assert.Equal(t, 0, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestBoxPrinter_WithInvalidLeftPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithLeftPadding(-5)

	assert.Equal(t, 0, p2.LeftPadding)
	assert.Empty(t, p.LeftPadding)
}

func TestBoxPrinter_WithInvalidRightPadding(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithRightPadding(-5)

	assert.Equal(t, 0, p2.RightPadding)
	assert.Empty(t, p.RightPadding)
}

func TestBoxPrinter_WithTopRightCornerString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTopRightCornerString("-")

	assert.Equal(t, "-", p2.TopRightCornerString)
	assert.Empty(t, p.TopRightCornerString)
}

func TestBoxPrinter_WithVerticalString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithVerticalString("-")

	assert.Equal(t, "-", p2.VerticalString)
	assert.Empty(t, p.VerticalString)
}

func TestBoxPrinter_WithHorizontalString(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithHorizontalString("-")

	assert.Equal(t, "-", p2.HorizontalString)
	assert.Empty(t, p.HorizontalString)
}
