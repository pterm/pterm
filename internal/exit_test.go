package internal_test

import (
	"os"
	"testing"

	"github.com/pterm/pterm/internal"
)

func TestExit(t *testing.T) {
	var lastExitCode int
	internal.DefaultExitFunc = func(code int) {
		lastExitCode = code
	}

	defer func() { internal.DefaultExitFunc = os.Exit }()

	internal.Exit(1)

	if lastExitCode != 1 {
		t.Errorf("Expected exit code 1, got %d", lastExitCode)
	}
}
