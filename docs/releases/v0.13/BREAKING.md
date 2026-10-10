# Migrating to PTerm v0.13

v0.13 rebuilds PTerm. Most code that uses it needs edits. This guide is
written so a coding agent can run the migration from top to bottom: do the
steps in order, then use the reference sections to rewrite each call site.

## Steps

1. **Update the module.** Run `go get github.com/pterm/pterm@v0.13` and
   `go mod tidy`. PTerm now needs Go 1.27; raise the `go` line in `go.mod`
   if it is lower.
2. **Rename the default printers.** Apply the table in
   [Printer names](#printer-names) to every file. Done when
   `grep -rn 'pterm\.Default' --include='*.go' .` prints nothing.
3. **Replace global variables with functions.** Apply
   [Global settings](#global-settings). Done when no file assigns or reads
   `pterm.Output`, `pterm.PrintColor`, `pterm.RawOutput` or
   `pterm.PrintDebugMessages`.
4. **Replace colors and styles.** Apply [Colors, styles and
   themes](#colors-styles-and-themes). Done when no file uses `pterm.Fg*`,
   `pterm.Bg*`, `pterm.NewRGB*`, `pterm.RGB`, `pterm.ThemeDefault` or
   `*pterm.Style`.
5. **Fix each printer.** Build with `go build ./...` and go through the
   errors one printer at a time, using the printer's section under
   [Printers](#printers). Done when `go build ./...` and `go vet ./...`
   pass.
6. **Review behavior changes.** Read [Behavior](#behavior) and check the
   program against each item: output that moved to stderr, colors that turn
   off in pipes, prompts that fail without a terminal, and Ctrl+C. Fix
   tests that compare output.
7. **Run the tests.** `go test ./...` passes, and snapshot tests of PTerm
   output were updated on purpose, not blindly.

## Printer names

Printers lost the `Default` prefix. The printer types keep their names
unless the table says otherwise.

| v0.12 | v0.13 |
| --- | --- |
| `pterm.DefaultArea` | `pterm.Area` |
| `pterm.DefaultBarChart` | `pterm.BarChart` |
| `pterm.DefaultBasicText` | `pterm.Text` |
| `pterm.DefaultBigText` | `pterm.BigText` |
| `pterm.DefaultBox` | `pterm.Box` |
| `pterm.DefaultBulletList` | `pterm.BulletList` |
| `pterm.DefaultCenter` | `pterm.Center` |
| `pterm.DefaultHeader` | `pterm.Header` |
| `pterm.DefaultHeatmap` | `pterm.Heatmap` |
| `pterm.DefaultLogger` (type `pterm.Logger`) | `pterm.Logger` (type `pterm.LoggerPrinter`) |
| `pterm.DefaultMultiPrinter` | `pterm.Multi` |
| `pterm.DefaultPanel` | `pterm.Panel` |
| `pterm.DefaultParagraph` | `pterm.Paragraph` |
| `pterm.DefaultProgressbar` | `pterm.Progressbar` |
| `pterm.DefaultSection` | `pterm.Section` |
| `pterm.DefaultSpinner` | `pterm.Spinner` |
| `pterm.DefaultTable` | `pterm.Table` |
| `pterm.DefaultTree` | `pterm.Tree` |
| `pterm.Description` | `pterm.Info` |
| type `pterm.Panel` (one cell) | type `pterm.PanelItem` |

The interactive printers are listed in [Interactive
printers](#interactive-printers).

New printers, with nothing to migrate: `Badge`, `Callout`, `Definitions`,
`Flow`, `KeyHints`, `KeyValue`, `Meter`, `Rule`, `Sparkline`.

## Global settings

Settings are read and written through functions, which makes them safe to
use from many goroutines.

| v0.12 | v0.13 |
| --- | --- |
| `pterm.Output = false` | `pterm.DisableOutput()` |
| `pterm.Output = true` | `pterm.EnableOutput()` |
| reading `pterm.Output` | `pterm.OutputEnabled()` |
| `pterm.PrintColor = false` | `pterm.DisableColor()` |
| reading `pterm.PrintColor` | `pterm.ColorEnabled()` |
| `pterm.RawOutput = true` | `pterm.DisableStyling()` |
| reading `pterm.RawOutput` | `!pterm.StylingEnabled()` |
| `pterm.PrintDebugMessages = true` | `pterm.EnableDebugMessages()` |
| `pterm.GetDefaultOutput()` | `pterm.DefaultOutput()` |
| `pterm.GetTerminalWidth()` | `pterm.TerminalWidth()` |
| `pterm.GetTerminalSize()` | `pterm.TerminalSize()` |
| `pterm.GetTerminalHeight()` | `_, h := pterm.TerminalSize()` |
| `pterm.SetForcedTerminalSize(w, h)` | removed; tests write to a writer of their own size |
| `pterm.RecalculateTerminalSize()` | removed; the size is read on every render |
| `pterm.Printo`, `Fprinto`, `Sprinto` | removed; use `pterm.Area` |

`pterm.SetErrorOutput(w)` and `pterm.ErrorOutput()` are new; see
[Behavior](#behavior).

## Colors, styles and themes

PTerm no longer has color types of its own. Colors come from
[atomicgo.dev/color](https://atomicgo.dev/color), and printers color
themselves through **tones**: references to the nine slots of the active
theme (`Primary`, `Accent`, `Success`, `Warning`, `Error`, `Info`, `Text`,
`Muted`, `Border`). Prefer a tone over a fixed color, so output follows the
user's theme.

| v0.12 | v0.13 |
| --- | --- |
| `pterm.FgRed.Sprint(x)` | `pterm.ToneError.Sprint(x)` |
| `pterm.FgGreen.Sprint(x)` | `pterm.ToneSuccess.Sprint(x)` |
| `pterm.FgYellow.Sprint(x)` | `pterm.ToneWarning.Sprint(x)` |
| `pterm.FgBlue` / `pterm.FgCyan` | `pterm.ToneInfo` / `pterm.ToneAccent` |
| `pterm.FgGray` / `pterm.FgDarkGray` | `pterm.ToneMuted` |
| `pterm.FgMagenta`, `pterm.FgLightMagenta` | `pterm.TonePrimary` |
| a color that must never change | `pterm.Fixed(color.Red)` |
| `pterm.NewRGB(r, g, b).Sprint(x)` | `pterm.Fixed(color.RGB(r, g, b)).Sprint(x)` |
| `pterm.NewRGBFromHEX(s)` | `color.Parse(s)` |
| `rgb.Fade(...)` | `c.Mix(other, t)` on `color.Color` |
| `pterm.NewStyle(pterm.FgRed, pterm.Bold)` | `pterm.NewStyle(pterm.ToneError, color.Bold)` |
| `pterm.BgRed` and other backgrounds | `pterm.NewStyle(pterm.ToneText).On(pterm.ToneError)` |
| `*pterm.Style` fields and arguments | `pterm.Style` (a value) or `pterm.Tone` |
| `pterm.ThemeDefault.SomeStyle = ...` | `pterm.SetTheme(theme)` with changed slots, or `With*Style` on the printer |
| `pterm.RemoveColorFromString(s)` | unchanged |

`color` is `atomicgo.dev/color`. Most `With*Style` options of the printers
now take a `pterm.Style`, and options that only set a color take a
`pterm.Tone`.

The theme follows the terminal: Catppuccin Macchiato on dark backgrounds,
Latte on light ones. Users override it with `PTERM_THEME=macchiato|latte|ansi`
and single slots with `PTERM_THEME_<SLOT>`, for example
`PTERM_THEME_PRIMARY=#f5a97f`. Code sets it with `pterm.SetTheme`.

## Printers

Every printer keeps the builder pattern: `With*` methods return a changed
copy. Options that were removed are listed per printer. Options not listed
kept their name.

### Status messages (`Info`, `Success`, `Warning`, `Error`, `Fatal`, `Debug`)

- A message is a glyph and text, `✔ Deployed`, instead of a colored label.
  Multi-line text and titled messages print as a callout.
- `WithPrefix`, `WithScope`, `WithMessageStyle`, `WithShowLineNumber`,
  `WithLineNumberOffset`, `WithDebugger` and `GetFormattedPrefix` are gone.
  Use `WithTone`, `WithGlyph` and `WithTitle`.
- `Fatal` no longer has `WithFatal`; it always exits with code 1 after
  printing.
- `Warning`, `Error` and `Fatal` write to standard error.

### Text, header, section, box, center, paragraph

- `DefaultBasicText` is `pterm.Text`; `WithStyle` takes a `pterm.Style`.
- `Header` draws a badge, a title, a muted subtitle and a rule.
  `WithBackgroundStyle` is `WithBadgeTone` and `WithTextStyle` is
  `WithTitleStyle`. `WithMargin` is now the padding inside the badge.
  `WithBadge` and `WithSubtitle` are new; without a badge text, the title
  is drawn in the badge.
- `Section`: `WithIndentCharacter` is `WithPrefix`, and the prefix is
  empty by default. `WithStyle` styles the current level; `WithStyles` sets
  all levels.
- `Box`: `WithTitleTopLeft` and the other five title position methods are
  one `WithTitlePosition(pterm.BoxTitleTopLeft)`. `WithBoxStyle` is
  `WithBorderTone`. `WithTopLeftCornerString` and the other corner and line
  methods are one `WithGlyphs(pterm.BoxGlyphs{...})`. Corners are rounded.
- `Paragraph` and `Center` keep their API.

### Table

- `WithStyle` is `WithCellStyle`.
- `HasHeader` is true by default; call `WithHasHeader(false)` for tables
  without one.
- The table has no frame by default. `WithBoxed()` draws one;
  `WithBoxPrinter` is gone.
- `WithHeaderRowSeparator` and its style are `WithHeaderRule`.
  `WithRowSeparator` is `WithRowSeparators`. `WithSeparatorStyle` and
  `WithRowSeparatorStyle` are `WithBorderTone`.
- Numbers are right-aligned automatically. `WithAlignment`,
  `WithColumnAlignment` and `WithColumnAlignments` set alignment by hand.
- Cells that do not fit the terminal are cut with an ellipsis instead of
  wrapping.

### Tree, bullet list, panel

- `TreeNode` gained `Description`. `pterm.NewTreeFromLeveledList(l)` is
  `l.Tree()`.
- `BulletListItem.TextStyle` is gone; style all items with the list's
  `WithTextStyle`. `BulletListItem.BulletStyle` is `Tone`. `NewBulletListFromString` and
  friends moved to putils as `BulletsFromIndented` and `BulletsFromSlice`.
- `Panels` is `[][]pterm.PanelItem`. `WithBoxPrinter` is `WithBoxed`.

### Spinner

- `Stop() error` ends the spinner as a success; `Success`, `Fail`,
  `Warning`, `Info` and the new `Skip` end it with a status. None of them
  return the printer.
- `WithSequence` is `WithFrames`. `WithStyle` and `WithMessageStyle` are
  `WithTone`. `WithTimerStyle`, `WithTimerRoundingFactor`,
  `WithStartedAt`, `SetStartedAt`, `SetWriter`, `GenericStart` and
  `GenericStop` are gone.
- `WithContext(ctx)` stops the spinner as canceled when `ctx` ends.

### Progressbar

- `Stop() error` no longer returns the printer. `Fail` is new.
- The current value of a running bar is read with `Progress()`, not the
  `Current` field.
- `WithBarCharacter` is `WithFill` and `WithBarFiller` is `WithTrack`.
  `WithBarStyle` is `WithTone`, `WithBarFillerStyle` is `WithTrackTone` and
  `WithTitleStyle` is `WithTitleTone`. `WithLastCharacter`,
  `WithBarPartialCharacters`, `WithElapsedTimeRoundingFactor` and
  `WithStartedAt` are gone; the bar grows in eighth blocks by itself.
- `SetWriter`, `SetStartedAt`, `IsActive` and `GetElapsedTime` are gone;
  `Elapsed()` and `ResetTimer()` remain.
- The bar stops by itself at its total. A total of 0 never stops by itself.

### Area and Multi

- `Area.Start` and `Area.Stop() error`; `GetContent` is `Content`.
- `pterm.Multi` groups live printers: `Start(printers...)`, `Add`,
  `Stop() error`. Every live printer now draws through one renderer, so
  `WithUpdateDelay`, `NewWriter` buffering and `GenericStart` are no longer
  needed. `NewWriter()` returns the target writer for old code.

### Logger and slog

- The type is `pterm.LoggerPrinter`; `pterm.Logger` is the default value.
- Logs go to standard error. Timestamps are off by default
  (`WithTime()`), and the default format is `15:04:05`.
- `Fatal` always exits, even when the level hides the line.
- `NewSlogHandler(logger)` accepts nil for the default logger. `WithAttrs`
  adds to existing attributes, and groups become dotted keys
  (`req.user.name`).

### Bar chart

- `Bar.Value` is a `float64`. Bars are `pterm.Bars{{Label: "npm", Value: 12}}`.
- `WithHorizontal` is gone; charts are horizontal by default and
  `WithVertical()` turns them. `WithShowValue` is `WithShowValues`, on by
  default.
- `WithGroups`, `WithStacked` and `WithSeries` are new.

### Heatmap

- `HeatmapData` is `[][]float64`.
- Background-colored cells, complementary text colors and the discrete and
  RGB color modes are replaced by `WithGradient(tones...)`.

### Big text

- `pterm.NewLettersFromString` is `pterm.LettersFromString`.
- `BigCharacters map[string]string` is `WithFont(map[rune][]string)`;
  `pterm.BigTextFont()` returns a copy of the default font to edit.
- `WithGradient` colors the text column by column.

### Interactive printers

The `Interactive` prefix is gone, and every prompt is built on one base.

| v0.12 | v0.13 |
| --- | --- |
| `pterm.DefaultInteractiveConfirm` | `pterm.Confirm` (type `ConfirmPrinter`) |
| `pterm.DefaultInteractiveContinue` | `pterm.Continue` (type `ContinuePrinter`) |
| `pterm.DefaultInteractiveMultiselect` | `pterm.Multiselect` (type `MultiselectPrinter`) |
| `pterm.DefaultInteractiveSelect` | `pterm.Select` (type `SelectPrinter`) |
| `pterm.DefaultInteractiveTextInput` | `pterm.TextInput` (type `TextInputPrinter`) |

All prompts:

- `WithDefaultText` is `WithText`. `Show(text)` still overrides it.
- `WithTextStyle`, `WithSuffixStyle`, `WithDelimiter` and the other
  `*Style` options that took a `*pterm.Style` are gone. Select and
  Multiselect take `WithStyles(pterm.OptionStyles{...})`, TextInput
  `WithStyles(pterm.TextInputStyles{...})` and Continue
  `WithStyles(pterm.ContinueStyles{...})`. Start from
  `pterm.DefaultOptionStyles` to change a single style.
- `WithContext`, `WithOnInterrupt` and `WithInput` are new.
- `Show` returns an error that callers must check: `pterm.ErrNotInteractive`
  without a terminal and no default, `pterm.ErrCanceled` when the context
  ends, `pterm.ErrInterrupted` on Ctrl+C with `WithOnInterrupt` set, and
  `pterm.ErrEmpty` for lists without options.

Per prompt:

- **Select**: `WithSelectedOptionStyle` is `WithStyles`.
  `WithFilterInputPlaceholder` is `WithFilterPlaceholder`. `WithItems` takes
  `[]pterm.Option` with descriptions, details, groups and disabled options;
  `Show` returns the option's `Value`, which defaults to its `Label`.
- **Multiselect**: `WithCheckmark(&pterm.Checkmark{...})` is
  `WithCheckmark(checked, unchecked)`. Filtering starts with `/`, because
  `a` and `n` select all and none. `WithFilterInputPlaceholder` is
  `WithFilterPlaceholder`.
- **TextInput**: `WithDelimiter` and `WithSubmitHint` are gone. Multi-line
  input is submitted with `WithSubmitKey` (Ctrl+D by default).
- **Confirm**: `WithConfirmStyle` and `WithRejectStyle` are gone.
- **Continue**: `WithShowShortHandles` and `WithOptionsStyle` are gone; the
  prompt shows the handles and the option names together.

## putils

putils converts Go values into pterm data and leaves printing to the
printers.

| v0.12 | v0.13 |
| --- | --- |
| `putils.TableFromStructSlice(p, rows)` | `p.WithData(putils.TableFromStructs(rows))` |
| `putils.DefaultTableFromStructSlice(rows)` | `pterm.Table.WithData(putils.TableFromStructs(rows))` |
| `putils.TableDataFromCSV(s)` | `putils.TableFromCSV(s)` (returns an error) |
| `putils.TableDataFromTSV(s)` | `putils.TableFromSeparated(s, '\t')` |
| `putils.TableDataFromSeparatedValues(s, sep, "\n")` | `putils.TableFromSeparated(s, sep)` |
| `putils.TreeFromLeveledList(l)` | `l.Tree()` |
| `putils.BulletListFromString(s, pad)` | `pterm.BulletList.WithItems(putils.BulletsFromIndented(s))` |
| `putils.BulletListFromStrings(s, pad)` | `pterm.BulletList.WithItems(putils.BulletsFromSlice(s))` |
| `putils.RunWithSpinner(s, fn)` | `putils.RunWithSpinner(ctx, text, fn)` |
| `putils.DownloadFileWithProgressbar(...)` | `putils.DownloadWithProgress(ctx, title, url, w)` |
| `putils.LettersFromString*` | `pterm.LettersFromString` |
| `putils.CenterText(s)` | `pterm.Center.Sprint(s)` |
| `putils.RGBFromHEX(s)` | `color.Parse(s)` |
| `putils.PrintAverageExecutionTime` | removed |

New: `KeyValueFromStruct`, `KeyValueFromMap`, `TreeFromMap`,
`TreeFromIndented`, `BarsFromMap`.

## Behavior

- **Standard error.** `Warning`, `Error`, `Fatal` and the logger write to
  `pterm.ErrorOutput()`, which is standard error. Code that captured them
  from standard output must read standard error or call
  `pterm.SetErrorOutput`.
- **Colors in pipes.** Colors turn off when the writer is not a terminal.
  `FORCE_COLOR` turns them on, `NO_COLOR` off.
- **Live output in pipes.** Spinners, progress bars and areas print only
  their final line when the writer is not a terminal.
- **No terminal for prompts.** Interactive printers return their default
  value, or `pterm.ErrNotInteractive` when they have none, instead of
  blocking.
- **Ctrl+C.** Live and interactive printers restore the terminal, leave a
  canceled line and exit with code 130. `WithOnInterrupt(fn)` runs `fn`
  instead of exiting.
- **Width.** Output wraps or cuts to the terminal width, and never when the
  width is unknown.
- **Glyphs.** Status glyphs are Unicode (`✔ ✖ ▲ ●`). The Linux console,
  legacy Windows consoles and `PTERM_ASCII=1` get ASCII.
- **Spacing.** Printers add no blank lines around their output, except
  `Header` and `Section`, which print one blank line before themselves.
