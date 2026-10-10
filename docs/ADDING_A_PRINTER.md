# Adding a printer

This guide walks through adding a printer to PTerm. Read
[STYLEGUIDE.md](STYLEGUIDE.md) first: it decides what the output looks like,
this guide decides how the code is shaped.

## 1. Pick the family

Every printer belongs to one of four families, defined in `printer.go`:

| Family | Interface | Use it when | Copy from |
| --- | --- | --- | --- |
| Text | `TextPrinter` | the printer styles text the caller passes in, like `fmt` | `callout_printer.go` |
| Render | `RenderPrinter` | the printer draws data it was configured with | `table_printer.go` |
| Live | `LivePrinter` | the output changes in place until `Stop` | `spinner_printer.go` |
| Interactive | `InteractivePrinter[T]` | the printer asks the user for a value | `confirm_printer.go` |

Copy the closest printer of the same family and rename it. Do not invent a
fifth shape.

## 2. The struct and its default value

A printer is a struct with exported fields for every setting, and one
package-level value with the plain name users start from:

```go
type LabelPrinter struct {
	// Tone colors the label background.
	Tone Tone
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
}

// Label is the default label.
var Label = LabelPrinter{Tone: TonePrimary}
```

Rules for the fields:

- Colors are always a `Tone` or a `Style`, never a `color.Color`. A tone
  resolves against the active theme at render time, so `SetTheme` restyles
  the printer. Users who want a fixed color use `pterm.Fixed(c)`.
- Glyphs come from `c.g` (the canvas, see below). A field that overrides a
  glyph is empty by default and falls back to the glyph set, so the ASCII
  fallback keeps working.
- Every printer has `Writer io.Writer`, nil meaning `DefaultOutput()`.
  Printers that report problems default to `ErrorOutput()` instead.
- Live and interactive printers also have `Context context.Context` and
  `OnInterrupt func()`.
- The zero value of every field must be safe to render.

## 3. With methods

Each field gets a `With` method. It takes a value receiver, changes the copy
and returns a pointer, so chaining never touches the printer it started from:

```go
// WithTone returns a copy of p that uses tone t.
func (p LabelPrinter) WithTone(t Tone) *LabelPrinter {
	p.Tone = t
	return &p
}
```

Boolean options take `...bool` and use `flag(b)`, so `WithShowTimer()` turns
the option on and `WithShowTimer(false)` turns it off.

## 4. Render into a canvas

All drawing happens in one function that takes a `canvas` and returns a
string. The canvas carries the color level of the writer, the theme, the
glyphs and the width:

| Call | Does |
| --- | --- |
| `c.paint(tone, s)` | colors `s` |
| `c.bold(tone, s)` | colors `s` and makes it bold |
| `c.style(style, s)` | renders `s` with a full `Style` |
| `c.wrap(s, indent)` | wraps `s` to the width minus `indent` |
| `c.cut(s, w)` | shortens `s` to `w` columns with an ellipsis |
| `c.g` | the active `Glyphs` |
| `c.width` | the width, or 0 when unknown |

Never wrap or cut when `c.width` is 0: the output is going to a file or a
pipe, and lines stay as long as they are. Measure text with
`internal/text.Width`, which skips escape sequences and counts wide runes
correctly. `len` is wrong for both.

Start a line with `node(...)` when it is a status line (a glyph followed by
text). Inside a Flow, node lines sit on the rail instead of next to it.

### Text printers

Implement two methods and let the helpers in `printer.go` build the rest:

```go
func (p LabelPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p LabelPrinter) render(c canvas, s string) string {
	return c.style(Style{Fg: ToneText, Bg: p.Tone, Attrs: color.Bold}, " "+s+" ")
}
```

The nine `TextPrinter` methods are then one-liners on top of `sprintText`,
`printLine` and `printOnError`. Copy them from `callout_printer.go`. The
newline is added after styling, so it never ends up inside an escape sequence.

### Render printers

`Srender` builds the canvas with `canvasFor(p.out())`, renders, and returns
the string. `Render` is always:

```go
func (p TablePrinter) Render() error {
	s, err := p.Srender()
	return renderTo(p.out(), s, err)
}
```

Return an error from `Srender` for data the printer cannot draw, such as a
row with more cells than the header. Do not panic.

### Live printers

Live output goes through `internal/screen`. `Start` opens a region on the
writer's screen and returns a started copy of the printer; the shared state
of a running printer lives in a pointer field (see `spinnerRun`) behind a
mutex. Redraw with `region.Set`, end with `region.Close(final)`.

The screen handles what is hard to get right: it prints normal output above
the region, truncates live lines to the width, hides the cursor while
drawing, restores the terminal on Ctrl+C, and on writers that are not
terminals it prints only the final content. Pass `OnInterrupt` content and
the user's handler in `screen.RegionOptions` so Ctrl+C leaves a clean
"canceled" line behind.

Finish a live printer with a status line made with `status.glyph(c)`, so
spinners, steps and prompts all end the same way.

### Interactive printers

Build a `prompt[T]` (in `interactive.go`) and call `run`. You supply four
functions:

- `view` draws the prompt while it waits. Start with `question(c, text, hint)`.
- `key` handles one key press and reports whether the prompt is done.
- `result` returns the answer.
- `collapse` draws the line left behind. Use `answered(c, st, text, answer)`.

`run` handles the rest: one prompt at a time, Ctrl+C, context
cancellation, and the fallback when standard input is not a terminal
(`ErrNotInteractive`, or the default value when there is one). Support
`WithInput(r io.Reader)` so tests and scripted demos can feed keys.

## 5. Register the family

Add a compile-time assertion to `printers.go`:

```go
_ TextPrinter = LabelPrinter{}
```

Live printers are asserted on the pointer type, interactive printers with
their result type (`InteractivePrinter[string]`).

## 6. Test with snapshots

Tests live in `<name>_printer_test.go` in package `pterm_test`. `snap`
renders the printer twice, with true color and with colors off, and compares
both against `testdata/snapshots/`:

```go
func TestLabel(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Label.Sprint("beta") })
}
```

Create or update snapshots with `UPDATE_SNAPSHOTS=1 go test ./...`, then read
the `.snap` diff before committing. For live printers, render into the
virtual terminal that `setup(t)` returns and assert on `term.Screen()`,
which is what a user would see after all cursor movement.

Cover at least: the default, every option that changes the output, empty
input, a line wider than the terminal, and `GlyphsASCII`.

## 7. Examples

Add `_examples/<name>/demo/main.go`, plus one folder per main feature (see
AGENTS.md). Interactive examples need a `ci.go` or `keys.tape` so CI can
record them. Check that each one builds:

```sh
go vet ./_examples/<name>/...
```

## Checklist

- [ ] Struct with documented exported fields, plus a plain-named default value
- [ ] `With` method for every field, value receiver, pointer result
- [ ] Colors are tones, glyphs come from the glyph set
- [ ] No wrapping or cutting when the width is unknown
- [ ] Assertion in `printers.go`
- [ ] Snapshot tests, styled and plain
- [ ] `_examples/<name>/demo`
- [ ] Output matches every rule in STYLEGUIDE.md
- [ ] `task lint` and `task test` pass
