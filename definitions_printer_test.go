package pterm_test

import (
	"errors"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

func TestDefinitionsSnapshots(t *testing.T) {
	snap(t, "Groups", srender(t, pterm.Definitions.
		WithGroup("Commands",
			pterm.Def("doctor", "Check the installation"),
			pterm.Def("disk clean", "Free space used by caches"),
		).
		WithGroup("Flags",
			pterm.Def("-v, --verbose", "Print what is happening"),
			pterm.Def("-h, --help", "Show this help"),
		).Srender))
	snap(t, "Plain", srender(t, pterm.Definitions.WithDefinitions(
		pterm.Def("NO_COLOR", "turn colors off"),
		pterm.Def("PTERM_ASCII", "use ASCII glyphs"),
	).Srender))
}

func TestDefinitionsLook(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		_ = pterm.Definitions.WithGroup("Flags",
			pterm.Def("-v, --verbose", "print what is happening"),
			pterm.Def("-o, --output", "where to write the result"),
		).Render()
	})
	want := "FLAGS\n" +
		"  -v, --verbose  print what is happening\n" +
		"  -o, --output   where to write the result\n"

	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestDefinitionsWrapAndLongTerms(t *testing.T) {
	setup(t)

	term := vt.New(40, 10)
	_ = pterm.Definitions.WithWriter(term).WithDefinitions(
		pterm.Def("-v", "print what is happening while it runs"),
		pterm.Def("--a-very-long-flag-name", "short"),
	).Render()

	want := "-v  print what is happening while it\n" +
		"    runs\n" +
		"--a-very-long-flag-name\n" +
		"    short"

	if term.Screen() != want {
		t.Errorf("screen:\n%s\nwant:\n%s", term.Screen(), want)
	}
}

func TestDefinitionsEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.Definitions.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}
