package internal

import (
	"strings"

	"github.com/gookit/color"
)

func AddTitleToLine(title, line string, length int, left bool) string {
	var ret string
	if left {
		ret += line + " " + title + " " + strings.Repeat(line, length-(3+len(color.ClearCode(title))))
	} else {
		ret += strings.Repeat(line, length-(3+len(color.ClearCode(title)))) + " " + title + " " + line
	}

	return ret
}
