package pterm_test

import (
	"reflect"
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

func TestInteractiveSelectPrinter_Show_MaxHeightSlidingWindow(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c", "d", "e", "f"}).WithDefaultOption("e").Show()
	testza.AssertEqual(t, "c", result)
}

func TestInteractiveSelectPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithDefaultText("default")
	testza.AssertEqual(t, p.DefaultText, "default")
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

func TestInteractiveSelectPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveSelectPrinter{}
	testza.AssertNil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveSelect.WithOnInterruptFunc(exitfunc)
	testza.AssertEqual(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}

func TestInteractiveSelectPrinter_WithFilter(t *testing.T) {
	p := pterm.DefaultInteractiveSelect.WithFilter(false)
	testza.AssertEqual(t, p.Filter, false)
}

func TestInteractiveSelectPrinter_GetSelectedOption(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Down)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	is := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"})
	is.Show()
	testza.AssertEqual(t, is.GetSelectedOption(), 2)
}

func TestInteractiveSelectPrinter_WithDefaultSelectedOption(t *testing.T) {
	// Check selectedOption value
	p := pterm.DefaultInteractiveSelect.WithDefaultSelectedOption(2)
	testza.AssertEqual(t, p.GetSelectedOption(), 2)

	// Check behavior
	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}).WithDefaultSelectedOption(2).Show()
	testza.AssertEqual(t, result, "a")

	// Check default values both with index and string
	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ = pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}).WithDefaultSelectedOption(2).WithDefaultOption("b").Show()
	testza.AssertEqual(t, result, "a")

	go func() {
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Up)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ = pterm.DefaultInteractiveSelect.WithOptions([]string{"a", "b", "c"}).WithDefaultOption("b").WithDefaultSelectedOption(2).Show()
	testza.AssertEqual(t, result, "a")
}
