package pterm

import (
	"math"
	"strconv"

	"atomicgo.dev/color"
)

// Helpers shared by the chart printers: the series tone cycle, thresholds,
// number formatting and gradients.

// seriesTones is the cycle charts color their series from, see the
// styleguide.
var seriesTones = []Tone{TonePrimary, ToneAccent, ToneInfo, ToneSuccess, ToneWarning, ToneError}

// SeriesTones returns the tones charts color their series with, in order.
// A chart with more series starts over at the first tone.
func SeriesTones() []Tone {
	return append([]Tone(nil), seriesTones...)
}

// cycleTone returns tone i of tones, or of the default series cycle when
// tones is empty.
func cycleTone(tones []Tone, i int) Tone {
	if len(tones) == 0 {
		tones = seriesTones
	}

	return tones[i%len(tones)]
}

// TrafficTones returns Success, Warning and Error: a gradient or a cycle
// for charts where high values are bad.
func TrafficTones() []Tone {
	return []Tone{ToneSuccess, ToneWarning, ToneError}
}

// Threshold gives values from a point of a scale upward a tone. Meters
// and sparklines pick the tone of the highest threshold a value reaches.
type Threshold struct {
	// From is where the threshold starts, as a fraction of the scale:
	// 0 is the bottom, 1 the top.
	From float64
	// Tone colors values at or above From.
	Tone Tone
}

// LoadThresholds is the usual traffic light: Success below 70%, Warning
// below 90%, Error from there.
func LoadThresholds() []Threshold {
	return []Threshold{{0, ToneSuccess}, {0.7, ToneWarning}, {0.9, ToneError}}
}

// thresholdTone returns the tone of the highest threshold ratio reaches,
// or fallback when it reaches none.
func thresholdTone(ts []Threshold, ratio float64, fallback Tone) Tone {
	best := math.Inf(-1)

	for _, t := range ts {
		if ratio >= t.From && t.From >= best {
			best, fallback = t.From, t.Tone
		}
	}

	return fallback
}

// formatNumber prints v with at most two decimals and no trailing zeros.
func formatNumber(v float64) string {
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return strconv.FormatFloat(v, 'f', -1, 64)
	}

	s := strconv.FormatFloat(math.Round(v*100)/100, 'f', -1, 64)
	if s == "-0" {
		return "0"
	}

	return s
}

// gradientAt returns the color at position t (0 to 1) of a gradient
// through stops, mixed in the colors of theme th.
func gradientAt(th Theme, stops []Tone, t float64) color.Color {
	switch len(stops) {
	case 0:
		return color.None
	case 1:
		return stops[0].In(th)
	}

	t = min(max(t, 0), 1)
	pos := t * float64(len(stops)-1)
	i := min(int(pos), len(stops)-2)

	return stops[i].In(th).Mix(stops[i+1].In(th), pos-float64(i))
}

// scale maps values onto a range. Charts with an unset range (Min equal
// to Max) fit it to their data.
type scale struct {
	lo, hi float64
}

// fitScale returns the scale for values: [lo, hi] when set, otherwise the
// data range, widened to include zero when zero is true.
func fitScale(lo, hi float64, values []float64, zero bool) scale {
	if lo != hi {
		return scale{min(lo, hi), max(lo, hi)}
	}

	lo, hi = math.Inf(1), math.Inf(-1)

	for _, v := range values {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			continue
		}

		lo, hi = min(lo, v), max(hi, v)
	}

	if lo > hi {
		lo, hi = 0, 0
	}

	if zero {
		lo, hi = min(lo, 0), max(hi, 0)
	}

	return scale{lo, hi}
}

// ratio returns where v sits on the scale, from 0 to 1. A scale without
// a range puts everything at the top, unless everything is zero.
func (s scale) ratio(v float64) float64 {
	if s.hi == s.lo {
		if v == 0 && s.hi == 0 {
			return 0
		}

		return 1
	}

	return min(max((v-s.lo)/(s.hi-s.lo), 0), 1)
}
