# Example animations

Every example in `_examples/<printer>/<example>/` is recorded as an
`animation.gif` with [VHS](https://github.com/charmbracelet/vhs). The GIFs are
committed next to the example, and the example READMEs, `_examples/README.md`
and the example section of the main `README.md` link them by relative path.

## Running it

You need Go, Python 3 and Docker. VHS runs from the
`ghcr.io/charmbracelet/vhs` image, so you do not have to install it.

| Command | What it does |
| --- | --- |
| `task animations` | Records every example and regenerates the docs. |
| `task animations:dry` | Records every example, leaves the docs alone. |
| `task animate -- table/demo` | Records one example. |
| `task animations:measure` | Builds and measures every example and prints the tapes it would record. Needs no Docker. |
| `task examples:build` | Builds and vets every example. |

Recordings run in parallel. Set `ANIMATE_JOBS` to change how many run at once
(default: the number of CPUs). Each VHS run starts its own Chromium, so on a
small machine fewer jobs can be faster.

## In CI

`.github/workflows/animator.yml` records the animations:

- On pull requests it uploads the GIFs as the `example-animations` artifact,
  laid out as `<printer>/<example>/animation.gif`. The docs are not touched.
- On pushes to `master` it also regenerates the docs and commits everything as
  `docs: autoupdate animations`. That commit does not trigger another run.
- It can be started by hand. Tick "commit" to push the result to the branch.

`go.yml` runs `scripts/build-examples.sh` on every push and pull request, so a
broken example fails CI before anything gets recorded.

## How the terminal is sized

The terminal fits the example, so the GIF shows everything without large empty
areas.

1. The example runs once without a terminal: output piped, stdin closed,
   `CI=true`, `TERM=xterm-256color`, `PTERM_THEME=macchiato` and `COLUMNS=80`.
   Live printers only print their final state when piped, which is what the
   height has to fit anyway.
2. Escape sequences are stripped and the widest line is measured in terminal
   cells. Wide characters such as CJK and most emoji count as two.
3. Columns are the widest line plus 4, kept between 80 and 120. If that is
   more than 80, the example runs again at the new width to count the rows
   there.
4. Rows are the counted lines (long lines count once per wrap) plus 2: one for
   the shell prompt after the example exits and one spare. The minimum is 8.

The tape then sets `Set Columns` and `Set Rows`, and VHS works out the pixel
size from the font. The recorded run gets the same `COLUMNS` value. Font
settings are fixed in the tape: `FontSize 18`, `LineHeight 1.2`, `Padding 24`,
theme Catppuccin Macchiato.

Interactive examples wait for input, so they cannot be measured this way. Any
example with a `keys.tape` is treated as interactive and gets 100x20 unless its
`size.env` says otherwise. If a measuring run takes longer than 60 seconds, the
example gets the same default size.

`Set Columns` and `Set Rows` need VHS 0.12 or newer. With an older image, set
`ANIMATE_SIZE_MODE=pixels` and the scripts write `Set Width` and `Set Height`
computed from `VHS_CELL_WIDTH` and `VHS_CELL_HEIGHT` instead.

## Files in an example folder

| File | Purpose |
| --- | --- |
| `main.go` | The example. Required. |
| `keys.tape` | VHS commands that drive an interactive example, for example `Type`, `Enter`, `Down` and `Sleep`. They run right after the example starts. Every interactive example needs one. |
| `size.env` | Optional size override, see below. |
| `options.tape` | Optional VHS `Set` commands, added after the generated settings. |
| `README.md`, `animation.gif` | Generated. Do not edit them. |

`size.env` holds `KEY=VALUE` lines. It is read, never executed, and `#` starts
a comment.

```sh
COLUMNS=100   # fixed width, skips the width measurement
ROWS=24       # fixed height, skips the height measurement
EXTRA_ROWS=4  # rows added on top of the measured height (default 2)
```

Set `EXTRA_ROWS` for a live example whose output is taller while it runs than
when it is done, and `COLUMNS` and `ROWS` for interactive examples that do not
fit 100x20.

A `Set Width`, `Set Height`, `Set Columns` or `Set Rows` in `options.tape` wins
over both the measurement and `size.env`. Prefer `size.env`; pixel sizes in
`options.tape` do not follow font changes.

## Settings

All of these are environment variables.

| Variable | Default | Meaning |
| --- | --- | --- |
| `ANIMATE_JOBS` | CPU count | Parallel builds and recordings. |
| `ANIMATE_PUBLISH` | `0` | `1` uploads the GIFs to vhs.charm.sh and links those URLs instead of the committed files. Same as `--publish`. |
| `ANIMATE_MIN_COLUMNS` / `ANIMATE_MAX_COLUMNS` | `80` / `120` | Width range. |
| `ANIMATE_COLUMN_PADDING` | `4` | Cells added to the widest line. |
| `ANIMATE_MIN_ROWS` | `8` | Smallest terminal height. |
| `ANIMATE_EXTRA_ROWS` | `2` | Rows added to the measured height. |
| `ANIMATE_INTERACTIVE_COLUMNS` / `ANIMATE_INTERACTIVE_ROWS` | `100` / `20` | Size for interactive examples and failed measurements. |
| `ANIMATE_MEASURE_TIMEOUT` | `60` | Seconds a measuring run may take. |
| `ANIMATE_RECORD_TIMEOUT` | `600` | Seconds a VHS recording may take. |
| `ANIMATE_SIZE_MODE` | `cells` | `pixels` for VHS images older than 0.12. |
| `ANIMATE_UPDATE_UNITTESTCOUNT` | `1` | `0` skips the README test count update on full runs. |
| `ANIMATE_VERBOSE` | `0` | `1` prints the log of every example, not only failed ones. |
| `VHS_DOCKER_IMAGE` | `ghcr.io/charmbracelet/vhs:v0.12.1` | VHS image. |
| `VHS_FONT_SIZE`, `VHS_LINE_HEIGHT`, `VHS_PADDING` | `18`, `1.2`, `24` | Font and spacing. |
| `VHS_THEME` | `Catppuccin Macchiato` | VHS theme. |
| `VHS_FRAMERATE` | `30` | GIF frame rate. |
| `VHS_END_SLEEP` | `4s` | How long the final frame stays. |

## The scripts

- `animate.sh` finds the examples, builds them all up front, records them in
  parallel and then calls `render-animation-docs.sh`.
- `record-example.sh` measures one example, writes its tape and runs VHS in
  Docker.
- `measure-output.py` measures captured output.
- `render-animation-docs.sh` writes the example READMEs and the generated
  sections of `README.md`.
- `build-examples.sh` builds and vets all examples. `go build ./_examples/...`
  cannot do this, because the go tool skips directories starting with `_` and
  reports success.
