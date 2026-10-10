//go:build windows

package pterm

import (
	"os"

	"golang.org/x/sys/windows"
)

// codePageUTF8 is the Windows code page number of UTF-8.
const codePageUTF8 = 65001

// legacyConsole reports whether output goes to a classic Windows console
// whose code page is not UTF-8. Windows Terminal and other modern hosts set
// WT_SESSION or similar variables and render Unicode with any code page.
func legacyConsole() bool {
	if os.Getenv("WT_SESSION") != "" || os.Getenv("TERM_PROGRAM") != "" || os.Getenv("ConEmuANSI") == "ON" {
		return false
	}

	var mode uint32
	if windows.GetConsoleMode(windows.Handle(os.Stdout.Fd()), &mode) != nil {
		return false
	}

	cp, err := windows.GetConsoleOutputCP()

	return err == nil && cp != codePageUTF8
}
