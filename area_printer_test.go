package pterm_test

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/vt"
)

func TestAreaUpdatesInPlace(t *testing.T) {
	term := setup(t)

	area, err := pterm.Area.Start("waiting\nfor data")
	if err != nil {
		t.Fatal(err)
	}

	if got := term.Screen(); got != "waiting\nfor data" {
		t.Errorf("screen = %q", got)
	}

	area.Update("3 rows")
	pterm.Println("log line")

	if got := term.Screen(); got != "log line\n3 rows" {
		t.Errorf("screen = %q", got)
	}

	if area.Content() != "3 rows" {
		t.Errorf("content = %q", area.Content())
	}

	_ = area.Stop()
	area.Update("ignored")

	if got := term.Screen(); got != "log line\n3 rows" {
		t.Errorf("screen after stop = %q", got)
	}

	if term.CursorHidden() {
		t.Error("cursor still hidden")
	}
}

func TestAreaClearAndRemove(t *testing.T) {
	term := setup(t)

	area, _ := pterm.Area.WithRemoveWhenDone().Start("temporary")
	area.Clear()

	if got := term.Screen(); got != "" {
		t.Errorf("screen after clear = %q", got)
	}

	area.Update("back")
	_ = area.Stop()

	if got := term.Screen(); got != "" {
		t.Errorf("screen after stop = %q", got)
	}
}

func TestAreaNotTerminal(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		area, _ := pterm.Area.WithFullscreen().WithCenter().Start("first")
		area.Update("second")
		area.Update("final\n")
		_ = area.Stop()
	})
	if out != "final\n" {
		t.Errorf("output = %q", out)
	}
}

func TestAreaCenter(t *testing.T) {
	setup(t)

	term := vt.New(21, 10)
	area, _ := pterm.Area.WithWriter(term).WithCenter().Start("ab\nabcd")

	// 20 usable columns: the block is 4 wide, so it moves 8 to the right
	// and its lines stay aligned with each other.
	if got, want := term.Screen(), "        ab\n        abcd"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}

	_ = area.Stop()
}

func TestAreaFullscreen(t *testing.T) {
	setup(t)

	term := vt.New(20, 7)
	area, _ := pterm.Area.WithWriter(term).WithFullscreen().WithCenter().Start("hi")

	// Six rows are drawn; the line sits in the middle of them.
	lines := strings.Split(term.Raw(), "\n")
	if len(lines) != 6 || !strings.Contains(lines[2], "hi") {
		t.Errorf("rows = %q", lines)
	}

	area.Update("a\nb\nc\nd\ne\nf\ng\nh")

	// Content taller than the terminal keeps its last rows.
	if got := strings.ReplaceAll(term.Screen(), " ", ""); !strings.HasSuffix(got, "c\nd\ne\nf\ng\nh") {
		t.Errorf("screen = %q", got)
	}

	_ = area.Stop()
}

func TestAreaWraps(t *testing.T) {
	setup(t)

	term := vt.New(11, 10)
	area, _ := pterm.Area.WithWriter(term).Start("one two three four")

	if got := term.Screen(); got != "one two\nthree four" {
		t.Errorf("screen = %q", got)
	}

	_ = area.Stop()
}

func TestAreaContextStops(t *testing.T) {
	term := setup(t)

	ctx, cancel := context.WithCancel(context.Background())
	area, _ := pterm.Area.WithContext(ctx).Start("live")

	cancel()

	deadline := time.Now().Add(time.Second)
	for term.CursorHidden() && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}

	area.Update("ignored")

	if got := term.Screen(); got != "live" || term.CursorHidden() {
		t.Errorf("screen = %q, cursor hidden = %v", got, term.CursorHidden())
	}
}

func TestAreaConcurrent(t *testing.T) {
	term := setup(t)

	area, _ := pterm.Area.Start()

	var wg sync.WaitGroup

	for i := range 8 {
		wg.Go(func() {
			for range 50 {
				area.Update(strings.Repeat("x", i+1))
				pterm.Print("")

				_ = area.Content()
			}
		})
	}

	wg.Wait()
	area.Update("done")
	_ = area.Stop()

	if got := term.Screen(); got != "done" {
		t.Errorf("screen = %q", got)
	}
}
