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
			{{Data: Sprint(a)}},
		}
		p := PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderMultiplePanels(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		panels := Panels{
			{{Data: Sprint("a\nbc\ndef")}, {Data: Sprint("abcd")}},
			{{Data: Sprint(a)}},
		}
		p := PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderMultiplePanelsWithBorder(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		panels := Panels{
			{{Data: Sprint("a\nbc\ndef")}, {Data: Sprint("abcd")}},
			{{Data: Sprint(a)}},
		}
		p := PanelPrinter{}.WithPanels(panels).WithBoxPrinter(DefaultBox)
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderWithSameColumnWidth(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		panels := Panels{
			{{Data: Sprint(a)}},
			{{Data: Sprint("test")}},
			{{Data: Sprint("Hello, World!")}},
		}
		p := PanelPrinter{}.WithPanels(panels).WithSameColumnWidth()
		err := p.Render()
		assert.NoError(t, err)
	})
}

func TestPanelPrinter_RenderWithBottomPadding(t *testing.T) {
	internal.TestPrintContains(t, func(w io.Writer, a interface{}) {
		panels := Panels{
			{{Data: Sprint(a)}},
			{{Data: Sprint("test")}},
			{{Data: Sprint("Hello, World!")}},
		}
		p := PanelPrinter{}.WithPanels(panels).WithBottomPadding(1)
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

func TestPanelPrinter_WithInvalidPadding(t *testing.T) {
	padding := -5
	p := PanelPrinter{}
	p2 := p.WithPadding(padding)

	assert.Equal(t, 0, p2.Padding)
	assert.Empty(t, p.Padding)
}

func TestPanelPrinter_WithBottomPadding(t *testing.T) {
	padding := 1337
	p := PanelPrinter{}
	p2 := p.WithBottomPadding(padding)

	assert.Equal(t, padding, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestPanelPrinter_WithInvalidBottomPadding(t *testing.T) {
	padding := -5
	p := PanelPrinter{}
	p2 := p.WithBottomPadding(padding)

	assert.Equal(t, 0, p2.BottomPadding)
	assert.Empty(t, p.BottomPadding)
}

func TestPanelPrinter_WithSameColumnWidth(t *testing.T) {
	p := PanelPrinter{}
	p2 := p.WithSameColumnWidth()

	assert.True(t, p2.SameColumnWidth)
	assert.False(t, p.SameColumnWidth)
}

func TestPanelPrinter_WithBoxPrinter(t *testing.T) {
	p := PanelPrinter{}
	p2 := p.WithBoxPrinter(DefaultBox)

	assert.Equal(t, DefaultBox, p2.BoxPrinter)
	assert.Empty(t, p.BoxPrinter)
}
