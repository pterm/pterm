package pterm

import (
	"fmt"
	"sort"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/pterm/pterm/internal"
)

var (
	// DefaultInteractiveMultiselect is the default InteractiveMultiselect printer.
	DefaultInteractiveMultiselect = InteractiveMultiselectPrinter{
		TextStyle:      &ThemeDefault.PrimaryStyle,
		DefaultText:    "Please select your options",
		Options:        []string{},
		OptionStyle:    &ThemeDefault.DefaultText,
		DefaultOptions: []string{},
		MaxHeight:      5,
		Selector:       ">",
		SelectorStyle:  &ThemeDefault.SecondaryStyle,
		Filter:         true,
		KeySelect:      keys.Enter,
		KeyConfirm:     keys.Tab,
	}
)

// InteractiveMultiselectPrinter is a printer for interactive multiselect menus.
type InteractiveMultiselectPrinter struct {
	DefaultText    string
	TextStyle      *Style
	Options        []string
	OptionStyle    *Style
	DefaultOptions []string
	MaxHeight      int
	Selector       string
	SelectorStyle  *Style
	Filter         bool

	selectedOption        int
	selectedOptions       []int
	text                  string
	fuzzySearchString     string
	fuzzySearchMatches    []string
	displayedOptions      []string
	displayedOptionsStart int
	displayedOptionsEnd   int

	KeySelect  keys.KeyCode
	KeyConfirm keys.KeyCode
}

// WithOptions sets the options.
func (p InteractiveMultiselectPrinter) WithOptions(options []string) *InteractiveMultiselectPrinter {
	p.Options = options
	return &p
}

// WithDefaultOptions sets the default options.
func (p InteractiveMultiselectPrinter) WithDefaultOptions(options []string) *InteractiveMultiselectPrinter {
	p.DefaultOptions = options
	return &p
}

// WithDefaultText sets the default text.
func (p InteractiveMultiselectPrinter) WithDefaultText(text string) *InteractiveMultiselectPrinter {
	p.DefaultText = text
	return &p
}

// WithMaxHeight sets the maximum height of the select menu.
func (p InteractiveMultiselectPrinter) WithMaxHeight(maxHeight int) *InteractiveMultiselectPrinter {
	p.MaxHeight = maxHeight
	return &p
}

// WithFilter sets the Filter option
func (p InteractiveMultiselectPrinter) WithFilter(filter bool) *InteractiveMultiselectPrinter {
	p.Filter = filter
	return &p
}

// WithKeySelect sets the confirm key
func (p InteractiveMultiselectPrinter) WithKeySelect(keySelect keys.KeyCode) *InteractiveMultiselectPrinter {
	p.KeySelect = keySelect
	return &p
}

// WithKeyConfirm sets the confirm key
func (p InteractiveMultiselectPrinter) WithKeyConfirm(keyConfirm keys.KeyCode) *InteractiveMultiselectPrinter {
	p.KeyConfirm = keyConfirm
	return &p
}

// Show shows the interactive multiselect menu and returns the selected entry.
func (p *InteractiveMultiselectPrinter) Show(text ...string) ([]string, error) {
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
		p.MaxHeight = DefaultInteractiveMultiselect.MaxHeight
	}

	maxHeight := p.MaxHeight
	if maxHeight > len(p.fuzzySearchMatches) {
		maxHeight = len(p.fuzzySearchMatches)
	}

	if len(p.Options) == 0 {
		return nil, fmt.Errorf("no options provided")
	}

	p.displayedOptions = append([]string{}, p.fuzzySearchMatches[:maxHeight]...)
	p.displayedOptionsStart = 0
	p.displayedOptionsEnd = maxHeight

	for _, option := range p.DefaultOptions {
		p.selectOption(option)
	}

	area, err := DefaultArea.Start(p.renderSelectMenu())
	defer area.Stop()
	if err != nil {
		return nil, fmt.Errorf("could not start area: %w", err)
	}

	if p.Filter && (p.KeyConfirm == keys.Space || p.KeySelect == keys.Space) {
		return nil, fmt.Errorf("if filter/search is active, keys.Space can not be used for KeySelect or KeyConfirm")
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
		case p.KeyConfirm:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			area.Update(p.renderFinishedMenu())
			return true, nil
		case p.KeySelect:
			if len(p.fuzzySearchMatches) > 0 {
				// Select option if not already selected
				p.selectOption(p.fuzzySearchMatches[p.selectedOption])
			}
			area.Update(p.renderSelectMenu())
		case keys.RuneKey:
			if p.Filter {
				// Fuzzy search for options
				// append to fuzzy search string
				p.fuzzySearchString += keyInfo.String()
				p.selectedOption = 0
				p.displayedOptionsStart = 0
				p.displayedOptionsEnd = maxHeight
				p.displayedOptions = append([]string{}, p.fuzzySearchMatches[:maxHeight]...)
			}
			area.Update(p.renderSelectMenu())
		case keys.Space:
			if p.Filter {
				p.fuzzySearchString += " "
				p.selectedOption = 0
				area.Update(p.renderSelectMenu())
			}
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
		case keys.Left:
			// Unselect all options
			p.selectedOptions = []int{}
			area.Update(p.renderSelectMenu())
		case keys.Right:
			// Select all options
			p.selectedOptions = []int{}
			for i := 0; i < len(p.Options); i++ {
				p.selectedOptions = append(p.selectedOptions, i)
			}
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
		}

		return false, nil
	})
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to start keyboard listener: %w", err)
	}

	var result []string
	for _, selectedOption := range p.selectedOptions {
		result = append(result, p.Options[selectedOption])
	}

	return result, nil
}

func (p InteractiveMultiselectPrinter) findOptionByText(text string) int {
	for i, option := range p.Options {
		if option == text {
			return i
		}
	}
	return -1
}

func (p *InteractiveMultiselectPrinter) isSelected(optionText string) bool {
	for _, selectedOption := range p.selectedOptions {
		if p.Options[selectedOption] == optionText {
			return true
		}
	}

	return false
}

func (p *InteractiveMultiselectPrinter) selectOption(optionText string) {
	if p.isSelected(optionText) {
		// Remove from selected options
		for i, selectedOption := range p.selectedOptions {
			if p.Options[selectedOption] == optionText {
				p.selectedOptions = append(p.selectedOptions[:i], p.selectedOptions[i+1:]...)
				break
			}
		}
	} else {
		// Add to selected options
		p.selectedOptions = append(p.selectedOptions, p.findOptionByText(optionText))
	}
}

func (p *InteractiveMultiselectPrinter) renderSelectMenu() string {
	var content string
	content += Sprintf("%s: %s\n", p.text, p.fuzzySearchString)

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
		var checkmark string
		if p.isSelected(option) {
			checkmark = fmt.Sprintf("[%s]", Green("✓"))
		} else {
			checkmark = fmt.Sprintf("[%s]", Red("✗"))
		}
		if i == p.selectedOption {
			content += Sprintf("%s %s %s\n", p.renderSelector(), checkmark, option)
		} else {
			content += Sprintf("  %s %s\n", checkmark, option)
		}
	}

	help := fmt.Sprintf("%s: %s | %s: %s | left: %s | right: %s", p.KeySelect, Bold.Sprint("select"), p.KeyConfirm, Bold.Sprint("confirm"), Bold.Sprint("none"), Bold.Sprint("all"))
	if p.Filter {
		help += fmt.Sprintf("| type to %s", Bold.Sprint("filter"))
	}
	content += ThemeDefault.SecondaryStyle.Sprintfln(help)

	return content
}

func (p InteractiveMultiselectPrinter) renderFinishedMenu() string {
	var content string
	content += Sprintf("%s: %s\n", p.text, p.fuzzySearchString)
	for _, option := range p.selectedOptions {
		content += Sprintf("  %s %s\n", p.renderSelector(), p.Options[option])
	}

	return content
}

func (p InteractiveMultiselectPrinter) renderSelector() string {
	return p.SelectorStyle.Sprint(p.Selector)
}
