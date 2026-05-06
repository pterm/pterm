package pterm_test

import (
	"reflect"
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/pterm/pterm/internal/testhelper"

	"github.com/pterm/pterm"
)

func TestInteractiveSelectPrinter_Show(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOption("b").Show()
	testhelper.AssertEqual(t, "d", result)
}

func TestInteractiveSelectPrinter_Show_MaxHeightSlidingWindow(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e", "f"}).WithDefaultOption("e").Show()
	testhelper.AssertEqual(t, "c", result)
}

func TestInteractiveSelectPrinter_Show_AlternateNavigationKeys(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.CtrlN)
		keyboard.SimulateKeyPress(keys.CtrlN)
		keyboard.SimulateKeyPress(keys.CtrlP)
		keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e"}).WithDefaultOption("b").Show()
	testhelper.AssertEqual(t, "c", result)
}

func TestInteractiveSelectPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithDefaultText("default")
	testhelper.AssertEqual(t, p.DefaultText, "default")
}

func TestInteractiveSelectPrinter_WithDefaultOption(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithDefaultOption("default")
	testhelper.AssertEqual(t, p.DefaultOption, "default")
}

func TestInteractiveSelectPrinter_WithOptions(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"})
	testhelper.AssertEqual(t, p.Options, []string{"a", "b", "c"})
}

func TestInteractiveSelectPrinter_WithMaxHeight(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithMaxHeight(1337)
	testhelper.AssertEqual(t, p.MaxHeight, 1337)
}

func TestInteractiveSelectPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveSelectPrinter{}
	testhelper.AssertNil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveSelect.WithOnInterruptFunc(exitfunc)
	testhelper.AssertEqual(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}

func TestInteractiveSelectPrinter_WithFilter(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithFilter(false)
	testhelper.AssertEqual(t, p.Filter, false)
}
