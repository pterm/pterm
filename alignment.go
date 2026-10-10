package pterm

import "github.com/pterm/pterm/internal/text"

// Alignment places text inside a wider column.
type Alignment uint8

const (
	// AlignAuto right-aligns numbers and sizes and left-aligns everything
	// else.
	AlignAuto Alignment = iota
	// AlignLeft aligns text to the left edge.
	AlignLeft
	// AlignRight aligns text to the right edge.
	AlignRight
	// AlignCenter centers text.
	AlignCenter
)

// String names the alignment: "auto", "left", "right" or "center".
func (a Alignment) String() string {
	switch a {
	case AlignLeft:
		return "left"
	case AlignRight:
		return "right"
	case AlignCenter:
		return "center"
	default:
		return "auto"
	}
}

// pad fills s with spaces to width columns. AlignAuto counts as left.
func (a Alignment) pad(s string, width int) string {
	switch a {
	case AlignRight:
		return text.PadLeft(s, width)
	case AlignCenter:
		return text.Center(s, width)
	default:
		return text.PadRight(s, width)
	}
}

// setAt returns a copy of s with s[i] = v, grown with zero values when s is
// too short. Builders use it so a copy of a printer never shares a slice
// with the printer it came from.
func setAt[T any](s []T, i int, v T) []T {
	if i < 0 {
		return s
	}

	out := make([]T, max(len(s), i+1))
	copy(out, s)
	out[i] = v

	return out
}

// at returns s[i], or the zero value when i is out of range.
func at[T any](s []T, i int) T {
	var zero T
	if i < 0 || i >= len(s) {
		return zero
	}

	return s[i]
}
