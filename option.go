package pterm

import (
	"cmp"
	"slices"
	"strconv"
	"strings"

	"atomicgo.dev/color"

	"github.com/pterm/pterm/internal/fuzzy"
	"github.com/pterm/pterm/internal/text"
)

// Option is one choice of a Select or a Multiselect. Only Label is needed;
// the other fields turn a plain list into one with groups and columns:
//
//	PACKAGE MANAGERS
//	❯ ◉ npm cache     ~/.npm            1.2 GB
//	  ○ yarn cache    ~/.cache/yarn   640.0 MB
type Option struct {
	// Label names the option. It is what the list shows and what the
	// filter matches.
	Label string
	// Value is returned when the option is chosen. Empty means Label.
	Value string
	// Description is shown muted in a column after the label, and cut to
	// fit the terminal.
	Description string
	// Detail is shown right-aligned in the last column, for sizes, counts
	// or versions.
	Detail string
	// Group puts the option under an uppercase group title. Keep the
	// options of a group next to each other.
	Group string
	// Disabled options are shown muted and cannot be chosen or toggled.
	Disabled bool
	// Selected marks the option as chosen from the start: the highlighted
	// option of a Select, or a checked option of a Multiselect.
	Selected bool
}

func (o Option) value() string {
	if o.Value != "" {
		return o.Value
	}

	return o.Label
}

// is reports whether s names o by its label or its value.
func (o Option) is(s string) bool {
	return s == o.Label || s == o.value()
}

// optionsOf turns plain labels into options.
func optionsOf(labels []string) []Option {
	options := make([]Option, len(labels))
	for i, l := range labels {
		options[i] = Option{Label: l}
	}

	return options
}

// OptionStyles style the rows of a Select or a Multiselect.
type OptionStyles struct {
	// Pointer styles the pointer in front of the active row.
	Pointer Style
	// Active styles the label of the active row.
	Active Style
	// Label styles the labels of the other rows.
	Label Style
	// Disabled styles the label and detail of disabled options.
	Disabled Style
	// Description styles the description column.
	Description Style
	// Detail styles the detail column.
	Detail Style
	// Group styles the group titles.
	Group Style
	// Match styles the letters that matched the filter.
	Match Style
	// Checked and Unchecked style the checkboxes of a Multiselect.
	Checked   Style
	Unchecked Style
	// Hint styles the scroll indicators and the "no matches" line.
	Hint Style
}

// DefaultOptionStyles are the styles Select and Multiselect start with.
var DefaultOptionStyles = OptionStyles{
	Pointer:     Style{Fg: TonePrimary},
	Active:      Style{Fg: TonePrimary, Attrs: color.Bold},
	Label:       Style{Fg: ToneText},
	Disabled:    Style{Fg: ToneMuted},
	Description: Style{Fg: ToneMuted},
	Detail:      Style{Fg: ToneText},
	Group:       Style{Fg: ToneMuted},
	Match:       Style{Fg: ToneAccent, Attrs: color.Underline},
	Checked:     Style{Fg: ToneSuccess},
	Unchecked:   Style{Fg: ToneMuted},
	Hint:        Style{Fg: ToneMuted},
}

// optionList is the state of a filterable, scrolling list of options,
// shared by Select and Multiselect.
type optionList struct {
	items []Option
	// height is the number of options shown at once; 0 shows all.
	height int
	// skipDisabled keeps the cursor off disabled options.
	skipDisabled bool

	query   []rune
	matches []optionMatch
	cursor  int // index into matches
	top     int // first visible match
}

type optionMatch struct {
	index     int
	positions []int
	score     int
}

func newOptionList(items []Option, height int, skipDisabled bool) *optionList {
	l := &optionList{items: items, height: max(height, 0), skipDisabled: skipDisabled}
	l.filter()

	return l
}

// setQuery filters the list by q and puts the cursor on the best match.
func (l *optionList) setQuery(q []rune) {
	l.query = q
	l.filter()
}

func (l *optionList) filter() {
	q := string(l.query)
	l.matches = l.matches[:0]

	for i, o := range l.items {
		if score, pos, ok := fuzzy.Match(q, o.Label); ok {
			l.matches = append(l.matches, optionMatch{index: i, positions: pos, score: score})
		}
	}

	if q != "" {
		l.rank()
	}

	l.cursor, l.top = 0, 0
	l.settle(1)
}

// rank orders matches by score but keeps groups together: groups come in
// the order of their best match, and the matches inside a group by score.
func (l *optionList) rank() {
	best := map[string]int{}
	first := map[string]int{}

	for _, m := range l.matches {
		g := l.items[m.index].Group
		if s, ok := best[g]; !ok || m.score > s {
			best[g] = m.score
		}

		if _, ok := first[g]; !ok {
			first[g] = m.index
		}
	}

	slices.SortStableFunc(l.matches, func(a, b optionMatch) int {
		ga, gb := l.items[a.index].Group, l.items[b.index].Group

		return cmp.Or(
			cmp.Compare(best[gb], best[ga]),
			cmp.Compare(first[ga], first[gb]),
			cmp.Compare(b.score, a.score),
			cmp.Compare(a.index, b.index),
		)
	})
}

// current returns the option under the cursor.
func (l *optionList) current() (Option, int, bool) {
	if len(l.matches) == 0 {
		return Option{}, -1, false
	}

	i := l.matches[l.cursor].index

	return l.items[i], i, true
}

// moveTo puts the cursor on the match at index i, which may be the
// default option.
func (l *optionList) moveTo(item int) {
	for k, m := range l.matches {
		if m.index == item {
			l.cursor = k
			l.scroll()

			return
		}
	}
}

// move moves the cursor by delta rows and wraps around at the ends.
func (l *optionList) move(delta int) {
	n := len(l.matches)
	if n == 0 {
		return
	}

	l.cursor = ((l.cursor+delta)%n + n) % n
	l.settle(sign(delta))
}

// page moves the cursor by one screen of rows without wrapping.
func (l *optionList) page(dir int) {
	if len(l.matches) == 0 {
		return
	}

	step := l.height
	if step == 0 {
		step = len(l.matches)
	}

	l.cursor = min(max(l.cursor+dir*step, 0), len(l.matches)-1)
	l.settle(dir)
}

// edge moves the cursor to the first or the last row.
func (l *optionList) edge(last bool) {
	if len(l.matches) == 0 {
		return
	}

	l.cursor, l.top = 0, 0
	dir := 1

	if last {
		l.cursor = len(l.matches) - 1
		dir = -1
	}

	l.settle(dir)
}

// settle moves the cursor in direction dir until it is on an option it may
// rest on, then scrolls it into view.
func (l *optionList) settle(dir int) {
	if l.skipDisabled && dir != 0 {
		for range l.matches {
			if !l.items[l.matches[l.cursor].index].Disabled {
				break
			}

			n := len(l.matches)
			l.cursor = ((l.cursor+dir)%n + n) % n
		}
	}

	l.scroll()
}

// scroll moves the visible window so it shows the cursor.
func (l *optionList) scroll() {
	if l.height == 0 {
		l.top = 0
		return
	}

	if l.cursor < l.top {
		l.top = l.cursor
	}

	if l.cursor >= l.top+l.height {
		l.top = l.cursor - l.height + 1
	}

	l.top = max(0, min(l.top, len(l.matches)-l.height))
}

// window returns the range of matches that is visible.
func (l *optionList) window() (start, end int) {
	if l.height == 0 {
		return 0, len(l.matches)
	}

	return l.top, min(l.top+l.height, len(l.matches))
}

func sign(n int) int {
	if n < 0 {
		return -1
	}

	return 1
}

// listView is what a printer adds to the rows of an optionList.
type listView struct {
	styles  OptionStyles
	pointer string
	// check returns the styled checkbox of an option, or "" for a list
	// without checkboxes.
	check func(index int) string
	// noMatches is shown when the filter matches nothing.
	noMatches string
}

// render returns the body lines of the list: group titles, rows and
// scroll indicators.
func (l *optionList) render(c canvas, v listView) []string {
	if len(l.matches) == 0 {
		return []string{promptIndent + c.style(v.styles.Hint, v.noMatches)}
	}

	cols := l.columns(c, v)
	start, end := l.window()
	grouped := l.hasGroups()
	overflow := l.height > 0 && len(l.matches) > l.height

	var lines []string

	// The up indicator takes the place of the blank line that separates
	// the first group, so a grouped list keeps its height while scrolling.
	up := ""
	if overflow {
		up = l.more(c, v, c.g.ArrowUp, start)
	}

	if overflow && !grouped {
		lines = append(lines, up)
	}

	group := "\x00"

	for k := start; k < end; k++ {
		m := l.matches[k]
		o := l.items[m.index]

		if grouped && o.Group != group {
			group = o.Group

			lines = append(lines, up)
			up = ""

			if group != "" {
				lines = append(lines, promptIndent+c.style(v.styles.Group, strings.ToUpper(group)))
			}
		}

		lines = append(lines, l.row(c, v, cols, m, k == l.cursor))
	}

	if overflow {
		lines = append(lines, l.more(c, v, c.g.ArrowDown, len(l.matches)-end))
	}

	return lines
}

// more renders a scroll indicator. It is an empty line when nothing is
// hidden in that direction, so the list keeps its height while scrolling.
func (l *optionList) more(c canvas, v listView, arrow string, n int) string {
	if n == 0 {
		return ""
	}

	return promptIndent + c.style(v.styles.Hint, arrow+" "+strconv.Itoa(n)+" more")
}

func (l *optionList) hasGroups() bool {
	for _, o := range l.items {
		if o.Group != "" {
			return true
		}
	}

	return false
}

// listColumns are the widths of the columns of a list. They are measured
// over every option, not just the visible ones, so the columns stay put
// while the list scrolls or gets filtered.
type listColumns struct {
	pointer, check, label, desc, detail int
}

func (l *optionList) columns(c canvas, v listView) listColumns {
	cols := listColumns{pointer: text.Width(v.pointer)}

	for i, o := range l.items {
		cols.label = max(cols.label, text.Width(o.Label))
		cols.desc = max(cols.desc, text.Width(o.Description))
		cols.detail = max(cols.detail, text.Width(o.Detail))

		if v.check != nil {
			cols.check = max(cols.check, text.Width(v.check(i)))
		}
	}

	if c.width <= 0 {
		return cols
	}

	// The screen cuts live lines one column before the edge.
	avail := c.width - 1

	fixed := len(promptIndent) + cols.pointer + 1 + cols.label
	if cols.check > 0 {
		fixed += cols.check + 1
	}

	if cols.detail > 0 {
		fixed += 2 + cols.detail
	}

	if cols.desc > 0 {
		room := avail - fixed - 2
		if room < 8 {
			// Too little room to say anything useful.
			room = 0
		}

		cols.desc = min(cols.desc, room)
	}

	if cols.desc > 0 {
		fixed += 2 + cols.desc
	}

	if fixed > avail {
		cols.label = max(cols.label-(fixed-avail), 1)
	}

	return cols
}

func (l *optionList) row(c canvas, v listView, cols listColumns, m optionMatch, active bool) string {
	o := l.items[m.index]
	st := v.styles

	var b strings.Builder

	b.WriteString(promptIndent)

	if active {
		b.WriteString(c.style(st.Pointer, v.pointer))
	} else {
		b.WriteString(strings.Repeat(" ", cols.pointer))
	}

	b.WriteString(" ")

	if v.check != nil {
		b.WriteString(text.PadRight(v.check(m.index), cols.check) + " ")
	}

	labelStyle, detailStyle := st.Label, st.Detail

	switch {
	case o.Disabled:
		labelStyle, detailStyle = st.Disabled, st.Disabled
	case active:
		labelStyle = st.Active
	}

	match := st.Match
	if active && !o.Disabled {
		match = match.With(color.Bold)
	}

	label := c.cut(highlight(c, o.Label, m.positions, labelStyle, match), cols.label)

	// Pad only when another column follows, so rows have no trailing
	// spaces.
	desc := ""
	if cols.desc > 0 && o.Description != "" {
		desc = c.style(st.Description, text.Truncate(o.Description, cols.desc, c.g.Ellipsis))
	}

	switch {
	case cols.detail > 0 && o.Detail != "":
		b.WriteString(text.PadRight(label, cols.label))

		if cols.desc > 0 {
			b.WriteString("  " + text.PadRight(desc, cols.desc))
		}

		b.WriteString("  " + text.PadLeft(c.style(detailStyle, o.Detail), cols.detail))

	case desc != "":
		b.WriteString(text.PadRight(label, cols.label) + "  " + desc)
	default:
		b.WriteString(label)
	}

	return b.String()
}

// highlight styles label with base, and the runes at positions with match.
func highlight(c canvas, label string, positions []int, base, match Style) string {
	if len(positions) == 0 {
		return c.style(base, label)
	}

	hit := make(map[int]bool, len(positions))
	for _, p := range positions {
		hit[p] = true
	}

	var (
		b   strings.Builder
		run []rune
		on  bool
	)

	flush := func() {
		if len(run) == 0 {
			return
		}

		st := base
		if on {
			st = match
		}

		b.WriteString(c.style(st, string(run)))
		run = run[:0]
	}

	for i, r := range []rune(label) {
		if hit[i] != on {
			flush()

			on = hit[i]
		}

		run = append(run, r)
	}

	flush()

	return b.String()
}
