package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestTableDataFromCSV(t *testing.T) {
	expected := pterm.TableData{
		[]string{"firstname", "lastname", "username"},
		[]string{"Marvin", "Wendt", "MarvinJWendt"},
	}

	input := "firstname,lastname,username\nMarvin,Wendt,MarvinJWendt"

	assert.EqualValues(t, expected, TableDataFromCSV(input))
}
