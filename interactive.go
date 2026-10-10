package pterm

import (
	"context"
	"io"
	"os"
	"slices"
	"sync"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/text"
)

// promptConfig is what every interactive printer shares: where it reads
// keys, where it draws, and what ends it early.
type promptConfig struct {
	writer      io.Writer
	input       io.Reader
	ctx         context.Context
	onInterrupt func()
}

// prompt drives one interactive printer: it draws a view in a live region,
// feeds keys to the printer until it is done, and collapses the region into
// a status line.
type prompt[T any] struct {
	cfg promptConfig

	// view renders the prompt while it waits for input.
	view func(c canvas) string
	// key handles one key press. It reports whether the prompt is done.
	key func(k keys.Key) (done bool)
	// result returns the answer once key reported done.
	result func() T
	// collapse renders the status line the prompt leaves behind.
	collapse func(c canvas, st status) string
}

// inputMu makes sure only one prompt reads keys at a time. Two prompts
// fighting over the terminal would each get half of the key presses.
var inputMu sync.Mutex

// stdinIsTerminal reports whether keys can be read from the terminal.
// Tests replace it.
var stdinIsTerminal = func() bool { return screen.IsTerminal(os.Stdin) }

// listenKeys reads key presses from the terminal. Tests replace it to feed
// key presses the way a terminal groups them, such as pasted text.
var listenKeys = keyboard.Listen

// viewHook receives every view a prompt draws. Tests use it to look at a
// prompt while it waits.
var viewHook func(view string)

// run shows the prompt and returns the answer. When standard input is not
// a terminal and the printer reads no other input, it returns fallback and
// ErrNotInteractive, or fallback and nil when hasDefault is set.
func (p prompt[T]) run(fallback T, hasDefault bool) (T, error) {
	if p.cfg.input == nil && !stdinIsTerminal() {
		if hasDefault {
			return fallback, nil
		}

		return fallback, ErrNotInteractive
	}

	inputMu.Lock()
	defer inputMu.Unlock()

	w := orDefault(p.cfg.writer)
	ctx := orBackground(p.cfg.ctx)

	var region *screen.Region
	if OutputEnabled() {
		region = screen.For(w).NewRegion(screen.RegionOptions{
			OnInterrupt: func() string { return p.collapse(canvasFor(w), statusCanceled) },
			Handler:     p.cfg.onInterrupt,
		})
	}

	draw := func() {
		if region == nil {
			return
		}

		view := p.view(canvasFor(w))
		if viewHook != nil {
			viewHook(view)
		}

		region.Set(view)
	}

	draw()

	var (
		interrupted bool
		canceled    bool
	)

	onKey := func(k keys.Key) (bool, error) {
		if ctx.Err() != nil {
			canceled = true
			return true, nil //nolint:nilerr // the canceled flag carries the outcome.
		}

		if k.Code == keys.CtrlC {
			interrupted = true
			return true, nil
		}

		if p.key(k) {
			return true, nil
		}

		draw()

		return false, nil
	}

	var err error

	if p.cfg.input != nil {
		err = keyboard.ListenReader(p.cfg.input, onKey)
	} else {
		stop := context.AfterFunc(ctx, func() {
			// Wake the listener so it sees the canceled context.
			_ = keyboard.SimulateKeyPress(keys.Null)
		})
		err = listenKeys(onKey)

		stop()
	}

	var zero T

	switch {
	case interrupted:
		p.finish(region, w, statusCanceled)

		if p.cfg.onInterrupt == nil {
			screen.Interrupted()
		} else {
			p.cfg.onInterrupt()
		}

		return zero, ErrInterrupted

	case canceled:
		p.finish(region, w, statusCanceled)
		return zero, ErrCanceled
	case err != nil:
		p.finish(region, w, statusCanceled)
		return zero, err
	}

	p.finish(region, w, statusSuccess)

	return p.result(), nil
}

func (p prompt[T]) finish(region *screen.Region, w io.Writer, st status) {
	if region != nil {
		region.Close(p.collapse(canvasFor(w), st))
	}
}

// question renders the first line of a prompt: the question glyph, the
// question in bold, and a muted hint.
func question(c canvas, text, hint string) string {
	line := c.paint(TonePrimary, c.g.Question) + " " + c.bold(ToneText, text)
	if hint != "" {
		line += " " + c.paint(ToneMuted, hint)
	}

	return node(line)
}

// answered renders the status line an answered prompt leaves behind.
func answered(c canvas, st status, text, answer string) string {
	line := st.glyph(c) + " " + c.paint(ToneText, text)

	switch {
	case st == statusCanceled:
		line += " " + c.paint(ToneMuted, "canceled")
	case answer != "":
		line += " " + c.paint(ToneAccent, answer)
	}

	return node(line)
}

// keyHint is a key and what it does, shown in the footer of a prompt.
type keyHint struct {
	key    string
	action string
	// drop orders the hints that go when the footer is too wide: higher
	// goes first. Zero never goes.
	drop int
}

// hints renders the key hint footer of a prompt with the KeyHints
// printer: [space] toggle · [enter] confirm. It never cuts the line;
// footer decides which hints fit.
func hints(c canvas, hs ...keyHint) string {
	kh := make([]KeyHint, len(hs))
	for i, h := range hs {
		kh[i] = KeyHint{Key: h.key, Desc: h.action}
	}

	c.width = 0

	return KeyHints.WithHints(kh...).render(c)
}

// footer renders the key hints of a prompt as its last line. When they do
// not fit the width, the least important hints are left out.
func footer(c canvas, hs ...keyHint) string {
	line := promptIndent + hints(c, hs...)

	for c.width > 0 && text.Width(line) > c.width-1 {
		worst := -1

		for i, h := range hs {
			if h.drop > 0 && (worst < 0 || h.drop >= hs[worst].drop) {
				worst = i
			}
		}

		if worst < 0 {
			break
		}

		hs = slices.Delete(slices.Clone(hs), worst, worst+1)
		line = promptIndent + hints(c, hs...)
	}

	return line
}

// promptIndent indents the body of a prompt under its question.
const promptIndent = "  "
