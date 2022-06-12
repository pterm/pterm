package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestInteractiveSelectPrinter_Show(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOption("b").Show()
	testza.AssertEqual(t, "d", result)
}

func TestInteractiveSelectPrinter_WithDefaultOption(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithDefaultOption("default")
	testza.AssertEqual(t, p.DefaultOption, "default")
}

func TestInteractiveSelectPrinter_WithOptions(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"})
	testza.AssertEqual(t, p.Options, []string{"a", "b", "c"})
}

func TestInteractiveSelectPrinter_WithMaxHeight(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithMaxHeight(1337)
	testza.AssertEqual(t, p.MaxHeight, 1337)
}
