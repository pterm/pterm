package pterm_test

import (
	"errors"
	"testing"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

var steps = []pterm.BulletListItem{
	{Text: "Measure what can be freed"},
	{Level: 1, Text: "npm and bun caches"},
	{Level: 1, Text: "Docker images"},
	{Text: "Clean up"},
}

func TestBulletListSnapshots(t *testing.T) {
	snap(t, "Default", srender(t, pterm.BulletList.WithItems(steps).Srender))
	snap(t, "Ordered", srender(t, pterm.BulletList.WithOrdered().WithItems(steps).Srender))
	snap(t, "Custom", srender(t, pterm.BulletList.WithItems([]pterm.BulletListItem{
		{Text: "done", Bullet: "✔", Tone: pterm.ToneSuccess},
		{Text: "plain"},
		{Text: "arrow", Bullet: ">"},
	}).WithBulletTone(pterm.ToneAccent).Srender))
}

func TestBulletListLook(t *testing.T) {
	setup(t)

	out := capture(t, func() { _ = pterm.BulletList.WithItems(steps).Render() })
	want := "• Measure what can be freed\n  • npm and bun caches\n  • Docker images\n• Clean up\n"

	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestBulletListOrderedNumbering(t *testing.T) {
	setup(t)

	items := make([]pterm.BulletListItem, 0, 12)
	for range 10 {
		items = append(items, pterm.BulletListItem{Text: "x"})
	}

	items = append(items, pterm.BulletListItem{Level: 1, Text: "child"}, pterm.BulletListItem{Text: "y"})

	got, _ := pterm.BulletList.WithOrdered().WithItems(items).Srender()
	want := " 1. x\n 2. x\n 3. x\n 4. x\n 5. x\n 6. x\n 7. x\n 8. x\n 9. x\n10. x\n    1. child\n11. y"

	if plain(got) != want {
		t.Errorf("got:\n%s\nwant:\n%s", plain(got), want)
	}
}

func TestBulletListWraps(t *testing.T) {
	setup(t)

	term := vt.New(20, 10)
	_ = pterm.BulletList.WithWriter(term).WithTexts("one two three four five six").Render()

	if want := "• one two three four\n  five six"; term.Screen() != want {
		t.Errorf("screen:\n%s\nwant:\n%s", term.Screen(), want)
	}
}

func TestBulletListEmpty(t *testing.T) {
	setup(t)

	if _, err := pterm.BulletList.Srender(); !errors.Is(err, pterm.ErrEmpty) {
		t.Errorf("err = %v, want ErrEmpty", err)
	}
}
