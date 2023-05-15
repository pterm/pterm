package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestInteractiveTextInputPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.InteractiveTextInputPrinter{}, "WithCaller")
}
