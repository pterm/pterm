package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/gookit/color"
)

// SetDefaultOutput sets the default output of pterm.
func SetDefaultOutput(w io.Writer) {
	color.SetOutput(w)
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...interface{}) string {
	return color.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string {
	return color.Sprintf(format, a...)
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintfln(format string, a ...interface{}) string {
	return color.Sprintf(format, a...) + "\n"
}

// Sprintln returns what Println would print to the terminal.
func Sprintln(a ...interface{}) string {
	str := fmt.Sprintln(a...)
	return Sprint(str)
}

// Sprinto returns what Printo would print.
func Sprinto(a ...interface{}) string {
	return "\r" + Sprint(a...)
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func Print(a ...interface{}) {
	if !Output {
		return
	}

	var ret string
	var printed bool

	for _, bar := range ActiveProgressBarPrinters {
		if bar.IsActive {
			ret += sClearLine()
			ret += Sprinto(a...)
			printed = true
		}
	}

	for _, spinner := range activeSpinnerPrinters {
		if spinner.IsActive {
			ret += sClearLine()
			ret += Sprinto(a...)
			printed = true
		}
	}

	if !printed {
		ret = color.Sprint(Sprint(a...))
	}

	color.Print(Sprint(ret))
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) {
	Print(Sprintln(a...))
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) {
	Print(Sprintf(format, a...))
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Printfln(format string, a ...interface{}) {
	Print(Sprintfln(format, a...))
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func PrintOnError(a ...interface{}) {
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
func PrintOnErrorf(format string, a ...interface{}) {
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
func Fprint(writer io.Writer, a ...interface{}) {
	if !Output {
		return
	}

	color.Fprint(writer, Sprint(a...))
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Fprintln(writer io.Writer, a ...interface{}) {
	Fprint(writer, Sprint(a...)+"\n")
}

// Printo overrides the current line in a terminal.
// If the current line is empty, the text will be printed like with pterm.Print.
// Example:
//   pterm.Printo("Hello, World")
//   time.Sleep(time.Second)
//   pterm.Printo("Hello, Earth!")
func Printo(a ...interface{}) {
	if !Output {
		return
	}

	color.Print("\r" + Sprint(a...))
}

// Fprinto prints Printo to a custom writer.
func Fprinto(w io.Writer, a ...interface{}) {
	Fprint(w, "\r", Sprint(a...))
}

// RemoveColorFromString removes color codes from a string.
func RemoveColorFromString(a ...interface{}) string {
	return color.ClearCode(Sprint(a...))
}

func clearLine() {
	Printo(strings.Repeat(" ", GetTerminalWidth()))
}

func sClearLine() string {
	return Sprinto(strings.Repeat(" ", GetTerminalWidth()))
}
