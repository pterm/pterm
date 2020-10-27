package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

var prefixPrinters = []PrefixPrinter{Info, Success, Warning, Error, *Fatal.WithFatal(false)}

func TestPrefixPrinterNilPrint(t *testing.T) {
	p := PrefixPrinter{}
	p.Println("Hello, World!")
}

func TestPrefixPrinterPrintMethods(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("Print", func(t *testing.T) {
			internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
				p.Print(a)
			})
		})

		t.Run("PrintWithScope", func(t *testing.T) {
			internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
				p2 := p.WithScope(Scope{
					Text:  "test",
					Style: NewStyle(FgRed, BgBlue, Bold),
				})
				p2.Print(a)
			})
		})

		t.Run("PrintWithMultipleLines", func(t *testing.T) {
			p2 := p.WithScope(Scope{
				Text:  "test",
				Style: NewStyle(FgRed, BgBlue, Bold),
			})
			p2.Print("This text\nhas\nmultiple\nlines")
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
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			assert.NotEmpty(t, p.GetFormattedPrefix())
		})
	}
}

func TestPrefixPrinter_WithFatal(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithFatal()

			assert.Equal(t, true, p2.Fatal)
		})
	}
}

func TestPrefixPrinter_WithMessageStyle(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := NewStyle(FgRed, BgBlue, Bold)
			p2 := p.WithMessageStyle(s)

			assert.Equal(t, s, p2.MessageStyle)
		})
	}
}

func TestPrefixPrinter_WithPrefix(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := Prefix{
				Text:  "test",
				Style: NewStyle(FgRed, BgBlue, Bold),
			}
			p2 := p.WithPrefix(s)

			assert.Equal(t, s, p2.Prefix)
		})
	}
}

func TestPrefixPrinter_WithScope(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := Scope{
				Text:  "test",
				Style: NewStyle(FgRed, BgBlue, Bold),
			}
			p2 := p.WithScope(s)

			assert.Equal(t, s, p2.Scope)
		})
	}
}

func Test_checkFatal(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithFatal()
			assert.Panics(t, func() {
				p2.Println("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_WithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()

			assert.True(t, p2.Debugger)
		})
	}
}
