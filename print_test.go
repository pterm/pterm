package pterm_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

// The Sprint* family must behave exactly like the fmt equivalents.

func TestSprintSpacingRules(t *testing.T) {
	// Spaces are added between operands only when neither is a string.
	assert.Equal(t, "ab", pterm.Sprint("a", "b"))
	assert.Equal(t, "1 2", pterm.Sprint(1, 2))
	assert.Equal(t, "a1", pterm.Sprint("a", 1))
	assert.Equal(t, "true false", pterm.Sprint(true, false))
	assert.Equal(t, "", pterm.Sprint())

	verb := "%s" // via a variable, so vet's printf check does not flag the intentional literal
	assert.Equal(t, "%s", pterm.Sprint(verb), "Sprint must not interpret format verbs")
}

func TestSprintf(t *testing.T) {
	assert.Equal(t, "Hello, World!", pterm.Sprintf("Hello, %s!", "World"))
	assert.Equal(t, "42 true x", pterm.Sprintf("%d %v %s", 42, true, "x"))
}

func TestSprintfln(t *testing.T) {
	assert.Equal(t, "Hello, World!\n", pterm.Sprintfln("Hello, %s!", "World"))
	assert.Equal(t, "already newlined\n\n", pterm.Sprintfln("already newlined\n"),
		"Sprintfln always appends a newline, even if one is already present")
}

func TestSprintln(t *testing.T) {
	// fmt.Sprintln semantics: always spaces between operands, newline appended.
	assert.Equal(t, "a b\n", pterm.Sprintln("a", "b"))
	assert.Equal(t, "1 2\n", pterm.Sprintln(1, 2))
	assert.Equal(t, "\n", pterm.Sprintln())
}

func TestSprintoPrefixesCarriageReturn(t *testing.T) {
	assert.Equal(t, "\rover", pterm.Sprinto("over"))
	assert.Equal(t, "\r", pterm.Sprinto())
}

// The Print* family must write exactly what the Sprint* counterparts return.

func TestPrintWritesExactlySprint(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Print("a", 1, true) })
	assert.Equal(t, pterm.Sprint("a", 1, true), out)
}

func TestPrintlnWritesExactlySprintln(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Println("a", 1) })
	assert.Equal(t, pterm.Sprintln("a", 1), out)
}

func TestPrintfWritesExactlySprintf(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Printf("%s=%d", "x", 7) })
	assert.Equal(t, "x=7", out)
}

func TestPrintflnWritesExactlySprintfln(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Printfln("%s=%d", "x", 7) })
	assert.Equal(t, "x=7\n", out)
}

func TestPrintoOverwritesCurrentLine(t *testing.T) {
	out := captureStdout(func(_ io.Writer) {
		pterm.Printo("first")
		pterm.Printo("second")
	})

	// Each Printo call returns to the start of the line before printing, so
	// the second call visually overwrites the first.
	assert.Equal(t, "\rfirst\rsecond", out)
}

func TestPrintoInRawModePrintsWithoutCarriageReturn(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	out := captureStdout(func(_ io.Writer) { pterm.Printo("raw line") })
	assert.Equal(t, "raw line", out)
}

// The Fprint* family must honor custom writers and fall back to the default.

func TestFprintWritesToCustomWriter(t *testing.T) {
	var buf strings.Builder

	stdout := captureStdout(func(_ io.Writer) { pterm.Fprint(&buf, "custom target") })

	assert.Equal(t, "custom target", buf.String())
	assert.Empty(t, stdout, "output must not leak to the default writer")
}

func TestFprintNilWriterFallsBackToDefault(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Fprint(nil, "fallback") })
	assert.Equal(t, "fallback", out)
}

func TestFprintlnAppendsNewline(t *testing.T) {
	var buf strings.Builder

	pterm.Fprintln(&buf, "a", 1)
	assert.Equal(t, pterm.Sprint("a", 1)+"\n", buf.String())
}

func TestFprintlnNilWriterFallsBackToDefault(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Fprintln(nil, "fallback") })
	assert.Equal(t, "fallback\n", out)
}

func TestFprintoWritesCarriageReturnToCustomWriter(t *testing.T) {
	var buf strings.Builder

	stdout := captureStdout(func(_ io.Writer) { pterm.Fprinto(&buf, "over") })

	assert.Equal(t, "\rover", buf.String())
	assert.Empty(t, stdout)
}

func TestFprintoNilWriterFallsBackToDefault(t *testing.T) {
	out := captureStdout(func(_ io.Writer) { pterm.Fprinto(nil, "over") })
	assert.Equal(t, "\rover", out)
}

// DisableOutput must suppress every print path, EnableOutput must restore it.

func TestDisableOutputSuppressesEveryPrintPath(t *testing.T) {
	paths := map[string]func(w io.Writer){
		"Print":         func(_ io.Writer) { pterm.Print("x") },
		"Println":       func(_ io.Writer) { pterm.Println("x") },
		"Printf":        func(_ io.Writer) { pterm.Printf("%s", "x") },
		"Printfln":      func(_ io.Writer) { pterm.Printfln("%s", "x") },
		"Printo":        func(_ io.Writer) { pterm.Printo("x") },
		"Fprint":        func(w io.Writer) { pterm.Fprint(w, "x") },
		"Fprintln":      func(w io.Writer) { pterm.Fprintln(w, "x") },
		"Fprinto":       func(w io.Writer) { pterm.Fprinto(w, "x") },
		"PrintOnError":  func(_ io.Writer) { pterm.PrintOnError(errors.New("x")) },
		"PrintOnErrorf": func(_ io.Writer) { pterm.PrintOnErrorf("%w", errors.New("x")) },
	}

	pterm.DisableOutput()
	t.Cleanup(pterm.EnableOutput)

	for name, print := range paths {
		t.Run(name, func(t *testing.T) {
			var custom strings.Builder

			stdout := captureStdout(func(_ io.Writer) { print(&custom) })

			assert.Empty(t, stdout, "%s must not write to the default writer while output is disabled", name)
			assert.Empty(t, custom.String(), "%s must not write to a custom writer while output is disabled", name)
		})
	}
}

func TestEnableOutputRestoresPrinting(t *testing.T) {
	pterm.DisableOutput()
	pterm.EnableOutput()

	out := captureStdout(func(_ io.Writer) { pterm.Print("back again") })
	assert.Equal(t, "back again", out)
}

// PrintOnError / PrintOnErrorf top-level helpers.

func TestPrintOnError(t *testing.T) {
	t.Run("prints each non-nil error on its own line", func(t *testing.T) {
		out := captureStdout(func(_ io.Writer) {
			pterm.PrintOnError(errors.New("first"), nil, errors.New("second"))
		})
		assert.Equal(t, "first\nsecond\n", out)
	})

	t.Run("ignores nil errors and non-error values", func(t *testing.T) {
		out := captureStdout(func(_ io.Writer) {
			pterm.PrintOnError(nil, "not an error", 42)
		})
		assert.Empty(t, out)
	})
}

func TestPrintOnErrorf(t *testing.T) {
	t.Run("wraps non-nil errors with the format", func(t *testing.T) {
		out := captureStdout(func(_ io.Writer) {
			pterm.PrintOnErrorf("wrapped: %w", errors.New("inner"))
		})
		assert.Equal(t, "wrapped: inner\n", out)
	})

	t.Run("prints nothing for nil errors and non-error values", func(t *testing.T) {
		out := captureStdout(func(_ io.Writer) {
			pterm.PrintOnErrorf("wrapped: %w", nil, "not an error")
		})
		assert.Empty(t, out)
	})
}

// Default output routing.

func TestSetDefaultOutputRoutesAllPrints(t *testing.T) {
	prev := pterm.GetDefaultOutput()

	t.Cleanup(func() { pterm.SetDefaultOutput(prev) })

	var buf strings.Builder

	pterm.SetDefaultOutput(&buf)
	pterm.Print("routed")

	assert.Equal(t, "routed", buf.String())
	assert.Same(t, &buf, pterm.GetDefaultOutput())
}
