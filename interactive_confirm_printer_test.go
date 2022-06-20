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

func TestInteractiveConfirmPrinter_WithRejectStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveConfirm.WithRejectStyle(style)
	testza.AssertEqual(t, p.RejectStyle, style)
}

func TestInteractiveConfirmPrinter_WithRejectText(t *testing.T) {
	p := pterm.DefaultInteractiveConfirm.WithRejectText("reject")
	testza.AssertEqual(t, p.RejectText, "reject")
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
