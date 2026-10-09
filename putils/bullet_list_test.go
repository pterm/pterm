package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestBulletListItemFromString(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		padding  string
		expected pterm.BulletListItem
	}{
		{
			name:     "No padding",
			text:     "item",
			padding:  " ",
			expected: pterm.BulletListItem{Level: 0, Text: "item"},
		},
		{
			name:     "Two spaces",
			text:     "  item",
			padding:  " ",
			expected: pterm.BulletListItem{Level: 2, Text: "item"},
		},
		{
			name:     "Tab padding",
			text:     "\t\t\titem",
			padding:  "\t",
			expected: pterm.BulletListItem{Level: 3, Text: "item"},
		},
		{
			name:     "Inner padding is kept in the text",
			text:     " item one",
			padding:  " ",
			expected: pterm.BulletListItem{Level: 1, Text: "item one"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, BulletListItemFromString(tt.text, tt.padding))
		})
	}
}

func TestBulletListFromStrings(t *testing.T) {
	list := BulletListFromStrings([]string{"a", " b", "  c"}, " ")

	expected := []pterm.BulletListItem{
		{Level: 0, Text: "a"},
		{Level: 1, Text: "b"},
		{Level: 2, Text: "c"},
	}
	assert.Equal(t, expected, list.Items)
}

func TestBulletListFromString(t *testing.T) {
	list := BulletListFromString("a\n b\n  c", " ")

	expected := []pterm.BulletListItem{
		{Level: 0, Text: "a"},
		{Level: 1, Text: "b"},
		{Level: 2, Text: "c"},
	}
	assert.Equal(t, expected, list.Items)
}

func TestBulletListFromStringEmptyInput(t *testing.T) {
	list := BulletListFromString("", " ")

	expected := []pterm.BulletListItem{
		{Level: 0, Text: ""},
	}
	assert.Equal(t, expected, list.Items)
}
