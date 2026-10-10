package pterm_test

import (
	"strings"
	"testing"

	"atomicgo.dev/color"

	"github.com/pterm/pterm"
)

func TestBadgeSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Badge.Sprint("new") })
	snap(t, "Warning", func() string { return pterm.Badge.WithTone(pterm.ToneWarning).Sprint("dry run") })
	snap(t, "Chip", func() string { return pterm.Badge.Chip().WithTone(pterm.ToneSuccess).Sprint("running") })
}

func TestBadgeText(t *testing.T) {
	setup(t)

	cases := []struct {
		p    *pterm.BadgePrinter
		want string
	}{
		{&pterm.Badge, " DRY RUN "},
		{pterm.Badge.Chip(), "[dry run]"},
		{pterm.Badge.Chip().WithUppercase(), "[DRY RUN]"},
		{pterm.Badge.WithUppercase(false).WithPadding(2), "  dry run  "},
		{pterm.Badge.Chip().WithBrackets("<", ">"), "<dry run>"},
		{pterm.Badge.Chip().Solid().WithPadding(0), "DRY RUN"},
	}

	for _, c := range cases {
		if got := pterm.RemoveColorFromString(c.p.Sprint("dry   run")); got != c.want {
			t.Errorf("got %q, want %q", got, c.want)
		}
	}

	if got := pterm.Badge.Sprint(""); got != "" {
		t.Errorf("empty badge = %q", got)
	}
}

func TestBadgeInline(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Badge.Print("x")
		pterm.Print(" text\n")
	})
	if out != " X  text\n" {
		t.Errorf("output = %q", out)
	}
}

func TestBadgeContrast(t *testing.T) {
	setup(t)

	// Every Macchiato tone is light, so badge text is a dark shade of it.
	dark := pterm.Badge.Sprint("x")
	if !strings.Contains(dark, "38;2;37;38;50") {
		t.Errorf("light badge text is not dark: %q", dark)
	}

	pterm.SetTheme(pterm.Theme{Primary: color.RGB(0x1e, 0x1e, 0x2e)})
	defer pterm.SetTheme(pterm.ThemeMacchiato)

	light := pterm.Badge.Sprint("x")
	if strings.Contains(light, "38;2;37;38;50") || !strings.Contains(light, "48;2;30;30;46") {
		t.Errorf("dark badge has wrong colors: %q", light)
	}
}
