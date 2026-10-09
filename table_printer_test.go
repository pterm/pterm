package pterm_test

// Behavioral tests for TablePrinter: column width math, alignment, separators,
// header styling, boxing and CSV input. The builder/contract plumbing is
// covered generically in contract_test.go, one representative output is locked
// in snapshot_test.go.

import (
	"encoding/csv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

// srenderPlain renders a RenderPrinter and strips all ANSI escape codes, so
// tests can assert on the exact visible layout. Shared by the render printer
// test files (table, tree, bullet list, panel, bar chart, heatmap, big text).
func srenderPlain(t *testing.T, p pterm.RenderPrinter) string {
	t.Helper()

	out, err := p.Srender()
	require.NoError(t, err)

	return stripANSI(out)
}

// styledPrefix returns the escape sequence a style emits directly before its
// content, so tests can detect which style a rendered region starts with.
func styledPrefix(t *testing.T, style *pterm.Style) string {
	t.Helper()

	parts := strings.SplitN(style.Sprint("X"), "X", 2)
	require.NotEmpty(t, parts[0], "style must emit an escape sequence")

	return parts[0]
}

func TestTablePrinter_ColumnsPaddedToWidestCell(t *testing.T) {
	printer := pterm.DefaultTable.WithData(pterm.TableData{
		{"Name", "Age"},
		{"Alice", "1"},
		{"Bob", "22"},
	})

	// Column 0 is padded to "Alice" (5), column 1 to "Age" (3); the last
	// column is padded too, so every separator lines up across all rows.
	expected := "" +
		"Name  │ Age\n" +
		"Alice │ 1  \n" +
		"Bob   │ 22 \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_WideUnicodeCellsAlign(t *testing.T) {
	printer := pterm.DefaultTable.WithData(pterm.TableData{
		{"汉字", "x"},
		{"abc", "y"},
	})

	// "汉字" occupies 4 terminal cells, so "abc" gets one space of padding.
	expected := "" +
		"汉字 │ x\n" +
		"abc  │ y\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_ANSICodesInCellsDoNotAffectPadding(t *testing.T) {
	printer := pterm.DefaultTable.WithData(pterm.TableData{
		{pterm.FgRed.Sprint("a"), "b"},
		{"cc", "d"},
	})

	// The styled "a" is much longer in bytes than "cc", but only its visible
	// width (1) may count for the column width.
	expected := "" +
		"a  │ b\n" +
		"cc │ d\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_RightAlignment(t *testing.T) {
	printer := pterm.DefaultTable.WithRightAlignment().WithData(pterm.TableData{
		{"Name", "Age"},
		{"Alice", "1"},
		{"Bob", "22"},
	})

	expected := "" +
		" Name │ Age\n" +
		"Alice │   1\n" +
		"  Bob │  22\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_MultilineCellsStayInTheirColumn(t *testing.T) {
	printer := pterm.DefaultTable.WithData(pterm.TableData{
		{"a\nb", "x"},
		{"c", "y"},
	})

	// The second line of the multiline cell stays in column 0; column 1 is
	// padded with spaces on that line.
	expected := "" +
		"a │ x\n" +
		"b │  \n" +
		"c │ y\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_HeaderRowSeparatorSpansTableWidth(t *testing.T) {
	printer := pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("=").WithData(pterm.TableData{
		{"Name", "Age"},
		{"Bob", "1"},
	})

	expected := "" +
		"Name │ Age\n" +
		"==========\n" +
		"Bob  │ 1  \n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_RowSeparatorBetweenRowsOnly(t *testing.T) {
	printer := pterm.DefaultTable.WithRowSeparator("-").WithData(pterm.TableData{
		{"a"},
		{"b"},
		{"c"},
	})

	// A separator line (as wide as the widest row) between rows, but not
	// after the last one.
	expected := "a\n-\nb\n-\nc\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_HasHeaderStylesOnlyFirstRow(t *testing.T) {
	data := pterm.TableData{
		{"Name", "Age"},
		{"Alice", "1"},
	}

	styled, err := pterm.DefaultTable.WithHasHeader().WithData(data).Srender()
	require.NoError(t, err)

	prefix := styledPrefix(t, &pterm.ThemeDefault.TableHeaderStyle)
	assert.Contains(t, styled, prefix+"Name", "header row must be wrapped in the header style")
	assert.NotContains(t, styled, prefix+"Alice", "data rows must not get the header style")

	// Apart from the header separator line, the header style must not change
	// the visible layout.
	expected := "" +
		"Name  │ Age\n" +
		"───────────\n" +
		"Alice │ 1  \n"
	assert.Equal(t, expected, stripANSI(styled))
}

func TestTablePrinter_AlternateRowStyleOnEverySecondRow(t *testing.T) {
	alt := pterm.NewStyle(pterm.FgMagenta)
	printer := pterm.DefaultTable.WithHasHeader().WithAlternateRowStyle(alt).WithData(pterm.TableData{
		{"Name", "Age"},
		{"Bob", "1"},
		{"Callie", "2"},
		{"Dean", "3"},
	})

	styled, err := printer.Srender()
	require.NoError(t, err)

	prefix := styledPrefix(t, alt)
	assert.Contains(t, styled, prefix+"Bob", "row 1 must use the alternate style")
	assert.Contains(t, styled, prefix+"Dean", "row 3 must use the alternate style")
	assert.NotContains(t, styled, prefix+"Callie", "row 2 must not use the alternate style")
}

func TestTablePrinter_BoxedWrapsTableInAlignedBox(t *testing.T) {
	printer := pterm.DefaultTable.WithBoxed().WithData(pterm.TableData{
		{"a", "b"},
		{"cc", "dd"},
	})

	expected := "" +
		"╭─────────╮\n" +
		"│ a  │ b  │\n" +
		"│ cc │ dd │\n" +
		"╰─────────╯"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestTablePrinter_EmptyDataRendersEmptyString(t *testing.T) {
	// Documents current behavior: a table without data renders as an empty
	// string and does not return an error.
	out, err := pterm.DefaultTable.Srender()
	require.NoError(t, err)
	assert.Empty(t, out)
}

func TestTablePrinter_WithCSVReaderParsesRecordsIntoData(t *testing.T) {
	reader := csv.NewReader(strings.NewReader("Name,Age\nAlice,30\n"))
	printer := pterm.DefaultTable.WithCSVReader(reader)

	require.Equal(t, pterm.TableData{{"Name", "Age"}, {"Alice", "30"}}, printer.Data)

	// The CSV-fed table must render exactly like the same data set directly.
	fromData := srenderPlain(t, pterm.DefaultTable.WithData(pterm.TableData{{"Name", "Age"}, {"Alice", "30"}}))
	assert.Equal(t, fromData, srenderPlain(t, printer))
}

func TestTablePrinter_WithCSVReaderKeepsDataOnParseError(t *testing.T) {
	reader := csv.NewReader(strings.NewReader("a,b\nc\n")) // inconsistent field count
	printer := pterm.DefaultTable.WithCSVReader(reader)

	assert.Nil(t, printer.Data, "malformed CSV must leave Data unchanged")
}

func TestTablePrinter_SrenderIsPure(t *testing.T) {
	data := pterm.TableData{
		{"Name", "Age"},
		{"Alice", "1"},
	}
	printer := pterm.DefaultTable.WithHasHeader().WithData(data)

	first, err := printer.Srender()
	require.NoError(t, err)

	second, err := printer.Srender()
	require.NoError(t, err)

	assert.Equal(t, first, second, "rendering twice must yield identical output")
	assert.Equal(t, pterm.TableData{{"Name", "Age"}, {"Alice", "1"}}, data, "rendering must not modify the input data")
}
