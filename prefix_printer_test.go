package pterm_test

import (
	"errors"
	"testing"

	"github.com/pterm/pterm"
)

func TestPrefixSnapshots(t *testing.T) {
	snap(t, "Info", func() string { return pterm.Info.Sprint("Fetching the latest release") })
	snap(t, "Success", func() string { return pterm.Success.Sprint("Deployed") })
	snap(t, "Warning", func() string { return pterm.Warning.Sprint("The cache is stale") })
	snap(t, "Error", func() string { return pterm.Error.Sprint("Unknown command") })
	snap(t, "ErrorMultiline", func() string {
		return pterm.Error.Sprint("Unknown command disk\nRun `mjw --help` to see available commands.")
	})
	snap(t, "WithTitle", func() string { return pterm.Info.WithTitle("Note").Sprint("Neutral information.") })
	snap(t, "Wrapped", func() string {
		return pterm.Info.Sprint("This message is long enough that it has to wrap onto a second line on an eighty column terminal.")
	})
}

func TestPrefixWriters(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Info.Println("to stdout")
	})
	if out != "● to stdout\n" {
		t.Errorf("output = %q", out)
	}
}

func TestDebugHiddenByDefault(t *testing.T) {
	setup(t)

	out := capture(t, func() { pterm.Debug.Println("hidden") })
	if out != "" {
		t.Errorf("debug printed %q while disabled", out)
	}

	pterm.EnableDebugMessages()
	defer pterm.DisableDebugMessages()

	out = capture(t, func() { pterm.Debug.Println("shown") })
	if out != "◦ shown\n" {
		t.Errorf("debug output = %q", out)
	}
}

func TestPrintOnError(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Error.PrintOnError(nil, errors.New("boom"), "not an error")
	})
	if out != "✖ boom\n" {
		t.Errorf("output = %q", out)
	}
}

func TestOutputDisabled(t *testing.T) {
	setup(t)
	pterm.DisableOutput()

	defer pterm.EnableOutput()

	out := capture(t, func() { pterm.Success.Println("quiet") })
	if out != "" {
		t.Errorf("printed %q with output disabled", out)
	}
}
