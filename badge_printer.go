package pterm

import (
	"fmt"
	"io"
	"strings"

	"atomicgo.dev/color"
)

// BadgeVariant is the look of a badge.
type BadgeVariant uint8

// The badge variants.
const (
	// BadgeSolid draws the text on a background in the badge's tone:
	// " WARN ".
	BadgeSolid BadgeVariant = iota
	// BadgeChip draws the text in the badge's tone between brackets:
	// "[warn]".
	BadgeChip
)

// BadgePrinter prints a short label that marks the state or kind of
// something. It is made to be used inline:
//
//	fmt.Println(pterm.Badge.WithTone(pterm.ToneWarning).Sprint("dry run"), "nothing was deleted")
type BadgePrinter struct {
	// Tone is the background of a solid badge and the color of a chip.
	Tone Tone
	// Variant picks the solid or the chip look.
	Variant BadgeVariant
	// Uppercase prints the text in upper case.
	Uppercase bool
	// Padding is the number of spaces on each side of the text of a solid
	// badge.
	Padding int
	// Open and Close surround the text of a chip.
	Open, Close string
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Badge is the default badge: upper case text on a Primary background.
var Badge = BadgePrinter{
	Tone:      TonePrimary,
	Variant:   BadgeSolid,
	Uppercase: true,
	Padding:   1,
	Open:      "[",
	Close:     "]",
}

// WithTone returns a copy of p in tone t.
func (p BadgePrinter) WithTone(t Tone) *BadgePrinter {
	p.Tone = t
	return &p
}

// WithVariant returns a copy of p with variant v.
func (p BadgePrinter) WithVariant(v BadgeVariant) *BadgePrinter {
	p.Variant = v
	return &p
}

// WithUppercase returns a copy of p that prints its text in upper case.
func (p BadgePrinter) WithUppercase(b ...bool) *BadgePrinter {
	p.Uppercase = flag(b)
	return &p
}

// WithPadding returns a copy of p with n spaces on each side of a solid
// badge's text.
func (p BadgePrinter) WithPadding(n int) *BadgePrinter {
	p.Padding = max(n, 0)
	return &p
}

// WithBrackets returns a copy of p whose chips are wrapped in open and
// close.
func (p BadgePrinter) WithBrackets(open, closing string) *BadgePrinter {
	p.Open, p.Close = open, closing
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p BadgePrinter) WithWriter(w io.Writer) *BadgePrinter {
	p.Writer = w
	return &p
}

// Chip returns a copy of p that prints chips, "[warn]", with the text as
// given.
func (p BadgePrinter) Chip() *BadgePrinter {
	p.Variant = BadgeChip
	p.Uppercase = false

	return &p
}

// Solid returns a copy of p that prints solid badges, " WARN ".
func (p BadgePrinter) Solid() *BadgePrinter {
	p.Variant = BadgeSolid
	p.Uppercase = true

	return &p
}

func (p BadgePrinter) out() io.Writer { return orDefault(p.Writer) }

func (p BadgePrinter) render(c canvas, s string) string {
	s = strings.Join(strings.Fields(s), " ")
	if s == "" {
		return ""
	}

	if p.Uppercase {
		s = strings.ToUpper(s)
	}

	if p.Variant == BadgeChip {
		return c.paint(p.Tone, p.Open+s+p.Close)
	}

	return c.badge(p.Tone, s, p.Padding)
}

// badge draws s on a background in tone t, with pad spaces on each side
// and text that contrasts with the background.
func (c canvas) badge(t Tone, s string, pad int) string {
	s = strings.Repeat(" ", pad) + s + strings.Repeat(" ", pad)

	bg := t.In(c.theme)
	if bg.IsNone() {
		return c.style(Style{Attrs: color.Bold}, s)
	}

	return c.style(Style{Fg: Fixed(contrast(bg)), Bg: t, Attrs: color.Bold}, s)
}

// contrast returns a readable text color on top of bg: a deep shade of bg
// on light colors and a pale tint on dark ones.
func contrast(bg color.Color) color.Color {
	if bg.IsDark() {
		return bg.Mix(color.RGB(0xff, 0xff, 0xff), 0.9)
	}

	return bg.Mix(color.RGB(0, 0, 0), 0.8)
}

// Sprint returns the badge.
func (p BadgePrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p BadgePrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the text and returns the badge.
func (p BadgePrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p BadgePrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the badge without a newline, so text can follow it.
func (p BadgePrinter) Print(a ...any) { printText(p, fmt.Sprint(a...)) }

// Println prints the badge and a newline.
func (p BadgePrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats the text and prints the badge without a newline.
func (p BadgePrinter) Printf(format string, a ...any) { printText(p, fmt.Sprintf(format, a...)) }

// Printfln formats the text and prints the badge and a newline.
func (p BadgePrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints a badge for every non-nil error in a.
func (p BadgePrinter) PrintOnError(a ...any) { printOnError(p, a...) }
