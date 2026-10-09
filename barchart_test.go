package pterm_test

// Behavioral tests for BarChartPrinter: the bar length math (values mapped to
// Height/Width), positive/negative chart parts, value display, label
// alignment and render purity. The builder/contract plumbing is covered in
// contract_test.go, representative outputs in snapshot_test.go.

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func TestBarChartPrinter_VerticalBarsProportionalToValues(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHeight(4).WithBars(pterm.Bars{
		{Label: "a", Value: 1},
		{Label: "b", Value: 2},
		{Label: "c", Value: 4},
	})

	// The maximum value (4) fills the full height; the other bars are
	// scaled linearly (1 and 2 cells) and share the baseline.
	expected := "" +
		"      ██ \n" +
		"      ██ \n" +
		"   ██ ██ \n" +
		"██ ██ ██ \n" +
		"a  b  c  \n" +
		"         \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_VerticalShowValuePrintsValueAboveBar(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHeight(2).WithShowValue().WithBars(pterm.Bars{
		{Label: "a", Value: 1},
		{Label: "b", Value: 2},
	})

	expected := "" +
		"1  2  \n" +
		"   ██ \n" +
		"██ ██ \n" +
		"a  b  \n" +
		"      \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_VerticalMixedValuesSplitChartAtZeroLine(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHeight(4).WithBars(pterm.Bars{
		{Label: "p", Value: 2},
		{Label: "n", Value: -2},
	})

	// With mixed signs the height is split: the positive bar occupies the
	// top half, the negative bar the bottom half below the zero line.
	expected := "" +
		"██    \n" +
		"██    \n" +
		"      \n" +
		"   ██ \n" +
		"   ██ \n" +
		"p  n  \n" +
		"      \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_VerticalZeroValuesRenderNoBars(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHeight(3).WithBars(pterm.Bars{
		{Label: "a", Value: 0},
		{Label: "b", Value: 0},
	})

	expected := "" +
		"      \n" +
		"      \n" +
		"      \n" +
		"a  b  \n" +
		"      \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_VerticalEqualValuesRenderFullHeight(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHeight(2).WithBars(pterm.Bars{
		{Label: "a", Value: 5},
		{Label: "b", Value: 5},
	})

	expected := "" +
		"██ ██ \n" +
		"██ ██ \n" +
		"a  b  \n" +
		"      \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_HorizontalBarsProportionalToValues(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHorizontal().WithWidth(6).WithBars(pterm.Bars{
		{Label: "a", Value: 3},
		{Label: "bb", Value: 6},
	})

	// The maximum value (6) fills the full width, 3 exactly half; labels are
	// left of the bars and padded to the widest label.
	expected := "" +
		"          \n" +
		"a  ███    \n" +
		"bb ██████ \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_HorizontalMixedValuesRenderOnBothSidesOfZero(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHorizontal().WithWidth(8).WithBars(pterm.Bars{
		{Label: "n", Value: -3},
		{Label: "p", Value: 6},
	})

	// Negative bars grow leftwards in the left (negative) part of the chart,
	// positive bars start right of it.
	expected := "" +
		"         \n" +
		"n ██     \n" +
		"p   ████ \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_HorizontalShowValueAppendsAlignedValues(t *testing.T) {
	printer := pterm.DefaultBarChart.WithHorizontal().WithWidth(4).WithShowValue().WithBars(pterm.Bars{
		{Label: "a", Value: 2},
		{Label: "b", Value: 4},
	})

	expected := "" +
		"          \n" +
		"a ██    2 \n" +
		"b ████  4 \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBarChartPrinter_CustomBarCharacters(t *testing.T) {
	t.Run("horizontal", func(t *testing.T) {
		printer := pterm.DefaultBarChart.WithHorizontal().WithWidth(4).
			WithHorizontalBarCharacter("#").
			WithBars(pterm.Bars{{Label: "a", Value: 2}})

		expected := "" +
			"       \n" +
			"a #### \n"

		assert.Equal(t, expected, srenderPlain(t, printer))
	})

	t.Run("vertical", func(t *testing.T) {
		printer := pterm.DefaultBarChart.WithHeight(1).
			WithVerticalBarCharacter("*").
			WithBars(pterm.Bars{{Label: "a", Value: 1}})

		expected := "" +
			"* \n" +
			"a \n" +
			"  \n"

		assert.Equal(t, expected, srenderPlain(t, printer))
	})
}

func TestBarChartPrinter_RawOutputListsLabelsAndValues(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	out, err := pterm.DefaultBarChart.WithBars(pterm.Bars{
		{Label: "a", Value: 1},
		{Label: "b", Value: -2},
	}).Srender()
	require.NoError(t, err)

	assert.Equal(t, "a: 1\nb: -2\n", out)
}

// Regression test: Srender used to write the styled labels (and default
// styles) back into the caller's Bars slice, so every subsequent render
// re-styled the already styled labels.
func TestBarChartPrinter_SrenderIsPure(t *testing.T) {
	for _, horizontal := range []bool{false, true} {
		name := "vertical"
		if horizontal {
			name = "horizontal"
		}

		t.Run(name, func(t *testing.T) {
			bars := pterm.Bars{
				{Label: "a", Value: 1},
				{Label: "b", Value: 2},
			}
			printer := pterm.DefaultBarChart.WithHorizontal(horizontal).WithHeight(3).WithWidth(3).WithBars(bars)

			first, err := printer.Srender()
			require.NoError(t, err)

			second, err := printer.Srender()
			require.NoError(t, err)

			assert.Equal(t, first, second, "rendering twice must yield identical output")
			assert.Equal(t, pterm.Bars{
				{Label: "a", Value: 1},
				{Label: "b", Value: 2},
			}, bars, "rendering must not modify the input bars (labels or styles)")
		})
	}
}
