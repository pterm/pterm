package pterm

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestProgressbar_Add(t *testing.T) {
	p := DefaultProgressbar.WithTotal(2000)
	p.Add(1337)
	assert.Equal(t, 1337, p.Current)
	p.Stop()
}

func TestProgressbar_AddTotalEqualsCurrent(t *testing.T) {
	p := DefaultProgressbar.WithTotal(1)
	p.Start()
	p.Add(1)
	assert.Equal(t, 1, p.Current)
	assert.False(t, p.IsActive)
	p.Stop()
}

func TestProgressbarRemoveWhenDone(t *testing.T) {
	p := DefaultProgressbar.WithTotal(1).WithRemoveWhenDone()
	p.Start()
	p.Add(1)
	assert.Equal(t, 1, p.Current)
	assert.False(t, p.IsActive)
}

func TestProgressbar_GenericStart(t *testing.T) {
	p := DefaultProgressbar
	p.GenericStart()
}

func TestProgressbar_GenericStop(t *testing.T) {
	p := DefaultProgressbar
	p.GenericStop()
}

func TestProgressbar_GetElapsedTime(t *testing.T) {
	p := DefaultProgressbar
	p.Start()
	p.Stop()
	assert.NotEmpty(t, p.GetElapsedTime())
}

func TestProgressbar_Increment(t *testing.T) {
	p := DefaultProgressbar.WithTotal(2000)
	p.Increment()
	assert.Equal(t, 1, p.Current)
}

func TestProgressbar_WithBarStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := Progressbar{}
	p2 := p.WithBarStyle(s)

	assert.Equal(t, s, p2.BarStyle)
}

func TestProgressbar_WithCurrent(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithCurrent(10)

	assert.Equal(t, 10, p2.Current)
}

func TestProgressbar_WithElapsedTimeRoundingFactor(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithElapsedTimeRoundingFactor(time.Hour)

	assert.Equal(t, time.Hour, p2.ElapsedTimeRoundingFactor)
}

func TestProgressbar_WithLastCharacter(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithLastCharacter(">")

	assert.Equal(t, ">", p2.LastCharacter)
}

func TestProgressbar_WithBarCharacter(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithBarCharacter("-")

	assert.Equal(t, "-", p2.BarCharacter)
}

func TestProgressbar_WithRemoveWhenDone(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithRemoveWhenDone()

	assert.True(t, p2.RemoveWhenDone)
}

func TestProgressbar_WithShowCount(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithShowCount()

	assert.True(t, p2.ShowCount)
}

func TestProgressbar_WithShowElapsedTime(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithShowElapsedTime()

	assert.True(t, p2.ShowElapsedTime)
}

func TestProgressbar_WithShowPercentage(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithShowPercentage()

	assert.True(t, p2.ShowPercentage)
}

func TestProgressbar_WithShowTitle(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithShowTitle()

	assert.True(t, p2.ShowTitle)
}

func TestProgressbar_WithTitle(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithTitle("test")

	assert.Equal(t, "test", p2.Title)
}

func TestProgressbar_WithTitleStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := Progressbar{}
	p2 := p.WithTitleStyle(s)

	assert.Equal(t, s, p2.TitleStyle)
}

func TestProgressbar_WithTotal(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithTotal(1337)

	assert.Equal(t, 1337, p2.Total)
}

func TestProgressbar_WithUpdateDelay(t *testing.T) {
	p := Progressbar{}
	p2 := p.WithUpdateDelay(time.Hour)

	assert.Equal(t, time.Hour, p2.UpdateDelay)
}

func TestProgressbar_parseElapsedTime(t *testing.T) {
	p := Progressbar{}
	p.Start()
	p.Stop()
	assert.NotEmpty(t, p.parseElapsedTime())
}
