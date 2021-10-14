package pterm

import (
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// BoxPrinter is able to render a box around printables.
type BoxPrinter struct {
	Title                   string
	TitleTopLeft            bool
	TitleTopRight           bool
	TitleTopCenter          bool
	TitleBottomLeft         bool
	TitleBottomRight        bool
	TitleBottomCenter       bool
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

// DefaultBox is the default BoxPrinter.
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
	TopPadding:              0,
	BottomPadding:           0,
	TitleTopLeft:            true,
}

// WithTitle returns a new box with a specific Title.
func (p BoxPrinter) WithTitle(str string) *BoxPrinter {
	p.Title = str
	return &p
}

// WithTitleTopLeft returns a new box with a specific Title alignment.
func (p BoxPrinter) WithTitleTopLeft(b ...bool) *BoxPrinter {
	b2 := internal.WithBoolean(b)
	p.TitleTopLeft = b2
	p.TitleTopRight = false
	p.TitleTopCenter = false
	p.TitleBottomLeft = false
	p.TitleBottomRight = false
	p.TitleBottomCenter = false
	return &p
}

// WithTitleTopRight returns a new box with a specific Title alignment.
func (p BoxPrinter) WithTitleTopRight(b ...bool) *BoxPrinter {
	b2 := internal.WithBoolean(b)
	p.TitleTopLeft = false
	p.TitleTopRight = b2
	p.TitleTopCenter = false
	p.TitleBottomLeft = false
	p.TitleBottomRight = false
	p.TitleBottomCenter = false
	return &p
}

// WithTitleTopCenter returns a new box with a specific Title alignment.
func (p BoxPrinter) WithTitleTopCenter(b ...bool) *BoxPrinter {
	b2 := internal.WithBoolean(b)
	p.TitleTopLeft = false
	p.TitleTopRight = false
	p.TitleTopCenter = b2
	p.TitleBottomLeft = false
	p.TitleBottomRight = false
	p.TitleBottomCenter = false
	return &p
}

// WithTitleBottomLeft returns a new box with a specific Title alignment.
func (p BoxPrinter) WithTitleBottomLeft(b ...bool) *BoxPrinter {
	b2 := internal.WithBoolean(b)
	p.TitleTopLeft = false
	p.TitleTopRight = false
	p.TitleTopCenter = false
	p.TitleBottomLeft = b2
	p.TitleBottomRight = false
	p.TitleBottomCenter = false
	return &p
}

// WithTitleBottomRight returns a new box with a specific Title alignment.
func (p BoxPrinter) WithTitleBottomRight(b ...bool) *BoxPrinter {
	b2 := internal.WithBoolean(b)
	p.TitleTopLeft = false
	p.TitleTopRight = false
	p.TitleTopCenter = false
	p.TitleBottomLeft = false
	p.TitleBottomRight = b2
	p.TitleBottomCenter = false
	return &p
}

// WithTitleBottomCenter returns a new box with a specific Title alignment.
func (p BoxPrinter) WithTitleBottomCenter(b ...bool) *BoxPrinter {
	b2 := internal.WithBoolean(b)
	p.TitleTopLeft = false
	p.TitleTopRight = false
	p.TitleTopCenter = false
	p.TitleBottomLeft = false
	p.TitleBottomRight = false
	p.TitleBottomCenter = b2
	return &p
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

// WithHorizontalString returns a new box with a specific HorizontalString.
func (p BoxPrinter) WithHorizontalString(str string) *BoxPrinter {
	p.HorizontalString = str
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

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p BoxPrinter) Sprint(a ...interface{}) string {
	if p.BoxStyle == nil {
		p.BoxStyle = &ThemeDefault.BoxStyle
	}
	if p.TextStyle == nil {
		p.TextStyle = &ThemeDefault.BoxTextStyle
	}
	maxWidth := internal.GetStringMaxWidth(Sprint(a...))

	var topLine string
	var bottomLine string

	if p.Title == "" {
		topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
			maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
		bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
			maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.TopLeftCornerString)
	} else {
		p.Title = strings.ReplaceAll(p.Title, "\n", " ")
		if (maxWidth + p.RightPadding + p.LeftPadding - 4) < len(RemoveColorFromString(p.Title)) {
			p.RightPadding = len(RemoveColorFromString(p.Title)) - (maxWidth + p.RightPadding + p.LeftPadding - 5)
		}
		if p.TitleTopLeft {
			topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + internal.AddTitleToLine(p.Title, p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.LeftPadding+p.RightPadding, true) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
			bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
				maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.TopLeftCornerString)
		} else if p.TitleTopRight {
			topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + internal.AddTitleToLine(p.Title, p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.LeftPadding+p.RightPadding, false) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
			bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
				maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.TopLeftCornerString)
		} else if p.TitleTopCenter {
			topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + internal.AddTitleToLineCenter(p.Title, p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
			bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
				maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.TopLeftCornerString)
		} else if p.TitleBottomLeft {
			topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
				maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
			bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + internal.AddTitleToLine(p.Title, p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.LeftPadding+p.RightPadding, true) + p.BoxStyle.Sprint(p.TopLeftCornerString)
		} else if p.TitleBottomRight {
			topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
				maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
			bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + internal.AddTitleToLine(p.Title, p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.LeftPadding+p.RightPadding, false) + p.BoxStyle.Sprint(p.TopLeftCornerString)
		} else if p.TitleBottomCenter {
			topLine = p.BoxStyle.Sprint(p.BottomRightCornerString) + strings.Repeat(p.BoxStyle.Sprint(p.HorizontalString),
				maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.BottomLeftCornerString)
			bottomLine = p.BoxStyle.Sprint(p.TopRightCornerString) + internal.AddTitleToLineCenter(p.Title, p.BoxStyle.Sprint(p.HorizontalString), maxWidth+p.LeftPadding+p.RightPadding) + p.BoxStyle.Sprint(p.TopLeftCornerString)
		}
	}

	boxString := strings.Repeat("\n", p.TopPadding) + Sprint(a...) + strings.Repeat("\n", p.BottomPadding)

	ss := strings.Split(boxString, "\n")
	for i, s2 := range ss {
		if runewidth.StringWidth(RemoveColorFromString(s2)) < maxWidth {
			ss[i] = p.BoxStyle.Sprint(p.VerticalString) + strings.Repeat(" ", p.LeftPadding) + p.TextStyle.Sprint(s2) +
				strings.Repeat(" ", maxWidth-runewidth.StringWidth(RemoveColorFromString(s2))+p.RightPadding) +
				p.BoxStyle.Sprint(p.VerticalString)
		} else {
			ss[i] = p.BoxStyle.Sprint(p.VerticalString) + strings.Repeat(" ", p.LeftPadding) + p.TextStyle.Sprint(s2) +
				strings.Repeat(" ", p.RightPadding) + p.BoxStyle.Sprint(p.VerticalString)
		}
	}
	return topLine + "\n" + strings.Join(ss, "\n") + "\n" + bottomLine
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p BoxPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(strings.TrimSuffix(Sprintln(a...), "\n")) + "\n"
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p BoxPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p BoxPrinter) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p BoxPrinter) Print(a ...interface{}) *TextPrinter {
	Print(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p BoxPrinter) Println(a ...interface{}) *TextPrinter {
	Print(p.Sprintln(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p BoxPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p BoxPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p BoxPrinter) PrintOnError(a ...interface{}) *TextPrinter {
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
func (p BoxPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
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
