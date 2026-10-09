package pterm_test

import (
	"io"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// startTestSpinner starts a spinner writing into a fresh syncBuffer. The huge
// delay freezes the animation after its first frame, so tests observe exactly
// the frames they trigger themselves (Stop interrupts the delay, so stopping
// stays instant).
func startTestSpinner(t *testing.T, printer *pterm.SpinnerPrinter, text string) (*pterm.SpinnerPrinter, *syncBuffer) {
	t.Helper()

	buf := &syncBuffer{}

	spinner, err := printer.WithDelay(time.Hour).WithWriter(buf).Start(text)
	require.NoError(t, err)
	t.Cleanup(func() { _ = spinner.Stop() })

	waitForOutput(t, buf, text)

	return spinner, buf
}

func TestSpinnerPrinter_StartRendersFirstFrameAndText(t *testing.T) {
	printer := pterm.DefaultSpinner.WithSequence("+").WithShowTimer(false)

	_, buf := startTestSpinner(t, printer, "loading files")

	assert.Equal(t, "+ loading files", lastFrame(buf.String()),
		"the first frame must consist of the first sequence glyph and the text")
}

func TestSpinnerPrinter_SequenceCycles(t *testing.T) {
	buf := &syncBuffer{}

	spinner, err := pterm.DefaultSpinner.
		WithSequence("AAA", "BBB").
		WithShowTimer(false).
		WithDelay(time.Millisecond).
		WithWriter(buf).
		Start("cycling")
	require.NoError(t, err)

	defer spinner.Stop()

	waitForOutput(t, buf, "AAA")
	waitForOutput(t, buf, "BBB")
}

func TestSpinnerPrinter_UpdateTextReplacesText(t *testing.T) {
	printer := pterm.DefaultSpinner.WithSequence("+").WithShowTimer(false)

	spinner, buf := startTestSpinner(t, printer, "initial text")

	spinner.UpdateText("updated text")
	waitForOutput(t, buf, "updated text")

	frame := lastFrame(buf.String())
	assert.Contains(t, frame, "updated text")
	assert.Contains(t, frame, "+", "the current sequence glyph must be preserved")
	assert.NotContains(t, frame, "initial text", "the old text must be gone from the current frame")
}

func TestSpinnerPrinter_ShowTimerRendersElapsedTime(t *testing.T) {
	buf := &syncBuffer{}

	spinner, err := pterm.DefaultSpinner.
		WithShowTimer(true).
		WithTimerRoundingFactor(time.Second).
		WithDelay(time.Millisecond).
		WithWriter(buf).
		Start("timing")
	require.NoError(t, err)

	defer spinner.Stop()

	spinner.SetStartedAt(time.Now().Add(-90 * time.Second))

	timerRegexp := regexp.MustCompile(`\(1m3\ds\)`)

	waitFor(t, func() bool {
		return timerRegexp.MatchString(stripTerminalEscapes(buf.String()))
	}, func() string {
		return "the timer never rendered the elapsed time, got:\n" + buf.String()
	})
}

func TestSpinnerPrinter_ShowTimerDisabled(t *testing.T) {
	printer := pterm.DefaultSpinner.WithSequence("+").WithShowTimer(false)

	_, buf := startTestSpinner(t, printer, "no timer")

	assert.NotContains(t, lastFrame(buf.String()), "(", "no timer may be rendered when ShowTimer is off")
}

func TestSpinnerPrinter_ResultPrinters(t *testing.T) {
	tests := []struct {
		name   string
		act    func(spinner *pterm.SpinnerPrinter)
		prefix string
		want   string
	}{
		{name: "Info", act: func(s *pterm.SpinnerPrinter) { s.Info("info message") }, prefix: "INFO", want: "info message"},
		{name: "Success", act: func(s *pterm.SpinnerPrinter) { s.Success("all done") }, prefix: "SUCCESS", want: "all done"},
		{name: "Warning", act: func(s *pterm.SpinnerPrinter) { s.Warning("be careful") }, prefix: "WARNING", want: "be careful"},
		{name: "Fail", act: func(s *pterm.SpinnerPrinter) { s.Fail("it broke") }, prefix: "ERROR", want: "it broke"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			spinner, buf := startTestSpinner(t, &pterm.DefaultSpinner, "working")

			tc.act(spinner)

			assert.False(t, spinner.IsActive, "the result printer must stop the spinner")

			frame := lastFrame(buf.String())
			assert.Contains(t, frame, tc.prefix, "the final line must carry the result printer's prefix")
			assert.Contains(t, frame, tc.want, "the final line must carry the message")
		})
	}
}

func TestSpinnerPrinter_ResultPrinterReusesTextAsMessage(t *testing.T) {
	spinner, buf := startTestSpinner(t, &pterm.DefaultSpinner, "reused text")

	spinner.Success()

	frame := lastFrame(buf.String())
	assert.Contains(t, frame, "SUCCESS")
	assert.Contains(t, frame, "reused text", "without arguments the spinner text must be reused as the message")
}

func TestSpinnerPrinter_ResultPrintersFallBackToDefaultPrinters(t *testing.T) {
	spinner := pterm.SpinnerPrinter{}

	out := captureStdout(func(_ io.Writer) {
		spinner.Success("fallback message")
	})

	visible := stripTerminalEscapes(out)
	assert.Contains(t, visible, "SUCCESS", "a zero-value spinner must fall back to the default result printers")
	assert.Contains(t, visible, "fallback message")
}

func TestSpinnerPrinter_RemoveWhenDoneClearsFrame(t *testing.T) {
	printer := pterm.DefaultSpinner.WithRemoveWhenDone().WithShowTimer(false)

	spinner, buf := startTestSpinner(t, printer, "temporary")

	require.NoError(t, spinner.Stop())

	out := buf.String()
	assert.True(t, strings.HasSuffix(out, "\r"+strings.Repeat(" ", 80)+"\r"),
		"stopping must blank the line and return the cursor to its start, got:\n%q", out)
	assert.Empty(t, lastFrame(out), "no spinner content may remain visible after Stop")
}

func TestSpinnerPrinter_StopKeepsLastFrameWithoutRemoveWhenDone(t *testing.T) {
	printer := pterm.DefaultSpinner.WithSequence("+").WithShowTimer(false)

	spinner, buf := startTestSpinner(t, printer, "persistent")

	require.NoError(t, spinner.Stop())

	out := buf.String()
	assert.True(t, strings.HasSuffix(out, "\n"), "stopping must finish the spinner line with a newline")
	assert.Contains(t, lastFrame(out), "persistent", "the last frame must stay visible")
}

// TestSpinnerPrinter_RawOutput covers the raw output mode: no escape codes are
// emitted, UpdateText prints plain lines, and Stop still deactivates the
// spinner (regression for https://github.com/pterm/pterm/issues/763).
func TestSpinnerPrinter_RawOutput(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	buf := &syncBuffer{}

	spinner, err := pterm.DefaultSpinner.WithDelay(time.Millisecond).WithWriter(buf).Start("raw mode")
	require.NoError(t, err)

	defer spinner.Stop()

	assert.True(t, spinner.IsActive)

	spinner.UpdateText("raw update")
	waitForOutput(t, buf, "raw update")

	require.NoError(t, spinner.Stop())
	assert.False(t, spinner.IsActive, "Stop must deactivate the spinner in raw output mode")

	assert.NotContains(t, buf.String(), "\x1b[", "raw output must not contain escape codes")
}
