package pterm

import (
	"io"
	"math"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// SparklinePrinter renders a series of numbers as a one-line chart, one
// column per value, with an optional muted label after it:
//
//	▁▂▄▃▅▇█▆▄▂  requests per second
//
// Thresholds tone each column by its own level, for example to show the
// load of every CPU core at a glance.
type SparklinePrinter struct {
	// Values are drawn from left to right.
	Values []float64
	// Tone colors the columns when no thresholds are set.
	Tone Tone
	// Thresholds tone every column by its level on the scale.
	Thresholds []Threshold
	// Min and Max fix the scale. When they are equal, the scale runs from
	// the smallest to the largest value.
	Min, Max float64
	// Label follows the sparkline.
	Label string
	// LabelTone colors the label.
	LabelTone Tone
	// Width limits the number of columns. When there are more values,
	// the most recent ones are shown. Zero fits the terminal.
	Width int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Sparkline is the default sparkline printer.
var Sparkline = SparklinePrinter{
	Tone:      TonePrimary,
	LabelTone: ToneMuted,
}

// WithValues returns a copy of p that draws values.
func (p SparklinePrinter) WithValues(values ...float64) *SparklinePrinter {
	p.Values = values
	return &p
}

// WithTone returns a copy of p whose columns use tone t.
func (p SparklinePrinter) WithTone(t Tone) *SparklinePrinter {
	p.Tone = t
	return &p
}

// WithThresholds returns a copy of p that tones every column by its own
// level, for example LoadThresholds().
func (p SparklinePrinter) WithThresholds(thresholds ...Threshold) *SparklinePrinter {
	p.Thresholds = thresholds
	return &p
}

// WithRange returns a copy of p with a fixed scale from lo to hi, such as
// 0 to 100 for percentages.
func (p SparklinePrinter) WithRange(lo, hi float64) *SparklinePrinter {
	p.Min, p.Max = lo, hi
	return &p
}

// WithLabel returns a copy of p that prints label after the sparkline.
func (p SparklinePrinter) WithLabel(label string) *SparklinePrinter {
	p.Label = label
	return &p
}

// WithLabelTone returns a copy of p whose label uses tone t.
func (p SparklinePrinter) WithLabelTone(t Tone) *SparklinePrinter {
	p.LabelTone = t
	return &p
}

// WithWidth returns a copy of p that draws at most w columns.
func (p SparklinePrinter) WithWidth(w int) *SparklinePrinter {
	p.Width = w
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p SparklinePrinter) WithWriter(w io.Writer) *SparklinePrinter {
	p.Writer = w
	return &p
}

func (p SparklinePrinter) out() io.Writer { return orDefault(p.Writer) }

// Srender returns the sparkline. It returns ErrEmpty without values.
func (p SparklinePrinter) Srender() (string, error) {
	if len(p.Values) == 0 {
		return "", ErrEmpty
	}

	return p.render(canvasFor(p.out())), nil
}

// Render prints the sparkline.
func (p SparklinePrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

func (p SparklinePrinter) render(c canvas) string {
	values := p.Values
	limit := p.Width

	if c.width > 0 {
		room := c.width
		if p.Label != "" {
			room -= text.Width(p.Label) + 2
		}

		if limit <= 0 || limit > room {
			limit = max(room, 1)
		}
	}

	if limit > 0 && len(values) > limit {
		values = values[len(values)-limit:]
	}

	sc := fitScale(p.Min, p.Max, values, false)
	levels := len(c.g.Columns)

	var b strings.Builder

	// Runs of columns in the same tone are painted together.
	var (
		run     strings.Builder
		runTone Tone
	)

	flush := func() {
		if run.Len() > 0 {
			b.WriteString(c.paint(runTone, run.String()))
			run.Reset()
		}
	}

	for _, v := range values {
		if math.IsNaN(v) || levels == 0 {
			flush()
			b.WriteString(" ")

			continue
		}

		r := sc.ratio(v)
		if sc.hi == sc.lo {
			r = 0.5
		}

		tone := p.Tone
		if len(p.Thresholds) > 0 {
			tone = thresholdTone(p.Thresholds, r, p.Tone)
		}

		if tone != runTone {
			flush()

			runTone = tone
		}

		run.WriteString(c.g.Columns[int(math.Round(r*float64(levels-1)))])
	}

	flush()

	if p.Label != "" {
		b.WriteString("  " + c.paint(p.LabelTone, p.Label))
	}

	return b.String()
}
