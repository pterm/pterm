package pterm_test

import (
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
)

func TestRuleSnapshots(t *testing.T) {
	snap(t, "Plain", func() string { return pterm.Rule.Sprint() })
	snap(t, "Titled", func() string { return pterm.Rule.Sprint("Test results") })
	snap(t, "MutedTitle", func() string {
		return pterm.Rule.WithWidth(40).WithTitleStyle(pterm.ToneMuted.Style()).Sprint("Details")
	})
}

func TestRuleWidths(t *testing.T) {
	setup(t)

	if w := text.Width(pterm.Rule.Sprint("title")); w != 80 {
		t.Errorf("rule is %d wide, want the terminal width 80", w)
	}

	if w := text.Width(pterm.Rule.WithWidth(30).Sprint("title")); w != 30 {
		t.Errorf("rule is %d wide, want 30", w)
	}

	if w := text.Width(pterm.Rule.WithWidth(200).Sprint()); w != 80 {
		t.Errorf("rule is %d wide, want it capped at 80", w)
	}

	t.Setenv("COLUMNS", "")

	out := capture(t, func() { pterm.Rule.Println("Results") })
	if out != "── Results "+strings.Repeat("─", 49)+"\n" {
		t.Errorf("unknown width output = %q", out)
	}

	out = capture(t, func() { pterm.Rule.WithFallbackWidth(10).Println() })
	if out != strings.Repeat("─", 10)+"\n" {
		t.Errorf("fallback output = %q", out)
	}
}

func TestRuleEdgeCases(t *testing.T) {
	setup(t)

	if got := pterm.RemoveColorFromString(pterm.Rule.WithWidth(12).Sprint("a very long title")); got != "── a very… ─" {
		t.Errorf("long title = %q", got)
	}

	if got := pterm.RemoveColorFromString(pterm.Rule.WithWidth(5).Sprint("title")); got != "─────" {
		t.Errorf("tiny rule = %q", got)
	}

	if got := pterm.RemoveColorFromString(pterm.Rule.WithInset(0).WithWidth(10).Sprint("ab")); got != "ab ───────" {
		t.Errorf("no inset = %q", got)
	}

	if got := pterm.RemoveColorFromString(pterm.Rule.WithGlyph("=").WithWidth(4).Sprint()); got != "====" {
		t.Errorf("custom glyph = %q", got)
	}
}

func TestRuleRender(t *testing.T) {
	setup(t)

	s, err := pterm.Rule.WithTitle("x").WithWidth(10).Srender()
	if err != nil || pterm.RemoveColorFromString(s) != "── x ─────" {
		t.Errorf("Srender = %q, %v", s, err)
	}

	out := capture(t, func() { _ = pterm.Rule.WithWidth(5).Render() })
	if out != "─────\n" {
		t.Errorf("Render output = %q", out)
	}
}
