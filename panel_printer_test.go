package pterm_test

import (
	"io"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestPanelPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.PanelPrinter{})
}

func TestPanelPrinterNilPrint(t *testing.T) {
	p := pterm.PanelPrinter{}
	err := p.Render()
	testza.AssertNoError(t, err)
}

func TestPanelPrinter_Render(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		panels := pterm.Panels{
			{{Data: pterm.Sprint(a)}},
		}
		p := pterm.PanelPrinter{}.WithPanels(panels)
		err := p.Render()
		testza.AssertNoError(t, err)
	})
}
