package pterm

import (
	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

var prefixPrinters = []PrefixPrinter{Info, Success, Warning, Error, *Fatal.WithFatal(false)}

func TestPrefixPrinterNilPrint(t *testing.T) {
	proxyToDevNull()
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

func TestSprintfWithNewLineEnding(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			assert.NotContains(t, "\n\n", p.Sprintf("%s\n\n\n\n", "Hello, World!"))
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

func TestPrefixPrinter_WithShowLineNumber(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithShowLineNumber()

			assert.Equal(t, true, p2.ShowLineNumber)
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

func TestPrefixPrinter_PrintWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			EnableDebugMessages()
			internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
				p2.Print(a)
			})
		})
	}
}

func TestPrefixPrinter_PrintlnWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			EnableDebugMessages()
			internal.TestPrintlnContains(t, func(w io.Writer, a interface{}) {
				p2.Println(a)
			})
		})
	}
}

func TestPrefixPrinter_PrintfWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			EnableDebugMessages()
			internal.TestPrintfContains(t, func(w io.Writer, format string, a interface{}) {
				p2.Printf(format, a)
			})
		})
	}
}

func TestPrefixPrinter_SprintWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			EnableDebugMessages()
			internal.TestSprintContains(t, func(a interface{}) string {
				return p2.Sprint(a)
			})
		})
	}
}

func TestPrefixPrinter_SprintlnWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			EnableDebugMessages()
			internal.TestSprintlnContains(t, func(a interface{}) string {
				return p2.Sprintln(a)
			})
		})
	}
}

func TestPrefixPrinter_SprintfWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			EnableDebugMessages()
			internal.TestSprintfContains(t, func(format string, a interface{}) string {
				return p2.Sprintf(format, a)
			})
		})
	}
}

func TestPrefixPrinter_PrintWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			DisableDebugMessages()
			internal.TestDoesNotOutput(t, func(w io.Writer) {
				p2.Print("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_PrintlnWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			DisableDebugMessages()
			internal.TestDoesNotOutput(t, func(w io.Writer) {
				p2.Println("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_PrintfWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			DisableDebugMessages()
			internal.TestDoesNotOutput(t, func(w io.Writer) {
				p2.Printf("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_SprintWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			internal.TestEmpty(t, func(a interface{}) string {
				return p2.Sprint(a)
			})
		})
	}
}

func TestPrefixPrinter_SprintlnWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			DisableDebugMessages()
			internal.TestEmpty(t, func(a interface{}) string {
				return p2.Sprintln(a)
			})
		})
	}
}

func TestPrefixPrinter_SprintfWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			DisableDebugMessages()
			internal.TestEmpty(t, func(a interface{}) string {
				return p2.Sprintf("Hello, %s!", a)
			})
		})
	}
}
