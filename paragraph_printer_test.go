package pterm_test

import (
	"strings"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/text"
)

const lorem = "pterm prints beautiful terminal output with one line of code. Every printer follows one theme, so tables, spinners and prompts look like parts of one tool."

func TestParagraphSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Paragraph.Sprint(lorem + "\n\n" + lorem) })
	snap(t, "MaxWidth", func() string { return pterm.Paragraph.WithMaxWidth(30).Sprint(lorem) })
}

func TestParagraphCollapsesWhitespace(t *testing.T) {
	setup(t)

	got := pterm.RemoveColorFromString(pterm.Paragraph.Sprint("one   two\nthree\n\n\n\nfour"))
	if got != "one two three\n\nfour" {
		t.Errorf("got %q", got)
	}
}

func TestParagraphWidth(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	// With an unknown terminal width, MaxWidth still applies.
	out := capture(t, func() { pterm.Paragraph.WithMaxWidth(20).Println(lorem) })
	for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
		if text.Width(line) > 20 {
			t.Errorf("line %q is wider than 20 columns", line)
		}
	}

	// Without MaxWidth and terminal width, nothing wraps.
	out = capture(t, func() { pterm.Paragraph.WithMaxWidth(0).Println(lorem) })
	if out != lorem+"\n" {
		t.Errorf("output = %q", out)
	}
}

func TestParagraphEmpty(t *testing.T) {
	setup(t)

	if got := pterm.Paragraph.Sprint("  \n\n "); got != "" {
		t.Errorf("got %q", got)
	}
}
