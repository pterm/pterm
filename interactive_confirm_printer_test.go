package pterm_test

import (
	"reflect"
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

func TestInteractiveConfirmPrinter_WithConfirmStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveConfirm.WithConfirmStyle(style)
	testza.AssertEqual(t, p.ConfirmStyle, style)
}

func TestInteractiveConfirmPrinter_WithConfirmText(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithConfirmText("confirm")
	testza.AssertEqual(t, p.ConfirmText, "confirm")
}

func TestInteractiveConfirmPrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithDefaultText("default")
	testza.AssertEqual(t, p.DefaultText, "default")
}

func TestInteractiveConfirmPrinter_WithDelimiter(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithDelimiter(">>")
	testza.AssertEqual(t, p.Delimiter, ">>")
}

func TestInteractiveConfirmPrinter_WithRejectStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveConfirm.WithRejectStyle(style)
	testza.AssertEqual(t, p.RejectStyle, style)
}

func TestInteractiveConfirmPrinter_WithRejectText(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithRejectText("reject")
	testza.AssertEqual(t, p.RejectText, "reject")
}

func TestInteractiveConfirmPrinter_CustomAnswers(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithRejectText("reject").WithConfirmText("accept")
	tests := []struct {
		name     string
		key      rune
		expected bool
	}{
		{
			name:     "Accept_upper_case",
			key:      'A',
			expected: true,
		},
		{
			name:     "Accept_lower",
			key:      'a',
			expected: true,
		},
		{
			name:     "Reject_upper_case",
			key:      'R',
			expected: false,
		},
		{
			name:     "Reject_lower_case",
			key:      'r',
			expected: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			go func() {
				keyboard.SimulateKeyPress(tc.key)
			}()
			result, _ := p.Show()
			testza.AssertEqual(t, result, tc.expected)
		})
	}
}

func TestInteractiveConfirmPrinter_WithSuffixStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveConfirm.WithSuffixStyle(style)
	testza.AssertEqual(t, p.SuffixStyle, style)
}

func TestInteractiveConfirmPrinter_WithTextStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveConfirm.WithTextStyle(style)
	testza.AssertEqual(t, p.TextStyle, style)
}

func TestInteractiveConfirmPrinter_WithOnInterruptFunc(t *testing.T) {
	// OnInterrupt function defaults to nil
	pd := pterm.InteractiveConfirmPrinter{}
	testza.AssertNil(t, pd.OnInterruptFunc)

	// Verify OnInterrupt is set
	exitfunc := func() {}
	p := pterm.DefaultInteractiveConfirm.WithOnInterruptFunc(exitfunc)
	testza.AssertEqual(t, reflect.ValueOf(p.OnInterruptFunc).Pointer(), reflect.ValueOf(exitfunc).Pointer())
}
