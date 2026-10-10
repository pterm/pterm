package pterm

import "time"

// SetStdinIsTerminal makes interactive printers believe standard input is,
// or is not, a terminal. It returns a function that undoes the change.
func SetStdinIsTerminal(tty bool) (restore func()) {
	old := stdinIsTerminal
	stdinIsTerminal = func() bool { return tty }

	return func() { stdinIsTerminal = old }
}

// SetLogClock makes log lines use now as their time until the returned
// function restores the real clock.
func SetLogClock(now func() time.Time) (restore func()) {
	logNow = now
	return func() { logNow = time.Now }
}
