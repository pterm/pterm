package internal

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/gookit/color"
)

// CaptureStdout captures everything written to the terminal and returns it as a string.
func CaptureStdout(f func(w io.Writer)) string {
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	color.SetOutput(w)

	f(w)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = originalStdout
	color.SetOutput(originalStdout)

	return string(out)
}
