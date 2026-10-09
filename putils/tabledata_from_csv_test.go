package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestTableDataFromCSV(t *testing.T) {
	tests := []struct {
		name     string
		csv      string
		expected pterm.TableData
	}{
		{
			name: "Header and one row",
			csv:  "firstname,lastname,username\nMarvin,Wendt,MarvinJWendt",
			expected: pterm.TableData{
				{"firstname", "lastname", "username"},
				{"Marvin", "Wendt", "MarvinJWendt"},
			},
		},
		{
			name: "Single row",
			csv:  "a,b,c",
			expected: pterm.TableData{
				{"a", "b", "c"},
			},
		},
		{
			name: "Uneven rows are preserved",
			csv:  "a,b,c\nd,e",
			expected: pterm.TableData{
				{"a", "b", "c"},
				{"d", "e"},
			},
		},
		{
			name: "Empty fields are preserved",
			csv:  "a,,c\n,,",
			expected: pterm.TableData{
				{"a", "", "c"},
				{"", "", ""},
			},
		},
		{
			name: "Empty input yields one empty cell",
			csv:  "",
			expected: pterm.TableData{
				{""},
			},
		},
		{
			// Rows are split at plain LF only, so a CR of a CRLF pair stays in
			// the last field of its row.
			name: "CRLF line endings keep the carriage return",
			csv:  "a,b\r\nc,d",
			expected: pterm.TableData{
				{"a", "b\r"},
				{"c", "d"},
			},
		},
		{
			// The conversion is a plain split: RFC 4180 quoting is not
			// supported and quotes are kept literally.
			name: "Quoted fields are not unquoted",
			csv:  `"hello, world",b`,
			expected: pterm.TableData{
				{`"hello`, ` world"`, "b"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, TableDataFromCSV(tt.csv))
		})
	}
}
