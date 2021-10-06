package pterm_test

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestAreaPrinter_NilPrint(t *testing.T) {
	p := pterm.AreaPrinter{}
	p.Update("asd")
}

func TestAreaPrinter_GenericStart(t *testing.T) {
	p := pterm.DefaultArea
	p.GenericStart()
}

func TestAreaPrinter_GenericStartRawOutput(t *testing.T) {
	pterm.DisableStyling()
	p := pterm.DefaultArea
	p.GenericStart()
	pterm.EnableStyling()
}

func TestAreaPrinter_GenericStop(t *testing.T) {
	p := pterm.DefaultArea
	p.GenericStop()
}

func TestAreaPrinter_RemoveWhenDone(t *testing.T) {
	a, _ := pterm.DefaultArea.WithRemoveWhenDone().Start()

	a.Update("asd")
	a.Stop()
}

func TestAreaPrinter_CenterFullscreen(t *testing.T) {
	a, _ := pterm.DefaultArea.WithRemoveWhenDone().WithFullscreen().WithCenter().Start()

	a.Update("asd")
	a.Stop()
}

func TestAreaPrinter_GetContent(t *testing.T) {
	a, _ := pterm.DefaultArea.Start()

	for _, printable := range printables {
		a.Update(printable)
		testza.AssertEqual(t, a.GetContent(), pterm.Sprint(printable))
	}

	a.Stop()
}

func TestAreaPrinter_WithRemoveWhenDone(t *testing.T) {
	p := pterm.AreaPrinter{}
	p2 := p.WithRemoveWhenDone()

	testza.AssertTrue(t, p2.RemoveWhenDone)
}

func TestAreaPrinter_WithFullscreen(t *testing.T) {
	p := pterm.AreaPrinter{}
	p2 := p.WithFullscreen()

	testza.AssertTrue(t, p2.Fullscreen)
}

func TestAreaPrinter_Clear(t *testing.T) {
	p := pterm.AreaPrinter{}
	p.Update("asd")

	p.Clear()
}
