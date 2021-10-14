package pterm

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

var (
	// DefaultHeader returns the printer for a default header text.
	// Defaults to LightWhite, Bold Text and a Gray DefaultHeader background.
	DefaultHeader = HeaderPrinter{
		TextStyle:       &ThemeDefault.HeaderTextStyle,
		BackgroundStyle: &ThemeDefault.HeaderBackgroundStyle,
		Margin:          5,
	}
)

// HeaderPrinter contains the data used to craft a header.
// A header is printed as a big box with text in it.
// Can be used as title screens or section separator.
type HeaderPrinter struct {
	TextStyle       *Style
	BackgroundStyle *Style
	Margin          int
	FullWidth       bool
}

// WithTextStyle returns a new HeaderPrinter with changed
func (p HeaderPrinter) WithTextStyle(style *Style) *HeaderPrinter {
	p.TextStyle = style
	return &p
}

// WithBackgroundStyle changes the background styling of the header.
func (p HeaderPrinter) WithBackgroundStyle(style *Style) *HeaderPrinter {
	p.BackgroundStyle = style
	return &p
}

// WithMargin changes the background styling of the header.
func (p HeaderPrinter) WithMargin(margin int) *HeaderPrinter {
	p.Margin = margin
	return &p
}

// WithFullWidth enables full width on a HeaderPrinter.
func (p HeaderPrinter) WithFullWidth(b ...bool) *HeaderPrinter {
	p.FullWidth = internal.WithBoolean(b)
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p HeaderPrinter) Sprint(a ...interface{}) string {
	if RawOutput {
		return Sprint(a...)
	}

	if p.TextStyle == nil {
		p.TextStyle = NewStyle()
	}
	if p.BackgroundStyle == nil {
		p.BackgroundStyle = NewStyle()
	}

	text := Sprint(a...)

	var blankLine string

	longestLine := internal.ReturnLongestLine(text, "\n")
	longestLineLen := runewidth.StringWidth(RemoveColorFromString(longestLine)) + p.Margin*2

	if p.FullWidth {
		text = splitText(text, GetTerminalWidth()-p.Margin*2)
		blankLine = strings.Repeat(" ", GetTerminalWidth())
	} else {
		if longestLineLen > GetTerminalWidth() {
			text = splitText(text, GetTerminalWidth()-p.Margin*2)
			blankLine = strings.Repeat(" ", GetTerminalWidth())
		} else {
			text = splitText(text, longestLineLen-p.Margin*2)
			blankLine = strings.Repeat(" ", longestLineLen)
		}
	}

	var marginString string
	var ret string

	if p.FullWidth {
		longestLineLen = runewidth.StringWidth(RemoveColorFromString(internal.ReturnLongestLine(text, "\n")))
		marginString = strings.Repeat(" ", (GetTerminalWidth()-longestLineLen)/2)
	} else {
		marginString = strings.Repeat(" ", p.Margin)
	}

	ret += p.BackgroundStyle.Sprint(blankLine) + "\n"
	for _, line := range strings.Split(text, "\n") {
		line = strings.ReplaceAll(line, "\n", "")
		line = marginString + line + marginString
		if runewidth.StringWidth(line) < runewidth.StringWidth(blankLine) {
			line += strings.Repeat(" ", runewidth.StringWidth(blankLine)-runewidth.StringWidth(line))
		}
		ret += p.BackgroundStyle.Sprint(p.TextStyle.Sprint(line)) + "\n"
	}
	ret += p.BackgroundStyle.Sprint(blankLine) + "\n"

	return ret
}

func splitText(text string, width int) string {
	var lines []string
	linesTmp := strings.Split(text, "\n")
	for _, line := range linesTmp {
		if runewidth.StringWidth(RemoveColorFromString(line)) > width {
			extraLines := []string{""}
			extraLinesCounter := 0
			for i, letter := range line {
				if i%width == 0 && i != 0 {
					extraLinesCounter++
					extraLines = append(extraLines, "")
				}
				extraLines[extraLinesCounter] += string(letter)
			}
			for _, extraLine := range extraLines {
				extraLine += "\n"
				lines = append(lines, extraLine)
			}
		} else {
			line += "\n"
			lines = append(lines, line)
		}
	}

	var line string
	for _, s := range lines {
		line += s
	}

	return strings.TrimSuffix(line, "\n")
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p HeaderPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(strings.TrimSuffix(Sprintln(a...), "\n"))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p HeaderPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p HeaderPrinter) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *HeaderPrinter) Print(a ...interface{}) *TextPrinter {
	Print(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *HeaderPrinter) Println(a ...interface{}) *TextPrinter {
	Print(p.Sprintln(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *HeaderPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *HeaderPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p *HeaderPrinter) PrintOnError(a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				p.Println(err)
			}
		}
	}

	tp := TextPrinter(p)
	return &tp
}

// PrintOnErrorf wraps every error which is not nil and prints it.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p *HeaderPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				p.Println(fmt.Errorf(format, err))
			}
		}
	}

	tp := TextPrinter(p)
	return &tp
}
