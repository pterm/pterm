package pterm_test

import (
	"encoding/csv"
	"fmt"
	"strings"
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

func TestTablePrinter(t *testing.T) {
	data := []pterm.TableData{
		{
			{"Firstname", "Lastname", "Age"},
			{"Paul", "Dean", "18"},
			{"John", "Carter", "21"},
			{"Jeremy", "Johnson", "42"},
		},
		{
			{"     Firstname  ", "   Lastname  ", "  Age "},
			{"  Paul ", "Dean", "18"},
			{"  John   ", "Carter", " 21      "},
			{"  Jeremy   ", "Johnson", "    42"},
		},
		{
			{"     Firstname  ", "   Lastname  ", "  Age "},
			{"  P\naul ", "Dean\n", "1\n8"},
			{"  John   ", "Car\n\nter", " 21      "},
			{"  J\n\n\neremy   ", "Johnson", "    4\n2"},
		},
	}

	for i, d := range data {
		t.Run(fmt.Sprintf("Render_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).Render()
			})
		})

		t.Run(fmt.Sprintf("WithStyle_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithStyle(pterm.NewStyle(pterm.FgRed)).Render()
			})
		})

		t.Run(fmt.Sprintf("WithHasHeader_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithHasHeader().Render()
			})
		})

		t.Run(fmt.Sprintf("WithHeaderStyle_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithHasHeader().WithHeaderStyle(pterm.NewStyle(pterm.FgRed)).Render()
			})
		})

		t.Run(fmt.Sprintf("WithHeaderRowSeparator_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithHasHeader().WithHeaderRowSeparator("-").Render()
			})
		})

		t.Run(fmt.Sprintf("WithHeaderRowSeparatorStyle_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithHasHeader().WithHeaderRowSeparator("-").WithHeaderRowSeparatorStyle(pterm.NewStyle(pterm.FgRed)).Render()
			})
		})

		t.Run(fmt.Sprintf("WithSeparator_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithSeparator(":").Render()
			})
		})

		t.Run(fmt.Sprintf("WithSeparatorStyle_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithSeparatorStyle(pterm.NewStyle(pterm.FgRed)).Render()
			})
		})

		t.Run(fmt.Sprintf("WithRowSeparator_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithRowSeparator("-").Render()
			})
		})

		t.Run(fmt.Sprintf("WithRowSeparatorStyle_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithRowSeparatorStyle(pterm.NewStyle(pterm.FgRed)).Render()
			})
		})

		t.Run(fmt.Sprintf("WithCSVReader_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithCSVReader(csv.NewReader(strings.NewReader("a,b,c\n1,2,3\nx,y,z"))).Render()
			})
		})

		t.Run(fmt.Sprintf("WithBoxed_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithBoxed().Render()
			})
		})

		t.Run(fmt.Sprintf("WithLeftAlignment_%d", i), func(t *testing.T) {
			printerTest(t, func() {
				pterm.DefaultTable.WithData(d).WithLeftAlignment().Render()
			})
		})
	}
}
