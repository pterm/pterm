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

func TestProgressbarPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.ProgressbarPrinter{})
}

func TestProgressbarPrinter_Add_WithTotal(t *testing.T) {
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
