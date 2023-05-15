package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveConfirmPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.InteractiveConfirmPrinter{})
}

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
