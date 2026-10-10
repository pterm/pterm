package pterm_test

import (
	"errors"
	"testing"

	"github.com/pterm/pterm"
)

func TestTextSnapshots(t *testing.T) {
	snap(t, "Default", func() string { return pterm.Text.Sprint("Plain text in the text tone") })
	snap(t, "Muted", func() string { return pterm.Text.WithTone(pterm.ToneMuted).Sprint("A muted hint") })
	snap(t, "Wrapped", func() string {
		return pterm.Text.Sprint("This line is long enough that it has to wrap onto a second line on an eighty column terminal.")
	})
}

func TestTextNoWrap(t *testing.T) {
	setup(t)

	long := "This line is long enough that it would wrap onto a second line on an eighty column terminal."
	if got := pterm.Text.WithNoWrap().Sprint(long); pterm.RemoveColorFromString(got) != long {
		t.Errorf("NoWrap changed the text: %q", got)
	}
}

func TestTextPrint(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Text.Print("a")
		pterm.Text.Println("b", "c")
		pterm.Text.Printf("%d", 1)
		pterm.Text.Printfln("%d", 2)
		pterm.Text.PrintOnError(nil, errors.New("boom"))
	})
	if out != "ab c\n12\nboom\n" {
		t.Errorf("output = %q", out)
	}
}

func TestTextUnknownWidthDoesNotWrap(t *testing.T) {
	setup(t)
	t.Setenv("COLUMNS", "")

	long := "This line is long enough that it would wrap onto a second line on an eighty column terminal."

	out := capture(t, func() { pterm.Text.Println(long) })
	if out != long+"\n" {
		t.Errorf("output = %q", out)
	}
}
