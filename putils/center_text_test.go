package putils

import (
	"testing"

	"github.com/pterm/pterm/internal/testhelper"
)

func TestCenterText(t *testing.T) {
	testhelper.AssertEqual(t, "Hello Wolrd\n    !!!    ", CenterText("Hello Wolrd\n!!!"))
}
