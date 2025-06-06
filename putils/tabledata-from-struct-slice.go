package putils

import (
	"reflect"

	"github.com/pterm/pterm"
)

// TableFromStructSlice accepts a customized table printer and and a slice of a struct.
// The table will be populated with the values of the structs. The header will be set to the structs field name.
// Use .WithHasHeader() to color the header.
// The function will return the populated pterm.TablePrinter.
func TableFromStructSlice(tablePrinter pterm.TablePrinter, structSlice any) *pterm.TablePrinter {
	to := reflect.TypeOf(structSlice)
	if to.Kind() != reflect.Slice {
		return &tablePrinter
	}
	el := to.Elem()

	isPointer := false
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
		isPointer = true
	}

	if el.Kind() != reflect.Struct {
		return &tablePrinter
	}

	numFields := el.NumField()
	fieldNames := make([]string, numFields)

	for i := 0; i < numFields; i++ {
		fieldNames[i] = el.Field(i).Name
	}

	records := pterm.TableData{
		fieldNames,
	}

	obj := reflect.ValueOf(structSlice)

	items := make([]any, obj.Len())
	for i := 0; i < obj.Len(); i++ {
		if isPointer {
			items[i] = obj.Index(i).Elem().Interface()
		} else {
			items[i] = obj.Index(i).Interface()
		}
	}

	for _, v := range items {
		item := reflect.ValueOf(v)
		record := make([]string, numFields)
		for i := 0; i < numFields; i++ {
			fieldVal := item.Field(i).Interface()
			record[i] = pterm.Sprintf("%v", fieldVal)
		}
		records = append(records, record)
	}
	tablePrinter.Data = records

	return &tablePrinter
}

// DefaultTableFromStructSlice will be populate the pterm.DefaultTable with the values of the structs. The header will be set to the structs field name.
// Use .WithHasHeader() to color the header.
// The function will return the populated pterm.TablePrinter.
func DefaultTableFromStructSlice(structSlice any) *pterm.TablePrinter {
	return TableFromStructSlice(pterm.DefaultTable, structSlice)
}
