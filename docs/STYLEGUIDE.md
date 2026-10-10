# PTerm styleguide

Every printer follows these rules. They are what makes a table, a spinner and a prompt look like parts of one tool. When you add or change a printer, check it against this page, and when a rule here is wrong, change the rule here first.

## Principles

- One line of code gives finished output. Defaults are the design; options exist for the cases the defaults do not cover.
- Color carries meaning, never decoration. Output must stay readable with colors off, which is why every status also has a glyph.
- Quiet by default. Most text is plain, titles are bold, and everything secondary is muted. A screen where everything is colored has no emphasis left.
- No boxes around things that do not need them. Use space, alignment and muted rules before reaching for a frame.

## Color

Printers never use raw colors. They use the nine tones of the theme (`TonePrimary`, `ToneMuted`...), and the theme maps them to colors. The default theme is Catppuccin Macchiato on dark terminals and Latte on light ones.

| Tone | Means | Used for |
| --- | --- | --- |
| Primary | "this is the thing" | titles of sections and headers, the active item in a list, spinner frames, the pointer, badges of the program itself |
| Accent | a value worth noticing | values in key-value output, flags and commands, highlighted search matches, the second end of gradients |
| Success | it worked | success glyphs, completed steps, fills of healthy meters |
| Warning | double check | warning glyphs and callouts, meters getting full |
| Error | it failed | error glyphs and callouts, canceled steps, meters over budget |
| Info | neutral note | info glyphs and callouts, links |
| Text | plain text | message bodies, table cells, list items |
| Muted | secondary | hints, elapsed times, units, table headers, descriptions, key hints, disabled options |
| Border | structure | rules, box frames, tree lines, the flow rail, meter tracks |

Rules:

- Text in a printer's own tone is short: a glyph, a title, a word. Bodies are Text.
- Muted text is never essential. If the user must read it, it is Text.
- Backgrounds are only used for badges. A badge is a tone background with dark text on dark themes and light text on light themes.
- Bold is for titles and the selected item. Never bold a whole paragraph.
- No italics, blinking or reverse video.

## Glyphs

All glyphs come from `CurrentGlyphs()`, never from string literals in a printer. The Unicode set is the default. The ASCII set is used when `PTERM_ASCII` is set, on the Linux console (`TERM=linux`) and on Windows consoles that cannot show Unicode.

| Purpose | Unicode | ASCII | Tone |
| --- | --- | --- | --- |
| Success | ✔ | v | Success |
| Error | ✖ | x | Error |
| Warning | ▲ | ! | Warning |
| Info | ● | i | Info |
| Debug | ◦ | . | Muted |
| Question | ? | ? | Primary |
| Pending | ○ | o | Muted |
| Skipped | – | - | Muted |
| Canceled | ✖ | x | Error |
| Pointer | ❯ | > | Primary |
| Bullet | • | * | Muted |
| Checked / unchecked | ◉ ○ | * o | Success / Muted |
| Inline separator | · | - | Muted |
| Callout bar | ┃ | \| | the callout's tone |
| Flow rail | ┌ │ └ | + \| \` | Border |
| Rule | ─ | - | Border |
| Meter fill / track | █ ░ | # . | tone / Border |
| Box | ╭ ╮ ╰ ╯ ─ │ | + - \| | Border |
| Tree | ├─ └─ │ | \|- \`- \| | Border |
| Ellipsis | … | ~ | as the cut text |
| Spinner | ⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏ | \|/-\\ | Primary |

Emoji and Nerd Font icons are not used. They render at different widths in different terminals and break alignment.

## Status lines

A status line is a glyph, a space and the text. Extra information follows the text after a space, muted:

```
✔ Measuring what can be freed 5.4s
● 3 packets captured
▲ The cache is stale
```

A status line that ends a live printer (spinner, step, prompt) keeps the same shape, so a finished flow reads as a list of status lines.

## Callouts

Text that must stand apart gets a bar in the callout's tone, a title with a glyph, the body in Text and an optional muted hint:

```
┃ ✖ Error
┃ Something went wrong while reading the config
┃ caused by: ENOENT: no such file
```

Multi-line messages of `Info`, `Success`, `Warning` and `Error` render as callouts. Use a callout for warnings before something destructive and for errors with a cause. Do not stack more than two.

## Layout and spacing

- Indentation is two spaces. Nested content (list items under a group title, the body of a prompt) is indented by two.
- Printers do not add blank lines around their output. The program decides spacing, and `Flow` adds it inside a flow. The exceptions are `Section` and `Header`, which separate parts of the output and keep one blank line before them.
- Columns are aligned with spaces, never with tabs. Numbers and sizes are right-aligned, everything else left-aligned.
- Inline items are joined with a muted ` · `: `v0.1.0 · 3 commands · 5 minutes ago`.
- Key hints read `[space] toggle · [a] all · [enter] confirm`, with the key in Text inside muted brackets and the action muted.

## Width

- Text, callouts and key-value values wrap to the terminal width.
- Table cells and list descriptions that do not fit are cut and end in `…`. They never wrap, because a wrapped row breaks the columns.
- When the width is unknown (the output is a pipe and `COLUMNS` is unset), nothing wraps and nothing is cut.
- Live output is always cut to the terminal width, because a wrapped live line cannot be erased.

## Tables, trees and lists

- Tables have no outer frame by default. The header is muted and uppercase, with a rule under it in Border. Rows have no separators.
- Trees draw their lines in Border, node names in Text (or Accent for leaves that are values), and an optional description column in Muted.
- Bullet lists use the muted bullet. A list can tone its bullets.
- Key-value output aligns the keys, which are Muted, and prints values in Text. Values that are versions, paths or numbers may use Accent.

## Live output

- Live printers redraw in place on terminals and print only their final state everywhere else.
- Elapsed time is shown muted after the text: `0.4s`, `12.3s`, `1m 34s`, `2h 5m`.
- Spinners and progress bars end as a status line. They never leave a half-drawn frame, even on Ctrl+C: an interrupted printer ends with the canceled glyph and a muted "canceled".
- The cursor is hidden while live output is drawn and shown again afterwards, also on Ctrl+C.

## Meters and charts

- A meter is a fill in a tone on a track of `░` in Border, followed by a right-aligned percentage and a muted label.
- Bars use eighth blocks so their length is exact to an eighth of a column.
- A chart colors its series from a fixed cycle: Primary, Accent, Info, Success, Warning, Error.

## Prompts

- A waiting prompt is a status line with the question glyph: `? What should be cleaned?`, followed by a muted hint such as `16 of 23` or the default answer.
- The active option has the Primary pointer and is bold. Disabled options are Muted and cannot be selected.
- Key hints go on the last line.
- An answered prompt collapses to one status line: `✔ What should be cleaned? 16 of 23` with the answer in Accent.

## Writing

- Sentence case for titles and messages: "Measuring what can be freed", not "Measuring What Can Be Freed".
- Messages do not end with a period unless they have more than one sentence.
- Group titles in lists are uppercase and muted: `PACKAGE MANAGERS`.
- Say what happened and what to do next. "Run `app --help` to see available commands." beats "An error occurred."

## Checklist for a printer

- [ ] Colors only through tones, glyphs only through `CurrentGlyphs()`.
- [ ] Readable with colors off and with the ASCII glyph set.
- [ ] Respects the width rules above, including unknown width.
- [ ] Works on a writer that is not a terminal.
- [ ] Every visual part can be changed with a `With*` method.
- [ ] Snapshot tests for the default look, styled and plain.
