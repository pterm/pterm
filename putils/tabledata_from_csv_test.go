package putils

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestTableDataFromCSV(t *testing.T) {
	expected := pterm.TableData{
		[]string{"firstname", "lastname", "username"},
		[]string{"Marvin", "Wendt", "MarvinJWendt"},
	}

	input := "firstname,lastname,username\nMarvin,Wendt,MarvinJWendt"

	testza.AssertEqualValues(t, expected, TableDataFromCSV(input))
}
