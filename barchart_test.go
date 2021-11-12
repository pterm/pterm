package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
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

	printer := pterm.DefaultBarChart.WithBars(bars)
	err := printer.Render()
	if err != nil {
		panic(err)
	}
	// TODO enable the snapshot once CI == local dev
	// content, _ := printer.Srender()
	// testza.SnapshotCreateOrValidate(t, t.Name(), content)
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

	printer := pterm.DefaultBarChart.WithBars(bars)
	printer.Render()
	// TODO enable the snapshot once CI == local dev
	// content, _ := printer.Srender()
	// testza.SnapshotCreateOrValidate(t, t.Name(), content)
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

	printer := pterm.DefaultBarChart.WithBars(bars)
	printer.Render()
	content, _ := printer.Srender()
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
	pterm.EnableStyling()
}

func TestBarChartPrinter_RenderMultipleLineLabel(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
	})

	printer.Render()
	content, _ := printer.Srender()
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_RenderNegativeBarValues(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
	})
	printer.Render()
	content, _ := printer.Srender()
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_RenderZeroBarValuesHorizontal(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
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
	})
	printer.Render()
	content, _ := printer.Srender()
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_RenderZeroBarValues(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
	})
	printer.Render()
	content, _ := printer.Srender()
	testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_RenderLowBarValues(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
	})
	printer.Render()
	// TODO enable the snapshot once CI == local dev
	// content, _ := printer.Srender()
	// testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_Render(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithBars(pterm.Bars{
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
	})
	printer.Render()
	// TODO enable the snapshot once CI == local dev
	// content, _ := printer.Srender()
	// testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_RenderHorizonzal(t *testing.T) {
	printer := pterm.DefaultBarChart.WithShowValue().WithHorizontal().WithBars(pterm.Bars{
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
	})
	printer.Render()
	// TODO enable the snapshot once CI == local dev
	// content, _ := printer.Srender()
	// testza.SnapshotCreateOrValidate(t, t.Name(), content)
}

func TestBarChartPrinter_WithHorizontalBarCharacter(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := "X"
	p2 := p.WithHorizontalBarCharacter(s)

	testza.AssertEqual(t, s, p2.HorizontalBarCharacter)
	testza.AssertZero(t, p.HorizontalBarCharacter)
}

func TestBarChartPrinter_WithVerticalBarCharacter(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := "X"
	p2 := p.WithVerticalBarCharacter(s)

	testza.AssertEqual(t, s, p2.VerticalBarCharacter)
	testza.AssertZero(t, p.VerticalBarCharacter)
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

	testza.AssertEqual(t, s, p2.Bars)
	testza.AssertZero(t, p.Bars)
}

func TestBarChartPrinter_WithHeight(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := 1337
	p2 := p.WithHeight(s)

	testza.AssertEqual(t, s, p2.Height)
	testza.AssertZero(t, p.Height)
}

func TestBarChartPrinter_WithHorizontal(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := true
	p2 := p.WithHorizontal(s)

	testza.AssertEqual(t, s, p2.Horizontal)
	testza.AssertZero(t, p.Horizontal)
}

func TestBarChartPrinter_WithShowValue(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := true
	p2 := p.WithShowValue(s)

	testza.AssertEqual(t, s, p2.ShowValue)
	testza.AssertZero(t, p.ShowValue)
}

func TestBarChartPrinter_WithWidth(t *testing.T) {
	p := pterm.BarChartPrinter{}
	s := 1337
	p2 := p.WithWidth(s)

	testza.AssertEqual(t, s, p2.Width)
	testza.AssertZero(t, p.Width)
}
