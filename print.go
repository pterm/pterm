package pterm

import (
	"fmt"
	"io"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/screen"
)

// Print writes to the default output like fmt.Print. Use it instead of fmt
// while live printers run, so the text appears above them.
func Print(a ...any) {
	emit(DefaultOutput(), fmt.Sprint(a...))
}

// Println writes to the default output like fmt.Println.
func Println(a ...any) {
	emit(DefaultOutput(), fmt.Sprintln(a...))
}

// Printf writes to the default output like fmt.Printf.
func Printf(format string, a ...any) {
	emit(DefaultOutput(), fmt.Sprintf(format, a...))
}

// Printfln writes to the default output like fmt.Printf and adds a newline.
func Printfln(format string, a ...any) {
	emit(DefaultOutput(), fmt.Sprintf(format, a...)+"\n")
}

// Fprint writes to w like fmt.Fprint, through w's renderer.
func Fprint(w io.Writer, a ...any) {
	emit(orDefault(w), fmt.Sprint(a...))
}

// Fprintln writes to w like fmt.Fprintln, through w's renderer.
func Fprintln(w io.Writer, a ...any) {
	emit(orDefault(w), fmt.Sprintln(a...))
}

// RemoveColorFromString returns s without escape sequences.
func RemoveColorFromString(s string) string {
	return screen.Strip(color.Strip(s))
}

// TerminalSize returns the width and height of the terminal behind the
// default output. Both are 0 when it is not a terminal and COLUMNS and
// LINES are unset.
func TerminalSize() (width, height int) {
	return screen.Size(DefaultOutput())
}

// TerminalWidth returns the width of the terminal behind the default output,
// or 80 when it is unknown.
func TerminalWidth() int {
	if w, _ := TerminalSize(); w > 0 {
		return w
	}

	return 80
}
