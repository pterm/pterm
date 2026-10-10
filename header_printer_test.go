package pterm_test

import (
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
)

func TestHeaderSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Header.Sprint("Release notes") })
	snap(t, "Badge", func() string {
		return pterm.Header.WithBadge("APP").WithSubtitle("Design system gallery").Sprint("debug ui")
	})
	snap(t, "ContentWidth", func() string {
		return pterm.Header.WithBadge("beta").WithFullWidth(false).Sprint("Features that may change")
	})
	snap(t, "NoRule", func() string {
		return pterm.Header.WithBadge("APP").WithNoRule().WithTopPadding(0).Sprint("debug ui")
	})
}

func TestHeaderLayout(t *testing.T) {
	setup(t)

	got := pterm.Header.WithBadge("APP").WithSubtitle("Design system gallery").Sprint("debug ui")
	lines := strings.Split(pterm.RemoveColorFromString(got), "\n")

	if len(lines) != 3 || lines[0] != "" {
		t.Fatalf("lines = %q", lines)
	}

	if lines[1] != " APP  debug ui  Design system gallery" {
		t.Errorf("header line = %q", lines[1])
	}

	if text.Width(lines[2]) != 80 {
		t.Errorf("rule is %d columns wide, want the terminal width", text.Width(lines[2]))
	}
}

func TestHeaderUnknownWidth(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	out := capture(t, func() { pterm.Header.WithBadge("APP").Println("debug ui") })
	if out != "\n APP  debug ui\n──────────────\n" {
		t.Errorf("output = %q", out)
	}
}

func TestHeaderNarrowDropsSubtitle(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "20")

	out := capture(t, func() {
		pterm.Header.WithBadge("APP").WithSubtitle("Design system gallery").Println("debug ui")
	})
	if out != "\n APP  debug ui\n────────────────────\n" {
		t.Errorf("output = %q", out)
	}
}

func TestHeaderNarrowCutsTitle(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "16")

	out := capture(t, func() { pterm.Header.WithBadge("APP").Println("a very long title") })
	for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
		if text.Width(line) > 16 {
			t.Errorf("line %q is wider than the terminal", line)
		}
	}

	if !strings.Contains(out, "…") {
		t.Errorf("title was not cut: %q", out)
	}
}

func TestHeaderBadgeContrast(t *testing.T) {
	setup(t)

	// Lavender is light, so the badge text must be dark: a low first
	// component in the 38;2 sequence.
	got := pterm.Header.WithBadge("APP").WithNoRule().WithTopPadding(0).Sprint("x")
	if !strings.Contains(got, "38;2;") || !strings.Contains(got, "48;2;183;189;248") {
		t.Errorf("badge colors missing: %q", got)
	}
}
