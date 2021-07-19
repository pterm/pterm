package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestBarChartPrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	p := pterm.BarChartPrinter{}
	err := p.Render()
	if err != nil {
		panic(err)
	}
}

func TestBarChartPrinterNilStylePrint(t *testing.T) {
	proxyToDevNull()
	bars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 5,
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
		},
	}

	err := pterm.DefaultBarChart.WithBars(bars).Render()
	if err != nil {
		panic(err)
	}
}

func TestBarChartPrinter_RenderExample(t *testing.T) {
	proxyToDevNull()
	bars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 5,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
	}

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
}

func TestBarChartPrinter_RenderExampleRawOutput(t *testing.T) {
	proxyToDevNull()
	pterm.DisableStyling()
	bars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 5,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
	}

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
	pterm.EnableStyling()
}

func TestBarChartPrinter_RenderMultipleLineLabel(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: -1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test\nNew Line",
			Value: -1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderNegativeBarValues(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: -1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderZeroBarValuesHorizontal(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderZeroBarValues(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderLowBarValues(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 1,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 1,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_Render(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderHorizonzal(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

func TestBarChartPrinter_WithHorizontalBarCharacter(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := "X"
	p2 := p.WithHorizontalBarCharacter(s)

	assert.Equal(t, s, p2.HorizontalBarCharacter)
	assert.Empty(t, p.HorizontalBarCharacter)
}

func TestBarChartPrinter_WithVerticalBarCharacter(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := "X"
	p2 := p.WithVerticalBarCharacter(s)

	assert.Equal(t, s, p2.VerticalBarCharacter)
	assert.Empty(t, p.VerticalBarCharacter)
}

func TestBarChartPrinter_WithBars(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}
	p2 := p.WithBars(s)

	assert.Equal(t, s, p2.Bars)
	assert.Empty(t, p.Bars)
}

func TestBarChartPrinter_WithHeight(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := 1337
	p2 := p.WithHeight(s)

	assert.Equal(t, s, p2.Height)
	assert.Empty(t, p.Height)
}

func TestBarChartPrinter_WithHorizontal(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := true
	p2 := p.WithHorizontal(s)

	assert.Equal(t, s, p2.Horizontal)
	assert.Empty(t, p.Horizontal)
}

func TestBarChartPrinter_WithShowValue(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := true
	p2 := p.WithShowValue(s)

	assert.Equal(t, s, p2.ShowValue)
	assert.Empty(t, p.ShowValue)
}

func TestBarChartPrinter_WithWidth(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := 1337
	p2 := p.WithWidth(s)

	assert.Equal(t, s, p2.Width)
	assert.Empty(t, p.Width)
}
