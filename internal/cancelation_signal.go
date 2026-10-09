// Package internal contains helpers shared by pterm's printers.
package internal

import "sync/atomic"

// NewCancelationSignal for keeping track of a cancelation.
// The flag is atomic because cancel is typically invoked from a keyboard
// listener goroutine while exit runs deferred on the caller's goroutine.
func NewCancelationSignal(interruptFunc func()) (func(), func()) {
	var canceled atomic.Bool

	cancel := func() {
		canceled.Store(true)
	}

	exit := func() {
		if canceled.Load() {
			if interruptFunc != nil {
				interruptFunc()
			} else {
				Exit(1)
			}
		}
	}

	return cancel, exit
}
