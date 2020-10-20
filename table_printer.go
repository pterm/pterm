package pterm

import (
	"encoding/csv"
	"strings"

	"github.com/pterm/pterm/internal"
)

// DefaultTable contains standards, which can be used to print a Table.
var DefaultTable = Table{
	Style:          &ThemeDefault.TableStyle,
	HeaderStyle:    &ThemeDefault.TableHeaderStyle,
	Separator:      " | ",
	SeparatorStyle: &ThemeDefault.TableSeparatorStyle,
}

// TableData is the type that contains the data of a Table.
type TableData [][]string

// Table is able to render tables.
type Table struct {
	Style          *Style
	HasHeader      bool
	HeaderStyle    *Style
	Separator      string
	SeparatorStyle *Style
	Data           TableData
}

// WithStyle returns a new Table with a specific Style.
func (p Table) WithStyle(style *Style) *Table {
	p.Style = style
	return &p
}

// WithHasHeader returns a new Table, where the first line is marked as a header.
func (p Table) WithHasHeader(b ...bool) *Table {
	p.HasHeader = internal.WithBoolean(b)
	return &p
}

// WithHeaderStyle returns a new Table with a specific HeaderStyle.
func (p Table) WithHeaderStyle(style *Style) *Table {
	p.HeaderStyle = style
	return &p
}

// WithSeparator returns a new Table with a specific separator.
func (p Table) WithSeparator(separator string) *Table {
	p.Separator = separator
	return &p
}

// WithSeparatorStyle returns a new Table with a specific SeparatorStyle.
func (p Table) WithSeparatorStyle(style *Style) *Table {
	p.SeparatorStyle = style
	return &p
}

// WithData returns a new Table with specific Data.
func (p Table) WithData(data [][]string) *Table {
	p.Data = data
	return &p
}

// WithCSVReader return a new Table with specified Data extracted from CSV.
func (p Table) WithCSVReader(reader *csv.Reader) *Table {
	if records, err := reader.ReadAll(); err == nil {
		p.Data = records
	}
	return &p
}

// Srender renders the Table as a string.
func (p Table) Srender() string {
	if p.Style == nil {
		p.Style = NewStyle()
	}
	if p.SeparatorStyle == nil {
		p.SeparatorStyle = NewStyle()
	}
	if p.HeaderStyle == nil {
		p.HeaderStyle = NewStyle()
	}

	var ret string
	maxColumnWidth := make(map[int]int)

	for _, row := range p.Data {
		for ci, column := range row {
			columnLength := len(RemoveColorFromString(column))
			if columnLength > maxColumnWidth[ci] {
				maxColumnWidth[ci] = columnLength
			}
		}
	}

	for ri, row := range p.Data {
		for ci, column := range row {
			columnLength := len(RemoveColorFromString(column))
			columnString := column + strings.Repeat(" ", maxColumnWidth[ci]-columnLength)

			if ci != len(row) && ci != 0 {
				ret += p.Style.Sprint(p.SeparatorStyle.Sprint(p.Separator))
			}

			if p.HasHeader && ri == 0 {
				ret += p.Style.Sprint(p.HeaderStyle.Sprint(columnString))
			} else {
				ret += p.Style.Sprint(columnString)
			}
		}

		ret += "\n"
	}

	ret = strings.TrimSuffix(ret, "\n")

	return ret
}

// Render prints the Table to the terminal.
func (p Table) Render() {
	Println(p.Srender())
}
