package color_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal/color"
)

func TestRGBTo256(t *testing.T) {
	tests := []struct {
		name    string
		r, g, b uint8
		want    uint8
	}{
		{name: "black is the cube corner", r: 0, g: 0, b: 0, want: 16},
		{name: "white is the cube corner", r: 255, g: 255, b: 255, want: 231},
		{name: "pure red", r: 255, g: 0, b: 0, want: 196},
		{name: "pure green", r: 0, g: 255, b: 0, want: 46},
		{name: "pure blue", r: 0, g: 0, b: 255, want: 21},
		{name: "exact cube color", r: 95, g: 135, b: 175, want: 67},
		{name: "middle gray uses the grayscale ramp", r: 128, g: 128, b: 128, want: 244},
		{name: "gray between cube steps uses the grayscale ramp", r: 100, g: 100, b: 100, want: 241},
		{name: "darkest ramp gray", r: 8, g: 8, b: 8, want: 232},
		{name: "lightest ramp gray", r: 238, g: 238, b: 238, want: 255},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, color.RGBTo256(test.r, test.g, test.b))
		})
	}
}

func TestRGBToBasic(t *testing.T) {
	tests := []struct {
		name    string
		r, g, b uint8
		want    uint8
	}{
		{name: "black", r: 0, g: 0, b: 0, want: 30},
		{name: "xterm red", r: 205, g: 0, b: 0, want: 31},
		{name: "xterm white", r: 229, g: 229, b: 229, want: 37},
		{name: "middle gray is bright black", r: 128, g: 128, b: 128, want: 90},
		{name: "pure red is bright red", r: 255, g: 0, b: 0, want: 91},
		{name: "pure magenta is bright magenta", r: 255, g: 0, b: 255, want: 95},
		{name: "pure white is bright white", r: 255, g: 255, b: 255, want: 97},
		{name: "xterm blue", r: 0, g: 0, b: 238, want: 34},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, color.RGBToBasic(test.r, test.g, test.b))
		})
	}
}

func TestLevelForegroundRGB(t *testing.T) {
	assert.Equal(t, "\x1b[38;2;255;0;0m", color.LevelTrueColor.ForegroundRGB(255, 0, 0))
	assert.Equal(t, "\x1b[38;5;196m", color.Level256.ForegroundRGB(255, 0, 0))
	assert.Equal(t, "\x1b[91m", color.LevelBasic.ForegroundRGB(255, 0, 0))
	assert.Equal(t, "", color.LevelNone.ForegroundRGB(255, 0, 0))
}

func TestLevelBackgroundRGB(t *testing.T) {
	assert.Equal(t, "\x1b[48;2;255;0;0m", color.LevelTrueColor.BackgroundRGB(255, 0, 0))
	assert.Equal(t, "\x1b[48;5;196m", color.Level256.BackgroundRGB(255, 0, 0))
	assert.Equal(t, "\x1b[101m", color.LevelBasic.BackgroundRGB(255, 0, 0))
	assert.Equal(t, "", color.LevelNone.BackgroundRGB(255, 0, 0))
}

func TestLevelString(t *testing.T) {
	assert.Equal(t, "none", color.LevelNone.String())
	assert.Equal(t, "16", color.LevelBasic.String())
	assert.Equal(t, "256", color.Level256.String())
	assert.Equal(t, "truecolor", color.LevelTrueColor.String())
}

// Every color of the 6x6x6 cube must map back to its own palette index.
func TestRGBTo256CubeIdentity(t *testing.T) {
	steps := []uint8{0, 95, 135, 175, 215, 255}

	for ri, r := range steps {
		for gi, g := range steps {
			for bi, b := range steps {
				want := uint8(16 + 36*ri + 6*gi + bi) //nolint:gosec // 16 + 36*5 + 6*5 + 5 = 231 always fits.

				assert.Equal(t, want, color.RGBTo256(r, g, b), "RGBTo256(%d, %d, %d)", r, g, b)
			}
		}
	}
}

// Every gray of the grayscale ramp must map back to its own palette index.
func TestRGBTo256GrayRampIdentity(t *testing.T) {
	for i := range 24 {
		gray := uint8(8 + 10*i)
		want := uint8(232 + i)

		assert.Equal(t, want, color.RGBTo256(gray, gray, gray), "RGBTo256(%d, %d, %d)", gray, gray, gray)
	}
}

// Every base color of the xterm default palette must map back to its own SGR
// code.
func TestRGBToBasicPaletteIdentity(t *testing.T) {
	palette := [16][3]uint8{
		{0, 0, 0},       // 30: black
		{205, 0, 0},     // 31: red
		{0, 205, 0},     // 32: green
		{205, 205, 0},   // 33: yellow
		{0, 0, 238},     // 34: blue
		{205, 0, 205},   // 35: magenta
		{0, 205, 205},   // 36: cyan
		{229, 229, 229}, // 37: white
		{127, 127, 127}, // 90: bright black
		{255, 0, 0},     // 91: bright red
		{0, 255, 0},     // 92: bright green
		{255, 255, 0},   // 93: bright yellow
		{92, 92, 255},   // 94: bright blue
		{255, 0, 255},   // 95: bright magenta
		{0, 255, 255},   // 96: bright cyan
		{255, 255, 255}, // 97: bright white
	}

	for i, rgb := range palette {
		want := uint8(30 + i)
		if i >= 8 {
			want = uint8(90 + i - 8)
		}

		assert.Equal(t, want, color.RGBToBasic(rgb[0], rgb[1], rgb[2]), "RGBToBasic(%d, %d, %d)", rgb[0], rgb[1], rgb[2])
	}
}

// RGBTo256 must always return a color index (16-255), never one of the 16
// base colors, so the result renders identically in every terminal.
func TestRGBTo256NeverReturnsBaseColors(t *testing.T) {
	for _, c := range [][3]uint8{{0, 0, 0}, {1, 1, 1}, {50, 50, 50}, {255, 255, 255}, {12, 34, 56}, {200, 100, 0}} {
		assert.GreaterOrEqual(t, color.RGBTo256(c[0], c[1], c[2]), uint8(16), "RGBTo256(%d, %d, %d)", c[0], c[1], c[2])
	}
}

func TestLevelStringUnknown(t *testing.T) {
	assert.Equal(t, "unknown", color.Level(99).String())
}
