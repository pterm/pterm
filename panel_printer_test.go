package pterm_test

import (
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestPanelPrinterNilPrint(t *testing.T) {
	p := pterm.PanelPrinter{}
	err := p.Render()
	testza.AssertNoError(t, err)
}

func TestPanelPrinterNilPrintWithPanels(t *testing.T) {
	panels := pterm.Panels{
		{
			{Data: "Hello, World"},
		},
	}
	p := pterm.PanelPrinter{}.WithPanels(panels)
	err := p.Render()
	testza.AssertNoError(t, err)
}

func TestPanelPrinter_Render(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		testza.AssertNoError(t, err)
	})
}

func TestPanelPrinter_RenderMultiplePanels(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint("a\nbc\ndef")}, {Data: pterm.Sprint("abcd")}},
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		testza.AssertNoError(t, err)
	})
}

func TestPanelPrinter_RenderMultiplePanelsWithBorder(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint("a\nbc\ndef")}, {Data: pterm.Sprint("abcd")}},
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels).WithBoxPrinter(pterm.DefaultBox)
		err := p.Render()
		testza.AssertNoError(t, err)
	})
}

func TestPanelPrinter_RenderWithSameColumnWidth(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
			{{Data: pterm.Sprint("test")}},
			{{Data: pterm.Sprint("Hello, World!")}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels).WithSameColumnWidth()
		err := p.Render()
		testza.AssertNoError(t, err)
	})
}

func TestPanelPrinter_RenderWithBottomPadding(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
			{{Data: pterm.Sprint("test")}},
			{{Data: pterm.Sprint("Hello, World!")}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels).WithBottomPadding(1)
		err := p.Render()
		testza.AssertNoError(t, err)
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

	testza.AssertEqual(t, panels, p2.Panels)
	testza.AssertZero(t, p.Panels)
}

func TestPanelPrinter_WithPadding(t *testing.T) {
	padding := 1337
	p := pterm.PanelPrinter{}
	p2 := p.WithPadding(padding)

	testza.AssertEqual(t, padding, p2.Padding)
	testza.AssertZero(t, p.Padding)
}

func TestPanelPrinter_WithInvalidPadding(t *testing.T) {
	padding := -5
	p := pterm.PanelPrinter{}
	p2 := p.WithPadding(padding)

	testza.AssertEqual(t, 0, p2.Padding)
	testza.AssertZero(t, p.Padding)
}

func TestPanelPrinter_WithBottomPadding(t *testing.T) {
	padding := 1337
	p := pterm.PanelPrinter{}
	p2 := p.WithBottomPadding(padding)

	testza.AssertEqual(t, padding, p2.BottomPadding)
	testza.AssertZero(t, p.BottomPadding)
}

func TestPanelPrinter_WithInvalidBottomPadding(t *testing.T) {
	padding := -5
	p := pterm.PanelPrinter{}
	p2 := p.WithBottomPadding(padding)

	testza.AssertEqual(t, 0, p2.BottomPadding)
	testza.AssertZero(t, p.BottomPadding)
}

func TestPanelPrinter_WithSameColumnWidth(t *testing.T) {
	p := pterm.PanelPrinter{}
	p2 := p.WithSameColumnWidth()

	testza.AssertTrue(t, p2.SameColumnWidth)
	testza.AssertFalse(t, p.SameColumnWidth)
}

func TestPanelPrinter_WithBoxPrinter(t *testing.T) {
	p := pterm.PanelPrinter{}
	p2 := p.WithBoxPrinter(pterm.DefaultBox)

	testza.AssertEqual(t, pterm.DefaultBox, p2.BoxPrinter)
	testza.AssertZero(t, p.BoxPrinter)
}

func TestPanelPrinter_WithWriter(t *testing.T) {
	p := pterm.PanelPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}
