package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestTableDataFromSeparatedValues(t *testing.T) {
	tests := []struct {
		name           string
		text           string
		valueSeparator string
		rowSeparator   string
		expected       pterm.TableData
	}{
		{
			name:           "Semicolon separated values",
			text:           "firstname;lastname;username\nMarvin;Wendt;MarvinJWendt",
			valueSeparator: ";",
			rowSeparator:   "\n",
			expected: pterm.TableData{
				{"firstname", "lastname", "username"},
				{"Marvin", "Wendt", "MarvinJWendt"},
			},
		},
		{
			name:           "Multi character separators",
			text:           "a | b || c | d",
			valueSeparator: " | ",
			rowSeparator:   " || ",
			expected: pterm.TableData{
				{"a", "b"},
				{"c", "d"},
			},
		},
		{
			name:           "CRLF row separator",
			text:           "a,b\r\nc,d",
			valueSeparator: ",",
			rowSeparator:   "\r\n",
			expected: pterm.TableData{
				{"a", "b"},
				{"c", "d"},
			},
		},
		{
			name:           "Uneven rows are preserved",
			text:           "a;b;c\nd\ne;f",
			valueSeparator: ";",
			rowSeparator:   "\n",
			expected: pterm.TableData{
				{"a", "b", "c"},
				{"d"},
				{"e", "f"},
			},
		},
		{
			name:           "Empty fields are preserved",
			text:           "a;;c",
			valueSeparator: ";",
			rowSeparator:   "\n",
			expected: pterm.TableData{
				{"a", "", "c"},
			},
		},
		{
			name:           "Empty input yields one empty cell",
			text:           "",
			valueSeparator: ";",
			rowSeparator:   "\n",
			expected: pterm.TableData{
				{""},
			},
		},
		{
			name:           "Trailing row separator yields a trailing empty row",
			text:           "a;b\n",
			valueSeparator: ";",
			rowSeparator:   "\n",
			expected: pterm.TableData{
				{"a", "b"},
				{""},
			},
		},
		{
			// The conversion is a plain split: there is no quoting mechanism
			// that lets a field contain the separator.
			name:           "Quotes are kept literally",
			text:           `"a;b";c`,
			valueSeparator: ";",
			rowSeparator:   "\n",
			expected: pterm.TableData{
				{`"a`, `b"`, "c"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, TableDataFromSeparatedValues(tt.text, tt.valueSeparator, tt.rowSeparator))
		})
	}
}
