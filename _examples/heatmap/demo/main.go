package main

import (
	"github.com/pterm/pterm"
)

func main() {
	data := [][]float32{
		{0.9, 0.2, -0.7, 0.4, -0.5, 0.6, -0.3, 0.8, -0.1, -1.0, 0.1, -0.8, 0.3},
		{0.2, -0.7, -0.5, -0.3, -0.1, 0.1, 0.3, 0.5, 0.9, -0.9, -0.7, -0.5, -0.3},
		{0.4, 0.4, -0.3, -1.0, 0.3, -0.2, -0.9, 0.5, -0.3, -1.0, 0.6, -0.2, -0.9},
		{0.9, -0.5, -0.1, 0.3, 1, -0.7, -0.3, 0.1, 0.7, -0.9, -0.5, 0.2, 0.6},
		{0.5, 0.6, 0.1, -0.2, -0.7, 0.8, 0.6, 0.1, -0.5, -0.7, 0.7, 0.3, 0.0},
	}

	headerData := pterm.HeatmapAxis{
		XAxis: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"},
		YAxis: []string{"1", "2", "3", "4", "5"},
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

	table2 = pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithRGB().WithOnlyColoredCells().WithCellSize(5)
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithRGB().WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithRGB().WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells().WithLegendOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells().WithLegendOnlyColoredCells()
	table2.Render()

	table2 = pterm.DefaultHeatmap.WithData(data).WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells().WithLegend(false)
	table2.Render()

	table15 := pterm.DefaultHeatmap.WithData(data).WithCellSize(3).WithGrid(false).WithAxisData(headerData).WithOnlyColoredCells().WithLegendOnlyColoredCells().WithBoxed()
	table15.Render()

	table4 := pterm.DefaultHeatmap.WithData(data).WithBoxed().WithAxisData(headerData).WithRGB().WithLegendOnlyColoredCells().WithLegendTag("custom tag")
	table4.Render()

	table5 := pterm.DefaultHeatmap.WithData(data)
	table5.Render()
}
