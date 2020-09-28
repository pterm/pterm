package pterm

import (
	"io"
	"os"
)

// Printo overrides the current line in a terminal.
// If the current line is empty, the text will be printed like with pterm.Print.
// To create a new line, which
// Example:
// pterm.Printo("Hello, World")
// time.Sleep(time.Second)
// pterm.Oprint("Hello, Earth!")
func Printo(a ...interface{}) {
	Fprint(os.Stdout, "\r"+Sprint(a...))
}

// Fprinto prints Printo to a custom writer.
func Fprinto(w io.Writer, a ...interface{}) {
	Fprint(w, "\r", Sprint(a...))
}
