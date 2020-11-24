package pterm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBar_WithLabel(t *testing.T) {
	p := Bar{}
	s := "X"
	p2 := p.WithLabel(s)

	assert.Equal(t, s, p2.Label)
	assert.Empty(t, p.Label)
}

func TestBar_WithStyle(t *testing.T) {
	p := Bar{}
	s := NewStyle(FgRed, BgBlue, Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Empty(t, p.Style)
}

func TestBar_WithValue(t *testing.T) {
	p := Bar{}
	s := 1337
	p2 := p.WithValue(s)

	assert.Equal(t, s, p2.Value)
	assert.Empty(t, p.Value)
}
