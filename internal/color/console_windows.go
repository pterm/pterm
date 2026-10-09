//go:build windows

package color

import (
	"os"
	"sync"

	"golang.org/x/sys/windows"
)

var (
	consoleOnce    sync.Once
	consoleVerdict Level
	consoleDecided bool
)

// consoleSupport reports the color support of the attached Windows console.
// The first call tries to enable virtual terminal processing on stdout and
// stderr, which makes classic consoles like cmd.exe and powershell interpret
// ANSI escape sequences.
//
// Verdicts:
//   - (LevelTrueColor, true): a console is attached and interprets ANSI
//     sequences. Every supported Windows version renders 24-bit color once
//     virtual terminal processing is on.
//   - (LevelNone, true): a console is attached but cannot interpret ANSI
//     sequences (legacy consoles, e.g. on Windows 7/8).
//   - (LevelNone, false): no console is attached (output is redirected), so
//     the environment has to decide.
func consoleSupport() (Level, bool) {
	consoleOnce.Do(func() {
		sawConsole := false
		ansiOK := true

		for _, f := range []*os.File{os.Stdout, os.Stderr} {
			handle := windows.Handle(f.Fd())

			var mode uint32
			if windows.GetConsoleMode(handle, &mode) != nil {
				// Not a console (e.g. redirected to a file or pipe); escape
				// sequences pass through unchanged, so there is nothing to
				// enable and nothing to learn.
				continue
			}

			sawConsole = true

			if mode&windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING != 0 {
				continue
			}

			if windows.SetConsoleMode(handle, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING) != nil {
				ansiOK = false
			}
		}

		if !sawConsole {
			return
		}

		consoleDecided = true
		if ansiOK {
			consoleVerdict = LevelTrueColor
		}
	})

	return consoleVerdict, consoleDecided
}
