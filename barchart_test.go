package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestBarChartPrinterNilPrint(t *testing.T) {
	proxyToDevNull()
	p := pterm.BarChartPrinter{}
	err := p.Render()
	if err != nil {
		panic(err)
	}
}

func TestBarChartPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.BarChartPrinter{})
}

func TestBarChartPrinter_NilStylePrint(t *testing.T) {
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

	pterm.DefaultBarChart.WithBars(bars).Render()
}

// VERTICAL bars + MIXED values test
func TestBarChartPrinter_RenderExample(t *testing.T) {
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

	pterm.DefaultBarChart.WithBars(bars).Render()
}

// VERTICAL bars + NEGATIVE values test
func TestBarChartPrinter_RenderNegativeBarValues(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
func TestBarChartPrinter_RenderPositiveBarValues(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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

// HORIZONTAL bars + MIXED values test
func TestBarChartPrinter_RenderExampleHorizontal(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
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
func TestBarChartPrinter_RenderNegativeBarValuesHorizontal(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
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
func TestBarChartPrinter_RenderPositiveBarValuesHorizontal(t *testing.T) {
	pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
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

func TestBarChartPrinter_RenderExampleRawOutput(t *testing.T) {
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

	pterm.DefaultBarChart.WithBars(bars).Render()
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
