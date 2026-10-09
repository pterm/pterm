package pterm_test

import (
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

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
		pterm.EnableOutput()
		pterm.DisableDebugMessages()
		setupStdoutCapture()
	}()

	const iterations = 200

	var wg sync.WaitGroup

	// Global toggle writers.
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
	wg.Go(func() {
		for range iterations {
			pterm.EnableOutput()
			pterm.DisableOutput()
		}
	})

	// The default writer is swapped while other goroutines print to it.
	wg.Go(func() {
		var alt syncBuffer

		for range iterations {
			pterm.SetDefaultOutput(io.Discard)
			pterm.SetDefaultOutput(&alt)

			_ = pterm.GetDefaultOutput()
		}
	})

	// Printers rendering to strings.
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

	// Printers writing to the (concurrently swapped) default writer.
	wg.Go(func() {
		for i := range iterations {
			pterm.Println("to default writer", i)
			pterm.Printo("overwrite")
		}
	})

	// Loggers share a package-level mutex; print from two at once.
	for range 2 {
		wg.Go(func() {
			logger := pterm.DefaultLogger.WithWriter(io.Discard)
			for i := range iterations {
				logger.Info("concurrent log", logger.Args("i", i))
			}
		})
	}

	// Live printers.
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

// TestConcurrentSetDefaultOutputIsConsistent verifies that a writer swapped in
// under contention is observed consistently: every GetDefaultOutput result is
// one of the writers that were actually set.
func TestConcurrentSetDefaultOutputIsConsistent(t *testing.T) {
	t.Cleanup(setupStdoutCapture)

	writerA := &syncBuffer{}
	writerB := &syncBuffer{}

	var wg sync.WaitGroup

	wg.Go(func() {
		for range 500 {
			pterm.SetDefaultOutput(writerA)
			pterm.SetDefaultOutput(writerB)
		}
	})

	seen := make(map[io.Writer]bool)

	for range 500 {
		seen[pterm.GetDefaultOutput()] = true
	}

	wg.Wait()

	for writer := range seen {
		assert.Contains(t, []io.Writer{&outBuf, writerA, writerB}, writer,
			"GetDefaultOutput returned a writer that was never set")
	}
}
