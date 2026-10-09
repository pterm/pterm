package pterm_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// lastAreaFrameLines returns the individual visible lines of the most recent
// MultiPrinter repaint in s.
func lastAreaFrameLines(s string) []string {
	return strings.Split(strings.TrimRight(lastAreaFrame(s), "\n"), "\n")
}

// waitForFrameLines polls buf until the most recent repaint consists exactly
// of want (one string per line).
func waitForFrameLines(t *testing.T, buf *fdBuffer, want []string) {
	t.Helper()

	waitFor(t, func() bool {
		lines := lastAreaFrameLines(buf.String())
		if len(lines) != len(want) {
			return false
		}

		for i, line := range lines {
			if line != want[i] {
				return false
			}
		}

		return true
	}, func() string {
		return fmt.Sprintf("the area never rendered the frame %q, got:\n%q", want, buf.String())
	})
}

func TestMultiPrinter_NewWriterReturnsDistinctWriters(t *testing.T) {
	multi := pterm.DefaultMultiPrinter

	writerOne := multi.NewWriter()
	writerTwo := multi.NewWriter()

	require.NotNil(t, writerOne)
	require.NotNil(t, writerTwo)
	assert.NotSame(t, writerOne, writerTwo, "every NewWriter call must return its own writer")
}

func TestMultiPrinter_RendersWritersOnSeparateLines(t *testing.T) {
	buf := &fdBuffer{}
	multi := pterm.DefaultMultiPrinter.WithWriter(buf).WithUpdateDelay(time.Millisecond * 5)

	writerOne := multi.NewWriter()
	writerTwo := multi.NewWriter()

	started, err := multi.Start()
	require.NoError(t, err)

	defer started.Stop()

	fmt.Fprint(writerOne, "output of printer one")
	fmt.Fprint(writerTwo, "output of printer two")

	// Each writer must get its own line, in registration order, without the
	// contents bleeding into each other.
	waitForFrameLines(t, buf, []string{"output of printer one", "output of printer two"})
}

func TestMultiPrinter_CarriageReturnKeepsLastWrite(t *testing.T) {
	buf := &fdBuffer{}
	multi := pterm.DefaultMultiPrinter.WithWriter(buf).WithUpdateDelay(time.Millisecond * 5)

	writer := multi.NewWriter()

	started, err := multi.Start()
	require.NoError(t, err)

	defer started.Stop()

	// Live printers overwrite their line via "\r"; only the last override may
	// be rendered.
	fmt.Fprint(writer, "first pass\rsecond pass\rfinal pass")

	waitForFrameLines(t, buf, []string{"final pass"})
}

func TestMultiPrinter_SkipsWritersWithoutVisibleContent(t *testing.T) {
	buf := &fdBuffer{}
	multi := pterm.DefaultMultiPrinter.WithWriter(buf).WithUpdateDelay(time.Millisecond * 5)

	emptyWriter := multi.NewWriter()
	contentWriter := multi.NewWriter()

	started, err := multi.Start()
	require.NoError(t, err)

	defer started.Stop()

	fmt.Fprint(emptyWriter, "\r\n")
	fmt.Fprint(contentWriter, "visible line")

	waitForFrameLines(t, buf, []string{"visible line"})
}

func TestMultiPrinter_LivePrintersRenderInside(t *testing.T) {
	buf := &fdBuffer{}
	multi := pterm.DefaultMultiPrinter.WithWriter(buf).WithUpdateDelay(time.Millisecond * 5)

	spinner, err := pterm.DefaultSpinner.
		WithShowTimer(false).
		WithDelay(time.Millisecond * 5).
		WithWriter(multi.NewWriter()).
		Start("spinner alpha")
	require.NoError(t, err)

	defer spinner.Stop()

	bar, err := pterm.DefaultProgressbar.
		WithTotal(10).
		WithCurrent(5).
		WithShowElapsedTime(false).
		WithWriter(multi.NewWriter()).
		Start("bar beta")
	require.NoError(t, err)

	defer bar.Stop()

	started, err := multi.Start()
	require.NoError(t, err)

	defer started.Stop()

	waitFor(t, func() bool {
		lines := lastAreaFrameLines(buf.String())

		return len(lines) == 2 &&
			strings.Contains(lines[0], "spinner alpha") && !strings.Contains(lines[0], "bar beta") &&
			strings.Contains(lines[1], "bar beta") && strings.Contains(lines[1], "50%") &&
			!strings.Contains(lines[1], "spinner alpha")
	}, func() string {
		return fmt.Sprintf("spinner and progressbar never rendered on their own lines, got:\n%q", buf.String())
	})
}

// TestLivePrinter_StartStopContract verifies the shared live printer
// lifecycle for all four live printers: Start returns the started instance
// (leaving the original untouched where Start forks a copy), the activity
// state flips with Start/Stop, and stopping twice is safe.
func TestLivePrinter_StartStopContract(t *testing.T) {
	t.Run("SpinnerPrinter", func(t *testing.T) {
		buf := &syncBuffer{}
		original := pterm.DefaultSpinner.WithDelay(time.Hour).WithWriter(buf)

		started, err := original.Start()
		require.NoError(t, err)
		assert.NotSame(t, original, started, "Start must return a fork")
		assert.False(t, original.IsActive, "the original printer must stay untouched")
		assert.True(t, started.IsActive)

		require.NoError(t, started.Stop())
		assert.False(t, started.IsActive)
		require.NoError(t, started.Stop(), "a second Stop must be safe")
	})

	t.Run("ProgressbarPrinter", func(t *testing.T) {
		buf := &syncBuffer{}
		original := pterm.DefaultProgressbar.WithShowElapsedTime(false).WithWriter(buf)

		started, err := original.Start()
		require.NoError(t, err)
		assert.NotSame(t, original, started, "Start must return a fork")
		assert.False(t, original.IsActive, "the original printer must stay untouched")
		assert.True(t, started.IsActive)

		_, err = started.Stop()
		require.NoError(t, err)
		assert.False(t, started.IsActive)

		_, err = started.Stop()
		require.NoError(t, err, "a second Stop must be safe")
	})

	t.Run("AreaPrinter", func(t *testing.T) {
		buf := &fdBuffer{}
		original := pterm.AreaPrinter{}
		original.SetWriter(buf)

		started, err := original.Start("area content")
		require.NoError(t, err)
		assert.Same(t, &original, started, "the AreaPrinter starts in place")

		require.NoError(t, started.Stop())
		require.NoError(t, started.Stop(), "a second Stop must be safe")
	})

	t.Run("MultiPrinter", func(t *testing.T) {
		buf := &fdBuffer{}
		multi := pterm.DefaultMultiPrinter.WithWriter(buf).WithUpdateDelay(time.Millisecond * 5)

		started, err := multi.Start()
		require.NoError(t, err)
		assert.Same(t, multi, started, "the MultiPrinter starts in place")
		assert.True(t, started.IsActive)

		_, err = started.Stop()
		require.NoError(t, err)
		assert.False(t, started.IsActive)

		_, err = started.Stop()
		require.NoError(t, err, "a second Stop must be safe")
	})
}

// TestLivePrinter_GenericStartStopRoundtrip drives every live printer through
// the LivePrinter interface, which is how the MultiPrinter manages them.
func TestLivePrinter_GenericStartStopRoundtrip(t *testing.T) {
	tests := []struct {
		name    string
		printer func() pterm.LivePrinter
	}{
		{name: "SpinnerPrinter", printer: func() pterm.LivePrinter {
			return pterm.DefaultSpinner.WithDelay(time.Hour).WithWriter(&syncBuffer{})
		}},
		{name: "ProgressbarPrinter", printer: func() pterm.LivePrinter {
			return pterm.DefaultProgressbar.WithShowElapsedTime(false).WithWriter(&syncBuffer{})
		}},
		{name: "AreaPrinter", printer: func() pterm.LivePrinter {
			area := pterm.AreaPrinter{}
			area.SetWriter(&fdBuffer{})

			return &area
		}},
		{name: "MultiPrinter", printer: func() pterm.LivePrinter {
			return pterm.DefaultMultiPrinter.WithWriter(&fdBuffer{}).WithUpdateDelay(time.Millisecond * 5)
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			printer := tc.printer()

			started, err := printer.GenericStart()
			require.NoError(t, err)
			require.NotNil(t, started)

			stopped, err := (*started).GenericStop()
			require.NoError(t, err)
			require.NotNil(t, stopped)

			// Stopping again through the generic interface must be safe.
			_, err = (*stopped).GenericStop()
			require.NoError(t, err)
		})
	}
}
