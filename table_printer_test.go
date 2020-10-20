package pterm

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestTablePrinterNilPrint(t *testing.T) {
	p := Table{}
	p.Render()
}

func TestTable_Render(t *testing.T) {
	DefaultTable.WithHasHeader().WithData(TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}).Render()
}

func TestTable_WithCSVReader(t *testing.T) {
	r := csv.NewReader(os.Stdin)
	p := Table{}
	p2 := p.WithCSVReader(r)
	p2.Srender()
}

func TestTable_WithData(t *testing.T) {
	d := TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}
	p := Table{}
	p2 := p.WithData(d)

	assert.Equal(t, d, p2.Data)
}

func TestTable_WithHasHeader(t *testing.T) {
	p := Table{}
	p2 := p.WithHasHeader()

	assert.True(t, p2.HasHeader)
}

func TestTable_WithHeaderStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := Table{}
	p2 := p.WithHeaderStyle(s)

	assert.Equal(t, s, p2.HeaderStyle)
}

func TestTable_WithSeparator(t *testing.T) {
	p := Table{}
	p2 := p.WithSeparator("-")

	assert.Equal(t, "-", p2.Separator)
}

func TestTable_WithSeparatorStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := Table{}
	p2 := p.WithSeparatorStyle(s)

	assert.Equal(t, s, p2.SeparatorStyle)
}

func TestTable_WithStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := Table{}
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
}
