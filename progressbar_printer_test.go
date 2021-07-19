package pterm_test

import (
	"testing"
	"time"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestProgressbarPrinter_Add(t *testing.T) {
	proxyToDevNull()
	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Add(1337)
	assert.Equal(t, 1337, p.Current)
	p.Stop()
}

func TestProgressbarPrinter_AddWithNoStyle(t *testing.T) {
	proxyToDevNull()
	p := pterm.ProgressbarPrinter{}.WithTotal(2000)
	p.Add(1337)
	testza.AssertEqual(t, 1337, p.Current)
	p.Stop()
}

func TestProgressbarPrinter_AddWithTotalOfZero(t *testing.T) {
	proxyToDevNull()
	p := pterm.ProgressbarPrinter{}.WithTotal(0)
	p.Add(1337)
	assert.Equal(t, 0, p.Current)
	p.Stop()
}

func TestProgressbarPrinter_AddTotalEqualsCurrent(t *testing.T) {
	proxyToDevNull()
	p := pterm.DefaultProgressbar.WithTotal(1)
	p.Start()
	p.Add(1)
	assert.Equal(t, 1, p.Current)
	assert.False(t, p.IsActive)
	p.Stop()
}

func TestProgressbarPrinter_RemoveWhenDone(t *testing.T) {
	proxyToDevNull()
	p, err := pterm.DefaultProgressbar.WithTotal(2).WithRemoveWhenDone().Start()
	assert.NoError(t, err)
	p.Stop()
	p.Add(1)
	assert.Equal(t, 1, p.Current)
	assert.False(t, p.IsActive)
}

func TestProgressbarPrinter_GenericStart(t *testing.T) {
	p := pterm.DefaultProgressbar
	p.GenericStart()
}

func TestProgressbarPrinter_GenericStartRawOutput(t *testing.T) {
	pterm.DisableStyling()
	p := pterm.DefaultProgressbar
	p.GenericStart()
	pterm.EnableStyling()
}

func TestProgressbarPrinter_GenericStop(t *testing.T) {
	p, err := pterm.DefaultProgressbar.Start()
	assert.NoError(t, err)
	p.GenericStop()
}

func TestProgressbarPrinter_GetElapsedTime(t *testing.T) {
	p := pterm.DefaultProgressbar
	p.Start()
	p.Stop()
	assert.NotEmpty(t, p.GetElapsedTime())
}

func TestProgressbarPrinter_Increment(t *testing.T) {
	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Increment()
	assert.Equal(t, 1, p.Current)
}

func TestProgressbarPrinter_WithBarStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithBarStyle(s)

	assert.Equal(t, s, p2.BarStyle)
}

func TestProgressbarPrinter_WithCurrent(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithCurrent(10)

	assert.Equal(t, 10, p2.Current)
}

func TestProgressbarPrinter_WithElapsedTimeRoundingFactor(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithElapsedTimeRoundingFactor(time.Hour)

	assert.Equal(t, time.Hour, p2.ElapsedTimeRoundingFactor)
}

func TestProgressbarPrinter_WithLastCharacter(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithLastCharacter(">")

	assert.Equal(t, ">", p2.LastCharacter)
}

func TestProgressbarPrinter_WithBarCharacter(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithBarCharacter("-")

	assert.Equal(t, "-", p2.BarCharacter)
}

func TestProgressbarPrinter_WithRemoveWhenDone(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithRemoveWhenDone()

	assert.True(t, p2.RemoveWhenDone)
}

func TestProgressbarPrinter_WithShowCount(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowCount()

	assert.True(t, p2.ShowCount)
}

func TestProgressbarPrinter_WithShowElapsedTime(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowElapsedTime()

	assert.True(t, p2.ShowElapsedTime)
}

func TestProgressbarPrinter_WithShowPercentage(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowPercentage()

	assert.True(t, p2.ShowPercentage)
}

func TestProgressbarPrinter_WithShowTitle(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowTitle()

	assert.True(t, p2.ShowTitle)
}

func TestProgressbarPrinter_WithTitle(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTitle("test")

	assert.Equal(t, "test", p2.Title)
}

func TestProgressbarPrinter_WithTitleStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTitleStyle(s)

	assert.Equal(t, s, p2.TitleStyle)
}

func TestProgressbarPrinter_WithTotal(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTotal(1337)

	assert.Equal(t, 1337, p2.Total)
}
