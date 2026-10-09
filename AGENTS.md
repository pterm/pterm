# AGENTS.md

Guidance for AI agents working in the PTerm repository.

## What this repo is

PTerm (`github.com/pterm/pterm`) is a Go library for building beautiful,
cross-platform terminal output. It is a **library, not an application** — there
is no `main` package at the root and nothing to "run". Consumers import it and
call printers like `pterm.Info.Println(...)` or `pterm.DefaultTable.WithData(...).Render()`.

It works across Windows CMD, macOS terminals, Linux, and CI systems, degrading
gracefully (falls back from TrueColor → ANSI → no color, and to raw text when
styling is unsupported or disabled).

- Module path: `github.com/pterm/pterm` (Go 1.26+)
- Root package `pterm` holds all the printers. Each printer lives in its own
  `*_printer.go` file with a matching `*_printer_test.go`.
- `putils/` — optional helper utilities (build tables from CSV/structs, letters
  from strings for BigText, download-with-progressbar, etc.). Depends on `pterm`.
- `internal/` — non-exported helpers: color-level detection & downsampling
  (`internal/color`), snapshot test harness (`internal/snapshot`), text
  measurement/centering, cancelation signals, etc.
- `_examples/` — runnable example programs, one subfolder per printer. The
  leading underscore keeps them out of the normal build. These are the source
  of truth for the README examples and the VHS animations.
- `docs/`, `README.md` — docs. Large sections of the README are generated (see
  the `<!-- ... -->` marker comments); do not hand-edit generated regions.

## How printers are set up

Every printer follows the same conventions. When adding or changing one, match
the existing pattern exactly.

1. **A struct + a `Default*` value.** Each printer is a struct (e.g.
   `SectionPrinter`) with an exported package-level default instance
   (`DefaultSection`). Users start from the default and customize it.

2. **Builder pattern with value receivers.** Configuration methods are
   `WithX(...)` methods that take a **value receiver**, mutate the copy, and
   return a **pointer** to the copy:

   ```go
   func (p SectionPrinter) WithLevel(level int) *SectionPrinter {
       p.Level = level
       return &p
   }
   ```

   This means `With*` never mutates the printer it was derived from — chaining
   is safe and side-effect-free. Preserve this; do not switch `With*` to pointer
   receivers.

3. **Printers belong to one of four families** (documented and compile-time
   enforced in `printers.go`):

   - **`TextPrinter`** (`interface_text_printer.go`) — prints formatted text
     directly. Implements `Sprint/Sprintln/Sprintf/Sprintfln` (return string)
     and `Print/Println/Printf/Printfln` (write to output) plus
     `PrintOnError(f)`. Examples: `BasicTextPrinter`, `PrefixPrinter` (powers
     `Info`/`Success`/`Warning`/`Error`/`Fatal`/`Debug`), `HeaderPrinter`,
     `SectionPrinter`, `BoxPrinter`, `CenterPrinter`, `Color`, `RGB`.
   - **`RenderPrinter`** (`interface_renderable_printer.go`) — renders complex
     multi-line content via `Render()` (to output) and `Srender()` (to string).
     Examples: `TablePrinter`, `TreePrinter`, `BarChartPrinter`,
     `BulletListPrinter`, `PanelPrinter`, `HeatmapPrinter`, `BigTextPrinter`.
   - **`LivePrinter`** (`interface_live_printer.go`) — output updates in place
     over time. `Start()` returns the started instance; `Stop()` terminates it.
     Examples: `SpinnerPrinter`, `ProgressbarPrinter`, `AreaPrinter`,
     `MultiPrinter`.
   - **Interactive printers** — read user input and return a result from
     `Show()`. No shared interface (return types differ) but follow the same
     `Default*` + `With*` + `Show()` shape. Examples:
     `InteractiveConfirmPrinter`, `InteractiveSelectPrinter`,
     `InteractiveMultiselectPrinter`, `InteractiveTextInputPrinter`,
     `InteractiveContinuePrinter`.

   `printers.go` has compile-time assertions (`var _ TextPrinter = ...`) that
   enforce family membership. If you add a printer, add its assertion there.

4. **The `Sprint` method is the core.** For text printers, `Print*` methods
   delegate to `Sprint*`, which delegate to `Fprint`/the package writers. Put
   the actual rendering logic in `Sprint`; the rest is boilerplate that follows
   the section_printer.go template.

## Global output, color, and styling

- Output goes through `print.go` (`Print`, `Fprint`, `Printo`, etc.). The
  default writer is `os.Stdout`; override with `SetDefaultOutput` or a
  printer's `WithWriter`.
- Global toggles live in `pterm.go`: `Output` (all output), `PrintColor`
  (color), `RawOutput` (styling). **Do not read/write these vars directly for
  concurrency** — use the `Enable*`/`Disable*` functions, which take
  `globalMu`. `PrintColor` is auto-detected at init from the environment
  (`NO_COLOR`, `TERM=dumb`, `FORCE_COLOR`, legacy Windows console) via
  `internal/color`.
- Theme: all default styles come from `ThemeDefault` in `theme.go`. Printers
  reference theme styles by pointer (e.g. `&ThemeDefault.SectionStyle`).
- Color: `color.go` (ANSI 3/4-bit `Color`), `rgb.go` (TrueColor `RGB` with
  fading/gradients). Color downsampling to the terminal's capability lives in
  `internal/color`.

## Build, test, lint

A `Taskfile.yml` exists — prefer its tasks:

- `task test` — runs `go test -race ./...` with `CGO_ENABLED=1`. (Plain
  `go test ./...` also works if you don't need the race detector.)
- `task lint` — `golangci-lint run` (config in `.golangci.yml`; strict linter
  set including `gosec`, `revive`, `wsl_v5` whitespace rules).
- `task fmt` — `golangci-lint run --fix` (also applies `gofmt`/`goimports`).

## Snapshot tests — important

Printer output is locked with **snapshot tests** (`snapshot_test.go` +
`internal/snapshot`, snapshots stored in `testdata/snapshots/*.snap`). Each
printer is rendered twice — styled and plain (raw) — and compared against the
committed snapshot.

- If you intentionally change rendered output, regenerate snapshots by running
  the tests with `UPDATE_SNAPSHOTS=1` (e.g. `UPDATE_SNAPSHOTS=1 go test ./...`),
  then **review the `.snap` diff** to confirm the change is what you intended.
- `UPDATE_SNAPSHOTS=1` is ignored when `CI` is set, so snapshots can't be
  silently rewritten in CI.
- `.snap` files are pinned to LF newlines via `.gitattributes`; keep it that way.

## Examples — keep them in sync

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
- Interactive examples need CI automation so animations can be recorded:
  either a `ci.go` that simulates key presses when `CI=true` (see
  `interactive_confirm/demo/ci.go`) or a `keys.tape` file.
- Each example must build on its own: `go build ./...` inside the example
  folder (the underscore prefix keeps `_examples/` out of the root build).

## Conventions & gotchas

- **Commit messages use Conventional Commits** (`feat:`, `fix:`, `chore:`,
  etc.) — see `conventionalcommit.json` / `CONTRIBUTING.md`.
- Keep the four-family builder pattern intact; new printers should be copyable
  from the closest existing printer of the same family.
- `_examples/` is not built by `go build ./...` (underscore prefix). Examples
  feed the README/animations pipeline; see the "Examples" section above. You
  don't normally need to run VHS locally.
- `deprecated.go` holds deprecated aliases kept for backward compatibility;
  don't remove them, add new deprecations there.
- The public API is widely depended upon — avoid breaking exported signatures.
