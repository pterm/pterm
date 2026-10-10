package pterm

import (
	"encoding/csv"
	"io"
	"regexp"
	"strings"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/text"
)

// TableData is the content of a table: rows of cells. With a header, the
// first row is the header.
type TableData [][]string

// TablePrinter renders rows of data in aligned columns. The default look has
// no frame, a muted uppercase header with a rule under it and right-aligned
// numbers:
//
//	CACHE  STATE     SIZE
//	─────────────────────
//	npm    stale  1.20 GB
//	nub    fresh   393 MB
//
// Cells that do not fit the terminal are cut with an ellipsis, widest
// columns first. They never wrap, and nothing is cut when the width is
// unknown.
type TablePrinter struct {
	// Data holds the rows. With HasHeader, the first row is the header.
	Data TableData
	// HasHeader styles the first row as the header.
	HasHeader bool
	// HeaderStyle styles the header cells.
	HeaderStyle Style
	// HeaderUppercase prints the header in uppercase.
	HeaderUppercase bool
	// HeaderRule draws a rule under the header.
	HeaderRule bool
	// CellStyle styles the body cells.
	CellStyle Style
	// BorderTone colors rules, separators and the frame.
	BorderTone Tone
	// Separator goes between two columns.
	Separator string
	// RowSeparators draws a rule between body rows.
	RowSeparators bool
	// Boxed draws a rounded frame around the table.
	Boxed bool
	// Title is inset into the top of the frame. Only used when Boxed.
	Title string
	// TitleStyle styles the title.
	TitleStyle Style
	// Padding is the space between the frame and the cells. Only used when
	// Boxed.
	Padding int
	// Alignment aligns every column that has no alignment of its own.
	// AlignAuto right-aligns columns that hold only numbers and sizes.
	Alignment Alignment
	// ColumnAlignments aligns single columns, indexed from the left.
	// AlignAuto falls back to Alignment.
	ColumnAlignments []Alignment
	// ColumnMinWidths is the minimum width of each column, indexed from the
	// left. Content still widens a column, and a terminal that is too
	// narrow still cuts it.
	ColumnMinWidths []int
	// ColumnStyles style the body cells of single columns. The zero Style
	// keeps CellStyle.
	ColumnStyles []Style
	// RowStyles style single body rows, indexed from the first row after
	// the header. The zero Style keeps the other styles.
	RowStyles []Style
	// AlternateRowStyle styles every second body row. The zero Style turns
	// it off.
	AlternateRowStyle Style
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Table is the default table. Its first row is the header.
var Table = TablePrinter{
	HasHeader:       true,
	HeaderStyle:     ToneMuted.Style(),
	HeaderUppercase: true,
	HeaderRule:      true,
	CellStyle:       ToneText.Style(),
	BorderTone:      ToneBorder,
	Separator:       "  ",
	TitleStyle:      TonePrimary.Bold(),
	Padding:         1,
}

// WithData returns a copy of p that renders data.
func (p TablePrinter) WithData(data [][]string) *TablePrinter {
	p.Data = data
	return &p
}

// WithCSVReader returns a copy of p that renders the records of r. A
// reader that fails leaves the data unchanged.
func (p TablePrinter) WithCSVReader(r *csv.Reader) *TablePrinter {
	if records, err := r.ReadAll(); err == nil {
		p.Data = records
	}

	return &p
}

// WithHasHeader returns a copy of p that styles its first row as the
// header, or not with false.
func (p TablePrinter) WithHasHeader(b ...bool) *TablePrinter {
	p.HasHeader = flag(b)
	return &p
}

// WithHeaderStyle returns a copy of p whose header cells use s.
func (p TablePrinter) WithHeaderStyle(s Style) *TablePrinter {
	p.HeaderStyle = s
	return &p
}

// WithHeaderUppercase returns a copy of p that prints the header in
// uppercase, or as written with false.
func (p TablePrinter) WithHeaderUppercase(b ...bool) *TablePrinter {
	p.HeaderUppercase = flag(b)
	return &p
}

// WithHeaderRule returns a copy of p that draws a rule under the header,
// or none with false.
func (p TablePrinter) WithHeaderRule(b ...bool) *TablePrinter {
	p.HeaderRule = flag(b)
	return &p
}

// WithCellStyle returns a copy of p whose body cells use s.
func (p TablePrinter) WithCellStyle(s Style) *TablePrinter {
	p.CellStyle = s
	return &p
}

// WithBorderTone returns a copy of p that draws rules, separators and the
// frame in tone t.
func (p TablePrinter) WithBorderTone(t Tone) *TablePrinter {
	p.BorderTone = t
	return &p
}

// WithSeparator returns a copy of p that puts sep between two columns.
func (p TablePrinter) WithSeparator(sep string) *TablePrinter {
	p.Separator = sep
	return &p
}

// WithRowSeparators returns a copy of p that draws a rule between body
// rows, or none with false.
func (p TablePrinter) WithRowSeparators(b ...bool) *TablePrinter {
	p.RowSeparators = flag(b)
	return &p
}

// WithBoxed returns a copy of p with a rounded frame around the table, or
// without one with false.
func (p TablePrinter) WithBoxed(b ...bool) *TablePrinter {
	p.Boxed = flag(b)
	return &p
}

// WithTitle returns a copy of p that insets t into the top of its frame.
// The title only shows when the table is boxed.
func (p TablePrinter) WithTitle(t string) *TablePrinter {
	p.Title = t
	return &p
}

// WithTitleStyle returns a copy of p whose title uses s.
func (p TablePrinter) WithTitleStyle(s Style) *TablePrinter {
	p.TitleStyle = s
	return &p
}

// WithPadding returns a copy of p with n spaces between the frame and the
// cells.
func (p TablePrinter) WithPadding(n int) *TablePrinter {
	p.Padding = max(n, 0)
	return &p
}

// WithAlignment returns a copy of p that aligns every column without an
// alignment of its own with a.
func (p TablePrinter) WithAlignment(a Alignment) *TablePrinter {
	p.Alignment = a
	return &p
}

// WithLeftAlignment returns a copy of p that left-aligns every column, or
// goes back to automatic alignment with false.
func (p TablePrinter) WithLeftAlignment(b ...bool) *TablePrinter {
	p.Alignment = AlignAuto
	if flag(b) {
		p.Alignment = AlignLeft
	}

	return &p
}

// WithRightAlignment returns a copy of p that right-aligns every column, or
// goes back to automatic alignment with false.
func (p TablePrinter) WithRightAlignment(b ...bool) *TablePrinter {
	p.Alignment = AlignAuto
	if flag(b) {
		p.Alignment = AlignRight
	}

	return &p
}

// WithColumnAlignment returns a copy of p that aligns column col (counted
// from 0) with a.
func (p TablePrinter) WithColumnAlignment(col int, a Alignment) *TablePrinter {
	p.ColumnAlignments = setAt(p.ColumnAlignments, col, a)
	return &p
}

// WithColumnAlignments returns a copy of p that aligns its columns with a,
// from the left.
func (p TablePrinter) WithColumnAlignments(a ...Alignment) *TablePrinter {
	p.ColumnAlignments = a
	return &p
}

// WithColumnMinWidths returns a copy of p with a minimum width per column,
// from the left. Passing the same widths to several tables keeps their
// columns aligned.
func (p TablePrinter) WithColumnMinWidths(widths ...int) *TablePrinter {
	p.ColumnMinWidths = widths
	return &p
}

// WithColumnStyle returns a copy of p whose body cells in column col
// (counted from 0) use s.
func (p TablePrinter) WithColumnStyle(col int, s Style) *TablePrinter {
	p.ColumnStyles = setAt(p.ColumnStyles, col, s)
	return &p
}

// WithRowStyle returns a copy of p whose body row row (counted from 0,
// after the header) uses s.
func (p TablePrinter) WithRowStyle(row int, s Style) *TablePrinter {
	p.RowStyles = setAt(p.RowStyles, row, s)
	return &p
}

// WithAlternateRowStyle returns a copy of p that styles every second body
// row with s, which makes wide tables easier to follow.
func (p TablePrinter) WithAlternateRowStyle(s Style) *TablePrinter {
	p.AlternateRowStyle = s
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p TablePrinter) WithWriter(w io.Writer) *TablePrinter {
	p.Writer = w
	return &p
}

func (p TablePrinter) out() io.Writer { return orDefault(p.Writer) }

// Render prints the table.
func (p TablePrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

// Srender returns the table. It returns ErrEmpty when there are no cells.
func (p TablePrinter) Srender() (string, error) {
	rows, cols := p.cells()
	if cols == 0 {
		return "", ErrEmpty
	}

	c := canvasFor(p.out())
	widths := p.widths(c, rows, cols)
	aligns := p.alignments(rows, cols)
	sep := p.Separator

	if strings.TrimSpace(sep) != "" {
		sep = c.paint(p.BorderTone, sep)
	}

	inner := (cols - 1) * text.Width(p.Separator)
	for _, w := range widths {
		inner += w
	}

	var (
		lines []string
		frame = dataFrame{c: c, tone: p.BorderTone, inner: inner + 2*p.Padding}
		pad   = strings.Repeat(" ", p.Padding)
		rule  = c.paint(p.BorderTone, strings.Repeat(c.g.Rule, inner))
	)

	if p.Boxed {
		rule = frame.rule()
	}

	body := 0

	for r, row := range rows {
		header := r == 0 && p.HasHeader
		st := p.rowStyle(body, header)

		if !header && body > 0 && p.RowSeparators {
			lines = append(lines, rule)
		}

		for _, line := range p.renderRow(c, row, widths, aligns, header, st, sep) {
			if p.Boxed {
				line = frame.line(pad + line + pad)
			}

			lines = append(lines, line)
		}

		if header {
			if p.HeaderRule {
				lines = append(lines, rule)
			}

			continue
		}

		body++
	}

	if p.Boxed {
		lines = append([]string{frame.top(p.Title, p.TitleStyle)}, lines...)
		lines = append(lines, frame.bottom())
	}

	return strings.Join(lines, "\n"), nil
}

// tableCell is a cell split into its lines.
type tableCell []string

// cells splits every cell into lines and returns the rows and the number
// of columns. Short rows get empty cells.
func (p TablePrinter) cells() ([][]tableCell, int) {
	cols := 0
	for _, row := range p.Data {
		cols = max(cols, len(row))
	}

	rows := make([][]tableCell, len(p.Data))

	for r, row := range p.Data {
		rows[r] = make([]tableCell, cols)

		for col := range cols {
			s := at(row, col)
			if r == 0 && p.HasHeader && p.HeaderUppercase {
				s = upperVisible(s)
			}

			rows[r][col] = strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
		}
	}

	return rows, cols
}

// widths returns the width of every column: the widest line, at least the
// minimum width, then shrunk widest first until the table fits the canvas.
func (p TablePrinter) widths(c canvas, rows [][]tableCell, cols int) []int {
	widths := make([]int, cols)

	for _, row := range rows {
		for col, cell := range row {
			for _, line := range cell {
				widths[col] = max(widths[col], text.Width(line))
			}
		}
	}

	natural := append([]int(nil), widths...)

	for col := range widths {
		widths[col] = max(widths[col], at(p.ColumnMinWidths, col))
	}

	if c.width <= 0 {
		return widths
	}

	avail := c.width - (cols-1)*text.Width(p.Separator)
	if p.Boxed {
		avail -= 2 + 2*p.Padding
	}

	total := 0
	for _, w := range widths {
		total += w
	}

	// Minimum widths only pad; drop them first when space is short.
	for col := range widths {
		if total <= avail {
			break
		}

		cut := min(widths[col]-natural[col], total-avail)
		widths[col] -= cut
		total -= cut
	}

	for total > avail {
		widest := -1

		for col, w := range widths {
			if w > min(natural[col], 3) && (widest < 0 || w > widths[widest]) {
				widest = col
			}
		}

		if widest < 0 {
			break
		}

		widths[widest]--
		total--
	}

	return widths
}

// alignments resolves the alignment of every column.
func (p TablePrinter) alignments(rows [][]tableCell, cols int) []Alignment {
	aligns := make([]Alignment, cols)

	for col := range aligns {
		a := at(p.ColumnAlignments, col)
		if a == AlignAuto {
			a = p.Alignment
		}

		if a == AlignAuto {
			a = AlignLeft
			if numericColumn(rows, col, p.HasHeader) {
				a = AlignRight
			}
		}

		aligns[col] = a
	}

	return aligns
}

// rowStyle returns the style a whole row gets on top of the cell styles,
// or the zero Style.
func (p TablePrinter) rowStyle(body int, header bool) Style {
	switch {
	case header:
		return Style{}
	case at(p.RowStyles, body) != (Style{}):
		return at(p.RowStyles, body)
	case body%2 == 1:
		return p.AlternateRowStyle
	default:
		return Style{}
	}
}

// renderRow returns the lines of one row, without frame.
func (p TablePrinter) renderRow(c canvas, row []tableCell, widths []int, aligns []Alignment, header bool, rowStyle Style, sep string) []string {
	height := 0
	for _, cell := range row {
		height = max(height, len(cell))
	}

	// A row style with a background also fills the padding and the
	// separators, so the row reads as one band.
	filled := rowStyle.Bg != (Tone{})
	if filled {
		sep = c.style(rowStyle, p.Separator)
	}

	lines := make([]string, height)

	for i := range height {
		parts := make([]string, len(row))

		for col, cell := range row {
			s := c.cut(at(cell, i), widths[col])
			st := p.cellStyle(col, header, rowStyle)

			if filled {
				parts[col] = c.style(st, aligns[col].pad(s, widths[col]))
				continue
			}

			parts[col] = aligns[col].pad(c.style(st, s), widths[col])
		}

		line := strings.Join(parts, sep)
		if !filled && !p.Boxed {
			line = strings.TrimRight(line, " ")
		}

		lines[i] = line
	}

	return lines
}

func (p TablePrinter) cellStyle(col int, header bool, rowStyle Style) Style {
	if header {
		return p.HeaderStyle
	}

	if rowStyle != (Style{}) {
		return rowStyle
	}

	if st := at(p.ColumnStyles, col); st != (Style{}) {
		return st
	}

	return p.CellStyle
}

// numberPattern matches numbers and sizes: "42", "-3.5", "1,024", "12%",
// "1.20 GB", "393 MB", "$4.20", "12ms", "1m 34s".
var numberPattern = regexp.MustCompile(`^[-+]?[$€£¥]?\d[\d,_']*(\.\d+)?\s?(%|[A-Za-zµ]{1,3}(/s)?)?(\s\d+(\.\d+)?\s?[A-Za-zµ]{1,3})*$`)

// placeholders stand for a missing number and do not make a column text.
var placeholders = map[string]bool{"-": true, "\u2013": true, "\u2014": true, "?": true, "n/a": true, "N/A": true}

// numericColumn reports whether every non-empty body cell of column col is
// a number or size, and at least one is.
func numericColumn(rows [][]tableCell, col int, hasHeader bool) bool {
	found := false

	for r, row := range rows {
		if r == 0 && hasHeader {
			continue
		}

		for _, line := range row[col] {
			line = strings.TrimSpace(color.Strip(line))
			if line == "" || placeholders[line] {
				continue
			}

			if !numberPattern.MatchString(line) {
				return false
			}

			found = true
		}
	}

	return found
}

// upperVisible uppercases s but leaves escape sequences alone.
func upperVisible(s string) string {
	var b strings.Builder

	last := 0
	for _, loc := range text.Escapes(s) {
		b.WriteString(strings.ToUpper(s[last:loc[0]]))
		b.WriteString(s[loc[0]:loc[1]])
		last = loc[1]
	}

	b.WriteString(strings.ToUpper(s[last:]))

	return b.String()
}
