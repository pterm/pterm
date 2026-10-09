package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenterText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected string
	}{
		{
			name:     "Empty string",
			text:     "",
			expected: "",
		},
		{
			name:     "Single line is unchanged",
			text:     "Hello World",
			expected: "Hello World",
		},
		{
			name:     "Shorter lines are centered on the longest line",
			text:     "Hello World\n!!!",
			expected: "Hello World\n    !!!    ",
		},
		{
			name:     "Odd padding drops the spare space",
			text:     "a\nbcd",
			expected: " a \nbcd",
		},
		{
			name:     "Three lines",
			text:     "a\nabc\nabcde",
			expected: "  a  \n abc \nabcde",
		},
		{
			name:     "Escape codes do not count toward the width",
			text:     "\x1b[31mab\x1b[0m\nabcdef",
			expected: "  \x1b[31mab\x1b[0m  \nabcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, CenterText(tt.text))
		})
	}
}
