package pterm

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// MeterItem is one meter: how full something is, and what it is.
type MeterItem struct {
	// Label follows the percentage, muted.
	Label string
	// Value is how full the meter is, from 0 to 1. Values above 1 show
	// their real percentage on a full meter.
	Value float64
	// Tone colors the fill. The zero value picks the tone from the
	// printer's thresholds.
	Tone Tone
}

// MeterPrinter renders meters: a fill on a dotted track, the percentage
// and a muted label. Meters fill in Success, Warning and Error by default,
// depending on how full they are.
//
//	████████▍░░░░░░░░░░░   42%  disk
//	██████████████████▌░   93%  memory
type MeterPrinter struct {
	// Items are the meters, one per line.
	Items []MeterItem
	// Width is the width of the meter itself, without percentage and
	// label. It shrinks to fit a narrow terminal.
	Width int
	// Thresholds tone the fill by how full a meter is.
	Thresholds []Threshold
	// Tone colors every fill and replaces the thresholds when set.
	Tone Tone
	// TrackTone colors the empty part of a meter.
	TrackTone Tone
	// PercentTone colors the percentage.
	PercentTone Tone
	// LabelTone colors the labels.
	LabelTone Tone
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Meter is the default meter printer.
var Meter = MeterPrinter{
	Width:       20,
	Thresholds:  LoadThresholds(),
	TrackTone:   ToneBorder,
	PercentTone: ToneText,
	LabelTone:   ToneMuted,
}

// WithItems returns a copy of p that renders items.
func (p MeterPrinter) WithItems(items ...MeterItem) *MeterPrinter {
	p.Items = items
	return &p
}

// WithValue returns a copy of p with a single meter filled to value (0 to
// 1) and an optional label.
func (p MeterPrinter) WithValue(value float64, label ...string) *MeterPrinter {
	p.Items = []MeterItem{{Value: value, Label: strings.Join(label, " ")}}
	return &p
}

// WithWidth returns a copy of p whose meters are w columns wide.
func (p MeterPrinter) WithWidth(w int) *MeterPrinter {
	p.Width = w
	return &p
}

// WithThresholds returns a copy of p that tones fills by thresholds.
func (p MeterPrinter) WithThresholds(thresholds ...Threshold) *MeterPrinter {
	p.Thresholds = thresholds
	return &p
}

// WithTone returns a copy of p that fills every meter in tone t.
func (p MeterPrinter) WithTone(t Tone) *MeterPrinter {
	p.Tone = t
	return &p
}

// WithTrackTone returns a copy of p whose tracks use tone t.
func (p MeterPrinter) WithTrackTone(t Tone) *MeterPrinter {
	p.TrackTone = t
	return &p
}

// WithPercentTone returns a copy of p whose percentages use tone t.
func (p MeterPrinter) WithPercentTone(t Tone) *MeterPrinter {
	p.PercentTone = t
	return &p
}

// WithLabelTone returns a copy of p whose labels use tone t.
func (p MeterPrinter) WithLabelTone(t Tone) *MeterPrinter {
	p.LabelTone = t
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p MeterPrinter) WithWriter(w io.Writer) *MeterPrinter {
	p.Writer = w
	return &p
}

func (p MeterPrinter) out() io.Writer { return orDefault(p.Writer) }

// Srender returns the meters. It returns ErrEmpty without items.
func (p MeterPrinter) Srender() (string, error) {
	if len(p.Items) == 0 {
		return "", ErrEmpty
	}

	return p.render(canvasFor(p.out())), nil
}

// Render prints the meters.
func (p MeterPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

func (p MeterPrinter) render(c canvas) string {
	percents := make([]string, len(p.Items))
	percentWidth := 3

	for i, item := range p.Items {
		percents[i] = percent(item.Value)
		percentWidth = max(percentWidth, text.Width(percents[i]))
	}

	width := max(p.Width, 1)
	if c.width > 0 {
		width = max(min(width, c.width-percentWidth-2), 4)
	}

	lines := make([]string, len(p.Items))

	for i, item := range p.Items {
		tone := item.Tone
		if tone == (Tone{}) {
			tone = p.Tone
		}

		if tone == (Tone{}) {
			tone = thresholdTone(p.Thresholds, item.Value, TonePrimary)
		}

		line := drawMeter(c, item.Value, width, tone, p.TrackTone) +
			"  " + text.PadLeft(c.paint(p.PercentTone, percents[i]), percentWidth)

		if item.Label != "" {
			line += "  " + c.paint(p.LabelTone, item.Label)
		}

		lines[i] = c.cut(line, c.width)
	}

	return strings.Join(lines, "\n")
}

// percent formats a ratio as a whole percentage.
func percent(ratio float64) string {
	if math.IsNaN(ratio) {
		return "0%"
	}

	return fmt.Sprintf("%d%%", int(math.Round(ratio*100)))
}
