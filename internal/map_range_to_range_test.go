package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapRangeToRange(t *testing.T) {
	assert.Equal(t, 127, MapRangeToRange(0, 100, 0, 255, 50))
	assert.Equal(t, 127, MapRangeToRange(0, 400, 0, 255, 200))
}
