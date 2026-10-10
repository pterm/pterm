package pterm

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/pterm/pterm/internal/screen"
)

// SpinnerPrinter shows that something is running. It ends as a status line
// with the time it took:
//
//	⠹ Downloading assets 1.2s
//	✔ Downloaded assets 3.4s
type SpinnerPrinter struct {
	// Text describes what is running.
	Text string
	// Tone colors the spinner frames.
	Tone Tone
	// Frames are drawn in turn. Nil uses the spinner of the glyph set.
	Frames []string
	// Delay is the time between two frames.
	Delay time.Duration
	// ShowTimer adds the elapsed time, muted, after the text.
	ShowTimer bool
	// RemoveWhenDone removes the spinner when it stops instead of leaving a
	// status line.
	RemoveWhenDone bool
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Context stops the spinner as canceled when it ends.
	Context context.Context
	// OnInterrupt runs when the user presses Ctrl+C while the spinner runs,
	// after the terminal was restored. Nil exits the program with code 130.
	OnInterrupt func()

	run *spinnerRun
}

// spinnerRun is the state of a started spinner. Copies of a started
// SpinnerPrinter share it.
type spinnerRun struct {
	mu      sync.Mutex
	text    string
	started time.Time
	frame   int
	region  *screen.Region
	done    chan struct{}
	// finished prints the final line once. Callers that end the spinner at
	// the same time wait until it is printed, so a Stop right before the
	// program exits never loses it.
	finished sync.Once
}

// Spinner is the default spinner.
var Spinner = SpinnerPrinter{
	Tone:      TonePrimary,
	Delay:     80 * time.Millisecond,
	ShowTimer: true,
}

// WithText returns a copy of p with text t.
func (p SpinnerPrinter) WithText(t string) *SpinnerPrinter {
	p.Text = t
	return &p
}

// WithTone returns a copy of p whose frames use tone t.
func (p SpinnerPrinter) WithTone(t Tone) *SpinnerPrinter {
	p.Tone = t
	return &p
}

// WithFrames returns a copy of p that draws frames in turn.
func (p SpinnerPrinter) WithFrames(frames ...string) *SpinnerPrinter {
	p.Frames = frames
	return &p
}

// WithDelay returns a copy of p that waits d between frames.
func (p SpinnerPrinter) WithDelay(d time.Duration) *SpinnerPrinter {
	p.Delay = d
	return &p
}

// WithShowTimer returns a copy of p that shows or hides the elapsed time.
func (p SpinnerPrinter) WithShowTimer(show ...bool) *SpinnerPrinter {
	p.ShowTimer = flag(show)
	return &p
}

// WithRemoveWhenDone returns a copy of p that disappears when it stops.
func (p SpinnerPrinter) WithRemoveWhenDone(remove ...bool) *SpinnerPrinter {
	p.RemoveWhenDone = flag(remove)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p SpinnerPrinter) WithWriter(w io.Writer) *SpinnerPrinter {
	p.Writer = w
	return &p
}

// WithContext returns a copy of p that stops as canceled when ctx ends.
func (p SpinnerPrinter) WithContext(ctx context.Context) *SpinnerPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting.
func (p SpinnerPrinter) WithOnInterrupt(fn func()) *SpinnerPrinter {
	p.OnInterrupt = fn
	return &p
}

// Start shows the spinner with the given text, or with p.Text when none is
// given, and returns the running spinner.
func (p SpinnerPrinter) Start(text ...any) (*SpinnerPrinter, error) {
	if len(text) > 0 {
		p.Text = sprintln(text...)
	}

	if p.Delay <= 0 {
		p.Delay = Spinner.Delay
	}

	run := &spinnerRun{text: p.Text, started: time.Now(), done: make(chan struct{})}
	p.run = run

	if OutputEnabled() {
		run.region = screen.For(p.out()).NewRegion(screen.RegionOptions{
			OnInterrupt: func() string { return p.final(canvasFor(p.out()), statusCanceled, "") },
			Handler:     p.OnInterrupt,
		})
	}

	p.draw()

	go p.animate(orBackground(p.Context))

	return &p, nil
}

func (p *SpinnerPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p *SpinnerPrinter) animate(ctx context.Context) {
	ticker := time.NewTicker(p.Delay)
	defer ticker.Stop()

	for {
		select {
		case <-p.run.done:
			return
		case <-ctx.Done():
			p.finish(statusCanceled, "")
			return
		case <-ticker.C:
			p.run.mu.Lock()
			p.run.frame++
			p.run.mu.Unlock()
			p.draw()
		}
	}
}

// draw shows the current frame.
func (p *SpinnerPrinter) draw() {
	if p.run.region == nil || !StylingEnabled() {
		return
	}

	c := canvasFor(p.out())
	frames := p.Frames

	if len(frames) == 0 {
		frames = c.g.Spinner
	}

	p.run.mu.Lock()
	frame := frames[p.run.frame%len(frames)]
	line := c.paint(p.Tone, frame) + " " + c.paint(ToneText, p.run.text) + p.timer(c)
	p.run.mu.Unlock()

	p.run.region.Set(node(line))
}

// timer returns the muted elapsed time, with a leading space. The caller
// holds p.run.mu.
func (p *SpinnerPrinter) timer(c canvas) string {
	if !p.ShowTimer {
		return ""
	}

	return " " + c.paint(ToneMuted, formatDuration(time.Since(p.run.started)))
}

// UpdateText replaces the text of a running spinner.
func (p *SpinnerPrinter) UpdateText(text string) {
	if p.run == nil {
		return
	}

	p.run.mu.Lock()
	p.run.text = text
	p.run.mu.Unlock()
	p.draw()
}

// Stop ends the spinner as a success with its current text.
func (p *SpinnerPrinter) Stop() error {
	p.finish(statusSuccess, "")
	return nil
}

// Success ends the spinner with a success glyph and the given text, or its
// current text when none is given.
func (p *SpinnerPrinter) Success(text ...any) { p.finish(statusSuccess, sprintln(text...)) }

// Fail ends the spinner with an error glyph.
func (p *SpinnerPrinter) Fail(text ...any) { p.finish(statusError, sprintln(text...)) }

// Warning ends the spinner with a warning glyph.
func (p *SpinnerPrinter) Warning(text ...any) { p.finish(statusWarning, sprintln(text...)) }

// Info ends the spinner with an info glyph.
func (p *SpinnerPrinter) Info(text ...any) { p.finish(statusInfo, sprintln(text...)) }

// Skip ends the spinner as skipped.
func (p *SpinnerPrinter) Skip(text ...any) { p.finish(statusSkipped, sprintln(text...)) }

func (p *SpinnerPrinter) finish(st status, text string) {
	if p.run == nil {
		return
	}

	p.run.finished.Do(func() {
		close(p.run.done)

		if p.run.region == nil {
			return
		}

		final := ""
		if !p.RemoveWhenDone {
			final = p.final(canvasFor(p.out()), st, text)
		}

		p.run.region.Close(final)
	})
}

// final renders the status line the spinner leaves behind.
func (p *SpinnerPrinter) final(c canvas, st status, text string) string {
	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	if text == "" {
		text = p.run.text
	}

	textTone, timer := ToneText, p.timer(c)
	if st == statusSkipped {
		textTone, timer = ToneMuted, ""
	}

	line := st.glyph(c) + " " + c.paint(textTone, text) + timer
	if st == statusCanceled {
		line += " " + c.paint(ToneMuted, "canceled")
	}

	return node(line)
}

// status is how a step, a spinner or a prompt ended.
type status uint8

const (
	statusSuccess status = iota
	statusError
	statusWarning
	statusInfo
	statusSkipped
	statusCanceled
	statusPending
)

// glyph returns the colored glyph of st.
func (st status) glyph(c canvas) string {
	switch st {
	case statusError:
		return c.paint(ToneError, c.g.Error)
	case statusWarning:
		return c.paint(ToneWarning, c.g.Warning)
	case statusInfo:
		return c.paint(ToneInfo, c.g.Info)
	case statusSkipped:
		return c.paint(ToneMuted, c.g.Skipped)
	case statusCanceled:
		return c.paint(ToneError, c.g.Canceled)
	case statusPending:
		return c.paint(ToneMuted, c.g.Pending)
	default:
		return c.paint(ToneSuccess, c.g.Success)
	}
}

// formatDuration formats d the way pterm shows elapsed time: "0.4s",
// "12.3s", "1m 34s", "2h 5m".
func formatDuration(d time.Duration) string {
	switch {
	case d < time.Minute:
		return fmt.Sprintf("%.1fs", d.Seconds())
	case d < time.Hour:
		return fmt.Sprintf("%dm %ds", int(d.Minutes()), int(d.Seconds())%60)
	default:
		return fmt.Sprintf("%dh %dm", int(d.Hours()), int(d.Minutes())%60)
	}
}

// flag returns the value of an optional bool argument, true when omitted.
func flag(b []bool) bool {
	return len(b) == 0 || b[0]
}
