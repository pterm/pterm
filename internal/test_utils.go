package internal

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var printables = []interface{}{"Hello, World!", 1337, true, false, 'c', 1.5, -1337, "\\", "%s"}

// TestPrintContains can be used to test Print methods.
func TestPrintContains(t *testing.T, logic func(w io.Writer, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := CaptureStdout(func(w io.Writer) {
				logic(w, printable)
			})
			assert.Contains(t, s, fmt.Sprint(printable))
		})
	}
}

// TestPrintfContains can be used to test Printf methods.
func TestPrintfContains(t *testing.T, logic func(w io.Writer, format string, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			s := CaptureStdout(func(w io.Writer) {
				logic(w, "Hello, %v!", printable)
			})
			assert.Contains(t, s, fmt.Sprintf("Hello, %v!", fmt.Sprint(printable)))
		})
	}
}

// TestPrintlnContains can be used to test Println methods.
func TestPrintlnContains(t *testing.T, logic func(w io.Writer, a interface{})) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			TestPrintContains(t, logic)
		})
	}
}

// TestSprintContains can be used to test Sprint methods.
func TestSprintContains(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic(printable), fmt.Sprint(printable))
		})
	}
}

// TestSprintfContains can be used to test Sprintf methods.
func TestSprintfContains(t *testing.T, logic func(format string, a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			assert.Contains(t, logic("Hello, %v!", printable), fmt.Sprintf("Hello, %v!", printable))
		})
	}
}

// TestSprintlnContains can be used to test Sprintln methods.
func TestSprintlnContains(t *testing.T, logic func(a interface{}) string) {
	for _, printable := range printables {
		t.Run(fmt.Sprint(printable), func(t *testing.T) {
			TestSprintContains(t, logic)
		})
	}
}

// TestDoesOutput can be used to test if something is outputted to stdout.
func TestDoesOutput(t *testing.T, logic func(w io.Writer)) {
	assert.NotEmpty(t, CaptureStdout(logic))
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
