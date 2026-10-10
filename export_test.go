package pterm

// SetStdinIsTerminal makes interactive printers believe standard input is,
// or is not, a terminal. It returns a function that undoes the change.
func SetStdinIsTerminal(tty bool) (restore func()) {
	old := stdinIsTerminal
	stdinIsTerminal = func() bool { return tty }

	return func() { stdinIsTerminal = old }
}
