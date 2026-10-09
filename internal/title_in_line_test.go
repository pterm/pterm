package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestAddTitleToLine(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		line     string
		length   int
		left     bool
		expected string
	}{
		{
			name:     "Title on the left",
			title:    "Title",
			line:     "-",
			length:   20,
			left:     true,
			expected: "- Title ------------",
		},
		{
			name:     "Title on the right",
			title:    "Title",
			line:     "-",
			length:   20,
			left:     false,
			expected: "------------ Title -",
		},
		{
			name:     "Minimal length has no filler",
			title:    "Title",
			line:     "-",
			length:   9,
			left:     true,
			expected: "- Title -",
		},
		{
			name:     "Styled title only counts its visible width",
			title:    "\x1b[31mT\x1b[0m",
			line:     "-",
			length:   10,
			left:     true,
			expected: "- \x1b[31mT\x1b[0m ------",
		},
		{
			name:     "Wide runes in the title count double",
			title:    "日本",
			line:     "-",
			length:   12,
			left:     true,
			expected: "- 日本 -----",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := internal.AddTitleToLine(tt.title, tt.line, tt.length, tt.left)

			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.length, internal.GetStringMaxWidth(result))
		})
	}
}

func TestAddTitleToLineCenter(t *testing.T) {
	tests := []struct {
		name     string
		title    string
		line     string
		length   int
		expected string
	}{
		{
			name:     "Even filler is split equally",
			title:    "Ti",
			line:     "-",
			length:   10,
			expected: "--- Ti ---",
		},
		{
			name:     "Odd filler puts the spare character on the right",
			title:    "Title",
			line:     "-",
			length:   20,
			expected: "------ Title -------",
		},
		{
			name:     "Minimal length has one line character per side",
			title:    "Title",
			line:     "-",
			length:   9,
			expected: "- Title -",
		},
		{
			name:     "Styled title only counts its visible width",
			title:    "\x1b[31mT\x1b[0m",
			line:     "-",
			length:   9,
			expected: "--- \x1b[31mT\x1b[0m ---",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := internal.AddTitleToLineCenter(tt.title, tt.line, tt.length)

			assert.Equal(t, tt.expected, result)
			assert.Equal(t, tt.length, internal.GetStringMaxWidth(result))
		})
	}
}
