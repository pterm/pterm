package putils

import (
	"fmt"
	"reflect"

	"github.com/pterm/pterm"
)

// TableFromStructSlice Accepts a customized table printer and adds data to it, then returns it.
func TableFromStructSlice(tablePrinter pterm.TablePrinter, structSlice interface{}) *pterm.TablePrinter {
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

	items := make([]interface{}, obj.Len())
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
			record[i] = fmt.Sprintf("%v", fieldVal)
		}
		records = append(records, record)
	}
	tablePrinter.Data = records

	return &tablePrinter
}

// DefaultTableFromStructSlice feeds the default table printer into the first param and returns the table filled with your computed data.
func DefaultTableFromStructSlice(structSlice interface{}) *pterm.TablePrinter {
	return TableFromStructSlice(pterm.DefaultTable, structSlice)
}
