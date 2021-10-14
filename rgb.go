package pterm

import (
	"fmt"
	"strconv"
	"strings"

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
func (p RGB) GetValues() (r, g, b uint8) {
	return p.R, p.G, p.B
}

// NewRGB returns a new RGB.
func NewRGB(r, g, b uint8) RGB {
	return RGB{R: r, G: g, B: b}
}

// NewRGBFromHEX converts a HEX and returns a new RGB.
func NewRGBFromHEX(hex string) (RGB, error) {
	hex = strings.ToLower(hex)
	hex = strings.ReplaceAll(hex, "#", "")
	hex = strings.ReplaceAll(hex, "0x", "")

	if len(hex) == 3 {
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return RGB{}, ErrHexCodeIsInvalid
	}

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

// Fade fades one RGB value (over other RGB values) to another RGB value, by giving the function a minimum, maximum and current value.
func (p RGB) Fade(min, max, current float32, end ...RGB) RGB {
	if min < 0 {
		max -= min
		current -= min
		min = 0
	}
	if len(end) == 1 {
		return RGB{
			R: uint8(internal.MapRangeToRange(min, max, float32(p.R), float32(end[0].R), current)),
			G: uint8(internal.MapRangeToRange(min, max, float32(p.G), float32(end[0].G), current)),
			B: uint8(internal.MapRangeToRange(min, max, float32(p.B), float32(end[0].B), current)),
		}
	} else if len(end) > 1 {
		f := (max - min) / float32(len(end))
		tempCurrent := current
		if f > current {
			return p.Fade(min, f, current, end[0])
		} else {
			for i := 0; i < len(end)-1; i++ {
				tempCurrent -= f
				if f > tempCurrent {
					return end[i].Fade(min, min+f, tempCurrent, end[i+1])
				}
			}
		}
	}
	return p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p RGB) Sprint(a ...interface{}) string {
	return color.RGB(p.R, p.G, p.B).Sprint(a...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p RGB) Sprintln(a ...interface{}) string {
	return p.Sprint(Sprintln(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p RGB) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p RGB) Sprintfln(format string, a ...interface{}) string {
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p RGB) Print(a ...interface{}) *TextPrinter {
	Print(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p RGB) Println(a ...interface{}) *TextPrinter {
	Print(p.Sprintln(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p RGB) Printf(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p RGB) Printfln(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintfln(format, a...))
	tp := TextPrinter(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p RGB) PrintOnError(a ...interface{}) *TextPrinter {
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
func (p RGB) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
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
