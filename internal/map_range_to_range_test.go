package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestMapRangeToRange(t *testing.T) {
	tests := []struct {
		name                                    string
		fromMin, fromMax, toMin, toMax, current float32
		expected                                int
	}{
		{name: "Middle of range", fromMin: 0, fromMax: 100, toMin: 0, toMax: 255, current: 50, expected: 127},
		{name: "Lower boundary", fromMin: 0, fromMax: 100, toMin: 0, toMax: 255, current: 0, expected: 0},
		{name: "Upper boundary", fromMin: 0, fromMax: 100, toMin: 0, toMax: 255, current: 100, expected: 255},
		{name: "Result is truncated not rounded", fromMin: 0, fromMax: 100, toMin: 0, toMax: 255, current: 1, expected: 2}, // 2.55 -> 2
		{name: "Wider from range", fromMin: 0, fromMax: 400, toMin: 0, toMax: 255, current: 200, expected: 127},
		{name: "Negative from range", fromMin: -200, fromMax: 200, toMin: 0, toMax: 255, current: 0, expected: 127},
		{name: "Negative to range lower boundary", fromMin: -100, fromMax: 100, toMin: -50, toMax: 50, current: -100, expected: -50},
		{name: "Negative to range middle", fromMin: -100, fromMax: 100, toMin: -50, toMax: 50, current: 0, expected: 0},
		{name: "Negative to range upper boundary", fromMin: -100, fromMax: 100, toMin: -50, toMax: 50, current: 100, expected: 50},
		{name: "Inverted to range lower boundary", fromMin: 0, fromMax: 100, toMin: 255, toMax: 0, current: 0, expected: 255},
		{name: "Inverted to range middle", fromMin: 0, fromMax: 100, toMin: 255, toMax: 0, current: 50, expected: 127},
		{name: "Inverted to range upper boundary", fromMin: 0, fromMax: 100, toMin: 255, toMax: 0, current: 100, expected: 0},
		{name: "Inverted from range", fromMin: 100, fromMax: 0, toMin: 0, toMax: 255, current: 25, expected: 191},
		{name: "Fractional ranges", fromMin: 0, fromMax: 200.123, toMin: 0, toMax: 254.3, current: 100, expected: 127},
		{name: "Empty from range returns zero", fromMin: 5, fromMax: 5, toMin: 0, toMax: 255, current: 5, expected: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.MapRangeToRange(tt.fromMin, tt.fromMax, tt.toMin, tt.toMax, tt.current))
		})
	}
}
