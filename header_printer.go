package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/text"
)

// HeaderPrinter prints the head of a program's output: a badge, a bold
// title and a muted subtitle on one line, with a rule under them and one
// blank line before.
//
//	 APP  debug ui  Design system gallery
//	──────────────────────────────────────
//
// The text passed to Sprint and Print is the title. Without badge text,
// the title itself is shown in the badge.
type HeaderPrinter struct {
	// Badge is the text of the badge, usually the name of the program.
	Badge string
	// Subtitle follows the title, muted.
	Subtitle string
	// BadgeTone is the background of the badge.
	BadgeTone Tone
	// TitleStyle styles the title next to the badge.
	TitleStyle Style
	// SubtitleStyle styles the subtitle.
	SubtitleStyle Style
	// RuleTone colors the rule under the header.
	RuleTone Tone
	// NoRule leaves out the rule.
	NoRule bool
	// FullWidth makes the rule span the terminal. Otherwise, and when the
	// terminal width is unknown, it is as wide as the header line.
	FullWidth bool
	// Margin is the number of spaces on each side of the badge text.
	Margin int
	// TopPadding is the number of blank lines before the header.
	TopPadding int
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Header is the default header.
var Header = HeaderPrinter{
	BadgeTone:     TonePrimary,
	TitleStyle:    TonePrimary.Bold(),
	SubtitleStyle: ToneMuted.Style(),
	RuleTone:      ToneBorder,
	FullWidth:     true,
	Margin:        1,
	TopPadding:    1,
}

// WithBadge returns a copy of p with badge text s.
func (p HeaderPrinter) WithBadge(s string) *HeaderPrinter {
	p.Badge = s
	return &p
}

// WithSubtitle returns a copy of p with subtitle s.
func (p HeaderPrinter) WithSubtitle(s string) *HeaderPrinter {
	p.Subtitle = s
	return &p
}

// WithBadgeTone returns a copy of p whose badge is in tone t.
func (p HeaderPrinter) WithBadgeTone(t Tone) *HeaderPrinter {
	p.BadgeTone = t
	return &p
}

// WithTitleStyle returns a copy of p that styles the title with s.
func (p HeaderPrinter) WithTitleStyle(s Style) *HeaderPrinter {
	p.TitleStyle = s
	return &p
}

// WithSubtitleStyle returns a copy of p that styles the subtitle with s.
func (p HeaderPrinter) WithSubtitleStyle(s Style) *HeaderPrinter {
	p.SubtitleStyle = s
	return &p
}

// WithRuleTone returns a copy of p whose rule is in tone t.
func (p HeaderPrinter) WithRuleTone(t Tone) *HeaderPrinter {
	p.RuleTone = t
	return &p
}

// WithNoRule returns a copy of p without the rule under the header.
func (p HeaderPrinter) WithNoRule(b ...bool) *HeaderPrinter {
	p.NoRule = flag(b)
	return &p
}

// WithFullWidth returns a copy of p whose rule spans the terminal.
func (p HeaderPrinter) WithFullWidth(b ...bool) *HeaderPrinter {
	p.FullWidth = flag(b)
	return &p
}

// WithMargin returns a copy of p with n spaces on each side of the badge
// text.
func (p HeaderPrinter) WithMargin(n int) *HeaderPrinter {
	p.Margin = max(n, 0)
	return &p
}

// WithTopPadding returns a copy of p with n blank lines before the header.
func (p HeaderPrinter) WithTopPadding(n int) *HeaderPrinter {
	p.TopPadding = max(n, 0)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p HeaderPrinter) WithWriter(w io.Writer) *HeaderPrinter {
	p.Writer = w
	return &p
}

func (p HeaderPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p HeaderPrinter) render(c canvas, s string) string {
	title := strings.Join(strings.Fields(s), " ")
	badge := strings.Join(strings.Fields(p.Badge), " ")
	subtitle := strings.Join(strings.Fields(p.Subtitle), " ")

	if badge == "" {
		badge, title = title, ""
	}

	pad := 2 * p.Margin
	badgeW := text.Width(badge) + pad

	// Make the line fit: the subtitle goes first, then the title is cut,
	// and only then the badge.
	if c.width > 0 {
		used := badgeW
		if title != "" {
			used += 1 + text.Width(title)
		}

		if subtitle != "" && used+2+text.Width(subtitle) > c.width {
			if room := c.width - used - 2; room >= 8 {
				subtitle = c.cut(subtitle, room)
			} else {
				subtitle = ""
			}
		}

		if title != "" && used > c.width {
			if room := c.width - badgeW - 1; room >= 4 {
				title = c.cut(title, room)
			} else {
				title = ""
			}
		}

		if title == "" && badgeW > c.width {
			badge = c.cut(badge, max(c.width-pad, 1))
		}
	}

	line := c.badge(p.BadgeTone, badge, p.Margin)
	if badge == "" {
		line = ""
	}

	if title != "" {
		line += " " + c.style(p.TitleStyle, title)
	}

	if subtitle != "" {
		line += "  " + c.style(p.SubtitleStyle, subtitle)
	}

	out := strings.Repeat("\n", p.TopPadding) + line

	if !p.NoRule {
		width := text.Width(line)
		if p.FullWidth && c.width > 0 {
			width = c.width
		}

		out += "\n" + rule(c, width, c.g.Rule, p.RuleTone, 0, "")
	}

	return out
}

// Sprint returns the header with the text as title.
func (p HeaderPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p HeaderPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the title and returns the header.
func (p HeaderPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p HeaderPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the header with the text as title. A header is a whole
// block, so Print and Println print the same.
func (p HeaderPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the header with the text as title.
func (p HeaderPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats the title and prints the header.
func (p HeaderPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats the title and prints the header.
func (p HeaderPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints a header titled with every non-nil error in a.
func (p HeaderPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
