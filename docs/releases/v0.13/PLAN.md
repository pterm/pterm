# PTerm v0.13 plan

This is the agreed plan for the v0.13 overhaul. Marvin settled every point below in the "PTerm v0.13 overhaul" project thread on 2026-10-10. Work happens on `feat/new-pterm` in every repo and ends in one PR per repo. Marvin merges and tags them himself. Nothing merges to `master` or `main` from an agent.

## Goals

- One line of code gives polished output. No setup, no boilerplate.
- Idiomatic Go, few dependencies, fast startup.
- Every printer is configurable, and the theme drives every color.
- Race-free by design, not by convention.
- Breaking changes are allowed. Each one is listed in `BREAKING.md` with migration steps an agent can follow.

## Repos and order

1. `atomicgo/color` gets pterm's color engine and becomes the only color library pterm uses.
2. `atomicgo/keyboard` drops testza, which removes an old pterm and gookit/color from the module graph, and replaces containerd/console with `x/term`.
3. `atomicgo/cursor` gets a current `go.mod` and whatever fixes the new renderer needs.
4. `pterm/pterm` pins the three modules by pseudo-version until Marvin tags them, then switches to the tags before merge.

All four repos move to Go 1.27.

## Color: atomicgo/color

- pterm's `internal/color` moves into atomicgo/color: ANSI 16, ANSI 256 and TrueColor, level detection from `NO_COLOR`, `FORCE_COLOR`, `COLORTERM`, `TERM`, the legacy Windows fallback and downsampling.
- Detection works per writer. A piped writer gets no color while stdout still does. `FORCE_COLOR` wins over the TTY check.
- Coloring and text attributes only (bold, dim, italic, underline, strikethrough, reverse). No predefined styles.
- Usable outside pterm, with docs and examples written so an agent can use it without reading the source.

## Theme

- Nine base slots: Primary, Accent, Success, Warning, Error, Info, Text, Muted, Border. Printers derive every color from these.
- Presets: Catppuccin Macchiato (default) and Latte. `PTERM_THEME=macchiato|latte|auto` picks one. `auto` is the default.
- Background detection reads `COLORFGBG`, then sends an OSC 11 query with a 50 ms timeout, and only when stdin and stdout are terminals. It runs once, on first styled output, and the result is cached. CI and pipes skip it and get Macchiato.
- `PTERM_THEME_<SLOT>` overrides one slot. Values: hex (`#b7bdf8`), an ANSI name (`red`, `brightblue`) or a 0-255 index.

## Look

`docs/STYLEGUIDE.md` defines the look, and every printer follows it. It covers which slot means what, the glyph table with ASCII fallbacks, spacing (2-space indent, one blank line between blocks), typography (bold only for titles, muted for hints and metadata) and borders. The reference is Marvin's screenshots: glyph status lines, callouts with a colored left bar, muted-header tables, rounded panels.

- Unicode glyphs, no Nerd Font. ASCII fallback when the terminal is not UTF-8, or when `PTERM_ASCII=1` is set.
- Info, Success, Warning and Error print a glyph and a toned title. Multi-line messages become a callout.
- Table: muted uppercase header, a rule under it, no outer box by default.
- Tree: muted connectors, toned keys, a description column.
- Box and Panel: rounded border in the Border color, title inset.
- Header: a Primary badge, a bold title, a muted subtitle and a full-width rule.
- Spinner: a braille dot in Primary, the text, muted elapsed time.
- Progressbar: Meter style with a toned fill, a dotted track and a muted percentage and ETA.
- Logger: a toned level glyph and muted `key=value` pairs. The slog handler stays.
- BigText: a Primary to Accent gradient.
- Bar chart: rewritten. Eighth-block bars on a dotted track, toned from a cycle of theme colors, stacked and grouped modes, negative values.
- Select and multiselect: groups, a description column, a right-aligned detail column, disabled options and a key-hint footer.

## New printers

Flow, Callout, KeyValue, Rule, Badge, KeyHints, Definitions, Sparkline and Meter.

Flow is a live container:

```go
flow := pterm.Flow.Start("disk clean")
step := flow.Step("Measuring what can be freed")
step.Success()
flow.Add(pterm.Callout.Warning("Heads up", "Nothing here can be undone."))
flow.Section("Packets")
flow.Stop()
```

- `Step` is a spinner step that ends as success, failure, warning or skip.
- `Add` takes any printer, including live and interactive ones. A live printer animates inside the rail and collapses to one finished line with elapsed time when it stops (`WithKeepOutput` keeps the last frame). A prompt collapses to its question and answer.
- `Section` draws a bold divider inside the rail.
- Plain pterm output while a flow runs joins the flow. Flows do not nest in v0.13.

## API

- Package vars drop the `Default` prefix: `pterm.Table`, `pterm.Spinner`, `pterm.Select`, `pterm.Info`. Types keep the `Printer` suffix. The `Interactive` prefix goes. The `Panel` type becomes `PanelItem`, the `Logger` type becomes `LoggerPrinter` and `MultiPrinter` becomes `pterm.Multi`.
- `With*` builders keep value receivers and return a pointer to the copy.
- Global switches (`Output`, `PrintColor`, `RawOutput`, `PrintDebugMessages`) become unexported. Only functions change them.
- `WithContext(ctx)` on live and interactive printers handles cancellation and timeouts.
- Interactive printers return `(T, error)`. With no TTY on stdin they return the default value if one is set, and `ErrNotInteractive` otherwise.
- The printer families may be reworked as long as the result stays idiomatic and adding a printer stays easy. `docs/ADDING_A_PRINTER.md` explains the steps.
- `deprecated.go` is deleted.
- Logger output and Warning, Error and Fatal go to stderr. Everything else goes to stdout. `WithWriter` overrides both.

## Runtime behavior

- One internal renderer owns all terminal output under one lock. Live printers register a region pinned at the bottom. Normal output printed while something is live scrolls above that region.
- While a live or interactive printer runs, pterm catches SIGINT, restores the cursor, marks the running item as canceled and exits with code 130. `WithOnInterruptFunc` on every live and interactive printer overrides that.
- Text, callouts and key-value blocks wrap to the terminal width. Table cells and list descriptions that do not fit end in `…`. With an unknown width (pipes) nothing wraps or truncates.
- Colors turn off automatically when the writer is not a terminal.

## putils

Converters from common Go types to pterm input, named `XFromY`: `TableFromStructs`, `TableFromCSV`, `TreeFromMap`, `TreeFromIndented`, `KeyValueFromStruct`, `KeyValueFromMap`, `BarsFromMap`, `BulletsFromSlice`. Runners: `RunWithSpinner(ctx, text, fn)` and `DownloadWithProgress`. Generics where they remove casts.

## Tests

- Snapshot tests stay, styled and plain.
- An internal virtual terminal replays live and interactive output and checks the final screen. It replaces the hand-written frame parsing.
- `go test -race` on Linux, macOS and Windows.

## Examples and docs

- Every printer has a `demo` example plus one per main feature.
- CI renders the VHS animations. It runs each example once headless, measures the output and sizes the terminal to fit (80 to 120 columns). PRs upload the GIFs as artifacts. Pushes to master commit them.
- `_examples/demo` is a trailer of about 30 seconds: a BigText title with a gradient sweep, a Flow installing a fake project, a table and tree drawn row by row, a bar chart whose bars grow next to a sparkline and meters, a scripted grouped multiselect, and a closing card with the one-liner and the install command.
- The README is regenerated. The pterm.sh site (`docs/index.html`) is out of scope.
- `docs/releases/v0.13/BREAKING.md` lists every breaking change with steps written for agents.
