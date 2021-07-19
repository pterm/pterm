package internal_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm/internal"
)

func TestCenterText(t *testing.T) {
	testza.AssertEqual(t, "  Hello Wolrd  \n      !!!      ", internal.CenterText("Hello Wolrd\n!!!", 15))
	testza.AssertEqual(t, "Hello\n Wolr\n  d  \n !!! ", internal.CenterText("Hello Wolrd\n!!!", 5))
}
