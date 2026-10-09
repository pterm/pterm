//go:build windows

package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleSupportVerdict(t *testing.T) {
	console, known := consoleSupport()

	if known {
		// A console is attached to the test run. Every Windows version that
		// can run this test suite supports virtual terminal processing, so
		// enabling it must have succeeded.
		assert.Equal(t, LevelTrueColor, console)

		return
	}

	// Output is redirected (the usual case under `go test`): no verdict.
	assert.Equal(t, LevelNone, console)
}
