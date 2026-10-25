package pterm

import (
	"errors"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// PieSlices is used to display multiple PieSlice.
type PieSlices []PieSlice

// PieSlice is a single slice of a pie chart.
type PieSlice struct {
	Label string
	Value float64
	// Style colors the slice and its legend marker.
	// If nil, the next color of PieChartPrinter.Colors is used.
	Style *Style
}

// PieChartPrinter is used to print pie charts.
type PieChartPrinter struct {
	Writer io.Writer
	Slices PieSlices
	// Radius is the radius of the pie in terminal rows.
	// Every row cell is two columns wide, so the pie looks round in most terminals.
	Radius int
	// Colors are cycled through for slices without their own Style.
	Colors []Color
	// Characters fill the slices when color output is disabled,
	// so the slices can still be told apart. Each entry should be two columns wide.
	Characters []string
	// LabelStyle styles the slice labels in the legend.
	LabelStyle *Style
	// ShowValue adds the raw value of every slice to the legend.
	ShowValue bool
	// ShowPercentage adds the share of every slice to the legend.
	ShowPercentage bool
}

var (
	// DefaultPieChart is the default PieChartPrinter.
	DefaultPieChart = PieChartPrinter{
		Radius:         6,
		Colors:         ThemeDefault.PieChartColors,
		Characters:     []string{"██", "▓▓", "▒▒", "░░", "##", "**", "++", "::"},
		LabelStyle:     &ThemeDefault.PieChartLabelStyle,
		ShowPercentage: true,
	}
)

// WithSlices returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithSlices(slices PieSlices) *PieChartPrinter {
	p.Slices = slices
	return &p
}

// WithRadius returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithRadius(radius int) *PieChartPrinter {
	p.Radius = radius
	return &p
}

// WithColors returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithColors(colors []Color) *PieChartPrinter {
	p.Colors = colors
	return &p
}

// WithCharacters returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithCharacters(characters []string) *PieChartPrinter {
	p.Characters = characters
	return &p
}

// WithLabelStyle returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithLabelStyle(style *Style) *PieChartPrinter {
	p.LabelStyle = style
	return &p
}

// WithShowValue returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithShowValue(b ...bool) *PieChartPrinter {
	p.ShowValue = internal.WithBoolean(b)
	return &p
}

// WithShowPercentage returns a new PieChartPrinter with a specific option.
func (p PieChartPrinter) WithShowPercentage(b ...bool) *PieChartPrinter {
	p.ShowPercentage = internal.WithBoolean(b)
	return &p
}

// WithWriter sets the custom Writer.
func (p PieChartPrinter) WithWriter(writer io.Writer) *PieChartPrinter {
	p.Writer = writer
	return &p
}

// Srender renders the PieChartPrinter as a string.
func (p PieChartPrinter) Srender() (string, error) {
	total := 0.0

	for _, slice := range p.Slices {
		if slice.Value < 0 || math.IsNaN(slice.Value) || math.IsInf(slice.Value, 0) {
			return "", errors.New("pie chart values must be finite and not negative")
		}

		total += slice.Value
	}

	if total == 0 {
		return "", errors.New("pie chart needs at least one slice with a value above zero")
	}

	if rawOutput() {
		return p.getRawOutput(total), nil
	}

	markers := p.markers()
	pie := p.renderPie(total, markers)
	legend := p.renderLegend(total, markers)

	// Center the legend vertically next to the pie.
	offset := max((len(pie)-len(legend))/2, 0)
	lineCount := max(len(pie), offset+len(legend))
	pieWidth := (p.radius()*2 + 1) * 2

	var ret strings.Builder

	for i := range lineCount {
		line := strings.Repeat(" ", pieWidth)
		if i < len(pie) {
			line = pie[i]
		}

		if j := i - offset; j >= 0 && j < len(legend) {
			line += "   " + legend[j]
		}

		ret.WriteString(strings.TrimRight(line, " "))
		ret.WriteByte('\n')
	}

	return ret.String(), nil
}

// Render prints the PieChartPrinter to the terminal.
func (p PieChartPrinter) Render() error {
	s, err := p.Srender()
	if err != nil {
		return err
	}

	Fprint(p.Writer, s)

	return nil
}

func (p PieChartPrinter) radius() int {
	if p.Radius < 1 {
		return 1
	}

	return p.Radius
}

// pieMarker is the fill of one slice: a two-column character and the style
// it is printed with.
type pieMarker struct {
	char  string
	style *Style
}

func (m pieMarker) sprint(count int) string {
	s := strings.Repeat(m.char, count)
	if m.style == nil {
		return s
	}

	return m.style.Sprint(s)
}

// markers returns the fill of every slice. With color enabled, slices share
// one fill character and differ by color; without color, every slice gets
// its own character instead.
func (p PieChartPrinter) markers() []pieMarker {
	markers := make([]pieMarker, len(p.Slices))
	colorEnabled := printColorEnabled()
	colorIndex := 0

	for i, slice := range p.Slices {
		markers[i] = pieMarker{char: "██", style: slice.Style}

		if !colorEnabled && len(p.Characters) > 0 {
			markers[i].char = p.Characters[i%len(p.Characters)]
		}

		if markers[i].style == nil && len(p.Colors) > 0 {
			markers[i].style = NewStyle(p.Colors[colorIndex%len(p.Colors)])
			colorIndex++
		}
	}

	return markers
}

// renderPie draws the pie row by row. The first slice starts at twelve
// o'clock and the slices follow clockwise.
func (p PieChartPrinter) renderPie(total float64, markers []pieMarker) []string {
	radius := p.radius()
	// Treat a cell as inside when its center is within half a cell of the
	// circle, which rounds off the edges nicely at small sizes.
	limit := (float64(radius) + 0.5) * (float64(radius) + 0.5)

	ends := make([]float64, len(p.Slices))
	sum := 0.0

	for i, slice := range p.Slices {
		sum += slice.Value
		ends[i] = sum / total
	}

	lines := make([]string, 0, radius*2+1)

	for y := -radius; y <= radius; y++ {
		var line strings.Builder

		// Consecutive cells of the same slice are printed as one run, so
		// every run needs only one pair of color codes. -1 marks a cell
		// outside the pie.
		run, runLength := -1, 0
		flush := func() {
			if run < 0 {
				line.WriteString(strings.Repeat("  ", runLength))
			} else {
				line.WriteString(markers[run].sprint(runLength))
			}
		}

		for x := -radius; x <= radius; x++ {
			index := pieSliceAt(x, y, limit, ends)
			if index != run && runLength > 0 {
				flush()

				runLength = 0
			}

			run = index
			runLength++
		}

		flush()

		lines = append(lines, line.String())
	}

	return lines
}

// pieSliceAt returns the index of the slice that covers the cell at x, y
// (relative to the center of the pie), or -1 if the cell is outside the pie.
func pieSliceAt(x, y int, limit float64, ends []float64) int {
	if float64(x*x+y*y) > limit {
		return -1
	}

	angle := math.Atan2(float64(x), float64(-y))
	if angle < 0 {
		angle += 2 * math.Pi
	}

	fraction := angle / (2 * math.Pi)

	for i, end := range ends {
		if fraction < end {
			return i
		}
	}

	return len(ends) - 1
}

func (p PieChartPrinter) renderLegend(total float64, markers []pieMarker) []string {
	labelStyle := p.LabelStyle
	if labelStyle == nil {
		labelStyle = NewStyle()
	}

	labelWidth := 0
	for _, slice := range p.Slices {
		labelWidth = max(labelWidth, runewidth.StringWidth(slice.Label))
	}

	lines := make([]string, 0, len(p.Slices))

	for i, slice := range p.Slices {
		line := markers[i].sprint(1) + " " + labelStyle.Sprint(slice.Label)

		// Pad outside of the label style, and only when there is a detail
		// column to align, so lines never end in styled blanks.
		if details := p.details(slice, total); details != "" {
			line += strings.Repeat(" ", labelWidth-runewidth.StringWidth(slice.Label)) + "  " + details
		}

		lines = append(lines, line)
	}

	return lines
}

// details returns the value and percentage of a slice, depending on the
// printer settings.
func (p PieChartPrinter) details(slice PieSlice, total float64) string {
	var parts []string

	if p.ShowValue {
		parts = append(parts, strconv.FormatFloat(slice.Value, 'f', -1, 64))
	}

	if p.ShowPercentage {
		percentage := strconv.FormatFloat(slice.Value/total*100, 'f', 1, 64) + "%"
		if p.ShowValue {
			percentage = "(" + percentage + ")"
		}

		parts = append(parts, percentage)
	}

	return strings.Join(parts, " ")
}

func (p PieChartPrinter) getRawOutput(total float64) string {
	var ret strings.Builder

	for _, slice := range p.Slices {
		ret.WriteString(slice.Label + ": " + strconv.FormatFloat(slice.Value, 'f', -1, 64))

		if p.ShowPercentage {
			ret.WriteString(" (" + strconv.FormatFloat(slice.Value/total*100, 'f', 1, 64) + "%)")
		}

		ret.WriteByte('\n')
	}

	return ret.String()
}
