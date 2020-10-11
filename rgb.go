package pterm

import "github.com/pterm/pterm/internal"

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

// Fade fades one RGB value to another RGB value, by giving the function a minimum, maximum and current value.
func (rgb RGB) Fade(min, max, current float32, end RGB) RGB {
	return RGB{
		R: uint8(internal.MapRangeToRange(min, max, float32(rgb.R), float32(end.R), current)),
		G: uint8(internal.MapRangeToRange(min, max, float32(rgb.G), float32(end.G), current)),
		B: uint8(internal.MapRangeToRange(min, max, float32(rgb.B), float32(end.B), current)),
	}
}
