package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestTableGroupGroupPrinter_NilPrint(t *testing.T) {
	p := pterm.TableGroup{}
	_ = p.Render()
}

func TestTableGroupPrinter_Append(t *testing.T) {
	// pterm.proxyToDevNull()

	d3 := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}

	d2 := pterm.TableData{
		{"Firstname", "Lastname"},
		{"Paul", "Dean"},
		{"Callie", "Mckay"},
		{"Libby", "Camacho"},
	}

	t1 := pterm.DefaultTable.WithHasHeader().WithData(d3)
	t2 := pterm.DefaultTable.WithHasHeader().WithData(d2)

	tg := pterm.DefaultGroupTable
	tg.Append(t1, t2)

	testza.AssertEqual(t, d3, tg.Tables[0].Data)
	testza.AssertEqual(t, d2, tg.Tables[1].Data)
}

func TestTableGroupPrinter_WithBoxed(t *testing.T) {
	_, err := pterm.DefaultGroupTable.WithBoxed(true).Srender()
	if err != nil {
		t.Error(err)
	}
}

func TestTableGroupPrinter_WithShare(t *testing.T) {
	d3 := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}

	d2 := pterm.TableData{
		{"Firstname", "Lastname"},
		{"Paul", "Dean"},
		{"Callie", "Mckay"},
		{"Libby", "Camacho"},
	}

	t1 := pterm.DefaultTable.WithHasHeader().WithData(d3).WithSeparator("*")
	t2 := pterm.DefaultTable.WithHasHeader().WithData(d2).WithSeparator("^")

	tg := pterm.DefaultGroupTable.WithShare(true)
	tg.Append(t1, t2).Srender()

	testza.AssertEqual(t, tg.Tables[0].Separator, tg.Tables[1].Separator)
}

func TestTableGroupPrinter_WithNoShare(t *testing.T) {
	d3 := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}

	d2 := pterm.TableData{
		{"Firstname", "Lastname"},
		{"Paul", "Dean"},
		{"Callie", "Mckay"},
		{"Libby", "Camacho"},
	}

	p1 := pterm.TablePrinter{}
	t1 := p1.WithHasHeader().WithData(d3).WithSeparator("*")
	p2 := pterm.TablePrinter{}
	t2 := p2.WithHasHeader().WithData(d2).WithSeparator("^")

	tg := pterm.TableGroup{}
	tg.Append(t1, t2).Srender()

	testza.AssertNotEqual(t, tg.Tables[0].Separator, tg.Tables[1].Separator)
}

func TestTableGroupPrinter_WithHasHeader(t *testing.T) {
	tg := pterm.TableGroup{}
	p2 := tg.WithHasHeader()

	testza.AssertTrue(t, p2.HasHeader)
}

func TestTableGroupPrinter_WithHeaderStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TableGroup{}
	p2 := p.WithHeaderStyle(s)

	testza.AssertEqual(t, s, p2.HeaderStyle)
}

func TestTableGroupPrinter_WithSeparator(t *testing.T) {
	p := pterm.TableGroup{}
	p2 := p.WithSeparator("-")

	testza.AssertEqual(t, "-", p2.Separator)
}

func TestTableGroupPrinter_WithSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TableGroup{}
	p2 := p.WithSeparatorStyle(s)

	testza.AssertEqual(t, s, p2.SeparatorStyle)
}

func TestTableGroupPrinter_WithHorizontalSeparator(t *testing.T) {
	p := pterm.TableGroup{}
	p2 := p.WithHorizontalSeparator("-")

	testza.AssertEqual(t, "-", p2.HSeparator)
}

func TestTableGroupPrinter_WithHorizontalSeparatorWrongSize(t *testing.T) {
	p := pterm.DefaultGroupTable
	p2 := p.WithHorizontalSeparator(" - ")

	testza.AssertEqual(t, "-", p2.HSeparator)
}

func TestTableGroupPrinter_WithHorizontalSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TableGroup{}
	p2 := p.WithHorizontalSeparatorStyle(s)

	testza.AssertEqual(t, s, p2.HSeparatorStyle)
}

func TestTableGroupPrinter_WithStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TableGroup{}
	p2 := p.WithStyle(s)

	testza.AssertEqual(t, s, p2.Style)
}
