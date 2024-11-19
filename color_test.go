package pterm_test

import (
	"errors"
	"io"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/gookit/color"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

func TestStylePrinterPrintMethods(t *testing.T) {
	p := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)

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
}

func TestRemoveColorFromString(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testString := pterm.Cyan(randomString)
		testza.AssertEqual(t, randomString, pterm.RemoveColorFromString(testString))
	}
}

func TestColorPrinterPrintMethods(t *testing.T) {
	p := pterm.Color(16)

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

func TestNewStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	testza.AssertEqual(t, s, &pterm.Style{pterm.FgRed, pterm.BgBlue, pterm.Bold})
}

func TestColor_ToStyle(t *testing.T) {
	s := pterm.FgCyan.ToStyle()
	testza.AssertEqual(t, s, &pterm.Style{pterm.FgCyan})
}

func TestStyle_Add(t *testing.T) {
	testza.AssertEqual(t, pterm.Style{pterm.FgRed, pterm.BgGreen}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}))
	testza.AssertEqual(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}).Add(pterm.Style{pterm.Bold}))
	testza.AssertEqual(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen, pterm.Bold}))
	testza.AssertEqual(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}, pterm.Style{pterm.Bold}))
}

func TestStyle_Code(t *testing.T) {
	testza.AssertNotZero(t, pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold).Code())
}

func TestStyle_String(t *testing.T) {
	testza.AssertNotZero(t, pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold).String())
}

func TestEnableColor(t *testing.T) {
	pterm.EnableColor()
	testza.AssertTrue(t, color.Enable)
	testza.AssertTrue(t, pterm.PrintColor)
}

func TestDisableColor(t *testing.T) {
	pterm.DisableColor()
	testza.AssertFalse(t, color.Enable)
	testza.AssertFalse(t, pterm.PrintColor)
}

func TestDisabledColorDoesPrintPlainString(t *testing.T) {
	pterm.DisableColor()
	testza.AssertEqual(t, "Hello, World!", pterm.FgRed.Sprint("Hello, World!"))
}
