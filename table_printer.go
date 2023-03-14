package pterm

import (
	"encoding/csv"
	"github.com/pterm/pterm/internal"
	"io"
	"strings"
)

// DefaultTable contains standards, which can be used to print a TablePrinter.
var DefaultTable = TablePrinter{
	Style:                   &ThemeDefault.TableStyle,
	HeaderStyle:             &ThemeDefault.TableHeaderStyle,
	HeaderRowSeparator:      "",
	HeaderRowSeparatorStyle: &ThemeDefault.TableSeparatorStyle,
	Separator:               " | ",
	SeparatorStyle:          &ThemeDefault.TableSeparatorStyle,
	RowSeparator:            "",
	RowSeparatorStyle:       &ThemeDefault.TableSeparatorStyle,
	LeftAlignment:           true,
	RightAlignment:          false,
}

// TableData is the type that contains the data of a TablePrinter.
type TableData [][]string

// TablePrinter is able to render tables.
type TablePrinter struct {
	Style                   *Style
	HasHeader               bool
	HeaderStyle             *Style
	HeaderRowSeparator      string
	HeaderRowSeparatorStyle *Style
	Separator               string
	SeparatorStyle          *Style
	RowSeparator            string
	RowSeparatorStyle       *Style
	Data                    TableData
	Boxed                   bool
	LeftAlignment           bool
	RightAlignment          bool
	Writer                  io.Writer
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

// WithHeaderRowSeparator returns a new TablePrinter with a specific header HeaderRowSeparator.
func (p TablePrinter) WithHeaderRowSeparator(separator string) *TablePrinter {
	p.HeaderRowSeparator = separator
	return &p
}

// WithHeaderRowSeparatorStyle returns a new TablePrinter with a specific header HeaderRowSeparatorStyle.
func (p TablePrinter) WithHeaderRowSeparatorStyle(style *Style) *TablePrinter {
	p.HeaderRowSeparatorStyle = style
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

// WithRowSeparator returns a new TablePrinter with a specific RowSeparator.
func (p TablePrinter) WithRowSeparator(separator string) *TablePrinter {
	p.RowSeparator = separator
	return &p
}

// WithRowSeparatorStyle returns a new TablePrinter with a specific RowSeparatorStyle.
func (p TablePrinter) WithRowSeparatorStyle(style *Style) *TablePrinter {
	p.RowSeparatorStyle = style
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

// WithWriter sets the Writer.
func (p TablePrinter) WithWriter(writer io.Writer) *TablePrinter {
	p.Writer = writer
	return &p
}

type table struct {
	rows            []row
	maxColumnWidths []int
}

type row struct {
	height       int
	cells        []cell
	columnWidths []int
}

type cell struct {
	width  int
	height int
	lines  []string
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
	if p.HeaderRowSeparatorStyle == nil {
		p.HeaderRowSeparatorStyle = NewStyle()
	}
	if p.RowSeparatorStyle == nil {
		p.RowSeparatorStyle = NewStyle()
	}

	var t table

	// convert data to table and calculate values
	for _, rRaw := range p.Data {
		var r row
		for _, cRaw := range rRaw {
			var c cell
			c.lines = strings.Split(cRaw, "\n")
			c.height = len(c.lines)
			for _, l := range c.lines {
				if internal.GetStringMaxWidth(l) > c.width {
					c.width = internal.GetStringMaxWidth(l)
				}
			}
			r.cells = append(r.cells, c)
			if c.height > r.height {
				r.height = c.height
			}
		}

		// set max column widths of table
		for i, c := range r.cells {
			if len(t.maxColumnWidths) <= i {
				t.maxColumnWidths = append(t.maxColumnWidths, c.width)
			} else if c.width > t.maxColumnWidths[i] {
				t.maxColumnWidths[i] = c.width
			}
		}

		t.rows = append(t.rows, r)
	}

	var maxRowWidth int
	for _, r := range t.rows {
		rowWidth := internal.GetStringMaxWidth(p.renderRow(t, r))
		if rowWidth > maxRowWidth {
			maxRowWidth = rowWidth
		}
	}

	// render table
	var s string

	for i, r := range t.rows {
		if i == 0 && p.HasHeader {
			s += p.HeaderStyle.Sprint(p.renderRow(t, r))

			if p.HeaderRowSeparator != "" {
				s += strings.Repeat(p.HeaderRowSeparatorStyle.Sprint(p.HeaderRowSeparator), maxRowWidth) + "\n"
			}
			continue
		}

		s += p.renderRow(t, r)

		if p.RowSeparator != "" {
			s += strings.Repeat(p.RowSeparatorStyle.Sprint(p.RowSeparator), maxRowWidth) + "\n"
		}
	}

	if p.Boxed {
		s = DefaultBox.Sprint(strings.TrimSuffix(s, "\n"))
	}

	return s, nil
}

// renderRow renders a row.
// It merges the cells of a row into one string.
// Each line of each cell is merged with the same line of the other cells.
func (p TablePrinter) renderRow(t table, r row) string {
	var s string

	// merge lines of cells and add separator
	// use the t.maxColumnWidths to add padding to the corresponding cell
	// a newline in a cell should be in the same column as the original cell
	for i := 0; i < r.height; i++ {
		for j, c := range r.cells {
			var currentLine string
			if i < len(c.lines) {
				currentLine = c.lines[i]
			}
			paddingForLine := t.maxColumnWidths[j] - internal.GetStringMaxWidth(currentLine)

			if p.RightAlignment {
				s += strings.Repeat(" ", paddingForLine)
			}

			if i < len(c.lines) {
				s += strings.TrimSpace(c.lines[i])
			}

			if j < len(r.cells)-1 {
				if p.LeftAlignment {
					s += strings.Repeat(" ", paddingForLine)
				}
				s += p.SeparatorStyle.Sprint(p.Separator)
			}
		}
		s += "\n"
	}

	return s
}

// Render prints the TablePrinter to the terminal.
func (p TablePrinter) Render() error {
	s, _ := p.Srender()
	Fprintln(p.Writer, s)

	return nil
}
