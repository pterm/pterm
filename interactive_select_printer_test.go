package pterm_test

import (
	"reflect"
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestInteractiveSelectPrinter_Show(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Down)
		_ = keyboard.SimulateKeyPress(keys.Down)
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOption("b").Show()
	assert.Equal(t, "d", result)
}

func TestInteractiveSelectPrinter_Show_MaxHeightSlidingWindow(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e", "f"}).WithDefaultOption("e").Show()
	assert.Equal(t, "c", result)
}

func TestInteractiveSelectPrinter_Show_AlternateNavigationKeys(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.CtrlN)
		_ = keyboard.SimulateKeyPress(keys.CtrlN)
		_ = keyboard.SimulateKeyPress(keys.CtrlP)
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOption("b").Show()
	assert.Equal(t, "c", result)
}

func TestInteractiveSelectPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithDefaultText("default")
	assert.Equal(t, p.DefaultText, "default")
}

func TestInteractiveSelectPrinter_WithDefaultOption(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithDefaultOption("default")
	assert.Equal(t, p.DefaultOption, "default")
}

func TestInteractiveSelectPrinter_WithOptions(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"})
	assert.Equal(t, p.Options, []string{"a", "b", "c"})
}

func TestInteractiveSelectPrinter_WithMaxHeight(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithMaxHeight(1337)
	assert.Equal(t, p.MaxHeight, 1337)
}

func TestInteractiveSelectPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveSelectPrinter{}
	assert.Nil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveSelect.WithOnInterruptFunc(exitfunc)
	assert.Equal(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}

func TestInteractiveSelectPrinter_WithFilter(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithFilter(false)
	assert.Equal(t, p.Filter, false)
}
