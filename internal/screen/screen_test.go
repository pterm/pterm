package screen_test

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"testing"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/vt"
)

func TestWriteAboveRegion(t *testing.T) {
	t.Parallel()

	term := vt.New(40, 20)
	s := screen.For(term)

	r := s.NewRegion(screen.RegionOptions{})
	r.Set("⠋ loading")
	s.Write("first\n")
	r.Set("⠙ loading 1s")
	s.Write("second\n")

	if got, want := term.Screen(), "first\nsecond\n⠙ loading 1s"; got != want {
		t.Fatalf("screen:\n%s\nwant:\n%s", got, want)
	}

	if !term.CursorHidden() {
		t.Error("cursor visible while a region is drawn")
	}

	r.Close("✔ loaded")

	if got, want := term.Screen(), "first\nsecond\n✔ loaded"; got != want {
		t.Fatalf("screen:\n%s\nwant:\n%s", got, want)
	}

	if term.CursorHidden() {
		t.Error("cursor still hidden after the last region closed")
	}
}

func TestPartialLineWaitsForRegion(t *testing.T) {
	t.Parallel()

	term := vt.New(40, 20)
	s := screen.For(term)

	r := s.NewRegion(screen.RegionOptions{})
	r.Set("live")
	s.Write("par")
	s.Write("tial")

	if got, want := term.Screen(), "live"; got != want {
		t.Fatalf("screen = %q, want %q", got, want)
	}

	r.Close("")

	if got, want := term.Screen(), "partial"; got != want {
		t.Fatalf("screen = %q, want %q", got, want)
	}
}

func TestMultipleRegions(t *testing.T) {
	t.Parallel()

	term := vt.New(40, 20)
	s := screen.For(term)

	a := s.NewRegion(screen.RegionOptions{})
	b := s.NewRegion(screen.RegionOptions{})

	a.Set("a1\na2")
	b.Set("b1")
	a.Set("a1")
	a.Close("a done")
	b.Set("b2")

	if got, want := term.Screen(), "a done\nb2"; got != want {
		t.Fatalf("screen:\n%s\nwant:\n%s", got, want)
	}
}

func TestLongLinesAreCut(t *testing.T) {
	t.Parallel()

	term := vt.New(10, 20)
	s := screen.For(term)

	r := s.NewRegion(screen.RegionOptions{})
	r.Set(strings.Repeat("x", 30))
	r.Set("short")
	r.Close("done")

	if got, want := term.Screen(), "done"; got != want {
		t.Fatalf("screen = %q, want %q", got, want)
	}
}

func TestTallRegionIsCapped(t *testing.T) {
	t.Parallel()

	term := vt.New(20, 5)
	s := screen.For(term)

	r := s.NewRegion(screen.RegionOptions{})
	r.Set("1\n2\n3\n4\n5\n6\n7")
	r.Set("a\nb")
	r.Close("c")

	if got, want := term.Screen(), "c"; got != want {
		t.Fatalf("screen = %q, want %q", got, want)
	}
}

func TestNonTerminalPrintsFinalOnly(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	s := screen.For(&buf)
	r := s.NewRegion(screen.RegionOptions{})
	r.Set("frame 1")
	s.Write("log\n")
	r.Set("frame 2")
	r.Close("done\n")

	if got, want := buf.String(), "log\ndone\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestDecorator(t *testing.T) {
	t.Parallel()

	var buf bytes.Buffer

	s := screen.For(&buf)
	pop := s.PushDecorator(func(str string, _ bool) string {
		if strings.HasPrefix(str, screen.Node) {
			return str
		}

		return "| " + str
	})

	s.Write(screen.Node + "* node\n")
	s.Write("content\n")
	pop()
	s.Write("after\n")

	if got, want := buf.String(), "* node\n| content\nafter\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestConcurrentWrites(t *testing.T) {
	t.Parallel()

	term := vt.New(60, 40)
	s := screen.For(term)
	r := s.NewRegion(screen.RegionOptions{})

	var wg sync.WaitGroup

	for i := range 8 {
		wg.Go(func() {
			for j := range 50 {
				s.Write(fmt.Sprintf("goroutine %d line %d\n", i, j))
				r.Set(fmt.Sprintf("progress %d/%d", i, j))
			}
		})
	}

	wg.Wait()
	r.Close("finished")

	lines := strings.Split(term.Screen(), "\n")
	if len(lines) != 8*50+1 {
		t.Fatalf("got %d lines, want %d", len(lines), 8*50+1)
	}

	for _, line := range lines[:len(lines)-1] {
		if !strings.HasPrefix(line, "goroutine ") {
			t.Fatalf("torn line %q", line)
		}
	}

	if lines[len(lines)-1] != "finished" {
		t.Fatalf("last line = %q", lines[len(lines)-1])
	}
}
