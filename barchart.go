package pterm

import (
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// BarChartPrinter is used to print bar charts.
type BarChartPrinter struct {
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
	renderPositiveVerticalBar := func(renderedBarRef *string, repeatCount int, bar Bar, chartAbsHeight int, indent string, showValue bool, moveUpBy int) {
		if showValue {
			*renderedBarRef += Sprint(indent + strconv.Itoa(bar.Value) + indent + "\n")
		}

		for i := chartAbsHeight; i > 0; i-- {
			if i > repeatCount {
				*renderedBarRef += indent + "  " + indent + " \n"
			} else {
				*renderedBarRef += indent + bar.Style.Sprint(p.VerticalBarCharacter) + indent + " \n"
			}
		}

		// Used when we draw diagram with both POSITIVE and NEGATIVE values.
		// In such case we separately draw top and bottom half of chart.
		// And we need MOVE UP positive part to top part of chart,
		// technically by adding empty pillars with height == height of chart's bottom part.
		if moveUpBy > 0 {
			for i := 0; i <= moveUpBy; i++ {
				*renderedBarRef += indent + "  " + indent + " \n"
			}
		}
	}

	renderNegativeVerticalBar := func(renderedBarRef *string, repeatCount int, bar Bar, chartAbsHeight int, indent string, showValue bool) {
		for i := 0; i > -chartAbsHeight; i-- {
			if i > repeatCount {
				*renderedBarRef += indent + bar.Style.Sprint(p.VerticalBarCharacter) + indent + " \n"
			} else {
				*renderedBarRef += indent + "  " + indent + " \n"
			}
		}

		if showValue {
			*renderedBarRef += Sprint(indent + strconv.Itoa(bar.Value) + indent + "\n")
		}
	}

	// =================================== HORIZONTAL BARS RENDERER ====================================================
	renderPositiveHorizontalBar := func(renderedBarRef *string, repeatCount int, bar Bar, chartAbsWidth int, showValue bool, moveRightBy int) {
		if moveRightBy > 0 {
			for i := 0; i < moveRightBy; i++ {
				*renderedBarRef += " "
			}
		}

		for i := 0; i <= chartAbsWidth; i++ {
			if i < repeatCount {
				*renderedBarRef += bar.Style.Sprint(p.HorizontalBarCharacter)
			} else {
				*renderedBarRef += " "
			}
		}

		if showValue {
			*renderedBarRef += " " + strconv.Itoa(bar.Value)
		}
	}

	renderNegativeHorizontalBar := func(renderedBarRef *string, repeatCount int, bar Bar, negativeChartPartWidth int, positiveChartPartWidth int, showValue bool) {
		for i := -negativeChartPartWidth; i < 0; i++ {
			if i < repeatCount {
				*renderedBarRef += " "
			} else {
				*renderedBarRef += bar.Style.Sprint(p.HorizontalBarCharacter)
			}
		}

		// In order to print values well-aligned, we should insert an indent with width == width of positive chart part
		if positiveChartPartWidth > 0 {
			for i := 0; i < positiveChartPartWidth; i++ {
				*renderedBarRef += " "
			}
		}

		/*

		This is in order to achieve this effect:
		 0
		-15
		-19
		 0
		-5

		INSTEAD OF THIS:

		0
		-15
		-19
		0
		-5

		*/
		if repeatCount == 0 {
			*renderedBarRef += " "
		}

		if showValue {
			*renderedBarRef += " " + strconv.Itoa(bar.Value)
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

	maxAbsBarValue = maxAbsValue(minBarValue, maxBarValue)

	if p.Horizontal {
		panels := Panels{[]Panel{{}, {}}}

		maxBarWidthInPositivePart := p.Width
		maxBarWidthInNegativePart := p.Width

		// If chart will consist of two parts - positive and negative - we should recalculate max bars WIDTH in LEFT and RIGHT parts
		if minBarValue < 0 && maxBarValue > 0 {
			maxBarWidthInPositivePart = internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(maxBarValue))
			maxBarWidthInNegativePart = abs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(minBarValue)))
		}

		for _, bar := range p.Bars {
			panels[0][0].Data += "\n" + bar.Label
			panels[0][1].Data += "\n"

			if minBarValue >= 0 {
				// As we don't have negative values, draw only positive (right) part of the chart:
				repeatCount := internal.MapRangeToRange(0, float32(maxAbsBarValue), 0, float32(p.Width), float32(bar.Value))
				renderPositiveHorizontalBar(&panels[0][1].Data, repeatCount, bar, maxBarWidthInPositivePart, p.ShowValue, 0)
			} else if maxBarValue <= 0 {
				// As we have only negative values, draw only negative (left) part of the chart:
				repeatCount := internal.MapRangeToRange(-float32(maxAbsBarValue), 0, -float32(p.Width), 0, float32(bar.Value))
				renderNegativeHorizontalBar(&panels[0][1].Data, repeatCount, bar, maxBarWidthInNegativePart, 0, p.ShowValue)
			} else {
				// We have positive and negative values, so draw both (left+right) parts of the chart:
				repeatCount := internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Width)/2, float32(p.Width)/2, float32(bar.Value))
				if bar.Value >= 0 {
					renderPositiveHorizontalBar(&panels[0][1].Data, repeatCount, bar, maxBarWidthInPositivePart, p.ShowValue, maxBarWidthInNegativePart)
				}

				if bar.Value < 0 {
					renderNegativeHorizontalBar(&panels[0][1].Data, repeatCount, bar, maxBarWidthInNegativePart, maxBarWidthInPositivePart, p.ShowValue)
				}
			}
		}
		ret, _ = DefaultPanel.WithPanels(panels).Srender()
		return ret, nil
	} else {
		renderedBars := make([]string, len(p.Bars))

		maxBarHeightInPositivePart := p.Height
		maxBarHeightInNegativePart := p.Height

		// If chart will consist of two parts - positive and negative - we should recalculate max bars height in top and bottom parts
		if minBarValue < 0 && maxBarValue > 0 {
			maxBarHeightInPositivePart = internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(maxBarValue))
			maxBarHeightInNegativePart = abs(internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(minBarValue)))
		}

		for i, bar := range p.Bars {
			var renderedBar string
			indent := strings.Repeat(" ", internal.GetStringMaxWidth(RemoveColorFromString(bar.Label))/2)

			if minBarValue >= 0 {
				// As we don't have negative values, draw only positive (top) part of the chart:
				repeatCount := internal.MapRangeToRange(0, float32(maxAbsBarValue), 0, float32(p.Height), float32(bar.Value))
				renderPositiveVerticalBar(&renderedBar, repeatCount, bar, maxBarHeightInPositivePart, indent, p.ShowValue, 0) // Don't MOVE UP as we have ONLY positive part of chart.
			} else if maxBarValue <= 0 {
				// As we have only negative values, draw only negative (bottom) part of the chart:
				repeatCount := internal.MapRangeToRange(-float32(maxAbsBarValue), 0, -float32(p.Height), 0, float32(bar.Value))
				renderNegativeVerticalBar(&renderedBar, repeatCount, bar, maxBarHeightInNegativePart, indent, p.ShowValue)
			} else {
				// We have positive and negative values, so draw both (top+bottom) parts of the chart:
				repeatCount := internal.MapRangeToRange(-float32(maxAbsBarValue), float32(maxAbsBarValue), -float32(p.Height)/2, float32(p.Height)/2, float32(bar.Value))
				if bar.Value >= 0 {
					renderPositiveVerticalBar(&renderedBar, repeatCount, bar, maxBarHeightInPositivePart, indent, p.ShowValue, maxBarHeightInNegativePart) // MOVE UP positive part, because we have both positive and negative parts of chart.
				}

				if bar.Value < 0 {
					renderNegativeVerticalBar(&renderedBar, repeatCount, bar, maxBarHeightInNegativePart, indent, p.ShowValue)
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
	Println(s)

	return nil
}
