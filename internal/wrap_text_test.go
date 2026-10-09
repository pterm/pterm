package internal

import (
	"strings"
	"testing"
)

func TestWrapText(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		width    int
		expected []string
	}{
		{
			name:     "Fits on one line",
			input:    "hello world",
			width:    20,
			expected: []string{"hello world"},
		},
		{
			name:     "Wraps at spaces",
			input:    "aaa bbb ccc ddd",
			width:    7,
			expected: []string{"aaa bbb", "ccc ddd"},
		},
		{
			name:     "Preserves explicit newlines",
			input:    "first\nsecond line",
			width:    20,
			expected: []string{"first", "second line"},
		},
		{
			name:     "Wraps each paragraph separately",
			input:    "aaa bbb\nccc ddd eee",
			width:    7,
			expected: []string{"aaa bbb", "ccc ddd", "eee"},
		},
		{
			name:     "Breaks a word wider than the line",
			input:    "abcdefghij",
			width:    4,
			expected: []string{"abcd", "efgh", "ij"},
		},
		{
			name:     "Breaks a long word inside a sentence",
			input:    "see abcdefghij end",
			width:    4,
			expected: []string{"see", "abcd", "efgh", "ij", "end"},
		},
		{
			name:     "Zero width disables wrapping",
			input:    "aaa bbb\nccc",
			width:    0,
			expected: []string{"aaa bbb", "ccc"},
		},
		{
			name:     "Negative width disables wrapping",
			input:    "aaa bbb\nccc",
			width:    -5,
			expected: []string{"aaa bbb", "ccc"},
		},
		{
			name:     "Word exactly as wide as the line",
			input:    "abcd efgh",
			width:    4,
			expected: []string{"abcd", "efgh"},
		},
		{
			name:     "Trailing spaces on wrapped lines are trimmed",
			input:    "a  b",
			width:    3,
			expected: []string{"a", "b"},
		},
		{
			name:     "Mixed narrow and wide runes",
			input:    "Go 言語 rocks",
			width:    5,
			expected: []string{"Go", "言語", "rocks"},
		},
		{
			name:     "Ignores escape codes when measuring",
			input:    "\x1b[31mred\x1b[0m and blue",
			width:    8,
			expected: []string{"\x1b[31mred\x1b[0m and", "blue"},
		},
		{
			name:     "Breaks fully styled words and restyles every chunk",
			input:    "\x1b[31mabcdefgh\x1b[0m",
			width:    4,
			expected: []string{"\x1b[31mabcd\x1b[0m", "\x1b[31mefgh\x1b[0m"},
		},
		{
			name:     "Keeps words with inner escape codes unbroken",
			input:    "abc\x1b[31mdefghijklm\x1b[0mnop",
			width:    4,
			expected: []string{"abc\x1b[31mdefghijklm\x1b[0mnop"},
		},
		{
			name:     "Handles wide runes",
			input:    "日本語のテキスト",
			width:    6,
			expected: []string{"日本語", "のテキ", "スト"},
		},
		{
			name:     "Normalizes CRLF",
			input:    "first\r\nsecond",
			width:    20,
			expected: []string{"first", "second"},
		},
		{
			name:     "Empty string",
			input:    "",
			width:    10,
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapText(tt.input, tt.width)
			if len(result) != len(tt.expected) {
				t.Fatalf("WrapText() = %q (%d lines), want %q (%d lines)", result, len(result), tt.expected, len(tt.expected))
			}

			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("WrapText() line %d = %q, want %q", i, result[i], tt.expected[i])
				}
			}

			if tt.width > 0 {
				for i, line := range result {
					if !strings.Contains(line, "\x1b") && GetStringMaxWidth(line) > tt.width {
						t.Errorf("WrapText() line %d is wider than %d: %q", i, tt.width, line)
					}
				}
			}
		})
	}
}
