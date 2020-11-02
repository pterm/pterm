package pterm

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestPanelPrinterNilPrint(t *testing.T) {
	p := PanelPrinter{}
	err := p.Render()
	assert.NoError(t, err)
}

func TestPanelPrinterNilPrintWithPanels(t *testing.T) {
	panels := Panels{
		{
			{Data: "Hello, World"},
		},
	}
	p := PanelPrinter{}.WithPanels(panels)
	err := p.Render()
	assert.NoError(t, err)
}

func TestPanelPrinter_Render(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		panels := Panels{
			{
				{Data: Sprint(a)},
			},
		}
		p := PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_WithPanels(t *testing.T) {
	panels := Panels{
		{
			{Data: "Hello, World!"},
		},
	}
	p := PanelPrinter{}
	p2 := p.WithPanels(panels)

	assert.Equal(t, panels, p2.Panels)
	assert.Empty(t, p.Panels)
}

func TestPanelPrinter_WithPadding(t *testing.T) {
	padding := 1337
	p := PanelPrinter{}
	p2 := p.WithPadding(padding)

	assert.Equal(t, padding, p2.Padding)
	assert.Empty(t, p.Padding)
}
