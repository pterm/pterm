package pterm_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/internal/screen"
)

// sizedBuffer is a buffer with a known width that is not a terminal, so
// log lines keep their colors and wrap like on a terminal of that width.
type sizedBuffer struct {
	bytes.Buffer

	width int
}

func (b *sizedBuffer) IsTerminal() bool          { return false }
func (b *sizedBuffer) Size() (width, height int) { return b.width, 24 }

// logTo returns what fn logs with a logger that writes to a buffer of the
// given width.
func logTo(width int, fn func(l *pterm.LoggerPrinter)) string {
	buf := &sizedBuffer{width: width}
	fn(pterm.Logger.WithWriter(buf).WithLevel(pterm.LogLevelTrace))

	return strings.TrimSuffix(buf.String(), "\n")
}

func fixClock(t *testing.T) {
	t.Helper()

	restore := pterm.SetLogClock(func() time.Time { return time.Date(2026, 10, 10, 14, 3, 9, 0, time.UTC) })
	t.Cleanup(restore)
}

func TestLoggerSnapshots(t *testing.T) {
	fixClock(t)

	levels := func(l *pterm.LoggerPrinter) {
		l.Trace("Reading config", l.Args("path", "./config.yaml"))
		l.Debug("Config loaded", l.Args("workers", 4))
		l.Info("Server started", l.Args("addr", ":8080", "env", "production"))
		l.Warn("Slow request", l.Args("path", "/api/search", "took", 1240*time.Millisecond))
		l.Error("Payment failed", l.Args("order", 1042, "error", errors.New("card declined")))
		l.Print("Plain line", l.Args("n", 1))
	}

	snap(t, "Levels", func() string { return logTo(80, levels) })
	snap(t, "Time", func() string {
		return logTo(80, func(l *pterm.LoggerPrinter) { l.WithTime().Info("Job queued", l.Args("job", "nightly")) })
	})
	snap(t, "Block", func() string {
		return logTo(60, func(l *pterm.LoggerPrinter) {
			l.Info("Retrying the upload of the nightly backup to the secondary region",
				l.Args("bucket", "s3://backups", "reason", "connection reset", "trace", "line one\nline two"))
		})
	})
	snap(t, "KeyStyles", func() string {
		return logTo(80, func(l *pterm.LoggerPrinter) {
			l.AppendKeyStyle("user", pterm.ToneAccent.Style()).Info("Signed in", l.Args("user", "ada", "err", "none"))
		})
	})
}

func TestLoggerPlainLines(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		l := pterm.Logger.WithLevel(pterm.LogLevelTrace)
		l.Trace("trace")
		l.Debug("debug")
		l.Info("info", l.Args("quoted", "two words", "empty", "", "eq", "a=b"))
		l.Warn("warn")
		l.Error("error")
		l.Print("print")
	})

	want := `◦ trace trace
◦ debug debug
● info  info quoted="two words" empty="" eq="a=b"
▲ warn  warn
✖ error error
•       print
`
	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestLoggerLevels(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Logger.Trace("hidden")
		pterm.Logger.Debug("hidden")
		pterm.Logger.Info("shown")
		pterm.Logger.WithLevel(pterm.LogLevelError).Warn("hidden")
		pterm.Logger.WithLevel(pterm.LogLevelDisabled).Print("hidden")
		pterm.Logger.WithLevel(pterm.LogLevelError).Print("print shows")
	})

	if out != "● info  shown\n•       print shows\n" {
		t.Errorf("output = %q", out)
	}

	if pterm.Logger.CanPrint(pterm.LogLevelDebug) || !pterm.Logger.CanPrint(pterm.LogLevelWarn) {
		t.Error("CanPrint does not follow the level")
	}
}

func TestLoggerWritesToErrorOutput(t *testing.T) {
	setup(t)

	var stdout, stderr bytes.Buffer

	pterm.SetDefaultOutput(&stdout)
	pterm.SetErrorOutput(&stderr)

	pterm.Logger.Info("to stderr")

	if stdout.Len() != 0 || !strings.Contains(stderr.String(), "to stderr") {
		t.Errorf("stdout = %q, stderr = %q", stdout.String(), stderr.String())
	}
}

func TestLoggerArgs(t *testing.T) {
	args := pterm.Logger.Args("a", 1, "orphan")
	if len(args) != 2 || args[1].Key != pterm.ErrKeyWithoutValue || args[1].Value != "orphan" {
		t.Errorf("args = %v", args)
	}

	m := pterm.Logger.ArgsFromMap(map[string]any{"b": 2, "a": 1, "c": 3})
	if m[0].Key != "a" || m[1].Key != "b" || m[2].Key != "c" {
		t.Errorf("map args not sorted: %v", m)
	}
}

func TestLoggerSortArguments(t *testing.T) {
	out := logTo(0, func(l *pterm.LoggerPrinter) {
		l = l.WithSortArguments()
		l.Info("sorted", l.Args("z", 1, "a", 2))
	})

	if !strings.HasSuffix(pterm.RemoveColorFromString(out), "sorted a=2 z=1") {
		t.Errorf("output = %q", out)
	}
}

func TestLoggerUnknownWidthStaysOnOneLine(t *testing.T) {
	out := logTo(0, func(l *pterm.LoggerPrinter) {
		l.Info(strings.Repeat("long ", 40), l.Args("value", "a\nb"))
	})

	if strings.Contains(out, "\n") || !strings.Contains(out, `value="a\nb"`) {
		t.Errorf("output = %q", out)
	}
}

func TestLoggerJSON(t *testing.T) {
	setup(t)
	fixClock(t)

	var buf bytes.Buffer

	l := pterm.Logger.WithWriter(&buf).WithFormatter(pterm.LogFormatterJSON).WithTimeFormat(time.RFC3339)
	l.Warn("Slow request", l.Args("ms", 1240, "error", errors.New("timeout")))

	var got map[string]any
	if err := json.Unmarshal(buf.Bytes(), &got); err != nil {
		t.Fatalf("invalid JSON %q: %v", buf.String(), err)
	}

	want := map[string]any{"level": "WARN", "msg": "Slow request", "timestamp": "2026-10-10T14:03:09Z", "ms": 1240.0, "error": "timeout"}
	for k, v := range want {
		if got[k] != v {
			t.Errorf("%s = %v, want %v", k, got[k], v)
		}
	}
}

func TestLoggerCaller(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Logger.WithCaller().Info("here")
	})

	if !strings.Contains(out, "caller=pterm/logger_test.go:") && !strings.Contains(out, "/logger_test.go:") {
		t.Errorf("output = %q", out)
	}

	wrapped := func(l *pterm.LoggerPrinter) { l.Info("wrapped") }

	out = capture(t, func() {
		wrapped(pterm.Logger.WithCaller().WithCallerOffset(1))
	})

	if !strings.Contains(out, "logger_test.go:") {
		t.Errorf("output with offset = %q", out)
	}
}

func TestLoggerFatalExits(t *testing.T) {
	setup(t)

	code := -1
	exit := screen.Exit
	screen.Exit = func(c int) { code = c }

	defer func() { screen.Exit = exit }()

	out := capture(t, func() { pterm.Logger.Fatal("boom", pterm.Logger.Args("code", 7)) })

	if out != "✖ fatal boom code=7\n" || code != 1 {
		t.Errorf("output = %q, exit code = %d", out, code)
	}
}

func TestLoggerLabels(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		pterm.Logger.WithLabel(pterm.LogLevelInfo, pterm.LevelLabel{Word: "note", Tone: pterm.ToneAccent}).Info("custom")
	})

	if out != "● note  custom\n" {
		t.Errorf("output = %q", out)
	}

	if _, ok := pterm.Logger.Labels[pterm.LogLevelInfo]; ok {
		t.Error("WithLabel changed the package logger")
	}
}

func TestLoggerAppendKeyStyleDoesNotShare(t *testing.T) {
	_ = pterm.Logger.AppendKeyStyle("user", pterm.ToneAccent.Style())

	if _, ok := pterm.Logger.KeyStyles["user"]; ok {
		t.Error("AppendKeyStyle changed the package logger")
	}
}

func TestLoggerAboveLivePrinters(t *testing.T) {
	term := setup(t)

	bar, _ := quietBar.WithTotal(2).Start("Working")

	pterm.Logger.Info("step", pterm.Logger.Args("n", 1))
	bar.Add(2)

	if got, want := term.Screen(), "● info  step n=1\n✔ Working 2/2"; got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}
}

func TestLoggerConcurrent(t *testing.T) {
	term := setup(t)

	var wg sync.WaitGroup

	for i := range 8 {
		wg.Go(func() {
			for range 25 {
				pterm.Logger.Info("tick", pterm.Logger.Args("worker", i))
			}
		})
	}

	wg.Wait()

	if n := strings.Count(term.Screen(), "● info  tick worker="); n != 200 {
		t.Errorf("%d lines", n)
	}
}
