package pterm

import (
	"io"
	"strings"

	"github.com/mattn/go-runewidth"

	"github.com/pterm/pterm/internal"
)

// Panel contains the data, which should be printed inside a PanelPrinter.
type Panel struct {
	Data string
}

// Panels is a two dimensional coordinate system for Panel.
type Panels [][]Panel

// DefaultPanel is the default PanelPrinter.
var DefaultPanel = PanelPrinter{
	Padding: 1,
}

// PanelPrinter prints content in boxes.
type PanelPrinter struct {
	Panels          Panels
	Padding         int
	BottomPadding   int
	SameColumnWidth bool
	BoxPrinter      BoxPrinter
	Writer          io.Writer
}

// WithPanels returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithPanels(panels Panels) *PanelPrinter {
	p.Panels = panels
	return &p
}

// WithPadding returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithPadding(padding int) *PanelPrinter {
	if padding < 0 {
		padding = 0
	}

	p.Padding = padding

	return &p
}

// WithBottomPadding returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithBottomPadding(bottomPadding int) *PanelPrinter {
	if bottomPadding < 0 {
		bottomPadding = 0
	}

	p.BottomPadding = bottomPadding

	return &p
}

// WithSameColumnWidth returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithSameColumnWidth(b ...bool) *PanelPrinter {
	p.SameColumnWidth = internal.WithBoolean(b)
	return &p
}

// WithBoxPrinter returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithBoxPrinter(boxPrinter BoxPrinter) *PanelPrinter {
	p.BoxPrinter = boxPrinter
	return &p
}

// WithWriter sets the custom Writer.
func (p PanelPrinter) WithWriter(writer io.Writer) *PanelPrinter {
	p.Writer = writer
	return &p
}

func (p PanelPrinter) getRawOutput() string {
	var ret strings.Builder

	for _, panel := range p.Panels {
		for _, panel2 := range panel {
			ret.WriteString(panel2.Data)
			ret.WriteString("\n\n")
		}

		ret.WriteByte('\n')
	}

	return ret.String()
}

// Srender renders the Template as a string.
func (p PanelPrinter) Srender() (string, error) {
	if rawOutput() {
		return p.getRawOutput(), nil
	}

	p.Panels = p.preparedPanels()

	columnWidths := p.sameColumnWidths()

	var ret strings.Builder

	for _, row := range p.Panels {
		p.renderRow(&ret, row, columnWidths)
	}

	return ret.String(), nil
}

// preparedPanels returns a copy of the configured panels with trailing
// newlines trimmed, the box applied and the bottom padding appended.
// Rendering must not mutate the caller's Panels.
func (p PanelPrinter) preparedPanels() Panels {
	panels := make(Panels, len(p.Panels))
	for i, row := range p.Panels {
		panels[i] = append([]Panel(nil), row...)
	}

	boxed := p.BoxPrinter != (BoxPrinter{})

	for i, row := range panels {
		for j := range row {
			row[j].Data = strings.TrimSuffix(row[j].Data, "\n")

			if boxed {
				row[j].Data = p.BoxPrinter.Sprint(row[j].Data)
			}

			if i != len(panels)-1 {
				row[j].Data += strings.Repeat("\n", p.BottomPadding)
			}
		}
	}

	return panels
}

// sameColumnWidths returns the width of the widest panel per column, so all
// panels of a column can be padded to the same width. Only used with
// SameColumnWidth.
func (p PanelPrinter) sameColumnWidths() map[int]int {
	columnWidths := make(map[int]int)

	if p.SameColumnWidth {
		for _, row := range p.Panels {
			for i, panel := range row {
				columnWidths[i] = max(columnWidths[i], internal.GetStringMaxWidth(panel.Data))
			}
		}
	}

	return columnWidths
}

// renderRow writes one row of panels side by side: every panel is padded to
// its width (or its column's width with SameColumnWidth) plus the configured
// padding, and shorter panels are filled up with blank lines.
func (p PanelPrinter) renderRow(ret *strings.Builder, row []Panel, columnWidths map[int]int) {
	panelLines := make([][]string, len(row))
	panelWidths := make([]int, len(row))

	var maxHeight int

	for i, panel := range row {
		// Terminate every line's styling so one panel's colors cannot leak
		// into its right-hand neighbor.
		data := strings.ReplaceAll(panel.Data, "\n", Reset.Sprint()+"\n")

		panelLines[i] = strings.Split(data, "\n")
		maxHeight = max(maxHeight, len(panelLines[i]))

		if p.SameColumnWidth {
			panelWidths[i] = columnWidths[i]
		} else {
			panelWidths[i] = internal.GetStringMaxWidth(data)
		}
	}

	for line := 0; line < maxHeight; line++ {
		for i := range row {
			var cell string
			if line < len(panelLines[i]) {
				cell = panelLines[i][line]
			}

			cellWidth := runewidth.StringWidth(RemoveColorFromString(cell))
			if cellWidth < panelWidths[i] {
				cell += strings.Repeat(" ", panelWidths[i]-cellWidth)
			}

			ret.WriteString(cell)
			ret.WriteString(strings.Repeat(" ", p.Padding))
		}

		ret.WriteByte('\n')
	}
}

// Render prints the Template to the terminal.
func (p PanelPrinter) Render() error {
	s, err := p.Srender()
	if err != nil {
		return err
	}

	Fprintln(p.Writer, s)

	return nil
}
