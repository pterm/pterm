// Package pterm prints polished terminal output with one line of code.
//
//	pterm.Success.Println("Deployed")
//	pterm.Table.WithData(data).Render()
//	spinner, _ := pterm.Spinner.Start("Downloading")
//
// Every printer starts from a ready-to-use package value (Info, Table,
// Spinner, Select...) and is customized with With* methods, which return a
// modified copy and never change the value they were called on.
//
// Colors come from the active Theme, which holds nine base colors. Printers
// refer to them by Tone, so changing the theme restyles every printer. The
// default theme follows the terminal background (Catppuccin Macchiato on
// dark terminals, Latte on light ones) and can be changed with SetTheme or
// the PTERM_THEME and PTERM_THEME_<SLOT> environment variables.
//
// All output goes through one renderer per writer, so printers can be used
// from many goroutines at once: normal output appears above running live
// printers instead of tearing them.
package pterm

import (
	"io"
	"os"
	"sync/atomic"

	"atomicgo.dev/color"
)

// settings are the global switches. They are replaced as a whole, never
// modified in place, so readers never see a half-written state.
type settings struct {
	output bool
	color  bool
	raw    bool
	debug  bool
	stdout io.Writer
	stderr io.Writer
}

var current atomic.Pointer[settings]

func init() {
	current.Store(&settings{output: true, color: true})
}

func load() settings {
	return *current.Load()
}

func update(change func(s *settings)) {
	for {
		old := current.Load()
		next := *old
		change(&next)

		if current.CompareAndSwap(old, &next) {
			return
		}
	}
}

// EnableOutput turns pterm output back on after DisableOutput.
func EnableOutput() { update(func(s *settings) { s.output = true }) }

// DisableOutput silences every printer, for example for a --quiet flag.
// Sprint methods still return their text.
func DisableOutput() { update(func(s *settings) { s.output = false }) }

// OutputEnabled reports whether pterm prints anything.
func OutputEnabled() bool { return load().output }

// EnableColor turns colors back on after DisableColor. Whether a writer
// shows colors still depends on what it supports.
func EnableColor() { update(func(s *settings) { s.color = true }) }

// DisableColor prints everything without colors and text attributes.
func DisableColor() { update(func(s *settings) { s.color = false }) }

// ColorEnabled reports whether pterm prints colors where the writer supports
// them.
func ColorEnabled() bool { return load().color }

// EnableStyling undoes DisableStyling.
func EnableStyling() { update(func(s *settings) { s.raw = false; s.color = true }) }

// DisableStyling prints plain text only: no colors, and live printers print
// their final state instead of animating. Use it when writing to files.
func DisableStyling() { update(func(s *settings) { s.raw = true; s.color = false }) }

// StylingEnabled reports whether styling is on.
func StylingEnabled() bool { return !load().raw }

// EnableDebugMessages makes the Debug printer and debug log lines print.
func EnableDebugMessages() { update(func(s *settings) { s.debug = true }) }

// DisableDebugMessages hides the Debug printer and debug log lines again.
func DisableDebugMessages() { update(func(s *settings) { s.debug = false }) }

// DebugMessagesEnabled reports whether debug output prints.
func DebugMessagesEnabled() bool { return load().debug }

// SetDefaultOutput sets the writer for normal output. The default is
// os.Stdout. Nil restores it.
func SetDefaultOutput(w io.Writer) { update(func(s *settings) { s.stdout = w }) }

// SetErrorOutput sets the writer for warnings, errors and log lines. The
// default is os.Stderr. Nil restores it.
func SetErrorOutput(w io.Writer) { update(func(s *settings) { s.stderr = w }) }

// DefaultOutput returns the writer for normal output.
func DefaultOutput() io.Writer {
	if w := load().stdout; w != nil {
		return w
	}

	return os.Stdout
}

// ErrorOutput returns the writer for warnings, errors and log lines.
func ErrorOutput() io.Writer {
	if w := load().stderr; w != nil {
		return w
	}

	return os.Stderr
}

// ColorLevel returns how many colors pterm uses for w: none when colors are
// disabled, otherwise what color.Detect reports for w. Override detection
// with color.SetLevel.
func ColorLevel(w io.Writer) color.Level {
	if !load().color {
		return color.LevelNone
	}

	return color.Detect(w)
}
