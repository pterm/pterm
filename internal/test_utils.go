package internal

import (
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// TestPrintContains can be used to test Print methods.
func TestPrintContains(t *testing.T, logic func(w io.Writer, a string)) {
	s := CaptureStdout(func(w io.Writer) {
		logic(w, "Hello, World!")
	})
	assert.Contains(t, s, "Hello, World!")
}

// TestPrintfContains can be used to test Printf methods.
func TestPrintfContains(t *testing.T, logic func(w io.Writer, format string, a string)) {
	s := CaptureStdout(func(w io.Writer) {
		logic(w, "Hello, %s!", "World")
	})
	assert.Contains(t, s, "Hello, World!")
}

// TestPrintlnContains can be used to test Println methods.
func TestPrintlnContains(t *testing.T, logic func(w io.Writer, a string)) {
	TestPrintContains(t, logic)
}

// TestSprintContains can be used to test Sprint methods.
func TestSprintContains(t *testing.T, logic func(a string) string) {
	assert.Contains(t, logic("Hello, World!"), "Hello, World!")
}

// TestSprintfContains can be used to test Sprintf methods.
func TestSprintfContains(t *testing.T, logic func(format string, a string) string) {
	assert.Contains(t, logic("Hello, %s!", "World"), "Hello, World!")
}

// TestSprintlnContains can be used to test Sprintln methods.
func TestSprintlnContains(t *testing.T, logic func(a string) string) {
	TestSprintContains(t, logic)
}

// CaptureStdout captures everything written to the terminal and returns it as a string.
func CaptureStdout(f func(w io.Writer)) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	color.SetOutput(w)

	f(w)

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = originalStdout
	color.SetOutput(w)

	return string(out)
}
