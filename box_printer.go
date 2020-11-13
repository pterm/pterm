package pterm

import (
	"strings"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

type BoxPrinter struct {
	Text                    string
	TextStyle               *Style
	VerticalString          string
	BoxStyle                *Style
	HorizontalString        string
	TopRightCornerString    string
	TopLeftCornerString     string
	BottomLeftCornerString  string
	BottomRightCornerString string
	TopPadding              int
	BottomPadding           int
	RightPadding            int
	LeftPadding             int
}

var DefaultBox = BoxPrinter{
	VerticalString:          "|",
	TopRightCornerString:    "└",
	TopLeftCornerString:     "┘",
	BottomLeftCornerString:  "┐",
	BottomRightCornerString: "┌",
	HorizontalString:        "─",
	BoxStyle:                &ThemeDefault.BoxStyle,
	TextStyle:               &ThemeDefault.BoxTextStyle,
	RightPadding:            1,
	LeftPadding:             1,
	TopPadding:              1,
	BottomPadding:           1,
}

// WithBoxStyle returns a new box with a specific box Style.
func (p BoxPrinter) WithBoxStyle(style *Style) *BoxPrinter {
	p.BoxStyle = style
	return &p
}

// WithTextStyle returns a new box with a specific text Style.
func (p BoxPrinter) WithTextStyle(style *Style) *BoxPrinter {
	p.TextStyle = style
	return &p
}

// WithText returns a new box with a specific Text.
func (p BoxPrinter) WithText(str string) *BoxPrinter {
	p.Text = str
	return &p
}

// WithTopRightCornerString returns a new box with a specific TopRightCornerString.
func (p BoxPrinter) WithTopRightCornerString(str string) *BoxPrinter {
	p.TopRightCornerString = str
	return &p
}

// WithTopLeftCornerString returns a new box with a specific TopLeftCornerString.
func (p BoxPrinter) WithTopLeftCornerString(str string) *BoxPrinter {
	p.TopLeftCornerString = str
	return &p
}

// WithBottomRightCornerString returns a new box with a specific BottomRightCornerString.
func (p BoxPrinter) WithBottomRightCornerString(str string) *BoxPrinter {
	p.BottomRightCornerString = str
	return &p
}

// WithBottomLeftCornerString returns a new box with a specific BottomLeftCornerString.
func (p BoxPrinter) WithBottomLeftCornerString(str string) *BoxPrinter {
	p.BottomLeftCornerString = str
	return &p
}

// WithVerticalString returns a new box with a specific VerticalString.
func (p BoxPrinter) WithVerticalString(str string) *BoxPrinter {
	p.VerticalString = str
	return &p
}

// WithTopPadding returns a new box with a specific TopPadding.
func (p BoxPrinter) WithTopPadding(padding int) *BoxPrinter {
	if padding < 0 {
		padding = 0
	}
	p.TopPadding = padding
	return &p
}

// WithBottomPadding returns a new box with a specific BottomPadding.
func (p BoxPrinter) WithBottomPadding(padding int) *BoxPrinter {
	if padding < 0 {
		padding = 0
	}
	p.BottomPadding = padding
	return &p
}

// WithRightPadding returns a new box with a specific RightPadding.
func (p BoxPrinter) WithRightPadding(padding int) *BoxPrinter {
	if padding < 0 {
		padding = 0
	}
	p.RightPadding = padding
	return &p
}

// WithLeftPadding returns a new box with a specific LeftPadding.
func (p BoxPrinter) WithLeftPadding(padding int) *BoxPrinter {
	if padding < 0 {
		padding = 0
	}
	p.LeftPadding = padding
	return &p
}

// Srender renders the Template as a string.
func (p BoxPrinter) Srender() (string, error) {
	if p.BoxStyle == nil {
		p.BoxStyle = NewStyle()
	}
	if p.TextStyle == nil {
		p.TextStyle = NewStyle()
	}
	maxWidth := internal.GetStringMaxWidth(p.Text)
	topLine := p.BoxStyle.Sprint(p.BottomRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.RightPadding+p.RightPadding) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
	bottomLine := p.BoxStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.RightPadding+p.RightPadding) + p.BoxStyle.Sprint(p.TopLeftCornerString)

	p.Text = strings.Repeat("\n", p.TopPadding) + p.Text + strings.Repeat("\n", p.BottomPadding)

	ss := strings.Split(p.Text, "\n")
	for i, s2 := range ss {
		if i != len(ss) {
			if runewidth.StringWidth(color.ClearCode(s2)) < maxWidth {
				ss[i] = p.BoxStyle.Sprint(p.VerticalString) + strings.Repeat(" ", p.LeftPadding) + p.TextStyle.Sprint(s2) + strings.Repeat(" ", maxWidth-runewidth.StringWidth(color.ClearCode(s2))+p.RightPadding) + p.BoxStyle.Sprint(p.VerticalString)
			} else {
				ss[i] = p.BoxStyle.Sprint(p.VerticalString) + strings.Repeat(" ", p.LeftPadding) + p.TextStyle.Sprint(s2) + strings.Repeat(" ", p.RightPadding) + p.BoxStyle.Sprint(p.VerticalString)
			}
		}
		if i == len(ss)-1 {
			ss[i] += "\n"
		}
	}

	return topLine + "\n" + strings.Join(ss, "\n") + bottomLine, nil
}

// Render prints the Template to the terminal.
func (p BoxPrinter) Render() error {
	s, _ := p.Srender()
	Println(s)
	return nil
}
