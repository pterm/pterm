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
	var ret string
	for _, panel := range p.Panels {
		for _, panel2 := range panel {
			ret += panel2.Data + "\n\n"
		}
		ret += "\n"
	}
	return ret
}

// Srender renders the Template as a string.
func (p PanelPrinter) Srender() (string, error) {
	var ret string

	if RawOutput {
		return p.getRawOutput(), nil
	}

	for i := range p.Panels {
		for i2 := range p.Panels[i] {
			p.Panels[i][i2].Data = strings.TrimSuffix(p.Panels[i][i2].Data, "\n")
		}
	}

	if p.BoxPrinter != (BoxPrinter{}) {
		for i := range p.Panels {
			for i2 := range p.Panels[i] {
				p.Panels[i][i2].Data = p.BoxPrinter.Sprint(p.Panels[i][i2].Data)
			}
		}
	}

	for i := range p.Panels {
		if len(p.Panels)-1 != i {
			for i2 := range p.Panels[i] {
				p.Panels[i][i2].Data += strings.Repeat("\n", p.BottomPadding)
			}
		}
	}

	columnMaxHeightMap := make(map[int]int)

	if p.SameColumnWidth {
		for _, panel := range p.Panels {
			for i, p2 := range panel {
				if columnMaxHeightMap[i] < internal.GetStringMaxWidth(p2.Data) {
					columnMaxHeightMap[i] = internal.GetStringMaxWidth(p2.Data)
				}
			}
		}
	}

	for _, boxLine := range p.Panels {
		var maxHeight int

		var renderedPanels []string

		for _, box := range boxLine {
			renderedPanels = append(renderedPanels, box.Data)
		}

		for i, panel := range renderedPanels {
			renderedPanels[i] = strings.ReplaceAll(panel, "\n", Reset.Sprint()+"\n")
		}

		for _, box := range renderedPanels {
			height := len(strings.Split(box, "\n"))
			if height > maxHeight {
				maxHeight = height
			}
		}

		for i := 0; i < maxHeight; i++ {
			if maxHeight != i {
				for j, letter := range renderedPanels {
					var letterLine string
					letterLines := strings.Split(letter, "\n")
					var maxLetterWidth int
					if !p.SameColumnWidth {
						maxLetterWidth = internal.GetStringMaxWidth(letter)
					}
					if len(letterLines) > i {
						letterLine = letterLines[i]
					}
					letterLineLength := runewidth.StringWidth(RemoveColorFromString(letterLine))
					if !p.SameColumnWidth {
						if letterLineLength < maxLetterWidth {
							letterLine += strings.Repeat(" ", maxLetterWidth-letterLineLength)
						}
					} else {
						if letterLineLength < columnMaxHeightMap[j] {
							letterLine += strings.Repeat(" ", columnMaxHeightMap[j]-letterLineLength)
						}
					}
					letterLine += strings.Repeat(" ", p.Padding)
					ret += letterLine
				}
				ret += "\n"
			}
		}
	}

	return ret, nil
}

// Render prints the Template to the terminal.
func (p PanelPrinter) Render() error {
	s, _ := p.Srender()
	Println(s)

	return nil
}
