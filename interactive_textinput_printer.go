package pterm

import (
	"strings"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm/internal"
)

// DefaultInteractiveTextInput is the default InteractiveTextInput printer.
var DefaultInteractiveTextInput = InteractiveTextInputPrinter{
	DefaultText: "Input text",
	Delimiter:   ": ",
	TextStyle:   &ThemeDefault.PrimaryStyle,
	Mask:        "",
}

// InteractiveTextInputPrinter is a printer for interactive select menus.
type InteractiveTextInputPrinter struct {
	TextStyle       *Style
	DefaultText     string
	DefaultValue    string
	Delimiter       string
	MultiLine       bool
	Mask            string
	OnInterruptFunc func()

	input      []string
	cursorXPos int
	cursorYPos int
	text       string
}

// WithDefaultText sets the default text.
func (p InteractiveTextInputPrinter) WithDefaultText(text string) *InteractiveTextInputPrinter {
	p.DefaultText = text
	return &p
}

// WithDefaultValue sets the default value.
func (p InteractiveTextInputPrinter) WithDefaultValue(value string) *InteractiveTextInputPrinter {
	p.DefaultValue = value
	return &p
}

// WithTextStyle sets the text style.
func (p InteractiveTextInputPrinter) WithTextStyle(style *Style) *InteractiveTextInputPrinter {
	p.TextStyle = style
	return &p
}

// WithMultiLine sets the multi line flag.
func (p InteractiveTextInputPrinter) WithMultiLine(multiLine ...bool) *InteractiveTextInputPrinter {
	p.MultiLine = internal.WithBoolean(multiLine)
	return &p
}

// WithMask sets the mask.
func (p InteractiveTextInputPrinter) WithMask(mask string) *InteractiveTextInputPrinter {
	p.Mask = mask
	return &p
}

// WithOnInterruptFunc sets the function to execute on exit of the input reader
func (p InteractiveTextInputPrinter) WithOnInterruptFunc(exitFunc func()) *InteractiveTextInputPrinter {
	p.OnInterruptFunc = exitFunc
	return &p
}

// WithDelimiter sets the delimiter between the message and the input.
func (p InteractiveTextInputPrinter) WithDelimiter(delimiter string) *InteractiveTextInputPrinter {
	p.Delimiter = delimiter
	return &p
}

// Show shows the interactive select menu and returns the selected entry.
func (p InteractiveTextInputPrinter) Show(text ...string) (string, error) {
	// should be the first defer statement to make sure it is executed last
	// and all the needed cleanup can be done before
	cancel, exit := internal.NewCancelationSignal(p.OnInterruptFunc)
	defer exit()

	var areaText string

	if len(text) == 0 || Sprint(text[0]) == "" {
		text = []string{p.DefaultText}
	}

	if p.MultiLine {
		areaText = p.TextStyle.Sprintfln("%s %s %s", text[0], ThemeDefault.SecondaryStyle.Sprint("[Press tab to submit]"), p.Delimiter)
	} else {
		areaText = p.TextStyle.Sprintf("%s%s", text[0], p.Delimiter)
	}

	p.text = areaText
	area := cursor.NewArea()
	area.Update(areaText)
	area.StartOfLine()

	if !p.MultiLine {
		cursor.Right(len(RemoveColorFromString(areaText)))
	}

	if p.DefaultValue != "" {
		p.input = append(p.input, p.DefaultValue)
		p.updateArea(&area)
	}

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if !p.MultiLine {
			p.cursorYPos = 0
		}
		if len(p.input) == 0 {
			p.input = append(p.input, "")
		}

		switch key.Code {
		case keys.Tab:
			if p.MultiLine {
				area.Bottom()
				return true, nil
			}
		case keys.Enter:
			if p.MultiLine {
				if key.AltPressed {
					p.cursorXPos = 0
				}
				appendAfterY := append([]string{}, p.input[p.cursorYPos+1:]...)
				appendAfterX := string(append([]rune{}, []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...))
				p.input[p.cursorYPos] = string(append([]rune{}, []rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos]...))
				p.input = append(p.input[:p.cursorYPos+1], appendAfterX)
				p.input = append(p.input, appendAfterY...)
				p.cursorYPos++
				p.cursorXPos = -internal.GetStringMaxWidth(p.input[p.cursorYPos])
				cursor.StartOfLine()
			} else {
				return true, nil
			}
		case keys.RuneKey:
			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(key.String()), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))
		case keys.Space:
			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(" "), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))
		case keys.Backspace:
			if len([]rune(p.input[p.cursorYPos]))+p.cursorXPos > 0 {
				p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))-1+p.cursorXPos], []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...))
			} else if p.cursorYPos > 0 {
				p.input[p.cursorYPos-1] += p.input[p.cursorYPos]
				appendAfterY := append([]string{}, p.input[p.cursorYPos+1:]...)
				p.input = append(p.input[:p.cursorYPos], appendAfterY...)
				p.cursorXPos = 0
				p.cursorYPos--
			}
		case keys.Delete:
			if len([]rune(p.input[p.cursorYPos]))+p.cursorXPos < len([]rune(p.input[p.cursorYPos])) {
				p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos+1:]...))
				p.cursorXPos++
			} else if p.cursorYPos < len(p.input)-1 {
				p.input[p.cursorYPos] += p.input[p.cursorYPos+1]
				appendAfterY := append([]string{}, p.input[p.cursorYPos+2:]...)
				p.input = append(p.input[:p.cursorYPos+1], appendAfterY...)
				p.cursorXPos = 0
			}
		case keys.CtrlC:
			cancel()
			return true, nil
		case keys.Down:
			if p.cursorYPos+1 < len(p.input) {
				p.cursorXPos = (internal.GetStringMaxWidth(p.input[p.cursorYPos]) + p.cursorXPos) - internal.GetStringMaxWidth(p.input[p.cursorYPos+1])
				if p.cursorXPos > 0 {
					p.cursorXPos = 0
				}
				p.cursorYPos++
			}
		case keys.Up:
			if p.cursorYPos > 0 {
				p.cursorXPos = (internal.GetStringMaxWidth(p.input[p.cursorYPos]) + p.cursorXPos) - internal.GetStringMaxWidth(p.input[p.cursorYPos-1])
				if p.cursorXPos > 0 {
					p.cursorXPos = 0
				}
				p.cursorYPos--
			}
		}

		if internal.GetStringMaxWidth(p.input[p.cursorYPos]) > 0 {
			switch key.Code {
			case keys.Right:
				if p.cursorXPos < 0 {
					p.cursorXPos++
				} else if p.cursorYPos < len(p.input)-1 {
					p.cursorYPos++
					p.cursorXPos = -internal.GetStringMaxWidth(p.input[p.cursorYPos])
				}
			case keys.Left:
				if p.cursorXPos+internal.GetStringMaxWidth(p.input[p.cursorYPos]) > 0 {
					p.cursorXPos--
				} else if p.cursorYPos > 0 {
					p.cursorYPos--
					p.cursorXPos = 0
				}
			}
		}

		p.updateArea(&area)

		return false, nil
	})
	if err != nil {
		return "", err
	}

	// Add new line
	Println()

	for i, s := range p.input {
		if i < len(p.input)-1 {
			areaText += s + "\n"
		} else {
			areaText += s
		}
	}

	return strings.ReplaceAll(areaText, p.text, ""), nil
}

func (p InteractiveTextInputPrinter) updateArea(area *cursor.Area) string {
	if !p.MultiLine {
		p.cursorYPos = 0
	}
	areaText := p.text

	for i, s := range p.input {
		if i < len(p.input)-1 {
			areaText += s + "\n"
		} else {
			areaText += s
		}
	}

	if p.Mask != "" {
		areaText = p.text + strings.Repeat(p.Mask, internal.GetStringMaxWidth(areaText)-internal.GetStringMaxWidth(p.text))
	}

	if p.cursorXPos+internal.GetStringMaxWidth(p.input[p.cursorYPos]) < 1 {
		p.cursorXPos = -internal.GetStringMaxWidth(p.input[p.cursorYPos])
	}

	area.Update(areaText)
	area.Top()
	area.Down(p.cursorYPos + 1)
	area.StartOfLine()
	if p.MultiLine {
		cursor.Right(internal.GetStringMaxWidth(p.input[p.cursorYPos]) + p.cursorXPos)
	} else {
		cursor.Right(internal.GetStringMaxWidth(areaText) + p.cursorXPos)
	}
	return areaText
}
