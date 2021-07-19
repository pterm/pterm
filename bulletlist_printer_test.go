package pterm_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestBulletListPrinterNilPrint(t *testing.T) {
	p := pterm.BulletListPrinter{}
	p.Render()
}

func TestBulletListPrinter_Render(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: fmt.Sprint(a)},
		}).Render()
	})
}

func TestBulletListPrinter_RenderWithoutStyle(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		pterm.BulletListPrinter{}.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: fmt.Sprint(a)},
		}).Render()
	})
}

func TestBulletListPrinter_RenderWithBullet(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a interface{}) {
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{
				Level:  0,
				Text:   fmt.Sprint(a),
				Bullet: "-",
			},
		}).Render()
	})
}

func TestBulletListPrinter_Srender(t *testing.T) {
	testSprintContainsWithoutError(t, func(a interface{}) (string, error) {
		return pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: fmt.Sprint(a)},
		}).Srender()
	})
}

func TestBulletListPrinter_WithBullet(t *testing.T) {
	p := pterm.BulletListPrinter{}
	p2 := p.WithBullet("-")

	assert.Equal(t, "-", p2.Bullet)
	assert.Empty(t, p.Bullet)
}

func TestBulletListPrinter_WithBulletStyle(t *testing.T) {
	p := pterm.BulletListPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithBulletStyle(s)

	assert.Equal(t, s, p2.BulletStyle)
	assert.Empty(t, p.BulletStyle)
}

func TestBulletListPrinter_WithItems(t *testing.T) {
	p := pterm.BulletListPrinter{}
	li := []pterm.BulletListItem{{
		Level:       0,
		Text:        "test",
		TextStyle:   nil,
		Bullet:      "+",
		BulletStyle: nil,
	}}
	p2 := p.WithItems(li)

	assert.Equal(t, li, p2.Items)
	assert.Empty(t, p.Items)
}

func TestBulletListPrinter_WithTextStyle(t *testing.T) {
	p := pterm.BulletListPrinter{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Empty(t, p.TextStyle)
}

func TestBulletListItem_WithBullet(t *testing.T) {
	p := pterm.BulletListItem{}
	p2 := p.WithBullet("-")

	assert.Equal(t, "-", p2.Bullet)
	assert.Empty(t, p.Bullet)
}

func TestBulletListItem_WithBulletStyle(t *testing.T) {
	p := pterm.BulletListItem{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithBulletStyle(s)

	assert.Equal(t, s, p2.BulletStyle)
	assert.Empty(t, p.BulletStyle)
}

func TestBulletListItem_WithLevel(t *testing.T) {
	p := pterm.BulletListItem{}
	p2 := p.WithLevel(1)

	assert.Equal(t, 1, p2.Level)
	assert.Empty(t, p.Level)
}

func TestBulletListItem_WithText(t *testing.T) {
	p := pterm.BulletListItem{}
	p2 := p.WithText("test")

	assert.Equal(t, "test", p2.Text)
	assert.Empty(t, p.Text)
}

func TestBulletListItem_WithTextStyle(t *testing.T) {
	p := pterm.BulletListItem{}
	s := pterm.NewStyle(pterm.FgRed, pterm.BgRed, pterm.Bold)
	p2 := p.WithTextStyle(s)

	assert.Equal(t, s, p2.TextStyle)
	assert.Empty(t, p.TextStyle)
}

func TestNewBulletListFromString(t *testing.T) {
	p := *pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
		{Level: 0, Text: "0"},
		{Level: 1, Text: "1"},
		{Level: 2, Text: "2"},
		{Level: 3, Text: "3"},
		{Level: 4, Text: "4"},
		{Level: 5, Text: "5"},
	})

	s := `0
 1
  2
   3
    4
     5`
	p2 := pterm.NewBulletListFromString(s, " ")

	assert.Equal(t, p, p2)
}
