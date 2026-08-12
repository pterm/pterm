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

	input         []string
	fitInput      []string
	cursorXPos    int
	cursorYPos    int
	text          string
	startedTyping bool
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

	if len(text) == 0 || text[0] == "" {
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

	p.input = append(p.input, strings.Split(p.DefaultValue, "\n")...)
	p.cursorYPos = len(p.input) - 1
	p.updateArea(&area, p.input, p.cursorXPos, p.cursorYPos)

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {

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
			if !p.startedTyping {
				p.startedTyping = true
			}

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
				p.cursorXPos = -getMaxW(p.input[p.cursorYPos])

				cursor.StartOfLine()
			} else {
				return true, nil
			}

		case keys.RuneKey:
			if !p.startedTyping {
				p.startedTyping = true
			}

			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(key.String()), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))

		case keys.Space:
			if !p.startedTyping {
				p.startedTyping = true
			}

			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(" "), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))

		case keys.Backspace:
			if !p.startedTyping {
				p.startedTyping = true
			}

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
			if !p.startedTyping {
				p.input = []string{""}
				p.startedTyping = true

				return false, nil
			}

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
			if !p.MultiLine {
				return false, nil
			}

			if !p.startedTyping {
				p.startedTyping = true
			}

			if p.cursorYPos+1 < len(p.input) {
				p.cursorXPos = min((getMaxW(p.input[p.cursorYPos])+p.cursorXPos)-getMaxW(p.input[p.cursorYPos+1]), 0)

				p.cursorYPos++
			}

		case keys.Up:
			if !p.MultiLine {
				return false, nil
			}

			if !p.startedTyping {
				p.startedTyping = true
			}

			if p.cursorYPos > 0 {
				p.cursorXPos = min((getMaxW(p.input[p.cursorYPos])+p.cursorXPos)-getMaxW(p.input[p.cursorYPos-1]), 0)

				p.cursorYPos--
			}
		}

		if getMaxW(p.input[p.cursorYPos]) > 0 {
			switch key.Code {
			case keys.Right:
				if p.cursorXPos < 0 {
					p.cursorXPos++
				} else if p.cursorYPos < len(p.input)-1 {
					p.cursorYPos++
					p.cursorXPos = -getMaxW(p.input[p.cursorYPos])
				}

			case keys.Left:
				if p.cursorXPos+getMaxW(p.input[p.cursorYPos]) > 0 {
					p.cursorXPos--
				} else if p.cursorYPos > 0 {
					p.cursorYPos--
					p.cursorXPos = 0
				}
			}
		}

		// for test
		{
			l := func(a ...any) { p.text += Sprintln(a...) }
			pink := NewRGB(255, 0, 200).Sprint
			y := NewRGB(251, 255, 0).Sprint
			b := NewRGB(88, 91, 255).Sprint
			g := NewRGB(21, 255, 0).Sprint

			p.text = LightRed("--------------\n")
			l(Sprint(pink("███"), y("███"), b("███"), g("███")))
			l(Sprintf("%v Y:%v X:%v", pink("logic"), y(p.cursorYPos), g(p.cursorXPos)))
			l(pink("input:"), b(Sprintf("%q", p.input)))
			p.text += LightRed("--------------\n")
		}

		// update the input buffer
		inputBuffer := make([]string, len(p.input))
		// handle the mask
		if p.Mask != "" {
			for _, s := range p.input {
				inputBuffer = append(inputBuffer, strings.Repeat(p.Mask, getMaxW(s)))
			}
		} else {
			inputBuffer = p.input
		}

		// TODO update area with actual coord
		p.updateArea(&area, inputBuffer, p.cursorXPos, p.cursorYPos)

		return false, nil
	})
	if err != nil {
		return "", err
	}

	// Add new line
	Println()

	if !p.startedTyping {
		return p.DefaultValue, nil
	}

	return strings.Join(p.input, "\n"), nil
}

func (p InteractiveTextInputPrinter) updateArea(area *cursor.Area, input []string, x, y int) string {

	areaText := textFitWidth(p.text)
	areaContent := areaText

	// TODO fit input
	x, y = p.cursorXPos, p.cursorYPos
	areaInput := input

	// // reserved code
	// if x+getMaxW(areaInput[y]) < 1 {
	// 	x = -getMaxW(areaInput[y])
	// }

	area.Update(areaContent)
	// cursor down offset
	area.Top()
	area.Down(strings.Count(areaText, "\n") + y)
	// cursor right offset
	area.StartOfLine()
	if p.MultiLine || y != 0 {
		cursor.Right(getMaxW(areaInput[y]) + x)
	} else {
		lines := strings.Split(p.text, "\n")
		cursor.Right(getMaxW(lines[len(lines)-1]) + getMaxW(areaInput[y]) + x)
	}

	return areaContent
}
