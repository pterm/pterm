package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestSectionSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Section.Sprint("Installation") })
	snap(t, "Level2", func() string { return pterm.Section.WithLevel(2).Sprint("Runtime") })
	snap(t, "Level3", func() string { return pterm.Section.WithLevel(3).Sprint("Indirect") })
	snap(t, "Prefix", func() string {
		return pterm.Section.WithPrefix("#").WithLevel(2).WithBottomPadding(1).Sprint("Development")
	})
}

func TestSectionPadding(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Section.Println("One")
		pterm.Section.WithTopPadding(0).WithBottomPadding(2).Println("Two")
	})
	if out != "\nOne\nTwo\n\n\n" {
		t.Errorf("output = %q", out)
	}
}

func TestSectionDeepLevelUsesLastStyle(t *testing.T) {
	setup(t)

	deep := pterm.Section.WithLevel(9).WithTopPadding(0).Sprint("x")
	if want := pterm.Section.WithLevel(3).WithTopPadding(0).Sprint("x"); deep != want {
		t.Errorf("level 9 = %q, want the level 3 look %q", deep, want)
	}

	if got := pterm.Section.WithLevel(0).WithTopPadding(0).Sprint("x"); got != pterm.Section.WithTopPadding(0).Sprint("x") {
		t.Errorf("level 0 = %q, want level 1", got)
	}
}

func TestSectionWrapsUnderPrefix(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "20")

	out := capture(t, func() {
		pterm.Section.WithPrefix("#").WithTopPadding(0).Println("a title that needs two lines")
	})
	if out != "# a title that needs\n  two lines\n" {
		t.Errorf("output = %q", out)
	}
}
