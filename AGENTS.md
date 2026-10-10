# AGENTS.md

Guidance for AI agents working in the PTerm repository.

## What this repo is

PTerm (`github.com/pterm/pterm`) is a Go library for building pretty
terminal output. It is a **library, not an application**: there is no `main`
package at the root. Consumers call printers like `pterm.Info.Println(...)` or
`pterm.Table.WithData(...).Render()`.

- Module path `github.com/pterm/pterm`, Go 1.27.
- The root package `pterm` is flat and holds every printer, one per
  `<name>_printer.go` with a matching `_test.go`.
- `putils/` converts common Go types into what printers accept (structs to
  tables, maps to trees or bars, indented text to trees) and wraps common
  patterns (`RunWithSpinner`, `DownloadWithProgress`).
- `internal/`: `screen` (the live renderer: regions, redraws, terminal size),
  `text` (cell width, wrapping, cutting), `vt` (a virtual terminal for tests),
  `snapshot` (the snapshot harness), `fuzzy` (prompt filtering).
- `_examples/`: one folder per printer, the source of the README examples and
  the GIFs. See [Examples](#examples-keep-them-in-sync).
- `docs/STYLEGUIDE.md` sets the look (tones, glyphs, spacing). Read it before
  changing rendered output. `docs/ADDING_A_PRINTER.md` is the recipe for a
  new printer. `docs/releases/v0.13/BREAKING.md` maps the v0.12 API to the
  current one.
- Large parts of `README.md` are generated between `<!-- ... -->` markers;
  edit the text outside them only.

## How printers are set up

1. **A struct and a package-level value of the same plain name.** The type
   is `TablePrinter`, the value users start from is `pterm.Table`. There is
   no `Default` prefix.
2. **Builder pattern with value receivers.** `WithX(...)` takes a value
   receiver, changes the copy and returns a pointer to it, so deriving a
   printer never changes the one it came from:

   ```go
   func (p SectionPrinter) WithLevel(level int) *SectionPrinter {
       p.Level = level
       return &p
   }
   ```

   Every exported field has a `With*` method.
3. **Four families**, defined in `printer.go` and asserted at compile time in
   `printers.go`. Add an assertion for every new printer.
   - `TextPrinter`: `Sprint*` returns styled text, `Print*` writes it.
     Header, Section, Box, Callout, Badge, Prefix (Info, Success, Warning,
     Error, Fatal, Debug) and others.
   - `RenderPrinter`: `Srender()` and `Render()` for configured data. Table,
     Tree, BarChart, Sparkline, KeyValue and others.
   - `LivePrinter`: `Start()` and `Stop()`, redraws in place. Spinner,
     Progressbar, Area, Multi, Flow. On a writer that is not a terminal they
     print their final state once.
   - `InteractivePrinter[T]`: `Show()` returns a `T`. Confirm, Select,
     Multiselect, TextInput, Continue. Without a terminal they return
     `ErrNotInteractive` unless a default is set.
4. **Render into a `canvas`** (`printer.go`): the color level, theme, glyphs
   and width of the target writer. A width of 0 means unknown; never wrap or
   cut then. Rendering returns a string, so `Sprint` and `Print` share it.

## Output, color and themes

- Writers: `print.go`. Normal output goes to `DefaultOutput()` (stdout);
  Logger, Warning, Error and Fatal go to `ErrorOutput()` (stderr).
- Global settings in `pterm.go` sit behind an atomic pointer and are only
  reached through functions (`DisableColor`, `ColorEnabled`, ...), so they
  are race safe. Keep it that way.
- Color comes only from `atomicgo.dev/color`. Printers color through
  **tones** (`ToneSuccess`, `ToneMuted`, ...), slots of the active `Theme` in
  `theme.go`. Use `Fixed(color)` only for a color that must ignore the theme.
- Glyphs (`glyphs.go`) have a Unicode and an ASCII set; render through
  `c.g`, never with a literal box or check character.

## Build, test, lint

Prefer the `Taskfile.yml` tasks:

- `task test`: `go test -race ./...`. Run with the race detector; race
  safety is a feature.
- `task lint`: golangci-lint v2.14 (`.golangci.yml`, includes `wsl_v5`,
  `revive`, `gosec`). `task fmt` applies fixes.
- `task examples:build`: builds and vets every example.

## Snapshot tests

Printers are locked with snapshots in `testdata/snapshots/*.snap`, rendered
styled and plain. Use the `snap`, `setup` and `capture` helpers in
`helpers_test.go`; `setup` gives a `vt` terminal for live printers.

- After an intended output change, run `UPDATE_SNAPSHOTS=1 go test ./...`
  and review the `.snap` diff.
- `UPDATE_SNAPSHOTS=1` is ignored when `CI` is set.
- `.snap` files stay LF (`.gitattributes`).

## Examples: keep them in sync

`_examples/` is the source of truth for user-facing documentation: the README
example sections, the per-example READMEs, and the VHS animations are all
generated from the `main.go` files in there (`scripts/`, `task animations`).
Never hand-edit the generated `README.md` or `animation.gif` files inside
`_examples/`.

**Important changes must be reflected in the examples.** Concretely:

- **New printer** → add a matching `_examples/<printer>/` directory with at
  least a `demo` example (the `demo` example is what the docs feature first).
- **New major feature or option on an existing printer** → add a small, focused
  example for it (one folder per feature, e.g. `table/from-csv`). Not every
  setting needs an example, but the main features of a printer should each
  have one.
- **Changed behavior or API** → update the affected examples so they still
  compile and show the current best practice.

Example conventions:

- Layout: `_examples/<printer>/<example-name>/main.go`, package `main`.
- Keep examples short and self-contained; use small, realistic demo data.
- Comments should sound human-written and explain intent or PTerm behavior,
  not narrate syntax. Do not use em-dashes in comments.
- Interactive examples need a `keys.tape` with the VHS key presses that drive
  them, so the animation can be recorded. The terminal size is measured
  automatically for other examples; interactive ones get 100x20 unless a
  `size.env` sets `COLUMNS` and `ROWS`. See `scripts/README.md`.
- Each example must build on its own. `task examples:build` builds and vets
  all of them (`go build ./_examples/...` matches nothing because of the
  underscore and still passes).

## Conventions & gotchas

- **Conventional Commits** (`feat:`, `fix:`, `docs:`, ...), see
  `CONTRIBUTING.md`.
- New printers are copied from the closest printer of the same family; see
  `docs/ADDING_A_PRINTER.md`.
- Breaking changes to the exported API go into
  `docs/releases/<version>/BREAKING.md`, with the old and the new call.
- No em dashes in comments, docs or examples.
