package pterm_test

import (
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestBarChartPrinterNilPrint(_ *testing.T) {
	proxyToDevNull()

	p := pterm.BarChartPrinter{}

	err := p.Render()
	if err != nil {
		panic(err)
	}
}

func TestBarChartPrinter_NilStylePrint(_ *testing.T) {
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

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
}

// VERTICAL bars + MIXED values test
func TestBarChartPrinter_RenderExample(_ *testing.T) {
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
			Label: "Long Label Example",
			Value: 7,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Zero",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Negative Value",
			Value: -4,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "NV",
			Value: -5,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
	}

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
}

// VERTICAL bars + NEGATIVE values test
func TestBarChartPrinter_RenderNegativeBarValues(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: -1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -1000,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -950,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -1500,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -10,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -100,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

// VERTICAL bars + POSITIVE values test
func TestBarChartPrinter_RenderPositiveBarValues(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 1000,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 1400,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 900,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

// VERTICAL bars + ZERO values test
func TestBarChartPrinter_RenderZeroBarValues(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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

// HORIZONTAL bars + MIXED values test
func TestBarChartPrinter_RenderExampleHorizontal(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 1337,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 1000,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Zero",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -800,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -500,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

// HORIZONTAL bars + NEGATIVE values test
func TestBarChartPrinter_RenderNegativeBarValuesHorizontal(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: -999,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -500,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -653,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -20,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -100,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: -30,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

// HORIZONTAL bars + POSITIVE values test
func TestBarChartPrinter_RenderPositiveBarValuesHorizontal(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
		pterm.Bar{
			Label: "Test",
			Value: 30,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 70,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 80,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 90,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 40,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
		pterm.Bar{
			Label: "Test",
			Value: 30,
			Style: pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold),
		},
	}).Render()
}

// HORIZONTAL bars + ZERO values test
func TestBarChartPrinter_RenderZeroBarValuesHorizontal(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
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

func TestBarChartPrinter_RenderExampleRawOutput(_ *testing.T) {
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
			Label: "Long Label Example",
			Value: 7,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Zero",
			Value: 0,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Negative Value",
			Value: -4,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "NV",
			Value: -5,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
	}

	_ = pterm.DefaultBarChart.WithBars(bars).Render()

	pterm.EnableStyling()
}

func TestBarChartPrinter_RenderMultipleLineLabel(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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

func TestBarChartPrinter_RenderLowBarValues(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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

func TestBarChartPrinter_Render(_ *testing.T) {
	_ = pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
	assert.Zero(t, p.HorizontalBarCharacter)
}

func TestBarChartPrinter_WithVerticalBarCharacter(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := "X"
	p2 := p.WithVerticalBarCharacter(s)

	assert.Equal(t, s, p2.VerticalBarCharacter)
	assert.Zero(t, p.VerticalBarCharacter)
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
	assert.Zero(t, p.Bars)
}

func TestBarChartPrinter_WithHeight(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := 1337
	p2 := p.WithHeight(s)

	assert.Equal(t, s, p2.Height)
	assert.Zero(t, p.Height)
}

func TestBarChartPrinter_WithHorizontal(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := true
	p2 := p.WithHorizontal(s)

	assert.Equal(t, s, p2.Horizontal)
	assert.Zero(t, p.Horizontal)
}

func TestBarChartPrinter_WithShowValue(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := true
	p2 := p.WithShowValue(s)

	assert.Equal(t, s, p2.ShowValue)
	assert.Zero(t, p.ShowValue)
}

func TestBarChartPrinter_WithWidth(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := 1337
	p2 := p.WithWidth(s)

	assert.Equal(t, s, p2.Width)
	assert.Zero(t, p.Width)
}

func TestBarChartPrinter_WithWriter(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}
