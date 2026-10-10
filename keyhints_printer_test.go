package pterm_test

import (
	"errors"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
)

var hints = []pterm.KeyHint{
	{Key: "space", Desc: "toggle"},
	{Key: "a", Desc: "all"},
	{Key: "enter", Desc: "confirm"},
}

func TestKeyHintsSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.KeyHints.WithHints(hints...).Srender))
	snap(t, "WithInfo", srender(t, pterm.KeyHints.WithHint("q", "quit").WithHint("", "every 5.0s").Srender))
}

func TestKeyHintsPlain(t *testing.T) {
	setup(t)

	out := capture(t, func() { _ = pterm.KeyHints.WithHints(hints...).Render() })
	if out != "[space] toggle · [a] all · [enter] confirm\n" {
		t.Errorf("output = %q", out)
	}

	out = capture(t, func() {
		_ = pterm.KeyHints.WithHints(pterm.KeyHint{Key: "esc"}, pterm.KeyHint{Desc: "info"}).
			WithSeparator(" | ").WithBrackets("<", ">").Render()
	})
	if out != "<esc> | info\n" {
		t.Errorf("custom output = %q", out)
	}
}

func TestKeyHintsEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.KeyHints.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("Srender error = %v, want ErrEmpty", err)
	}

	if err := pterm.KeyHints.Render(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("Render error = %v, want ErrEmpty", err)
	}
}

func TestKeyHintsNarrow(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "20")

	out := capture(t, func() { _ = pterm.KeyHints.WithHints(hints...).Render() })
	if out != "[space] toggle · [a…\n" || text.Width(out[:len(out)-1]) != 20 {
		t.Errorf("output = %q", out)
	}
}

func TestKeyHintsWithHintDoesNotShare(t *testing.T) {
	setup(t)

	base := pterm.KeyHints.WithHints(make([]pterm.KeyHint, 0, 4)...)
	a := base.WithHint("a", "one")
	b := base.WithHint("b", "two")

	if s, _ := a.Srender(); pterm.RemoveColorFromString(s) != "[a] one" {
		t.Errorf("a = %q", s)
	}

	if s, _ := b.Srender(); pterm.RemoveColorFromString(s) != "[b] two" {
		t.Errorf("b = %q", s)
	}
}
