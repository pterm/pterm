package internal_test

import (
	"testing"

	"github.com/pterm/pterm/internal"
	"github.com/stretchr/testify/assert"
)

func TestCenterText(t *testing.T) {
	assert.Equal(t, "  Hello Wolrd  \n      !!!      ", internal.CenterText("Hello Wolrd\n!!!", 15))
	assert.Equal(t, "Hello\n Wolr\n  d  \n !!! ", internal.CenterText("Hello Wolrd\n!!!", 5))
}
