package pterm_test

// Behavioral tests for BulletListPrinter: indentation math, bullet selection,
// multiline continuation lines and per-item styling. The builder/contract
// plumbing is covered in contract_test.go, one representative output in
// snapshot_test.go.

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func TestBulletListPrinter_IndentationAndBullets(t *testing.T) {
	printer := pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "zero"},
		{Level: 1, Text: "one", Bullet: "-"},
		{Level: 3, Text: "three"},
	})

	// Indentation is exactly two spaces per Level; items without their own
	// bullet use the printer's default bullet.
	expected := "" +
		"• zero\n" +
		"  - one\n" +
		"      • three\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBulletListPrinter_CustomPrinterBullet(t *testing.T) {
	printer := pterm.DefaultBulletList.WithBullet(">").WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "a"},
		{Level: 0, Text: "b", Bullet: "*"}, // item bullet wins over printer bullet
	})

	expected := "" +
		"> a\n" +
		"* b\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBulletListPrinter_MultilineItemContinuationLines(t *testing.T) {
	printer := pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 1, Text: "first\nsecond"},
	})

	// Continuation lines keep the item's indentation and align under the
	// text, not under the bullet.
	expected := "" +
		"  • first\n" +
		"    second\n"

	assert.Equal(t, expected, srenderPlain(t, printer))
}

func TestBulletListPrinter_PerItemStyles(t *testing.T) {
	textStyle := pterm.NewStyle(pterm.FgRed)
	bulletStyle := pterm.NewStyle(pterm.FgCyan)
	printer := pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "styled", TextStyle: textStyle, BulletStyle: bulletStyle},
		{Level: 0, Text: "unstyled"},
	})

	styled, err := printer.Srender()
	require.NoError(t, err)

	assert.Contains(t, styled, textStyle.Sprint("styled"), "item text must use the item's TextStyle")
	assert.Contains(t, styled, bulletStyle.Sprint("•"), "bullet must use the item's BulletStyle")
	assert.NotContains(t, styled, textStyle.Sprint("unstyled"), "other items must not inherit the style")
}

func TestBulletListPrinter_PrinterStylesAreItemFallback(t *testing.T) {
	printerStyle := pterm.NewStyle(pterm.FgGreen)
	itemStyle := pterm.NewStyle(pterm.FgRed)
	printer := pterm.DefaultBulletList.WithTextStyle(printerStyle).WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "inherits"},
		{Level: 0, Text: "own style", TextStyle: itemStyle},
	})

	styled, err := printer.Srender()
	require.NoError(t, err)

	assert.Contains(t, styled, printerStyle.Sprint("inherits"), "items without a style must use the printer style")
	assert.Contains(t, styled, itemStyle.Sprint("own style"), "an item's own style must win over the printer style")
}

func TestNewBulletListFromString(t *testing.T) {
	expected := *pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "0"},
		{Level: 1, Text: "1"},
		{Level: 2, Text: "2"},
	})

	assert.Equal(t, expected, pterm.NewBulletListFromString("0\n 1\n  2", " "))
}

func TestBulletListPrinter_SrenderIsPure(t *testing.T) {
	items := []pterm.BulletListItem{
		{Level: 0, Text: "a"},
		{Level: 1, Text: "b"},
	}
	printer := pterm.DefaultBulletList.WithItems(items)

	first, err := printer.Srender()
	require.NoError(t, err)

	second, err := printer.Srender()
	require.NoError(t, err)

	assert.Equal(t, first, second, "rendering twice must yield identical output")
	assert.Equal(t, []pterm.BulletListItem{
		{Level: 0, Text: "a"},
		{Level: 1, Text: "b"},
	}, items, "rendering must not modify the input items (e.g. write back styles)")
}

// BulletListItem is a builder-style helper type that is not part of the
// printer list in contract_test.go, so its With* methods are verified here.
func TestBulletListItem_Builders(t *testing.T) {
	style := pterm.NewStyle(pterm.FgRed)
	item := pterm.BulletListItem{}.
		WithLevel(2).
		WithText("text").
		WithBullet("-").
		WithTextStyle(style).
		WithBulletStyle(style)

	assert.Equal(t, pterm.BulletListItem{
		Level:       2,
		Text:        "text",
		Bullet:      "-",
		TextStyle:   style,
		BulletStyle: style,
	}, *item)
}
