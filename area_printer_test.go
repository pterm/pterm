package pterm_test

import (
	"os"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestAreaPrinter_NilPrint(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p.Update("asd")

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GenericStart(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.DefaultArea
	p.GenericStart()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GenericStartRawOutput(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	pterm.DisableStyling()
	p := pterm.DefaultArea
	p.GenericStart()
	pterm.EnableStyling()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GenericStop(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.DefaultArea
	p.GenericStop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_RemoveWhenDone(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	a, _ := pterm.DefaultArea.WithRemoveWhenDone().Start()

	a.Update("asd")
	a.Stop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_CenterFullscreen(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	a, _ := pterm.DefaultArea.WithRemoveWhenDone().WithFullscreen().WithCenter(false).Start()

	a.Update("asd")
	a.Stop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GetContent(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	a, _ := pterm.DefaultArea.Start()

	for _, printable := range printables {
		a.Update(printable)
		testza.AssertEqual(t, a.GetContent(), pterm.Sprint(printable))
	}

	a.Stop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_WithRemoveWhenDone(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p2 := p.WithRemoveWhenDone()

	testza.AssertTrue(t, p2.RemoveWhenDone)

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_WithFullscreen(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p2 := p.WithFullscreen()

	testza.AssertTrue(t, p2.Fullscreen)

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_Clear(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p.Update("asd")

	p.Clear()

	os.Stdout = originalStdout // Restore original os.Stdout
}
