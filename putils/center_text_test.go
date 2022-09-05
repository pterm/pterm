package putils

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestCenterText(t *testing.T) {
	testza.AssertEqual(t, "Hello Wolrd\n    !!!    ", CenterText("Hello Wolrd\n!!!"))
}
