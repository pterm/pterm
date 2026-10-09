package pterm_test

// Helpers shared by the interactive printer tests
// (interactive_*_printer_test.go).
//
// All interactive printers read key presses through the global keyboard mock
// (keyboard.SimulateKeyPress), so their tests are inherently bound to global
// state: they must never run in parallel, every Show() call has to consume
// exactly the key presses a test simulates, and a Show() that never terminates
// would hang the whole suite. The helpers below make those invariants explicit
// and enforce them with timeouts instead of hangs.

import (
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"atomicgo.dev/keyboard"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// interactiveTimeout is the maximum time an interactive Show() call may take
// before its test is failed. Simulated key presses are handled in
// microseconds, so hitting this always means a key was dropped or never sent.
const interactiveTimeout = 10 * time.Second

// simulateKeys queues the given key presses for the keyboard listener started
// by the next Show() call. Like keyboard.SimulateKeyPress it accepts runes,
// strings (one press per rune), keys.KeyCode and keys.Key values.
//
// The key presses are sent from a goroutine because the mock channel is
// unbuffered: each send blocks until the listener consumes it, which also
// guarantees the presses arrive strictly in order. The test fails if Show()
// returned without consuming every simulated press, because leftover presses
// would be delivered to the keyboard listener of a later test.
func simulateKeys(t *testing.T, keySequence ...any) {
	t.Helper()

	done := make(chan struct{})

	go func() {
		defer close(done)

		// keyboard.SimulateKeyPress only handles its first argument (it
		// returns inside its loop), so the keys are sent one by one.
		for _, key := range keySequence {
			_ = keyboard.SimulateKeyPress(key)
		}
	}()

	t.Cleanup(func() {
		select {
		case <-done:
		case <-time.After(interactiveTimeout):
			t.Error("Show() returned without consuming all simulated key presses; the leftover presses would corrupt the next interactive test")
		}
	})
}

// showInteractive runs an interactive printer's Show() function and fails the
// test if it does not return in time (e.g. because it is waiting for a key
// press that was never simulated) instead of hanging the whole suite.
func showInteractive[T any](t *testing.T, show func() (T, error)) (T, error) {
	t.Helper()

	type showResult struct {
		value T
		err   error
	}

	resultChan := make(chan showResult, 1)

	go func() {
		value, err := show()
		resultChan <- showResult{value: value, err: err}
	}()

	select {
	case result := <-resultChan:
		return result.value, result.err
	case <-time.After(interactiveTimeout):
		t.Fatal("Show() did not return; it is probably waiting for a key press that was never simulated")
		panic("unreachable")
	}
}

// ansiControlSequences matches every ANSI CSI sequence, including the cursor
// movement and line clearing codes written between live area frames
// (stripANSI only removes SGR color sequences).
var ansiControlSequences = regexp.MustCompile(`\x1b\[[0-9;?]*[A-Za-z]`)

// areaWriter is an in-memory cursor.Writer, so everything an interactive
// printer renders through pterm.DefaultArea can be captured and asserted on.
type areaWriter struct {
	syncBuffer
}

// Fd makes areaWriter satisfy cursor.Writer; there is no real file behind it.
func (w *areaWriter) Fd() uintptr { return 0 }

// frames splits everything that was written to the area into the individual
// rendered frames. Every frame starts with the prompt text, so the stream is
// stripped of ANSI control codes, normalized, and split on that marker.
func (w *areaWriter) frames(promptMarker string) []string {
	clean := ansiControlSequences.ReplaceAllString(w.String(), "")
	clean = strings.ReplaceAll(clean, "\r\n", "\n")
	clean = strings.ReplaceAll(clean, "\r", "")

	var frames []string

	for _, frame := range strings.Split(clean, promptMarker) {
		if frame == "" {
			continue
		}

		frames = append(frames, promptMarker+frame)
	}

	return frames
}

// captureAreaOutput redirects pterm.DefaultArea (which the select and
// multiselect printers render their live frames into) to an in-memory buffer
// and restores the default writer when the test finishes.
func captureAreaOutput(t *testing.T) *areaWriter {
	t.Helper()

	buf := &areaWriter{}
	pterm.DefaultArea.SetWriter(buf)

	t.Cleanup(func() { pterm.DefaultArea.SetWriter(os.Stdout) })

	return buf
}

// captureUserFacingStdout captures what f writes to the process' real
// os.Stdout. The text input printer renders its live frames through a raw
// cursor.Area that always writes to os.Stdout (it does not honor
// pterm.SetDefaultOutput), so asserting on its rendering requires swapping
// the os.Stdout file itself.
func captureUserFacingStdout(t *testing.T, f func()) string {
	t.Helper()

	readEnd, writeEnd, err := os.Pipe()
	require.NoError(t, err)

	original := os.Stdout
	os.Stdout = writeEnd

	defer func() {
		os.Stdout = original

		_ = writeEnd.Close()
		_ = readEnd.Close()
	}()

	f()

	os.Stdout = original

	require.NoError(t, writeEnd.Close())

	captured, err := io.ReadAll(readEnd)
	require.NoError(t, err)

	return string(captured)
}
