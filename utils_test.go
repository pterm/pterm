package pterm_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

var printables = []any{"Hello, World!", 1337, true, false, -1337, 'c', 1.5, "\\", "%s"}
var terminalWidth = 80
var terminalHeight = 60

func TestMain(m *testing.M) {
	pterm.SetForcedTerminalSize(terminalWidth, terminalHeight)
	setupStdoutCapture()
	exitVal := m.Run()
	teardownStdoutCapture()
	os.Exit(exitVal)
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
			testza.AssertContains(t, s, fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			t.Helper()

			s := captureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			testza.AssertContains(t, s, fmt.Sprint(printable))
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
			testza.AssertContains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := captureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			testza.AssertContains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
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
			testza.AssertContains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic(printable), fmt.Sprint(printable))
		})
		pterm.EnableStyling()
	}
}

// testSprintContainsWithoutError can be used to test Sprint methods which return an error.
func testSprintContainsWithoutError(t *testing.T, logic func(a any) (string, error)) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			testza.AssertContains(t, s, fmt.Sprint(printable))
			testza.AssertNoError(t, err)
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s, err := logic(printable)
			testza.AssertContains(t, s, fmt.Sprint(printable))
			testza.AssertNoError(t, err)
		})
		pterm.EnableStyling()
	}
}

// testSprintfContains can be used to test Sprintf methods.
func testSprintfContains(t *testing.T, logic func(format string, a any) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
		pterm.DisableStyling()
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			testza.AssertContains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
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

// testDoesOutput can be used to test if something is outputted to stdout.
func testDoesOutput(t *testing.T, logic func(w io.Writer)) {
	testza.AssertNotZero(t, captureStdout(logic))
	pterm.DisableStyling()
	testza.AssertNotZero(t, captureStdout(logic))
	pterm.EnableStyling()
}

// testEmpty checks that a function does not return a string.
func testEmpty(t *testing.T, logic func(a any) string) {
	for _, printable := range printables {
		testza.AssertZero(t, logic(printable))
		pterm.DisableStyling()
		testza.AssertZero(t, logic(printable))
		pterm.EnableStyling()
	}
}

// testDoesNotOutput can be used, to test that something does not output anything to stdout.
func testDoesNotOutput(t *testing.T, logic func(w io.Writer)) {
	testza.AssertZero(t, captureStdout(logic))
	pterm.DisableStyling()
	testza.AssertZero(t, captureStdout(logic))
	pterm.EnableStyling()
}

var outBuf bytes.Buffer

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

// readStdout reads the current stdout buffor. Assumes setupStdoutCapture() has been called before.
func readStdout() string {
	content := outBuf.String()
	outBuf.Reset()
	return content
}

func proxyToDevNull() {
	pterm.SetDefaultOutput(os.NewFile(0, os.DevNull))
}
