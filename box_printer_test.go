package pterm_test

import (
	"errors"
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

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
		testza.AssertContains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutError", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnError(nil)
		})
		testza.AssertZero(t, result)
	})

	t.Run("PrintOnErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
		})
		testza.AssertContains(t, result, "hello world")
	})

	t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
		result := captureStdout(func(w io.Writer) {
			p.PrintOnErrorf("", nil)
		})
		testza.AssertZero(t, result)
	})
}

func TestBoxPrinter_WithBottomLeftCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomLeftCornerString("-")

	testza.AssertEqual(t, "-", p2.BottomLeftCornerString)
	testza.AssertZero(t, p.BottomLeftCornerString)
}

func TestBoxPrinter_WithBottomPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomPadding(5)

	testza.AssertEqual(t, 5, p2.BottomPadding)
	testza.AssertZero(t, p.BottomPadding)
}

func TestBoxPrinter_WithBottomRightCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomRightCornerString("-")

	testza.AssertEqual(t, "-", p2.BottomRightCornerString)
	testza.AssertZero(t, p.BottomRightCornerString)
}

func TestBoxPrinter_WithTitle(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitle("-")

	testza.AssertEqual(t, "-", p2.Title)
	testza.AssertZero(t, p.Title)
}

func TestBoxPrinter_WithTitleTopLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopLeft()

	testza.AssertEqual(t, true, p2.TitleTopLeft)
	testza.AssertEqual(t, false, p.TitleTopLeft)
}

func TestBoxPrinter_WithTitleTopRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopRight()

	testza.AssertEqual(t, true, p2.TitleTopRight)
	testza.AssertEqual(t, false, p.TitleTopRight)
}

func TestBoxPrinter_WithTitleTopCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopCenter()

	testza.AssertEqual(t, true, p2.TitleTopCenter)
	testza.AssertEqual(t, false, p.TitleTopCenter)
}

func TestBoxPrinter_WithTitleBottomRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomRight()

	testza.AssertEqual(t, true, p2.TitleBottomRight)
	testza.AssertEqual(t, false, p.TitleBottomRight)
}

func TestBoxPrinter_WithTitleBottomLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomLeft()

	testza.AssertEqual(t, true, p2.TitleBottomLeft)
	testza.AssertEqual(t, false, p.TitleBottomLeft)
}

func TestBoxPrinter_WithTitleBottomCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomCenter()

	testza.AssertEqual(t, true, p2.TitleBottomCenter)
	testza.AssertEqual(t, false, p.TitleBottomCenter)
}

func TestBoxPrinter_WithTitleWithTitleBottomLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomLeft().WithTitle("a").Sprint("Lorem Ipsum")

	testza.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopLeft(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopLeft().WithTitle("a").Sprint("Lorem Ipsum")

	testza.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleBottomRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomRight().WithTitle("a").Sprint("Lorem Ipsum")

	testza.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopRight(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopRight().WithTitle("a").Sprint("Lorem Ipsum")

	testza.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleTopCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleTopCenter().WithTitle("a").Sprint("Lorem Ipsum")

	testza.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithTitleWithTitleBottomCenter(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTitleBottomCenter().WithTitle("a").Sprint("Lorem Ipsum")

	testza.AssertContains(t, p2, "Lorem Ipsum")
}

func TestBoxPrinter_WithBoxStyle(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithBoxStyle(s)

	testza.AssertEqual(t, s, p2.BoxStyle)
	testza.AssertZero(t, p.BoxStyle)
}

func TestBoxPrinter_WithLeftPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithLeftPadding(5)

	testza.AssertEqual(t, 5, p2.LeftPadding)
	testza.AssertZero(t, p.LeftPadding)
}

func TestBoxPrinter_WithRightPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithRightPadding(5)

	testza.AssertEqual(t, 5, p2.RightPadding)
	testza.AssertZero(t, p.RightPadding)
}

func TestBoxPrinter_WithTextStyle(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	testza.AssertEqual(t, s, p2.TextStyle)
	testza.AssertZero(t, p.TextStyle)
}

func TestBoxPrinter_WithTopLeftCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopLeftCornerString("-")

	testza.AssertEqual(t, "-", p2.TopLeftCornerString)
	testza.AssertZero(t, p.TopLeftCornerString)
}

func TestBoxPrinter_WithTopPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopPadding(5)

	testza.AssertEqual(t, 5, p2.TopPadding)
	testza.AssertZero(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidTopPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopPadding(-5)

	testza.AssertEqual(t, 0, p2.TopPadding)
	testza.AssertZero(t, p.TopPadding)
}

func TestBoxPrinter_WithInvalidBottomPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithBottomPadding(-5)

	testza.AssertEqual(t, 0, p2.BottomPadding)
	testza.AssertZero(t, p.BottomPadding)
}

func TestBoxPrinter_WithInvalidLeftPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithLeftPadding(-5)

	testza.AssertEqual(t, 0, p2.LeftPadding)
	testza.AssertZero(t, p.LeftPadding)
}

func TestBoxPrinter_WithInvalidRightPadding(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithRightPadding(-5)

	testza.AssertEqual(t, 0, p2.RightPadding)
	testza.AssertZero(t, p.RightPadding)
}

func TestBoxPrinter_WithTopRightCornerString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithTopRightCornerString("-")

	testza.AssertEqual(t, "-", p2.TopRightCornerString)
	testza.AssertZero(t, p.TopRightCornerString)
}

func TestBoxPrinter_WithVerticalString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithVerticalString("-")

	testza.AssertEqual(t, "-", p2.VerticalString)
	testza.AssertZero(t, p.VerticalString)
}

func TestBoxPrinter_WithHorizontalString(t *testing.T) {
	p := pterm.BoxPrinter{}
	p2 := p.WithHorizontalString("-")

	testza.AssertEqual(t, "-", p2.HorizontalString)
	testza.AssertZero(t, p.HorizontalString)
}

func TestBoxPrinter_WithWriter(t *testing.T) {
	p := pterm.BoxPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}
