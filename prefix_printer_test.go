package pterm_test

import (
	"errors"
	"fmt"
	"io"
	"os"
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

func TestPrefixPrinterPrintMethods(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("Print", func(t *testing.T) {
			testPrintContains(t, func(w io.Writer, a any) {
				p.WithWriter(w).Print(a)
			})
		})

		t.Run("PrintWithScope", func(t *testing.T) {
			testPrintContains(t, func(w io.Writer, a any) {
				p2 := p.WithScope(pterm.Scope{
					Text:  "test",
					Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
				})
				p2.WithWriter(w).Print(a)
			})
		})

		t.Run("PrintWithShowLineNumber", func(t *testing.T) {
			testPrintContains(t, func(w io.Writer, a any) {
				p2 := p.WithShowLineNumber().WithWriter(w)
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
			testPrintfContains(t, func(w io.Writer, format string, a any) {
				p.WithWriter(w).Printf(format, a)
			})
		})

		t.Run("Printfln", func(t *testing.T) {
			testPrintflnContains(t, func(w io.Writer, format string, a any) {
				p.WithWriter(w).Printfln(format, a)
			})
		})

		t.Run("Println", func(t *testing.T) {
			testPrintlnContains(t, func(w io.Writer, a any) {
				p.WithWriter(w).Println(a)
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
				p.WithWriter(w).PrintOnError(errors.New("hello world"))
			})
			testza.AssertContains(t, result, "hello world")
		})

		t.Run("PrintIfError_WithoutError", func(t *testing.T) {
			result := captureStdout(func(w io.Writer) {
				p.WithWriter(w).PrintOnError(nil)
			})
			testza.AssertZero(t, result)
		})

		t.Run("PrintOnErrorf", func(t *testing.T) {
			result := captureStdout(func(w io.Writer) {
				p.WithWriter(w).PrintOnErrorf("wrapping error : %w", errors.New("hello world"))
			})
			testza.AssertContains(t, result, "hello world")
		})

		t.Run("PrintIfError_WithoutErrorf", func(t *testing.T) {
			result := captureStdout(func(w io.Writer) {
				p.WithWriter(w).PrintOnErrorf("", nil)
			})
			testza.AssertZero(t, result)
		})
	}
}

func TestPrefixPrinterWithoutPrefix(t *testing.T) {
	pterm.DisableStyling()
	for _, p := range prefixPrinters {
		p2 := p.WithPrefix(pterm.Prefix{})
		t.Run("", func(t *testing.T) {
			for _, printable := range printables {
				ret := captureStdout(func(w io.Writer) {
					p2.WithWriter(w).Print(printable)
				})
				testza.AssertEqual(t, ret, fmt.Sprint(printable))
			}
		})
	}
	pterm.EnableStyling()
}

func TestSprintfWithNewLineEnding(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			testza.AssertNotContains(t, "\n\n", p.Sprintf("%s\n\n\n\n", "Hello, World!"))
		})
	}
}

func TestPrefixPrinter_GetFormattedPrefix(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			testza.AssertNotZero(t, p.GetFormattedPrefix())
		})
	}
}

func TestPrefixPrinter_WithFatal(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithFatal()

			testza.AssertEqual(t, true, p2.Fatal)
		})
	}
}

func TestPrefixPrinter_WithShowLineNumber(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithShowLineNumber()

			testza.AssertEqual(t, true, p2.ShowLineNumber)
		})
	}
}

func TestPrefixPrinter_WithMessageStyle(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
			p2 := p.WithMessageStyle(s)

			testza.AssertEqual(t, s, p2.MessageStyle)
		})
	}
}

func TestPrefixPrinter_WithPrefix(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := pterm.Prefix{
				Text:  "test",
				Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
			}
			p2 := p.WithPrefix(s)

			testza.AssertEqual(t, s, p2.Prefix)
		})
	}
}

func TestPrefixPrinter_WithScope(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := pterm.Scope{
				Text:  "test",
				Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
			}
			p2 := p.WithScope(s)

			testza.AssertEqual(t, s, p2.Scope)
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

func TestPrefixPrinter_WithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()

			testza.AssertTrue(t, p2.Debugger)
		})
	}
}

func TestPrefixPrinter_PrintWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.EnableDebugMessages()
			testPrintContains(t, func(w io.Writer, a any) {
				p2.WithWriter(w).Print(a)
			})
		})
	}
}

func TestPrefixPrinter_PrintlnWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.EnableDebugMessages()
			testPrintlnContains(t, func(w io.Writer, a any) {
				p2.WithWriter(w).Println(a)
			})
		})
	}
}

func TestPrefixPrinter_PrintfWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.EnableDebugMessages()
			testPrintfContains(t, func(w io.Writer, format string, a any) {
				p2.WithWriter(w).Printf(format, a)
			})
		})
	}
}

func TestPrefixPrinter_SprintWithDebugger(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithDebugger()
			pterm.EnableDebugMessages()
			testSprintContains(t, func(a any) string {
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
			testSprintlnContains(t, func(a any) string {
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
			testSprintfContains(t, func(format string, a any) string {
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
			testEmpty(t, func(a any) string {
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
			testEmpty(t, func(a any) string {
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
			testEmpty(t, func(a any) string {
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
			testEmpty(t, func(a any) string {
				return p2.Sprintfln("Hello, %s!", a)
			})
		})
	}
}

func TestPrefixPrinter_WithLineNumberOffset(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			p2 := p.WithLineNumberOffset(1337)

			testza.AssertEqual(t, 1337, p2.LineNumberOffset)
		})
	}
}

func TestPrefixPrinter_WithWriter(t *testing.T) {
	for _, p := range prefixPrinters {
		t.Run("", func(t *testing.T) {
			s := os.Stderr
			p2 := p.WithWriter(s)

			testza.AssertEqual(t, s, p2.Writer)
		})
	}
}
