package pterm_test

import (
	"github.com/pterm/pterm"
	"testing"
)

func TestLogger_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.Logger{}, "WithCaller", "WithTime")
}
