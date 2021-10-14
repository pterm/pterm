package pterm

import (
	"encoding/csv"
	"strings"
	"unicode/utf8"

	"github.com/pterm/pterm/internal"
)

// DefaultTable contains standards, which can be used to print a TablePrinter.
var DefaultTable = TablePrinter{
	Style:          &ThemeDefault.TableStyle,
	HeaderStyle:    &ThemeDefault.TableHeaderStyle,
	Separator:      " | ",
	SeparatorStyle: &ThemeDefault.TableSeparatorStyle,
	LeftAlignment:  true,
	RightAlignment: false,
}

// TableData is the type that contains the data of a TablePrinter.
type TableData [][]string

// TablePrinter is able to render tables.
type TablePrinter struct {
	Style          *Style
	HasHeader      bool
	HeaderStyle    *Style
	Separator      string
	SeparatorStyle *Style
	Data           TableData
	Boxed          bool
	LeftAlignment  bool
	RightAlignment bool
}

// WithStyle returns a new TablePrinter with a specific Style.
func (p TablePrinter) WithStyle(style *Style) *TablePrinter {
	p.Style = style
	return &p
}

// WithHasHeader returns a new TablePrinter, where the first line is marked as a header.
func (p TablePrinter) WithHasHeader(b ...bool) *TablePrinter {
	p.HasHeader = internal.WithBoolean(b)
	return &p
}

// WithHeaderStyle returns a new TablePrinter with a specific HeaderStyle.
func (p TablePrinter) WithHeaderStyle(style *Style) *TablePrinter {
	p.HeaderStyle = style
	return &p
}

// WithSeparator returns a new TablePrinter with a specific separator.
func (p TablePrinter) WithSeparator(separator string) *TablePrinter {
	p.Separator = separator
	return &p
}

// WithSeparatorStyle returns a new TablePrinter with a specific SeparatorStyle.
func (p TablePrinter) WithSeparatorStyle(style *Style) *TablePrinter {
	p.SeparatorStyle = style
	return &p
}

// WithData returns a new TablePrinter with specific Data.
func (p TablePrinter) WithData(data [][]string) *TablePrinter {
	p.Data = data
	return &p
}

// WithCSVReader return a new TablePrinter with specified Data extracted from CSV.
func (p TablePrinter) WithCSVReader(reader *csv.Reader) *TablePrinter {
	if records, err := reader.ReadAll(); err == nil {
		p.Data = records
	}
	return &p
}

// WithBoxed returns a new TablePrinter with a box around the table.
func (p TablePrinter) WithBoxed(b ...bool) *TablePrinter {
	p.Boxed = internal.WithBoolean(b)
	return &p
}

// WithLeftAlignment returns a new TablePrinter with left alignment.
func (p TablePrinter) WithLeftAlignment(b ...bool) *TablePrinter {
	b2 := internal.WithBoolean(b)
	p.LeftAlignment = b2
	p.RightAlignment = false
	return &p
}

// WithRightAlignment returns a new TablePrinter with right alignment.
func (p TablePrinter) WithRightAlignment(b ...bool) *TablePrinter {
	b2 := internal.WithBoolean(b)
	p.LeftAlignment = false
	p.RightAlignment = b2
	return &p
}

// Srender renders the TablePrinter as a string.
func (p TablePrinter) Srender() (string, error) {
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
			columnLength := utf8.RuneCountInString(RemoveColorFromString(column))
			if columnLength > maxColumnWidth[ci] {
				maxColumnWidth[ci] = columnLength
			}
		}
	}

	for ri, row := range p.Data {
		for ci, column := range row {
			columnString := p.createColumnString(column, maxColumnWidth[ci])

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

	if p.Boxed {
		ret = DefaultBox.Sprint(ret)
	}

	return ret, nil
}

func (p TablePrinter) createColumnString(data string, maxColumnWidth int) string {
	columnLength := utf8.RuneCountInString(RemoveColorFromString(data))
	if p.RightAlignment {
		return strings.Repeat(" ", maxColumnWidth-columnLength) + data
	}
	return data + strings.Repeat(" ", maxColumnWidth-columnLength)
}

// Render prints the TablePrinter to the terminal.
func (p TablePrinter) Render() error {
	s, _ := p.Srender()
	Println(s)

	return nil
}
