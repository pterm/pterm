package color

import (
	"os"
	"strconv"
	"strings"
)

// Level describes how many colors a terminal can render.
type Level uint8

const (
	// LevelNone means the terminal renders no ANSI colors at all.
	LevelNone Level = iota

	// LevelBasic means the terminal renders the 16 base ANSI colors.
	LevelBasic

	// Level256 means the terminal renders the xterm 256-color palette.
	Level256

	// LevelTrueColor means the terminal renders 24-bit RGB colors.
	LevelTrueColor
)

// String returns a human-readable name for the level.
func (l Level) String() string {
	switch l {
	case LevelNone:
		return "none"
	case LevelBasic:
		return "16"
	case Level256:
		return "256"
	case LevelTrueColor:
		return "truecolor"
	default:
		return "unknown"
	}
}

// ForegroundRGB returns the escape sequence that best approximates the given
// 24-bit foreground color at this level. Terminals without true color support
// get the nearest color of their palette instead of a sequence they cannot
// render. LevelNone returns an empty string.
func (l Level) ForegroundRGB(r, g, b uint8) string {
	switch l {
	case LevelTrueColor:
		return ForegroundRGB(r, g, b)
	case Level256:
		return Foreground256(RGBTo256(r, g, b))
	case LevelBasic:
		return Sequence(strconv.Itoa(int(RGBToBasic(r, g, b))))
	default:
		return ""
	}
}

// BackgroundRGB returns the escape sequence that best approximates the given
// 24-bit background color at this level. See Level.ForegroundRGB.
func (l Level) BackgroundRGB(r, g, b uint8) string {
	// The background code of each base color is its foreground code + 10
	// (30-37 becomes 40-47, 90-97 becomes 100-107).
	const backgroundOffset = 10

	switch l {
	case LevelTrueColor:
		return BackgroundRGB(r, g, b)
	case Level256:
		return Background256(RGBTo256(r, g, b))
	case LevelBasic:
		return Sequence(strconv.Itoa(int(RGBToBasic(r, g, b)) + backgroundOffset))
	default:
		return ""
	}
}

// DetectLevel reports how many colors the current terminal can render.
//
// The verdict is based on well-known environment variables (NO_COLOR,
// FORCE_COLOR, CLICOLOR, CLICOLOR_FORCE, COLORTERM, TERM and several
// terminal-specific ones) and, on Windows, on whether the attached console
// interprets ANSI sequences. The first call on Windows switches the console
// into virtual terminal mode, so ANSI colors also work in classic consoles
// like cmd.exe.
func DetectLevel() Level {
	console, consoleKnown := consoleSupport()

	return detectLevel(os.Getenv, console, consoleKnown)
}

// SupportsANSI reports whether the current terminal renders ANSI escape
// sequences at all. See DetectLevel for how the verdict is reached.
func SupportsANSI() bool {
	return DetectLevel() != LevelNone
}

// SupportsTrueColor reports whether the current terminal renders 24-bit (RGB)
// colors. See DetectLevel for how the verdict is reached.
func SupportsTrueColor() bool {
	return DetectLevel() == LevelTrueColor
}

// detectLevel is the testable core of DetectLevel: all inputs are explicit.
// console and consoleKnown carry the platform console verdict (see
// consoleSupport).
func detectLevel(getenv func(string) string, console Level, consoleKnown bool) Level {
	// FORCE_COLOR both forces colors on (optionally at a specific level) and,
	// set to 0, forces them off. It wins over every other signal.
	if force := getenv("FORCE_COLOR"); force != "" {
		switch strings.ToLower(force) {
		case "0", "false", "none":
			return LevelNone
		case "2":
			return Level256
		case "3":
			return LevelTrueColor
		default: // "1", "true" and any other value force at least basic colors.
			return max(envLevel(getenv, console, consoleKnown), LevelBasic)
		}
	}

	// CLICOLOR_FORCE (BSD convention) forces colors on but names no level.
	if force := getenv("CLICOLOR_FORCE"); force != "" && force != "0" {
		return max(envLevel(getenv, console, consoleKnown), LevelBasic)
	}

	// NO_COLOR (https://no-color.org) and CLICOLOR=0 opt out of colors.
	if getenv("NO_COLOR") != "" || getenv("CLICOLOR") == "0" {
		return LevelNone
	}

	return envLevel(getenv, console, consoleKnown)
}

// envLevel determines the color level from descriptive (non-forcing)
// environment variables and the platform console verdict.
func envLevel(getenv func(string) string, console Level, consoleKnown bool) Level {
	colorTerm := strings.ToLower(getenv("COLORTERM"))
	term := strings.ToLower(getenv("TERM"))

	if strings.Contains(colorTerm, "truecolor") || strings.Contains(colorTerm, "24bit") {
		return LevelTrueColor
	}

	if strings.Contains(term, "truecolor") || strings.Contains(term, "24bit") || strings.Contains(term, "direct") {
		return LevelTrueColor
	}

	if term == "dumb" {
		return LevelNone
	}

	// Terminals that support true color but do not always set COLORTERM:
	// Windows Terminal, ConEmu (with ANSI enabled) and JetBrains' terminal.
	if getenv("WT_SESSION") != "" || getenv("ConEmuANSI") == "ON" || getenv("TERMINAL_EMULATOR") == "JetBrains-JediTerm" {
		return LevelTrueColor
	}

	switch getenv("TERM_PROGRAM") {
	case "iTerm.app", "Hyper", "vscode", "ghostty", "WezTerm", "Apple_Terminal":
		return LevelTrueColor
	}

	if strings.Contains(term, "256color") {
		return Level256
	}

	// ANSICON translates ANSI sequences for legacy Windows consoles.
	if getenv("ANSICON") != "" {
		return LevelBasic
	}

	// The platform console gave a definitive verdict: on Windows this is
	// LevelTrueColor when the console interprets ANSI sequences (every
	// supported Windows version renders 24-bit color then) and LevelNone when
	// it cannot (legacy consoles, e.g. cmd.exe on Windows 7/8).
	if consoleKnown {
		return console
	}

	// Nothing indicates otherwise: assume a basic ANSI terminal. Whether
	// colors are printed at all is the caller's decision (pterm gates on
	// pterm.PrintColor), so this only affects how RGB colors degrade.
	return LevelBasic
}
