package pterm

import (
	"fmt"
	"io"
	"strings"
)

// CalloutPrinter prints a block of text that must stand out, with a colored
// bar on its left and an optional title:
//
//	┃ ▲ Heads up
//	┃ Nothing here can be undone.
type CalloutPrinter struct {
	// Tone colors the bar, the glyph and the title.
	Tone Tone
	// Glyph is drawn before the title. Empty draws no glyph.
	Glyph string
	// Title is the first line, in the printer's tone.
	Title string
	// Hint is a muted last line, such as the cause of an error.
	Hint string
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Callout is the default callout, in the Info tone.
var Callout = CalloutPrinter{Tone: ToneInfo}

// WithTone returns a copy of p that uses tone t.
func (p CalloutPrinter) WithTone(t Tone) *CalloutPrinter {
	p.Tone = t
	return &p
}

// WithGlyph returns a copy of p that draws g before the title.
func (p CalloutPrinter) WithGlyph(g string) *CalloutPrinter {
	p.Glyph = g
	return &p
}

// WithTitle returns a copy of p with title t.
func (p CalloutPrinter) WithTitle(t string) *CalloutPrinter {
	p.Title = t
	return &p
}

// WithHint returns a copy of p that ends with the muted line h.
func (p CalloutPrinter) WithHint(h string) *CalloutPrinter {
	p.Hint = h
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p CalloutPrinter) WithWriter(w io.Writer) *CalloutPrinter {
	p.Writer = w
	return &p
}

// Note returns a copy of p styled as a note: Info tone, info glyph, and
// title t.
func (p CalloutPrinter) Note(t string) *CalloutPrinter {
	return p.preset(ToneInfo, CurrentGlyphs().Info, t)
}

// Tip returns a copy of p styled as a success callout with title t.
func (p CalloutPrinter) Tip(t string) *CalloutPrinter {
	return p.preset(ToneSuccess, CurrentGlyphs().Success, t)
}

// Warn returns a copy of p styled as a warning with title t.
func (p CalloutPrinter) Warn(t string) *CalloutPrinter {
	return p.preset(ToneWarning, CurrentGlyphs().Warning, t)
}

// Danger returns a copy of p styled as an error with title t.
func (p CalloutPrinter) Danger(t string) *CalloutPrinter {
	return p.preset(ToneError, CurrentGlyphs().Error, t)
}

func (p CalloutPrinter) preset(t Tone, glyph, title string) *CalloutPrinter {
	p.Tone, p.Glyph, p.Title = t, glyph, title
	return &p
}

func (p CalloutPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p CalloutPrinter) render(c canvas, s string) string {
	bar := c.paint(p.Tone, c.g.Bar) + " "

	var lines []string

	if p.Title != "" {
		title := p.Title
		if p.Glyph != "" {
			title = p.Glyph + " " + title
		}

		lines = append(lines, c.paint(p.Tone, title))
	}

	if s = strings.TrimRight(s, "\n"); s != "" {
		lines = append(lines, c.wrap(c.paint(ToneText, s), 2)...)
	}

	if p.Hint != "" {
		lines = append(lines, c.wrap(c.paint(ToneMuted, p.Hint), 2)...)
	}

	for i, line := range lines {
		lines[i] = bar + line
	}

	return strings.Join(lines, "\n")
}

// Sprint returns the callout with the text as body.
func (p CalloutPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p CalloutPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the body and returns the callout.
func (p CalloutPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p CalloutPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the callout with the text as body.
func (p CalloutPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the callout with the text as body.
func (p CalloutPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats the body and prints the callout.
func (p CalloutPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats the body and prints the callout.
func (p CalloutPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints a callout for every non-nil error in a.
func (p CalloutPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
