package pterm

import (
	"io"

	"atomicgo.dev/keyboard/keys"
)

// SetStdinIsTerminal makes interactive printers believe standard input is,
// or is not, a terminal. It returns a function that undoes the change.
func SetStdinIsTerminal(tty bool) (restore func()) {
	old := stdinIsTerminal
	stdinIsTerminal = func() bool { return tty }

	return func() { stdinIsTerminal = old }
}

// SetTerminalKeys makes interactive printers read ks as if they came from
// the terminal. Unlike WithInput, which delivers every rune as its own key,
// this keeps keys as given, so a test can paste text as one key the way
// the terminal reader groups it.
func SetTerminalKeys(ks ...keys.Key) (restore func()) {
	oldTTY, oldListen := stdinIsTerminal, listenKeys
	stdinIsTerminal = func() bool { return true }
	listenKeys = func(onKey func(keys.Key) (bool, error)) error {
		for _, k := range ks {
			stop, err := onKey(k)
			if err != nil || stop {
				return err
			}
		}

		return io.EOF
	}

	return func() { stdinIsTerminal, listenKeys = oldTTY, oldListen }
}

// SetViewHook makes interactive printers pass every view they draw to fn.
func SetViewHook(fn func(view string)) (restore func()) {
	old := viewHook
	viewHook = fn

	return func() { viewHook = old }
}
