package pterm

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTablePrinter_NilPrint(t *testing.T) {
	p := TablePrinter{}
	p.Render()
}

func TestTablePrinter_Render(t *testing.T) {
	proxyToDevNull()
	DefaultTable.WithHasHeader().WithData(TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).Render()
}

func TestTablePrinter_WithCSVReader(t *testing.T) {
	r := csv.NewReader(os.Stdin)
	p := TablePrinter{}
	p2 := p.WithCSVReader(r)
	p2.Srender()
}

func TestTablePrinter_WithData(t *testing.T) {
	proxyToDevNull()
	d := TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}
	p := TablePrinter{}
	p2 := p.WithData(d)

	assert.Equal(t, d, p2.Data)
}

func TestTablePrinter_WithHasHeader(t *testing.T) {
	p := TablePrinter{}
	p2 := p.WithHasHeader()

	assert.True(t, p2.HasHeader)
}

func TestTablePrinter_WithHeaderStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := TablePrinter{}
	p2 := p.WithHeaderStyle(s)

	assert.Equal(t, s, p2.HeaderStyle)
}

func TestTablePrinter_WithSeparator(t *testing.T) {
	p := TablePrinter{}
	p2 := p.WithSeparator("-")

	assert.Equal(t, "-", p2.Separator)
}
