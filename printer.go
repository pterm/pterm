package pterm

import (
	"context"
	"fmt"
	"io"
	"strings"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/text"
)

// Every printer belongs to one of four families. Each family is an
// interface, and printers.go asserts at compile time that every printer
// implements its family.

// TextPrinter prints formatted text, like fmt does, with a printer's style
// around it. Sprint methods return the styled text; Print methods write it.
type TextPrinter interface {
	Sprint(a ...any) string
	Sprintln(a ...any) string
	Sprintf(format string, a ...any) string
	Sprintfln(format string, a ...any) string
	Print(a ...any)
	Println(a ...any)
	Printf(format string, a ...any)
	Printfln(format string, a ...any)
	PrintOnError(a ...any)
}

// RenderPrinter renders data that the printer was configured with, such as
// a table or a tree. Srender returns the output; Render writes it.
type RenderPrinter interface {
	Srender() (string, error)
	Render() error
}

// LivePrinter redraws its output in place until it is stopped. On writers
// that are not terminals it prints its final state only, once Stop is
// called.
type LivePrinter interface {
	Stop() error
}

// InteractivePrinter asks the user for a value of type T.
//
// Show returns ErrNotInteractive when standard input is not a terminal and
// the printer has no default value, and ErrCanceled when the context set
// with WithContext ends first.
type InteractivePrinter[T any] interface {
	Show(prompt ...string) (T, error)
}

// canvas is what a printer renders with: the color level of its writer,
// the theme, the glyphs and the width available. Printers render into a
// string first, so the same code serves Sprint and Print.
type canvas struct {
	level color.Level
	theme Theme
	g     Glyphs
	width int // 0 when the width is unknown; never wrap or cut then
}

// canvasFor returns the canvas for output to w.
func canvasFor(w io.Writer) canvas {
	width, _ := screen.Size(w)

	return canvas{
		level: ColorLevel(w),
		theme: CurrentTheme(),
		g:     CurrentGlyphs(),
		width: width,
	}
}

// paint colors s with tone t.
func (c canvas) paint(t Tone, s string) string {
	return c.style(Style{Fg: t}, s)
}

// style renders s with st.
func (c canvas) style(st Style, s string) string {
	return st.Render(c.level, c.theme, s)
}

// bold renders s bold in tone t.
func (c canvas) bold(t Tone, s string) string {
	return c.style(Style{Fg: t, Attrs: color.Bold}, s)
}

// wrap wraps s to the canvas width minus indent columns. With an unknown
// width it only splits at newlines.
func (c canvas) wrap(s string, indent int) []string {
	if c.width <= 0 {
		return strings.Split(s, "\n")
	}

	return text.Wrap(s, max(c.width-indent, 10))
}

// cut shortens s to w columns with an ellipsis. A w of zero or less leaves
// s alone.
func (c canvas) cut(s string, w int) string { //nolint:unused // TODO: used by the render printers, remove this directive when they land.
	return text.Truncate(s, w, c.g.Ellipsis)
}

// node marks line as a node line, see screen.Node.
func node(line string) string {
	return screen.Node + line
}

// emit writes s to w through w's screen.
func emit(w io.Writer, s string) {
	if !OutputEnabled() || s == "" {
		return
	}

	screen.For(w).Write(s)
}

// textRenderer is the one method a text printer writes itself. The
// helpers below build the TextPrinter methods on top of it.
type textRenderer interface {
	out() io.Writer
	render(c canvas, s string) string
}

func sprintText(p textRenderer, s string) string {
	return screen.Strip(p.render(canvasFor(p.out()), s))
}

func printText(p textRenderer, s string) { //nolint:unused // TODO: used by the render printers, remove this directive when they land.
	if !OutputEnabled() {
		return
	}

	w := p.out()
	emit(w, p.render(canvasFor(w), s))
}

// printLine prints s and a newline. Text printers render s on its own and
// add the newline after the styling, so a trailing newline never ends up
// inside an escape sequence or a frame.
func printLine(p textRenderer, s string) {
	if !OutputEnabled() {
		return
	}

	w := p.out()
	emit(w, p.render(canvasFor(w), s)+"\n")
}

func printOnError(p textRenderer, a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok && err != nil {
			printLine(p, err.Error())
		}
	}
}

// sprintln is fmt.Sprintln without the newline, which text printers add
// after styling.
func sprintln(a ...any) string {
	return strings.TrimSuffix(fmt.Sprintln(a...), "\n")
}

// renderTo writes the output of a render printer to w.
func renderTo(w io.Writer, s string, err error) error { //nolint:unused // TODO: used by the render printers, remove this directive when they land.
	if err != nil {
		return err
	}

	emit(w, s+"\n")

	return nil
}

// orDefault returns w, or the default output when w is nil.
func orDefault(w io.Writer) io.Writer {
	if w == nil {
		return DefaultOutput()
	}

	return w
}

// orErrorOutput returns w, or the error output when w is nil.
func orErrorOutput(w io.Writer) io.Writer { //nolint:unused // TODO: used by the render printers, remove this directive when they land.
	if w == nil {
		return ErrorOutput()
	}

	return w
}

// orBackground returns ctx, or context.Background when ctx is nil.
func orBackground(ctx context.Context) context.Context {
	if ctx == nil {
		return context.Background()
	}

	return ctx
}
