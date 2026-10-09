package pterm_test

// Behavioral tests for SectionPrinter.
//
// Builder methods, Print*/Sprint* delegation and the global styling invariants
// are covered generically in contract_test.go. This file verifies the actual
// layout: indent characters per level, padding newlines and styling scope.

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestSectionPrinterDefaultLayout(t *testing.T) {
	// Level 1, one line of top and bottom padding, "»" indent character.
	assert.Equal(t, "\n» Title\n", stripANSI(pterm.DefaultSection.Sprint("Title")))
}

func TestSectionPrinterLevelRepeatsIndentCharacter(t *testing.T) {
	assert.Equal(t, "\n» Title\n", stripANSI(pterm.DefaultSection.WithLevel(1).Sprint("Title")))
	assert.Equal(t, "\n»» Title\n", stripANSI(pterm.DefaultSection.WithLevel(2).Sprint("Title")))
	assert.Equal(t, "\n»»» Title\n", stripANSI(pterm.DefaultSection.WithLevel(3).Sprint("Title")))
}

func TestSectionPrinterLevelZeroOmitsIndent(t *testing.T) {
	// Level 0 renders neither indent characters nor the separating space.
	assert.Equal(t, "\nTitle\n", stripANSI(pterm.DefaultSection.WithLevel(0).Sprint("Title")))
}

func TestSectionPrinterCustomIndentCharacter(t *testing.T) {
	p := pterm.DefaultSection.WithIndentCharacter("=>").WithLevel(2)

	assert.Equal(t, "\n=>=> Title\n", stripANSI(p.Sprint("Title")))
}

func TestSectionPrinterPaddingAddsExactNewlines(t *testing.T) {
	p := pterm.DefaultSection.WithTopPadding(3).WithBottomPadding(2)

	assert.Equal(t, "\n\n\n» Title\n\n", stripANSI(p.Sprint("Title")))
}

func TestSectionPrinterZeroPadding(t *testing.T) {
	p := pterm.DefaultSection.WithTopPadding(0).WithBottomPadding(0)

	assert.Equal(t, "» Title", stripANSI(p.Sprint("Title")))
}

func TestSectionPrinterStylesOnlyTheTitle(t *testing.T) {
	// The indent characters and the title are wrapped in the section style
	// (Bold=1, FgLightMagenta=95); the padding newlines stay unstyled.
	assert.Equal(t, "\n\x1b[1;95m»\x1b[0m \x1b[1;95mT\x1b[0m\n", pterm.DefaultSection.Sprint("T"))
}

func TestSectionPrinterZeroValueRendersBareText(t *testing.T) {
	p := pterm.SectionPrinter{}

	// No padding, level 0, no style: the text passes through unchanged.
	assert.Equal(t, "Title", p.Sprint("Title"))
}
