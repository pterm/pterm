package pterm_test

import (
	"io"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestPanelPrinterNilPrint(t *testing.T) {
	p := pterm.PanelPrinter{}
	err := p.Render()
	assert.NoError(t, err)
}

func TestPanelPrinterNilPrintWithPanels(t *testing.T) {
	panels := pterm.Panels{
		{
			{Data: "Hello, World"},
		},
	}
	p := pterm.PanelPrinter{}.WithPanels(panels)
	err := p.Render()
	assert.NoError(t, err)
}

func TestPanelPrinter_Render(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderMultiplePanels(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint("a\nbc\ndef")}, {Data: pterm.Sprint("abcd")}},
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderMultiplePanelsWithBorder(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint("a\nbc\ndef")}, {Data: pterm.Sprint("abcd")}},
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels).WithBoxPrinter(pterm.DefaultBox)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderWithSameColumnWidth(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
			{{Data: pterm.Sprint("test")}},
			{{Data: pterm.Sprint("Hello, World!")}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels).WithSameColumnWidth()
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderWithBottomPadding(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
			{{Data: pterm.Sprint("test")}},
			{{Data: pterm.Sprint("Hello, World!")}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels).WithBottomPadding(1)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_WithPanels(t *testing.T) {
	panels := pterm.Panels{
		{
			{Data: "Hello, World!"},
		},
	}
	p := pterm.PanelPrinter{}
	p2 := p.WithPanels(panels)

	assert.Equal(t, panels, p2.Panels)
	assert.Empty(t, p.Panels)
}

func TestPanelPrinter_WithPadding(t *testing.T) {
	padding := 1337
	p := pterm.PanelPrinter{}
	p2 := p.WithPadding(padding)

	assert.Equal(t, padding, p2.Padding)
	assert.Empty(t, p.Padding)
}

func TestPanelPrinter_WithInvalidPadding(t *testing.T) {
	padding := -5
	p := pterm.PanelPrinter{}
	p2 := p.WithPadding(padding)

	assert.Equal(t, 0, p2.Padding)
	assert.Empty(t, p.Padding)
}

func TestPanelPrinter_WithBottomPadding(t *testing.T) {
	padding := 1337
	p := pterm.PanelPrinter{}
	p2 := p.WithBottomPadding(padding)

	assert.Equal(t, padding, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestPanelPrinter_WithInvalidBottomPadding(t *testing.T) {
	padding := -5
	p := pterm.PanelPrinter{}
	p2 := p.WithBottomPadding(padding)

	assert.Equal(t, 0, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestPanelPrinter_WithSameColumnWidth(t *testing.T) {
	p := pterm.PanelPrinter{}
	p2 := p.WithSameColumnWidth()

	assert.True(t, p2.SameColumnWidth)
	assert.False(t, p.SameColumnWidth)
}

func TestPanelPrinter_WithBoxPrinter(t *testing.T) {
	p := pterm.PanelPrinter{}
	p2 := p.WithBoxPrinter(pterm.DefaultBox)

	assert.Equal(t, pterm.DefaultBox, p2.BoxPrinter)
	assert.Empty(t, p.BoxPrinter)
}
