package pterm

import (
	"strings"
	"sync"
	"time"

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
	actualX       int
	actualY       int
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

	type coord struct{ y, end int }
	// coordinate mapping
	// c:=cMap[ay],py=c.y,px=ax+c.end
	var cMap []coord
	updateCoords := func() {
		cMap = make([]coord, 0, len(p.fitInput))
		for py, logicLine := range p.input {
			if getMaxW(logicLine) < GetTerminalWidth() {
				cMap = append(cMap, coord{y: py, end: 0})
				continue
			}
			px := -getMaxW(logicLine)
			for _, line := range linesFitWidth([]string{logicLine}) {
				px += getMaxW(line)
				cMap = append(cMap, coord{y: py, end: px})
			}
		}
	}
	logicYX := func(ay, ax int) (int, int) {
		c := cMap[ay]
		return c.y, ax + c.end
	}
	actualYX := func(py, px int) (int, int) {
		var y, x int
		for ay, c := range cMap {
			if py == c.y && px <= c.end {
				y, x = ay, px-c.end
				break
			}
		}
		return y, x
	}
	updateFitInput := func() { p.fitInput = linesFitWidth(p.input) }
	updateLogicYX := func() {
		updateFitInput()
		updateCoords()
		p.cursorYPos, p.cursorXPos = logicYX(p.actualY, p.actualX)
	}
	updateActualYX := func() {
		updateFitInput()
		updateCoords()
		p.actualY, p.actualX = actualYX(p.cursorYPos, p.cursorXPos)
	}

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
	// !remove brefore PR
	textLog := func() {
		l := func(a ...any) { p.text += Sprintln(a...) }
		pink := NewRGB(255, 0, 200).Sprintf
		y := NewRGB(251, 255, 0).Sprint
		b := NewRGB(88, 91, 255).Sprint
		g := NewRGB(21, 255, 0).Sprint
		o := NewRGB(255, 94, 0).Sprint

		p.text = LightRed("--------------\n")
		l(Sprint(pink("█"), y("█"), b("█"), g("█"), o("█")))
		l(Sprint(pink("width:"), o(GetTerminalWidth())))
		l(Sprintf("%v Y:%v X:%v L:%v",
			pink("lgc"), g(p.cursorYPos), g(p.cursorXPos),
			y(getMaxW(p.input[p.cursorYPos]))))
		l(Sprintf("%v Y:%v X:%v L:%v",
			pink("act"), g(p.actualY), g(p.actualX),
			y(getMaxW(p.fitInput[p.actualY]))))

		l(b(Sprintf("%q", p.input)))
		l(b(Sprintf("%q", p.fitInput)))
		for k, v := range cMap {
			l(b(Sprintf("%v:(y:%v,end:%v)", k, v.y, v.end)))
		}

		p.text += LightRed("--------------\n")
	}

	p.input = append(p.input, strings.Split(p.DefaultValue, "\n")...)
	p.cursorYPos = len(p.input) - 1
	p.cursorXPos = 0
	p.fitInput = linesFitWidth(p.input)
	updateActualYX()

	// !remove brefore PR
	textLog()
	area := cursor.NewArea()

	p.updateArea(&area, p.fitInput)

	// watch and fit the terminal width
	var mu sync.Mutex
	fitDone := make(chan struct{})
	defer close(fitDone)
	go func() {
		watchWidth(fitDone, 100*time.Millisecond, func(w int) {
			mu.Lock()
			defer mu.Unlock()
			updateActualYX()
			p.updateArea(&area, p.fitInput)
		})
	}()

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		mu.Lock()
		defer mu.Unlock()

		if len(p.input) == 0 {
			p.input = append(p.input, "")
		}

		updateLogicYX()

		switch key.Code {
		case keys.Tab:
			if p.MultiLine {
				area.Bottom()
				return true, nil
			}

		case keys.Enter:
			// TODO review Enter case
			if !p.startedTyping {
				p.startedTyping = true
			}

			if p.MultiLine {
				if key.AltPressed {
					p.actualX = 0
				}

				updateLogicYX()

				appendAfterY := append([]string{}, p.input[p.cursorYPos+1:]...)
				appendAfterX := string(append([]rune{}, []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...))
				p.input[p.cursorYPos] = string(append([]rune{}, []rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos]...))
				p.input = append(p.input[:p.cursorYPos+1], appendAfterX)
				p.input = append(p.input, appendAfterY...)
				p.cursorYPos++
				p.cursorXPos = -getMaxW(p.input[p.cursorYPos])

				updateActualYX()

			} else {
				return true, nil
			}

		case keys.RuneKey:
			if !p.startedTyping {
				p.startedTyping = true
			}

			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(key.String()), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))
			updateActualYX()

		case keys.Space:
			if !p.startedTyping {
				p.startedTyping = true
			}

			p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], append([]rune(" "), []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...)...))
			updateActualYX()

		case keys.Backspace:
			// TODO Backspace case
			if !p.startedTyping {
				p.startedTyping = true
			}

			handle := func() {
				if len([]rune(p.input[p.cursorYPos]))+p.cursorXPos > 0 {
					p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))-1+p.cursorXPos], []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos:]...))
				} else if p.cursorYPos > 0 {
					p.input[p.cursorYPos-1] += p.input[p.cursorYPos]
					appendAfterY := append([]string{}, p.input[p.cursorYPos+1:]...)
					p.input = append(p.input[:p.cursorYPos], appendAfterY...)
					p.cursorXPos = 0
					p.cursorYPos--
				}
			}
			handle()

		case keys.Delete:
			// TODO Delete case
			if !p.startedTyping {
				p.startedTyping = true
			}

			handle := func() {
				if len([]rune(p.input[p.cursorYPos]))+p.cursorXPos < len([]rune(p.input[p.cursorYPos])) {
					p.input[p.cursorYPos] = string(append([]rune(p.input[p.cursorYPos])[:len([]rune(p.input[p.cursorYPos]))+p.cursorXPos], []rune(p.input[p.cursorYPos])[len([]rune(p.input[p.cursorYPos]))+p.cursorXPos+1:]...))
					p.cursorXPos++
				} else if p.cursorYPos < len(p.input)-1 {
					p.input[p.cursorYPos] += p.input[p.cursorYPos+1]
					appendAfterY := append([]string{}, p.input[p.cursorYPos+2:]...)
					p.input = append(p.input[:p.cursorYPos+1], appendAfterY...)
					p.cursorXPos = 0
				}
			}
			handle()

		case keys.CtrlC:
			cancel()
			return true, nil

		case keys.Down:
			if !p.startedTyping {
				p.startedTyping = true
			}

			if p.actualY+1 < len(p.fitInput) {
				p.actualX = min((getMaxW(p.fitInput[p.actualY])+p.actualX)-getMaxW(p.fitInput[p.actualY+1]), 0)

				p.actualY++
			}

		case keys.Up:
			if !p.startedTyping {
				p.startedTyping = true
			}

			if p.actualY > 0 {
				p.actualX = min((getMaxW(p.fitInput[p.actualY])+p.actualX)-getMaxW(p.fitInput[p.actualY-1]), 0)

				p.actualY--
			}

		case keys.Right:
			if !p.startedTyping {
				p.startedTyping = true
			}
			if p.actualX < 0 {
				p.actualX++
			} else if p.actualY < len(p.fitInput)-1 {
				p.actualY++
				p.actualX = -getMaxW(p.fitInput[p.actualY])
			}

		case keys.Left:
			if !p.startedTyping {
				p.startedTyping = true
			}
			if p.actualX+getMaxW(p.fitInput[p.actualY]) > 0 {
				p.actualX--
			} else if p.actualY > 0 {
				p.actualY--
				p.actualX = 0
			}

		case keys.Esc:
		}

		// update logic coord
		updateLogicYX()

		// !remove brefore PR
		textLog()

		// update the input buffer
		areaInput := make([]string, 0)
		// handle the mask
		if p.Mask != "" {
			for _, s := range p.input {
				areaInput = append(areaInput, strings.Repeat(p.Mask, getMaxW(s)))
			}
		} else {
			areaInput = p.input
		}

		areaInput = linesFitWidth(areaInput)
		p.updateArea(&area, areaInput)

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

func (p InteractiveTextInputPrinter) updateArea(area *cursor.Area, fitInput []string) string {

	areaText := textFitWidth(p.text)
	areaContent := areaText
	areaContent += strings.Join(fitInput, "\n")

	area.Update(areaContent)

	x, y := p.actualX, p.actualY
	// cursor down offset
	area.Top()
	area.Down(strings.Count(areaText, "\n") + y)
	// cursor right offset
	area.StartOfLine()
	if p.MultiLine || y != 0 {
		cursor.Right(getMaxW(fitInput[y]) + x)
	} else {
		lines := strings.Split(p.text, "\n")
		cursor.Right(getMaxW(lines[len(lines)-1]) + getMaxW(fitInput[y]) + x)
	}

	return areaContent
}
