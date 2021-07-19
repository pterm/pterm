package pterm_test

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestTablePrinter_NilPrint(t *testing.T) {
	p := pterm.TablePrinter{}
	p.Render()
}

func TestTablePrinter_Render(t *testing.T) {
	proxyToDevNull()
	pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).Render()
}

func TestTablePrinter_WithCSVReader(t *testing.T) {
	r := csv.NewReader(os.Stdin)
	p := pterm.TablePrinter{}
	p2 := p.WithCSVReader(r)
	p2.Srender()
}

func TestTablePrinter_WithBoxed(t *testing.T) {
	_, err := pterm.DefaultTable.WithBoxed().Srender()
	if err != nil {
		t.Error(err)
	}
}

func TestTablePrinter_WithData(t *testing.T) {
	proxyToDevNull()
	d := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}
	p := pterm.TablePrinter{}
	p2 := p.WithData(d)

	assert.Equal(t, d, p2.Data)
}

func TestTablePrinter_WithHasHeader(t *testing.T) {
	p := pterm.TablePrinter{}
	p2 := p.WithHasHeader()

	assert.True(t, p2.HasHeader)
}

func TestTablePrinter_WithHeaderStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithHeaderStyle(s)

	assert.Equal(t, s, p2.HeaderStyle)
}

func TestTablePrinter_WithSeparator(t *testing.T) {
	p := pterm.TablePrinter{}
	p2 := p.WithSeparator("-")

	assert.Equal(t, "-", p2.Separator)
}

func TestTablePrinter_WithSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithSeparatorStyle(s)

	assert.Equal(t, s, p2.SeparatorStyle)
}

func TestTablePrinter_WithStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
}
