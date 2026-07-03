package putils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCenterText(t *testing.T) {
	assert.Equal(t, "Hello Wolrd\n    !!!    ", CenterText("Hello Wolrd\n!!!"))
}
