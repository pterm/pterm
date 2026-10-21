package pterm_test

// Behavioral tests for BasicTextPrinter.
//
// The builder methods, Print*/Sprint* delegation, PrintOnError semantics,
// custom writers and the global styling invariants are covered generically in
// contract_test.go. This file verifies the printer's actual rendering: how the
// configured Style is (or is not) applied to the text.

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestBasicTextPrinterPassesTextThroughWithoutStyle(t *testing.T) {
	assert.Equal(t, "Hello, PTerm!", pterm.DefaultBasicText.Sprint("Hello, PTerm!"))
}

func TestBasicTextPrinterEmptyInput(t *testing.T) {
	assert.Equal(t, "", pterm.DefaultBasicText.Sprint(""))
	assert.Equal(t, "\n", pterm.DefaultBasicText.Sprintln(""))
}

func TestBasicTextPrinterFormatVerbsAreLiteralInSprint(t *testing.T) {
	// Sprint is not a formatting function; "%s" must survive as literal text.
	input := "100%s done"

	assert.Equal(t, input, pterm.DefaultBasicText.Sprint(input))
}

func TestBasicTextPrinterAppliesStyleExactly(t *testing.T) {
	p := pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed, pterm.Bold))

	// FgRed = SGR 31, Bold = SGR 1: the text must be wrapped in exactly one
	// combined SGR sequence and one reset.
	assert.Equal(t, "\x1b[31;1mHello\x1b[0m", p.Sprint("Hello"))
}

func TestBasicTextPrinterStylesEachLineSeparately(t *testing.T) {
	p := pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed))

	// Every line is wrapped individually so the style survives terminal
	// operations that work line by line.
	assert.Equal(t, "\x1b[31ma\x1b[0m\n\x1b[31mb\x1b[0m", p.Sprint("a\nb"))
}

func TestBasicTextPrinterReappliesStyleAfterEmbeddedReset(t *testing.T) {
	p := pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgGreen))
	inner := pterm.FgRed.Sprint("red") // "\x1b[31mred\x1b[0m"

	// The reset sequence of the embedded red segment must re-open the outer
	// green style, otherwise " b" would render unstyled.
	assert.Equal(t, "\x1b[32ma \x1b[31mred\x1b[0m\x1b[32m b\x1b[0m", p.Sprint("a "+inner+" b"))
}

func TestBasicTextPrinterKeepsVisibleTextOfPreStyledInput(t *testing.T) {
	out := pterm.DefaultBasicText.Sprint("\x1b[31mred\x1b[0m plain")

	// Input that already contains ANSI codes keeps its visible text and its
	// original color sequence.
	assert.Equal(t, "red plain", stripANSI(out))
	assert.Contains(t, out, "\x1b[31mred")
}

func TestBasicTextPrinterSprintlnAppendsExactlyOneNewline(t *testing.T) {
	assert.Equal(t, fmt.Sprintln("Hello"), pterm.DefaultBasicText.Sprintln("Hello"))
	assert.Equal(t, "Hello\n", pterm.DefaultBasicText.Sprintln("Hello"))

	// Multiple arguments format with intervening spaces and end with exactly one newline.
	assert.Equal(t, fmt.Sprintln("Hello", "World"), pterm.DefaultBasicText.Sprintln("Hello", "World"))
	assert.Equal(t, "Hello World\n", pterm.DefaultBasicText.Sprintln("Hello", "World"))

	assert.Equal(t, fmt.Sprintln("count:", 42, true), pterm.DefaultBasicText.Sprintln("count:", 42, true))
	assert.Equal(t, "count: 42 true\n", pterm.DefaultBasicText.Sprintln("count:", 42, true))

	// Verify no double newline is added.
	out := pterm.DefaultBasicText.Sprintln("hello")
	assert.True(t, strings.HasSuffix(out, "\n"))
	assert.False(t, strings.HasSuffix(out, "\n\n"))

	styled := pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed))
	assert.Equal(t, "\x1b[31mHello\x1b[0m\n", styled.Sprintln("Hello"))
	assert.Equal(t, "\x1b[31mHello World\x1b[0m\n", styled.Sprintln("Hello", "World"))
}

func TestBasicTextPrinterSprintfFormatsBeforeStyling(t *testing.T) {
	p := pterm.DefaultBasicText.WithStyle(pterm.NewStyle(pterm.FgRed))

	// The format expansion happens first, then the whole result is styled.
	assert.Equal(t, "\x1b[31m2 items\x1b[0m", p.Sprintf("%d %s", 2, "items"))
}

func TestBasicTextPrinterZeroValuePrintsPlainText(t *testing.T) {
	out := captureStdout(func(_ io.Writer) {
		p := pterm.BasicTextPrinter{}
		p.Println("Hello, World!")
	})

	assert.Equal(t, "Hello, World!\n", out)
}
