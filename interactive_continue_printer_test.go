package pterm_test

import (
	"testing"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestInteractiveContinuePrinter_Show_yes(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress('y')
	}()
	result, _ := pterm.DefaultInteractiveContinue.Show()
	testza.AssertEqual(t, result, "yes")
	go func() {
		keyboard.SimulateKeyPress('Y')
	}()
	result, _ = pterm.DefaultInteractiveContinue.Show()
	testza.AssertEqual(t, result, "yes")
}

func TestInteractiveContinuePrinter_Show_no(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress('n')
	}()
	result, _ := pterm.DefaultInteractiveContinue.Show()
	testza.AssertEqual(t, result, "no")
}

func TestInteractiveContinuePrinter_WithDefaultValueIndes(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithDefaultValueIndex(1)
	testza.AssertEqual(t, p.DefaultValueIndex, 1)
}

func TestInteractiveContinuePrinter_WithDefaultValue_yes(t *testing.T) {
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	p := pterm.DefaultInteractiveContinue.WithDefaultValue("yes")
	result, _ := p.Show()
	testza.AssertEqual(t, result, "yes")
}

func TestInteractiveContinuePrinter_WithDefaultValue_no(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithDefaultValue("no")
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := p.Show()
	testza.AssertEqual(t, result, "no")
	go func() {
		keyboard.SimulateKeyPress('n')
	}()
	result, _ = p.Show()
	testza.AssertEqual(t, result, "no")
	go func() {
		keyboard.SimulateKeyPress('N')
	}()
	result, _ = p.Show()
	testza.AssertEqual(t, result, "no")
}

func TestInteractiveContinuePrinter_WithShowShortHandles(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithShowShortHandles()
	testza.AssertTrue(t, p.ShowShortHandles)
	go func() {
		keyboard.SimulateKeyPress('n')
	}()
	result, _ := p.Show()
	testza.AssertEqual(t, result, "no")
}

func TestInteractiveContinuePrinter_WithOptionsStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveContinue.WithOptionsStyle(style)
	testza.AssertEqual(t, p.OptionsStyle, style)
}

func TestInteractiveContinuePrinter_WithOptions(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithOptions([]string{"next", "stop", "continue"})
	testza.AssertEqual(t, p.Options, []string{"next", "stop", "continue"})
}

func TestInteractiveContinuePrinter_WithHandles(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithOptions([]string{"yes", "no", "always", "never"}).WithHandles([]string{"y", "n", "a", "N"})
	testza.AssertEqual(t, p.Handles, []string{"y", "n", "a", "N"})
	tests := []struct {
		name     string
		key      rune
		expected string
	}{
		{
			name:     "Yes",
			key:      'y',
			expected: "yes",
		},
		{
			name:     "No",
			key:      'n',
			expected: "no",
		},
		{
			name:     "Always",
			key:      'a',
			expected: "always",
		},
		{
			name:     "Never",
			key:      'N',
			expected: "never",
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
	p.DefaultValueIndex = 1
	go func() {
		keyboard.SimulateKeyPress(keys.Enter)
	}()
	result, _ := p.Show()
	testza.AssertEqual(t, result, "no")
}

func TestInteractiveContinuePrinter_WithDefaultText(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithDefaultText("default")
	testza.AssertEqual(t, p.DefaultText, "default")
}

func TestInteractiveContinuePrinter_WithDelimiter(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithDelimiter(">>")
	testza.AssertEqual(t, p.Delimiter, ">>")
}

func TestInteractiveContinuePrinter_CustomAnswers(t *testing.T) {
	p := pterm.DefaultInteractiveContinue.WithOptions([]string{"next", "stop", "continue"})
	tests := []struct {
		name     string
		key      rune
		expected string
	}{
		{
			name:     "Next",
			key:      'n',
			expected: "next",
		},
		{
			name:     "Stop",
			key:      's',
			expected: "stop",
		},
		{
			name:     "Continue",
			key:      'c',
			expected: "continue",
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

func TestInteractiveContinuePrinter_WithSuffixStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveContinue.WithSuffixStyle(style)
	testza.AssertEqual(t, p.SuffixStyle, style)
}

func TestInteractiveContinuePrinter_WithTextStyle(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	p := pterm.DefaultInteractiveContinue.WithTextStyle(style)
	testza.AssertEqual(t, p.TextStyle, style)
}
