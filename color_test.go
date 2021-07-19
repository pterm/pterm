package pterm_test

import (
	"errors"
	"io"
	"testing"

	"github.com/gookit/color"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestStylePrinterPrintMethods(t *testing.T) {
	p := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)

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
}

func TestRemoveColorFromString(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testString := pterm.Cyan(randomString)
		assert.Equal(t, randomString, pterm.RemoveColorFromString(testString))
	}
}

func TestColorPrinterPrintMethods(t *testing.T) {
	p := pterm.Color(16)

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

func TestNewStyle(t *testing.T) {

}

func TestStyle_Add(t *testing.T) {
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}))
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}).Add(pterm.Style{pterm.Bold}))
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen, pterm.Bold}))
	assert.Equal(t, pterm.Style{pterm.FgRed, pterm.BgGreen, pterm.Bold}, pterm.Style{pterm.FgRed}.Add(pterm.Style{pterm.BgGreen}, pterm.Style{pterm.Bold}))
}

func TestStyle_Code(t *testing.T) {
	assert.NotEmpty(t, pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold).Code())
}

func TestStyle_String(t *testing.T) {
	assert.NotEmpty(t, pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold).String())
}

func TestEnableColor(t *testing.T) {
	pterm.EnableColor()
	assert.True(t, color.Enable)
	assert.True(t, pterm.PrintColor)
}

func TestDisableColor(t *testing.T) {
	pterm.DisableColor()
	assert.False(t, color.Enable)
	assert.False(t, pterm.PrintColor)
}

func TestDisabledColorDoesPrintPlainString(t *testing.T) {
	pterm.DisableColor()
	assert.Equal(t, "Hello, World!", pterm.FgRed.Sprint("Hello, World!"))
}
