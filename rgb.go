package pterm

import "github.com/pterm/pterm/internal"

type RGB struct {
	R uint8
	G uint8
	B uint8
}

func NewRGB(r, g, b uint8) RGB {
	return RGB{R: r, G: g, B: b}
}

func (rgb RGB) Fade(min, max, current float32, end RGB) RGB {
	internal.MapRangeToRange(min, max, float32(rgb.B), float32(end.B), current)
	return RGB{
		R: uint8(internal.MapRangeToRange(min, max, float32(rgb.R), float32(end.R), current)),
		G: uint8(internal.MapRangeToRange(min, max, float32(rgb.G), float32(end.G), current)),
		B: uint8(internal.MapRangeToRange(min, max, float32(rgb.B), float32(end.B), current)),
	}
}
