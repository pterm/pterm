package pterm_test

import (
	"io"
	"testing"

	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// showConfirm runs the given confirm prompt (which must terminate through the
// key presses simulated beforehand) and returns its result together with
// everything it printed.
func showConfirm(t *testing.T, printer *pterm.InteractiveConfirmPrinter, text ...string) (result bool, output string) {
	t.Helper()

	var err error

	output = captureStdout(func(_ io.Writer) {
		result, err = showInteractive(t, func() (bool, error) {
			return printer.Show(text...)
		})
	})

	require.NoError(t, err)

	return result, output
}

func TestInteractiveConfirmPrinter_AnswerKeys(t *testing.T) {
	tests := []struct {
		name     string
		key      rune
		expected bool
		answer   string
	}{
		{name: "y confirms", key: 'y', expected: true, answer: "Yes"},
		{name: "Y confirms case-insensitively", key: 'Y', expected: true, answer: "Yes"},
		{name: "n rejects", key: 'n', expected: false, answer: "No"},
		{name: "N rejects case-insensitively", key: 'N', expected: false, answer: "No"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			simulateKeys(t, tc.key)

			result, output := showConfirm(t, &pterm.DefaultInteractiveConfirm)

			assert.Equal(t, tc.expected, result)
			assert.Contains(t, stripANSI(output), tc.answer, "the chosen answer must be echoed after the prompt")
		})
	}
}

func TestInteractiveConfirmPrinter_EnterReturnsDefaultValue(t *testing.T) {
	tests := []struct {
		name         string
		defaultValue bool
		answer       string
	}{
		{name: "default false", defaultValue: false, answer: "No"},
		{name: "default true", defaultValue: true, answer: "Yes"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			simulateKeys(t, keys.Enter)

			result, output := showConfirm(t, pterm.DefaultInteractiveConfirm.WithDefaultValue(tc.defaultValue))

			assert.Equal(t, tc.defaultValue, result)
			assert.Contains(t, stripANSI(output), tc.answer, "pressing enter must echo the default answer")
		})
	}
}

func TestInteractiveConfirmPrinter_PromptShowsDefaultInSuffix(t *testing.T) {
	t.Run("default false uppercases n", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		_, output := showConfirm(t, pterm.DefaultInteractiveConfirm.WithDefaultValue(false), "Proceed with the deployment?")

		assert.Contains(t, stripANSI(output), "Proceed with the deployment? [y/N]: ")
	})

	t.Run("default true uppercases y", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		_, output := showConfirm(t, pterm.DefaultInteractiveConfirm.WithDefaultValue(true), "Proceed with the deployment?")

		assert.Contains(t, stripANSI(output), "Proceed with the deployment? [Y/n]: ")
	})
}

func TestInteractiveConfirmPrinter_CustomAnswerTexts(t *testing.T) {
	printer := pterm.DefaultInteractiveConfirm.WithConfirmText("Absolutely").WithRejectText("Denied")

	t.Run("suffix is derived from the custom texts", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		_, output := showConfirm(t, printer)

		assert.Contains(t, stripANSI(output), "[a/D]", "the short handles must come from the custom answer texts")
	})

	tests := []struct {
		name     string
		key      rune
		expected bool
		answer   string
	}{
		{name: "custom confirm lower", key: 'a', expected: true, answer: "Absolutely"},
		{name: "custom confirm upper", key: 'A', expected: true, answer: "Absolutely"},
		{name: "custom reject lower", key: 'd', expected: false, answer: "Denied"},
		{name: "custom reject upper", key: 'D', expected: false, answer: "Denied"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			simulateKeys(t, tc.key)

			result, output := showConfirm(t, printer)

			assert.Equal(t, tc.expected, result)
			assert.Contains(t, stripANSI(output), tc.answer)
		})
	}
}

func TestInteractiveConfirmPrinter_IgnoresUnrelatedKeys(t *testing.T) {
	simulateKeys(t, 'x', '1', keys.Down, 'y')

	result, output := showConfirm(t, &pterm.DefaultInteractiveConfirm)

	assert.True(t, result, "unrelated keys must be ignored until a valid answer is pressed")
	assert.Contains(t, stripANSI(output), "Yes")
}

func TestInteractiveConfirmPrinter_InterruptCallsOnInterruptFunc(t *testing.T) {
	interrupted := false
	printer := pterm.DefaultInteractiveConfirm.WithOnInterruptFunc(func() { interrupted = true })

	simulateKeys(t, keys.CtrlC)

	result, _ := showConfirm(t, printer)

	assert.True(t, interrupted, "Ctrl+C must invoke the OnInterruptFunc")
	assert.False(t, result, "an interrupted prompt must not confirm")
}
