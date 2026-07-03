package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestBoxPrinterNilPrint(_ *testing.T) {
	p := pterm.BoxPrinter{}
	p.Println("Hello, World!")
}

func TestBoxPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultBox

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(_ io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(_ io.Writer, format string, a any) {
			p.Printf(format, a)
		})
	})

	t.Run("Printfln", func(t *testing.T) {
		testPrintflnContains(t, func(_ io.Writer, format string, a any) {
			p.Printfln(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(_ io.Writer, a any) {
			p.Println(a)
		})
	})

	t.Run("Sprint", func(t *testing.T) {
		testSprintContains(t, func(a any) string {
			return p.Sprint(a)
		})
	})

	t.Run("SprintWithTitle", func(t *testing.T) {
		testSprintContains(t, func(a any) string {
			return p.WithTitle("a").Sprint(a)
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

	t.Run("SprintMultipleLines", func(t *testing.T) {
		testSprintContains(t, func(a any) string {
			return p.Sprint("testing\ntesting2" + pterm.Sprint(a))
		})
	})

	t.Run("PrintOnError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnError(errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnError(nil)
		})
		assert.Zero(t, result)
	})

	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		assert.Contains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
		result := captureStdout(func(_ io.Writer) {
			p.PrintOnErrorf("", nil)
		})
		assert.Zero(t, result)
	})
}

func TestBoxPrinter_WithBottomLeftCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomLeftCornerString("-")

	assert.Equal(t, "-", p2.BottomLeftCornerString)
	assert.Zero(t, p.BottomLeftCornerString)
}

func TestBoxPrinter_WithBottomPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomPadding(5)

	assert.Equal(t, 5, p2.BottomPadding)
	assert.Zero(t, p.BottomPadding)
}

func TestBoxPrinter_WithBottomRightCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomRightCornerString("-")

	assert.Equal(t, "-", p2.BottomRightCornerString)
	assert.Zero(t, p.BottomRightCornerString)
}

func TestBoxPrinter_WithTitle(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitle("-")

	assert.Equal(t, "-", p2.Title)
	assert.Zero(t, p.Title)
}

func TestBoxPrinter_WithTitleTopLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopLeft()

	assert.Equal(t, true, p2.TitleTopLeft)
	assert.Equal(t, false, p.TitleTopLeft)
}

func TestBoxPrinter_WithTitleTopRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopRight()

	assert.Equal(t, true, p2.TitleTopRight)
	assert.Equal(t, false, p.TitleTopRight)
}

func TestBoxPrinter_WithTitleTopCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopCenter()

	assert.Equal(t, true, p2.TitleTopCenter)
	assert.Equal(t, false, p.TitleTopCenter)
}

func TestBoxPrinter_WithTitleBottomRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomRight()

	assert.Equal(t, true, p2.TitleBottomRight)
	assert.Equal(t, false, p.TitleBottomRight)
}

func TestBoxPrinter_WithTitleBottomLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomLeft()

	assert.Equal(t, true, p2.TitleBottomLeft)
	assert.Equal(t, false, p.TitleBottomLeft)
}

func TestBoxPrinter_WithTitleBottomCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomCenter()

	assert.Equal(t, true, p2.TitleBottomCenter)
	assert.Equal(t, false, p.TitleBottomCenter)
}

func TestBoxPrinter_WithTitleWithTitleBottomLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomLeft().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Contains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopLeft().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Contains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleBottomRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomRight().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Contains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopRight().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Contains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopCenter().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Contains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleBottomCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomCenter().WithTitle("a").Sprint("Lorem Ipsum")

	assert.Contains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithBoxStyle(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithBoxStyle(s)

	assert.Equal(t, s, p2.BoxStyle)
	assert.Zero(t, p.BoxStyle)
}

func TestBoxPrinter_WithLeftPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithLeftPadding(5)

	assert.Equal(t, 5, p2.LeftPadding)
	assert.Zero(t, p.LeftPadding)
}

func TestBoxPrinter_WithRightPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithRightPadding(5)

	assert.Equal(t, 5, p2.RightPadding)
	assert.Zero(t, p.RightPadding)
}

func TestBoxPrinter_WithTextStyle(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Zero(t, p.TextStyle)
}

func TestBoxPrinter_WithTopLeftCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopLeftCornerString("-")

	assert.Equal(t, "-", p2.TopLeftCornerString)
	assert.Zero(t, p.TopLeftCornerString)
}

func TestBoxPrinter_WithTopPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopPadding(5)

	assert.Equal(t, 5, p2.TopPadding)
	assert.Zero(t, p.TopPadding)
}

func TestBoxPrinter_WithHorizontalPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithHorizontalPadding(5)

	assert.Equal(t, 5, p2.LeftPadding)
	assert.Equal(t, 5, p2.RightPadding)
	assert.Equal(t, 0, p.LeftPadding)
	assert.Equal(t, 0, p.RightPadding)
}

func TestBoxPrinter_WithVerticalPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithVerticalPadding(5)

	assert.Equal(t, 5, p2.TopPadding)
	assert.Equal(t, 5, p2.BottomPadding)
	assert.Equal(t, 0, p.TopPadding)
	assert.Equal(t, 0, p.BottomPadding)
}

func TestBoxPrinter_WithPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithPadding(5)

	assert.Equal(t, 5, p2.TopPadding)
	assert.Equal(t, 5, p2.BottomPadding)
	assert.Equal(t, 5, p2.LeftPadding)
	assert.Equal(t, 5, p2.RightPadding)
	assert.Equal(t, 0, p.TopPadding)
	assert.Equal(t, 0, p.BottomPadding)
	assert.Equal(t, 0, p.LeftPadding)
	assert.Equal(t, 0, p.RightPadding)
}

func TestBoxPrinter_WithInvalidTopPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopPadding(-5)

	assert.Equal(t, 0, p2.TopPadding)
	assert.Zero(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidBottomPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomPadding(-5)

	assert.Equal(t, 0, p2.BottomPadding)
	assert.Zero(t, p.BottomPadding)
}

func TestBoxPrinter_WithInvalidLeftPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithLeftPadding(-5)

	assert.Equal(t, 0, p2.LeftPadding)
	assert.Zero(t, p.LeftPadding)
}

func TestBoxPrinter_WithInvalidRightPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithRightPadding(-5)

	assert.Equal(t, 0, p2.RightPadding)
	assert.Zero(t, p.RightPadding)
}

func TestBoxPrinter_WithTopRightCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopRightCornerString("-")

	assert.Equal(t, "-", p2.TopRightCornerString)
	assert.Zero(t, p.TopRightCornerString)
}

func TestBoxPrinter_WithVerticalString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithVerticalString("-")

	assert.Equal(t, "-", p2.VerticalString)
	assert.Zero(t, p.VerticalString)
}

func TestBoxPrinter_WithHorizontalString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithHorizontalString("-")

	assert.Equal(t, "-", p2.HorizontalString)
	assert.Zero(t, p.HorizontalString)
}

func TestBoxPrinter_WithWriter(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}
