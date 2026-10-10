package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// RulePrinter prints a horizontal line across the terminal, with an
// optional title near its start:
//
//	── Results ─────────────────────────────
//
// The text passed to Sprint and Print is the title. Render prints the
// rule with the Title field.
type RulePrinter struct {
	// Title is shown inside the rule by Render and Srender.
	Title string
	// TitleStyle styles the title.
	TitleStyle Style
	// Tone colors the line.
	Tone Tone
	// Glyph is the character the line is drawn with. Empty uses the rule
	// glyph of the glyph set.
	Glyph string
	// Width is the width of the rule in columns. Zero spans the terminal,
	// or FallbackWidth when the width of the terminal is unknown.
	Width int
	// FallbackWidth is the width used when Width is zero and the terminal
	// width is unknown.
	FallbackWidth int
	// Inset is the number of line glyphs before the title.
	Inset int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Rule is the default rule.
var Rule = RulePrinter{
	TitleStyle:    ToneText.Style(),
	Tone:          ToneBorder,
	FallbackWidth: 60,
	Inset:         2,
}

// WithTitle returns a copy of p with title t, for Render.
func (p RulePrinter) WithTitle(t string) *RulePrinter {
	p.Title = t
	return &p
}

// WithTitleStyle returns a copy of p that styles the title with s.
func (p RulePrinter) WithTitleStyle(s Style) *RulePrinter {
	p.TitleStyle = s
	return &p
}

// WithTone returns a copy of p whose line is in tone t.
func (p RulePrinter) WithTone(t Tone) *RulePrinter {
	p.Tone = t
	return &p
}

// WithGlyph returns a copy of p that draws the line with g.
func (p RulePrinter) WithGlyph(g string) *RulePrinter {
	p.Glyph = g
	return &p
}

// WithWidth returns a copy of p that is w columns wide. Zero spans the
// terminal.
func (p RulePrinter) WithWidth(w int) *RulePrinter {
	p.Width = max(w, 0)
	return &p
}

// WithFallbackWidth returns a copy of p that is w columns wide when the
// terminal width is unknown.
func (p RulePrinter) WithFallbackWidth(w int) *RulePrinter {
	p.FallbackWidth = max(w, 0)
	return &p
}

// WithInset returns a copy of p with n line glyphs before the title.
func (p RulePrinter) WithInset(n int) *RulePrinter {
	p.Inset = max(n, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p RulePrinter) WithWriter(w io.Writer) *RulePrinter {
	p.Writer = w
	return &p
}

func (p RulePrinter) out() io.Writer { return orDefault(p.Writer) }

func (p RulePrinter) width(c canvas) int {
	switch {
	case p.Width > 0 && c.width > 0:
		return min(p.Width, c.width)
	case p.Width > 0:
		return p.Width
	case c.width > 0:
		return c.width
	default:
		return p.FallbackWidth
	}
}

func (p RulePrinter) render(c canvas, s string) string {
	glyph := p.Glyph
	if glyph == "" {
		glyph = c.g.Rule
	}

	return rule(c, p.width(c), glyph, p.Tone, p.Inset, c.style(p.TitleStyle, strings.Join(strings.Fields(s), " ")))
}

// rule draws a line of width columns in tone t. A non-empty title is set
// in after inset glyphs, with a space on each side.
func rule(c canvas, width int, glyph string, t Tone, inset int, title string) string {
	gw := max(text.Width(glyph), 1)
	line := func(cols int) string {
		return c.paint(t, strings.Repeat(glyph, max(cols, 0)/gw))
	}

	// Keep at least one glyph after the title, so the rule reads as a rule.
	head := min(inset*gw, max(width-2, 0))
	if title == "" || width-head-2-gw < 1 {
		return line(width)
	}

	if head == 0 {
		title = c.cut(title, width-1-gw)
		return title + " " + line(width-text.Width(title)-1)
	}

	title = c.cut(title, width-head-2-gw)

	return line(head) + " " + title + " " + line(width-head-text.Width(title)-2)
}

// Srender returns the rule with the Title field.
func (p RulePrinter) Srender() (string, error) {
	return p.Sprint(p.Title), nil
}

// Render prints the rule with the Title field.
func (p RulePrinter) Render() error {
	p.Println(p.Title)
	return nil
}

// Sprint returns the rule with the text as title.
func (p RulePrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p RulePrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the title and returns the rule.
func (p RulePrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p RulePrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the rule with the text as title. A rule is a whole line,
// so Print and Println print the same.
func (p RulePrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the rule with the text as title.
func (p RulePrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats the title and prints the rule.
func (p RulePrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats the title and prints the rule.
func (p RulePrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints a rule titled with every non-nil error in a.
func (p RulePrinter) PrintOnError(a ...any) { printOnError(p, a...) }
