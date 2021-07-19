package internal

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestCenterText(t *testing.T) {
	testza.AssertEqual(t, "  Hello Wolrd  \n      !!!      ", CenterText("Hello Wolrd\n!!!", 15))
	testza.AssertEqual(t, "Hello\n Wolr\n  d  \n !!! ", CenterText("Hello Wolrd\n!!!", 5))
}
