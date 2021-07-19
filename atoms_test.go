package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestBar_WithLabel(t *testing.T) {
	p := pterm.Bar{}
	s := "X"
	p2 := p.WithLabel(s)

	assert.Equal(t, s, p2.Label)
	assert.Empty(t, p.Label)
}

func TestBar_WithStyle(t *testing.T) {
	p := pterm.Bar{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
	assert.Empty(t, p.Style)
}

func TestBar_WithValue(t *testing.T) {
	p := pterm.Bar{}
	s := 1337
	p2 := p.WithValue(s)

	assert.Equal(t, s, p2.Value)
	assert.Empty(t, p.Value)
}

func TestBar_WithLabelStyle(t *testing.T) {
	p := pterm.Bar{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p2 := p.WithLabelStyle(s)

	assert.Equal(t, s, p2.LabelStyle)
	assert.Empty(t, p.LabelStyle)
}
