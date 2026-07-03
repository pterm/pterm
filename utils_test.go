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
	"github.com/stretchr/testify/assert"
)

var printables = []any{"Hello, World!", 1337, true, false, -1337, 'c', 1.5, "\\", "%s"}
var terminalWidth = 80
var terminalHeight = 60

func TestMain(m *testing.M) {
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

// testPrintContains can be used to test Print methods.
func testPrintContains(t *testing.T, logic func(w io.Writer, a any)) {
	t.Helper()

	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			t.Helper()

			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			assert.Contains(t, s, fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			t.Helper()

			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			assert.Contains(t, s, fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testPrintfContains can be used to test Printf methods.
func testPrintfContains(t *testing.T, logic func(w io.Writer, format string, a any)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			assert.Contains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			assert.Contains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.EnableStyling()
	}
}

// testPrintflnContains can be used to test Printfln methods.
func testPrintflnContains(t *testing.T, logic func(w io.Writer, format string, a any)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintfContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testPrintfContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testPrintlnContains can be used to test Println methods.
func testPrintlnContains(t *testing.T, logic func(w io.Writer, a any)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			t.Helper()

			testPrintContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			t.Helper()

			testPrintContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testSprintContains can be used to test Sprint methods.
func testSprintContains(t *testing.T, logic func(a any) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintContainsWithoutError can be used to test Sprint methods which return an error.
func testSprintContainsWithoutError(t *testing.T, logic func(a any) (string, error)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			assert.Contains(t, s, fmt.Sprint(printable))
			assert.NoError(t, err)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			assert.Contains(t, s, fmt.Sprint(printable))
			assert.NoError(t, err)
		})
		pterm.EnableStyling()
	}
}

// testSprintfContains can be used to test Sprintf methods.
func testSprintfContains(t *testing.T, logic func(format string, a any) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintflnContains can be used to test Sprintfln methods.
func testSprintflnContains(t *testing.T, logic func(format string, a any) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintfContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintfContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testSprintlnContains can be used to test Sprintln methods.
func testSprintlnContains(t *testing.T, logic func(a any) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintContains(t, logic)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testSprintContains(t, logic)
		})
		pterm.EnableStyling()
	}
}

// testEmpty checks that a function does not return a string.
func testEmpty(t *testing.T, logic func(a any) string) {
	for _, printable := range printables {
		assert.Zero(t, logic(printable))
		pterm.DisableStyling()
		assert.Zero(t, logic(printable))
		pterm.EnableStyling()
	}
}

// testDoesNotOutput can be used, to test that something does not output anything to stdout.
func testDoesNotOutput(t *testing.T, logic func(w io.Writer)) {
	assert.Zero(t, captureStdout(logic))
	pterm.DisableStyling()
	assert.Zero(t, captureStdout(logic))
	pterm.EnableStyling()
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

func proxyToDevNull() {
	pterm.SetDefaultOutput(os.NewFile(0, os.DevNull))
}
