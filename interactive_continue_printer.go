package pterm

import (
	"fmt"
	"os"
	"strings"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/pterm/pterm/internal"
)

var (
	// DefaultInteractiveContinue is the default InteractiveContinue printer.
	// Pressing "y" will return yes, "n" will return no, "a" returns all and "s" returns stop.
	// Pressing enter without typing any letter will return the configured default value (by default set to "yes", the fisrt option).
	DefaultInteractiveContinue = InteractiveContinuePrinter{
		DefaultValueIndex: 0,
		DefaultText:       "Do you want to continue",
		TextStyle:         &ThemeDefault.PrimaryStyle,
		Options:           []string{"yes", "no", "all", "cancel"},
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
	ShowShortHandles  bool
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
		Warning.Printf("%v is not a valid set of handles", handles)
		p.setDefaultHandles()
		return &p
	}
	p.Handles = handles
	return &p
}

// WithShowShortHandles will set ShowShortHandles to true
// this makes the printer display the shorthand options instead their shorthand version.
func (p InteractiveContinuePrinter) WithShowShortHandles(b ...bool) *InteractiveContinuePrinter {
	p.ShowShortHandles = internal.WithBoolean(b)
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
//	result, _ := pterm.DefaultInteractiveContinue.Show("Do you want to apply the changes?")
//	pterm.Println(result)
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
			for i, c := range p.Handles {
				if !p.ShowShortHandles {
					c = string([]rune(c)[0])
				}
				if char == c || (i == p.DefaultValueIndex && strings.EqualFold(c, char)) {
					p.OptionsStyle.Print(p.Options[i])
					Println()
					result = p.Options[i]
					return true, nil
				}
			}
		case keys.Enter:
			p.OptionsStyle.Print(p.Options[p.DefaultValueIndex])
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

// getShortHandles returns the short hand answers for the continueation prompt
func (p InteractiveContinuePrinter) getShortHandles() []string {
	var handles []string
	for _, option := range p.Options {
		handles = append(handles, strings.ToLower(string([]rune(option)[0])))
	}
	handles[p.DefaultValueIndex] = strings.ToUpper(handles[p.DefaultValueIndex])

	return handles
}

// setDefaultHandles initialises the handles
func (p *InteractiveContinuePrinter) setDefaultHandles() {
	if p.ShowShortHandles {
		p.Handles = p.getShortHandles()
	}

	if p.Handles == nil || len(p.Handles) == 0 {
		p.Handles = make([]string, len(p.Options))
		copy(p.Handles, p.Options)
		p.Handles[p.DefaultValueIndex] = cases.Title(language.Und, cases.Compact).String(p.Handles[p.DefaultValueIndex])
	}
}

// getSuffix returns the continuation prompt suffix
func (p *InteractiveContinuePrinter) getSuffix() string {
	if p.Handles == nil || len(p.Handles) != len(p.Options) {
		p.setDefaultHandles()
	}

	return p.SuffixStyle.Sprintf("[%s]", strings.Join(p.Handles, "/"))
}
