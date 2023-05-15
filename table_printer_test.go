package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestTablePrinter_NilPrint(t *testing.T) {
	p := pterm.TablePrinter{}
	p.Render()
}

func TestTablePrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.TablePrinter{}, "WithCSVReader")
}
