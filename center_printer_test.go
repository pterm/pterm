package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestCenterPrinter_WithCenterEachLineSeparately(t *testing.T) {
	p := CenterPrinter{}
	p2 := p.WithCenterEachLineSeparately()

	assert.True(t, p2.CenterEachLineSeparately)
	assert.False(t, p.CenterEachLineSeparately)
}

func TestCenterPrinterPrintMethods(t *testing.T) {
	p := DefaultCenter

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

func TestCenterPrinterPrintMethodsCenterSeparately(t *testing.T) {
	p := DefaultCenter.WithCenterEachLineSeparately()

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
