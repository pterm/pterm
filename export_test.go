package pterm

import "time"

// SetLogClock makes log lines use now as their time until the returned
// function restores the real clock.
func SetLogClock(now func() time.Time) (restore func()) {
	logNow = now
	return func() { logNow = time.Now }
}
