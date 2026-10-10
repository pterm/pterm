package pterm_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

func TestConfirmKeys(t *testing.T) {
	tests := []struct {
		name  string
		input string
		def   bool
		want  bool
	}{
		{"y", "y", false, true},
		{"n", "n", true, false},
		{"enter takes default", "\r", true, true},
		{"arrow toggles", "\x1b[C\r", false, true},
		{"uppercase", "Y", false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(t)

			got, err := pterm.Confirm.WithDefaultValue(tt.def).WithInput(strings.NewReader(tt.input)).Show("Delete?")
			if err != nil {
				t.Fatal(err)
			}

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmCollapses(t *testing.T) {
	term := setup(t)

	_, _ = pterm.Confirm.WithInput(strings.NewReader("y")).Show("Delete 23 caches?")

	if got, want := term.Screen(), "✔ Delete 23 caches? Yes"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}
}

func TestConfirmNotInteractive(t *testing.T) {
	setup(t)

	restore := pterm.SetStdinIsTerminal(false)
	defer restore()

	got, err := pterm.Confirm.WithDefaultValue(true).Show()
	if err != nil || !got {
		t.Errorf("got %v, %v; want the default value", got, err)
	}
}

func TestConfirmInterrupt(t *testing.T) {
	term := setup(t)

	called := false

	_, err := pterm.Confirm.
		WithInput(strings.NewReader("\x03")).
		WithOnInterrupt(func() { called = true }).
		Show("Delete?")
	if !errors.Is(err, pterm.ErrInterrupted) || !called {
		t.Errorf("err = %v, handler called = %v", err, called)
	}

	if got, want := term.Screen(), "✖ Delete? canceled"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}
}

func TestConfirmContext(t *testing.T) {
	setup(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := pterm.Confirm.WithContext(ctx).WithInput(strings.NewReader("y")).Show()
	if !errors.Is(err, pterm.ErrCanceled) {
		t.Errorf("err = %v, want ErrCanceled", err)
	}
}
