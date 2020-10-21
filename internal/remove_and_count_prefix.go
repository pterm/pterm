package internal

import (
	"strings"
)

func RemoveAndCountPrefix(input, subString string) (string, int) {
	inputLength := len(input)
	input = strings.TrimLeft(input, subString)
	return input, inputLength - len(input)
}
