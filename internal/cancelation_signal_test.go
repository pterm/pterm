package internal

import (
	"os"
	"testing"
)

func TestNewCancelationSignal(t *testing.T) {
	// Mock DefaultExitFunc
	exitCalled := false
	exitCode := 0
	DefaultExitFunc = func(code int) {
		exitCalled = true
		exitCode = code
	}
	defer func() { DefaultExitFunc = os.Exit }() // Reset after tests

	// Scenario 1: Testing cancel function
	cancel, exit := NewCancelationSignal(nil)
	if exitCalled {
		t.Error("Exit function should not be called immediately after NewCancelationSignal")
	}

	// Scenario 2: Testing exit function without cancel
	exit()
	if exitCalled {
		t.Error("Exit function should not be called when cancel is not set")
	}

	// Scenario 3: Testing cancel then exit with no interruptFunc
	cancel()
	exit()
	if !exitCalled || exitCode != 1 {
		t.Errorf("Expected Exit(1) to be called, exitCalled: %v, exitCode: %d", exitCalled, exitCode)
	}

	// Reset for next scenario
	exitCalled = false
	exitCode = 0

	// Scenario 4: Testing cancel then exit with interruptFunc
	interruptCalled := false
	cancel, exit = NewCancelationSignal(func() {
		interruptCalled = true
	})
	cancel()
	exit()
	if interruptCalled == false {
		t.Error("Expected interruptFunc to be called")
	}
	if exitCalled {
		t.Error("Exit should not be called when interruptFunc is provided")
	}
}
