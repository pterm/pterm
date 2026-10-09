package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestRemoveAndCountPrefix(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		subString     string
		expectedText  string
		expectedCount int
	}{
		{
			name:          "No prefix",
			input:         "item",
			subString:     " ",
			expectedText:  "item",
			expectedCount: 0,
		},
		{
			name:          "Single space prefix",
			input:         " item",
			subString:     " ",
			expectedText:  "item",
			expectedCount: 1,
		},
		{
			name:          "Multiple space prefixes",
			input:         "   item",
			subString:     " ",
			expectedText:  "item",
			expectedCount: 3,
		},
		{
			name:          "Tab prefix",
			input:         "\t\titem",
			subString:     "\t",
			expectedText:  "item",
			expectedCount: 2,
		},
		{
			name:          "Only trims leading occurrences",
			input:         "  item  ",
			subString:     " ",
			expectedText:  "item  ",
			expectedCount: 2,
		},
		{
			name:          "Input consisting only of the prefix",
			input:         "   ",
			subString:     " ",
			expectedText:  "",
			expectedCount: 3,
		},
		{
			name:          "Empty input",
			input:         "",
			subString:     " ",
			expectedText:  "",
			expectedCount: 0,
		},
		{
			// strings.TrimLeft treats subString as a cutset of runes, so any
			// combination of its runes is trimmed, not just whole repetitions,
			// and the count is per trimmed byte.
			name:          "Multi character prefix acts as a rune cutset",
			input:         "baab-item",
			subString:     "ab",
			expectedText:  "-item",
			expectedCount: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			text, count := internal.RemoveAndCountPrefix(tt.input, tt.subString)

			assert.Equal(t, tt.expectedText, text)
			assert.Equal(t, tt.expectedCount, count)
		})
	}
}
