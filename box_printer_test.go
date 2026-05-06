package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/pterm/pterm/internal/testhelper"

	"github.com/pterm/pterm"
)

func TestBoxPrinterNilPrint(t *testing.T) {
	p := pterm.BoxPrinter{}
	p.Println("Hello, World!")
}

func TestBoxPrinterPrintMethods(t *testing.T) {
	p := pterm.DefaultBox

	t.Run("Print", func(t *testing.T) {
		testPrintContains(t, func(w io.Writer, a any) {
			p.Print(a)
		})
	})

	t.Run("Printf", func(t *testing.T) {
		testPrintfContains(t, func(w io.Writer, format string, a any) {
			p.Printf(format, a)
		})
	})

	t.Run("Printfln", func(t *testing.T) {
		testPrintflnContains(t, func(w io.Writer, format string, a any) {
			p.Printfln(format, a)
		})
	})

	t.Run("Println", func(t *testing.T) {
		testPrintlnContains(t, func(w io.Writer, a any) {
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
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(errors.New("hello world"))
		})
		testhelper.AssertContains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(nil)
		})
		testhelper.AssertZero(t, result)
	})

	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		testhelper.AssertContains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnErrorf("", nil)
		})
		testhelper.AssertZero(t, result)
	})
}

func TestBoxPrinter_WithBottomLeftCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomLeftCornerString("-")

	testhelper.AssertEqual(t, "-", p2.BottomLeftCornerString)
	testhelper.AssertZero(t, p.BottomLeftCornerString)
}

func TestBoxPrinter_WithBottomPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomPadding(5)

	testhelper.AssertEqual(t, 5, p2.BottomPadding)
	testhelper.AssertZero(t, p.BottomPadding)
}

func TestBoxPrinter_WithBottomRightCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomRightCornerString("-")

	testhelper.AssertEqual(t, "-", p2.BottomRightCornerString)
	testhelper.AssertZero(t, p.BottomRightCornerString)
}

func TestBoxPrinter_WithTitle(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitle("-")

	testhelper.AssertEqual(t, "-", p2.Title)
	testhelper.AssertZero(t, p.Title)
}

func TestBoxPrinter_WithTitleTopLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopLeft()

	testhelper.AssertEqual(t, true, p2.TitleTopLeft)
	testhelper.AssertEqual(t, false, p.TitleTopLeft)
}

func TestBoxPrinter_WithTitleTopRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopRight()

	testhelper.AssertEqual(t, true, p2.TitleTopRight)
	testhelper.AssertEqual(t, false, p.TitleTopRight)
}

func TestBoxPrinter_WithTitleTopCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopCenter()

	testhelper.AssertEqual(t, true, p2.TitleTopCenter)
	testhelper.AssertEqual(t, false, p.TitleTopCenter)
}

func TestBoxPrinter_WithTitleBottomRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomRight()

	testhelper.AssertEqual(t, true, p2.TitleBottomRight)
	testhelper.AssertEqual(t, false, p.TitleBottomRight)
}

func TestBoxPrinter_WithTitleBottomLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomLeft()

	testhelper.AssertEqual(t, true, p2.TitleBottomLeft)
	testhelper.AssertEqual(t, false, p.TitleBottomLeft)
}

func TestBoxPrinter_WithTitleBottomCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomCenter()

	testhelper.AssertEqual(t, true, p2.TitleBottomCenter)
	testhelper.AssertEqual(t, false, p.TitleBottomCenter)
}

func TestBoxPrinter_WithTitleWithTitleBottomLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomLeft().WithTitle("a").Sprint("Lorem Ipsum")

	testhelper.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopLeft().WithTitle("a").Sprint("Lorem Ipsum")

	testhelper.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleBottomRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomRight().WithTitle("a").Sprint("Lorem Ipsum")

	testhelper.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopRight().WithTitle("a").Sprint("Lorem Ipsum")

	testhelper.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopCenter().WithTitle("a").Sprint("Lorem Ipsum")

	testhelper.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleBottomCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomCenter().WithTitle("a").Sprint("Lorem Ipsum")

	testhelper.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithBoxStyle(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithBoxStyle(s)

	testhelper.AssertEqual(t, s, p2.BoxStyle)
	testhelper.AssertZero(t, p.BoxStyle)
}

func TestBoxPrinter_WithLeftPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithLeftPadding(5)

	testhelper.AssertEqual(t, 5, p2.LeftPadding)
	testhelper.AssertZero(t, p.LeftPadding)
}

func TestBoxPrinter_WithRightPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithRightPadding(5)

	testhelper.AssertEqual(t, 5, p2.RightPadding)
	testhelper.AssertZero(t, p.RightPadding)
}

func TestBoxPrinter_WithTextStyle(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	testhelper.AssertEqual(t, s, p2.TextStyle)
	testhelper.AssertZero(t, p.TextStyle)
}

func TestBoxPrinter_WithTopLeftCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopLeftCornerString("-")

	testhelper.AssertEqual(t, "-", p2.TopLeftCornerString)
	testhelper.AssertZero(t, p.TopLeftCornerString)
}

func TestBoxPrinter_WithTopPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopPadding(5)

	testhelper.AssertEqual(t, 5, p2.TopPadding)
	testhelper.AssertZero(t, p.TopPadding)
}

func TestBoxPrinter_WithHorizontalPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithHorizontalPadding(5)

	testhelper.AssertEqual(t, 5, p2.LeftPadding)
	testhelper.AssertEqual(t, 5, p2.RightPadding)
	testhelper.AssertEqual(t, 0, p.LeftPadding)
	testhelper.AssertEqual(t, 0, p.RightPadding)
}

func TestBoxPrinter_WithVerticalPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithVerticalPadding(5)

	testhelper.AssertEqual(t, 5, p2.TopPadding)
	testhelper.AssertEqual(t, 5, p2.BottomPadding)
	testhelper.AssertEqual(t, 0, p.TopPadding)
	testhelper.AssertEqual(t, 0, p.BottomPadding)
}

func TestBoxPrinter_WithPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithPadding(5)

	testhelper.AssertEqual(t, 5, p2.TopPadding)
	testhelper.AssertEqual(t, 5, p2.BottomPadding)
	testhelper.AssertEqual(t, 5, p2.LeftPadding)
	testhelper.AssertEqual(t, 5, p2.RightPadding)
	testhelper.AssertEqual(t, 0, p.TopPadding)
	testhelper.AssertEqual(t, 0, p.BottomPadding)
	testhelper.AssertEqual(t, 0, p.LeftPadding)
	testhelper.AssertEqual(t, 0, p.RightPadding)
}

func TestBoxPrinter_WithInvalidTopPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopPadding(-5)

	testhelper.AssertEqual(t, 0, p2.TopPadding)
	testhelper.AssertZero(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidBottomPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomPadding(-5)

	testhelper.AssertEqual(t, 0, p2.BottomPadding)
	testhelper.AssertZero(t, p.BottomPadding)
}

func TestBoxPrinter_WithInvalidLeftPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithLeftPadding(-5)

	testhelper.AssertEqual(t, 0, p2.LeftPadding)
	testhelper.AssertZero(t, p.LeftPadding)
}

func TestBoxPrinter_WithInvalidRightPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithRightPadding(-5)

	testhelper.AssertEqual(t, 0, p2.RightPadding)
	testhelper.AssertZero(t, p.RightPadding)
}

func TestBoxPrinter_WithTopRightCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopRightCornerString("-")

	testhelper.AssertEqual(t, "-", p2.TopRightCornerString)
	testhelper.AssertZero(t, p.TopRightCornerString)
}

func TestBoxPrinter_WithVerticalString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithVerticalString("-")

	testhelper.AssertEqual(t, "-", p2.VerticalString)
	testhelper.AssertZero(t, p.VerticalString)
}

func TestBoxPrinter_WithHorizontalString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithHorizontalString("-")

	testhelper.AssertEqual(t, "-", p2.HorizontalString)
	testhelper.AssertZero(t, p.HorizontalString)
}

func TestBoxPrinter_WithWriter(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testhelper.AssertEqual(t, s, p2.Writer)
	testhelper.AssertZero(t, p.Writer)
}
