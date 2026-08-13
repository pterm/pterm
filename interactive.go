package pterm

import (
	"strings"
	"time"

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
	if internal.GetStringMaxWidth(strings.Join(ss, "\n")) >= w {
		// find the last index that GetStringWidth(s[:index])<=width
		findIndex := func(s string, width int) int {
			l, r := 0, len(s)+1
			for l+1 < r {
				mid := (l + r) >> 1
				// if getMaxW(string([]rune(s)[:mid])) <= width {
				if getMaxW(s[:mid]) <= width {
					l = mid
				} else {
					r = mid
				}
			}
			return l
		}
		buffer := make([]string, 0)
		for _, s := range ss {
			if getMaxW(s) <= w {
				buffer = append(buffer, s)
				continue
			}
			for len(s) > 0 {
				i := findIndex(s, w)
				// buffer = append(buffer, string([]rune(s)[:i]))
				// s = string([]rune(s)[i:])
				buffer = append(buffer, s[:i])
				s = s[i:]
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
