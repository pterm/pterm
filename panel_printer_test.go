package pterm_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

var dashboard = pterm.Panels{
	{
		{Title: "Uptime", Data: "99.99%\nlast 30 days"},
		{Title: "Errors", Data: "0.02%"},
	},
	{
		{Title: "Region", Data: "eu-central-1"},
	},
}

func TestPanelSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.Panel.WithPanels(pterm.Panels{{
		{Data: "first\npanel"}, {Data: "second"}, {Data: "third\nhas\nlines"},
	}}).Srender))
	snap(t, "Titles", srender(t, pterm.Panel.WithPanels(dashboard).Srender))
	snap(t, "Boxed", srender(t, pterm.Panel.WithBoxed().WithPanels(dashboard).Srender))
	snap(t, "SameColumnWidth", srender(t, pterm.Panel.WithBoxed().WithSameColumnWidth().WithBottomPadding(1).
		WithPanels(dashboard).Srender))
}

func TestPanelBoxedLook(t *testing.T) {
	setup(t)

	out := capture(t, func() { _ = pterm.Panel.WithBoxed().WithPanels(dashboard[:1]).Render() })
	want := "╭─ Uptime ─────╮  ╭─ Errors ─╮\n" +
		"│ 99.99%       │  │ 0.02%    │\n" +
		"│ last 30 days │  │          │\n" +
		"╰──────────────╯  ╰──────────╯\n"

	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestPanelBreaksWideRows(t *testing.T) {
	setup(t)

	term := vt.New(30, 20)
	_ = pterm.Panel.WithWriter(term).WithBoxed().WithRow(
		pterm.PanelItem{Title: "One", Data: "aaaaaaaaaa"},
		pterm.PanelItem{Title: "Two", Data: "bbbbbbbbbb"},
		pterm.PanelItem{Title: "Three", Data: strings.Repeat("c", 50)},
	).Render()

	lines := strings.Split(term.Screen(), "\n")
	// Two items fit next to each other; the third gets its own row and its
	// content is cut to fit.
	if len(lines) != 6 || !strings.Contains(lines[4], "…") {
		t.Errorf("want 2 rows of frames, got:\n%s", term.Screen())
	}

	for _, line := range lines {
		if w := len([]rune(line)); w > 30 {
			t.Errorf("line %q is %d columns wide", line, w)
		}
	}
}

func TestPanelStylesDoNotLeak(t *testing.T) {
	setup(t)

	got, _ := pterm.Panel.WithRow(
		pterm.PanelItem{Data: "\x1b[31mred"},
		pterm.PanelItem{Data: "plain"},
	).Srender()

	if !strings.Contains(got, "red\x1b[0m") {
		t.Errorf("no reset after styled line: %q", got)
	}
}

func TestPanelBuildersDoNotShare(t *testing.T) {
	setup(t)

	base := pterm.Panel.WithRow(pterm.PanelItem{Data: "a"})
	one := base.WithRow(pterm.PanelItem{Data: "b"})
	two := base.WithRow(pterm.PanelItem{Data: "c"})

	if one.Panels[1][0].Data != "b" || two.Panels[1][0].Data != "c" || len(base.Panels) != 1 {
		t.Error("builders share their rows")
	}
}

func TestPanelEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.Panel.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}
