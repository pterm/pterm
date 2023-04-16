package main

import (
	"github.com/pterm/pterm"
)

func main() {
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3, -0.6, 0.5, -0.2, 0.9, 0.1, -0.9, 0.2, -0.7},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3, -0.1, 0.1, 0.6, 0.7, 1.0, -0.9, -0.6, -0.5},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9, 0.4, -0.3, -0.7, 0.6, -0.2, -0.9, 0.7, 0.0},
		{0.9, -0.5, -0.1, 0.3, 0.9, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6, 1.0, -0.6, 0.0, 0.4, 0.8, -0.6, -0.2, 0.2},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0, -0.5, 1.0, 0.7, 0.2, -0.1, -0.6, 0.9, 0.6},
		{0.4, -0.3, 0.3, -0.9, -0.3, 0.6, -0.9, -0.3, 0.6, -0.9, 0.0, 0.6, -0.6, 0.0, 0.6, -0.6, 0.0, 0.9, -0.6, 0.2, 0.9},
		{0.1, 0.8, 0.5, 0.5, 0.1, 0.1, -0.3, -0.3, -0.6, -0.7, 1.0, 0.9, 0.6, 0.5, 0.2, 0.1, -0.2, -0.2, -0.6, -0.6, -0.9},
		{0.9, -0.1, 0.9, -0.3, 0.7, -0.5, 0.6, -0.6, 0.4, -0.6, 0.2, -0.7, 0.0, -1.0, -0.2, 0.9, -0.1, 0.7, -0.3, 0.5, -0.5},
		{-0.9, -1.0, -0.9, -1.0, -0.9, -0.7, -0.9, -0.7, -0.6, -0.7, -0.6, -0.5, -0.6, -0.5, -0.6, -0.5, -0.3, -0.4, -0.3, -0.2, -0.3},
		{-0.3, 0.1, -0.7, 0.6, -0.5, 0.7, 0.0, 1.0, 0.2, -0.6, 0.6, -0.4, 0.9, 0.1, -0.9, 0.3, -0.5, 0.7, -0.3, 0.9, 0.2},
	}

	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u"},
		// XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"},
		YAxis: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		// YAxis: []string{"1", "2", "3", "4"},
	}

	table := pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithRGB().WithGrid(false)
	table.Render()

	table = pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithGrid(false)
	table.Render()

	table = pterm.DefaultHeatmap.WithData(data).WithBoxed().WithAxisData(headerData).WithRGB()
	table.Render()

	table = pterm.DefaultHeatmap.WithData(data).WithBoxed().WithAxisData(headerData).WithRGB().WithLegendOnlyColoredCells()
	table.Render()

	table = pterm.DefaultHeatmap.WithData(data).WithBoxed().WithAxisData(headerData)
	table.Render()

	table3 := pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithRGB()
	table3.Render()

	table2 := pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithRGB()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithBoxed().WithRGB()
	table2.Render()

	table = pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithRGB().WithGrid(false)
	table.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithRGB().WithAxisData(headerData).WithOnlyColoredCells().WithCellSize(3)
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithRGB().WithOnlyColoredCells().WithCellSize(3)
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithRGB().WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithRGB().WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells().WithLegendOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells().WithLegendOnlyColoredCells()
	table2.Render()
}
