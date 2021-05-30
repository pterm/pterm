package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenterText(t *testing.T) {
	assert.Equal(t, "  Hello Wolrd  \n      !!!      ", CenterText("Hello Wolrd\n!!!", 15))
	assert.Equal(t, "Hello\n Wolr\n  d  \n !!! ", CenterText("Hello Wolrd\n!!!", 5))
}
