package pterm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

// The default checkmark must render the well-known glyphs (the styling around
// them depends on the environment at package init, the glyphs must not).
func TestThemeDefaultCheckmarkGlyphs(t *testing.T) {
	assert.Equal(t, "✓", stripANSI(pterm.ThemeDefault.Checkmark.Checked))
	assert.Equal(t, "✗", stripANSI(pterm.ThemeDefault.Checkmark.Unchecked))
}

// Bar is not part of the printer builder contract (it is a data atom), so its
// With* methods are verified here: chaining must accumulate all fields and
// never mutate the value it was derived from.
func TestBarBuilderMethods(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	labelStyle := pterm.NewStyle(pterm.FgCyan, pterm.Bold)

	original := pterm.Bar{}
	modified := original.
		WithLabel("cpu").
		WithValue(42).
		WithStyle(style).
		WithLabelStyle(labelStyle)

	assert.Equal(t, pterm.Bar{Label: "cpu", Value: 42, Style: style, LabelStyle: labelStyle}, *modified)
	assert.Equal(t, pterm.Bar{}, original, "With* must not mutate the original Bar")
}
