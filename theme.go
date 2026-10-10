package pterm

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/screen"
)

// Theme holds the nine base colors every printer builds on. Printers never
// store colors of their own; they refer to these slots by Tone, so a new
// theme restyles everything.
type Theme struct {
	// Primary marks titles, the active item and anything that is "pterm".
	Primary color.Color
	// Accent marks values and highlights that stand next to Primary.
	Accent color.Color
	// Success marks things that worked.
	Success color.Color
	// Warning marks things to double check.
	Warning color.Color
	// Error marks failures.
	Error color.Color
	// Info marks neutral notes.
	Info color.Color
	// Text is the color of plain text.
	Text color.Color
	// Muted is for hints, metadata and anything secondary.
	Muted color.Color
	// Border is for lines, frames and tracks that structure output.
	Border color.Color
}

// ThemeMacchiato is the Catppuccin Macchiato theme, made for dark terminals.
var ThemeMacchiato = Theme{
	Primary: color.RGB(0xb7, 0xbd, 0xf8), // lavender
	Accent:  color.RGB(0x8b, 0xd5, 0xca), // teal
	Success: color.RGB(0xa6, 0xda, 0x95), // green
	Warning: color.RGB(0xee, 0xd4, 0x9f), // yellow
	Error:   color.RGB(0xed, 0x87, 0x96), // red
	Info:    color.RGB(0x8a, 0xad, 0xf4), // blue
	Text:    color.RGB(0xca, 0xd3, 0xf5), // text
	Muted:   color.RGB(0x80, 0x87, 0xa2), // overlay1
	Border:  color.RGB(0x5b, 0x60, 0x78), // surface2
}

// ThemeLatte is the Catppuccin Latte theme, made for light terminals.
var ThemeLatte = Theme{
	Primary: color.RGB(0x72, 0x87, 0xfd), // lavender
	Accent:  color.RGB(0x17, 0x92, 0x99), // teal
	Success: color.RGB(0x40, 0xa0, 0x2b), // green
	Warning: color.RGB(0xdf, 0x8e, 0x1d), // yellow
	Error:   color.RGB(0xd2, 0x0f, 0x39), // red
	Info:    color.RGB(0x1e, 0x66, 0xf5), // blue
	Text:    color.RGB(0x4c, 0x4f, 0x69), // text
	Muted:   color.RGB(0x8c, 0x8f, 0xa1), // overlay1
	Border:  color.RGB(0xac, 0xb0, 0xbe), // surface2
}

// ThemeANSI uses the 16 basic terminal colors, so output follows whatever
// palette the user configured in their terminal.
var ThemeANSI = Theme{
	Primary: color.Magenta,
	Accent:  color.Cyan,
	Success: color.Green,
	Warning: color.Yellow,
	Error:   color.Red,
	Info:    color.Blue,
	Text:    color.None,
	Muted:   color.BrightBlack,
	Border:  color.BrightBlack,
}

// themes are the presets PTERM_THEME can name.
var themes = map[string]Theme{
	"macchiato": ThemeMacchiato,
	"latte":     ThemeLatte,
	"ansi":      ThemeANSI,
}

// slots maps each slot name to its field, for PTERM_THEME_<SLOT>.
var slots = map[string]func(*Theme) *color.Color{
	"PRIMARY": func(t *Theme) *color.Color { return &t.Primary },
	"ACCENT":  func(t *Theme) *color.Color { return &t.Accent },
	"SUCCESS": func(t *Theme) *color.Color { return &t.Success },
	"WARNING": func(t *Theme) *color.Color { return &t.Warning },
	"ERROR":   func(t *Theme) *color.Color { return &t.Error },
	"INFO":    func(t *Theme) *color.Color { return &t.Info },
	"TEXT":    func(t *Theme) *color.Color { return &t.Text },
	"MUTED":   func(t *Theme) *color.Color { return &t.Muted },
	"BORDER":  func(t *Theme) *color.Color { return &t.Border },
}

var (
	activeTheme atomic.Pointer[Theme]
	themeOnce   sync.Once

	// backgroundTimeout bounds the OSC 11 query on first use.
	backgroundTimeout = 50 * time.Millisecond
)

// SetTheme makes t the theme of every printer. Environment variables do not
// apply to a theme set this way; use ThemeFromEnv to layer them on top.
func SetTheme(t Theme) {
	themeOnce.Do(func() {})
	activeTheme.Store(&t)
}

// CurrentTheme returns the active theme. On first use it is picked from the
// environment, see ThemeFromEnv.
func CurrentTheme() Theme {
	themeOnce.Do(func() {
		t := ThemeFromEnv()
		activeTheme.Store(&t)
	})

	return *activeTheme.Load()
}

// ThemeFromEnv builds the theme the environment asks for.
//
// PTERM_THEME picks a preset: "macchiato", "latte", "ansi" or "auto" (the
// default). "auto" asks the terminal for its background color and picks
// Latte for light backgrounds and Macchiato otherwise. The question is only
// asked when standard input and output are both terminals and CI is unset,
// takes at most 50 milliseconds, and is asked once per program.
//
// PTERM_THEME_<SLOT> then overrides single colors, for example
// PTERM_THEME_PRIMARY=#f5a97f. Values can be hex colors, basic color names
// such as "brightblue", or 256-color indexes. Invalid values are ignored.
func ThemeFromEnv() Theme {
	name := strings.ToLower(strings.TrimSpace(os.Getenv("PTERM_THEME")))

	t, ok := themes[name]
	if !ok {
		t = autoTheme()
	}

	for slot, field := range slots {
		v, set := os.LookupEnv("PTERM_THEME_" + slot)
		if !set {
			continue
		}

		if c, err := color.Parse(v); err == nil {
			*field(&t) = c
		}
	}

	return t
}

func autoTheme() Theme {
	if os.Getenv("CI") != "" || !screen.IsTerminal(os.Stdout) || !screen.IsTerminal(os.Stdin) {
		return ThemeMacchiato
	}

	if bg, ok := color.Background(backgroundTimeout); ok && !bg.IsDark() {
		return ThemeLatte
	}

	return ThemeMacchiato
}

// Tone is a color for a part of a printer: one of the theme slots, or a
// fixed color made with Fixed. Slots resolve against the active theme when
// the printer renders, so they follow SetTheme.
type Tone struct {
	slot  uint8
	fixed color.Color
}

const (
	slotNone uint8 = iota
	slotPrimary
	slotAccent
	slotSuccess
	slotWarning
	slotError
	slotInfo
	slotText
	slotMuted
	slotBorder
	slotFixed
)

// The theme tones. ToneDefault leaves the terminal's own color alone.
var (
	ToneDefault = Tone{}
	TonePrimary = Tone{slot: slotPrimary}
	ToneAccent  = Tone{slot: slotAccent}
	ToneSuccess = Tone{slot: slotSuccess}
	ToneWarning = Tone{slot: slotWarning}
	ToneError   = Tone{slot: slotError}
	ToneInfo    = Tone{slot: slotInfo}
	ToneText    = Tone{slot: slotText}
	ToneMuted   = Tone{slot: slotMuted}
	ToneBorder  = Tone{slot: slotBorder}
)

// Fixed returns a Tone that is always c, whatever the theme.
func Fixed(c color.Color) Tone {
	return Tone{slot: slotFixed, fixed: c}
}

// In returns the color of t in theme th.
func (t Tone) In(th Theme) color.Color {
	switch t.slot {
	case slotPrimary:
		return th.Primary
	case slotAccent:
		return th.Accent
	case slotSuccess:
		return th.Success
	case slotWarning:
		return th.Warning
	case slotError:
		return th.Error
	case slotInfo:
		return th.Info
	case slotText:
		return th.Text
	case slotMuted:
		return th.Muted
	case slotBorder:
		return th.Border
	case slotFixed:
		return t.fixed
	default:
		return color.None
	}
}

// Color returns the color of t in the active theme.
func (t Tone) Color() color.Color {
	return t.In(CurrentTheme())
}

// String names the tone: "primary", "muted"..., or the fixed color.
func (t Tone) String() string {
	names := [...]string{"default", "primary", "accent", "success", "warning", "error", "info", "text", "muted", "border"}
	if t.slot == slotFixed {
		return t.fixed.String()
	}

	return names[t.slot]
}

// Style returns a Style with t as foreground.
func (t Tone) Style() Style {
	return Style{Fg: t}
}

// Bold returns a bold Style with t as foreground.
func (t Tone) Bold() Style {
	return Style{Fg: t, Attrs: color.Bold}
}

// Sprint colors the text with t for the default output.
func (t Tone) Sprint(a ...any) string {
	return t.Style().Sprint(a...)
}

// Sprintf formats and colors the text with t for the default output.
func (t Tone) Sprintf(format string, a ...any) string {
	return t.Style().Sprintf(format, a...)
}

// Style is a foreground tone, a background tone and text attributes.
// The zero value leaves text unchanged.
type Style struct {
	Fg    Tone
	Bg    Tone
	Attrs color.Attr
}

// NewStyle returns a Style with fg as foreground and the given attributes.
func NewStyle(fg Tone, attrs ...color.Attr) Style {
	s := Style{Fg: fg}
	for _, a := range attrs {
		s.Attrs |= a
	}

	return s
}

// On returns a copy of s with bg as background.
func (s Style) On(bg Tone) Style {
	s.Bg = bg
	return s
}

// With returns a copy of s with the given attributes added.
func (s Style) With(attrs ...color.Attr) Style {
	for _, a := range attrs {
		s.Attrs |= a
	}

	return s
}

// Render styles text at level l with theme th.
func (s Style) Render(l color.Level, th Theme, text string) string {
	return color.Style{Foreground: s.Fg.In(th), Background: s.Bg.In(th), Attrs: s.Attrs}.Render(l, text)
}

// Sprint styles the text for the default output.
func (s Style) Sprint(a ...any) string {
	return s.Render(ColorLevel(DefaultOutput()), CurrentTheme(), fmt.Sprint(a...))
}

// Sprintf formats and styles the text for the default output.
func (s Style) Sprintf(format string, a ...any) string {
	return s.Sprint(fmt.Sprintf(format, a...))
}
