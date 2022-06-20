package pterm

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

// PrintColor is false if PTerm should not print colored output.
var PrintColor = true

// EnableColor enables colors.
func EnableColor() {
	color.Enable = true
	PrintColor = true
}

// DisableColor disables colors.
func DisableColor() {
	color.Enable = false
	PrintColor = false
}

// Foreground colors. basic foreground colors 30 - 37.
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
	// FgDefault revert default FG.
	FgDefault Color = 39
)

// Extra foreground color 90 - 97.
const (
	FgDarkGray Color = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgLightWhite
	// FgGray is an alias of FgDarkGray.
	FgGray Color = 90
)

// Background colors. basic background colors 40 - 47.
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	// BgDefault reverts to the default background.
	BgDefault Color = 49
)

// Extra background color 100 - 107.
const (
	BgDarkGray Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgLightWhite
	// BgGray is an alias of BgDarkGray.
	BgGray Color = 100
)

// Option settings.
const (
	Reset Color = iota
	Bold
	Fuzzy
	Italic
	Underscore
	Blink
	FastBlink
	Reverse
	Concealed
	Strikethrough
)

var (
	// Red is an alias for FgRed.Sprint.
	Red = FgRed.Sprint
	// Cyan is an alias for FgCyan.Sprint.
	Cyan = FgCyan.Sprint
	// Gray is an alias for FgGray.Sprint.
	Gray = FgGray.Sprint
	// Blue is an alias for FgBlue.Sprint.
	Blue = FgBlue.Sprint
	// Black is an alias for FgBlack.Sprint.
	Black = FgBlack.Sprint
	// Green is an alias for FgGreen.Sprint.
	Green = FgGreen.Sprint
	// White is an alias for FgWhite.Sprint.
	White = FgWhite.Sprint
	// Yellow is an alias for FgYellow.Sprint.
	Yellow = FgYellow.Sprint
	// Magenta is an alias for FgMagenta.Sprint.
	Magenta = FgMagenta.Sprint

	// Normal is an alias for FgDefault.Sprint.
	Normal = FgDefault.Sprint

	// extra light.

	// LightRed is a shortcut for FgLightRed.Sprint.
	LightRed = FgLightRed.Sprint
	// LightCyan is a shortcut for FgLightCyan.Sprint.
	LightCyan = FgLightCyan.Sprint
	// LightBlue is a shortcut for FgLightBlue.Sprint.
	LightBlue = FgLightBlue.Sprint
	// LightGreen is a shortcut for FgLightGreen.Sprint.
	LightGreen = FgLightGreen.Sprint
	// LightWhite is a shortcut for FgLightWhite.Sprint.
	LightWhite = FgLightWhite.Sprint
	// LightYellow is a shortcut for FgLightYellow.Sprint.
	LightYellow = FgLightYellow.Sprint
	// LightMagenta is a shortcut for FgLightMagenta.Sprint.
	LightMagenta = FgLightMagenta.Sprint
)

// Color is a number which will be used to color strings in the terminal.
type Color uint8

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
// Input will be colored with the parent Color.
func (c Color) Sprintln(a ...interface{}) string {
	str := fmt.Sprintln(a...)
	return c.Sprint(str)
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
// Input will be colored with the parent Color.
func (c Color) Sprint(a ...interface{}) string {
	message := Sprint(a...)
	messageLines := strings.Split(message, "\n")
	for i, line := range messageLines {
		messageLines[i] = color.RenderCode(c.String(), strings.ReplaceAll(line, color.ResetSet, Sprintf("\x1b[0m\u001B[%sm", c.String())))
	}
	message = strings.Join(messageLines, "\n")
	return message
}

// Sprintf formats according to a format specifier and returns the resulting string.
// Input will be colored with the parent Color.
func (c Color) Sprintf(format string, a ...interface{}) string {
	return c.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
// Input will be colored with the parent Color.
func (c Color) Sprintfln(format string, a ...interface{}) string {
	return c.Sprint(Sprintf(format, a...) + "\n")
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color.
func (c Color) Println(a ...interface{}) *TextPrinter {
	Print(c.Sprintln(a...))
	tc := TextPrinter(c)
	return &tc
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color.
func (c Color) Print(a ...interface{}) *TextPrinter {
	Print(c.Sprint(a...))
	tc := TextPrinter(c)
	return &tc
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color.
func (c Color) Printf(format string, a ...interface{}) *TextPrinter {
	Print(c.Sprintf(format, a...))
	tc := TextPrinter(c)
	return &tc
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Color.
func (c Color) Printfln(format string, a ...interface{}) *TextPrinter {
	Print(c.Sprintfln(format, a...))
	tp := TextPrinter(c)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (c Color) PrintOnError(a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				c.Println(err)
			}
		}
	}

	tp := TextPrinter(c)
	return &tp
}

// PrintOnErrorf wraps every error which is not nil and prints it.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (c Color) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				c.Println(fmt.Errorf(format, err))
			}
		}
	}

	tp := TextPrinter(c)
	return &tp
}

// String converts the color to a string. eg "35".
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

// Style is a collection of colors.
// Can include foreground, background and styling (eg. Bold, Underscore, etc.) colors.
type Style []Color

// NewStyle returns a new Style.
// Accepts multiple colors.
func NewStyle(colors ...Color) *Style {
	ret := Style{}
	for _, c := range colors {
		ret = append(ret, c)
	}
	return &ret
}

// Add styles to the current Style.
func (s Style) Add(styles ...Style) Style {
	ret := s

	for _, st := range styles {
		ret = append(ret, st...)
	}

	return ret
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
// Input will be colored with the parent Style.
func (s Style) Sprint(a ...interface{}) string {
	message := Sprint(a...)
	messageLines := strings.Split(message, "\n")
	for i, line := range messageLines {
		messageLines[i] = color.RenderCode(s.String(), strings.ReplaceAll(line, color.ResetSet, Sprintf("\x1b[0m\u001B[%sm", s.String())))
	}
	message = strings.Join(messageLines, "\n")
	return color.RenderCode(s.String(), message)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
// Input will be colored with the parent Style.
func (s Style) Sprintln(a ...interface{}) string {
	return s.Sprint(a...) + "\n"
}

// Sprintf formats according to a format specifier and returns the resulting string.
// Input will be colored with the parent Style.
func (s Style) Sprintf(format string, a ...interface{}) string {
	return s.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
// Input will be colored with the parent Style.
func (s Style) Sprintfln(format string, a ...interface{}) string {
	return s.Sprint(Sprintf(format, a...) + "\n")
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style.
func (s Style) Print(a ...interface{}) {
	Print(s.Sprint(a...))
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style.
func (s Style) Println(a ...interface{}) {
	Println(s.Sprint(a...))
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style.
func (s Style) Printf(format string, a ...interface{}) {
	Print(s.Sprintf(format, a...))
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
// Input will be colored with the parent Style.
func (s Style) Printfln(format string, a ...interface{}) {
	Print(s.Sprintfln(format, a...))
}

// Code convert to code string. returns like "32;45;3".
func (s Style) Code() string {
	return s.String()
}

// String convert to code string. returns like "32;45;3".
func (s Style) String() string {
	return colors2code(s...)
}

// Converts colors to code.
// Return format: "32;45;3".
func colors2code(colors ...Color) string {
	if len(colors) == 0 {
		return ""
	}

	var codes []string
	for _, c := range colors {
		codes = append(codes, c.String())
	}

	return strings.Join(codes, ";")
}
