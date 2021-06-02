package pterm

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/pterm/pterm/internal"
)

// DefaultTable contains standards, which can be used to print a TablePrinter.
var DefaultTable = TablePrinter{
	Style:          &ThemeDefault.TableStyle,
	HeaderStyle:    &ThemeDefault.TableHeaderStyle,
	Separator:      " | ",
	SeparatorStyle: &ThemeDefault.TableSeparatorStyle,
}

// TableData is the type that contains the data of a TablePrinter.
type TableData [][]string

// TablePrinter is able to render tables.
type TablePrinter struct {
	Style          *Style
	HasHeader      bool
	HeaderStyle    *Style
	Separator      string
	SeparatorStyle *Style
	Data           TableData
}

// WithStyle returns a new TablePrinter with a specific Style.
func (p TablePrinter) WithStyle(style *Style) *TablePrinter {
	p.Style = style
	return &p
}

// WithHasHeader returns a new TablePrinter, where the first line is marked as a header.
func (p TablePrinter) WithHasHeader(b ...bool) *TablePrinter {
	p.HasHeader = internal.WithBoolean(b)
	return &p
}

// WithHeaderStyle returns a new TablePrinter with a specific HeaderStyle.
func (p TablePrinter) WithHeaderStyle(style *Style) *TablePrinter {
	p.HeaderStyle = style
	return &p
}

// WithSeparator returns a new TablePrinter with a specific separator.
func (p TablePrinter) WithSeparator(separator string) *TablePrinter {
	p.Separator = separator
	return &p
}

// WithSeparatorStyle returns a new TablePrinter with a specific SeparatorStyle.
func (p TablePrinter) WithSeparatorStyle(style *Style) *TablePrinter {
	p.SeparatorStyle = style
	return &p
}

// WithData returns a new TablePrinter with specific Data.
func (p TablePrinter) WithData(data [][]string) *TablePrinter {
	p.Data = data
	return &p
}

// WithCSVReader return a new TablePrinter with specified Data extracted from CSV.
func (p TablePrinter) WithCSVReader(reader *csv.Reader) *TablePrinter {
	if records, err := reader.ReadAll(); err == nil {
		p.Data = records
	}
	return &p
}

// WithSliceOfStruct returns a table printer with data extracted from the slice.
// if the input isn't a slice of structs no data will be extracted.
func (p TablePrinter) WithSliceOfStruct(slice interface{}) *TablePrinter {

	to := reflect.TypeOf(slice)
	if to.Kind() != reflect.Slice {
		return &p
	}
	el := to.Elem()

	isPointer := false
	if el.Kind() == reflect.Ptr {
		el = el.Elem()
		isPointer = true
	}

	if el.Kind() != reflect.Struct {
		return &p
	}

	numFields := el.NumField()
	fieldNames := make([]string, numFields)

	for i := 0; i < numFields; i++ {
		fieldNames[i] = el.Field(i).Name
	}

	records := TableData{
		fieldNames,
	}

	obj := reflect.ValueOf(slice)

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
	p.Data = records
	return &p
}

// Srender renders the TablePrinter as a string.
func (p TablePrinter) Srender() (string, error) {
	if p.Style == nil {
		p.Style = NewStyle()
	}
	if p.SeparatorStyle == nil {
		p.SeparatorStyle = NewStyle()
	}
	if p.HeaderStyle == nil {
		p.HeaderStyle = NewStyle()
	}

	var ret string
	maxColumnWidth := make(map[int]int)

	for _, row := range p.Data {
		for ci, column := range row {
			columnLength := utf8.RuneCountInString(RemoveColorFromString(column))
			if columnLength > maxColumnWidth[ci] {
				maxColumnWidth[ci] = columnLength
			}
		}
	}

	for ri, row := range p.Data {
		for ci, column := range row {
			columnLength := utf8.RuneCountInString(RemoveColorFromString(column))
			columnString := column + strings.Repeat(" ", maxColumnWidth[ci]-columnLength)

			if ci != len(row) && ci != 0 {
				ret += p.Style.Sprint(p.SeparatorStyle.Sprint(p.Separator))
			}

			if p.HasHeader && ri == 0 {
				ret += p.Style.Sprint(p.HeaderStyle.Sprint(columnString))
			} else {
				ret += p.Style.Sprint(columnString)
			}
		}

		ret += "\n"
	}

	ret = strings.TrimSuffix(ret, "\n")

	return ret, nil
}

// Render prints the TablePrinter to the terminal.
func (p TablePrinter) Render() error {
	s, _ := p.Srender()
	Println(s)

	return nil
}
