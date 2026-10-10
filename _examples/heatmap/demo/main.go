package main

import "github.com/pterm/pterm"

func main() {
	// Commits per weekday, in two-hour slots. Each row is one day, and the
	// colors run from Border through Primary to Accent.
	data := pterm.HeatmapData{
		{3, 6, 4, 7, 12, 24, 34, 32, 28, 12, 3, 6},
		{5, 3, 6, 4, 14, 26, 31, 34, 25, 14, 5, 3},
		{7, 5, 3, 6, 11, 28, 33, 31, 27, 11, 7, 5},
		{4, 7, 5, 3, 13, 25, 35, 33, 24, 13, 4, 7},
		{6, 4, 7, 5, 10, 27, 32, 35, 26, 10, 6, 4},
		{1, 4, 2, 5, 5, 6, 11, 9, 10, 5, 1, 4},
		{3, 1, 4, 2, 7, 8, 8, 11, 7, 7, 3, 1},
	}

	// One X label per column and one Y label per row.
	axis := pterm.HeatmapAxis{
		XAxis: []string{"00", "02", "04", "06", "08", "10", "12", "14", "16", "18", "20", "22"},
		YAxis: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
	}

	pterm.Heatmap.WithData(data).WithAxisData(axis).WithLegendLabel("commits").Render()
}
