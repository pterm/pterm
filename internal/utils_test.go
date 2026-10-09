package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestRunsInCi(t *testing.T) {
	tests := []struct {
		name     string
		ci       string
		expected bool
	}{
		{name: "Unset CI means no CI", ci: "", expected: false},
		{name: "CI=true", ci: "true", expected: true},
		{name: "CI=1", ci: "1", expected: true},
		// Only presence is checked, the value is not interpreted.
		{name: "CI=false still counts as CI", ci: "false", expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CI", tt.ci)

			assert.Equal(t, tt.expected, internal.RunsInCi())
		})
	}
}

func TestRemoveEscapeCodes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Plain text is unchanged",
			input:    "plain text",
			expected: "plain text",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Removes SGR color codes",
			input:    "\x1b[31mred\x1b[0m",
			expected: "red",
		},
		{
			name:     "Removes truecolor codes",
			input:    "\x1b[38;2;1;2;3mrgb\x1b[0m",
			expected: "rgb",
		},
		{
			name:     "Removes OSC 8 hyperlinks but keeps the link text",
			input:    "\x1b]8;;https://example.com\x1b\\link\x1b]8;;\x1b\\",
			expected: "link",
		},
		{
			name:     "Removes hyperlink and color codes together",
			input:    "\x1b]8;;https://example.com\x1b\\\x1b[31mred link\x1b[0m\x1b]8;;\x1b\\ end",
			expected: "red link end",
		},
		{
			name:     "Keeps non-SGR sequences like cursor movement",
			input:    "\x1b[2A\x1b[31mup\x1b[0m",
			expected: "\x1b[2Aup",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.RemoveEscapeCodes(tt.input))
		})
	}
}
