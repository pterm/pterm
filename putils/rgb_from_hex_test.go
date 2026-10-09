package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestRGBFromHEX(t *testing.T) {
	tests := []struct {
		name     string
		hex      string
		expected pterm.RGB
	}{
		{name: "Six digits with hash", hex: "#ff0000", expected: pterm.RGB{R: 255, G: 0, B: 0}},
		{name: "Six digits without prefix", hex: "00ff00", expected: pterm.RGB{R: 0, G: 255, B: 0}},
		{name: "Six digits with 0x prefix", hex: "0x0000ff", expected: pterm.RGB{R: 0, G: 0, B: 255}},
		{name: "Uppercase digits", hex: "#FFAA00", expected: pterm.RGB{R: 255, G: 170, B: 0}},
		{name: "Mixed case digits", hex: "#FfAa0f", expected: pterm.RGB{R: 255, G: 170, B: 15}},
		{name: "Black", hex: "#000000", expected: pterm.RGB{R: 0, G: 0, B: 0}},
		{name: "White", hex: "#ffffff", expected: pterm.RGB{R: 255, G: 255, B: 255}},
		{name: "Distinct channels", hex: "#010203", expected: pterm.RGB{R: 1, G: 2, B: 3}},
		{name: "Shorthand with hash", hex: "#fba", expected: pterm.RGB{R: 255, G: 187, B: 170}},
		{name: "Shorthand without prefix", hex: "fba", expected: pterm.RGB{R: 255, G: 187, B: 170}},
		{name: "Shorthand with 0x prefix", hex: "0xfba", expected: pterm.RGB{R: 255, G: 187, B: 170}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rgb, err := RGBFromHEX(tt.hex)

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, rgb)
		})
	}
}

func TestRGBFromHEXInvalidLength(t *testing.T) {
	for _, hex := range []string{"", "#", "f", "ff", "ffff", "fffff", "fffffff", "#ff00", "0xff"} {
		t.Run(hex, func(t *testing.T) {
			rgb, err := RGBFromHEX(hex)

			assert.ErrorIs(t, err, pterm.ErrHexCodeIsInvalid)
			assert.Equal(t, pterm.RGB{}, rgb)
		})
	}
}

func TestRGBFromHEXInvalidDigits(t *testing.T) {
	for _, hex := range []string{"zzzzzz", "#gggggg", "12345x"} {
		t.Run(hex, func(t *testing.T) {
			rgb, err := RGBFromHEX(hex)

			assert.Error(t, err)
			assert.Equal(t, pterm.RGB{}, rgb)
		})
	}
}
