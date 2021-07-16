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
		Height:                 GetTerminalHeight() * 2 / 3,
		Width:                  GetTerminalWidth() * 2 / 3,
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

	for _, bar := range p.Bars {
		if bar.Value > maxBarValue {
			maxBarValue = bar.Value
		}
		labelHeight := len(strings.Split(bar.Label, "\n"))
		if labelHeight > maxLabelHeight {
			maxLabelHeight = labelHeight
		}
	}

	if p.Horizontal {
		panels := Panels{[]Panel{{}, {}}}
		for _, bar := range p.Bars {
			panels[0][0].Data += "\n" + bar.Label
			repeatCount := internal.MapRangeToRange(0, float32(maxBarValue), 0, float32(p.Width), float32(bar.Value))
			if repeatCount < 0 {
				repeatCount = 0
			}
			panels[0][1].Data += "\n" + bar.Style.Sprint(strings.Repeat(p.HorizontalBarCharacter, repeatCount))
			if p.ShowValue {
				panels[0][1].Data += " " + strconv.Itoa(bar.Value)
			}
		}
		ret, _ = DefaultPanel.WithPanels(panels).Srender()
		return ret, nil
	} else {
		renderedBars := make([]string, len(p.Bars))

		for i, bar := range p.Bars {
			var renderedBar string
			indent := strings.Repeat(" ", internal.GetStringMaxWidth(RemoveColorFromString(bar.Label))/2)

			if p.ShowValue {
				renderedBar += Sprint(indent + strconv.Itoa(bar.Value) + indent + "\n")
			}

			bar.Value = internal.MapRangeToRange(0, float32(maxBarValue), 0, float32(p.Height), float32(bar.Value))
			for i := 0; i < bar.Value; i++ {
				renderedBar += indent + bar.Style.Sprint(p.VerticalBarCharacter) + indent + " \n"
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
