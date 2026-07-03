package pterm_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal"
)

func TestInteractiveTextInputPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDefaultText("default")
	assert.Equal(t, p.DefaultText, "default")
}

func TestInteractiveTextInputPrinter_WithDefaultValue(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDefaultValue("default")
	assert.Equal(t, p.DefaultValue, "default")
}

func TestInteractiveTextInputPrinter_WithDelimiter(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithDelimiter(">>")
	assert.Equal(t, p.Delimiter, ">>")
}

func TestInteractiveTextInputPrinter_WithMultiLine_true(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithMultiLine()
	assert.True(t, p.MultiLine)
}

func TestInteractiveTextInputPrinter_WithMultiLine_false(t *testing.T) {
	p := pterm.DefaultInteractiveTextInput.WithMultiLine(false)
	assert.False(t, p.MultiLine)
}

func TestInteractiveTextInputPrinter_WithTextStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveTextInput.WithTextStyle(style)
	assert.Equal(t, p.TextStyle, style)
}

func TestInteractiveTextInputPrinter_WithMask(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress('a')
		_ = keyboard.SimulateKeyPress('b')
		_ = keyboard.SimulateKeyPress('c')
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show()
	assert.Equal(t, "abc", result)
}

func TestInteractiveTextInputPrinter_WithCancel(t *testing.T) {
	exitCalled := false
	internal.DefaultExitFunc = func(_ int) {
		exitCalled = true
	}

	defer func() { internal.DefaultExitFunc = os.Exit }()

	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress(keys.CtrlC)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show()
	assert.Equal(t, "", result)

	if !exitCalled {
		t.Errorf("Expected exit to be called on Ctrl+C")
	}
}

func TestInteractiveTextInputPrinter_OnEnter(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("default").Show()
	assert.Equal(t, "default", result)
}

func TestInteractiveTextInputPrinter_Editable(t *testing.T) {
	go func() {
		// change `default` to `deffaultt` by simulating cursor moves `left`, `right` and inserting
		// keys on current cursor positions.
		_ = keyboard.SimulateKeyPress(keys.Left)
		_ = keyboard.SimulateKeyPress(keys.Left)
		_ = keyboard.SimulateKeyPress(keys.Left)
		_ = keyboard.SimulateKeyPress(keys.Left)
		_ = keyboard.SimulateKeyPress(keys.Key{Code: keys.RuneKey, Runes: []rune{'f'}})
		_ = keyboard.SimulateKeyPress(keys.Right)
		_ = keyboard.SimulateKeyPress(keys.Right)
		_ = keyboard.SimulateKeyPress(keys.Right)
		_ = keyboard.SimulateKeyPress(keys.Key{Code: keys.RuneKey, Runes: []rune{'t'}})
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue("default").Show()
	assert.Equal(t, "deffaultt", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnTab(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("default").Show()
	assert.Equal(t, "default", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnUp(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress(keys.Backspace)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress("second line")
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress("first line")
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	assert.Equal(t, "first line\nsecond line", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnDown(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress("a")
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress("b")
		_ = keyboard.SimulateKeyPress(keys.Down)
		_ = keyboard.SimulateKeyPress("c")
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	assert.Equal(t, "a\nb\nc", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnLeft(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress(keys.Backspace)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress("a")
		_ = keyboard.SimulateKeyPress(keys.Left)
		_ = keyboard.SimulateKeyPress(keys.Left)
		_ = keyboard.SimulateKeyPress("b")
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	assert.Equal(t, "b\na", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnRight(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress('a')
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress(keys.Right)
		_ = keyboard.SimulateKeyPress(keys.Right)
		_ = keyboard.SimulateKeyPress("b")
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).Show()
	assert.Equal(t, "a\nb", result)
}

func TestInteractiveTextInputPrinter_OnBackspace(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Backspace)
		_ = keyboard.SimulateKeyPress(keys.Enter)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithDefaultValue("a").Show()
	assert.Equal(t, "", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnDelete(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Backspace)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress('a')
		_ = keyboard.SimulateKeyPress(keys.Up)
		_ = keyboard.SimulateKeyPress(keys.Delete)
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("a").Show()
	assert.Equal(t, "a", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnBackspace(t *testing.T) {
	go func() {
		_ = keyboard.SimulateKeyPress(keys.Backspace)
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress(keys.Backspace)
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		WithDefaultValue("a").Show()
	assert.Equal(t, "", result)
}

func TestInteractiveTextInputPrinter_WithMultiLineOnLeftRight(t *testing.T) {
	go func() {
		time.Sleep(1 * time.Millisecond)

		_ = keyboard.SimulateKeyPress("a")
		_ = keyboard.SimulateKeyPress(keys.Enter)
		_ = keyboard.SimulateKeyPress("b")
		_ = keyboard.SimulateKeyPress(keys.Tab)
	}()

	result, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine(true).
		Show("Enter")
	assert.Equal(t, "a\nb", result)
}

func TestInteractiveTextInputPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveTextInputPrinter{}
	assert.Nil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveTextInput.WithOnInterruptFunc(exitfunc)
	assert.Equal(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}
