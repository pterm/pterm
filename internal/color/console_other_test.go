//go:build !windows

package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsoleSupportGivesNoVerdict(t *testing.T) {
	console, known := consoleSupport()

	assert.False(t, known)
	assert.Equal(t, LevelNone, console)
}
