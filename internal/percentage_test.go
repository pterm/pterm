package internal_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestPercentage(t *testing.T) {
	tests := []struct {
		name           string
		total, current float64
		expected       float64
	}{
		{name: "Half of hundred", total: 100, current: 50, expected: 50},
		{name: "Half of two hundred", total: 200, current: 100, expected: 50},
		{name: "Quarter of five hundred", total: 500, current: 100, expected: 20},
		{name: "Zero current", total: 100, current: 0, expected: 0},
		{name: "Full total", total: 100, current: 100, expected: 100},
		{name: "Above the total", total: 100, current: 150, expected: 150},
		{name: "Negative current", total: 100, current: -50, expected: -50},
		{name: "Negative total", total: -100, current: 50, expected: -50},
		{name: "Fractional result", total: 8, current: 1, expected: 12.5},
		{name: "Fractional inputs", total: 0.5, current: 0.25, expected: 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.Percentage(tt.total, tt.current))
		})
	}
}

func TestPercentageRepeatingFraction(t *testing.T) {
	assert.InDelta(t, 33.333333, internal.Percentage(3, 1), 0.000001)
	assert.InDelta(t, 66.666666, internal.Percentage(3, 2), 0.000001)
}

func TestPercentageZeroTotal(t *testing.T) {
	assert.True(t, math.IsInf(internal.Percentage(0, 50), 1))
	assert.True(t, math.IsInf(internal.Percentage(0, -50), -1))
	assert.True(t, math.IsNaN(internal.Percentage(0, 0)))
}

func TestPercentageRound(t *testing.T) {
	tests := []struct {
		name           string
		total, current float64
		expected       float64
	}{
		{name: "Exact value stays untouched", total: 200, current: 100, expected: 50},
		{name: "Rounds down", total: 3, current: 1, expected: 33},                           // 33.33...
		{name: "Rounds up", total: 3, current: 2, expected: 67},                             // 66.66...
		{name: "Rounds half away from zero", total: 8, current: 1, expected: 13},            // 12.5
		{name: "Rounds negative half away from zero", total: 8, current: -1, expected: -13}, // -12.5
		{name: "Zero current", total: 100, current: 0, expected: 0},
		{name: "Full total", total: 100, current: 100, expected: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, internal.PercentageRound(tt.total, tt.current))
		})
	}
}
