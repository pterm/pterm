package internal

import (
	"os"
)

// NewCancelationSignal for keeping track of a cancelation
func NewCancelationSignal() (func(), func()) {
	canceled := false

	cancel := func() {
		canceled = true
	}
	exit := func() {
		if canceled {
			os.Exit(1)
		}
	}

	return cancel, exit
}
