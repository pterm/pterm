package pterm

import (
	"strconv"

	"github.com/gookit/color"

	"github.com/pterm/pterm/internal"
)

// RGB color model is an additive color model in which red, green, and blue light are added together in various ways to reproduce a broad array of colors.
// The name of the model comes from the initials of the three additive primary colors, red, green, and blue.
// https://en.wikipedia.org/wiki/RGB_color_model
type RGB struct {
	R uint8
	G uint8
	B uint8
}

// GetValues returns the RGB values separately.
func (rgb RGB) GetValues() (r, g, b uint8) {
	return rgb.R, rgb.G, rgb.B
}

// NewRGB returns a new RGB.
func NewRGB(r, g, b uint8) RGB {
	return RGB{R: r, G: g, B: b}
}

// NewRGBFromHEX converts a HEX and returns a new RGB.
func NewRGBFromHEX(hex string) (RGB, error) {
	i64, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return RGB{}, err
	}
	c := int(i64)

	return RGB{
		R: uint8(c >> 16),
		G: uint8((c & 0x00FF00) >> 8),
		B: uint8(c & 0x0000FF),
	}, nil
}

// Fade fades one RGB value to another RGB value, by giving the function a minimum, maximum and current value.
func (rgb RGB) Fade(min, max, current float32, end RGB) RGB {
	return RGB{
		R: uint8(internal.MapRangeToRange(min, max, float32(rgb.R), float32(end.R), current)),
		G: uint8(internal.MapRangeToRange(min, max, float32(rgb.G), float32(end.G), current)),
		B: uint8(internal.MapRangeToRange(min, max, float32(rgb.B), float32(end.B), current)),
	}
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (rgb RGB) Sprint(a ...interface{}) string {
	return color.RGB(rgb.R, rgb.G, rgb.B).Sprint(a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (rgb RGB) Sprintln(a ...interface{}) string {
	return Sprintln(rgb.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (rgb RGB) Sprintf(format string, a ...interface{}) string {
	return rgb.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (rgb RGB) Print(a ...interface{}) TextPrinter {
	Print(rgb.Sprint(a...))
	return &rgb
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (rgb RGB) Println(a ...interface{}) TextPrinter {
	Println(rgb.Sprint(a...))
	return &rgb
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (rgb RGB) Printf(format string, a ...interface{}) TextPrinter {
	Print(rgb.Sprintf(format, a...))
	return &rgb
}
