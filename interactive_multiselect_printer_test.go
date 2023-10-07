package pterm_test

import (
	"reflect"
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveMultiselectPrinter_Show(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOptions([]string{"b"}).Show()
	testza.AssertEqual(t, []string{"b", "c"}, result)
}

func TestInteractiveMultiselectPrinter_Show_MaxHeightSlidingWindow(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e", "f"}).WithDefaultOptions([]string{"b"}).Show()
	testza.AssertEqual(t, []string{"b", "e"}, result)
}

func TestInteractiveMultiselectPrinter_Show_AlternateNavigationKeys(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.CtrlN)
		keyboard.SimulateKeyPress(keys.CtrlN)
		keyboard.SimulateKeyPress(keys.CtrlN)
		keyboard.SimulateKeyPress(keys.CtrlP)
		keyboard.SimulateKeyPress(keys.Enter)
		keyboard.SimulateKeyPress(keys.Tab)
	}()
	result, _ := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOptions([]string{"b"}).Show()
	testza.AssertEqual(t, []string{"b", "c"}, result)
}

func TestInteractiveMultiselectPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithDefaultText("default")
	testza.AssertEqual(t, p.DefaultText, "default")
}

func TestInteractiveMultiselectPrinter_WithDefaultOption(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithDefaultOptions([]string{"default"})
	testza.AssertEqual(t, p.DefaultOptions, []string{"default"})
}

func TestInteractiveMultiselectPrinter_WithOptions(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c"})
	testza.AssertEqual(t, p.Options, []string{"a", "b", "c"})
}

func TestInteractiveMultiselectPrinter_WithMaxHeight(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithMaxHeight(1337)
	testza.AssertEqual(t, p.MaxHeight, 1337)
}

func TestInteractiveMultiselectPrinter_WithKeySelect(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithKeySelect(keys.Left).WithOptions([]string{"a", "b", "c"})
	testza.AssertEqual(t, p.KeySelect, keys.Left)
}

func TestInteractiveMultiselectPrinter_WithKeyConfirm(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithKeyConfirm(keys.Left).WithOptions([]string{"a", "b", "c"})
	testza.AssertEqual(t, p.KeyConfirm, keys.Left)
}

func TestInteractiveMultiselectPrinter_WithCheckmark(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithCheckmark(&pterm.Checkmark{Checked: "+", Unchecked: "-"}).WithOptions([]string{"a", "b", "c"})
	testza.AssertEqual(t, p.Checkmark, &pterm.Checkmark{Checked: "+", Unchecked: "-"})
}

func TestInteractiveMultiselectPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveMultiselectPrinter{}
	testza.AssertNil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveMultiselect.WithOnInterruptFunc(exitfunc)
	testza.AssertEqual(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}
