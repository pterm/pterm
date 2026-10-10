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

// quietBar is a bar without timers, so its frames are deterministic.
var quietBar = pterm.Progressbar.WithShowElapsedTime(false).WithShowETA(false)

func TestProgressbarSnapshots(t *testing.T) {
	frame := func(bar *pterm.ProgressbarPrinter, n int) func() string {
		return func() string {
			term := vt.New(80, 24)
			b, _ := bar.WithWriter(term).Start()
			b.Add(n)

			raw := term.Raw()
			_ = b.Stop()

			return raw + "\n" + strings.TrimPrefix(term.Raw(), raw)
		}
	}

	snap(t, "Running", frame(quietBar.WithTitle("Downloading"), 52))
	snap(t, "Empty", frame(quietBar.WithTitle("Downloading"), 0))
	snap(t, "Tone", frame(quietBar.WithTitle("Migrating").WithTone(pterm.ToneSuccess).WithBarWidth(20), 7))
	snap(t, "BarOnly", frame(quietBar.WithShowTitle(false).WithShowCount(false), 33))
}

func TestProgressbarRedrawsInPlace(t *testing.T) {
	term := setup(t)

	bar, err := quietBar.WithTitle("Downloading").Start()
	if err != nil {
		t.Fatal(err)
	}

	bar.Add(52)

	// 52% of 30 columns is 15.6: 15 full blocks, a five-eighths block (0.6
	// rounds to 5/8) and the track.
	want := "Downloading  " + strings.Repeat("█", 15) + "▋" + strings.Repeat("░", 14) + "   52%   52/100"
	if got := term.Screen(); got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}

	pterm.Println("log line")
	bar.Add(10)

	if got := term.Screen(); !strings.HasPrefix(got, "log line\nDownloading") || strings.Count(got, "\n") != 1 {
		t.Errorf("screen after log line:\n%s", got)
	}

	if !term.CursorHidden() {
		t.Error("cursor visible while the bar runs")
	}

	bar.Add(100)

	if got, want := term.Screen(), "log line\n✔ Downloading 100/100"; got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}

	if term.CursorHidden() {
		t.Error("cursor still hidden")
	}
}

func TestProgressbarTimers(t *testing.T) {
	term := setup(t)

	bar, _ := pterm.Progressbar.WithTotal(10).Start("Working")
	bar.Add(5)

	if got := term.Screen(); !strings.Contains(got, "s · eta ") {
		t.Errorf("screen has no elapsed time and ETA: %q", got)
	}

	_ = bar.Stop()

	got := term.Screen()
	if !strings.HasPrefix(got, "✔ Working 5/10 ") || !strings.HasSuffix(got, "s") {
		t.Errorf("final line = %q", got)
	}
}

func TestProgressbarUpdateTitle(t *testing.T) {
	term := setup(t)

	bar, _ := quietBar.WithTotal(4).Start("Installing react")
	bar.Increment()
	bar.UpdateTitle("Installing vite")

	// The title column keeps the width of the longest title, so the bar
	// stays where it was.
	if got := term.Screen(); !strings.HasPrefix(got, "Installing vite   █") {
		t.Errorf("screen = %q", got)
	}

	bar.Add(3)

	if got := term.Screen(); got != "✔ Installing vite 4/4" {
		t.Errorf("screen = %q", got)
	}
}

func TestProgressbarNotTerminal(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		bar, _ := quietBar.WithTotal(3).Start("Building")
		for range 3 {
			bar.Increment()
		}
	})
	if out != "✔ Building 3/3\n" {
		t.Errorf("output = %q", out)
	}
}

func TestProgressbarNarrowTerminal(t *testing.T) {
	setup(t)

	term := vt.New(40, 10)
	bar, _ := quietBar.WithWriter(term).Start("Downloading")
	bar.Add(50)

	line := term.Screen()
	if w := len([]rune(line)); w > 39 || !strings.Contains(line, "50%") || !strings.Contains(line, "█") {
		t.Errorf("line of %d columns: %q", w, line)
	}

	_ = bar.Stop()
}

func TestProgressbarContextCancels(t *testing.T) {
	term := setup(t)

	ctx, cancel := context.WithCancel(context.Background())
	bar, _ := quietBar.WithTotal(10).WithContext(ctx).Start("Uploading")
	bar.Add(2)
	cancel()

	deadline := time.Now().Add(time.Second)
	for term.Screen() != "✖ Uploading 2/10 canceled" && time.Now().Before(deadline) {
		time.Sleep(5 * time.Millisecond)
	}

	if got := term.Screen(); got != "✖ Uploading 2/10 canceled" {
		t.Errorf("screen = %q", got)
	}

	bar.Add(8) // no effect after cancel

	if got := term.Screen(); got != "✖ Uploading 2/10 canceled" {
		t.Errorf("screen after Add = %q", got)
	}
}

func TestProgressbarStopAfterCancel(t *testing.T) {
	setup(t)

	// The bar's own goroutine sees the canceled context while Stop runs.
	// Stop must still return only after the canceled line was printed.
	for range 50 {
		term := vt.New(80, 24)
		ctx, cancel := context.WithCancel(context.Background())

		bar, _ := quietBar.WithWriter(term).WithTotal(10).WithContext(ctx).Start("Uploading")

		cancel()

		_ = bar.Stop()

		if got := term.Screen(); got != "✖ Uploading 0/10 canceled" {
			t.Fatalf("screen = %q", got)
		}
	}
}

func TestProgressbarFailAndRemove(t *testing.T) {
	term := setup(t)

	bar, _ := quietBar.Start("Downloading")
	bar.Add(40)
	bar.Fail("Download failed")

	gone, _ := quietBar.WithRemoveWhenDone().Start("Temporary")
	gone.Add(100)

	if got := term.Screen(); got != "✖ Download failed 40/100" {
		t.Errorf("screen = %q", got)
	}
}

func TestProgressbarStartFull(t *testing.T) {
	term := setup(t)

	bar, _ := quietBar.WithTotal(5).WithCurrent(5).Start("Nothing to do")
	if current, total := bar.Progress(); current != 5 || total != 5 {
		t.Errorf("progress = %d/%d", current, total)
	}

	if got := term.Screen(); got != "✔ Nothing to do 5/5" {
		t.Errorf("screen = %q", got)
	}
}

func TestProgressbarASCII(t *testing.T) {
	term := setup(t)

	pterm.SetGlyphs(pterm.GlyphsASCII)

	defer pterm.SetGlyphs(pterm.GlyphsUnicode)

	bar, _ := quietBar.WithBarWidth(10).Start("Copying")
	bar.Add(55)

	if got, want := term.Screen(), "Copying  ######....   55%   55/100"; got != want {
		t.Errorf("screen = %q, want %q", got, want)
	}

	_ = bar.Stop()
}

func TestProgressbarConcurrent(t *testing.T) {
	term := setup(t)

	bar, _ := pterm.Progressbar.WithTotal(800).Start("Hashing")

	var wg sync.WaitGroup

	for range 8 {
		wg.Go(func() {
			for range 100 {
				bar.Increment()
				pterm.Print("")
			}
		})
	}

	wg.Go(func() {
		for range 20 {
			bar.UpdateTitle("Hashing files")
			pterm.Println("progress", bar.Elapsed() >= 0)
		}
	})

	wg.Wait()

	if current, _ := bar.Progress(); current != 800 {
		t.Errorf("current = %d", current)
	}

	lines := strings.Split(term.Screen(), "\n")
	if last := lines[len(lines)-1]; !strings.HasPrefix(last, "✔ Hashing files 800/800") {
		t.Errorf("last line = %q", last)
	}
}

func TestProgressbarOutputDisabled(t *testing.T) {
	term := setup(t)

	pterm.DisableOutput()

	defer pterm.EnableOutput()

	bar, _ := quietBar.Start("Silent")
	bar.Add(100)

	if got := term.Raw(); got != "" {
		t.Errorf("output = %q", got)
	}
}
