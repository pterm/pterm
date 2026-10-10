package pterm_test

import (
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
)

func TestBoxSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Box.Sprint("3 files changed\n128 insertions, 41 deletions") })
	snap(t, "Title", func() string { return pterm.Box.WithTitle("Result").Sprint("3 files changed") })
	snap(t, "TitleCenter", func() string {
		return pterm.Box.WithTitle("Result").WithTitlePosition(pterm.BoxTitleTopCenter).WithWidth(30).Sprint("centered")
	})
	snap(t, "TitleBottomRight", func() string {
		return pterm.Box.WithTitle("Result").WithTitlePosition(pterm.BoxTitleBottomRight).WithWidth(30).Sprint("bottom")
	})
	snap(t, "Padding", func() string { return pterm.Box.WithPadding(1).Sprint("padded") })
	snap(t, "Nested", func() string {
		return pterm.Box.WithTitle("Outer").Sprint(pterm.Box.WithTitle("Inner").Sprint("content"))
	})
	snap(t, "Wrapped", func() string {
		return pterm.Box.Sprint(strings.Repeat("A box wraps text that is wider than the terminal. ", 3))
	})
}

func TestBoxPlainLayout(t *testing.T) {
	setup(t)

	out := capture(t, func() { pterm.Box.WithTitle("Result").Println("ok") })

	want := "╭─ Result ─╮\n│ ok       │\n╰──────────╯\n"
	if out != want {
		t.Errorf("output =\n%s\nwant\n%s", out, want)
	}
}

func TestBoxLinesAreEqualWidth(t *testing.T) {
	setup(t)

	cases := []string{
		"short",
		"wide 中文 characters\nand a second line",
		pterm.Text.WithTone(pterm.ToneAccent).Sprint("styled") + " text",
		strings.Repeat("word ", 40),
	}

	for _, c := range cases {
		lines := strings.Split(pterm.Box.WithTitle("t").Sprint(c), "\n")
		for _, line := range lines {
			if text.Width(line) != text.Width(lines[0]) {
				t.Errorf("%q: line %q has width %d, want %d", c, line, text.Width(line), text.Width(lines[0]))
			}

			if text.Width(line) > 80 {
				t.Errorf("%q: line %q is wider than the terminal", c, line)
			}
		}
	}
}

func TestBoxCustomGlyphs(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Box.WithGlyphs(pterm.BoxGlyphs{TopLeft: "+", TopRight: "+", BottomLeft: "+", BottomRight: "+"}).Println("x")
	})
	if out != "+───+\n│ x │\n+───+\n" {
		t.Errorf("output = %q", out)
	}
}

func TestBoxLongTitleIsCut(t *testing.T) {
	setup(t)

	got := pterm.RemoveColorFromString(pterm.Box.WithTitle(strings.Repeat("title ", 30)).Sprint("x"))
	for line := range strings.SplitSeq(got, "\n") {
		if text.Width(line) > 80 {
			t.Errorf("line %q is wider than the terminal", line)
		}
	}

	if !strings.Contains(got, "…") {
		t.Errorf("long title was not cut:\n%s", got)
	}
}
