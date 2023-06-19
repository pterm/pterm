package internal

import (
	"strings"
)

// AddTitleToLine adds a title to a site of a line ex: "─ This is the title ──────"
func AddTitleToLine(title, line string, length int, left bool) string {
	var ret string
	if left {
		ret += line + " " + title + " " + line + strings.Repeat(line, length-(4+GetStringMaxWidth(title)))
	} else {
		ret += strings.Repeat(line, length-(4+GetStringMaxWidth(title))) + line + " " + title + " " + line
	}

	return ret
}

// AddTitleToLineCenter adds a title to the center of a line ex: "─ This is the title ──────"
func AddTitleToLineCenter(title, line string, length int) string {
	var ret string
	repeatString := length - (4 + GetStringMaxWidth(title))
	unevenRepeatString := repeatString % 2

	ret += strings.Repeat(line, repeatString/2) + line + " " + title + " " + line + strings.Repeat(line, repeatString/2+unevenRepeatString)

	return ret
}
