package internal

import (
	"strings"
)

// RemoveAndCountPrefix removes the prefix from the input string and returns
// the trimmed string and the number of leading occurrences.
func RemoveAndCountPrefix(input, subString string) (string, int) {
	inputLength := len(input)
	input = strings.TrimLeft(input, subString)

	return input, inputLength - len(input)
}
