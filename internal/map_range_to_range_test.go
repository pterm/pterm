package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapRangeToRange(t *testing.T) {
	assert.Equal(t, 127, MapRangeToRange(0, 100, 0, 255, 50))
	assert.Equal(t, 127, MapRangeToRange(0, 400, 0, 255, 200))
	assert.Equal(t, 127, MapRangeToRange(-200, 200, 0, 255, 0))
	assert.Equal(t, 127, MapRangeToRange(0, 200.123, 0, 254.3, 100))
}
