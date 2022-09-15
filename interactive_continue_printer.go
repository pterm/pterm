package pterm

import (
	"fmt"
	"os"
	"strings"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var (
	// DefaultInteractiveContinue is the default InteractiveContinue printer.
	// Pressing "y" will return yes, "n" will return no, "a" returns all and "s" returns stop.
	// Pressing enter without typing any letter will return the configured default value (by default set to "yes", the fisrt option).
	DefaultInteractiveContinue = InteractiveContinuePrinter{
		DefaultValue: "Y",
		DefaultText:  "Do you want to continue",
		TextStyle:    &ThemeDefault.PrimaryStyle,
		Options:      map[string]string{"Y": "yes", "n": "no", "a": "allways", "s": "stop"},
		OptionsStyle: &ThemeDefault.SuccessMessageStyle,
		SuffixStyle:  &ThemeDefault.SecondaryStyle,
	}
)

// InteractiveContinuePrinter is a printer for interactive continue prompts.
type InteractiveContinuePrinter struct {
	DefaultValue string
	DefaultText  string
	TextStyle    *Style
	Options      map[string]string
	OptionsStyle *Style
	SuffixStyle  *Style
}

// WithDefaultText sets the default text.
func (p InteractiveContinuePrinter) WithDefaultText(text string) *InteractiveContinuePrinter {
	p.DefaultText = text
	return &p
}

// WithDefaultValueIndex sets the default value, which will be returned when the user presses enter without typing any letter.
func (p InteractiveContinuePrinter) WithDefaultValue(value string) *InteractiveContinuePrinter {
	if _, ok := p.Options[value]; !ok {
		panic("Invalid value: " + value)
	}
	p.DefaultValue = value
	return &p
}

// WithTextStyle sets the text style.
func (p InteractiveContinuePrinter) WithTextStyle(style *Style) *InteractiveContinuePrinter {
	p.TextStyle = style
	return &p
}

// WithOptions sets the options.
func (p InteractiveContinuePrinter) WithOptions(options map[string]string) *InteractiveContinuePrinter {
	p.Options = options
	return &p
}

// WithOptionsStyle sets the continue style.
func (p InteractiveContinuePrinter) WithOptionsStyle(style *Style) *InteractiveContinuePrinter {
	p.OptionsStyle = style
	return &p
}

// WithSuffixStyle sets the suffix style.
func (p InteractiveContinuePrinter) WithSuffixStyle(style *Style) *InteractiveContinuePrinter {
	p.SuffixStyle = style
	return &p
}

// Show shows the continue prompt.
//
// Example:
//
//	 result, _ := pterm.DefaultInteractiveContinue.Show("Do you want to apply the changes?")
//		pterm.Println(result)
func (p InteractiveContinuePrinter) Show(text ...string) (string, error) {
	var result string

	if len(text) == 0 || text[0] == "" {
		text = []string{p.DefaultText}
	}

	p.TextStyle.Print(text[0] + " " + p.getSuffix() + ": ")

	err := keyboard.Listen(func(keyInfo keys.Key) (stop bool, err error) {
		if err != nil {
			return false, fmt.Errorf("failed to get key: %w", err)
		}
		key := keyInfo.Code
		char := keyInfo.String()

		switch key {
		case keys.RuneKey:
			for k := range p.Options {
				c := string([]rune(k)[0])
				if char == c || (k == p.DefaultValue && strings.EqualFold(c, char)) {
					Println()
					result = p.Options[k]
					return true, nil
				}
			}
		case keys.Enter:
			Println()
			result = p.Options[p.DefaultValue]
			return true, nil
		case keys.CtrlC:
			os.Exit(1)
			return true, nil
		}
		return false, nil
	})
	cursor.StartOfLine()
	return result, err
}

// getSuffix returns the continueation prompt suffix
func (p InteractiveContinuePrinter) getSuffix() string {
	if p.Options == nil || len(p.Options) == 0 {
		panic("Handles not initialized")
	}
	var (
		builder strings.Builder
		step    string
	)
	for k, v := range p.Options {
		builder.WriteString(fmt.Sprintf("%s%s (%s)", step, k, v))
		step = " / "
	}

	return p.SuffixStyle.Sprintf("[%s]", builder.String())
}
