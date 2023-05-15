package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveSelectPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.InteractiveSelectPrinter{})
}

func TestInteractiveSelectPrinter_Show(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOption("b").Show()
	testza.AssertEqual(t, "d", result)
}

func TestInteractiveSelectPrinter_Show_MaxHeightSlidingWindow(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e", "f"}).WithDefaultOption("e").Show()
	testza.AssertEqual(t, "c", result)
}
