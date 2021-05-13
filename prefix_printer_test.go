package pterm

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
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
			testPrintContains(t, func(w io.Writer, a interface{}) {
				p.Print(a)
			})
		})

		t.Run("PrintWithScope", func(t *testing.T) {
			testPrintContains(t, func(w io.Writer, a interface{}) {
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

		t.Run("PrintIfError", func(t *testing.T) {
			result := captureStdout(func(w io.Writer) {
				p.PrintIfError(errors.New("hello world"))
			})
			assert.Contains(t, result, "hello world")
		})

		t.Run("PrintIfError_WithoutError", func(t *testing.T) {
			result := captureStdout(func(w io.Writer) {
				p.PrintIfError(nil)
			})
			assert.Empty(t, result)
		})
	}
}

func TestPrefixPrinterWithoutPrefix(t *testing.T) {
	DisableStyling()
	for _, p := range prefixPrinters {
		p2 := p.WithPrefix(Prefix{})
		t.Run("", func(t *testing.T) {
			for _, printable := range printables {
				ret := captureStdout(func(w io.Writer) {
					p2.Print(printable)
				})
				assert.Equal(t, ret, fmt.Sprint(printable))
			}
		})
	}
	EnableStyling()
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
			testPrintContains(t, func(w io.Writer, a interface{}) {
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
			testPrintlnContains(t, func(w io.Writer, a interface{}) {
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
			testPrintfContains(t, func(w io.Writer, format string, a interface{}) {
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
			testSprintContains(t, func(a interface{}) string {
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
			testSprintlnContains(t, func(a interface{}) string {
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
			testSprintfContains(t, func(format string, a interface{}) string {
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
			testDoesNotOutput(t, func(w io.Writer) {
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
			testDoesNotOutput(t, func(w io.Writer) {
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
			testDoesNotOutput(t, func(w io.Writer) {
				p2.Printf("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_SprintWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			testEmpty(t, func(a interface{}) string {
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
			testEmpty(t, func(a interface{}) string {
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
			testEmpty(t, func(a interface{}) string {
				return p2.Sprintf("Hello, %s!", a)
			})
		})
	}
}
