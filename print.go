package pterm

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pterm/pterm/internal"
)

// defaultWriter is the package-level default output writer.
//
// Reads from inside pterm should go through getDefaultWriter (or the public
// GetDefaultOutput) so they share the lock with SetDefaultOutput.
var defaultWriter io.Writer = os.Stdout

// SetDefaultOutput sets the default output of pterm.
func SetDefaultOutput(w io.Writer) {
	globalMu.Lock()
	defer globalMu.Unlock()

	defaultWriter = w
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...any) string {
	return fmt.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...any) string {
	return fmt.Sprintf(format, a...)
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintfln(format string, a ...any) string {
	return fmt.Sprintf(format, a...) + "\n"
}

// Sprintln returns what Println would print to the terminal.
func Sprintln(a ...any) string {
	str := fmt.Sprintln(a...)
	return Sprint(str)
}

// Sprinto returns what Printo would print.
func Sprinto(a ...any) string {
	return "\r" + Sprint(a...)
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...any) {
	Fprint(getDefaultWriter(), a...)
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...any) {
	Print(Sprintln(a...))
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...any) {
	Print(Sprintf(format, a...))
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Printfln(format string, a ...any) {
	Print(Sprintfln(format, a...))
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func PrintOnError(a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				Println(err)
			}
		}
	}
}

// PrintOnErrorf wraps every error which is not nil and prints it.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func PrintOnErrorf(format string, a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				Println(fmt.Errorf(format, err))
			}
		}
	}
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Fprint(writer io.Writer, a ...any) {
	if !outputEnabled() {
		return
	}

	var ret string
	var printed bool

	bars := snapshotProgressBars()
	for _, bar := range bars {
		if bar.isActive() && (bar.writer() == writer || bar.writer() == os.Stderr) {
			ret += sClearLine()
			ret += Sprinto(a...)
			printed = true
		}
	}

	spinners := snapshotSpinners()
	for _, spinner := range spinners {
		if spinner.isActive() && (spinner.writer() == writer || spinner.writer() == os.Stderr) {
			ret += sClearLine()

			ret += Sprinto(a...)
			printed = true
		}
	}

	if !printed {
		ret = Sprint(a...)
	}

	if writer == nil {
		writer = getDefaultWriter()
	}

	_, _ = fmt.Fprint(writer, ret)

	// Refresh all progressbars in case they were overwritten previously. Reference: #302
	for _, bar := range bars {
		if bar.isActive() {
			bar.UpdateTitle(bar.title())
		}
	}
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(writer io.Writer, a ...any) {
	Fprint(writer, Sprint(a...)+"\n")
}

// Printo overrides the current line in a terminal.
// If the current line is empty, the text will be printed like with pterm.Print.
// Example:
//
//	pterm.Printo("Hello, World")
//	time.Sleep(time.Second)
//	pterm.Printo("Hello, Earth!")
func Printo(a ...any) {
	if !outputEnabled() {
		return
	}

	if rawOutput() {
		Sprint(a...)
		return
	}

	_, _ = fmt.Fprint(getDefaultWriter(), "\r"+Sprint(a...))
}

// Fprinto prints Printo to a custom writer.
func Fprinto(w io.Writer, a ...any) {
	if !outputEnabled() {
		return
	}

	if w == nil {
		w = getDefaultWriter()
	}

	_, _ = fmt.Fprint(w, "\r", Sprint(a...))
}

// RemoveColorFromString removes color codes and OSC 8 hyperlinks from a string.
func RemoveColorFromString(a ...any) string {
	return internal.RemoveEscapeCodes(Sprint(a...))
}

func fClearLine(writer io.Writer) {
	if rawOutput() || writer == nil || !outputEnabled() {
		return
	}

	Fprinto(writer, strings.Repeat(" ", GetTerminalWidth()))
}

func sClearLine() string {
	if rawOutput() || !outputEnabled() {
		return ""
	}

	return Sprinto(strings.Repeat(" ", GetTerminalWidth()))
}
