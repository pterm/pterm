package internal_test

import (
	"testing"

	"github.com/pterm/pterm/internal/testhelper"
	"github.com/pterm/pterm/internal"
)

func TestMapRangeToRange(t *testing.T) {
	testhelper.AssertEqual(t, 127, internal.MapRangeToRange(0, 100, 0, 255, 50))
	testhelper.AssertEqual(t, 127, internal.MapRangeToRange(0, 400, 0, 255, 200))
	testhelper.AssertEqual(t, 127, internal.MapRangeToRange(-200, 200, 0, 255, 0))
	testhelper.AssertEqual(t, 127, internal.MapRangeToRange(0, 200.123, 0, 254.3, 100))
}
