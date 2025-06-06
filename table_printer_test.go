package pterm_test

import (
	"encoding/csv"
	"io"
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestTablePrinter_NilPrint(t *testing.T) {
	p := pterm.TablePrinter{}
	p.Render()
}

func TestTablePrinter_Render(t *testing.T) {
	d := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}
	// WithLeftAlignment
	printer := pterm.DefaultTable.WithHasHeader().WithLeftAlignment().WithData(d)
	content, err := printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
	// WithRightAlignment
	printer = pterm.DefaultTable.WithHasHeader().WithRightAlignment().WithData(d)
	content, err = printer.Srender()

	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
}

func TestTablePrinterWithAlternateStyle_Render(t *testing.T) {
	d := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}

	// Define the alternate row style
	alternateStyle := pterm.NewStyle(pterm.BgDarkGray)

	// Create a printer with the alternate row style
	printer := pterm.DefaultTable.WithHasHeader().WithAlternateRowStyle(alternateStyle).WithData(d)
	content, err := printer.Srender()
	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
}

func TestTablePrinterWithRowSeparators_Render(t *testing.T) {
	d := pterm.TableData{
		{"Firstname", "Lastname", "Email"},
		{"Paul", "Dean", "nisi.dictum.augue@velitAliquam.co.uk"},
		{"Callie", "Mckay", "egestas.nunc.sed@est.com"},
		{"Libby", "Camacho", "aliquet.lobortis@semper.com"},
	}
	// WithHeaderSeparator
	printer := pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("=").WithData(d)
	content, err := printer.Srender()
	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
	// WithRowSeparator
	printer = pterm.DefaultTable.WithHasHeader().WithRowSeparator("-").WithData(d)
	content, err = printer.Srender()
	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
	// WithHeaderRowSeparator & WithRowSeparator
	printer = pterm.DefaultTable.WithHasHeader().WithHeaderRowSeparator("=").WithRowSeparator("-").WithData(d)
	content, err = printer.Srender()
	testza.AssertNoError(t, err)
	testza.AssertNotNil(t, content)
}

func TestTablePrinter_WithCSVReader(t *testing.T) {
	content := captureStdout(func(w io.Writer) {
		r := csv.NewReader(&outBuf)
		p := pterm.TablePrinter{}
		p.WithCSVReader(r)
	})
	testza.AssertNotNil(t, content)
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

	testza.AssertEqual(t, d, p2.Data)
}

func TestTablePrinter_WithHasHeader(t *testing.T) {
	p := pterm.TablePrinter{}
	p2 := p.WithHasHeader()

	testza.AssertTrue(t, p2.HasHeader)
}

func TestTablePrinter_WithHeaderStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithHeaderStyle(s)

	testza.AssertEqual(t, s, p2.HeaderStyle)
}

func TestTablePrinter_WithSeparator(t *testing.T) {
	p := pterm.TablePrinter{}
	p2 := p.WithSeparator("-")

	testza.AssertEqual(t, "-", p2.Separator)
}

func TestTablePrinter_WithSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithSeparatorStyle(s)

	testza.AssertEqual(t, s, p2.SeparatorStyle)
}

func TestTablePrinter_WithHeaderRowSeparator(t *testing.T) {
	p := pterm.TablePrinter{}
	p2 := p.WithHeaderRowSeparator("-")

	testza.AssertEqual(t, "-", p2.HeaderRowSeparator)
}

func TestTablePrinter_WithHeaderRowSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithHeaderRowSeparatorStyle(s)

	testza.AssertEqual(t, s, p2.HeaderRowSeparatorStyle)
}

func TestTablePrinter_WithRowSeparator(t *testing.T) {
	p := pterm.TablePrinter{}
	p2 := p.WithRowSeparator("-")

	testza.AssertEqual(t, "-", p2.RowSeparator)
}

func TestTablePrinter_WithRowSeparatorStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithRowSeparatorStyle(s)

	testza.AssertEqual(t, s, p2.RowSeparatorStyle)
}

func TestTablePrinter_WithStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.TablePrinter{}
	p2 := p.WithStyle(s)

	testza.AssertEqual(t, s, p2.Style)
}

func TestTablePrinter_WithLeftAlignment(t *testing.T) {
	s := true
	p := pterm.TablePrinter{}
	p2 := p.WithLeftAlignment(s)

	testza.AssertEqual(t, s, p2.LeftAlignment)
}

func TestTablePrinter_WithRightAlignment(t *testing.T) {
	s := true
	p := pterm.TablePrinter{}
	p2 := p.WithRightAlignment(s)

	testza.AssertEqual(t, s, p2.RightAlignment)
}

func TestTablePrinter_WithWriter(t *testing.T) {
	p := pterm.TablePrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
}
