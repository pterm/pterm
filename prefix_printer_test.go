package pterm

import (
	"github.com/pterm/pterm/internal"
	"io"
	"testing"
)

func TestPrefixPrinterNilPrint(t *testing.T) {
	p := PrefixPrinter{}
	p.Println("Hello, World!")
}

func TestPrefixPrinterPrintMethods(t *testing.T) {
	prefixPrinters := []PrefixPrinter{Info, Success, Warning, Error, *Fatal.WithFatal(false)}

	for _, p := range prefixPrinters {
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
}

func TestPrefixPrinter_GetFormattedPrefix(t *testing.T) {

}

func TestPrefixPrinter_WithFatal(t *testing.T) {

}

func TestPrefixPrinter_WithMessageStyle(t *testing.T) {

}

func TestPrefixPrinter_WithPrefix(t *testing.T) {

}

func TestPrefixPrinter_WithScope(t *testing.T) {

}

func Test_checkFatal(t *testing.T) {

}
