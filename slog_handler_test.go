package pterm_test

import (
	"context"
	"log/slog"
	"strings"
	"testing"

	"github.com/pterm/pterm"
)

func TestSlogHandler(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		logger := slog.New(pterm.NewSlogHandler(pterm.Logger.WithLevel(pterm.LogLevelDebug)))

		logger.Log(context.Background(), slog.LevelDebug-4, "trace hidden")
		logger.Debug("debug", "a", 1)
		logger.Info("info")
		logger.Warn("warn")
		logger.Error("error", "err", "boom")
		logger.Log(context.Background(), slog.LevelError+4, "above error")
	})

	want := `◦ debug debug a=1
● info  info
▲ warn  warn
✖ error error err=boom
✖ error above error
`
	if out != want {
		t.Errorf("output:\n%s\nwant:\n%s", out, want)
	}
}

func TestSlogHandlerAttrsAndGroups(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		logger := slog.New(pterm.NewSlogHandler(nil)).With("app", "shop").WithGroup("req").With("id", 7)
		logger.Info("handled", "status", 200, slog.Group("user", "name", "ada"), slog.Group("", "flat", true))
	})

	if want := "● info  handled app=shop req.id=7 req.status=200 req.user.name=ada req.flat=true\n"; out != want {
		t.Errorf("output = %q, want %q", out, want)
	}
}

func TestSlogHandlerFollowsLogger(t *testing.T) {
	setup(t)

	l := pterm.Logger.WithCaller()
	h := pterm.NewSlogHandler(l)

	if h.Enabled(context.Background(), slog.LevelDebug) {
		t.Error("debug enabled at info level")
	}

	l.Level = pterm.LogLevelDebug

	if !h.Enabled(context.Background(), slog.LevelDebug) {
		t.Error("level change on the logger not picked up")
	}

	out := capture(t, func() { slog.New(h).Debug("here") })

	// The caller comes from the record, so it points at this file and not
	// at log/slog.
	if !strings.Contains(out, "slog_handler_test.go:") {
		t.Errorf("output = %q", out)
	}
}
