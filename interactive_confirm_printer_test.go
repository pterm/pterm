package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestInteractiveConfirmPrinter_Show_yes(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress('y')
	}()
	result, _ := pterm.DefaultInteractiveConfirm.Show()
	testza.AssertTrue(t, result)
}

func TestInteractiveConfirmPrinter_Show_no(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress('n')
	}()
	result, _ := pterm.DefaultInteractiveConfirm.Show()
	testza.AssertFalse(t, result)
}

func TestInteractiveConfirmPrinter_WithDefaultValue(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithDefaultValue(true)
	testza.AssertTrue(t, p.DefaultValue)
}

func TestInteractiveConfirmPrinter_WithDefaultValue_false(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	p := pterm.DefaultInteractiveConfirm.WithDefaultValue(false)
	result, _ := p.Show()
	testza.AssertFalse(t, result)
}

func TestInteractiveConfirmPrinter_WithDefaultValue_true(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	p := pterm.DefaultInteractiveConfirm.WithDefaultValue(true)
	result, _ := p.Show()
	testza.AssertTrue(t, result)
}
