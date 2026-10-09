package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestTableDataFromTSV(t *testing.T) {
	tests := []struct {
		name     string
		tsv      string
		expected pterm.TableData
	}{
		{
			name: "Header and one row",
			tsv:  "firstname\tlastname\tusername\nMarvin\tWendt\tMarvinJWendt",
			expected: pterm.TableData{
				{"firstname", "lastname", "username"},
				{"Marvin", "Wendt", "MarvinJWendt"},
			},
		},
		{
			name: "Fields may contain spaces and commas",
			tsv:  "full name\tlocation\nMarvin Wendt\tHamburg, Germany",
			expected: pterm.TableData{
				{"full name", "location"},
				{"Marvin Wendt", "Hamburg, Germany"},
			},
		},
		{
			name: "Empty fields are preserved",
			tsv:  "a\t\tc",
			expected: pterm.TableData{
				{"a", "", "c"},
			},
		},
		{
			name: "Uneven rows are preserved",
			tsv:  "a\tb\tc\nd",
			expected: pterm.TableData{
				{"a", "b", "c"},
				{"d"},
			},
		},
		{
			name: "Empty input yields one empty cell",
			tsv:  "",
			expected: pterm.TableData{
				{""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, TableDataFromTSV(tt.tsv))
		})
	}
}
