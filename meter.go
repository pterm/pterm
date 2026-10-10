package pterm

import (
	"math"
	"strings"
)

// This file holds the drawing core that meters, bar charts and progress
// bars share. A bar is a fill in a tone, exact to an eighth of a column,
// on an optional track that runs up to a fixed width.

// drawMeter draws a meter width columns wide: a fill in tone fill that
// covers ratio (0 to 1) of the width, exact to an eighth of a column, on
// a track of the glyph set's meter track in tone track (ToneBorder by
// default in every printer). Ratios outside 0 to 1 are clamped. The
// result is always exactly width columns wide.
//
//	████████▍░░░░░░░
func drawMeter(c canvas, ratio float64, width int, fill, track Tone) string {
	if math.IsNaN(ratio) {
		ratio = 0
	}

	ratio = min(max(ratio, 0), 1)

	return drawBar(c, ratio*float64(width), width, fill, c.g.MeterTrack, track)
}

// segment is one part of a stacked bar: length columns in tone.
type segment struct {
	length float64
	tone   Tone
}

// drawBar draws a bar of length columns in tone, then fills the rest of
// width with track in trackTone. An empty track pads with spaces, so the
// result is always width columns wide.
func drawBar(c canvas, length float64, width int, tone Tone, track string, trackTone Tone) string {
	return drawStack(c, []segment{{length, tone}}, width, track, trackTone)
}

// drawStack draws segments one after the other, growing to the right. The
// boundaries between segments fall on whole columns, because a cell can
// only have one color, and the end of the last segment is exact to an
// eighth of a column.
func drawStack(c canvas, segs []segment, width int, track string, trackTone Tone) string {
	var b strings.Builder

	used := 0
	end := 0.0

	for i, s := range segs {
		end += max(s.length, 0)

		if i < len(segs)-1 {
			cells := min(int(math.Round(end)), width) - used
			if cells > 0 {
				b.WriteString(c.paint(s.tone, strings.Repeat(c.g.MeterFill, cells)))
				used += cells
			}

			continue
		}

		fill, cells := c.eighthBar(min(end, float64(width))-float64(used), width-used)
		b.WriteString(c.paint(s.tone, fill))

		used += cells
	}

	b.WriteString(c.track(track, trackTone, width-used))

	return b.String()
}

// drawStackLeft is drawStack for bars that grow to the left, such as
// negative values: the track comes first and the first segment ends at
// the right edge. Without eighth blocks for this direction, the free end
// is exact to half a column.
func drawStackLeft(c canvas, segs []segment, width int, track string, trackTone Tone) string {
	// bounds[i] is the column, counted from the right edge, where segment
	// i ends. Inner bounds are whole columns.
	bounds := make([]int, len(segs))
	end := 0.0

	for i, s := range segs {
		end += max(s.length, 0)
		bounds[i] = min(int(math.Round(end)), width)
	}

	end = min(end, float64(width))
	halves := int(math.Round(end * 2))
	full, half := halves/2, halves%2 == 1

	if half && c.g.RightHalf == "" {
		full, half = min(full+1, width), false
	}

	cells := full
	if half {
		cells++
	}

	var b strings.Builder

	b.WriteString(c.track(track, trackTone, width-cells))

	for i := len(segs) - 1; i >= 0; i-- {
		inner := 0
		if i > 0 {
			inner = min(bounds[i-1], full)
		}

		outer := min(bounds[i], full)
		if i == len(segs)-1 {
			outer = full

			if half {
				b.WriteString(c.paint(segs[i].tone, c.g.RightHalf))
			}
		}

		if n := outer - inner; n > 0 {
			b.WriteString(c.paint(segs[i].tone, strings.Repeat(c.g.MeterFill, n)))
		}
	}

	return b.String()
}

// eighthBar returns a fill length columns long, exact to an eighth, and
// the number of columns it takes, never more than limit. Glyph sets
// without eighth blocks round to whole columns.
func (c canvas) eighthBar(length float64, limit int) (string, int) {
	if limit <= 0 || length <= 0 {
		return "", 0
	}

	eighths := int(math.Round(min(length, float64(limit)) * 8))
	full, part := eighths/8, eighths%8

	partial := ""
	if part > 0 && part <= len(c.g.Eighths) {
		partial = c.g.Eighths[part-1]
	}

	if part > 0 && partial == "" && part >= 4 {
		full++
	}

	fill := strings.Repeat(c.g.MeterFill, full)
	if partial != "" {
		fill += partial
		full++
	}

	return fill, full
}

// track returns n columns of track glyphs in tone, or n spaces when
// glyph is empty.
func (c canvas) track(glyph string, tone Tone, n int) string {
	if n <= 0 {
		return ""
	}

	if glyph == "" {
		return strings.Repeat(" ", n)
	}

	return c.paint(tone, strings.Repeat(glyph, n))
}
