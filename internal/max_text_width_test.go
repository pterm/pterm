package internal

import (
	"fmt"
	"testing"
)

func TestStripOSC8Hyperlinks(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Basic OSC 8 hyperlink with BEL terminator",
			input:    "\x1b]8;;https://example.com\x07link text\x1b]8;;\x07",
			expected: "link text",
		},
		{
			name:     "Basic OSC 8 hyperlink with ST terminator",
			input:    "\x1b]8;;https://example.com\x1b\\link text\x1b]8;;\x1b\\",
			expected: "link text",
		},
		{
			name:     "OSC 8 hyperlink created with fmt.Sprintf",
			input:    fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", "https://example.com", "link text"),
			expected: "link text",
		},
		{
			name:     "Multiple OSC 8 hyperlinks",
			input:    "\x1b]8;;https://example.com\x1b\\first\x1b]8;;\x1b\\ and \x1b]8;;https://example.org\x1b\\second\x1b]8;;\x1b\\",
			expected: "first and second",
		},
		{
			name:     "No hyperlinks",
			input:    "plain text",
			expected: "plain text",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Hyperlink with special characters in URL",
			input:    "\x1b]8;;https://example.com/path?query=value&foo=bar\x1b\\link\x1b]8;;\x1b\\",
			expected: "link",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := StripOSC8Hyperlinks(tt.input)
			if result != tt.expected {
				t.Errorf("StripOSC8Hyperlinks() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGetStringMaxWidth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Plain text",
			input:    "hello",
			expected: 5,
		},
		{
			name:     "Text with OSC 8 hyperlink",
			input:    fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", "https://example.com", "link"),
			expected: 4, // Only "link" should be counted
		},
		{
			name:     "Text with OSC 8 hyperlink using BEL",
			input:    "\x1b]8;;https://example.com\x07link text\x1b]8;;\x07",
			expected: 9, // Only "link text" should be counted
		},
		{
			name:     "Multiple lines with hyperlinks",
			input:    "plain\n\x1b]8;;https://example.com\x1b\\link\x1b]8;;\x1b\\",
			expected: 5, // "plain" is longer than "link"
		},
		{
			name:     "Hyperlink longer than plain text",
			input:    "hi\n\x1b]8;;https://example.com\x1b\\longer link\x1b]8;;\x1b\\",
			expected: 11, // "longer link" is the longest
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Hyperlink with ANSI color codes",
			input:    "\x1b]8;;https://example.com\x1b\\\x1b[31mred link\x1b[0m\x1b]8;;\x1b\\",
			expected: 8, // "red link" without color codes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStringMaxWidth(tt.input)
			if result != tt.expected {
				t.Errorf("GetStringMaxWidth() = %d, want %d", result, tt.expected)
			}
		})
	}
}
