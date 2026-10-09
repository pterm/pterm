package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestWithBoolean(t *testing.T) {
	tests := []struct {
		name     string
		input    []bool
		expected bool
	}{
		{name: "No value defaults to true", input: []bool{}, expected: true},
		{name: "Nil slice defaults to true", input: nil, expected: true},
		{name: "Explicit true", input: []bool{true}, expected: true},
		{name: "Explicit false", input: []bool{false}, expected: false},
		{name: "Only the first value counts (false wins)", input: []bool{false, true}, expected: false},
		{name: "Only the first value counts (true wins)", input: []bool{true, false}, expected: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.WithBoolean(tt.input))
		})
	}
}
