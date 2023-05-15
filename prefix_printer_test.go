package pterm_test

import (
	"errors"
	"io"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

var prefixPrinters = []pterm.PrefixPrinter{pterm.Info, pterm.Success, pterm.Warning, pterm.Error, *pterm.Fatal.WithFatal(false)}

func TestPrefixPrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	p := pterm.PrefixPrinter{}
	p.Println("Hello, World!")
}

func TestPrefixPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.PrefixPrinter{})
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
				p2 := p.WithScope(pterm.Scope{
					Text:  "test",
					Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
				})
				p2.Print(a)
			})
		})

		t.Run("PrintWithShowLineNumber", func(t *testing.T) {
			testPrintContains(t, func(w io.Writer, a interface{}) {
				p2 := p.WithShowLineNumber()
				p2.Print(a)
			})
		})

		t.Run("PrintWithMultipleLines", func(t *testing.T) {
			p2 := p.WithScope(pterm.Scope{
				Text:  "test",
				Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
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
}

func Test_checkFatal(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithFatal()
			testza.AssertPanics(t, func() {
				p2.Println("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_PrintWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.EnableDebugMessages()
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
			pterm.EnableDebugMessages()
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
			pterm.EnableDebugMessages()
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
			pterm.EnableDebugMessages()
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
			pterm.EnableDebugMessages()
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
			pterm.EnableDebugMessages()
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
			pterm.DisableDebugMessages()
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
			pterm.DisableDebugMessages()
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
			pterm.DisableDebugMessages()
			testDoesNotOutput(t, func(w io.Writer) {
				p2.Printf("Hello, World!")
			})
		})
	}
}

func TestPrefixPrinter_PrintflnWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.DisableDebugMessages()
			testDoesNotOutput(t, func(w io.Writer) {
				p2.Printfln("Hello, World!")
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
			pterm.DisableDebugMessages()
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
			pterm.DisableDebugMessages()
			testEmpty(t, func(a interface{}) string {
				return p2.Sprintf("Hello, %s!", a)
			})
		})
	}
}

func TestPrefixPrinter_SprintflnWithoutDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.DisableDebugMessages()
			testEmpty(t, func(a interface{}) string {
				return p2.Sprintfln("Hello, %s!", a)
			})
		})
	}
}
