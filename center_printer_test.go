package pterm_test

import (
	"testing"

	"github.com/pterm/pterm"
)

func TestCenterSnapshots(t *testing.T) {
	snap(t, "Block", func() string { return pterm.Center.Sprint("short\na longer line") })
	snap(t, "EachLine", func() string { return pterm.Center.WithCenterEachLineSeparately().Sprint("short\na longer line") })
}

func TestCenterWidth(t *testing.T) {
	setup(t)

	if got := pterm.Center.WithWidth(10).Sprint("ab"); got != "    ab" {
		t.Errorf("got %q", got)
	}

	if got := pterm.Center.WithWidth(4).Sprint("too wide"); got != "too wide" {
		t.Errorf("text wider than the width changed: %q", got)
	}
}

func TestCenterUnknownWidth(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	out := capture(t, func() { pterm.Center.Println("as is") })
	if out != "as is\n" {
		t.Errorf("output = %q", out)
	}
}
