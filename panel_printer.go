package pterm

import (
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
	Panels      Panels
	Padding     int
	Border      bool
	BorderStyle *Style
}

// WithPanels returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithPanels(panels Panels) *PanelPrinter {
	p.Panels = panels
	return &p
}

// WithPadding returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithPadding(padding int) *PanelPrinter {
	p.Padding = padding
	return &p
}

// WithBorder returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithBorder(b ...bool) *PanelPrinter {
	b2 := internal.WithBoolean(b)
	p.Border = b2
	return &p
}

// WithBorderStyle returns a new PanelPrinter with specific options.
func (p PanelPrinter) WithBorderStyle(style *Style) *PanelPrinter {
	p.BorderStyle = style
	return &p
}

// Srender renders the Template as a string.
func (p PanelPrinter) Srender() (string, error) {
	var ret string

	for _, boxLine := range p.Panels {
		var maxHeight int

		var renderedPanels []string

		for _, box := range boxLine {
			renderedPanels = append(renderedPanels, box.Data)
		}

		if p.Border {
			var panels []string
			var tmh int
			for _, panel := range renderedPanels {
				h := len(strings.Split(panel, "\n"))
				if h > tmh {
					tmh = h
				}
			}
			for _, panel := range renderedPanels {
				panel += strings.Repeat("\n", tmh-len(strings.Split(panel, "\n")))
				s, _ := DefaultPanel.WithPanels(Panels{[]Panel{{Data: panel}}}).Srender()
				panels = append(panels, s)
			}
			for i, panel := range panels {
				renderedPanels[i] = boxed(panel)
			}

		}

		for _, box := range renderedPanels {
			height := len(strings.Split(box, "\n"))
			if height > maxHeight {
				maxHeight = height
			}
		}

		for i := 0; i < maxHeight; i++ {
			for _, letter := range renderedPanels {
				var letterLine string
				letterLines := strings.Split(letter, "\n")
				maxLetterWidth := internal.GetStringMaxWidth(letter)
				if len(letterLines) > i {
					letterLine = letterLines[i]
				}
				letterLineLength := runewidth.StringWidth(letterLine)
				if letterLineLength < maxLetterWidth {
					letterLine += strings.Repeat(" ", maxLetterWidth-letterLineLength)
				}
				letterLine += strings.Repeat(" ", p.Padding)
				ret += letterLine
			}
			ret += "\n"
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

func boxed(s string) string {
	maxWidth := internal.GetStringMaxWidth(s)
	topLine := "┌" + strings.Repeat("─", maxWidth) + "┐"

	ss := strings.Split(s, "\n")
	for i, s2 := range ss {
		if i != len(ss)-1 {
			ss[i] = "|" + s2 + "|"
		}
	}

	bottomLine := "└" + strings.Repeat("─", maxWidth) + "┘"

	return topLine + "\n" + strings.Join(ss, "\n") + bottomLine
}
