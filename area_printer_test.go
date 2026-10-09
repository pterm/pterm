package pterm_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// fdBuffer is a syncBuffer that also reports a (fake) file descriptor. The
// AreaPrinter — and the MultiPrinter, which renders through an area — only
// accepts writers on which the cursor can be moved, i.e. writers with an Fd
// method, so tests capture area output through this type.
type fdBuffer struct {
	syncBuffer
}

// Fd pretends the buffer is a terminal file descriptor.
func (b *fdBuffer) Fd() uintptr {
	return 0
}

// clearLineSequence is the ANSI sequence the area emits to erase a line
// before repainting.
const clearLineSequence = "\x1b[2K"

// lastAreaFrame returns the visible content of the most recent area repaint
// in s. Every repaint first clears the previous content with clear-line
// sequences, so the text after the final clear is what stays on screen.
func lastAreaFrame(s string) string {
	if idx := strings.LastIndex(s, clearLineSequence); idx >= 0 {
		s = s[idx+len(clearLineSequence):]
	}

	return stripTerminalEscapes(s)
}

// startTestArea starts the given area printer writing into a fresh fdBuffer.
func startTestArea(t *testing.T, printer *pterm.AreaPrinter, text ...any) (*pterm.AreaPrinter, *fdBuffer) {
	t.Helper()

	buf := &fdBuffer{}
	printer.SetWriter(buf)

	started, err := printer.Start(text...)
	require.NoError(t, err)
	t.Cleanup(func() { _ = started.Stop() })

	return started, buf
}

func TestAreaPrinter_StartRendersInitialContent(t *testing.T) {
	printer := pterm.AreaPrinter{}

	area, buf := startTestArea(t, &printer, "hello area")

	assert.Same(t, &printer, area, "Start must return the started instance")
	assert.Equal(t, "hello area", area.GetContent())
	assert.Equal(t, "hello area", lastAreaFrame(buf.String()))
}

func TestAreaPrinter_UpdateReplacesContentInPlace(t *testing.T) {
	area, buf := startTestArea(t, &pterm.AreaPrinter{}, "first content")

	area.Update("second content")

	out := buf.String()
	firstEnd := strings.Index(out, "first content") + len("first content")
	assert.Contains(t, out[firstEnd:], clearLineSequence, "the old content must be cleared before repainting")

	frame := lastAreaFrame(out)
	assert.Equal(t, "second content", frame, "only the new content may remain visible")
	assert.Equal(t, "second content", area.GetContent())
}

func TestAreaPrinter_MultipleUpdates(t *testing.T) {
	area, buf := startTestArea(t, &pterm.AreaPrinter{})

	for _, content := range []string{"one", "two", "three"} {
		area.Update(content)

		assert.Equal(t, content, area.GetContent())
		assert.Equal(t, content, lastAreaFrame(buf.String()))
	}
}

func TestAreaPrinter_MultilineUpdateClearsAllLines(t *testing.T) {
	area, buf := startTestArea(t, &pterm.AreaPrinter{}, "line 1\nline 2\nline 3")

	area.Update("replacement")

	out := buf.String()
	assert.Contains(t, out, "\x1b[1A"+clearLineSequence, "multi-line content must be cleared upwards line by line")
	assert.Equal(t, "replacement", lastAreaFrame(out))
}

func TestAreaPrinter_CenterCentersContent(t *testing.T) {
	area, buf := startTestArea(t, pterm.DefaultArea.WithCenter())

	area.Update("ab")

	line := strings.TrimRight(lastAreaFrame(buf.String()), "\n")
	assert.Equal(t, strings.Repeat(" ", 39)+"ab", line,
		"the content must be padded to the center of the 80 column terminal")
}

func TestAreaPrinter_FullscreenPadsToTerminalHeight(t *testing.T) {
	area, buf := startTestArea(t, pterm.DefaultArea.WithFullscreen())

	area.Update("content line")

	frame := lastAreaFrame(buf.String())
	assert.Equal(t, "content line", strings.TrimSpace(frame))
	assert.Equal(t, 58, strings.Count(frame, "\n"),
		"the content must be padded with blank lines to fill the 60 row terminal")
}

func TestAreaPrinter_FullscreenCenterPadsOnBothSides(t *testing.T) {
	area, buf := startTestArea(t, pterm.DefaultArea.WithFullscreen().WithCenter())

	area.Update("x")

	frame := lastAreaFrame(buf.String())
	assert.Equal(t, "x", strings.TrimSpace(frame))
	assert.Equal(t, 59, strings.Count(frame, "\n"))

	lines := strings.Split(frame, "\n")
	require.Greater(t, len(lines), 31)
	assert.Equal(t, "x", strings.TrimSpace(lines[30]), "the content must sit in the vertical center")
}

func TestAreaPrinter_ClearEmptiesArea(t *testing.T) {
	area, buf := startTestArea(t, &pterm.AreaPrinter{}, "visible content")

	buf.Reset()
	area.Clear()

	out := buf.String()
	assert.Contains(t, out, clearLineSequence, "clearing must erase the painted line")
	assert.Empty(t, strings.TrimSpace(stripTerminalEscapes(out)), "clearing must not paint any visible content")
}

func TestAreaPrinter_RemoveWhenDoneClearsOnStop(t *testing.T) {
	area, buf := startTestArea(t, pterm.DefaultArea.WithRemoveWhenDone(), "temporary content")

	buf.Reset()
	require.NoError(t, area.Stop())

	out := buf.String()
	assert.Contains(t, out, clearLineSequence, "stopping must clear the area")
	assert.Empty(t, strings.TrimSpace(stripTerminalEscapes(out)), "no visible content may remain after Stop")

	// A second Stop must be a silent no-op.
	buf.Reset()
	require.NoError(t, area.Stop())
	assert.Empty(t, buf.String())
}

func TestAreaPrinter_StopKeepsContentWithoutRemoveWhenDone(t *testing.T) {
	area, buf := startTestArea(t, &pterm.AreaPrinter{}, "kept content")

	buf.Reset()
	require.NoError(t, area.Stop())

	assert.Empty(t, buf.String(), "without RemoveWhenDone, Stop must not touch the rendered content")
}
