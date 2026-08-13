package pterm

import (
	"strings"
	"time"
	"unicode/utf8"

	"github.com/pterm/pterm/internal"
)

func getMaxW(s string) int {
	return internal.GetStringMaxWidth(s)
}

// fit the text for terminal width
func textFitWidth(text string) string {
	return strings.Join(linesFitWidth(strings.Split(text, "\n")), "\n")
}

// fit the []string for terminal width
func linesFitWidth(ss []string) []string {
	w := max(GetTerminalWidth()-1, 1)
	if getMaxW(strings.Join(ss, "\n")) >= w {
		// find the last index that GetStringWidth(s[:index])<=width
		findIndex := func(s string, width int) int {
			l, r := 0, len(s)+1
			for l+1 < r {
				mid := (l + r) >> 1
				if getMaxW(s[:mid]) <= width {
					l = mid
				} else {
					r = mid
				}
			}
			return l
		}

		// count runes maybe invalid bytes at the end
		countValidRunes := func(s string) int {
			if s == "" {
				return 0
			}
			// Remove all invalid bytes at the end
			for {
				r, size := utf8.DecodeLastRuneInString(s)
				if r != utf8.RuneError || size != 1 {
					break
				}
				s = s[:len(s)-1]
			}
			return utf8.RuneCountInString(s)
		}
		buffer := make([]string, 0)
		for _, s := range ss {
			if getMaxW(s) <= w {
				buffer = append(buffer, s)
				continue
			}
			for len(s) > 0 {
				i := countValidRunes(s[:findIndex(s, w)])
				front, end := string([]rune(s)[:i]), string([]rune(s)[i:])
				buffer = append(buffer, front)
				s = end
			}
		}
		return buffer
	}

	return ss
}

// this create a [time.Ticker] that
// call the onChange(width)
// when the terminal's width changed
func watchWidth(done <-chan struct{}, interval time.Duration, onChange func(w int)) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	last := GetTerminalWidth()
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			if w := GetTerminalWidth(); w != last {
				last = w
				onChange(w)
			}
		}
	}
}
