package pterm

import (
	"fmt"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var (
	// DefaultInteractiveConfirm is the default InteractiveConfirm printer.
	DefaultInteractiveConfirm = InteractiveConfirmPrinter{
		DefaultValue: false,
		TextStyle:    &ThemeDefault.PrimaryStyle,
		ConfirmText:  "Yes",
		ConfirmStyle: &ThemeDefault.SuccessMessageStyle,
		RejectText:   "No",
		RejectStyle:  &ThemeDefault.ErrorMessageStyle,
		SuffixStyle:  &ThemeDefault.SecondaryStyle,
	}
)

type InteractiveConfirmPrinter struct {
	DefaultValue bool
	TextStyle    *Style
	ConfirmText  string
	ConfirmStyle *Style
	RejectText   string
	RejectStyle  *Style
	SuffixStyle  *Style
}

// WithDefaultValue sets the default value, which will be returned when the user presses enter without typing "y" or "n".
func (p InteractiveConfirmPrinter) WithDefaultValue(value bool) *InteractiveConfirmPrinter {
	p.DefaultValue = value
	return &p
}

// WithTextStyle sets the text style.
func (p InteractiveConfirmPrinter) WithTextStyle(style *Style) *InteractiveConfirmPrinter {
	p.TextStyle = style
	return &p
}

// WithConfirmText sets the confirm text.
func (p InteractiveConfirmPrinter) WithConfirmText(text string) *InteractiveConfirmPrinter {
	p.ConfirmText = text
	return &p
}

// WithConfirmStyle sets the confirm style.
func (p InteractiveConfirmPrinter) WithConfirmStyle(style *Style) *InteractiveConfirmPrinter {
	p.ConfirmStyle = style
	return &p
}

// WithRejectText sets the reject text.
func (p InteractiveConfirmPrinter) WithRejectText(text string) *InteractiveConfirmPrinter {
	p.RejectText = text
	return &p
}

// WithRejectStyle sets the reject style.
func (p InteractiveConfirmPrinter) WithRejectStyle(style *Style) *InteractiveConfirmPrinter {
	p.RejectStyle = style
	return &p
}

// WithSuffixStyle sets the suffix style.
func (p InteractiveConfirmPrinter) WithSuffixStyle(style *Style) *InteractiveConfirmPrinter {
	p.SuffixStyle = style
	return &p
}

// Show shows the confirm prompt.
//
// Example:
//  result, _ := pterm.DefaultInteractiveConfirm.Show("Are you sure?")
//	pterm.Println(result)
func (p InteractiveConfirmPrinter) Show(text ...string) (bool, error) {
	err := keyboard.StartListener()
	if err != nil {
		return false, fmt.Errorf("failed to start keyboard listener: %w", err)
	}

	if text == nil {
		text = []string{"Please confirm"}
	}

	p.TextStyle.Print(text[0] + " " + p.getSuffix() + ": ")

	for {
		keyInfo, err := keyboard.GetKey()
		key := keyInfo.Code
		char := keyInfo.String()
		if err != nil {
			return false, fmt.Errorf("failed to get key: %w", err)
		}

		switch key {
		case keys.RuneKey:
			switch char {
			case "y", "Y":
				p.ConfirmStyle.Print(p.ConfirmText)
				Println()
				return true, keyboard.StopListener()
			case "n", "N":
				p.RejectStyle.Print(p.RejectText)
				Println()
				return false, keyboard.StopListener()
			}
		case keys.Enter:
			if p.DefaultValue {
				p.ConfirmStyle.Print(p.ConfirmText)
			} else {
				p.RejectStyle.Print(p.RejectText)
			}
			Println()
			return p.DefaultValue, keyboard.StopListener()
		case keys.CtrlC:
			os.Exit(1)
			return false, keyboard.StopListener()
		}
	}
}

func (p InteractiveConfirmPrinter) getSuffix() string {
	var y string
	var n string
	if p.DefaultValue {
		y = "Y"
		n = "n"
	} else {
		y = "y"
		n = "N"
	}

	return p.SuffixStyle.Sprintf("[%s/%s]", y, n)
}
