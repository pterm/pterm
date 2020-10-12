package pterm

import (
	"encoding/csv"
	"strings"

	"github.com/pterm/pterm/internal"
)

// DefaultTable contains standards, which can be used to print a Table.
var DefaultTable = Table{
	HeaderStyle:    NewStyle(FgLightCyan),
	Separator:      " | ",
	SeparatorStyle: NewStyle(FgGray),
}

// TableData is the type that contains the data of a Table.
type TableData [][]string

// Table is able to render tables.
type Table struct {
	Style          Style
	HasHeader      bool
	HeaderStyle    Style
	Separator      string
	SeparatorStyle Style
	Data           TableData
}

// WithStyle returns a new Table with a specific Style.
func (t Table) WithStyle(style Style) *Table {
	t.Style = style
	return &t
}

// WithHasHeader returns a new Table, where the first line is marked as a header.
func (t Table) WithHasHeader(b ...bool) *Table {
	t.HasHeader = internal.WithBoolean(b)
	return &t
}

// WithHeaderStyle returns a new Table with a specific HeaderStyle.
func (t Table) WithHeaderStyle(style Style) *Table {
	t.HeaderStyle = style
	return &t
}

// WithSeparator returns a new Table with a specific separator.
func (t Table) WithSeparator(separator string) *Table {
	t.Separator = separator
	return &t
}

// WithSeparatorStyle returns a new Table with a specific SeparatorStyle.
func (t Table) WithSeparatorStyle(style Style) *Table {
	t.SeparatorStyle = style
	return &t
}

// WithData returns a new Table with specific Data.
func (t Table) WithData(data [][]string) *Table {
	t.Data = data
	return &t
}

// WithCSVReader return a new Table with specified Data extracted from CSV.
func (t Table) WithCSVReader(reader *csv.Reader) *Table {
	if records, err := reader.ReadAll(); err == nil {
		t.Data = records
	}
	return &t
}

// Srender renders the Table as a string.
func (t Table) Srender() string {
	var ret string
	maxColumnWidth := make(map[int]int)

	for _, row := range t.Data {
		for ci, column := range row {
			columnLength := len(RemoveColorFromString(column))
			if columnLength > maxColumnWidth[ci] {
				maxColumnWidth[ci] = columnLength
			}
		}
	}

	for ri, row := range t.Data {
		for ci, column := range row {
			columnLength := len(RemoveColorFromString(column))
			columnString := column + strings.Repeat(" ", maxColumnWidth[ci]-columnLength)

			if ci != len(row) && ci != 0 {
				ret += t.Style.Sprint(t.SeparatorStyle.Sprint(t.Separator))
			}

			if t.HasHeader && ri == 0 {
				ret += t.Style.Sprint(t.HeaderStyle.Sprint(columnString))
			} else {
				ret += t.Style.Sprint(columnString)
			}
		}

		ret += "\n"
	}

	ret = strings.TrimSuffix(ret, "\n")

	return ret
}

// Render prints the Table to the terminal.
func (t Table) Render() {
	Println(t.Srender())
}
