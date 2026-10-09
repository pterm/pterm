package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

func TestComplementary(t *testing.T) {
	tests := []struct {
		name                string
		r, g, b             uint8
		wantR, wantG, wantB uint8
	}{
		{name: "Black becomes white", r: 0, g: 0, b: 0, wantR: 255, wantG: 255, wantB: 255},
		{name: "White becomes black", r: 255, g: 255, b: 255, wantR: 0, wantG: 0, wantB: 0},
		{name: "Red becomes cyan", r: 255, g: 0, b: 0, wantR: 0, wantG: 255, wantB: 255},
		{name: "Green becomes magenta", r: 0, g: 255, b: 0, wantR: 255, wantG: 0, wantB: 255},
		{name: "Blue becomes yellow", r: 0, g: 0, b: 255, wantR: 255, wantG: 255, wantB: 0},
		{name: "Each channel is inverted independently", r: 1, g: 2, b: 3, wantR: 254, wantG: 253, wantB: 252},
		{name: "Middle gray flips to its neighbor", r: 128, g: 128, b: 128, wantR: 127, wantG: 127, wantB: 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := internal.Complementary(tt.r, tt.g, tt.b)

			assert.Equal(t, tt.wantR, r)
			assert.Equal(t, tt.wantG, g)
			assert.Equal(t, tt.wantB, b)
		})
	}
}

func TestComplementaryIsItsOwnInverse(t *testing.T) {
	r, g, b := internal.Complementary(internal.Complementary(12, 99, 200))

	assert.Equal(t, uint8(12), r)
	assert.Equal(t, uint8(99), g)
	assert.Equal(t, uint8(200), b)
}
