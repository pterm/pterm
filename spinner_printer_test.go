package pterm_test

import (
	"context"
	"testing"
	"time"

	"github.com/pterm/pterm"
)

func TestSpinnerLeavesStatusLine(t *testing.T) {
	term := setup(t)

	s, err := pterm.Spinner.WithShowTimer(false).Start("Downloading")
	if err != nil {
		t.Fatal(err)
	}

	pterm.Println("log line")
	s.Success("Downloaded")

	if got, want := term.Screen(), "log line\n✔ Downloaded"; got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}

	if term.CursorHidden() {
		t.Error("cursor still hidden")
	}
}

func TestSpinnerStatuses(t *testing.T) {
	term := setup(t)
	sp := pterm.Spinner.WithShowTimer(false)

	for _, end := range []func(*pterm.SpinnerPrinter){
		func(s *pterm.SpinnerPrinter) { s.Fail("failed") },
		func(s *pterm.SpinnerPrinter) { s.Warning("warned") },
		func(s *pterm.SpinnerPrinter) { s.Info("noted") },
		func(s *pterm.SpinnerPrinter) { s.Skip("skipped") },
	} {
		s, _ := sp.Start("working")
		end(s)
	}

	want := "✖ failed\n▲ warned\n● noted\n– skipped"
	if got := term.Screen(); got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}
}

func TestSpinnerContextCancels(t *testing.T) {
	term := setup(t)

	ctx, cancel := context.WithCancel(context.Background())
	s, _ := pterm.Spinner.WithShowTimer(false).WithDelay(time.Millisecond).WithContext(ctx).Start("waiting")

	cancel()

	deadline := time.Now().Add(time.Second)
	for term.Screen() != "✖ waiting canceled" && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}

	if got := term.Screen(); got != "✖ waiting canceled" {
		t.Errorf("screen = %q", got)
	}

	s.Success() // no effect after cancel
}

func TestSpinnerNotTerminal(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		s, _ := pterm.Spinner.WithShowTimer(false).Start("building")

		time.Sleep(20 * time.Millisecond)
		s.Success("built")
	})
	if out != "✔ built\n" {
		t.Errorf("output = %q", out)
	}
}

func TestSpinnerRemoveWhenDone(t *testing.T) {
	term := setup(t)

	s, _ := pterm.Spinner.WithRemoveWhenDone().Start("temporary")
	_ = s.Stop()

	if got := term.Screen(); got != "" {
		t.Errorf("screen = %q", got)
	}
}
