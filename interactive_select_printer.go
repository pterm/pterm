package pterm

import (
	"fmt"
	"math"
	"sort"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/pterm/pterm/internal"
)

var (
	// DefaultInteractiveSelect is the default InteractiveSelect printer.
	DefaultInteractiveSelect = InteractiveSelectPrinter{
		TextStyle:     &ThemeDefault.PrimaryStyle,
		DefaultText:   "Please select an option",
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
	DefaultText   string
	Options       []string
	OptionStyle   *Style
	DefaultOption string
	MaxHeight     int
	Selector      string
	SelectorStyle *Style

	selectedOption        int
	result                string
	text                  string
	fuzzySearchString     string
	fuzzySearchMatches    []string
	displayedOptions      []string
	displayedOptionsStart int
	displayedOptionsEnd   int
}

// WithDefaultText sets the default text.
func (p InteractiveSelectPrinter) WithDefaultText(text string) *InteractiveSelectPrinter {
	p.DefaultText = text
	return &p
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

// Show shows the interactive select menu and returns the selected entry.
func (p *InteractiveSelectPrinter) Show(text ...string) (string, error) {
	// should be the first defer statement to make sure it is executed last
	// and all the needed cleanup can be done before
	cancel, exit := internal.NewCancelationSignal()
	defer exit()

	if len(text) == 0 || Sprint(text[0]) == "" {
		text = []string{p.DefaultText}
	}

	p.text = p.TextStyle.Sprint(text[0])
	p.fuzzySearchMatches = append([]string{}, p.Options...)

	if p.MaxHeight == 0 {
		p.MaxHeight = DefaultInteractiveSelect.MaxHeight
	}

	maxHeight := p.MaxHeight
	if maxHeight > len(p.fuzzySearchMatches) {
		maxHeight = len(p.fuzzySearchMatches)
	}

	if len(p.Options) == 0 {
		return "", fmt.Errorf("no options provided")
	}

	p.displayedOptions = append([]string{}, p.fuzzySearchMatches[:maxHeight]...)
	p.displayedOptionsStart = 0
	p.displayedOptionsEnd = maxHeight

	// Get index of default option
	if p.DefaultOption != "" {
		for i, option := range p.Options {
			if option == p.DefaultOption {
				p.selectedOption = i
				if i > 0 && len(p.Options) > maxHeight {
					p.displayedOptionsEnd = int(math.Min(float64(i-1+maxHeight), float64(len(p.Options))))
					p.displayedOptionsStart = p.displayedOptionsEnd - maxHeight
				} else {
					p.displayedOptionsStart = 0
					p.displayedOptionsEnd = maxHeight
				}
				p.displayedOptions = p.Options[p.displayedOptionsStart:p.displayedOptionsEnd]
				break
			}
		}
	}

	area, err := DefaultArea.Start(p.renderSelectMenu())
	defer area.Stop()
	if err != nil {
		return "", fmt.Errorf("could not start area: %w", err)
	}

	area.Update(p.renderSelectMenu())

	cursor.Hide()
	defer cursor.Show()

	err = keyboard.Listen(func(keyInfo keys.Key) (stop bool, err error) {
		key := keyInfo.Code

		if p.MaxHeight > len(p.fuzzySearchMatches) {
			maxHeight = len(p.fuzzySearchMatches)
		} else {
			maxHeight = p.MaxHeight
		}

		switch key {
		case keys.RuneKey:
			// Fuzzy search for options
			// append to fuzzy search string
			p.fuzzySearchString += keyInfo.String()
			p.selectedOption = 0
			p.displayedOptionsStart = 0
			p.displayedOptionsEnd = maxHeight
			p.displayedOptions = append([]string{}, p.fuzzySearchMatches[:maxHeight]...)
			area.Update(p.renderSelectMenu())
		case keys.Space:
			p.fuzzySearchString += " "
			p.selectedOption = 0
			area.Update(p.renderSelectMenu())
		case keys.Backspace:
			// Remove last character from fuzzy search string
			if len(p.fuzzySearchString) > 0 {
				// Handle UTF-8 characters
				p.fuzzySearchString = string([]rune(p.fuzzySearchString)[:len([]rune(p.fuzzySearchString))-1])
			}

			if p.fuzzySearchString == "" {
				p.fuzzySearchMatches = append([]string{}, p.Options...)
			}

			p.renderSelectMenu()

			if len(p.fuzzySearchMatches) > p.MaxHeight {
				maxHeight = p.MaxHeight
			} else {
				maxHeight = len(p.fuzzySearchMatches)
			}

			p.selectedOption = 0
			p.displayedOptionsStart = 0
			p.displayedOptionsEnd = maxHeight
			p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)

			area.Update(p.renderSelectMenu())
		case keys.Up:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			if p.selectedOption > 0 {
				p.selectedOption--
				if p.selectedOption < p.displayedOptionsStart {
					p.displayedOptionsStart--
					p.displayedOptionsEnd--
					if p.displayedOptionsStart < 0 {
						p.displayedOptionsStart = 0
						p.displayedOptionsEnd = maxHeight
					}
					p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
				}
			} else {
				p.selectedOption = len(p.fuzzySearchMatches) - 1
				p.displayedOptionsStart = len(p.fuzzySearchMatches) - maxHeight
				p.displayedOptionsEnd = len(p.fuzzySearchMatches)
				p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
			}

			area.Update(p.renderSelectMenu())
		case keys.Down:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			p.displayedOptions = p.fuzzySearchMatches[:maxHeight]
			if p.selectedOption < len(p.fuzzySearchMatches)-1 {
				p.selectedOption++
				if p.selectedOption >= p.displayedOptionsEnd {
					p.displayedOptionsStart++
					p.displayedOptionsEnd++
					p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
				}
			} else {
				p.selectedOption = 0
				p.displayedOptionsStart = 0
				p.displayedOptionsEnd = maxHeight
				p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
			}

			area.Update(p.renderSelectMenu())
		case keys.CtrlC:
			cancel()
			return true, nil
		case keys.Enter:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
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

func (p *InteractiveSelectPrinter) renderSelectMenu() string {
	var content string
	content += Sprintf("%s %s: %s\n", p.text, p.SelectorStyle.Sprint("[type to search]"), p.fuzzySearchString)

	// find options that match fuzzy search string
	rankedResults := fuzzy.RankFindFold(p.fuzzySearchString, p.Options)
	// map rankedResults to fuzzySearchMatches
	p.fuzzySearchMatches = []string{}
	if len(rankedResults) != len(p.Options) {
		sort.Sort(rankedResults)
	}
	for _, result := range rankedResults {
		p.fuzzySearchMatches = append(p.fuzzySearchMatches, result.Target)
	}

	if len(p.fuzzySearchMatches) != 0 {
		p.result = p.fuzzySearchMatches[p.selectedOption]
	}

	indexMapper := make([]string, len(p.fuzzySearchMatches))
	for i := 0; i < len(p.fuzzySearchMatches); i++ {
		// if in displayed options range
		if i >= p.displayedOptionsStart && i < p.displayedOptionsEnd {
			indexMapper[i] = p.fuzzySearchMatches[i]
		}
	}

	for i, option := range indexMapper {
		if option == "" {
			continue
		}
		if i == p.selectedOption {
			content += Sprintf("%s %s\n", p.renderSelector(), p.OptionStyle.Sprint(option))
		} else {
			content += Sprintf("  %s\n", p.OptionStyle.Sprint(option))
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
