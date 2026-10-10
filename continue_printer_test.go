package pterm_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

func TestContinueKeys(t *testing.T) {
	tests := []struct {
		name  string
		p     *pterm.ContinuePrinter
		input string
		want  string
	}{
		{"handle", &pterm.Continue, "a", "all"},
		{"uppercase handle", &pterm.Continue, "N", "no"},
		{"enter takes the default", &pterm.Continue, keyEnter, "yes"},
		{"default value", pterm.Continue.WithDefaultValue("cancel"), keyEnter, "cancel"},
		{"default index", pterm.Continue.WithDefaultValueIndex(2), keyEnter, "all"},
		{"index out of range", pterm.Continue.WithDefaultValueIndex(9), keyEnter, "cancel"},
		{"arrows move", &pterm.Continue, keyRight + keyRight + keyEnter, "all"},
		{"left wraps", &pterm.Continue, keyLeft + keyEnter, "cancel"},
		{"unknown keys are ignored", &pterm.Continue, "xz" + "c", "cancel"},
		{"custom options", pterm.Continue.WithOptions([]string{"overwrite", "skip", "quit"}), "s", "skip"},
		{"custom handles", pterm.Continue.WithOptions([]string{"overwrite", "skip"}).WithHandles([]string{"w", "k"}), "k", "skip"},
		{"wrong number of handles", pterm.Continue.WithHandles([]string{"x"}), "n", "no"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(t)

			got, err := tt.p.WithInput(strings.NewReader(tt.input)).Show()
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestContinueSnapshots(t *testing.T) {
	view := func(p *pterm.ContinuePrinter, input string) func() string {
		return func() string {
			return waiting(t, func() {
				_, _ = p.WithInput(strings.NewReader(input)).Show("Overwrite config.yaml?")
			})
		}
	}

	snap(t, "Default", view(&pterm.Continue, keyRight+keyEnter))
	snap(t, "NoOptions", view(pterm.Continue.WithShowOptions(false), keyRight+keyEnter))
	snap(t, "Custom", view(pterm.Continue.WithOptions([]string{"overwrite", "skip", "quit"}).WithDefaultValue("skip"), keyLeft+keyEnter))
}

func TestContinueCollapses(t *testing.T) {
	term := setup(t)

	_, _ = pterm.Continue.WithInput(strings.NewReader("a")).Show("Overwrite config.yaml?")

	if got, want := term.Screen(), "✔ Overwrite config.yaml? all"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}
}

func TestContinueNotInteractive(t *testing.T) {
	setup(t)

	restore := pterm.SetStdinIsTerminal(false)
	defer restore()

	got, err := pterm.Continue.WithDefaultValue("no").Show()
	if err != nil || got != "no" {
		t.Errorf("got %q, %v; want the default value", got, err)
	}
}

func TestContinueErrors(t *testing.T) {
	term := setup(t)

	if _, err := pterm.Continue.WithOptions(nil).Show(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}

	_, err := pterm.Continue.WithInput(strings.NewReader(keyCtrlC)).WithOnInterrupt(func() {}).Show("Overwrite?")
	if !errors.Is(err, pterm.ErrInterrupted) {
		t.Errorf("err = %v, want ErrInterrupted", err)
	}

	if got, want := term.Screen(), "✖ Overwrite? canceled"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	if _, err := pterm.Continue.WithContext(ctx).WithInput(strings.NewReader("y")).Show(); !errors.Is(err, pterm.ErrCanceled) {
		t.Errorf("err = %v, want ErrCanceled", err)
	}
}
