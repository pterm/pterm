package pterm_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestNewRGB(t *testing.T) {
	assert.Equal(t, pterm.RGB{R: 0, G: 0, B: 0}, pterm.NewRGB(0, 0, 0))
	assert.Equal(t, pterm.RGB{R: 255, G: 255, B: 255}, pterm.NewRGB(255, 255, 255))
	assert.Equal(t, pterm.RGB{R: 1, G: 2, B: 3}, pterm.NewRGB(1, 2, 3))
	assert.Equal(t, pterm.RGB{R: 1, G: 2, B: 3, Background: true}, pterm.NewRGB(1, 2, 3, true))
	assert.Equal(t, pterm.RGB{R: 1, G: 2, B: 3, Background: false}, pterm.NewRGB(1, 2, 3, false))
}

func TestNewRGBFromHEX(t *testing.T) {
	tests := []struct {
		hex  string
		want pterm.RGB
	}{
		{hex: "#ff0009", want: pterm.RGB{R: 255, G: 0, B: 9}},
		{hex: "ff0009", want: pterm.RGB{R: 255, G: 0, B: 9}},
		{hex: "ff00090x", want: pterm.RGB{R: 255, G: 0, B: 9}},
		{hex: "ff00090X", want: pterm.RGB{R: 255, G: 0, B: 9}},
		{hex: "#fba", want: pterm.RGB{R: 255, G: 187, B: 170}},
		{hex: "fba", want: pterm.RGB{R: 255, G: 187, B: 170}},
		{hex: "fba0x", want: pterm.RGB{R: 255, G: 187, B: 170}},
	}
	for _, test := range tests {
		t.Run(test.hex, func(t *testing.T) {
			rgb, err := pterm.NewRGBFromHEX(test.hex)
			assert.Equal(t, test.want, rgb)
			assert.NoError(t, err)
		})
	}

	testsFail := []struct {
		hex  string
		want error
	}{
		{hex: "faba0x", want: pterm.ErrHexCodeIsInvalid},
		{hex: "faba", want: pterm.ErrHexCodeIsInvalid},
		{hex: "#faba", want: pterm.ErrHexCodeIsInvalid},
		{hex: "fax", want: strconv.ErrSyntax},
	}
	for _, test := range testsFail {
		t.Run(test.hex, func(t *testing.T) {
			_, err := pterm.NewRGBFromHEX(test.hex)
			assert.True(t, errors.Is(err, test.want))
		})
	}
}

func TestRGB_GetValues(t *testing.T) {
	r, g, b := pterm.RGB{R: 1, G: 2, B: 3}.GetValues()

	assert.Equal(t, uint8(1), r)
	assert.Equal(t, uint8(2), g)
	assert.Equal(t, uint8(3), b)
}

// Fade interpolation math: the exact RGB values at fraction points.

func TestRGB_FadeEndpoints(t *testing.T) {
	from := pterm.RGB{R: 10, G: 20, B: 30}
	to := pterm.RGB{R: 210, G: 220, B: 230}

	assert.Equal(t, from, from.Fade(0, 100, 0, to), "at the minimum the fade returns the start color")
	assert.Equal(t, to, from.Fade(0, 100, 100, to), "at the maximum the fade returns the end color")
}

func TestRGB_FadeMidpointAverages(t *testing.T) {
	black := pterm.RGB{}
	white := pterm.RGB{R: 255, G: 255, B: 255}

	assert.Equal(t, pterm.RGB{R: 127, G: 127, B: 127}, black.Fade(0, 100, 50, white))
}

func TestRGB_FadePreservesBackgroundFlag(t *testing.T) {
	from := pterm.RGB{R: 0, G: 0, B: 0, Background: true}

	assert.True(t, from.Fade(0, 100, 50, pterm.RGB{R: 255, G: 255, B: 255}).Background)
}

func TestRGB_Fade(t *testing.T) {
	type args struct {
		min     float32
		max     float32
		current float32
		end     []pterm.RGB
	}

	tests := []struct {
		name  string
		start pterm.RGB
		args  args
		want  pterm.RGB
	}{
		{name: "Middle", start: pterm.RGB{}, args: args{min: 0, max: 100, current: 50, end: []pterm.RGB{{R: 255, G: 255, B: 255}}}, want: pterm.RGB{R: 127, G: 127, B: 127}},
		{name: "ZeroToZero", start: pterm.RGB{}, args: args{min: 0, max: 100, current: 50, end: []pterm.RGB{{}}}, want: pterm.RGB{}},
		{name: "SameStartAndEnd", start: pterm.RGB{R: 0, G: 1, B: 2}, args: args{min: 0, max: 100, current: 50, end: []pterm.RGB{{R: 0, G: 1, B: 2}}}, want: pterm.RGB{R: 0, G: 1, B: 2}},
		{name: "NegativeRangeMiddle", start: pterm.RGB{}, args: args{min: -50, max: 50, current: 0, end: []pterm.RGB{{R: 255, G: 255, B: 255}}}, want: pterm.RGB{R: 127, G: 127, B: 127}},
		{name: "NegativeRangeMiddleMultipleRGB", start: pterm.RGB{}, args: args{min: -50, max: 50, current: 0, end: []pterm.RGB{{R: 127, G: 127, B: 127}, {R: 255, G: 255, B: 255}}}, want: pterm.RGB{R: 127, G: 127, B: 127}},
		{name: "MiddleMultipleRGB", start: pterm.RGB{}, args: args{min: 0, max: 100, current: 50, end: []pterm.RGB{{R: 127, G: 127, B: 127}, {R: 255, G: 255, B: 255}}}, want: pterm.RGB{R: 127, G: 127, B: 127}},
		{name: "QuarterTwoStops", start: pterm.RGB{}, args: args{min: 0, max: 100, current: 25, end: []pterm.RGB{{R: 255, G: 255, B: 255}, {R: 255, G: 255, B: 255}}}, want: pterm.RGB{R: 127, G: 127, B: 127}},
		{name: "ThreeQuartersTwoStops", start: pterm.RGB{}, args: args{min: 0, max: 100, current: 75, end: []pterm.RGB{{R: 255, G: 0, B: 0}, {R: 0, G: 0, B: 255}}}, want: pterm.RGB{R: 127, G: 0, B: 127}},
		{name: "MaxMultipleRGB", start: pterm.RGB{}, args: args{min: 0, max: 100, current: 100, end: []pterm.RGB{{R: 127, G: 127, B: 127}, {R: 9, G: 8, B: 7}}}, want: pterm.RGB{R: 9, G: 8, B: 7}},
		{name: "MiddleMultipleRGBPositiveMin", start: pterm.RGB{}, args: args{min: 10, max: 110, current: 60, end: []pterm.RGB{{R: 127, G: 127, B: 127}, {R: 255, G: 255, B: 255}}}, want: pterm.RGB{R: 127, G: 127, B: 127}},
		{name: "MiddleNoRGB", start: pterm.RGB{}, args: args{min: 10, max: 110, current: 60, end: []pterm.RGB{}}, want: pterm.RGB{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.start.Fade(tt.args.min, tt.args.max, tt.args.current, tt.args.end...))
		})
	}
}

// Exact TrueColor escape sequences. The test environment forces true color
// (COLORTERM=truecolor in TestMain, or CI), so RGB printers render at full
// depth here; degradation to lower levels is covered further down.

func TestRGBSprintEmitsExactTrueColorSequence(t *testing.T) {
	assert.Equal(t, "\x1b[38;2;1;2;3mx\x1b[0m", pterm.NewRGB(1, 2, 3).Sprint("x"))
	assert.Equal(t, "\x1b[38;2;255;0;255mx\x1b[0m", pterm.NewRGB(255, 0, 255).Sprint("x"))
}

func TestRGBBackgroundSprintClearsToEndOfLine(t *testing.T) {
	// Background colors additionally clear to the end of the line so the
	// color fills the rest of the row.
	assert.Equal(t, "\x1b[48;2;1;2;3mx\x1b[0m\x1b[K", pterm.NewRGB(1, 2, 3, true).Sprint("x"))
}

func TestRGBSprintFamilyExactOutput(t *testing.T) {
	rgb := pterm.NewRGB(1, 2, 3)

	assert.Equal(t, "\x1b[38;2;1;2;3mn=1\x1b[0m", rgb.Sprintf("n=%d", 1))
	assert.Equal(t, "\x1b[38;2;1;2;3mn=1\x1b[0m\n", rgb.Sprintfln("n=%d", 1))
	assert.Equal(t, "\x1b[38;2;1;2;3mx\n\x1b[0m", rgb.Sprintln("x"), "Sprintln wraps the newline inside the sequence")
	assert.Equal(t, "", rgb.Sprint(""), "empty input must not emit stray sequences")
}

func TestNewRGBStyle(t *testing.T) {
	fg := pterm.RGB{R: 1, G: 2, B: 3}
	bg := pterm.RGB{R: 4, G: 5, B: 6}

	onlyFg := pterm.NewRGBStyle(fg)
	assert.Equal(t, fg, onlyFg.Foreground)
	assert.Equal(t, "\x1b[38;2;1;2;3mx\x1b[0m", onlyFg.Sprint("x"), "without a background only the foreground sequence is emitted")

	withBg := pterm.NewRGBStyle(fg, bg)
	assert.Equal(t, fg, withBg.Foreground)
	assert.Equal(t, bg, withBg.Background)
	assert.Equal(t, "\x1b[38;2;1;2;3m\x1b[48;2;4;5;6mx\x1b[0m", withBg.Sprint("x"))
}

func TestRGBStyleSprintCombinesForegroundBackgroundAndOptions(t *testing.T) {
	style := pterm.NewRGBStyle(pterm.RGB{R: 1, G: 2, B: 3}, pterm.RGB{R: 4, G: 5, B: 6}).AddOptions(pterm.Bold, pterm.Italic)

	expected := "\x1b[38;2;1;2;3m\x1b[48;2;4;5;6m\x1b[1m\x1b[3mx\x1b[0m"
	assert.Equal(t, expected, style.Sprint("x"))
	assert.Equal(t, expected, style.Sprintf("%s", "x"))
	assert.Equal(t, expected+"\n", style.Sprintln("x"))
	assert.Equal(t, expected+"\n", style.Sprintfln("%s", "x"))
	assert.Equal(t, "", style.Sprint(""), "empty input must not emit stray sequences")
}

func TestRGBStyleAddOptionsDoesNotMutateOriginal(t *testing.T) {
	original := pterm.NewRGBStyle(pterm.RGB{R: 1, G: 2, B: 3})
	modified := original.AddOptions(pterm.Bold)

	assert.Empty(t, original.Options)
	assert.Equal(t, []pterm.Color{pterm.Bold}, modified.Options)
}

func TestRGBToRGBStyleRespectsBackgroundFlag(t *testing.T) {
	fg := pterm.NewRGB(1, 2, 3)
	bg := pterm.NewRGB(1, 2, 3, true)

	assert.Equal(t, "\x1b[38;2;1;2;3mx\x1b[0m", fg.ToRGBStyle().Sprint("x"))

	// A background RGB becomes the style's background; without a foreground
	// the style still emits the zero-value foreground sequence.
	assert.Equal(t, pterm.RGB{R: 1, G: 2, B: 3, Background: true}, bg.ToRGBStyle().Background)
	assert.Equal(t, pterm.RGB{}, bg.ToRGBStyle().Foreground)
}

// Disabled colors: RGB printers must return plain text and strip escape codes
// already embedded in the input.

func TestRGBSprintWithDisabledColors(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableColor()

	assert.Equal(t, "plain", pterm.NewRGB(255, 0, 0).Sprint("plain"))
	assert.Equal(t, "plain", pterm.NewRGB(255, 0, 0, true).Sprint("plain"))
	assert.Equal(t, "green text", pterm.NewRGB(255, 0, 0).Sprint("\x1b[32mgreen\x1b[0m text"))
	assert.Equal(t, "plain", pterm.NewRGBStyle(pterm.RGB{R: 255}, pterm.RGB{B: 255}).AddOptions(pterm.Bold).Sprint("plain"))
}

// The RGB printers degrade to the closest color the terminal can actually
// render: 256-color palette entries on 256-color terminals, the 16 base
// colors everywhere else. True color terminals and CI systems (checked via
// snapshot tests) get the full 24-bit sequences.
func TestRGBSprintDegradesToTerminalColorLevel(t *testing.T) {
	// Scrub every variable the detection looks at, including CI, which
	// forces true color rendering.
	for _, name := range []string{"CI", "NO_COLOR", "FORCE_COLOR", "CLICOLOR", "CLICOLOR_FORCE", "COLORTERM", "WT_SESSION", "ConEmuANSI", "TERMINAL_EMULATOR", "TERM_PROGRAM", "ANSICON"} {
		t.Setenv(name, "")
	}

	t.Setenv("TERM", "xterm-256color")
	assert.Equal(t, "\x1b[38;5;196mred\x1b[0m", pterm.NewRGB(255, 0, 0).Sprint("red"))
	assert.Equal(t, "\x1b[48;5;196mred\x1b[0m\x1b[K", pterm.NewRGB(255, 0, 0, true).Sprint("red"))
	assert.Equal(t, "\x1b[38;5;196m\x1b[48;5;16m\x1b[1mred\x1b[0m", pterm.NewRGBStyle(pterm.RGB{R: 255}, pterm.RGB{}).AddOptions(pterm.Bold).Sprint("red"))

	t.Setenv("TERM", "xterm")
	assert.Equal(t, "\x1b[91mred\x1b[0m", pterm.NewRGB(255, 0, 0).Sprint("red"))
	assert.Equal(t, "\x1b[101mred\x1b[0m\x1b[K", pterm.NewRGB(255, 0, 0, true).Sprint("red"))
}
