package pterm_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveConfirmPrinter_Show_yes(t *testing.T) {
	go func() {
		time.Sleep(500 * time.Millisecond)
		keyboard.SimulateKeyPress('y')
	}()
	result, _ := pterm.DefaultInteractiveConfirm.Show()
	testza.AssertTrue(t, result)
}

func TestInteractiveConfirmPrinter_Show_no(t *testing.T) {
	go func() {
		time.Sleep(500 * time.Millisecond)
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
		time.Sleep(500 * time.Millisecond)
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	p := pterm.DefaultInteractiveConfirm.WithDefaultValue(false)
	result, _ := p.Show()
	testza.AssertFalse(t, result)
}

func TestInteractiveConfirmPrinter_WithDefaultValue_true(t *testing.T) {
	go func() {
		time.Sleep(500 * time.Millisecond)
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

func TestInteractiveConfirmPrinter_WithTimeout(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithTimeout(3 * time.Second)
	testza.AssertEqual(t, p.Timeout, 3*time.Second)
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
				time.Sleep(500 * time.Millisecond)
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

func TestInteractiveConfirmPrinter_WithWriter(t *testing.T) {
	p := pterm.InteractiveConfirmPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}

func TestInteractiveConfirmPrinter_WithTimerStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.InteractiveConfirmPrinter{}
	p2 := p.WithTimeoutTimerStyle(s)

	testza.AssertEqual(t, s, p2.TimeoutTimerStyle)
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

func TestInteractiveConfirmPrinter_ConfirmWithTimeout(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithTimeout(150 * time.Millisecond).WithDefaultValue(true)
	result, _ := p.Show()
	testza.AssertEqual(t, result, p.DefaultValue)
}
