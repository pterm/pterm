package pterm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaPrinter_NilPrint(t *testing.T) {
	p := AreaPrinter{}
	p.Update("asd")
}

func TestAreaPrinter_GenericStart(t *testing.T) {
	p := DefaultArea
	p.GenericStart()
}

func TestAreaPrinter_GenericStartRawOutput(t *testing.T) {
	DisableStyling()
	p := DefaultArea
	p.GenericStart()
	EnableStyling()
}

func TestAreaPrinter_GenericStop(t *testing.T) {
	p := DefaultArea
	p.GenericStop()
}

func TestAreaPrinter_RemoveWhenDone(t *testing.T) {
	a, _ := DefaultArea.WithRemoveWhenDone().Start()

	a.Update("asd")
	a.Stop()
}

func TestAreaPrinter_CenterFullscreen(t *testing.T) {
	a, _ := DefaultArea.WithRemoveWhenDone().WithFullscreen().WithCenter().Start()

	a.Update("asd")
	a.Stop()
}

func TestAreaPrinter_GetContent(t *testing.T) {
	a, _ := DefaultArea.Start()

	for _, printable := range printables {
		a.Update(printable)
		assert.Equal(t, a.GetContent(), Sprint(printable))
		assert.Equal(t, a.GetContent(), a.content)
	}

	a.Stop()
}

func TestAreaPrinter_WithRemoveWhenDone(t *testing.T) {
	p := AreaPrinter{}
	p2 := p.WithRemoveWhenDone()

	assert.True(t, p2.RemoveWhenDone)
}

func TestAreaPrinter_WithFullscreen(t *testing.T) {
	p := AreaPrinter{}
	p2 := p.WithFullscreen()

	assert.True(t, p2.Fullscreen)
}
