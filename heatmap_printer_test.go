package pterm_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestHeatmapPrinter_NilPrint(_ *testing.T) {
	p := pterm.HeatmapPrinter{}
	_ = p.Render()
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

	assert.NoError(t, err)
	assert.NotNil(t, content)

	// WithoutGrid
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithGrid(false)
	content, err = printer.Srender()

	assert.NoError(t, err)
	assert.NotNil(t, content)

	// WithColouredCells
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithOnlyColoredCells()
	content, err = printer.Srender()

	assert.NoError(t, err)
	assert.NotNil(t, content)

	// WithoutStyle
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithAxisStyle(nil)
	content, err = printer.Srender()

	assert.NoError(t, err)
	assert.NotNil(t, content)

	// WithoutSeparatorStyle
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithSeparatorStyle(nil)
	content, err = printer.Srender()

	assert.NoError(t, err)
	assert.NotNil(t, content)

	// WithEnableRGB
	printer = pterm.DefaultHeatmap.WithAxisData(hd).WithData(d).WithEnableRGB(true)
	content, err = printer.Srender()

	assert.NoError(t, err)
	assert.NotNil(t, content)
}

func TestHeatmapPrinter_WithAxisData(t *testing.T) {
	hd := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c"},
		YAxis: []string{"1", "2", "3"},
	}

	h := pterm.DefaultHeatmap.WithAxisData(hd)

	assert.True(t, h.HasHeader)
	assert.Equal(t, hd, h.Axis)
}

func TestHeatmapPrinter_WithAxisStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.HeatmapPrinter{}
	p2 := p.WithAxisStyle(s)

	assert.Equal(t, s, p2.AxisStyle)
}

func TestHeatmapPrinter_WithSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	h := pterm.HeatmapPrinter{}
	h2 := h.WithSeparatorStyle(s)

	assert.Equal(t, s, h2.SeparatorStyle)
}

func TestHeatmapPrinter_WithData(t *testing.T) {
	proxyToDevNull()

	d := [][]float32{
		{-1, -0.9, -0.8, -0.7, -0.6},
		{-1, -0.9, -0.8, -0.7, -0.6},
	}

	h := pterm.HeatmapPrinter{}
	h2 := h.WithData(d)

	assert.EqualValues(t, d, h2.Data)
}

func TestHeatmapPrinter_WithBoxed(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithBoxed(true)

	assert.True(t, h2.Boxed)
}

func TestHeatmapPrinter_WithGrid(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithGrid(true)

	assert.True(t, h2.Grid)
}

func TestHeatmapPrinter_WithoutGrid(t *testing.T) {
	h := pterm.DefaultHeatmap
	h2 := h.WithGrid(false)

	assert.False(t, h2.Grid)
	assert.False(t, h2.Boxed)
}

func TestHeatmapPrinter_WithRGB(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithEnableRGB(true)

	assert.True(t, h2.EnableRGB)
}

func TestHeatmapPrinter_WithOnlyColoredCells(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithOnlyColoredCells(true)

	assert.True(t, h2.OnlyColoredCells)
}

func TestHeatmapPrinter_WithCellSize(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	h2 := h.WithCellSize(1)

	assert.Equal(t, 1, h2.CellSize)
}

func TestHeatmapPrinter_WithWriter(t *testing.T) {
	h := pterm.HeatmapPrinter{}
	s := os.Stderr
	h2 := h.WithWriter(s)

	assert.Equal(t, s, h2.Writer)
	assert.Zero(t, h.Writer)
}
