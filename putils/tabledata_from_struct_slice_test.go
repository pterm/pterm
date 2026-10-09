package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

type person struct {
	Firstname string
	Lastname  string
	Age       int
	Admin     bool
}

func TestTableFromStructSlice(t *testing.T) {
	printer := TableFromStructSlice(pterm.DefaultTable, []person{
		{Firstname: "Marvin", Lastname: "Wendt", Age: 20, Admin: true},
		{Firstname: "John", Lastname: "Doe", Age: 42, Admin: false},
	})

	expected := pterm.TableData{
		{"Firstname", "Lastname", "Age", "Admin"},
		{"Marvin", "Wendt", "20", "true"},
		{"John", "Doe", "42", "false"},
	}
	assert.Equal(t, expected, printer.Data)
}

func TestTableFromStructSlicePointerSlice(t *testing.T) {
	printer := TableFromStructSlice(pterm.DefaultTable, []*person{
		{Firstname: "Marvin", Lastname: "Wendt", Age: 20, Admin: true},
	})

	expected := pterm.TableData{
		{"Firstname", "Lastname", "Age", "Admin"},
		{"Marvin", "Wendt", "20", "true"},
	}
	assert.Equal(t, expected, printer.Data)
}

func TestTableFromStructSliceEmptySliceYieldsHeaderOnly(t *testing.T) {
	printer := TableFromStructSlice(pterm.DefaultTable, []person{})

	expected := pterm.TableData{
		{"Firstname", "Lastname", "Age", "Admin"},
	}
	assert.Equal(t, expected, printer.Data)
}

func TestTableFromStructSliceNonSliceInputIsIgnored(t *testing.T) {
	printer := TableFromStructSlice(pterm.DefaultTable, 42)

	assert.Equal(t, pterm.DefaultTable.Data, printer.Data)
}

func TestTableFromStructSliceNonStructElementsAreIgnored(t *testing.T) {
	printer := TableFromStructSlice(pterm.DefaultTable, []int{1, 2, 3})

	assert.Equal(t, pterm.DefaultTable.Data, printer.Data)
}

func TestTableFromStructSliceDoesNotMutateInputPrinter(t *testing.T) {
	original := pterm.DefaultTable
	printer := TableFromStructSlice(original, []person{{Firstname: "Marvin"}})

	assert.NotSame(t, &original, printer)
	assert.Nil(t, original.Data)
	assert.NotNil(t, printer.Data)
}

func TestDefaultTableFromStructSlice(t *testing.T) {
	printer := DefaultTableFromStructSlice([]person{
		{Firstname: "Marvin", Lastname: "Wendt", Age: 20, Admin: true},
	})

	expected := pterm.TableData{
		{"Firstname", "Lastname", "Age", "Admin"},
		{"Marvin", "Wendt", "20", "true"},
	}
	assert.Equal(t, expected, printer.Data)
	assert.Nil(t, pterm.DefaultTable.Data, "the global default table must stay untouched")
}
