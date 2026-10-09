package pterm_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/pterm/pterm"
)

var terminalWidth = 80
var terminalHeight = 60

func TestMain(m *testing.M) {
	// Make color rendering independent of the terminal (or CI runner) the
	// tests happen to run in: colors on, detection reporting true color.
	for _, name := range []string{"NO_COLOR", "FORCE_COLOR", "CLICOLOR", "CLICOLOR_FORCE"} {
		os.Unsetenv(name)
	}

	_ = os.Setenv("COLORTERM", "truecolor")

	pterm.EnableColor()

	pterm.SetForcedTerminalSize(terminalWidth, terminalHeight)
	setupStdoutCapture()

	exitVal := m.Run()

	teardownStdoutCapture()

	// A test that starts a live printer (spinner, progressbar, area) without
	// stopping it leaks a goroutine that keeps writing into shared state and
	// makes *later* tests fail in confusing ways. Fail deterministically here
	// instead, with the stacks of the leaked goroutines.
	if leaks := leakedPtermGoroutines(); leaks != "" && exitVal == 0 {
		fmt.Fprintf(os.Stderr, "FAIL: leaked pterm goroutines after test run:\n\n%s\n", leaks)

		exitVal = 1
	}

	os.Exit(exitVal)
}

// leakedPtermGoroutines returns the stacks of goroutines still running inside
// the pterm package after all tests finished. It retries for a short grace
// period so goroutines that are just shutting down are not reported.
func leakedPtermGoroutines() string {
	deadline := time.Now().Add(2 * time.Second)

	for {
		var leaks []string
		buf := make([]byte, 1<<20)

		n := runtime.Stack(buf, true)
		for stack := range strings.SplitSeq(string(buf[:n]), "\n\n") {
			if strings.Contains(stack, "github.com/pterm/pterm.") && !strings.Contains(stack, "pterm_test.TestMain") {
				leaks = append(leaks, stack)
			}
		}

		if len(leaks) == 0 {
			return ""
		}

		if time.Now().After(deadline) {
			return strings.Join(leaks, "\n\n")
		}

		time.Sleep(50 * time.Millisecond)
	}
}

// syncBuffer wraps bytes.Buffer with a mutex so live printer goroutines
// (progressbars, spinners) writing to it cannot race with the test goroutine
// reading or resetting it. bytes.Buffer is not safe for concurrent use, so
// without this, the race detector flags spinner/progressbar tests whenever a
// printer goroutine is still alive while readStdout runs.
type syncBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *syncBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.buf.Write(p)
}

func (b *syncBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.buf.String()
}

func (b *syncBuffer) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.buf.Reset()
}

func (b *syncBuffer) Read(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.buf.Read(p)
}

func (b *syncBuffer) Bytes() []byte {
	b.mu.Lock()
	defer b.mu.Unlock()

	out := make([]byte, b.buf.Len())
	copy(out, b.buf.Bytes())

	return out
}

var outBuf syncBuffer

// setupStdoutCapture sets up a fake stdout capture.
func setupStdoutCapture() {
	outBuf.Reset()
	pterm.SetDefaultOutput(&outBuf)
}

// teardownStdoutCapture restores the real stdout.
func teardownStdoutCapture() {
	pterm.SetDefaultOutput(os.Stdout)
}

// captureStdout simulates capturing of os.stdout with a buffer and returns what was written to the screen
func captureStdout(f func(w io.Writer)) string {
	setupStdoutCapture()
	f(&outBuf)

	return readStdout()
}

// readStdout reads the current stdout buffer. Assumes setupStdoutCapture() has been called before.
func readStdout() string {
	content := outBuf.String()
	outBuf.Reset()

	return content
}
