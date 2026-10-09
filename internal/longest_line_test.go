package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestReturnLongestLine(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		sep      string
		expected string
	}{
		{
			name:     "Empty string",
			text:     "",
			sep:      "\n",
			expected: "",
		},
		{
			name:     "Single line",
			text:     "hello",
			sep:      "\n",
			expected: "hello",
		},
		{
			name:     "Longest of multiple lines",
			text:     "a\nlongest line\nbb",
			sep:      "\n",
			expected: "longest line",
		},
		{
			name:     "First line wins a tie",
			text:     "aaa\nbbb",
			sep:      "\n",
			expected: "aaa",
		},
		{
			name:     "Custom separator",
			text:     "a|bbb|cc",
			sep:      "|",
			expected: "bbb",
		},
		{
			name:     "Separator not found returns the whole text",
			text:     "a b c",
			sep:      "\n",
			expected: "a b c",
		},
		{
			name:     "Escape codes do not count toward the width",
			text:     "\x1b[31mab\x1b[0m\nabc",
			sep:      "\n",
			expected: "abc",
		},
		{
			name:     "Longest styled line is returned with its codes intact",
			text:     "\x1b[31mabcd\x1b[0m\nabc",
			sep:      "\n",
			expected: "\x1b[31mabcd\x1b[0m",
		},
		{
			name:     "OSC 8 hyperlinks only count their visible text",
			text:     "\x1b]8;;https://example.com\x1b\\ab\x1b]8;;\x1b\\\nabc",
			sep:      "\n",
			expected: "abc",
		},
		{
			name:     "Wide runes count double",
			text:     "日本語\nabcde",
			sep:      "\n",
			expected: "日本語", // visible width 6 beats 5
		},
		{
			name:     "Trailing separator adds an empty line",
			text:     "abc\n",
			sep:      "\n",
			expected: "abc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.ReturnLongestLine(tt.text, tt.sep))
		})
	}
}
