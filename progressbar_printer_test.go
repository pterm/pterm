package pterm_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestProgressbarPrinter_Add(t *testing.T) {
	proxyToDevNull()
	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Add(1337)
	testza.AssertEqual(t, 1337, p.Current)
	p.Stop()
}

func TestProgressbarPrinter_Add_With(t *testing.T) {
	proxyToDevNull()
	w := pterm.GetTerminalWidth()
	h := pterm.GetTerminalHeight()
	pterm.SetForcedTerminalSize(1, 1)
	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Add(1337)
	testza.AssertEqual(t, 1337, p.Current)
	p.Stop()
	pterm.SetForcedTerminalSize(w, h)
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
	testza.AssertEqual(t, 0, p.Current)
	p.Stop()
}

func TestProgressbarPrinter_AddTotalEqualsCurrent(t *testing.T) {
	proxyToDevNull()
	p := pterm.DefaultProgressbar.WithTotal(1)
	p.Start()
	p.Add(1)
	testza.AssertEqual(t, 1, p.Current)
	testza.AssertFalse(t, p.IsActive)
	p.Stop()
}

func TestProgressbarPrinter_RemoveWhenDone(t *testing.T) {
	proxyToDevNull()
	p, err := pterm.DefaultProgressbar.WithTotal(2).WithRemoveWhenDone().Start()
	testza.AssertNoError(t, err)
	p.Stop()
	p.Add(1)
	testza.AssertEqual(t, 1, p.Current)
	testza.AssertFalse(t, p.IsActive)
}

func TestProgressbarPrinter_StartWithTitle(t *testing.T) {
	p := pterm.DefaultProgressbar
	p2, _ := p.Start("Title")
	testza.AssertEqual(t, "Title", p2.Title)
	p.Stop()
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
	testza.AssertNoError(t, err)
	p.GenericStop()
}

func TestProgressbarPrinter_GetElapsedTime(t *testing.T) {
	p := pterm.DefaultProgressbar
	p.Start()
	p.Stop()
	testza.AssertNotZero(t, p.GetElapsedTime())
}

func TestProgressbarPrinter_Increment(t *testing.T) {
	p := pterm.DefaultProgressbar.WithTotal(2000)
	p.Increment()
	testza.AssertEqual(t, 1, p.Current)
}

func TestProgressbarPrinter_WithBarStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithBarStyle(s)

	testza.AssertEqual(t, s, p2.BarStyle)
}

func TestProgressbarPrinter_WithCurrent(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithCurrent(10)

	testza.AssertEqual(t, 10, p2.Current)
}

func TestProgressbarPrinter_WithElapsedTimeRoundingFactor(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithElapsedTimeRoundingFactor(time.Hour)

	testza.AssertEqual(t, time.Hour, p2.ElapsedTimeRoundingFactor)
}

func TestProgressbarPrinter_WithLastCharacter(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithLastCharacter(">")

	testza.AssertEqual(t, ">", p2.LastCharacter)
}

func TestProgressbarPrinter_WithBarCharacter(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithBarCharacter("-")

	testza.AssertEqual(t, "-", p2.BarCharacter)
}

func TestProgressbarPrinter_WithRemoveWhenDone(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithRemoveWhenDone()

	testza.AssertTrue(t, p2.RemoveWhenDone)
}

func TestProgressbarPrinter_WithShowCount(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowCount()

	testza.AssertTrue(t, p2.ShowCount)
}

func TestProgressbarPrinter_WithShowElapsedTime(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowElapsedTime()

	testza.AssertTrue(t, p2.ShowElapsedTime)
}

func TestProgressbarPrinter_WithShowPercentage(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowPercentage()

	testza.AssertTrue(t, p2.ShowPercentage)
}

func TestProgressbarPrinter_WithShowTitle(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithShowTitle()

	testza.AssertTrue(t, p2.ShowTitle)
}

func TestProgressbarPrinter_WithTitle(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTitle("test")

	testza.AssertEqual(t, "test", p2.Title)
}

func TestProgressbarPrinter_WithTitleStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTitleStyle(s)

	testza.AssertEqual(t, s, p2.TitleStyle)
}

func TestProgressbarPrinter_WithTotal(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTotal(1337)

	testza.AssertEqual(t, 1337, p2.Total)
}

func TestProgressbarPrinter_WithMaxWidth(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithMaxWidth(1337)

	testza.AssertEqual(t, 1337, p2.MaxWidth)
}

func TestProgressbarPrinter_WithBarFiller(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithBarFiller("-")

	testza.AssertEqual(t, "-", p2.BarFiller)
}

func TestProgressbarPrinter_UpdateTitle(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	p2 := p.WithTitle("test")
	p2.UpdateTitle("test2")

	testza.AssertEqual(t, "test2", p2.Title)
}

func TestProgressbarPrinter_WithWriter(t *testing.T) {
	p := pterm.ProgressbarPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	testza.AssertEqual(t, s, p2.Writer)
	testza.AssertZero(t, p.Writer)
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
			stderr, err := testza.CaptureStderr(func(w io.Writer) error {
				pb, err := pterm.DefaultProgressbar.WithTitle("Hello world").WithWriter(os.Stderr).Start()
				time.Sleep(time.Second) // Required otherwise the goroutine doesn't run and the text isnt outputted
				testza.AssertNoError(t, err)
				testCase.action(pb)
				time.Sleep(time.Second) // Required otherwise the goroutine doesn't run and the text isnt updated
				return nil
			})

			testza.AssertNoError(t, err)
			testza.AssertContains(t, stderr, "Hello world")
			testza.AssertContains(t, stderr, testCase.expectOutputToContain)
		})
	}
}
