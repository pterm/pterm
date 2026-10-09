package pterm_test

import (
	"os"
	"testing"

	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

// showTextInput runs the given text input prompt (which must terminate
// through the key presses simulated beforehand) and returns the entered text.
func showTextInput(t *testing.T, printer *pterm.InteractiveTextInputPrinter, text ...string) string {
	t.Helper()

	result, err := showInteractive(t, func() (string, error) {
		return printer.Show(text...)
	})
	require.NoError(t, err)

	return result
}

func TestInteractiveTextInputPrinter_TypeAndSubmit(t *testing.T) {
	simulateKeys(t, "hi", keys.Enter)

	result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

	assert.Equal(t, "hi", result)
}

func TestInteractiveTextInputPrinter_SpacesArePartOfTheInput(t *testing.T) {
	simulateKeys(t, "hello", keys.Space, "world", keys.Enter)

	result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

	assert.Equal(t, "hello world", result)
}

func TestInteractiveTextInputPrinter_RendersPromptAndTypedText(t *testing.T) {
	var result string

	output := captureUserFacingStdout(t, func() {
		simulateKeys(t, "hello", keys.Enter)

		result = showTextInput(t, &pterm.DefaultInteractiveTextInput, "Your name")
	})

	assert.Equal(t, "hello", result)
	assert.Contains(t, stripANSI(output), "Your name: ", "the prompt and delimiter must be rendered")
	assert.Contains(t, stripANSI(output), "hello", "the typed text must be rendered")
}

func TestInteractiveTextInputPrinter_BackspaceRemovesLastRune(t *testing.T) {
	simulateKeys(t, "hix", keys.Backspace, keys.Enter)

	result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

	assert.Equal(t, "hi", result)
}

func TestInteractiveTextInputPrinter_CursorMovementEditsAtPosition(t *testing.T) {
	t.Run("left moves the insert position back", func(t *testing.T) {
		simulateKeys(t, "ac", keys.Left, 'b', keys.Enter)

		result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

		assert.Equal(t, "abc", result)
	})

	t.Run("right moves the insert position forward again", func(t *testing.T) {
		simulateKeys(t, "ab", keys.Left, keys.Left, keys.Right, 'c', keys.Enter)

		result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

		assert.Equal(t, "acb", result)
	})

	t.Run("delete removes the rune under the cursor", func(t *testing.T) {
		simulateKeys(t, "abc", keys.Left, keys.Left, keys.Delete, keys.Enter)

		result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

		assert.Equal(t, "ac", result)
	})
}

func TestInteractiveTextInputPrinter_DefaultValue(t *testing.T) {
	t.Run("enter submits the default value", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithDefaultValue("pterm"))

		assert.Equal(t, "pterm", result)
	})

	t.Run("typing appends to the default value", func(t *testing.T) {
		simulateKeys(t, 'd', keys.Enter)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithDefaultValue("abc"))

		assert.Equal(t, "abcd", result)
	})

	t.Run("delete discards the default value", func(t *testing.T) {
		simulateKeys(t, keys.Delete, 'x', keys.Enter)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithDefaultValue("abc"))

		assert.Equal(t, "x", result)
	})

	t.Run("backspace edits the default value", func(t *testing.T) {
		simulateKeys(t, keys.Backspace, keys.Enter)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithDefaultValue("abc"))

		assert.Equal(t, "ab", result)
	})
}

func TestInteractiveTextInputPrinter_MaskRendersMaskButReturnsText(t *testing.T) {
	var result string

	output := captureUserFacingStdout(t, func() {
		simulateKeys(t, "abc", keys.Enter)

		result = showTextInput(t, pterm.DefaultInteractiveTextInput.WithMask("*"))
	})

	assert.Equal(t, "abc", result, "the returned text must be the real input, not the mask")
	assert.Contains(t, stripANSI(output), "***", "the rendered input must be masked")
	assert.NotContains(t, output, "abc", "the real input must never be rendered")
}

func TestInteractiveTextInputPrinter_MultiLine(t *testing.T) {
	t.Run("enter inserts a newline and tab submits", func(t *testing.T) {
		var result string

		output := captureUserFacingStdout(t, func() {
			simulateKeys(t, "a", keys.Enter, "b", keys.Tab)

			result = showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine())
		})

		assert.Equal(t, "a\nb", result)
		assert.Contains(t, stripANSI(output), "[Press tab to submit]", "the multi-line hint must be rendered")
	})

	t.Run("up navigates to the previous line", func(t *testing.T) {
		simulateKeys(t, keys.Enter, "second", keys.Up, "first", keys.Tab)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine())

		assert.Equal(t, "first\nsecond", result)
	})

	t.Run("down navigates to the next line", func(t *testing.T) {
		simulateKeys(t, "a", keys.Enter, keys.Enter, keys.Up, "b", keys.Down, "c", keys.Tab)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine())

		assert.Equal(t, "a\nb\nc", result)
	})

	t.Run("backspace joins lines", func(t *testing.T) {
		simulateKeys(t, "a", keys.Enter, "b", keys.Backspace, keys.Backspace, 'c', keys.Tab)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine())

		assert.Equal(t, "ac", result)
	})

	t.Run("delete joins with the next line", func(t *testing.T) {
		simulateKeys(t, "a", keys.Enter, "b", keys.Up, keys.Delete, keys.Tab)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine())

		assert.Equal(t, "ab", result)
	})

	t.Run("tab submits an untouched default value", func(t *testing.T) {
		simulateKeys(t, keys.Tab)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine().WithDefaultValue("keep"))

		assert.Equal(t, "keep", result)
	})

	t.Run("enter submits an untouched default value", func(t *testing.T) {
		simulateKeys(t, keys.Enter)

		result := showTextInput(t, pterm.DefaultInteractiveTextInput.WithMultiLine().WithDefaultValue("keep"))

		assert.Equal(t, "keep", result)
	})
}

func TestInteractiveTextInputPrinter_InterruptCallsOnInterruptFunc(t *testing.T) {
	interrupted := false
	printer := pterm.DefaultInteractiveTextInput.WithOnInterruptFunc(func() { interrupted = true })

	simulateKeys(t, "abc", keys.CtrlC)

	result := showTextInput(t, printer)

	assert.True(t, interrupted, "Ctrl+C must invoke the OnInterruptFunc")
	// The printer returns the partial input on interrupt; aborting is the
	// OnInterruptFunc's responsibility (the default one exits the process).
	assert.Equal(t, "abc", result)
}

func TestInteractiveTextInputPrinter_InterruptExitsByDefault(t *testing.T) {
	exitCode := -1
	internal.DefaultExitFunc = func(code int) { exitCode = code }

	t.Cleanup(func() { internal.DefaultExitFunc = os.Exit })

	simulateKeys(t, keys.CtrlC)

	result := showTextInput(t, &pterm.DefaultInteractiveTextInput)

	assert.Empty(t, result)
	assert.Equal(t, 1, exitCode, "without an OnInterruptFunc, Ctrl+C must exit with code 1")
}
