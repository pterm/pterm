package pterm

import (
	"strings"
	"unicode/utf8"

	"github.com/pterm/pterm/internal"
)

// DefaultGroupTable contains standards, which can be used to print a TablePrinter.
var DefaultGroupTable = TableGroup{
	Style:           &ThemeDefault.TableStyle,
	HasHeader:       true,
	HeaderStyle:     &ThemeDefault.TableHeaderStyle,
	Separator:       " | ",
	SeparatorStyle:  &ThemeDefault.TableSeparatorStyle,
	HSeparator:      "-",
	HSeparatorStyle: &ThemeDefault.TableSeparatorStyle,
}

type TableGroup struct {
	Style           *Style
	HasHeader       bool
	HeaderStyle     *Style
	Separator       string
	SeparatorStyle  *Style
	Boxed           bool
	HSeparator      string
	HSeparatorStyle *Style
	Share           bool
	ColumnLength    map[int]int
	Columns         int
	Tables          []*TablePrinter
}

// WithShare returns a new TableGroup that all tables within share same values.
func (t *TableGroup) WithShare(b ...bool) *TableGroup {
	t.Share = internal.WithBoolean(b)
	return t
}

// WithStyle returns a new TableGroup with a specific Style.
func (t *TableGroup) WithStyle(style *Style) *TableGroup {
	t.Style = style
	return t
}

// WithHasHeader returns a new TablePrinter, where the first line is marked as a header.
func (t *TableGroup) WithHasHeader(b ...bool) *TableGroup {
	t.HasHeader = internal.WithBoolean(b)
	return t
}

// WithHeaderStyle returns a new TableGroup with a specific HeaderStyle.
func (t *TableGroup) WithHeaderStyle(style *Style) *TableGroup {
	t.HeaderStyle = style
	return t
}

// WithSeparator returns a new TableGroup with a specific separator.
func (t *TableGroup) WithSeparator(separator string) *TableGroup {
	t.Separator = separator
	return t
}

// WithSeparatorStyle returns a new TableGroup with a specific SeparatorStyle.
func (t *TableGroup) WithSeparatorStyle(style *Style) *TableGroup {
	t.SeparatorStyle = style
	return t
}

// WithHorizontalSeparator returns a new TableGroup with a specific horizontal separator.
func (t *TableGroup) WithHorizontalSeparator(separator string) *TableGroup {
	if utf8.RuneCountInString(separator) != 1 {
		t.HSeparator = DefaultGroupTable.HSeparator

		return t
	}

	t.HSeparator = separator

	return t
}

// WithHorizontalSeparatorStyle returns a new TableGroup with a specific HorizontalSeparatorStyle.
func (t *TableGroup) WithHorizontalSeparatorStyle(style *Style) *TableGroup {
	t.HSeparatorStyle = style
	return t
}

// WithBoxed returns a new TableGroup with a box around the table.
func (t *TableGroup) WithBoxed(b ...bool) *TableGroup {
	t.Boxed = internal.WithBoolean(b)
	return t
}

// Append appends the table to the TableGroup.
func (t *TableGroup) Append(items ...interface{}) Grouper {
	if t.ColumnLength == nil {
		t.ColumnLength = make(map[int]int)
	}

	for _, item := range items {
		if table, ok := item.(*TablePrinter); ok {
			if len(table.Data[0]) > t.Columns {
				t.Columns = len(table.Data[0])
			}

			for _, row := range table.Data {
				for ci, column := range row {
					columnLength := utf8.RuneCountInString(RemoveColorFromString(column))
					if columnLength > t.ColumnLength[ci] {
						t.ColumnLength[ci] = columnLength
					}
				}
			}

			t.Tables = append(t.Tables, table)
		}
	}

	g := Grouper(t)
	return g
}

// Srender renders the TableGroup as a string.
func (t *TableGroup) Srender() (string, error) {
	var ret string

	for ti, table := range t.Tables {
		if t.Share {
			table.Style = t.Style
			table.HasHeader = t.HasHeader
			table.HeaderStyle = t.HeaderStyle
			table.Separator = t.Separator
			table.SeparatorStyle = t.SeparatorStyle
		}

		if table.Style == nil {
			table.Style = NewStyle()
		}
		if table.SeparatorStyle == nil {
			table.SeparatorStyle = NewStyle()
		}
		if table.HeaderStyle == nil {
			table.HeaderStyle = NewStyle()
		}

		for ri, row := range table.Data {
			for i := 0; i < t.Columns; i++ {
				columnString := ""
				if i < len(row) {
					columnLength := utf8.RuneCountInString(RemoveColorFromString(row[i]))
					columnString = row[i] + strings.Repeat(" ", t.ColumnLength[i]-columnLength)
				} else {
					columnString = strings.Repeat(" ", t.ColumnLength[i])
				}

				if i != 0 {
					ret += table.Style.Sprint(table.SeparatorStyle.Sprint(table.Separator))
				}

				if table.HasHeader && ri == 0 {
					ret += table.Style.Sprint(table.HeaderStyle.Sprint(columnString))
				} else {
					ret += table.Style.Sprint(columnString)
				}
			}

			ret += "\n"
		}

		if t.HSeparator != "" && ti != len(t.Tables)-1 {
			ret += t.horizontalSeparator()

			ret += "\n"
		}
	}

	ret = strings.TrimSuffix(ret, "\n")

	if t.Boxed {
		ret = DefaultBox.Sprint(ret)
	}

	return ret, nil
}

// Render prints the TableGroup to the terminal.
func (t *TableGroup) Render() error {
	s, _ := t.Srender()
	Println(s)

	return nil
}

func (t *TableGroup) horizontalSeparator() string {
	width := 0

	for i, l := range t.ColumnLength {
		width += l

		if i != len(t.ColumnLength) && i != 0 {
			width += utf8.RuneCountInString(t.Separator)
		}
	}

	return t.Style.Sprint(t.HSeparatorStyle.Sprint(strings.Repeat(t.HSeparator, width)))
}
