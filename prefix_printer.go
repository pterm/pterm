package pterm

import (
	"fmt"
	"io"
	"strings"

	"github.com/pterm/pterm/internal/screen"
)

// PrefixPrinter prints status messages: a glyph in the printer's tone, then
// the message. A message with several lines, or a printer with a title,
// becomes a callout.
//
//	✔ Deployed to production
type PrefixPrinter struct {
	// Tone colors the glyph and the title.
	Tone Tone
	// Glyph is drawn before the message. Empty means the glyph of Kind
	// from the active glyph set.
	Glyph string
	// Kind picks the default glyph and the default writer.
	Kind MessageKind
	// Title is shown above multi-line messages. Empty uses the first line
	// of the message as title.
	Title string
	// Writer receives the output. Nil means the default output, or the
	// error output for warnings, errors and fatal messages.
	Writer io.Writer
}

// MessageKind is the kind of a status message.
type MessageKind uint8

// The kinds of status messages.
const (
	KindInfo MessageKind = iota
	KindSuccess
	KindWarning
	KindError
	KindFatal
	KindDebug
)

// The status printers.
var (
	// Info prints neutral notes.
	Info = PrefixPrinter{Tone: ToneInfo, Kind: KindInfo}
	// Success prints things that worked.
	Success = PrefixPrinter{Tone: ToneSuccess, Kind: KindSuccess}
	// Warning prints things to double check, to the error output.
	Warning = PrefixPrinter{Tone: ToneWarning, Kind: KindWarning}
	// Error prints failures, to the error output.
	Error = PrefixPrinter{Tone: ToneError, Kind: KindError}
	// Fatal prints like Error and then exits the program with code 1.
	Fatal = PrefixPrinter{Tone: ToneError, Kind: KindFatal}
	// Debug prints only after EnableDebugMessages.
	Debug = PrefixPrinter{Tone: ToneMuted, Kind: KindDebug}
)

// WithTone returns a copy of p that uses tone t.
func (p PrefixPrinter) WithTone(t Tone) *PrefixPrinter {
	p.Tone = t
	return &p
}

// WithGlyph returns a copy of p that draws g before messages.
func (p PrefixPrinter) WithGlyph(g string) *PrefixPrinter {
	p.Glyph = g
	return &p
}

// WithTitle returns a copy of p that prints messages as a callout with
// title t.
func (p PrefixPrinter) WithTitle(t string) *PrefixPrinter {
	p.Title = t
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p PrefixPrinter) WithWriter(w io.Writer) *PrefixPrinter {
	p.Writer = w
	return &p
}

func (p PrefixPrinter) out() io.Writer {
	switch {
	case p.Writer != nil:
		return p.Writer
	case p.Kind == KindWarning || p.Kind == KindError || p.Kind == KindFatal:
		return ErrorOutput()
	default:
		return DefaultOutput()
	}
}

func (p PrefixPrinter) glyph(g Glyphs) string {
	if p.Glyph != "" {
		return p.Glyph
	}

	switch p.Kind {
	case KindSuccess:
		return g.Success
	case KindWarning:
		return g.Warning
	case KindError, KindFatal:
		return g.Error
	case KindDebug:
		return g.Debug
	default:
		return g.Info
	}
}

func (p PrefixPrinter) render(c canvas, s string) string {
	if p.Kind == KindDebug && !DebugMessagesEnabled() {
		return ""
	}

	s = strings.TrimRight(s, "\n")
	glyph := c.paint(p.Tone, p.glyph(c.g))

	if p.Title == "" && !strings.Contains(s, "\n") {
		body := c.paint(ToneText, s)
		if p.Kind == KindDebug {
			body = c.paint(ToneMuted, s)
		}

		lines := c.wrap(body, 2)
		for i := range lines[1:] {
			lines[i+1] = "  " + lines[i+1]
		}

		return node(glyph + " " + strings.Join(lines, "\n"))
	}

	title, body := p.Title, s
	if title == "" {
		title, body, _ = strings.Cut(s, "\n")
	}

	return Callout.WithTone(p.Tone).WithGlyph(p.glyph(c.g)).WithTitle(title).render(c, body)
}

// Sprint returns the message as Print would print it.
func (p PrefixPrinter) Sprint(a ...any) string { return sprintText(p, fmt.Sprint(a...)) }

// Sprintln is like Sprint and adds a newline.
func (p PrefixPrinter) Sprintln(a ...any) string { return sprintText(p, sprintln(a...)) + "\n" }

// Sprintf formats the message and returns it as Print would print it.
func (p PrefixPrinter) Sprintf(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...))
}

// Sprintfln is like Sprintf and adds a newline.
func (p PrefixPrinter) Sprintfln(format string, a ...any) string {
	return sprintText(p, fmt.Sprintf(format, a...)) + "\n"
}

// Print prints the message. Messages always end with a newline, so Print
// and Println print the same.
func (p PrefixPrinter) Print(a ...any) { p.print(fmt.Sprint(a...)) }

// Println prints the message.
func (p PrefixPrinter) Println(a ...any) { p.print(sprintln(a...)) }

// Printf formats and prints the message.
func (p PrefixPrinter) Printf(format string, a ...any) { p.print(fmt.Sprintf(format, a...)) }

// Printfln formats and prints the message.
func (p PrefixPrinter) Printfln(format string, a ...any) { p.print(fmt.Sprintf(format, a...)) }

// PrintOnError prints every non-nil error in a.
func (p PrefixPrinter) PrintOnError(a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok && err != nil {
			p.print(err.Error())
		}
	}
}

// PrintOnErrorf prints every non-nil error in a, wrapped with format, which
// must contain one %w or %v verb.
func (p PrefixPrinter) PrintOnErrorf(format string, a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok && err != nil {
			p.print(fmt.Errorf(format, err).Error())
		}
	}
}

func (p PrefixPrinter) print(s string) {
	if out := sprintRaw(p, s); out != "" {
		emit(p.out(), out+"\n")
	}

	if p.Kind == KindFatal {
		screen.Exit(1)
	}
}

// sprintRaw renders for the printer's own writer and keeps node markers,
// for output that goes to a screen.
func sprintRaw(p textRenderer, s string) string {
	return p.render(canvasFor(p.out()), s)
}
