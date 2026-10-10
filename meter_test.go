package pterm

import (
	"strings"
	"testing"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/text"
)

func testCanvas(g Glyphs) canvas {
	return canvas{level: color.LevelNone, theme: ThemeMacchiato, g: g, width: 80}
}

func TestDrawMeterIsAlwaysExactlyWide(t *testing.T) {
	for _, g := range []Glyphs{GlyphsUnicode, GlyphsASCII} {
		c := testCanvas(g)

		for width := 1; width <= 12; width++ {
			for i := -8; i <= 120; i++ {
				ratio := float64(i) / 100

				if got := text.Width(drawMeter(c, ratio, width, ToneSuccess, ToneBorder)); got != width {
					t.Fatalf("drawMeter(%v, %d) is %d columns wide", ratio, width, got)
				}
			}
		}
	}
}

func TestDrawMeterEighths(t *testing.T) {
	c := testCanvas(GlyphsUnicode)

	cases := map[float64]string{
		0:      "░░░░",
		0.25:   "█░░░",
		0.5:    "██░░",
		1.0:    "████",
		1.5:    "████",
		0.3125: "█▎░░", // 1.25 columns
		0.03:   "▏░░░",
	}

	for ratio, want := range cases {
		if got := drawMeter(c, ratio, 4, ToneSuccess, ToneBorder); got != want {
			t.Errorf("drawMeter(%v) = %q, want %q", ratio, got, want)
		}
	}
}

func TestDrawMeterASCIIRoundsToColumns(t *testing.T) {
	c := testCanvas(GlyphsASCII)

	if got := drawMeter(c, 0.3125, 4, ToneSuccess, ToneBorder); got != "#..." {
		t.Errorf("got %q", got)
	}

	if got := drawMeter(c, 0.375, 4, ToneSuccess, ToneBorder); got != "##.." {
		t.Errorf("got %q", got)
	}
}

func TestDrawStack(t *testing.T) {
	c := testCanvas(GlyphsUnicode)
	segs := []segment{{2.4, ToneSuccess}, {1.3, ToneWarning}, {0.5, ToneError}}

	got := drawStack(c, segs, 6, "", ToneBorder)
	if want := "████▎ "; got != want {
		t.Errorf("drawStack = %q, want %q", got, want)
	}

	for width := 1; width <= 8; width++ {
		if w := text.Width(drawStack(c, segs, width, "░", ToneBorder)); w != width {
			t.Errorf("width %d: drawStack is %d wide", width, w)
		}

		if w := text.Width(drawStackLeft(c, segs, width, "░", ToneBorder)); w != width {
			t.Errorf("width %d: drawStackLeft is %d wide", width, w)
		}
	}
}

func TestDrawStackLeft(t *testing.T) {
	c := testCanvas(GlyphsUnicode)

	cases := map[float64]string{
		0:   "░░░░",
		1:   "░░░█",
		1.5: "░░▐█",
		2.2: "░░██",
		2.8: "░███",
		9:   "████",
	}

	for length, want := range cases {
		if got := drawStackLeft(c, []segment{{length, ToneError}}, 4, "░", ToneBorder); got != want {
			t.Errorf("drawStackLeft(%v) = %q, want %q", length, got, want)
		}
	}
}

func TestColumns(t *testing.T) {
	c := testCanvas(GlyphsUnicode)

	up := columnUp(c, []segment{{1.5, ToneSuccess}, {1, ToneError}}, 4)

	glyphs := make([]string, len(up))
	for i, cl := range up {
		glyphs[i] = cl.glyph
	}

	if got := strings.Join(glyphs, ""); got != "██▄" {
		t.Errorf("columnUp = %q", got)
	}

	if up[0].tone != ToneSuccess || up[2].tone != ToneError {
		t.Errorf("tones = %v %v", up[0].tone, up[2].tone)
	}

	down := columnDown(c, []segment{{1.5, ToneError}}, 4)
	if len(down) != 2 || down[1].glyph != "▀" {
		t.Errorf("columnDown = %v", down)
	}

	if got := columnUp(c, []segment{{9, ToneError}}, 3); len(got) != 3 {
		t.Errorf("columnUp is %d tall, want 3", len(got))
	}
}

func TestThresholdTone(t *testing.T) {
	ts := LoadThresholds()

	cases := map[float64]Tone{0: ToneSuccess, 0.69: ToneSuccess, 0.7: ToneWarning, 0.95: ToneError, 2: ToneError, -1: TonePrimary}
	for ratio, want := range cases {
		if got := thresholdTone(ts, ratio, TonePrimary); got != want {
			t.Errorf("thresholdTone(%v) = %v, want %v", ratio, got, want)
		}
	}
}

func TestFormatNumber(t *testing.T) {
	cases := map[float64]string{12: "12", 932.5: "932.5", 1.005: "1", 9.256: "9.26", -0.001: "0", -4: "-4"}
	for v, want := range cases {
		if got := formatNumber(v); got != want {
			t.Errorf("formatNumber(%v) = %q, want %q", v, got, want)
		}
	}
}
