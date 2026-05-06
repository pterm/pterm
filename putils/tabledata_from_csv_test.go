package putils

import (
	"testing"

	"github.com/pterm/pterm/internal/testhelper"
	"github.com/pterm/pterm"
)

func TestTableDataFromCSV(t *testing.T) {
	expected := pterm.TableData{
		[]string{"firstname", "lastname", "username"},
		[]string{"Marvin", "Wendt", "MarvinJWendt"},
	}

	input := "firstname,lastname,username\nMarvin,Wendt,MarvinJWendt"

	testhelper.AssertEqualValues(t, expected, TableDataFromCSV(input))
}
