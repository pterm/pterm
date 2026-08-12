package pterm

import (
	"strings"

	"github.com/pterm/pterm/internal"
)

func getMaxW(s string) int {
	return internal.GetStringMaxWidth(s)
}

// fit the text for terminal width
func textFitWidth(text string) string {
	return strings.Join(linesFitWidth(strings.Split(text, "\n")), "\n")
}

func linesFitWidth(ss []string) []string {
	getWidth := internal.GetStringWidth
	w := GetTerminalWidth()
	if internal.GetStringMaxWidth(strings.Join(ss, "\n")) > w {
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

// returns the coords typed [][2]int that index is the actual line number in terminal
// the param ss should be the []string that fit the terminal width
// areaInput := [inputFitWidth](p.input)
// if y,x is the actual coords in terminal,y ∈ [0,len(areaInput)), x ∈[-len(areaInput[y]),0]
// p.cursorYPos,p.cursorXPos:=coords[y][0],coords[y][1]+x
func inputCoords(ss []string) [][2]int {
	coords := [][2]int{}
	for py, s := range ss {
		px := -getMaxW(s)
		for _, line := range linesFitWidth([]string{s}) {
			px += getMaxW(line)
			coords = append(coords, [2]int{py, px})
		}
	}
	return coords
}

// returns the coords in logic input
func logicYX(ay, ax int, coords [][2]int) (int, int) {
	return coords[ay][0], coords[ay][1] + ax
}

// returns actual coord in terminal,if not returns itself
func actualYX(py, px int, coords [][2]int) (int, int) {
	for ay, c := range coords {
		if c[0] == py && c[1] >= px {
			return ay, px - c[1]
		}
	}
	return py, px
}
