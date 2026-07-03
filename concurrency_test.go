package pterm_test

import (
	"io"
	"sync"
	"testing"

	"github.com/pterm/pterm"
)

// TestConcurrencySmoke hammers pterm from many goroutines at once while the
// global configuration is being toggled. It exists to give the race detector
// something to chew on: `go test -race` fails this test if any printer or
// setter touches shared state without synchronization.
func TestConcurrencySmoke(_ *testing.T) {
	defer func() {
		pterm.EnableStyling()
		pterm.EnableColor()
		pterm.DisableDebugMessages()
	}()

	const iterations = 200
	var wg sync.WaitGroup

	wg.Go(func() {
		for range iterations {
			pterm.EnableColor()
			pterm.DisableColor()
		}
	})
	wg.Go(func() {
		for range iterations {
			pterm.EnableStyling()
			pterm.DisableStyling()
		}
	})
	wg.Go(func() {
		for range iterations {
			pterm.EnableDebugMessages()
			pterm.DisableDebugMessages()
		}
	})

	for range 4 {
		wg.Go(func() {
			for i := range iterations {
				pterm.Info.Sprint("concurrent sprint ", i)
				pterm.FgRed.Sprint("red text")
				pterm.NewRGB(1, 2, 3).Sprint("rgb text")
				pterm.NewStyle(pterm.FgGreen, pterm.Bold).Sprint("styled text")
				pterm.DefaultBox.Sprint("boxed text")
			}
		})
	}

	wg.Go(func() {
		for range 10 {
			spinner, err := pterm.DefaultSpinner.WithWriter(io.Discard).Start("spinning")
			if err == nil {
				_ = spinner.Stop()
			}
		}
	})
	wg.Go(func() {
		for range 10 {
			bar, err := pterm.DefaultProgressbar.WithWriter(io.Discard).WithTotal(10).Start()
			if err == nil {
				for range 10 {
					bar.Increment()
				}

				_, _ = bar.Stop()
			}
		}
	})

	wg.Wait()
}
