package internal_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm/internal"
)

// RandomStrings is shared test input for printer tests; the printers must be
// able to cope with every entry, so the collection has to keep covering the
// interesting shapes: plain text, special characters and different newline
// styles.
func TestRandomStringsCoverInterestingShapes(t *testing.T) {
	assert.NotEmpty(t, internal.RandomStrings)

	var hasLF, hasCRLF, hasSpecialChars bool

	for _, s := range internal.RandomStrings {
		assert.NotEmpty(t, s)

		hasLF = hasLF || strings.Contains(s, "\n")
		hasCRLF = hasCRLF || strings.Contains(s, "\r\n")
		hasSpecialChars = hasSpecialChars || strings.ContainsAny(s, "²³§€")
	}

	assert.True(t, hasLF, "RandomStrings should contain a multiline string")
	assert.True(t, hasCRLF, "RandomStrings should contain windows line endings")
	assert.True(t, hasSpecialChars, "RandomStrings should contain special characters")
}
