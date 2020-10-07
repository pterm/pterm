package pterm

import "github.com/pterm/pterm/internal"

var (
	// DefaultHeader returns the printer for a default header text.
	// Defaults to LightWhite, Bold Text and a Gray DefaultHeader background.
	DefaultHeader = HeaderPrinter{
		TextStyle:       Style{FgLightWhite, Bold},
		BackgroundStyle: Style{BgGray},
		Margin:          5,
	}
)

// HeaderPrinter contains the data used to craft a header.
// A header is printed as a big box with text in it.
// Can be used as title screens or section separator.
type HeaderPrinter struct {
	TextStyle       Style
	BackgroundStyle Style
	Margin          int
	FullWidth       bool
}

// WithTextStyle returns a new HeaderPrinter with changed
func (p HeaderPrinter) WithTextStyle(style Style) *HeaderPrinter {
	p.TextStyle = style
	return &p
}

// WithBackgroundStyle changes the background styling of the header.
func (p HeaderPrinter) WithBackgroundStyle(style Style) *HeaderPrinter {
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
	text := Sprint(a...)

	if p.FullWidth {
		p.Margin = (GetTerminalWidth() - len(text)) / 2
	}

	renderedTextLength := len(text) + p.Margin*2

	var marginString string
	for i := 0; i < p.Margin; i++ {
		marginString += " "
	}
	var blankLine string
	for i := 0; i < renderedTextLength; i++ {
		blankLine += " "
	}

	var ret string

	ret += p.BackgroundStyle.Sprint(blankLine) + "\n"
	ret += p.BackgroundStyle.Sprint(p.TextStyle.Sprint(marginString+text+marginString)) + "\n"
	ret += p.BackgroundStyle.Sprint(blankLine) + "\n"

	return ret
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p HeaderPrinter) Sprintln(a ...interface{}) string {
	return Sprint(p.Sprint(a...) + "\n")
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p HeaderPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p HeaderPrinter) Print(a ...interface{}) GenericPrinter {
	Print(p.Sprint(a...))
	return &p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p HeaderPrinter) Println(a ...interface{}) GenericPrinter {
	Println(p.Sprint(a...))
	return &p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p HeaderPrinter) Printf(format string, a ...interface{}) GenericPrinter {
	Print(p.Sprintf(format, a...))
	return &p
}
