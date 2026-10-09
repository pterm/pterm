package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestCenterText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		width    int
		expected string
	}{
		{
			name:     "Empty string",
			text:     "",
			width:    0,
			expected: "",
		},
		{
			name:     "Even padding is split equally",
			text:     "abc",
			width:    7,
			expected: "  abc  ",
		},
		{
			name:     "Odd padding drops the spare space",
			text:     "abc",
			width:    6,
			expected: " abc ", // padding of 3 is halved to 1 on each side
		},
		{
			name:     "Text as wide as the width is unchanged",
			text:     "abc",
			width:    3,
			expected: "abc",
		},
		{
			name:     "Multiline centers every line",
			text:     "Hello World\n!!!",
			width:    15,
			expected: "  Hello World  \n      !!!      ",
		},
		{
			name:     "Width zero uses the longest line",
			text:     "Hello\n!",
			width:    0,
			expected: "Hello\n  !  ",
		},
		{
			name:     "Text wider than width is broken into chunks",
			text:     "Hello World\n!!!",
			width:    5,
			expected: "Hello\n Worl\n  d  \n !!! ",
		},
		{
			name:     "Escape codes do not count toward the width",
			text:     "\x1b[31mab\x1b[0m",
			width:    6,
			expected: "  \x1b[31mab\x1b[0m  ",
		},
		{
			name:     "Styled and plain lines align",
			text:     "\x1b[31mab\x1b[0m\nabcdef",
			width:    6,
			expected: "  \x1b[31mab\x1b[0m  \nabcdef",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.CenterText(tt.text, tt.width))
		})
	}
}
