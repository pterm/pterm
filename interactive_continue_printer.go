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
		DefaultValueIndex: 0,
		DefaultText:       "Do you want to continue",
		TextStyle:         &ThemeDefault.PrimaryStyle,
		Options:           []string{"yes", "no", "all", "stop"},
		OptionsStyle:      &ThemeDefault.SuccessMessageStyle,
		SuffixStyle:       &ThemeDefault.SecondaryStyle,
	}
)

// InteractiveContinuePrinter is a printer for interactive continue prompts.
type InteractiveContinuePrinter struct {
	DefaultValueIndex int
	DefaultText       string
	TextStyle         *Style
	Options           []string
	OptionsStyle      *Style
	Handles           []string
	ShowFullHandles   bool
	SuffixStyle       *Style
}

// WithDefaultText sets the default text.
func (p InteractiveContinuePrinter) WithDefaultText(text string) *InteractiveContinuePrinter {
	p.DefaultText = text
	return &p
}

// WithDefaultValueIndex sets the default value, which will be returned when the user presses enter without typing any letter.
func (p InteractiveContinuePrinter) WithDefaultValueIndex(value int) *InteractiveContinuePrinter {
	if value >= len(p.Options) {
		panic("Index out of range")
	}
	p.DefaultValueIndex = value
	return &p
}

// WithDefaultValue sets the default value, which will be returned when the user presses enter without typing any letter.
func (p InteractiveContinuePrinter) WithDefaultValue(value string) *InteractiveContinuePrinter {
	for i, o := range p.Options {
		if o == value {
			p.DefaultValueIndex = i
			break
		}
	}
	return &p
}

// WithTextStyle sets the text style.
func (p InteractiveContinuePrinter) WithTextStyle(style *Style) *InteractiveContinuePrinter {
	p.TextStyle = style
	return &p
}

// WithOptions sets the options.
func (p InteractiveContinuePrinter) WithOptions(options []string) *InteractiveContinuePrinter {
	p.Options = options
	return &p
}

// WithHandles allows you to customize the short handles for the answers.
func (p InteractiveContinuePrinter) WithHandles(handles []string) *InteractiveContinuePrinter {
	if len(handles) != len(p.Options) {
		panic("Invalid number of handles")
	}
	p.Handles = handles
	return &p
}

// WithFullHandles will set ShowFullHandles to true
// this makes the printer display the full options instead their shorthand version.
func (p InteractiveContinuePrinter) WithFullHandles() *InteractiveContinuePrinter {
	p.ShowFullHandles = true
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
//  result, _ := pterm.DefaultInteractiveContinue.Show("Do you want to apply the changes?")
//	pterm.Println(result)
func (p InteractiveContinuePrinter) Show(text ...string) (string, error) {
	var result string

	if len(text) == 0 || text[0] == "" {
		text = []string{p.DefaultText}
	}

	if p.ShowFullHandles {
		p.Handles = p.Options
	}

	if p.Handles == nil || len(p.Handles) == 0 {
		p.Handles = p.getDefaultHandles()
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
			for i, c := range p.Handles {
				if p.ShowFullHandles {
					c = string([]rune(c)[0])
				}
				if char == c || (i == p.DefaultValueIndex && strings.EqualFold(c, char)) {
					Println()
					result = p.Options[i]
					return true, nil
				}
			}
		case keys.Enter:
			Println()
			result = p.Options[p.DefaultValueIndex]
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

// getDefaultHandles returns the short hand answers for the continueation prompt
func (p InteractiveContinuePrinter) getDefaultHandles() []string {
	handles := []string{}
	for _, option := range p.Options {
		handles = append(handles, strings.ToLower(string([]rune(option)[0])))
	}
	handles[p.DefaultValueIndex] = strings.ToUpper(handles[p.DefaultValueIndex])

	return handles
}

// getSuffix returns the continueation prompt suffix
func (p InteractiveContinuePrinter) getSuffix() string {
	if p.Handles == nil || len(p.Handles) != len(p.Options) {
		panic("Handles not initialized")
	}

	return p.SuffixStyle.Sprintf("[%s]", strings.Join(p.Handles, "/"))
}
