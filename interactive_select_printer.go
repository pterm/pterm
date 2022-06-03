package pterm

import (
	"fmt"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

var (
	// DefaultInteractiveSelect is the default InteractiveSelect printer.
	DefaultInteractiveSelect = InteractiveSelectPrinter{}
)

type InteractiveSelectPrinter struct {
	Options       []string
	DefaultOption string
	MaxHeight     int

	selectedOption int
	result         string
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

func (p InteractiveSelectPrinter) Show(text ...string) (string, error) {
	err := keyboard.Listen(func(keyInfo keys.Key) (stop bool, err error) {
		if text == nil {
			text = []string{""}
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
		if err != nil {
			return true, fmt.Errorf("could not start area: %w", err)
		}

		for p.result == "" {
			key := keyInfo.Code

			switch key {
			case keys.Up:
				if p.selectedOption > 0 {
					p.selectedOption--
				}
				area.Update(p.renderSelectMenu())
			case keys.Down:
				if p.selectedOption < len(p.Options)-1 {
					p.selectedOption++
				}
				area.Update(p.renderSelectMenu())
			case keys.CtrlC:
				os.Exit(1)
			case keys.Enter:
				p.result = p.Options[p.selectedOption]
			}
		}
		return true, nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to start keyboard listener: %w", err)
	}

	return p.result, nil
}

func (p InteractiveSelectPrinter) renderSelectMenu() string {
	var content string
	for i, option := range p.Options {
		if i == p.selectedOption {
			content += fmt.Sprintf("> %s\n", option)
		} else {
			content += fmt.Sprintf("  %s\n", option)
		}
	}

	return content
}
