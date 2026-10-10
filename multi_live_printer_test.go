package pterm_test

import (
	"bytes"
	"testing"

	"github.com/pterm/pterm"
)

func TestMultiStopsGroup(t *testing.T) {
	term := setup(t)

	multi, _ := pterm.Multi.Start()
	if multi.NewWriter() != pterm.DefaultOutput() {
		t.Error("NewWriter does not return the default output")
	}

	spinner, _ := pterm.Spinner.WithShowTimer(false).WithWriter(multi.NewWriter()).Start("Running tests")
	bar, _ := quietBar.WithTotal(4).WithWriter(multi.NewWriter()).Start("Building")
	multi.Add(spinner, bar)

	bar.Add(4)

	if err := multi.Stop(); err != nil {
		t.Fatal(err)
	}

	if got, want := term.Screen(), "✔ Building 4/4\n✔ Running tests"; got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}

	// Printers added after Stop are stopped right away.
	late, _ := pterm.Spinner.WithShowTimer(false).Start("Late")
	multi.Add(late)

	if got := term.Screen(); got != "✔ Building 4/4\n✔ Running tests\n✔ Late" {
		t.Errorf("screen = %q", got)
	}
}

func TestMultiWriter(t *testing.T) {
	setup(t)

	var buf bytes.Buffer

	multi, _ := pterm.Multi.WithWriter(&buf).Start()
	if multi.NewWriter() != &buf {
		t.Error("NewWriter does not return the configured writer")
	}

	var unstarted pterm.MultiPrinter

	unstarted.Add(nil)

	if err := unstarted.Stop(); err != nil {
		t.Errorf("Stop on an unstarted group = %v", err)
	}
}
