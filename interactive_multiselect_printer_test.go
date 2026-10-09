package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// multiselectPrompt is passed to every Show() call in this file. It doubles as
// the frame marker for areaWriter.frames, so it must not occur in any option.
const multiselectPrompt = "Pick some"

// showMultiselect runs the given multiselect menu (which must terminate
// through the key presses simulated beforehand) and returns the chosen
// options.
func showMultiselect(t *testing.T, printer *pterm.InteractiveMultiselectPrinter) ([]string, error) {
	t.Helper()

	return showInteractive(t, func() ([]string, error) {
		return printer.Show(multiselectPrompt)
	})
}

// checked and unchecked render an option row's checkbox with the plain theme
// checkmark atoms, e.g. "[✓] b" and "[✗] b".
func checked(option string) string {
	return "[" + stripANSI(pterm.ThemeDefault.Checkmark.Checked) + "] " + option
}

func unchecked(option string) string {
	return "[" + stripANSI(pterm.ThemeDefault.Checkmark.Unchecked) + "] " + option
}

func TestInteractiveMultiselectPrinter_SelectAndConfirm(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e"})

	// Enter is the default select key, Tab the default confirm key.
	simulateKeys(t, keys.Enter, keys.Down, keys.Enter, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, result, "confirm must return exactly the selected options in selection order")

	frames := area.frames(multiselectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)

	lastMenu := frames[len(frames)-2]
	assert.Contains(t, lastMenu, checked("a"), "selected options must show the checked checkmark")
	assert.Contains(t, lastMenu, "❯ "+checked("b"), "the highlighted row must show the selector")
	assert.Contains(t, lastMenu, unchecked("c"), "unselected options must show the unchecked checkmark")
	assert.Contains(t, lastMenu, "select")
	assert.Contains(t, lastMenu, "confirm")
}

func TestInteractiveMultiselectPrinter_SelectKeyTogglesSelection(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b"})

	simulateKeys(t, keys.Enter, keys.Enter, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Empty(t, result, "selecting an option twice must deselect it again")

	frames := area.frames(multiselectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)
	assert.Contains(t, frames[len(frames)-2], "❯ "+unchecked("a"), "a toggled-off option must render unchecked")
}

func TestInteractiveMultiselectPrinter_DefaultOptionsArePreselected(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"a", "b", "c"}).
		WithDefaultOptions([]string{"c", "a"})

	simulateKeys(t, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"c", "a"}, result, "the default options must be returned in the given order")

	frames := area.frames(multiselectPrompt)
	require.NotEmpty(t, frames)
	assert.Contains(t, frames[0], checked("a"))
	assert.Contains(t, frames[0], unchecked("b"))
	assert.Contains(t, frames[0], checked("c"))
}

func TestInteractiveMultiselectPrinter_SelectAllAndSelectNone(t *testing.T) {
	t.Run("right selects all options", func(t *testing.T) {
		captureAreaOutput(t)
		simulateKeys(t, keys.Right, keys.Tab)

		result, err := showMultiselect(t, pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c"}))

		require.NoError(t, err)
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("left deselects all options", func(t *testing.T) {
		captureAreaOutput(t)
		simulateKeys(t, keys.Right, keys.Left, keys.Tab)

		result, err := showMultiselect(t, pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c"}))

		require.NoError(t, err)
		assert.Empty(t, result)
	})
}

func TestInteractiveMultiselectPrinter_MaxHeightScrollsWindow(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"a", "b", "c", "d", "e", "f"}).
		WithMaxHeight(3)

	simulateKeys(t, keys.Down, keys.Down, keys.Down, keys.Enter, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"d"}, result)

	frames := area.frames(multiselectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)

	lastMenu := frames[len(frames)-2]
	assert.Contains(t, lastMenu, unchecked("b"))
	assert.Contains(t, lastMenu, unchecked("c"))
	assert.Contains(t, lastMenu, "❯ "+checked("d"))
	assert.NotContains(t, lastMenu, unchecked("a"), "options scrolled out at the top must not be rendered")
	assert.NotContains(t, lastMenu, unchecked("e"), "options below the window must not be rendered")
}

func TestInteractiveMultiselectPrinter_FilterSelectsMatch(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"apple", "banana", "cherry"})

	simulateKeys(t, 'b', 'a', 'n', keys.Enter, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"banana"}, result)

	frames := area.frames(multiselectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)

	lastMenu := frames[len(frames)-2]
	assert.Contains(t, lastMenu, "[type to search]: ban", "the typed filter must be rendered next to the prompt")
	assert.NotContains(t, lastMenu, "apple", "options that do not match the filter must disappear")
}

func TestInteractiveMultiselectPrinter_BackspaceRestoresOptions(t *testing.T) {
	captureAreaOutput(t)

	printer := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"apple", "banana"})

	// 'z' matches nothing, so the first select key press must be ignored;
	// backspace removes the filter and the second one selects "apple".
	simulateKeys(t, 'z', keys.Enter, keys.Backspace, keys.Enter, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"apple"}, result)
}

func TestInteractiveMultiselectPrinter_CustomSelectAndConfirmKeys(t *testing.T) {
	captureAreaOutput(t)

	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"a", "b", "c"}).
		WithFilter(false).
		WithKeySelect(keys.Space).
		WithKeyConfirm(keys.Enter)

	simulateKeys(t, keys.Space, keys.Down, keys.Space, keys.Enter)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"a", "b"}, result)
}

func TestInteractiveMultiselectPrinter_ShowSelectedOptions(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"a", "b"}).
		WithShowSelectedOptions()

	simulateKeys(t, keys.Enter, keys.Tab)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, []string{"a"}, result)

	frames := area.frames(multiselectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)
	assert.Contains(t, frames[len(frames)-2], "Selected: a", "the selected options must be listed below the menu")
}

func TestInteractiveMultiselectPrinter_SpaceKeyBindingsRequireDisabledFilter(t *testing.T) {
	t.Run("KeySelect", func(t *testing.T) {
		captureAreaOutput(t)

		_, err := showMultiselect(t, pterm.DefaultInteractiveMultiselect.
			WithOptions([]string{"a"}).
			WithKeySelect(keys.Space))

		assert.ErrorContains(t, err, "keys.Space")
	})

	t.Run("KeyConfirm", func(t *testing.T) {
		captureAreaOutput(t)

		_, err := showMultiselect(t, pterm.DefaultInteractiveMultiselect.
			WithOptions([]string{"a"}).
			WithKeyConfirm(keys.Space))

		assert.ErrorContains(t, err, "keys.Space")
	})
}

func TestInteractiveMultiselectPrinter_NoOptionsReturnsError(t *testing.T) {
	captureAreaOutput(t)

	_, err := showMultiselect(t, pterm.DefaultInteractiveMultiselect.WithOptions(nil))

	assert.ErrorContains(t, err, "no options provided")
}

func TestInteractiveMultiselectPrinter_InterruptCallsOnInterruptFunc(t *testing.T) {
	captureAreaOutput(t)

	interrupted := false
	printer := pterm.DefaultInteractiveMultiselect.
		WithOptions([]string{"a", "b"}).
		WithOnInterruptFunc(func() { interrupted = true })

	simulateKeys(t, keys.CtrlC)

	result, err := showMultiselect(t, printer)

	require.NoError(t, err)
	assert.True(t, interrupted, "Ctrl+C must invoke the OnInterruptFunc")
	assert.Empty(t, result, "an interrupted prompt must not return options")
}
