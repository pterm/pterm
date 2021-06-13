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

	t.Run("SprintWithTitle", func(t *testing.T) {
		testSprintContains(t, func(a interface{}) string {
			return p.WithTitle("a").Sprint(a)
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

func TestBoxPrinter_WithTitle(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitle("-")

	assert.Equal(t, "-", p2.Title)
	assert.Empty(t, p.Title)
}

func TestBoxPrinter_WithTitleTopLeft(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleTopLeft()

	assert.Equal(t, true, p2.TitleTopLeft)
	assert.Equal(t, false, p.TitleTopLeft)
}

func TestBoxPrinter_WithTitleTopRight(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleTopRight()

	assert.Equal(t, true, p2.TitleTopRight)
	assert.Equal(t, false, p.TitleTopRight)
}

func TestBoxPrinter_WithTitleTopCenter(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleTopCenter()

	assert.Equal(t, true, p2.TitleTopCenter)
	assert.Equal(t, false, p.TitleTopCenter)
}

func TestBoxPrinter_WithTitleBottomRight(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleBottomRight()

	assert.Equal(t, true, p2.TitleBottomRight)
	assert.Equal(t, false, p.TitleBottomRight)
}

func TestBoxPrinter_WithTitleBottomLeft(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleBottomLeft()

	assert.Equal(t, true, p2.TitleBottomLeft)
	assert.Equal(t, false, p.TitleBottomLeft)
}

func TestBoxPrinter_WithTitleBottomCenter(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleBottomCenter()

	assert.Equal(t, true, p2.TitleBottomCenter)
	assert.Equal(t, false, p.TitleBottomCenter)
}

func TestBoxPrinter_WithTitleWithTitleBottomLeft(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleBottomLeft().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Equal(t, "\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39mLorem Ipsum\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39m\x1b[0m a \x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m", p2)
}

func TestBoxPrinter_WithTitleWithTitleTopLeft(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleTopLeft().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Equal(t, "\x1b[39m\x1b[0m\x1b[39m\x1b[0m a \x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39mLorem Ipsum\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m", p2)
}

func TestBoxPrinter_WithTitleWithTitleBottomRight(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleBottomRight().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Equal(t, "\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39mLorem Ipsum\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m a \x1b[39m\x1b[0m\x1b[39m\x1b[0m", p2)
}

func TestBoxPrinter_WithTitleWithTitleTopRight(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleTopRight().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Equal(t, "\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m a \x1b[39m\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39mLorem Ipsum\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m", p2)
}

func TestBoxPrinter_WithTitleWithTitleTopCenter(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleTopCenter().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Equal(t, "\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m a \x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39mLorem Ipsum\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m", p2)
}

func TestBoxPrinter_WithTitleWithTitleBottomCenter(t *testing.T) {
	p := BoxPrinter{}
	p2 := p.WithTitleBottomCenter().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Equal(t, "\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39mLorem Ipsum\x1b[0m\x1b[39m\x1b[0m\n\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m a \x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m\x1b[39m\x1b[0m", p2)
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
