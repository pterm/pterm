package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal"
)

var DefaultTable = Table{
	HeaderStyle:    NewStyle(FgLightCyan),
	Separator:      " | ",
	SeparatorStyle: NewStyle(FgGray),
}

type TableData [][]string

type Table struct {
	Style          Style
	HasHeader      bool
	HeaderStyle    Style
	Separator      string
	SeparatorStyle Style
	Data           TableData
}

func (t Table) WithStyle(style Style) *Table {
	t.Style = style
	return &t
}

func (t Table) WithHasHeader(b ...bool) *Table {
	t.HasHeader = internal.WithBoolean(b)
	return &t
}

func (t Table) WithHeaderStyle(style Style) *Table {
	t.HeaderStyle = style
	return &t
}

func (t Table) WithSeparator(separator string) *Table {
	t.Separator = separator
	return &t
}

func (t Table) WithSeparatorStyle(style Style) *Table {
	t.SeparatorStyle = style
	return &t
}

func (t Table) WithData(data [][]string) *Table {
	t.Data = data
	return &t
}

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

func (t Table) Render() *Table {
	Println(t.Srender())
	return &t
}
