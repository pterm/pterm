package internal

import (
	"strings"

	"github.com/gookit/color"
)

// CenterText returns a centered string with a padding left and right
// If width is 0, it will be calculated automatically
func CenterText(text string, width int) string {
	var lines []string
	if width == 0 {
		width = GetStringMaxWidth(text)
	}
	linesTmp := strings.Split(text, "\n")
	for _, line := range linesTmp {
		if len(color.ClearCode(line)) > width {
			extraLines := []string{""}
			extraLinesCounter := 0
			for i, letter := range line {
				if i%width == 0 && i != 0 {
					extraLinesCounter++
					extraLines = append(extraLines, "")
				}
				extraLines[extraLinesCounter] += string(letter)
			}
			for _, extraLine := range extraLines {
				padding := width - len(color.ClearCode(extraLine))
				extraLine = strings.Repeat(" ", padding/2) + extraLine + strings.Repeat(" ", padding/2) + "\n"
				lines = append(lines, extraLine)
			}
		} else {
			padding := width - len(color.ClearCode(line))
			line = strings.Repeat(" ", padding/2) + line + strings.Repeat(" ", padding/2) + "\n"
			lines = append(lines, line)
		}
	}

	var line string
	for _, s := range lines {
		line += s
	}

	return strings.TrimSuffix(line, "\n")
}
