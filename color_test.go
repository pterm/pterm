package pterm

import (
	"github.com/gookit/color"
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestStylePrinterPrintMethods(t *testing.T) {
	p := NewStyle(FgRed, BgBlue, Bold)

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

func TestRemoveColorFromString(t *testing.T) {
	for _, randomString := range internal.RandomStrings {
		testString := Cyan(randomString)
		assert.Equal(t, randomString, RemoveColorFromString(testString))
	}
}

func TestColorPrinterPrintMethods(t *testing.T) {
	p := Color(16)

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

func TestNewStyle(t *testing.T) {

}

func TestStyle_Add(t *testing.T) {
	assert.Equal(t, Style{FgRed, BgGreen}, Style{FgRed}.Add(Style{BgGreen}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen}).Add(Style{Bold}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen, Bold}))
	assert.Equal(t, Style{FgRed, BgGreen, Bold}, Style{FgRed}.Add(Style{BgGreen}, Style{Bold}))
}

func TestStyle_Code(t *testing.T) {
	assert.NotEmpty(t, NewStyle(FgRed, BgBlue, Bold).Code())
}

func TestStyle_String(t *testing.T) {
	assert.NotEmpty(t, NewStyle(FgRed, BgBlue, Bold).String())
}

func Test_colors2code(t *testing.T) {
	assert.NotEmpty(t, colors2code(FgRed, FgBlue))
}

func TestEnableColor(t *testing.T) {
	EnableColor()
	assert.True(t, color.Enable)
	assert.True(t, PrintColor)
}

func TestDisableColor(t *testing.T) {
	DisableColor()
	assert.False(t, color.Enable)
	assert.False(t, PrintColor)
}

func TestDisabledColorDoesPrintPlainString(t *testing.T) {
	DisableColor()
	assert.Equal(t, "Hello, World!", FgRed.Sprint("Hello, World!"))
}
