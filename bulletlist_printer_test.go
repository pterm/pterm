package pterm_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestBulletListPrinterNilPrint(t *testing.T) {
	p := pterm.BulletListPrinter{}
	p.Render()
}

func TestBulletListPrinter_WithMethods(t *testing.T) {
	testWithMethods(t, pterm.BulletListPrinter{})
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

	testza.AssertEqual(t, p, p2)
}
