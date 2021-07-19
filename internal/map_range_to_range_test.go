package internal_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm/internal"
)

func TestMapRangeToRange(t *testing.T) {
	testza.AssertEqual(t, 127, internal.MapRangeToRange(0, 100, 0, 255, 50))
	testza.AssertEqual(t, 127, internal.MapRangeToRange(0, 400, 0, 255, 200))
	testza.AssertEqual(t, 127, internal.MapRangeToRange(-200, 200, 0, 255, 0))
	testza.AssertEqual(t, 127, internal.MapRangeToRange(0, 200.123, 0, 254.3, 100))
}
