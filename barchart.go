package pterm

import (
	"io"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// BarChartPrinter is used to print bar charts.
type BarChartPrinter struct {
	Writer     io.Writer
	Bars       Bars
	Horizontal bool
	ShowValue  bool
	// Height sets the maximum height of a vertical bar chart.
	// The default is calculated to fit into the terminal.
	// Ignored if Horizontal is set to true.
	Height int
	// Width sets the maximum width of a horizontal bar chart.
	// The default is calculated to fit into the terminal.
	// Ignored if Horizontal is set to false.
	Width                  int
	VerticalBarCharacter   string
	HorizontalBarCharacter string
}

var (
	// DefaultBarChart is the default BarChartPrinter.
	DefaultBarChart = BarChartPrinter{
		Horizontal:             false,
		VerticalBarCharacter:   "██",
		HorizontalBarCharacter: "█",
		// keep in sync with RecalculateTerminalSize()
		Height: GetTerminalHeight() * 2 / 3,
		Width:  GetTerminalWidth() * 2 / 3,
	}
)

// WithBars returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithBars(bars Bars) *BarChartPrinter {
	p.Bars = bars
	return &p
}

// WithVerticalBarCharacter returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithVerticalBarCharacter(char string) *BarChartPrinter {
	p.VerticalBarCharacter = char
	return &p
}

// WithHorizontalBarCharacter returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithHorizontalBarCharacter(char string) *BarChartPrinter {
	p.HorizontalBarCharacter = char
	return &p
}

// WithHorizontal returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithHorizontal(b ...bool) *BarChartPrinter {
	b2 := internal.WithBoolean(b)
	p.Horizontal = b2
	return &p
}

// WithHeight returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithHeight(value int) *BarChartPrinter {
	p.Height = value
	return &p
}

// WithWidth returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithWidth(value int) *BarChartPrinter {
	p.Width = value
	return &p
}

// WithShowValue returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithShowValue(b ...bool) *BarChartPrinter {
	p.ShowValue = internal.WithBoolean(b)
	return &p
}

// WithWriter sets the custom Writer.
func (p BarChartPrinter) WithWriter(writer io.Writer) *BarChartPrinter {
	p.Writer = writer
	return &p
}

func (p BarChartPrinter) getRawOutput() string {
	var ret string

	for _, bar := range p.Bars {
		ret += Sprintfln("%s: %d", bar.Label, bar.Value)
	}

	return ret
}

// Srender renders the BarChart as a string.
func (p BarChartPrinter) Srender() (string, error) {
	maxAbsValue := func(value1 int, value2 int) int {
		min := value1
		max := value2

		if value1 > value2 {
			min = value2
			max = value1
		}

		maxAbs := max

		if min < 0 && -min > max { // This is to avoid something like "int(math.Abs(float64(minBarValue)))"
			maxAbs = -min // (--) == (+)
		}

		return maxAbs
	}

	abs := func(value int) int {
		if value < 0 {
			return -value
		}

		return value
	}
	// =================================== VERTICAL BARS RENDERER ======================================================

	type renderParams struct {
		repeatCount             int
		bar                     Bar
		positiveChartPartHeight int
		negativeChartPartHeight int
		positiveChartPartWidth  int
		negativeChartPartWidth  int
		indent                  string
		showValue               bool
		moveUp                  bool
		moveRight               bool
	}

	renderPositiveVerticalBar := func(renderedBarRef *string, rParams renderParams) {
		if rParams.showValue {
			*renderedBarRef += Sprint(rParams.indent + strconv.Itoa(rParams.bar.Value) + rParams.indent + "\n")
		}

		for i := rParams.positiveChartPartHeight; i > 0; i-- {
			if i > rParams.repeatCount {
				*renderedBarRef += rParams.indent + "  " + rParams.indent + " \n"
			} else {
				*renderedBarRef += rParams.indent + rParams.bar.Style.Sprint(p.VerticalBarCharacter) + rParams.indent + " \n"
			}
		}

		// Used when we draw diagram with both POSITIVE and NEGATIVE values.
		// In such case we separately draw top and bottom half of chart.
		// And we need MOVE UP positive part to top part of chart,
		// technically by adding empty pillars with height == height of chart's bottom part.
		if rParams.moveUp {
			for i := 0; i <= rParams.negativeChartPartHeight; i++ {
				*renderedBarRef += rParams.indent + "  " + rParams.indent + " \n"
			}
		}
	}

	renderNegativeVerticalBar := func(renderedBarRef *string, rParams renderParams) {
		for i := 0; i > -rParams.negativeChartPartHeight; i-- {
			if i > rParams.repeatCount {
				*renderedBarRef += rParams.indent + rParams.bar.Style.Sprint(p.VerticalBarCharacter) + rParams.indent + " \n"
			} else {
				*renderedBarRef += rParams.indent + "  " + rParams.indent + " \n"
			}
		}

		if rParams.showValue {
			*renderedBarRef += Sprint(rParams.indent + strconv.Itoa(rParams.bar.Value) + rParams.indent + "\n")
		}
	}

	// =================================== HORIZONTAL BARS RENDERER ====================================================
	renderPositiveHorizontalBar := func(renderedBarRef *string, rParams renderParams) {
		if rParams.moveRight {
			for i := 0; i < rParams.negativeChartPartWidth; i++ {
				*renderedBarRef += " "
			}
		}

		for i := 0; i < rParams.positiveChartPartWidth; i++ {
			if i < rParams.repeatCount {
				*renderedBarRef += rParams.bar.Style.Sprint(p.HorizontalBarCharacter)
			} else {
				*renderedBarRef += " "
			}
		}

		if rParams.showValue {
			// For positive horizontal bars we add one more space before adding value,
			// so they will be well aligned with negative values, which have "-" sign before them
			*renderedBarRef += " "

			*renderedBarRef += " " + strconv.Itoa(rParams.bar.Value)
		}
	}

	renderNegativeHorizontalBar := func(renderedBarRef *string, rParams renderParams) {
		for i := -rParams.negativeChartPartWidth; i < 0; i++ {
			if i < rParams.repeatCount {
				*renderedBarRef += " "
			} else {
				*renderedBarRef += rParams.bar.Style.Sprint(p.HorizontalBarCharacter)
			}
		}

		// In order to print values well-aligned (in case when we have both - positive and negative part of chart),
		// we should insert an indent with width == width of positive chart part
		if rParams.positiveChartPartWidth > 0 {
			for i := 0; i < rParams.positiveChartPartWidth; i++ {
				*renderedBarRef += " "
			}
		}

		if rParams.showValue {
			/*
				This is in order to achieve this effect:
				 0
				-15
				 0
				-19

				INSTEAD OF THIS:

				0
				-15
				0
				-19
			*/
			if rParams.repeatCount == 0 {
				*renderedBarRef += " "
			}

			*renderedBarRef += " " + strconv.Itoa(rParams.bar.Value)
		}
	}
	// =================================================================================================================

	if RawOutput {
		return p.getRawOutput(), nil
	}
	for i, bar := range p.Bars {
		if bar.Style == nil {
			p.Bars[i].Style = &ThemeDefault.BarStyle
		}

		if bar.LabelStyle == nil {
			p.Bars[i].LabelStyle = &ThemeDefault.BarLabelStyle
		}

		p.Bars[i].Label = p.Bars[i].LabelStyle.Sprint(bar.Label)
	}

	var ret string

	var maxLabelHeight int
	var maxBarValue int
	var minBarValue int
	var maxAbsBarValue int
	var rParams renderParams

	for _, bar := range p.Bars {
		if bar.Value > maxBarValue {
			maxBarValue = bar.Value
		}
		if bar.Value < minBarValue {
			minBarValue = bar.Value
		}
		labelHeight := len(strings.Split(bar.Label, "\n"))
		if labelHeight > maxLabelHeight {
			maxLabelHeight = labelHeight
		}
	}

	maxAbsBarValue = maxAbsValue(maxBarValue, minBarValue)

	if p.Horizontal {
		panels := Panels{[]Panel{{}, {}}}

		rParams.showValue = p.ShowValue
		rParams.positiveChartPartWidth = p.Width
		rParams.negativeChartPartWidth = p.Width

		// If chart will consist of two parts - positive and negative - we should recalculate max bars WIDTH in LEFT and RIGHT parts
		if minBarValue < 0 && maxBarValue > 0 {
			rParams.positiveChartPartWidth = abs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(maxBarValue)))
			rParams.negativeChartPartWidth = abs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(minBarValue)))
		}

		for _, bar := range p.Bars {
			rParams.bar = bar
			panels[0][0].Data += "\n" + bar.Label
			panels[0][1].Data += "\n"

			if minBarValue >= 0 {
				// As we don't have negative values, draw only positive (right) part of the chart:
				rParams.repeatCount = internal.MapRangeToRange(0, float32(maxAbsBarValue), 0, float32(p.Width), float32(bar.Value))
				rParams.moveRight = false

				renderPositiveHorizontalBar(&panels[0][1].Data, rParams)
			} else if maxBarValue <= 0 {
				// As we have only negative values, draw only negative (left) part of the chart:
				rParams.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), 0, -float32(p.Width), 0, float32(bar.Value))
				rParams.positiveChartPartWidth = 0

				renderNegativeHorizontalBar(&panels[0][1].Data, rParams)
			} else {
				// We have positive and negative values, so draw both (left+right) parts of the chart:
				rParams.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(bar.Value))

				if bar.Value >= 0 {
					rParams.moveRight = true

					renderPositiveHorizontalBar(&panels[0][1].Data, rParams)
				}

				if bar.Value < 0 {
					renderNegativeHorizontalBar(&panels[0][1].Data, rParams)
				}
			}
		}
		ret, _ = DefaultPanel.WithPanels(panels).Srender()
		return ret, nil
	} else {
		renderedBars := make([]string, len(p.Bars))

		rParams.showValue = p.ShowValue
		rParams.positiveChartPartHeight = p.Height
		rParams.negativeChartPartHeight = p.Height

		// If chart will consist of two parts - positive and negative - we should recalculate max bars height in top and bottom parts
		if minBarValue < 0 && maxBarValue > 0 {
			rParams.positiveChartPartHeight = abs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(maxBarValue)))
			rParams.negativeChartPartHeight = abs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(minBarValue)))
		}

		for i, bar := range p.Bars {
			var renderedBar string
			rParams.bar = bar
			rParams.indent = strings.Repeat(" ", internal.GetStringMaxWidth(RemoveColorFromString(bar.Label))/2)

			if minBarValue >= 0 {
				// As we don't have negative values, draw only positive (top) part of the chart:
				rParams.repeatCount = internal.MapRangeToRange(0, float32(maxAbsBarValue), 0, float32(p.Height), float32(bar.Value))
				rParams.moveUp = false // Don't MOVE UP as we have ONLY positive part of chart.

				renderPositiveVerticalBar(&renderedBar, rParams)
			} else if maxBarValue <= 0 {
				// As we have only negative values, draw only negative (bottom) part of the chart:
				rParams.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), 0, -float32(p.Height), 0, float32(bar.Value))

				renderNegativeVerticalBar(&renderedBar, rParams)
			} else {
				// We have positive and negative values, so draw both (top+bottom) parts of the chart:
				rParams.repeatCount = internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(bar.Value))

				if bar.Value >= 0 {
					rParams.moveUp = true // MOVE UP positive part, because we have both positive and negative parts of chart.

					renderPositiveVerticalBar(&renderedBar, rParams)
				}

				if bar.Value < 0 {
					renderNegativeVerticalBar(&renderedBar, rParams)
				}
			}

			labelHeight := len(strings.Split(bar.Label, "\n"))
			renderedBars[i] = renderedBar + bar.Label + strings.Repeat("\n", maxLabelHeight-labelHeight) + " "
		}

		var maxBarHeight int

		for _, bar := range renderedBars {
			totalBarHeight := len(strings.Split(bar, "\n"))
			if totalBarHeight > maxBarHeight {
				maxBarHeight = totalBarHeight
			}
		}

		for i, bar := range renderedBars {
			totalBarHeight := len(strings.Split(bar, "\n"))
			if totalBarHeight < maxBarHeight {
				renderedBars[i] = strings.Repeat("\n", maxBarHeight-totalBarHeight) + renderedBars[i]
			}
		}

		for i := 0; i <= maxBarHeight; i++ {
			for _, barString := range renderedBars {
				var barLine string
				letterLines := strings.Split(barString, "\n")
				maxBarWidth := internal.GetStringMaxWidth(RemoveColorFromString(barString))
				if len(letterLines) > i {
					barLine = letterLines[i]
				}
				letterLineLength := runewidth.StringWidth(RemoveColorFromString(barLine))
				if letterLineLength < maxBarWidth {
					barLine += strings.Repeat(" ", maxBarWidth-letterLineLength)
				}
				ret += barLine
			}
			ret += "\n"
		}
	}

	return ret, nil
}

// Render prints the Template to the terminal.
func (p BarChartPrinter) Render() error {
	s, _ := p.Srender()
	Fprintln(p.Writer, s)

	return nil
}
