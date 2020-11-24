package pterm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarChartPrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	p := BarChartPrinter{}
	err := p.Render()
	if err != nil {
		panic(err)
	}
}

func TestBarChartPrinter_RenderExample(t *testing.T) {
	proxyToDevNull()
	bars := Bars{
		Bar{
			Label: "Bar 1",
			Value: 5,
			Style: NewStyle(FgCyan),
		},
		Bar{
			Label: "Bar 2",
			Value: 3,
			Style: NewStyle(FgCyan),
		},
		Bar{
			Label: "Longer Label",
			Value: 7,
			Style: NewStyle(FgCyan),
		},
	}

	_ = DefaultBarChart.WithBars(bars).Render()
}

func TestBarChartPrinter_RenderMultipleLineLabel(t *testing.T) {
	DefaultBarChart.WithShowValue().WithBars(Bars{
		Bar{
			Label: "Test",
			Value: -1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test\nNew Line",
			Value: -1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: -1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderNegativeBarValues(t *testing.T) {
	DefaultBarChart.WithShowValue().WithBars(Bars{
		Bar{
			Label: "Test",
			Value: -1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: -1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderZeroBarValues(t *testing.T) {
	DefaultBarChart.WithShowValue().WithBars(Bars{
		Bar{
			Label: "Test",
			Value: 0,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: 0,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderLowBarValues(t *testing.T) {
	DefaultBarChart.WithShowValue().WithBars(Bars{
		Bar{
			Label: "Test",
			Value: 1,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: 1,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}).Render()
}

func TestBarChartPrinter_Render(t *testing.T) {
	DefaultBarChart.WithShowValue().WithBars(Bars{
		Bar{
			Label: "Test",
			Value: 1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: 1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}).Render()
}

func TestBarChartPrinter_RenderHorizonzal(t *testing.T) {
	DefaultBarChart.WithShowValue().WithHorizontal().WithBars(Bars{
		Bar{
			Label: "Test",
			Value: 1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: 1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}).Render()
}

func TestBarChartPrinter_WithHorizontalBarCharacter(t *testing.T) {
	p := BarChartPrinter{}
	s := "X"
	p2 := p.WithHorizontalBarCharacter(s)

	assert.Equal(t, s, p2.HorizontalBarCharacter)
	assert.Empty(t, p.HorizontalBarCharacter)
}

func TestBarChartPrinter_WithVerticalBarCharacter(t *testing.T) {
	p := BarChartPrinter{}
	s := "X"
	p2 := p.WithVerticalBarCharacter(s)

	assert.Equal(t, s, p2.VerticalBarCharacter)
	assert.Empty(t, p.VerticalBarCharacter)
}

func TestBarChartPrinter_WithBars(t *testing.T) {
	p := BarChartPrinter{}
	s := Bars{
		Bar{
			Label: "Test",
			Value: 1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
		Bar{
			Label: "Test",
			Value: 1337,
			Style: NewStyle(FgRed, BgBlue, Bold),
		},
	}
	p2 := p.WithBars(s)

	assert.Equal(t, s, p2.Bars)
	assert.Empty(t, p.Bars)
}

func TestBarChartPrinter_WithHeight(t *testing.T) {
	p := BarChartPrinter{}
	s := 1337
	p2 := p.WithHeight(s)

	assert.Equal(t, s, p2.Height)
	assert.Empty(t, p.Height)
}

func TestBarChartPrinter_WithHorizontal(t *testing.T) {
	p := BarChartPrinter{}
	s := true
	p2 := p.WithHorizontal(s)

	assert.Equal(t, s, p2.Horizontal)
	assert.Empty(t, p.Horizontal)
}

func TestBarChartPrinter_WithShowValue(t *testing.T) {
	p := BarChartPrinter{}
	s := true
	p2 := p.WithShowValue(s)

	assert.Equal(t, s, p2.ShowValue)
	assert.Empty(t, p.ShowValue)
}

func TestBarChartPrinter_WithWidth(t *testing.T) {
	p := BarChartPrinter{}
	s := 1337
	p2 := p.WithWidth(s)

	assert.Equal(t, s, p2.Width)
	assert.Empty(t, p.Width)
}
