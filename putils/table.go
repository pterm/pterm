package putils

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"

	"github.com/pterm/pterm"
)

// TableFromStructs turns a slice of structs into table data. The first row
// is the header, made of the field names. A `pterm:"Name"` tag renames a
// column and `pterm:"-"` hides it. Unexported fields are skipped, nil
// pointers become empty cells, and other values are formatted with
// fmt.Sprint.
//
//	pterm.Table.WithData(putils.TableFromStructs(users)).Render()
func TableFromStructs[T any](rows []T) pterm.TableData {
	typ := reflect.TypeFor[T]()
	for typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil
	}

	fields := structFields(typ)

	header := make([]string, len(fields))
	for i, f := range fields {
		header[i] = f.name
	}

	data := pterm.TableData{header}

	for _, row := range rows {
		v := reflect.ValueOf(row)
		for v.Kind() == reflect.Pointer && !v.IsNil() {
			v = v.Elem()
		}

		cells := make([]string, len(fields))

		if v.Kind() == reflect.Struct {
			for i, f := range fields {
				cells[i] = format(v.Field(f.index))
			}
		}

		data = append(data, cells)
	}

	return data
}

// TableFromCSV parses CSV text into table data. Rows may have different
// lengths; the table pads them.
func TableFromCSV(s string) (pterm.TableData, error) {
	return TableFromSeparated(s, ',')
}

// TableFromSeparated parses text whose values are separated by sep, such
// as '\t' for TSV, into table data. Quoting follows the CSV rules.
func TableFromSeparated(s string, sep rune) (pterm.TableData, error) {
	r := csv.NewReader(strings.NewReader(s))
	r.Comma = sep
	r.FieldsPerRecord = -1
	r.LazyQuotes = true

	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("putils: parse table: %w", err)
	}

	return records, nil
}

type field struct {
	name  string
	index int
}

// structFields returns the exported fields of typ that are not hidden with
// a `pterm:"-"` tag, in declaration order.
func structFields(typ reflect.Type) []field {
	var fields []field

	for i := range typ.NumField() {
		f := typ.Field(i)
		if !f.IsExported() {
			continue
		}

		name := f.Name

		switch tag := f.Tag.Get("pterm"); tag {
		case "-":
			continue
		case "":
		default:
			name = tag
		}

		fields = append(fields, field{name: name, index: i})
	}

	return fields
}

// format returns the text of a value: empty for nil pointers and
// interfaces, fmt.Sprint for everything else.
func format(v reflect.Value) string {
	for v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface {
		if v.IsNil() {
			return ""
		}

		v = v.Elem()
	}

	if !v.CanInterface() {
		return ""
	}

	return fmt.Sprint(v.Interface())
}
