package internal

// NewCancelationSignal for keeping track of a cancelation
func NewCancelationSignal(interruptFunc func()) (func(), func()) {
	canceled := false

	cancel := func() {
		canceled = true
	}

	exit := func() {
		if canceled {
			if interruptFunc != nil {
				interruptFunc()
			} else {
				Exit(1)
			}
		}
	}

	return cancel, exit
}
