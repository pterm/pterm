package pterm_test

import (
	"reflect"
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestInteractiveMultiselectPrinter_Show(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Down)
		_ = keyboard.SimulateKeyPress(keys.Down)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOptions([]string{"b"}).Show()
	assert.Equal(t, []string{"b", "c"}, result)
}

func TestInteractiveMultiselectPrinter_Show_MaxHeightSlidingWindow(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e", "f"}).WithDefaultOptions([]string{"b"}).Show()
	assert.Equal(t, []string{"b", "e"}, result)
}

func TestInteractiveMultiselectPrinter_Show_AlternateNavigationKeys(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.CtrlN)
		_ = keyboard.SimulateKeyPress(keys.CtrlN)
		_ = keyboard.SimulateKeyPress(keys.CtrlN)
		_ = keyboard.SimulateKeyPress(keys.CtrlP)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOptions([]string{"b"}).Show()
	assert.Equal(t, []string{"b", "c"}, result)
}

func TestInteractiveMultiselectPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithDefaultText("default")
	assert.Equal(t, p.DefaultText, "default")
}

func TestInteractiveMultiselectPrinter_WithDefaultOption(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithDefaultOptions([]string{"default"})
	assert.Equal(t, p.DefaultOptions, []string{"default"})
}

func TestInteractiveMultiselectPrinter_WithOptions(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithOptions([]string{"a", "b", "c"})
	assert.Equal(t, p.Options, []string{"a", "b", "c"})
}

func TestInteractiveMultiselectPrinter_WithMaxHeight(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithMaxHeight(1337)
	assert.Equal(t, p.MaxHeight, 1337)
}

func TestInteractiveMultiselectPrinter_WithKeySelect(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithKeySelect(keys.Left).WithOptions([]string{"a", "b", "c"})
	assert.Equal(t, p.KeySelect, keys.Left)
}

func TestInteractiveMultiselectPrinter_WithKeyConfirm(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithKeyConfirm(keys.Left).WithOptions([]string{"a", "b", "c"})
	assert.Equal(t, p.KeyConfirm, keys.Left)
}

func TestInteractiveMultiselectPrinter_WithCheckmark(t *testing.T) {
	p := pterm.DefaultInteractiveMultiselect.WithCheckmark(&pterm.Checkmark{Checked: "+", Unchecked: "-"}).WithOptions([]string{"a", "b", "c"})
	assert.Equal(t, p.Checkmark, &pterm.Checkmark{Checked: "+", Unchecked: "-"})
}

func TestInteractiveMultiselectPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveMultiselectPrinter{}
	assert.Nil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveMultiselect.WithOnInterruptFunc(exitfunc)
	assert.Equal(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}
