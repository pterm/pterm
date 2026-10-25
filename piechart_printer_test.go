package pterm_test

// Behavioral tests for PieChartPrinter.
//
// Builder methods and the Render/Srender contract are covered generically in
// contract_test.go; the styled default output is locked by snapshot tests.
// This file verifies the geometry, the legend and input validation.

import (
	"math"
	"strings"
	"testing"

	"github.com/mattn/go-runewidth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func TestPieChartPrinterPieIsSquareInCells(t *testing.T) {
	out, err := pterm.DefaultPieChart.WithRadius(3).WithShowPercentage(false).WithSlices(pterm.PieSlices{
		{Label: "a", Value: 1},
	}).Srender()
	require.NoError(t, err)

	lines := strings.Split(strings.TrimSuffix(stripANSI(out), "\n"), "\n")
	require.Len(t, lines, 7, "a radius of 3 spans 2*3+1 rows")

	// The middle row runs through the center, so it is the full diameter:
	// seven cells, two columns each, followed by the legend.
	assert.True(t, strings.HasPrefix(lines[3], strings.Repeat("█", 14)+"   "), "got %q", lines[3])
}

func TestPieChartPrinterFirstSliceStartsAtTwelveOClock(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableColor()

	// Two equal halves: the first fills the right half, the second the left.
	out, err := pterm.DefaultPieChart.WithRadius(3).WithSlices(pterm.PieSlices{
		{Label: "right", Value: 1},
		{Label: "left", Value: 1},
	}).Srender()
	require.NoError(t, err)

	middle := strings.Split(out, "\n")[3]
	assert.True(t, strings.HasPrefix(middle, strings.Repeat("▓", 6)+strings.Repeat("█", 8)), "got %q", middle)
}

func TestPieChartPrinterLegendAlignsLabelsAndShowsShares(t *testing.T) {
	out, err := pterm.DefaultPieChart.WithShowValue().WithSlices(pterm.PieSlices{
		{Label: "Go", Value: 3},
		{Label: "TypeScript", Value: 1},
	}).Srender()
	require.NoError(t, err)

	plain := stripANSI(out)
	assert.Contains(t, plain, "██ Go          3 (75.0%)\n")
	assert.Contains(t, plain, "██ TypeScript  1 (25.0%)\n")
}

func TestPieChartPrinterLegendWithoutDetails(t *testing.T) {
	out, err := pterm.DefaultPieChart.WithShowPercentage(false).WithSlices(pterm.PieSlices{
		{Label: "a", Value: 1},
		{Label: "bbb", Value: 1},
	}).Srender()
	require.NoError(t, err)

	plain := stripANSI(out)
	assert.Contains(t, plain, "██ a\n", "labels must not carry trailing padding")
	assert.NotContains(t, plain, "%")
}

func TestPieChartPrinterNoColorUsesDistinctCharacters(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableColor()

	out, err := pterm.DefaultPieChart.WithSlices(pterm.PieSlices{
		{Label: "a", Value: 1},
		{Label: "b", Value: 1},
		{Label: "c", Value: 1},
	}).Srender()
	require.NoError(t, err)

	assert.Contains(t, out, "██ a")
	assert.Contains(t, out, "▓▓ b")
	assert.Contains(t, out, "▒▒ c")
}

func TestPieChartPrinterLinesHaveNoTrailingSpaces(t *testing.T) {
	out, err := pterm.DefaultPieChart.WithSlices(pterm.PieSlices{
		{Label: "a", Value: 1},
		{Label: "b", Value: 2},
	}).Srender()
	require.NoError(t, err)

	for i, line := range strings.Split(stripANSI(out), "\n") {
		assert.Equalf(t, strings.TrimRight(line, " "), line, "line %d has trailing spaces", i)
		assert.LessOrEqualf(t, runewidth.StringWidth(line), 80, "line %d is wider than the terminal", i)
	}
}

func TestPieChartPrinterUsesSliceStyle(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.EnableStyling()

	custom := pterm.NewStyle(pterm.FgRed)

	out, err := pterm.DefaultPieChart.WithSlices(pterm.PieSlices{
		{Label: "a", Value: 1, Style: custom},
	}).Srender()
	require.NoError(t, err)

	assert.Contains(t, out, custom.Sprint("██"))
}

func TestPieChartPrinterRejectsInvalidData(t *testing.T) {
	tests := map[string]pterm.PieSlices{
		"no slices":  nil,
		"all zero":   {{Label: "a", Value: 0}},
		"negative":   {{Label: "a", Value: 2}, {Label: "b", Value: -1}},
		"not finite": {{Label: "a", Value: math.Inf(1)}},
		"NaN":        {{Label: "a", Value: math.NaN()}},
	}

	for name, slices := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := pterm.DefaultPieChart.WithSlices(slices).Srender()
			require.Error(t, err)
			assert.Error(t, pterm.DefaultPieChart.WithSlices(slices).Render())
		})
	}
}

func TestPieChartPrinterRawOutput(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	out, err := pterm.DefaultPieChart.WithSlices(pterm.PieSlices{
		{Label: "a", Value: 1.5},
		{Label: "b", Value: 4.5},
	}).Srender()
	require.NoError(t, err)

	assert.Equal(t, "a: 1.5 (25.0%)\nb: 4.5 (75.0%)\n", out)
}

func TestPieChartPrinterSrenderDoesNotMutateSlices(t *testing.T) {
	slices := pterm.PieSlices{{Label: "a", Value: 1}, {Label: "b", Value: 2}}
	before := append(pterm.PieSlices(nil), slices...)

	_, err := pterm.DefaultPieChart.WithSlices(slices).Srender()
	require.NoError(t, err)

	assert.Equal(t, before, slices)
}
