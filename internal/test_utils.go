package internal

import (
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestPrintContains(t *testing.T, logic func(w io.Writer, a string)) {
	s := CaptureStdout(func(w io.Writer) {
		logic(w, "Hello, World!")
	})
	assert.Contains(t, s, "Hello, World!")
}

func TestPrintfContains(t *testing.T, logic func(w io.Writer, format string, a string)) {
	s := CaptureStdout(func(w io.Writer) {
		logic(w, "Hello, %s!", "World")
	})
	assert.Contains(t, s, "Hello, World!")
}

func TestPrintlnContains(t *testing.T, logic func(w io.Writer, a string)) {
	TestPrintContains(t, logic)
}

func TestSprintContains(t *testing.T, logic func(a string) string) {
	assert.Contains(t, logic("Hello, World!"), "Hello, World!")
}

func TestSprintfContains(t *testing.T, logic func(format string, a string) string) {
	assert.Contains(t, logic("Hello, %s!", "World"), "Hello, World!")
}

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
