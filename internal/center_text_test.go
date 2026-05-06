package internal_test

import (
	"testing"

	"github.com/pterm/pterm/internal/testhelper"
	"github.com/pterm/pterm/internal"
)

func TestCenterText(t *testing.T) {
	testhelper.AssertEqual(t, "  Hello Wolrd  \n      !!!      ", internal.CenterText("Hello Wolrd\n!!!", 15))
	testhelper.AssertEqual(t, "Hello\n Wolr\n  d  \n !!! ", internal.CenterText("Hello Wolrd\n!!!", 5))
}
