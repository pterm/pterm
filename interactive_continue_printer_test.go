package pterm_test

import (
	"io"
	"testing"

	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// showContinue runs the given continue prompt (which must terminate through
// the key presses simulated beforehand) and returns its result together with
// everything it printed.
func showContinue(t *testing.T, printer *pterm.InteractiveContinuePrinter, text ...string) (result string, output string) {
	t.Helper()

	var err error

	output = captureStdout(func(_ io.Writer) {
		result, err = showInteractive(t, func() (string, error) {
			return printer.Show(text...)
		})
	})

	require.NoError(t, err)

	return result, output
}

func TestInteractiveContinuePrinter_HandleKeysPickOptions(t *testing.T) {
	tests := []struct {
		name     string
		key      rune
		expected string
	}{
		{name: "y picks yes", key: 'y', expected: "yes"},
		{name: "Y picks the default case-insensitively", key: 'Y', expected: "yes"},
		{name: "n picks no", key: 'n', expected: "no"},
		{name: "a picks all", key: 'a', expected: "all"},
		{name: "c picks cancel", key: 'c', expected: "cancel"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			simulateKeys(t, tc.key)

			result, output := showContinue(t, &pterm.DefaultInteractiveContinue)

			assert.Equal(t, tc.expected, result)
			assert.Contains(t, stripANSI(output), tc.expected, "the chosen option must be echoed after the prompt")
		})
	}
}

func TestInteractiveContinuePrinter_EnterReturnsDefaultOption(t *testing.T) {
	t.Run("first option is the default", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		result, _ := showContinue(t, &pterm.DefaultInteractiveContinue)

		assert.Equal(t, "yes", result)
	})

	t.Run("WithDefaultValueIndex changes the default", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		result, output := showContinue(t, pterm.DefaultInteractiveContinue.WithDefaultValueIndex(2))

		assert.Equal(t, "all", result)
		assert.Contains(t, stripANSI(output), "[yes/no/All/cancel]", "the default option must be title-cased in the suffix")
	})

	t.Run("WithDefaultValue changes the default by name", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		result, _ := showContinue(t, pterm.DefaultInteractiveContinue.WithDefaultValue("no"))

		assert.Equal(t, "no", result)
	})
}

func TestInteractiveContinuePrinter_PromptShowsTextAndSuffix(t *testing.T) {
	simulateKeys(t, keys.Enter)

	_, output := showContinue(t, &pterm.DefaultInteractiveContinue, "Apply these changes?")

	assert.Contains(t, stripANSI(output), "Apply these changes? [Yes/no/all/cancel]: ")
}

func TestInteractiveContinuePrinter_ShowShortHandles(t *testing.T) {
	printer := pterm.DefaultInteractiveContinue.WithShowShortHandles()

	simulateKeys(t, 'n')

	result, output := showContinue(t, printer)

	assert.Equal(t, "no", result)
	assert.Contains(t, stripANSI(output), "[Y/n/a/c]", "short handles must be rendered instead of the full options")
}

func TestInteractiveContinuePrinter_CustomOptions(t *testing.T) {
	printer := pterm.DefaultInteractiveContinue.WithOptions([]string{"retry", "abort"})

	t.Run("handle keys are derived from the options", func(t *testing.T) {
		for key, expected := range map[rune]string{'r': "retry", 'a': "abort"} {
			simulateKeys(t, key)

			result, _ := showContinue(t, printer)

			assert.Equal(t, expected, result)
		}
	})

	t.Run("enter picks the first option", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		result, output := showContinue(t, printer)

		assert.Equal(t, "retry", result)
		assert.Contains(t, stripANSI(output), "[Retry/abort]")
	})
}

func TestInteractiveContinuePrinter_CustomHandles(t *testing.T) {
	printer := pterm.DefaultInteractiveContinue.
		WithOptions([]string{"yes", "no", "always", "never"}).
		WithHandles([]string{"y", "n", "a", "N"})

	tests := []struct {
		name     string
		key      rune
		expected string
	}{
		{name: "lowercase n picks no", key: 'n', expected: "no"},
		{name: "uppercase N picks never", key: 'N', expected: "never"},
		{name: "a picks always", key: 'a', expected: "always"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			simulateKeys(t, tc.key)

			result, output := showContinue(t, printer)

			assert.Equal(t, tc.expected, result)
			assert.Contains(t, stripANSI(output), "[y/n/a/N]", "custom handles must be rendered as given")
		})
	}
}

func TestInteractiveContinuePrinter_MismatchedHandlesFallBackToDefaults(t *testing.T) {
	var printer *pterm.InteractiveContinuePrinter

	output := captureStdout(func(_ io.Writer) {
		printer = pterm.DefaultInteractiveContinue.WithHandles([]string{"only-one"})
	})

	assert.Contains(t, stripANSI(output), "not a valid set of handles", "a mismatched handle count must print a warning")
	assert.Equal(t, []string{"Yes", "no", "all", "cancel"}, printer.Handles, "invalid handles must fall back to the default handles")

	simulateKeys(t, 'n')

	result, _ := showContinue(t, printer)
	assert.Equal(t, "no", result)
}

func TestInteractiveContinuePrinter_InterruptCallsOnInterruptFunc(t *testing.T) {
	interrupted := false
	printer := pterm.DefaultInteractiveContinue.WithOnInterruptFunc(func() { interrupted = true })

	simulateKeys(t, keys.CtrlC)

	result, _ := showContinue(t, printer)

	assert.True(t, interrupted, "Ctrl+C must invoke the OnInterruptFunc")
	assert.Empty(t, result, "an interrupted prompt must not return an option")
}
