package pterm_test

import (
	"os"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestAreaPrinter_NilPrint(_ *testing.T) {
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
	started, err := p.GenericStart()
	assert.NoError(t, err)

	_, _ = (*started).GenericStop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GenericStartRawOutput(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	pterm.DisableStyling()

	defer pterm.EnableStyling()

	p := pterm.DefaultArea
	started, err := p.GenericStart()
	assert.NoError(t, err)

	_, _ = (*started).GenericStop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GenericStop(_ *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.DefaultArea
	_, _ = p.GenericStop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_RemoveWhenDone(_ *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	a, _ := pterm.DefaultArea.WithRemoveWhenDone().Start()

	a.Update("asd")
	_ = a.Stop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_CenterFullscreen(_ *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	a, _ := pterm.DefaultArea.WithRemoveWhenDone().WithFullscreen().WithCenter().Start()

	a.Update("asd")
	_ = a.Stop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_GetContent(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	a, _ := pterm.DefaultArea.Start()

	for _, printable := range printables {
		a.Update(printable)
		assert.Equal(t, a.GetContent(), pterm.Sprint(printable))
	}

	_ = a.Stop()

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_WithRemoveWhenDone(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p2 := p.WithRemoveWhenDone()

	assert.True(t, p2.RemoveWhenDone)

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_WithFullscreen(t *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p2 := p.WithFullscreen()

	assert.True(t, p2.Fullscreen)

	os.Stdout = originalStdout // Restore original os.Stdout
}

func TestAreaPrinter_Clear(_ *testing.T) {
	originalStdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull) // Set os.Stdout to DevNull to hide output from cursor.Area

	p := pterm.AreaPrinter{}
	p.Update("asd")

	p.Clear()

	os.Stdout = originalStdout // Restore original os.Stdout
}
