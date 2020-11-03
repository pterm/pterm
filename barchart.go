package pterm

import (
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// BarChartPrinter is used to print bar charts.
type BarChartPrinter struct {
	Bars       Bars
	Horizontal bool
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

// WithBarVerticalCharacter returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithBarVerticalCharacter(char string) *BarChartPrinter {
	p.VerticalBarCharacter = char
	return &p
}

// WithBarHorizontalCharacter returns a new BarChartPrinter with a specific option.
func (p BarChartPrinter) WithBarHorizontalCharacter(char string) *BarChartPrinter {
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

// Srender renders the Template as a string.
func (p BarChartPrinter) Srender() (string, error) {
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
			panels[0][1].Data += "\n" + bar.Style.Sprint(strings.Repeat(p.HorizontalBarCharacter, internal.MapRangeToRange(0, float32(maxBarValue), 0, float32(p.Width), float32(bar.Value))))
		}
		ret, _ = DefaultPanel.WithPanels(panels).Srender()
		return ret, nil
	} else {
		renderedBars := make([]string, len(p.Bars))

		for i, bar := range p.Bars {
			var renderedBar string
			bar.Value = internal.MapRangeToRange(0, float32(maxBarValue), 0, float32(p.Height), float32(bar.Value))
			for i := 0; i < bar.Value; i++ {
				indent := strings.Repeat(" ", internal.GetStringMaxWidth(RemoveColorFromString(bar.Label))/2)
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
				maxBarWidth := internal.GetStringMaxWidth(barString)
				if len(letterLines) > i {
					barLine = letterLines[i]
				}
				letterLineLength := runewidth.StringWidth(barLine)
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
	s, err := p.Srender()
	if err != nil {
		return err
	}
	Println(s)

	return nil
}
