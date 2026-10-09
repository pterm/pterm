package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// selectPrompt is passed to every Show() call in this file. It doubles as the
// frame marker for areaWriter.frames, so it must not occur in any option.
const selectPrompt = "Pick one"

// showSelect runs the given select menu (which must terminate through the key
// presses simulated beforehand) and returns the chosen option.
func showSelect(t *testing.T, printer *pterm.InteractiveSelectPrinter) (string, error) {
	t.Helper()

	return showInteractive(t, func() (string, error) {
		return printer.Show(selectPrompt)
	})
}

func TestInteractiveSelectPrinter_NavigateAndSelect(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"})

	simulateKeys(t, keys.Down, keys.Down, keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "c", result)

	frames := area.frames(selectPrompt)
	require.NotEmpty(t, frames)
	assert.Contains(t, frames[len(frames)-1], "❯ c", "the finished menu must show the chosen option next to the selector")
}

func TestInteractiveSelectPrinter_NavigationWrapsAroundEnds(t *testing.T) {
	t.Run("down from the last option wraps to the first", func(t *testing.T) {
		captureAreaOutput(t)
		simulateKeys(t, keys.Down, keys.Down, keys.Down, keys.Enter)

		result, err := showSelect(t, pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}))

		require.NoError(t, err)
		assert.Equal(t, "a", result)
	})

	t.Run("up from the first option wraps to the last", func(t *testing.T) {
		captureAreaOutput(t)
		simulateKeys(t, keys.Up, keys.Enter)

		result, err := showSelect(t, pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}))

		require.NoError(t, err)
		assert.Equal(t, "c", result)
	})
}

func TestInteractiveSelectPrinter_AlternateNavigationKeys(t *testing.T) {
	captureAreaOutput(t)
	simulateKeys(t, keys.CtrlN, keys.CtrlN, keys.CtrlP, keys.Enter)

	result, err := showSelect(t, pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}))

	require.NoError(t, err)
	assert.Equal(t, "b", result, "CtrlN must move down and CtrlP must move up")
}

func TestInteractiveSelectPrinter_DefaultOptionStartsSelection(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}).WithDefaultOption("b")

	simulateKeys(t, keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "b", result)

	frames := area.frames(selectPrompt)
	require.NotEmpty(t, frames)
	assert.Contains(t, frames[0], "❯ b", "the first frame must highlight the default option")
	assert.Contains(t, frames[0], "  a", "other options must be rendered without the selector")
}

func TestInteractiveSelectPrinter_MaxHeightScrollsWindow(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"a", "b", "c", "d", "e", "f"}).
		WithMaxHeight(3)

	simulateKeys(t, keys.Down, keys.Down, keys.Down, keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "d", result)

	frames := area.frames(selectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)

	// The last menu frame (the finished frame comes after it) must show the
	// scrolled three-option window b..d with the selector on d.
	lastMenu := frames[len(frames)-2]
	assert.Contains(t, lastMenu, "  b\n")
	assert.Contains(t, lastMenu, "  c\n")
	assert.Contains(t, lastMenu, "❯ d\n")
	assert.NotContains(t, lastMenu, "  a\n", "options scrolled out at the top must not be rendered")
	assert.NotContains(t, lastMenu, "  e\n", "options below the window must not be rendered")
}

func TestInteractiveSelectPrinter_DefaultOptionScrollsWindow(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"a", "b", "c", "d", "e", "f"}).
		WithMaxHeight(3).
		WithDefaultOption("e")

	simulateKeys(t, keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "e", result)

	frames := area.frames(selectPrompt)
	require.NotEmpty(t, frames)
	assert.Contains(t, frames[0], "  d\n")
	assert.Contains(t, frames[0], "❯ e\n")
	assert.Contains(t, frames[0], "  f\n")
	assert.NotContains(t, frames[0], "  a\n", "the window must scroll to make the default option visible")
}

func TestInteractiveSelectPrinter_FilterNarrowsOptions(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"apple", "banana", "cherry"})

	simulateKeys(t, 'b', 'a', 'n', keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "banana", result)

	frames := area.frames(selectPrompt)
	require.GreaterOrEqual(t, len(frames), 2)

	filtered := frames[len(frames)-2]
	assert.Contains(t, filtered, "[type to search]: ban", "the typed filter must be rendered next to the prompt")
	assert.Contains(t, filtered, "❯ banana")
	assert.NotContains(t, filtered, "apple", "options that do not match the filter must disappear")
	assert.NotContains(t, filtered, "cherry")
}

func TestInteractiveSelectPrinter_FilterIsCaseInsensitive(t *testing.T) {
	captureAreaOutput(t)
	simulateKeys(t, 'B', 'A', 'N', keys.Enter)

	result, err := showSelect(t, pterm.DefaultInteractiveSelect.WithOptions([]string{"apple", "banana", "cherry"}))

	require.NoError(t, err)
	assert.Equal(t, "banana", result)
}

func TestInteractiveSelectPrinter_FilterSupportsSpaces(t *testing.T) {
	captureAreaOutput(t)
	simulateKeys(t, 'd', keys.Space, 'a', keys.Enter)

	result, err := showSelect(t, pterm.DefaultInteractiveSelect.WithOptions([]string{"red apple", "green pear"}))

	require.NoError(t, err)
	assert.Equal(t, "red apple", result)
}

func TestInteractiveSelectPrinter_BackspaceRestoresOptions(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"apple", "banana"})

	// 'z' matches nothing, so the first enter must be ignored; backspace
	// removes the filter and the second enter picks the first option again.
	simulateKeys(t, 'z', keys.Enter, keys.Backspace, keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "apple", result)

	frames := area.frames(selectPrompt)
	require.GreaterOrEqual(t, len(frames), 3)
	assert.NotContains(t, frames[2], "apple", "a filter without matches must render no options")
}

func TestInteractiveSelectPrinter_SpaceIsIgnoredWithoutFilter(t *testing.T) {
	area := captureAreaOutput(t)
	printer := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}).WithFilter(false)

	simulateKeys(t, keys.Space, keys.Down, keys.Enter)

	result, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.Equal(t, "b", result, "space must not corrupt the menu when filtering is disabled")

	frames := area.frames(selectPrompt)
	require.NotEmpty(t, frames)
	assert.NotContains(t, frames[0], "[type to search]", "the filter placeholder must not be rendered when filtering is disabled")
}

func TestInteractiveSelectPrinter_NoOptionsReturnsError(t *testing.T) {
	captureAreaOutput(t)

	_, err := showSelect(t, pterm.DefaultInteractiveSelect.WithOptions(nil))

	assert.ErrorContains(t, err, "no options provided")
}

func TestInteractiveSelectPrinter_InterruptCallsOnInterruptFunc(t *testing.T) {
	captureAreaOutput(t)

	interrupted := false
	printer := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"a", "b"}).
		WithOnInterruptFunc(func() { interrupted = true })

	simulateKeys(t, keys.CtrlC)

	_, err := showSelect(t, printer)

	require.NoError(t, err)
	assert.True(t, interrupted, "Ctrl+C must invoke the OnInterruptFunc")
}
