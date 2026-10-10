// Package vt is a small virtual terminal for tests. It understands the
// subset of escape sequences pterm writes (cursor movement, erasing, cursor
// visibility) and ignores colors, so a test can assert what a user would
// see on screen after live output finished.
package vt

import (
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

// Terminal is a virtual terminal. Its zero value is not usable; call New.
type Terminal struct {
	mu            sync.Mutex
	width, height int
	rows          [][]rune
	row, col      int
	cursorHidden  bool
	raw           strings.Builder
	pending       []byte
}

// wide marks the second cell of a wide character.
const wide = rune(-1)

// New returns a terminal of the given size. The scrollback is unlimited, so
// output never scrolls away.
func New(width, height int) *Terminal {
	return &Terminal{width: width, height: height, rows: [][]rune{{}}}
}

// IsTerminal reports true, so pterm draws live output.
func (t *Terminal) IsTerminal() bool { return true }

// Size returns the size given to New.
func (t *Terminal) Size() (width, height int) { return t.width, t.height }

// CursorHidden reports whether the last cursor visibility sequence hid it.
func (t *Terminal) CursorHidden() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.cursorHidden
}

// Raw returns every byte written so far.
func (t *Terminal) Raw() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	return t.raw.String()
}

// Screen returns what the terminal shows: every row with trailing spaces
// removed, and trailing empty rows dropped.
func (t *Terminal) Screen() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	lines := make([]string, len(t.rows))

	for i, r := range t.rows {
		var b strings.Builder

		for _, c := range r {
			if c != wide {
				b.WriteRune(c)
			}
		}

		lines[i] = strings.TrimRight(b.String(), " ")
	}

	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return strings.Join(lines, "\n")
}

// Write processes p like a terminal would.
func (t *Terminal) Write(p []byte) (int, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.raw.Write(p)

	data := append(t.pending, p...) //nolint:gocritic // pending is consumed here.
	t.pending = nil

	for i := 0; i < len(data); {
		c := data[i]

		switch {
		case c == 0x1b:
			n := t.escape(data[i:])
			if n == 0 {
				t.pending = append([]byte(nil), data[i:]...)
				return len(p), nil
			}

			i += n
		case c == '\n':
			t.row++
			t.col = 0
			t.ensureRow()
			i++
		case c == '\r':
			t.col = 0
			i++
		case c < 0x20:
			i++
		default:
			if !utf8.FullRune(data[i:]) {
				t.pending = append([]byte(nil), data[i:]...)
				return len(p), nil
			}

			r, size := utf8.DecodeRune(data[i:])
			t.put(r)
			i += size
		}
	}

	return len(p), nil
}

func (t *Terminal) ensureRow() {
	for len(t.rows) <= t.row {
		t.rows = append(t.rows, []rune{})
	}
}

func (t *Terminal) put(r rune) {
	w := runewidth.RuneWidth(r)
	if w == 0 {
		return
	}

	if t.width > 0 && t.col+w > t.width {
		t.row++
		t.col = 0
		t.ensureRow()
	}

	line := t.rows[t.row]
	for len(line) < t.col+w {
		line = append(line, ' ')
	}

	line[t.col] = r
	if w == 2 {
		line[t.col+1] = wide
	}

	t.rows[t.row] = line
	t.col += w
}

// escape handles the escape sequence at the start of data and returns its
// length, or 0 when data ends before the sequence does.
func (t *Terminal) escape(data []byte) int {
	if len(data) < 2 {
		return 0
	}

	switch data[1] {
	case '[':
		for j := 2; j < len(data); j++ {
			if data[j] >= 0x40 && data[j] <= 0x7e {
				t.csi(string(data[2:j]), data[j])
				return j + 1
			}
		}

		return 0
	case ']':
		for j := 2; j < len(data); j++ {
			if data[j] == 0x07 {
				return j + 1
			}

			if data[j] == 0x1b && j+1 < len(data) && data[j+1] == '\\' {
				return j + 2
			}
		}

		return 0
	default:
		return 2
	}
}

func (t *Terminal) csi(params string, final byte) {
	n := 1
	if v, err := strconv.Atoi(params); err == nil && v > 0 {
		n = v
	}

	switch final {
	case 'A':
		t.row = max(0, t.row-n)
	case 'B':
		t.row += n
		t.ensureRow()
	case 'C':
		t.col += n
	case 'D':
		t.col = max(0, t.col-n)
	case 'G':
		t.col = n - 1
	case 'K':
		t.eraseLine(params)
	case 'J':
		t.eraseLine("")
		t.rows = t.rows[:t.row+1]
	case 'h', 'l':
		if params == "?25" {
			t.cursorHidden = final == 'l'
		}
	}
}

func (t *Terminal) eraseLine(params string) {
	line := t.rows[t.row]

	switch params {
	case "2":
		t.rows[t.row] = []rune{}
	case "1":
		for i := 0; i <= t.col && i < len(line); i++ {
			line[i] = ' '
		}
	default:
		if t.col < len(line) {
			t.rows[t.row] = line[:t.col]
		}
	}
}
