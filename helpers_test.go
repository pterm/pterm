package pterm_test

import (
	"bytes"
	"regexp"
	"testing"

	"atomicgo.dev/color"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/snapshot"
	"github.com/pterm/pterm/internal/vt"
)

// Tests in this package change global state (theme, color level, outputs),
// so they do not call t.Parallel.

// setup pins everything that would make output depend on the machine: the
// theme, the glyphs, the color level and the terminal size. It returns a
// virtual terminal that is both the default and the error output.
func setup(t *testing.T) *vt.Terminal {
	t.Helper()

	term := vt.New(80, 24)

	pterm.SetTheme(pterm.ThemeMacchiato)
	pterm.SetGlyphs(pterm.GlyphsUnicode)
	pterm.EnableOutput()
	pterm.EnableStyling()
	pterm.DisableDebugMessages()
	pterm.SetDefaultOutput(term)
	pterm.SetErrorOutput(term)
	color.SetLevel(color.LevelTrueColor)

	t.Cleanup(func() {
		pterm.SetDefaultOutput(nil)
		pterm.SetErrorOutput(nil)
		pterm.EnableStyling()
		color.ResetLevel()
	})

	return term
}

// snap asserts render's output twice against committed snapshots: styled
// with true color, and plain with colors off.
func snap(t *testing.T, name string, render func() string) {
	t.Helper()

	t.Run(name, func(t *testing.T) {
		setup(t)

		t.Run("styled", func(t *testing.T) {
			color.SetLevel(color.LevelTrueColor)
			snapshot.Assert(t, render())
		})

		t.Run("plain", func(t *testing.T) {
			color.SetLevel(color.LevelNone)
			snapshot.Assert(t, render())
		})
	})
}

// srender adapts a Srender method for snap and fails the test on errors.
func srender(t *testing.T, render func() (string, error)) func() string { //nolint:unused // TODO: used by the render printers, remove this directive when they land.
	t.Helper()

	return func() string {
		s, err := render()
		if err != nil {
			t.Fatalf("Srender: %v", err)
		}

		return s
	}
}

// capture returns what fn prints to a buffer that is not a terminal, the
// way a pipe would receive it: without colors.
func capture(t *testing.T, fn func()) string {
	t.Helper()

	var buf bytes.Buffer

	color.SetLevel(color.LevelNone)
	defer color.SetLevel(color.LevelTrueColor)

	pterm.SetDefaultOutput(&buf)
	pterm.SetErrorOutput(&buf)

	fn()

	return buf.String()
}

// timerRegex matches the muted elapsed time spinners and steps append.
var timerRegex = regexp.MustCompile(` \d+\.\ds`)

// stripTimer removes elapsed times so screens with timers can be compared.
func stripTimer(s string) string {
	return timerRegex.ReplaceAllString(s, "")
}
