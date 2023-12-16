package pterm_test

import (
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestHeatmapPrinter_NilPrint(t *testing.T) {
	p := pterm.HeatmapPrinter{}
	p.Render()
}

func TestHeatmapPrinter_SRender(t *testing.T) {
	d := [][]float32{
		{-1, -0.9, -0.8},
		{-1, -0.9, -0.8},
		{-1, -0.9, -0.8},
	}

	hd := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c"},
		YAxis: []string{"1", "2", "3"},
	}
	// WithGrid
	printer := pterm.DefaultHeatmap.WithAxisData(hd).WithData(d)
	content, err := printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)

	// WithoutGrid
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithGrid(false)
	content, err = printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)

	// WithColouredCells
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithOnlyColoredCells()
	content, err = printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)

	// WithoutStyle
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithAxisStyle(nil)
	content, err = printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)

	// WithoutSeparatorStyle
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithSeparatorStyle(nil)
	content, err = printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)

	// WithEnableRGB
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithEnableRGB(true)
	content, err = printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
}

func TestHeatmapPrinter_WithAxisData(t *testing.T) {
	hd := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c"},
		YAxis: []string{"1", "2", "3"},
	}

	h := pterm.DefaultHeatmap.WithAxisData(hd)

	testza.AssertTrue(t, h.HasHeader)
	testza.AssertEqual(t, hd, h.Axis)
}

func TestHeatmapPrinter_WithAxisStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.HeatmapPrinter{}
	p2 := p.WithAxisStyle(s)

	testza.AssertEqual(t, s, p2.AxisStyle)
}

func TestHeatmapPrinter_WithSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	h := pterm.HeatmapPrinter{}
	h2 := h.WithSeparatorStyle(s)

	testza.AssertEqual(t, s, h2.SeparatorStyle)
}

func TestHeatmapPrinter_WithData(t *testing.T) {
	proxyToDevNull()
	d := [][]float32{
		{-1, -0.9, -0.8, -0.7, -0.6},
		{-1, -0.9, -0.8, -0.7, -0.6},
	}

	h := pterm.HeatmapPrinter{}
	h2 := h.WithData(d)

	testza.AssertEqualValues(t, d, h2.Data)
}

func TestHeatmapPrinter_WithBoxed(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithBoxed(true)

	testza.AssertTrue(t, h2.Boxed)
}

func TestHeatmapPrinter_WithGrid(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithGrid(true)

	testza.AssertTrue(t, h2.Grid)
}

func TestHeatmapPrinter_WithoutGrid(t *testing.T) {
	h := pterm.DefaultHeatmap
	h2 := h.WithGrid(false)

	testza.AssertFalse(t, h2.Grid)
	testza.AssertFalse(t, h2.Boxed)
}

func TestHeatmapPrinter_WithRGB(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithEnableRGB(true)

	testza.AssertTrue(t, h2.EnableRGB)
}

func TestHeatmapPrinter_WithOnlyColoredCells(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithOnlyColoredCells(true)

	testza.AssertTrue(t, h2.OnlyColoredCells)
}

func TestHeatmapPrinter_WithCellSize(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithCellSize(1)

	testza.AssertEqual(t, 1, h2.CellSize)
}

func TestHeatmapPrinter_WithWriter(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	s := os.Stderr
	h2 := h.WithWriter(s)

	testza.AssertEqual(t, s, h2.Writer)
	testza.AssertZero(t, h.Writer)
}
