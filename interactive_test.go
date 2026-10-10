package pterm_test

import (
	"strings"
	"testing"

	"atomicgo.dev/color"
	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm"
)

// Key sequences for WithInput, which reads them like a terminal would
// send them.
const (
	keyUp        = "\x1b[A"
	keyDown      = "\x1b[B"
	keyRight     = "\x1b[C"
	keyLeft      = "\x1b[D"
	keyEnter     = "\r"
	keyBackspace = "\x7f"
	keyCtrlC     = "\x03"
)

// waiting runs show and returns the last view the prompt drew, which is
// what the prompt looked like right before its final key.
func waiting(t *testing.T, show func()) string {
	t.Helper()

	var last string

	restore := pterm.SetViewHook(func(v string) { last = v })
	defer restore()

	show()

	return strings.ReplaceAll(last, "\x00", "")
}

// typed is a key press of text the way the terminal reader delivers typed
// or pasted runes: all of them in one key.
func typed(s string) keys.Key {
	return keys.Key{Code: keys.RuneKey, Runes: []rune(s)}
}

func key(c keys.KeyCode) keys.Key {
	return keys.Key{Code: c}
}

// fromTerminal makes prompts read ks as terminal key presses until the
// test ends.
func fromTerminal(t *testing.T, ks ...keys.Key) {
	t.Helper()
	t.Cleanup(pterm.SetTerminalKeys(ks...))
}

func stripANSI(s string) string {
	return color.Strip(s)
}
