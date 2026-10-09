package pterm_test

// Behavioral tests for PanelPrinter: side-by-side line merging, padding math,
// column width normalization, bottom padding and boxed panels. The
// builder/contract plumbing is covered in contract_test.go, one representative
// output in snapshot_test.go.

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func TestPanelPrinter_SideBySideLayout(t *testing.T) {
	printer := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: "a\nbc"}, {Data: "xyz"}},
	})

	// Both panels share output lines: every panel line is padded to its
	// panel's width plus the default padding of one space; missing lines of
	// shorter panels become spaces.
	expected := "" +
		"a  xyz \n" +
		"bc     \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestPanelPrinter_PaddingBetweenColumns(t *testing.T) {
	printer := pterm.DefaultPanel.WithPadding(3).WithPanels(pterm.Panels{
		{{Data: "a"}, {Data: "b"}},
	})

	expected := "a   b   \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestPanelPrinter_SameColumnWidthPadsAllRowsToWidestPanel(t *testing.T) {
	printer := pterm.DefaultPanel.WithSameColumnWidth().WithPanels(pterm.Panels{
		{{Data: "a"}},
		{{Data: "cccc"}},
	})

	// Without SameColumnWidth the first row would be 2 cells wide; with it,
	// every row of the column is padded to the widest panel (4) plus padding.
	expected := "" +
		"a    \n" +
		"cccc \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestPanelPrinter_BottomPaddingAddsBlankLinesBetweenRows(t *testing.T) {
	printer := pterm.DefaultPanel.WithBottomPadding(1).WithPanels(pterm.Panels{
		{{Data: "a"}},
		{{Data: "b"}},
	})

	// One blank line after every row except the last.
	expected := "" +
		"a \n" +
		"  \n" +
		"b \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestPanelPrinter_BoxedPanelsAlign(t *testing.T) {
	printer := pterm.DefaultPanel.WithBoxPrinter(pterm.DefaultBox).WithPanels(pterm.Panels{
		{{Data: "a"}, {Data: "b\nc"}},
	})

	// Each panel is boxed individually; the shorter box is padded with
	// spaces so the columns stay aligned.
	expected := "" +
		"╭───╮ ╭───╮ \n" +
		"│ a │ │ b │ \n" +
		"╰───╯ │ c │ \n" +
		"      ╰───╯ \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestPanelPrinter_RawOutputListsPanelsSequentially(t *testing.T) {
	restoreGlobalStyling(t)
	pterm.DisableStyling()

	out, err := pterm.DefaultPanel.WithPanels(pterm.Panels{
		{{Data: "left"}, {Data: "right"}},
	}).Srender()
	require.NoError(t, err)

	// Documents current behavior: raw mode does not lay panels out
	// side-by-side but prints them one below the other.
	assert.Equal(t, "left\n\nright\n\n\n", out)
}

func TestPanelPrinter_SrenderIsPure(t *testing.T) {
	// Trailing newlines are trimmed, boxes and bottom padding are added
	// during rendering — none of that may leak into the caller's panels.
	panels := pterm.Panels{
		{{Data: "a\n"}, {Data: "b"}},
		{{Data: "c"}},
	}
	printer := pterm.DefaultPanel.WithBoxPrinter(pterm.DefaultBox).WithBottomPadding(2).WithPanels(panels)

	first, err := printer.Srender()
	require.NoError(t, err)

	second, err := printer.Srender()
	require.NoError(t, err)

	assert.Equal(t, first, second, "rendering twice must yield identical output")
	assert.Equal(t, pterm.Panels{
		{{Data: "a\n"}, {Data: "b"}},
		{{Data: "c"}},
	}, panels, "rendering must not modify the input panels")
}
