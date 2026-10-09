package pterm_test

// Behavioral tests for PrefixPrinter (pterm.Info, Success, Warning, Error,
// Fatal, Debug, Description).
//
// Builder methods, Print*/Sprint* delegation, PrintOnError semantics and the
// global styling invariants are covered generically in contract_test.go. This
// file verifies the actual rendering: the exact prefix layout, multiline
// indentation, scope, debug gating, fatal panics, line numbers and raw mode.

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// enableDebugMessages turns debug output on for one test and restores the
// default (disabled) state afterwards.
func enableDebugMessages(t *testing.T) {
	t.Helper()
	pterm.EnableDebugMessages()
	t.Cleanup(pterm.DisableDebugMessages)
}

func TestPrefixPrinterPlainOutput(t *testing.T) {
	// The prefix is rendered as " <text> " followed by a separating space, so
	// the visible output is "<padded prefix> <message>".
	tests := []struct {
		name     string
		printer  pterm.PrefixPrinter
		expected string
	}{
		{"Info", pterm.Info, "  INFO    message"},
		{"Success", pterm.Success, " SUCCESS  message"},
		{"Warning", pterm.Warning, " WARNING  message"},
		{"Error", pterm.Error, "  ERROR   message"},
		{"Fatal", *pterm.Fatal.WithFatal(false), "  FATAL   message"},
		{"Description", pterm.Description, " Description  message"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, stripANSI(tc.printer.Sprint("message")))
		})
	}
}

func TestPrefixPrinterStyledOutput(t *testing.T) {
	// Exact ANSI for one representative case: the Info prefix is FgBlack(30)
	// on BgCyan(46), the message is FgLightCyan(96).
	assert.Equal(t, "\x1b[30;46m  INFO   \x1b[0m \x1b[96mx\x1b[0m", pterm.Info.Sprint("x"))
}

func TestPrefixPrinterMultilineIndentsContinuationLines(t *testing.T) {
	// Continuation lines are indented with visual width of prefix plus 2 spaces
	// plus the separating space, so the message column stays aligned under the
	// first line.
	assert.Equal(t, "  INFO    first\n          second", stripANSI(pterm.Info.Sprint("first\nsecond")))
	assert.Equal(t, "  ERROR   first\n          second", stripANSI(pterm.Error.Sprint("first\nsecond")))
}

func TestPrefixPrinterMultilineWideCharacterAlignment(t *testing.T) {
	tests := []struct {
		name     string
		prefix   string
		input    string
		expected string
	}{
		{
			name:     "ASCII control",
			prefix:   "INFO",
			input:    "first line\nsecond line",
			expected: " INFO  first line\n       second line",
		},
		{
			name:     "CJK two characters",
			prefix:   "情報",
			input:    "first line\nsecond line",
			expected: " 情報  first line\n       second line",
		},
		{
			name:     "CJK four characters three lines",
			prefix:   "警告です",
			input:    "line 1\nline 2\nline 3",
			expected: " 警告です  line 1\n           line 2\n           line 3",
		},
		{
			name:     "Emoji wide character",
			prefix:   "\U0001F680",
			input:    "line 1\nline 2",
			expected: " \U0001F680  line 1\n     line 2",
		},
		{
			name:     "Mixed ASCII and CJK",
			prefix:   "LOG 日志",
			input:    "entry 1\nentry 2",
			expected: " LOG 日志  entry 1\n           entry 2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := pterm.PrefixPrinter{
				Prefix: pterm.Prefix{
					Text: tc.prefix,
				},
			}
			result := stripANSI(p.Sprint(tc.input))
			assert.Equal(t, tc.expected, result)

			// Verify visual column alignment: the message text on every line
			// must start at the exact same visual display column.
			lines := strings.Split(result, "\n")
			inputLines := strings.Split(tc.input, "\n")
			require.Equal(t, len(inputLines), len(lines))

			expectedIndentWidth := runewidth.StringWidth(tc.prefix) + 3 // 1 leading + visual width + 1 trailing + 1 separating space
			for i, line := range lines {
				require.Truef(t, strings.HasSuffix(line, inputLines[i]), "line %d %q does not end with message %q", i, line, inputLines[i])
				indent := strings.TrimSuffix(line, inputLines[i])
				assert.Equalf(t, expectedIndentWidth, runewidth.StringWidth(indent),
					"line %d indent visual width: got %d, want %d", i, runewidth.StringWidth(indent), expectedIndentWidth)
			}
		})
	}
}

func TestPrefixPrinterMultilineWideCharacterStyled(t *testing.T) {
	style := pterm.NewStyle(pterm.BgCyan, pterm.FgBlack)
	p := pterm.PrefixPrinter{
		Prefix: pterm.Prefix{
			Text:  "情報",
			Style: style,
		},
	}
	out := p.Sprint("first\nsecond")
	expected := style.Sprint(" 情報 ") + " first\n" + style.Sprint("      ") + " second"
	assert.Equal(t, expected, out)
}

func TestPrefixPrinterSprintlnAppendsNewline(t *testing.T) {
	assert.Equal(t, "  INFO    message\n", stripANSI(pterm.Info.Sprintln("message")))
}

func TestPrefixPrinterCollapsesTrailingNewlines(t *testing.T) {
	// Any number of trailing newlines in the input collapses into exactly one.
	assert.Equal(t, "  INFO    a\n", stripANSI(pterm.Info.Sprint("a\n\n\n")))
}

func TestPrefixPrinterScopeIsRenderedBetweenPrefixAndMessage(t *testing.T) {
	p := pterm.Info.WithScope(pterm.Scope{Text: "myscope"})

	assert.Equal(t, "  INFO     (myscope) message", stripANSI(p.Sprint("message")))
}

func TestPrefixPrinterRawOutput(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	// In raw mode the padded prefix block degrades to "<prefix>: <message>".
	assert.Equal(t, "INFO: message", pterm.Info.Sprint("message"))
	assert.Equal(t, "ERROR: message", pterm.Error.Sprint("message"))

	t.Run("empty prefix prints the bare message", func(t *testing.T) {
		assert.Equal(t, "message", pterm.Info.WithPrefix(pterm.Prefix{}).Sprint("message"))
	})
}

func TestPrefixPrinterDebugGating(t *testing.T) {
	t.Run("prints nothing while debug messages are disabled", func(t *testing.T) {
		pterm.DisableDebugMessages()

		assert.Empty(t, pterm.Debug.Sprint("dbg"))
		assert.Empty(t, pterm.Debug.Sprintln("dbg"))
		assert.Empty(t, pterm.Debug.Sprintf("%s", "dbg"))
		assert.Empty(t, pterm.Debug.Sprintfln("%s", "dbg"))

		out := captureStdout(func(_ io.Writer) {
			pterm.Debug.Print("dbg")
			pterm.Debug.Println("dbg")
			pterm.Debug.Printf("%s", "dbg")
			pterm.Debug.Printfln("%s", "dbg")
		})
		assert.Empty(t, out)
	})

	t.Run("prints like any other prefix printer when enabled", func(t *testing.T) {
		enableDebugMessages(t)

		assert.Equal(t, "  DEBUG   message", stripANSI(pterm.Debug.Sprint("message")))
	})
}

func TestPrefixPrinterFatal(t *testing.T) {
	t.Run("panics after printing the message", func(t *testing.T) {
		setupStdoutCapture()

		assert.Panics(t, func() {
			pterm.Fatal.Println("fatal message")
		})

		assert.Equal(t, "  FATAL   fatal message\n", stripANSI(readStdout()))
	})

	t.Run("WithFatal(false) prints without panicking", func(t *testing.T) {
		out := captureStdout(func(_ io.Writer) {
			pterm.Fatal.WithFatal(false).Println("recoverable")
		})

		assert.Equal(t, "  FATAL   recoverable\n", stripANSI(out))
	})
}

func TestPrefixPrinterShowLineNumberReportsTheCallSite(t *testing.T) {
	_, file, line, ok := runtime.Caller(0)
	require.True(t, ok)

	out := captureStdout(func(_ io.Writer) {
		pterm.Info.WithShowLineNumber().Println("message") // keep exactly 4 lines below runtime.Caller(0)
	})

	assert.Equal(t, fmt.Sprintf("  INFO    message\n└ (%s:%d)\n", file, line+4), stripANSI(out))
}

func TestPrefixPrinterZeroValueDoesNotPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = captureStdout(func(_ io.Writer) {
			p := pterm.PrefixPrinter{}
			p.Println("Hello, World!")
		})
	})
}
