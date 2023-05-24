package pterm_test

import (
	"io"
	"testing"

	"github.com/pterm/pterm"
)

func TestCenterPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.CenterPrinter{})
}

func TestCenterPrinter(t *testing.T) {
	t.Run("Println", func(t *testing.T) {
		printerTest(t, func() {
			pterm.DefaultCenter.Println("Lorem ipsum dolor sit amet")
		})
	})

	t.Run("WithCenterEachLineSeparately", func(t *testing.T) {
		printerTest(t, func() {
			pterm.DefaultCenter.WithCenterEachLineSeparately().Println("Lorem ipsum dolor sit amet\nconsetetur sadipscing elitre\nsed diam nonumy eirmod tempor invidunt ut")
		})
	})
}

func TestCenterPrinterPrintMethodsCenterSeparately(t *testing.T) {
	p := pterm.DefaultCenter.WithCenterEachLineSeparately()

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

	t.Run("Sprintln", func(t *testing.T) {
		testSprintlnContains(t, func(a interface{}) string {
			return p.Sprintln(a)
		})
	})
}
