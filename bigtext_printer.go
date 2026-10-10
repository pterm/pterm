package pterm

import (
	"fmt"
	"io"
	"maps"
	"math"
	"strings"
	"unicode"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/text"
)

// Letter is a character of big text with an optional style of its own.
type Letter struct {
	// String is the character.
	String string
	// Style colors the letter. The zero value takes the printer's tone or
	// gradient.
	Style Style
}

// Letters is a sequence of letters.
type Letters []Letter

// LettersFromString returns the characters of s as letters, all with
// style st, or with the printer's look when st is left out.
func LettersFromString(s string, st ...Style) Letters {
	var style Style
	if len(st) > 0 {
		style = st[0]
	}

	letters := make(Letters, 0, len(s))
	for _, r := range s {
		letters = append(letters, Letter{String: string(r), Style: style})
	}

	return letters
}

// BigTextPrinter prints text in large letters made of blocks, for title
// screens. By default the letters run through a gradient from Primary to
// Accent:
//
//	██████  ████████ ███████ ██████  ███    ███
//	██   ██    ██    ██      ██   ██ ████  ████
//	██████     ██    █████   ██████  ██ ████ ██
//	██         ██    ██      ██   ██ ██  ██  ██
//	██         ██    ███████ ██   ██ ██      ██
type BigTextPrinter struct {
	// Text is what Render prints.
	Text string
	// Letters replace Text when set, so single letters can have a style
	// of their own.
	Letters Letters
	// Font maps characters to their lines. Nil uses the built-in font,
	// see BigTextFont. Lowercase letters fall back to uppercase ones.
	Font map[rune][]string
	// Tone colors the letters when Gradient is empty.
	Tone Tone
	// Gradient lists the tones the letters run through from left to
	// right.
	Gradient []Tone
	// GradientOffset shifts the gradient to the right, in widths of the
	// text. The gradient bounces at its ends, so animating the offset
	// from 0 to 2 sweeps it across the text and back without a jump.
	GradientOffset float64
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// BigText is the default big text printer, with a gradient from Primary
// to Accent.
var BigText = BigTextPrinter{
	Tone:     TonePrimary,
	Gradient: []Tone{TonePrimary, ToneAccent},
}

// BigTextFont returns a copy of the built-in font, to change or extend
// and pass to WithFont.
func BigTextFont() map[rune][]string {
	return maps.Clone(bigFont)
}

// WithText returns a copy of p that prints s.
func (p BigTextPrinter) WithText(s string) *BigTextPrinter {
	p.Text = s
	return &p
}

// WithLetters returns a copy of p that prints letters, joined in order.
func (p BigTextPrinter) WithLetters(letters ...Letters) *BigTextPrinter {
	var all Letters
	for _, l := range letters {
		all = append(all, l...)
	}

	p.Letters = all

	return &p
}

// WithFont returns a copy of p that draws with font.
func (p BigTextPrinter) WithFont(font map[rune][]string) *BigTextPrinter {
	p.Font = font
	return &p
}

// WithTone returns a copy of p that colors the letters in tone t instead
// of a gradient.
func (p BigTextPrinter) WithTone(t Tone) *BigTextPrinter {
	p.Tone = t
	p.Gradient = nil

	return &p
}

// WithGradient returns a copy of p whose letters run through tones from
// left to right.
func (p BigTextPrinter) WithGradient(tones ...Tone) *BigTextPrinter {
	p.Gradient = tones
	return &p
}

// WithGradientOffset returns a copy of p whose gradient is shifted by
// offset widths of the text. See GradientOffset.
func (p BigTextPrinter) WithGradientOffset(offset float64) *BigTextPrinter {
	p.GradientOffset = offset
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p BigTextPrinter) WithWriter(w io.Writer) *BigTextPrinter {
	p.Writer = w
	return &p
}

func (p BigTextPrinter) out() io.Writer { return orDefault(p.Writer) }

// Srender returns the big text of Text or Letters. It returns ErrEmpty
// when there is nothing to print.
func (p BigTextPrinter) Srender() (string, error) {
	letters := p.Letters
	if len(letters) == 0 {
		letters = LettersFromString(p.Text)
	}

	if len(letters) == 0 {
		return "", ErrEmpty
	}

	return p.renderLetters(canvasFor(p.out()), letters), nil
}

// Render prints the big text of Text or Letters.
func (p BigTextPrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}

func (p BigTextPrinter) render(c canvas, s string) string {
	return p.renderLetters(c, LettersFromString(s))
}

// bigGlyph is one letter, drawn.
type bigGlyph struct {
	lines []string
	style Style
	width int
}

func (p BigTextPrinter) renderLetters(c canvas, letters Letters) string {
	if !StylingEnabled() {
		var b strings.Builder
		for _, l := range letters {
			b.WriteString(l.String)
		}

		return b.String()
	}

	// Each line of text becomes its own block of big letters.
	var blocks []string

	var line Letters

	for _, l := range letters {
		if l.String == "\n" {
			blocks = append(blocks, p.renderLine(c, line))
			line = nil

			continue
		}

		line = append(line, l)
	}

	blocks = append(blocks, p.renderLine(c, line))

	return strings.Join(blocks, "\n\n")
}

func (p BigTextPrinter) renderLine(c canvas, letters Letters) string {
	font := p.Font
	if font == nil {
		font = bigFont
	}

	var glyphs []bigGlyph

	height, width := 0, 0

	for _, l := range letters {
		for _, r := range l.String {
			lines, ok := font[r]
			if !ok {
				lines, ok = font[unicode.ToUpper(r)]
			}

			if !ok {
				continue
			}

			g := bigGlyph{lines: lines, style: l.Style}
			for _, line := range lines {
				g.width = max(g.width, text.Width(line))
			}

			glyphs = append(glyphs, g)
			height = max(height, len(lines))
			width += g.width
		}
	}

	// Letters end in a blank column, so the ink of the text ends one
	// column early. The gradient spans the ink.
	width = max(width-1, 1)
	rows := make([]string, height)

	for y := range height {
		var row strings.Builder

		x := 0

		for _, g := range glyphs {
			line := ""
			if y < len(g.lines) {
				line = g.lines[y]
			}

			row.WriteString(p.paintLine(c, g.style, text.PadRight(line, g.width), x, width))
			x += g.width
		}

		rows[y] = strings.TrimRight(row.String(), " ")
		if c.width > 0 {
			rows[y] = text.Truncate(rows[y], c.width, "")
		}
	}

	return strings.Join(rows, "\n")
}

// paintLine draws one line of a glyph that starts at column x of a text
// width columns wide. Runs of cells with the same color share one escape
// sequence, and spaces stay unpainted.
func (p BigTextPrinter) paintLine(c canvas, st Style, line string, x, width int) string {
	var (
		b   strings.Builder
		run strings.Builder
		cur color.Color
	)

	flush := func() {
		if run.Len() == 0 {
			return
		}

		if st != (Style{}) {
			b.WriteString(c.style(st, run.String()))
		} else {
			b.WriteString(c.paint(Fixed(cur), run.String()))
		}

		run.Reset()
	}

	for i, r := range []rune(line) {
		if r == ' ' {
			flush()
			b.WriteByte(' ')

			continue
		}

		col := p.colorAt(c, x+i, width)
		if col != cur {
			flush()

			cur = col
		}

		run.WriteString(p.glyph(c, r))
	}

	flush()

	return b.String()
}

// colorAt returns the color of column x of a text width columns wide.
func (p BigTextPrinter) colorAt(c canvas, x, width int) color.Color {
	if len(p.Gradient) == 0 {
		return p.Tone.In(c.theme)
	}

	t := p.GradientOffset
	if width > 1 {
		t += float64(x) / float64(width-1)
	}

	// Bounce at the ends: 0 to 1 runs forward, 1 to 2 back.
	t = math.Mod(math.Abs(t), 2)
	if t > 1 {
		t = 2 - t
	}

	return gradientAt(c.theme, p.Gradient, t)
}

// glyph maps a block character of the font to the glyph set, so the font
// also reads with ASCII glyphs.
func (p BigTextPrinter) glyph(c canvas, r rune) string {
	switch r {
	case '█':
		return c.g.MeterFill
	case '▀':
		return orFill(c.g.UpperHalf, c.g.MeterFill)
	case '▄':
		return orFill(c.g.LowerHalf, c.g.MeterFill)
	default:
		return string(r)
	}
}

func orFill(s, fill string) string {
	if s == "" {
		return fill
	}

	return s
}

// Sprint returns the text in big letters.
func (p BigTextPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p BigTextPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the text and returns it in big letters.
func (p BigTextPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p BigTextPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the text in big letters.
func (p BigTextPrinter) Print(a ...any) { printLine(p, fmt.Sprint(a...)) }

// Println prints the text in big letters.
func (p BigTextPrinter) Println(a ...any) { printLine(p, sprintln(a...)) }

// Printf formats the text and prints it in big letters.
func (p BigTextPrinter) Printf(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// Printfln formats the text and prints it in big letters.
func (p BigTextPrinter) Printfln(format string, a ...any) { printLine(p, fmt.Sprintf(format, a...)) }

// PrintOnError prints every non-nil error in a in big letters.
func (p BigTextPrinter) PrintOnError(a ...any) { printOnError(p, a...) }
