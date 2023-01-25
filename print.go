package pterm

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/gookit/color"
)

// Need to use this because "github.com/gookit/color" is NOT a thread-safe library for Print & Sprintf functions.
// Used to protect against some unsafe actions in Fprint as well
var pLock sync.RWMutex

// SetDefaultOutput sets the default output of pterm.
func SetDefaultOutput(w io.Writer) {
	pLock.Lock()
	defer pLock.Unlock()
	color.SetOutput(w)
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func Sprint(a ...interface{}) string {
	pLock.Lock()
	defer pLock.Unlock()
	return color.Sprint(a...)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string {
	pLock.Lock()
	defer pLock.Unlock()
	return color.Sprintf(format, a...)
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func Sprintfln(format string, a ...interface{}) string {
	pLock.Lock()
	defer pLock.Unlock()
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
	Fprint(nil, a...)
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
	pLock.Lock()
	defer pLock.Unlock()
	if !Output.Load() {
		return
	}

	var ret string
	var printed bool

	activeProgressBarPrinters.lock.Lock()
	for _, bar := range activeProgressBarPrinters.printers {
		if bar.IsActive && bar.Writer == writer {
			ret += sClearLine()
			ret += "\r" + color.Sprint(a...)
			printed = true
		}
	}
	activeProgressBarPrinters.lock.Unlock()

	activeSpinnerPrinters.lock.Lock()
	for _, spinner := range activeSpinnerPrinters.printers {
		if spinner.atomicIsActive.Load() && spinner.Writer == writer {
			ret += sClearLine()
			ret += "\r" + color.Sprint(a...)
			printed = true
		}
	}
	activeSpinnerPrinters.lock.Unlock()

	if !printed {
		ret = color.Sprint(a...)
	}

	if writer != nil {
		color.Fprint(writer, color.Sprint(ret))
	} else {
		color.Print(color.Sprint(ret))
	}
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
//
//	pterm.Printo("Hello, World")
//	time.Sleep(time.Second)
//	pterm.Printo("Hello, Earth!")
func Printo(a ...interface{}) {
	pLock.Lock()
	defer pLock.Unlock()
	if !Output.Load() {
		return
	}

	color.Print("\r" + color.Sprint(a...))
}

// Fprinto prints Printo to a custom writer.
func Fprinto(w io.Writer, a ...interface{}) {
	pLock.Lock()
	defer pLock.Unlock()
	if !Output.Load() {
		return
	}
	if w != nil {
		color.Fprint(w, "\r", color.Sprint(a...))
	} else {
		color.Print("\r", color.Sprint(a...))
	}
}

// RemoveColorFromString removes color codes from a string.
func RemoveColorFromString(a ...interface{}) string {
	return color.ClearCode(Sprint(a...))
}

func fClearLine(writer io.Writer) {
	Fprinto(writer, strings.Repeat(" ", GetTerminalWidth()))
}

func sClearLine() string {
	return "\r" + color.Sprint(strings.Repeat(" ", GetTerminalWidth()))
}
