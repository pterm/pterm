package pterm

import (
	"os"
	"strings"

	"github.com/pterm/pterm/internal"
	"golang.org/x/term"
)

func getWidth(s string) int {
	return min(GetTerminalWidth(), internal.GetStringMaxWidth(s))
}
func textFitWidth(text string) string {
	getWidth := internal.GetStringWidth
	w := GetTerminalWidth()
	if internal.GetStringMaxWidth(text) >= w {
		// find the last index that GetStringWidth(s[:index])<=width
		findIndex := func(s string, width int) int {
			l, r := 0, len(s)+1
			for l+1 < r {
				mid := (l + r) >> 1
				if getWidth(s[:mid]) <= width {
					l = mid
				} else {
					r = mid
				}
			}
			return l
		}
		buffer := make([]string, 0)
		for _, s := range strings.Split(text, "\n") {
			if getWidth(s) <= w {
				buffer = append(buffer, s)
				continue
			}
			for len(s) > 0 {
				i := findIndex(s, w)
				buffer = append(buffer, s[:i])
				s = s[i:]
			}
		}
		text = strings.Join(buffer, "\n")
	}
	return text
}
func inputFitWidth(ss []string) []string {
	getWidth := internal.GetStringWidth
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err == nil && internal.GetStringMaxWidth(strings.Join(ss, "\n")) > w {
		// find the last index that GetStringWidth(s[:index])<=width
		findIndex := func(s string, width int) int {
			l, r := 0, len(s)+1
			for l+1 < r {
				mid := (l + r) >> 1
				if getWidth(s[:mid]) <= width {
					l = mid
				} else {
					r = mid
				}
			}
			return l
		}
		buffer := make([]string, 0)
		for _, s := range ss {
			if getWidth(s) <= w {
				buffer = append(buffer, s)
				continue
			}
			for len(s) > 0 {
				i := findIndex(s, w)
				buffer = append(buffer, s[:i])
				s = s[i:]
			}
		}
		return buffer
	}
	return ss
}
