package pterm

import (
	"testing"
)

func TestTablePrinter_NilPrint(t *testing.T) {
	p := TablePrinter{}
	p.Render()
}
