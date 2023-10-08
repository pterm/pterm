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

	pterm.Println("The following table has no rgb (supported by every terminal), no axis data and a legend.\n")

	table := pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithLegend(false).WithColors([]pterm.Color{pterm.BgBlue, pterm.BgRed, pterm.BgGreen, pterm.BgYellow}...).WithLegend()
	table.Render()

	pterm.Println("The following table has rgb (not supported by every terminal), axis data and a legend.\n")
	table2 := pterm.DefaultHeatmap.WithData(data).WithBoxed(false).WithAxisData(headerData).WithRGBRange(pterm.NewRGB(0, 0, 255), pterm.NewRGB(255, 0, 0), pterm.NewRGB(0, 255, 0), pterm.NewRGB(255, 255, 0))
	table2.Render()
}
