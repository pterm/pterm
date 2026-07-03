package pterm_test

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestProgressbarPrinter_Add(t *testing.T) {
	proxyToDevNull()

	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Add(1337)
	assert.Equal(t, 1337, p.Current)
	_, _ = p.Stop()
}

func TestProgressbarPrinter_Add_With(t *testing.T) {
	proxyToDevNull()

	w := pterm.GetTerminalWidth()
	h := pterm.GetTerminalHeight()

	pterm.SetForcedTerminalSize(1, 1)

	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Add(1337)
	assert.Equal(t, 1337, p.Current)
	_, _ = p.Stop()

	pterm.SetForcedTerminalSize(w, h)
}

func TestProgressbarPrinter_AddWithNoStyle(t *testing.T) {
	proxyToDevNull()

	p := pterm.ProgressbarPrinter{}.WithTotal(2000)
	p.Add(1337)
	assert.Equal(t, 1337, p.Current)
	_, _ = p.Stop()
}

func TestProgressbarPrinter_AddWithTotalOfZero(t *testing.T) {
	proxyToDevNull()

	p := pterm.ProgressbarPrinter{}.WithTotal(0)
	p.Add(1337)
	assert.Equal(t, 0, p.Current)
	_, _ = p.Stop()
}

func TestProgressbarPrinter_AddTotalEqualsCurrent(t *testing.T) {
	proxyToDevNull()

	// Start returns the started instance; using the original would leak the
	// started bar's goroutine and pollute the output of later tests.
	p, err := pterm.DefaultProgressbar.WithTotal(1).Start()
	assert.NoError(t, err)

	defer p.Stop()

	p.Add(1)
	assert.Equal(t, 1, p.Current)
	assert.False(t, p.IsActive, "the progressbar should auto-stop when reaching its total")
}

func TestProgressbarPrinter_RemoveWhenDone(t *testing.T) {
	proxyToDevNull()

	p, err := pterm.DefaultProgressbar.WithTotal(2).WithRemoveWhenDone().Start()
	assert.NoError(t, err)

	_, _ = p.Stop()
	p.Add(1)
	assert.Equal(t, 1, p.Current)
	assert.False(t, p.IsActive)
}

func TestProgressbarPrinter_StartWithTitle(t *testing.T) {
	p := pterm.DefaultProgressbar
	p2, _ := p.Start("Title")
	assert.Equal(t, "Title", p2.Title)
	_, _ = p2.Stop()
}

func TestProgressbarPrinter_GenericStart(_ *testing.T) {
	p := pterm.DefaultProgressbar

	lp, _ := p.GenericStart()
	if lp != nil {
		_, _ = (*lp).GenericStop()
	}
}

func TestProgressbarPrinter_GenericStartRawOutput(_ *testing.T) {
	pterm.DisableStyling()

	p := pterm.DefaultProgressbar

	lp, _ := p.GenericStart()
	if lp != nil {
		_, _ = (*lp).GenericStop()
	}

	pterm.EnableStyling()
}

func TestProgressbarPrinter_GenericStop(t *testing.T) {
	p, err := pterm.DefaultProgressbar.Start()
	assert.NoError(t, err)

	_, _ = p.GenericStop()
}

func TestProgressbarPrinter_GetElapsedTime(t *testing.T) {
	p := pterm.DefaultProgressbar
	p2, _ := p.Start()
	_, _ = p2.Stop()
	assert.NotZero(t, p2.GetElapsedTime())
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

func TestProgressbarPrinter_WithMaxWidth(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithMaxWidth(1337)

	assert.Equal(t, 1337, p2.MaxWidth)
}

func TestProgressbarPrinter_WithBarFiller(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithBarFiller("-")

	assert.Equal(t, "-", p2.BarFiller)
}

func TestProgressbarPrinter_UpdateTitle(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTitle("test")
	p2.UpdateTitle("test2")

	assert.Equal(t, "test2", p2.Title)
}

func TestProgressbarPrinter_WithWriter(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}

func TestProgressbarPrinter_OutputToWriters(t *testing.T) {
	testCases := map[string]struct {
		action                func(*pterm.ProgressbarPrinter)
		expectOutputToContain string
	}{
		"ExpectUpdatedTitleToBeWrittenToStderr": {
			action: func(pb *pterm.ProgressbarPrinter) {
				pb.UpdateTitle("Updated text")
			},
			expectOutputToContain: "Updated text",
		},
	}

	for testTitle, testCase := range testCases {
		t.Run(testTitle, func(t *testing.T) {
			buf := &syncBuffer{}
			pb, err := pterm.DefaultProgressbar.WithTitle("Hello world").WithWriter(buf).Start()
			assert.NoError(t, err)

			defer pb.Stop()

			waitForOutput(t, buf, "Hello world")
			testCase.action(pb)
			waitForOutput(t, buf, testCase.expectOutputToContain)
		})
	}
}

// waitForOutput polls buf until it contains want, failing the test after a
// generous timeout. Polling keeps the test fast in the common case without
// relying on fixed sleeps that get flaky under -race.
func waitForOutput(t *testing.T, buf *syncBuffer, want string) {
	t.Helper()

	deadline := time.Now().Add(10 * time.Second)
	for time.Now().Before(deadline) {
		if strings.Contains(buf.String(), want) {
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Errorf("output did not contain %q within timeout, got:\n%s", want, buf.String())
}
