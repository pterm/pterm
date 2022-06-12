package pterm

import (
	"fmt"
	"os"
	"sort"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	// DefaultInteractiveSelect is the default InteractiveSelect printer.
	DefaultInteractiveSelect = InteractiveSelectPrinter{
		TextStyle:     &ThemeDefault.PrimaryStyle,
		Options:       []string{},
		OptionStyle:   &ThemeDefault.DefaultText,
		DefaultOption: "",
		MaxHeight:     5,
		Selector:      ">",
		SelectorStyle: &ThemeDefault.SecondaryStyle,
	}
)

// InteractiveSelectPrinter is a printer for interactive select menus.
type InteractiveSelectPrinter struct {
	TextStyle     *Style
	Options       []string
	OptionStyle   *Style
	DefaultOption string
	MaxHeight     int
	Selector      string
	SelectorStyle *Style

	selectedOption     int
	result             string
	text               string
	fuzzySearchString  string
	fuzzySearchMatches []string
}

// WithOptions sets the options.
func (p InteractiveSelectPrinter) WithOptions(options []string) *InteractiveSelectPrinter {
	p.Options = options
	return &p
}

// WithDefaultOption sets the default options.
func (p InteractiveSelectPrinter) WithDefaultOption(option string) *InteractiveSelectPrinter {
	p.DefaultOption = option
	return &p
}

// WithMaxHeight sets the maximum height of the select menu.
func (p InteractiveSelectPrinter) WithMaxHeight(maxHeight int) *InteractiveSelectPrinter {
	p.MaxHeight = maxHeight
	return &p
}

func (p *InteractiveSelectPrinter) Show(text ...string) (string, error) {
	if len(text) == 0 || text[0] == "" {
		text = []string{"Please select an option"}
	}

	p.text = p.TextStyle.Sprint(text[0])
	p.fuzzySearchMatches = p.Options

	if p.MaxHeight == 0 {
		p.MaxHeight = DefaultInteractiveSelect.MaxHeight
	}

	if len(p.Options) == 0 {
		return "", fmt.Errorf("no options provided")
	}

	// Get index of default option
	if p.DefaultOption != "" {
		for i, option := range p.Options {
			if option == p.DefaultOption {
				p.selectedOption = i
				break
			}
		}
	}

	area, err := DefaultArea.Start(p.renderSelectMenu())
	defer area.Stop()
	if err != nil {
		return "", fmt.Errorf("could not start area: %w", err)
	}

	cursor.Hide()
	defer cursor.Show()
	err = keyboard.Listen(func(keyInfo keys.Key) (stop bool, err error) {
		key := keyInfo.Code

		switch key {
		case keys.RuneKey:
			// Fuzzy search for options

			// append to fuzzy search string
			p.fuzzySearchString += keyInfo.String()

			p.selectedOption = 0
			area.Update(p.renderSelectMenu())
		case keys.Space:
			p.fuzzySearchString += " "
			p.selectedOption = 0
			area.Update(p.renderSelectMenu())
		case keys.Backspace:
			// Remove last character from fuzzy search string
			if len(p.fuzzySearchString) > 0 {
				p.fuzzySearchString = p.fuzzySearchString[:len(p.fuzzySearchString)-1]
			}

			if p.fuzzySearchString == "" {
				p.fuzzySearchMatches = p.Options
			}

			area.Update(p.renderSelectMenu())
		case keys.Up:
			if p.selectedOption > 0 {
				p.selectedOption--
			} else {
				p.selectedOption = len(p.fuzzySearchMatches) - 1
			}
			area.Update(p.renderSelectMenu())
		case keys.Down:
			if p.selectedOption < len(p.fuzzySearchMatches)-1 {
				p.selectedOption++
			} else {
				p.selectedOption = 0
			}
			area.Update(p.renderSelectMenu())
		case keys.CtrlC:
			os.Exit(1)
		case keys.Enter:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			p.result = p.fuzzySearchMatches[p.selectedOption]
			area.Update(p.renderFinishedMenu())
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("failed to start keyboard listener: %w", err)
	}

	return p.result, nil
}

func (p InteractiveSelectPrinter) renderSelectMenu() string {
	var content string
	content += Sprintf("%s %s: %s\n", p.text, ThemeDefault.SecondaryStyle.Sprint("[type to search]"), p.fuzzySearchString)

	// find options that match fuzzy search string
	rankedResults := fuzzy.RankFind(p.fuzzySearchString, p.fuzzySearchMatches)
	// map rankedResults to fuzzySearchMatches
	p.fuzzySearchMatches = []string{}
	sort.Sort(rankedResults)
	for _, result := range rankedResults {
		p.fuzzySearchMatches = append(p.fuzzySearchMatches, result.Target)
	}

	for i, option := range p.fuzzySearchMatches {
		if i == p.selectedOption {
			content += Sprintf("%s %s\n", p.renderSelector(), option)
		} else {
			content += Sprintf("  %s\n", option)
		}
	}

	return content
}

func (p InteractiveSelectPrinter) renderFinishedMenu() string {
	var content string
	content += Sprintf("%s: %s\n", p.text, p.fuzzySearchString)
	content += Sprintf("  %s %s\n", p.renderSelector(), p.result)

	return content
}

func (p InteractiveSelectPrinter) renderSelector() string {
	return p.SelectorStyle.Sprint(p.Selector)
}
