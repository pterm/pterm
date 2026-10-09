package pterm_test

// Behavioral tests for HeatmapPrinter: grid geometry, axis label placement,
// legend layout, color bucketing (min/max colors, complementary text colors),
// RGB fading and input validation. The builder/contract plumbing is covered
// in contract_test.go, one representative output in snapshot_test.go.

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func heatmapAxis() pterm.HeatmapAxis {
	return pterm.HeatmapAxis{
		XAxis: []string{"a", "b"},
		YAxis: []string{"c", "d"},
	}
}

func TestHeatmapPrinter_GridAxisAndLegendLayout(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithAxisData(heatmapAxis()).WithData(pterm.HeatmapData{
		{0, 1},
		{2, 3},
	})

	// One grid cell per data value, Y axis labels in the first column, X axis
	// labels in the last row, and a boxed legend from min (0) to max (3).
	expected := "" +
		"╭─┬─┬─╮\n" +
		"│c│0│1│\n" +
		"├─┼─┼─┤\n" +
		"│d│2│3│\n" +
		"├─┼─┼─┤\n" +
		"│ │a│b│\n" +
		"╰─┴─┴─╯\n" +
		"\n" +
		"╭──────┬───┬───┬───┬───┬───┬───╮\n" +
		"│Legend│ 0 │0.6│1.2│1.8│2.4│ 3 │\n" +
		"╰──────┴───┴───┴───┴───┴───┴───╯\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestHeatmapPrinter_WithLegendFalseRemovesLegend(t *testing.T) {
	data := pterm.HeatmapData{{0, 1}, {2, 3}}

	withLegend := srenderPlain(t, pterm.DefaultHeatmap.WithAxisData(heatmapAxis()).WithData(data))
	withoutLegend := srenderPlain(t, pterm.DefaultHeatmap.WithAxisData(heatmapAxis()).WithData(data).WithLegend(false))

	assert.Contains(t, withLegend, "Legend")
	assert.NotContains(t, withoutLegend, "Legend")
	assert.True(t, strings.HasPrefix(withLegend, strings.TrimSuffix(withoutLegend, "\n")),
		"the grid itself must not change when the legend is removed")
}

func TestHeatmapPrinter_WithLegendLabel(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithLegendLabel("Scale").WithData(pterm.HeatmapData{{0, 1}})

	out := srenderPlain(t, printer)
	assert.Contains(t, out, "Scale")
	assert.NotContains(t, out, "Legend")
}

func TestHeatmapPrinter_OnlyColoredCellsRenderCellSizedBlanks(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithOnlyColoredCells().WithCellSize(3).WithLegend(false).WithData(pterm.HeatmapData{
		{0, 1},
		{2, 3},
	})

	// Without a header the cells are CellSize wide and contain no values.
	expected := "" +
		"╭───┬───╮\n" +
		"│   │   │\n" +
		"├───┼───┤\n" +
		"│   │   │\n" +
		"╰───┴───╯\n"

	assert.Equal(t, expected, srenderPlain(t, printer))

	// In styled mode the cells carry their bucket color as background:
	// 0 -> BgRed, 1 -> BgYellow, 2 -> BgLightGreen, 3 -> BgGreen.
	styled, err := printer.Srender()
	require.NoError(t, err)

	for _, code := range []string{"\x1b[41m", "\x1b[43m", "\x1b[102m", "\x1b[42m"} {
		assert.Contains(t, styled, code+"   ", "each cell must be painted with its bucket's background color")
	}
}

// Regression test: a refactor dropped the background color from the non-RGB
// heatmap cells, rendering all cells with the text color only.
func TestHeatmapPrinter_MinAndMaxValuesGetFirstAndLastColor(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithLegend(false).WithData(pterm.HeatmapData{{0, 3}})

	styled, err := printer.Srender()
	require.NoError(t, err)

	// DefaultHeatmap colors range from BgRed (41) for the minimum to BgGreen
	// (42) for the maximum.
	assert.Contains(t, styled, "\x1b[41m0", "the minimum value must get the first color")
	assert.Contains(t, styled, "\x1b[42m3", "the maximum value must get the last color")
}

func TestHeatmapPrinter_ComplementaryTextColorPerBucket(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithEnableComplementaryColor().WithLegend(false).WithData(pterm.HeatmapData{{0, 3}})

	styled, err := printer.Srender()
	require.NoError(t, err)

	// BgRed cells get FgCyan (36) text, BgGreen cells FgMagenta (35).
	assert.Contains(t, styled, "\x1b[36m\x1b[41m0")
	assert.Contains(t, styled, "\x1b[35m\x1b[42m3")
}

func TestHeatmapPrinter_LegendUsesTheBucketColors(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithData(pterm.HeatmapData{{0, 3}})

	styled, err := printer.Srender()
	require.NoError(t, err)

	legendStart := strings.Index(styled, "Legend")
	require.GreaterOrEqual(t, legendStart, 0)

	legend := styled[legendStart:]
	assert.Contains(t, legend, "\x1b[41m", "the legend must paint the first color")
	assert.Contains(t, legend, "\x1b[42m", "the legend must paint the last color")
}

func TestHeatmapPrinter_RGBFadesFromFirstToLastColor(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithEnableRGB().WithLegend(false).WithData(pterm.HeatmapData{{0, 1}})

	styled, err := printer.Srender()
	require.NoError(t, err)

	// DefaultHeatmap fades from RGB(255,0,0) to RGB(0,255,0) backgrounds.
	assert.Contains(t, styled, "48;2;255;0;0", "the minimum value must get the first RGB color")
	assert.Contains(t, styled, "48;2;0;255;0", "the maximum value must get the last RGB color")
}

func TestHeatmapPrinter_WithoutGridRendersBareCells(t *testing.T) {
	printer := pterm.DefaultHeatmap.WithGrid(false).WithLegend(false).WithData(pterm.HeatmapData{{0, 1}})

	// Documents current behavior: without grid (and therefore without box)
	// the cells are rendered back-to-back.
	assert.Equal(t, "01\n", srenderPlain(t, printer))
}

func TestHeatmapPrinter_InputValidation(t *testing.T) {
	tests := []struct {
		name    string
		printer *pterm.HeatmapPrinter
		wantErr string
	}{
		{
			name:    "nil data",
			printer: &pterm.HeatmapPrinter{},
			wantErr: "data is nil",
		},
		{
			name:    "empty data",
			printer: pterm.DefaultHeatmap.WithData(pterm.HeatmapData{}),
			wantErr: "data is empty",
		},
		{
			name:    "non-rectangular data",
			printer: pterm.DefaultHeatmap.WithData(pterm.HeatmapData{{1, 2}, {3}}),
			wantErr: "data is not rectangular",
		},
		{
			name:    "x axis nil",
			printer: pterm.DefaultHeatmap.WithAxisData(pterm.HeatmapAxis{YAxis: []string{"y"}}).WithData(pterm.HeatmapData{{1}}),
			wantErr: "x axis is nil",
		},
		{
			name:    "y axis nil",
			printer: pterm.DefaultHeatmap.WithAxisData(pterm.HeatmapAxis{XAxis: []string{"x"}}).WithData(pterm.HeatmapData{{1}}),
			wantErr: "y axis is nil",
		},
		{
			name: "x axis length mismatch",
			printer: pterm.DefaultHeatmap.WithAxisData(pterm.HeatmapAxis{
				XAxis: []string{"a", "b"},
				YAxis: []string{"c", "d"},
			}).WithData(pterm.HeatmapData{{1, 2}, {3}}),
			wantErr: "x axis length does not match data",
		},
		{
			name: "y axis length mismatch",
			printer: pterm.DefaultHeatmap.WithAxisData(pterm.HeatmapAxis{
				XAxis: []string{"a"},
				YAxis: []string{"c", "d"},
			}).WithData(pterm.HeatmapData{{1}}),
			wantErr: "y axis length does not match data",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.printer.Srender()
			assert.EqualError(t, err, tc.wantErr)

			// Render must propagate the same error instead of printing.
			assert.EqualError(t, tc.printer.Render(), tc.wantErr)
		})
	}
}

func TestHeatmapPrinter_SrenderIsPure(t *testing.T) {
	data := pterm.HeatmapData{
		{0, 1},
		{2, 3},
	}
	printer := pterm.DefaultHeatmap.WithAxisData(heatmapAxis()).WithData(data)

	first, err := printer.Srender()
	require.NoError(t, err)

	second, err := printer.Srender()
	require.NoError(t, err)

	assert.Equal(t, first, second, "rendering twice must yield identical output")
	assert.Equal(t, pterm.HeatmapData{{0, 1}, {2, 3}}, data, "rendering must not modify the input data")
}
