package internal

import "os"

// NewCancelationSignal for keeping track of a cancelation
func NewCancelationSignal(exitFunc func()) (func(), func()) {
	canceled := false

	cancel := func() {
		canceled = true
	}

	if exitFunc == nil {
		exitFunc = func() { os.Exit(1) }
	}

	exit := func() {
		if canceled {
			exitFunc()
		}
	}

	return cancel, exit
}
