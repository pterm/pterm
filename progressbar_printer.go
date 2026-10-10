package pterm

import (
	"context"
	"io"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/text"
)

// ProgressbarPrinter shows how far a task with a known amount of work got.
// It ends as a status line with the count and the time it took:
//
//	Downloading  ██████████▍░░░░░░░░░   52%   52/100  1.2s · eta 1.1s
//	✔ Downloading 100/100 2.3s
//
// The bar stops by itself when Current reaches Total.
type ProgressbarPrinter struct {
	// Title describes the task.
	Title string
	// Total is the amount of work. The bar is full when Current reaches it.
	Total int
	// Current is the amount of work done when the bar starts.
	Current int

	// Tone colors the filled part of the bar.
	Tone Tone
	// TrackTone colors the empty part of the bar.
	TrackTone Tone
	// TitleTone colors the title.
	TitleTone Tone
	// PercentageTone colors the percentage.
	PercentageTone Tone
	// DetailTone colors the count, the elapsed time and the ETA.
	DetailTone Tone
	// Fill is the glyph of the filled part. Empty uses the meter fill of
	// the glyph set.
	Fill string
	// Track is the glyph of the empty part. Empty uses the meter track of
	// the glyph set.
	Track string

	// BarWidth is the width of the bar itself in columns. The bar gets
	// narrower when the line would not fit otherwise.
	BarWidth int
	// MaxWidth caps the width of the whole line. Zero or less uses the
	// terminal width.
	MaxWidth int

	// ShowTitle shows the title.
	ShowTitle bool
	// ShowPercentage shows the percentage after the bar.
	ShowPercentage bool
	// ShowCount shows "current/total" after the bar.
	ShowCount bool
	// ShowElapsedTime shows the time since the start.
	ShowElapsedTime bool
	// ShowETA shows an estimate of the time left.
	ShowETA bool
	// RemoveWhenDone removes the bar when it stops instead of leaving a
	// status line.
	RemoveWhenDone bool

	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Context stops the bar as canceled when it ends.
	Context context.Context
	// OnInterrupt runs when the user presses Ctrl+C while the bar runs,
	// after the terminal was restored. Nil exits the program with code 130.
	OnInterrupt func()

	run *progressRun
}

// progressRun is the state of a started progress bar. Copies of a started
// ProgressbarPrinter share it.
type progressRun struct {
	mu         sync.Mutex
	title      string
	current    int
	total      int
	started    time.Time
	titleWidth int // widest title so far, so the bar does not jump around
	region     *screen.Region
	done       chan struct{}
	stopped    bool
	finished   sync.Once // the final line is printed once, and callers wait for it
}

// Progressbar is the default progress bar.
var Progressbar = ProgressbarPrinter{
	Total:           100,
	Tone:            TonePrimary,
	TrackTone:       ToneBorder,
	TitleTone:       ToneText,
	PercentageTone:  ToneText,
	DetailTone:      ToneMuted,
	BarWidth:        30,
	MaxWidth:        100,
	ShowTitle:       true,
	ShowPercentage:  true,
	ShowCount:       true,
	ShowElapsedTime: true,
	ShowETA:         true,
}

// progressbarTick is how often a running bar redraws to update its timer.
const progressbarTick = 200 * time.Millisecond

// WithTitle returns a copy of p with title t.
func (p ProgressbarPrinter) WithTitle(t string) *ProgressbarPrinter {
	p.Title = t
	return &p
}

// WithTotal returns a copy of p that is full at total.
func (p ProgressbarPrinter) WithTotal(total int) *ProgressbarPrinter {
	p.Total = total
	return &p
}

// WithCurrent returns a copy of p that starts at current.
func (p ProgressbarPrinter) WithCurrent(current int) *ProgressbarPrinter {
	p.Current = current
	return &p
}

// WithTone returns a copy of p whose fill uses tone t.
func (p ProgressbarPrinter) WithTone(t Tone) *ProgressbarPrinter {
	p.Tone = t
	return &p
}

// WithTrackTone returns a copy of p whose track uses tone t.
func (p ProgressbarPrinter) WithTrackTone(t Tone) *ProgressbarPrinter {
	p.TrackTone = t
	return &p
}

// WithTitleTone returns a copy of p whose title uses tone t.
func (p ProgressbarPrinter) WithTitleTone(t Tone) *ProgressbarPrinter {
	p.TitleTone = t
	return &p
}

// WithPercentageTone returns a copy of p whose percentage uses tone t.
func (p ProgressbarPrinter) WithPercentageTone(t Tone) *ProgressbarPrinter {
	p.PercentageTone = t
	return &p
}

// WithDetailTone returns a copy of p whose count and times use tone t.
func (p ProgressbarPrinter) WithDetailTone(t Tone) *ProgressbarPrinter {
	p.DetailTone = t
	return &p
}

// WithFill returns a copy of p that draws the filled part with glyph g.
// A custom fill has no partial blocks, so the bar moves a column at a time.
func (p ProgressbarPrinter) WithFill(g string) *ProgressbarPrinter {
	p.Fill = g
	return &p
}

// WithTrack returns a copy of p that draws the empty part with glyph g.
func (p ProgressbarPrinter) WithTrack(g string) *ProgressbarPrinter {
	p.Track = g
	return &p
}

// WithBarWidth returns a copy of p whose bar is w columns wide.
func (p ProgressbarPrinter) WithBarWidth(w int) *ProgressbarPrinter {
	p.BarWidth = w
	return &p
}

// WithMaxWidth returns a copy of p whose line is at most w columns wide.
func (p ProgressbarPrinter) WithMaxWidth(w int) *ProgressbarPrinter {
	p.MaxWidth = w
	return &p
}

// WithShowTitle returns a copy of p that shows or hides the title.
func (p ProgressbarPrinter) WithShowTitle(show ...bool) *ProgressbarPrinter {
	p.ShowTitle = flag(show)
	return &p
}

// WithShowPercentage returns a copy of p that shows or hides the percentage.
func (p ProgressbarPrinter) WithShowPercentage(show ...bool) *ProgressbarPrinter {
	p.ShowPercentage = flag(show)
	return &p
}

// WithShowCount returns a copy of p that shows or hides "current/total".
func (p ProgressbarPrinter) WithShowCount(show ...bool) *ProgressbarPrinter {
	p.ShowCount = flag(show)
	return &p
}

// WithShowElapsedTime returns a copy of p that shows or hides the elapsed
// time.
func (p ProgressbarPrinter) WithShowElapsedTime(show ...bool) *ProgressbarPrinter {
	p.ShowElapsedTime = flag(show)
	return &p
}

// WithShowETA returns a copy of p that shows or hides the time left.
func (p ProgressbarPrinter) WithShowETA(show ...bool) *ProgressbarPrinter {
	p.ShowETA = flag(show)
	return &p
}

// WithRemoveWhenDone returns a copy of p that disappears when it stops.
func (p ProgressbarPrinter) WithRemoveWhenDone(remove ...bool) *ProgressbarPrinter {
	p.RemoveWhenDone = flag(remove)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p ProgressbarPrinter) WithWriter(w io.Writer) *ProgressbarPrinter {
	p.Writer = w
	return &p
}

// WithContext returns a copy of p that stops as canceled when ctx ends.
func (p ProgressbarPrinter) WithContext(ctx context.Context) *ProgressbarPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting.
func (p ProgressbarPrinter) WithOnInterrupt(fn func()) *ProgressbarPrinter {
	p.OnInterrupt = fn
	return &p
}

// Start shows the bar with the given title, or with p.Title when none is
// given, and returns the running bar.
func (p ProgressbarPrinter) Start(title ...any) (*ProgressbarPrinter, error) {
	if len(title) > 0 {
		p.Title = sprintln(title...)
	}

	run := &progressRun{
		title:      p.Title,
		current:    max(p.Current, 0),
		total:      p.Total,
		started:    time.Now(),
		titleWidth: text.Width(p.Title),
		done:       make(chan struct{}),
	}
	p.run = run

	if OutputEnabled() {
		run.region = screen.For(p.out()).NewRegion(screen.RegionOptions{
			OnInterrupt: func() string { return p.final(canvasFor(p.out()), statusCanceled, "") },
			Handler:     p.OnInterrupt,
		})
	}

	if run.total > 0 && run.current >= run.total {
		p.finish(statusSuccess, "")
		return &p, nil
	}

	p.draw()

	go p.animate(orBackground(p.Context))

	return &p, nil
}

func (p *ProgressbarPrinter) out() io.Writer { return orDefault(p.Writer) }

func (p *ProgressbarPrinter) animate(ctx context.Context) {
	ticker := time.NewTicker(progressbarTick)
	defer ticker.Stop()

	for {
		select {
		case <-p.run.done:
			return
		case <-ctx.Done():
			p.finish(statusCanceled, "")
			return
		case <-ticker.C:
			if p.ShowElapsedTime || p.ShowETA {
				p.draw()
			}
		}
	}
}

// Increment adds one to the progress.
func (p *ProgressbarPrinter) Increment() *ProgressbarPrinter {
	return p.Add(1)
}

// Add adds n to the progress. The bar stops as a success when it reaches
// its total.
func (p *ProgressbarPrinter) Add(n int) *ProgressbarPrinter {
	if p.run == nil {
		return p
	}

	p.run.mu.Lock()
	if p.run.stopped {
		p.run.mu.Unlock()
		return p
	}

	p.run.current = max(p.run.current+n, 0)
	full := p.run.total > 0 && p.run.current >= p.run.total

	if full {
		p.run.current = p.run.total
	}
	p.run.mu.Unlock()

	if full {
		p.finish(statusSuccess, "")
		return p
	}

	p.draw()

	return p
}

// UpdateTitle replaces the title of a running bar.
func (p *ProgressbarPrinter) UpdateTitle(title string) *ProgressbarPrinter {
	if p.run == nil {
		return p
	}

	p.run.mu.Lock()
	p.run.title = title
	p.run.titleWidth = max(p.run.titleWidth, text.Width(title))
	p.run.mu.Unlock()
	p.draw()

	return p
}

// Progress returns the current progress and the total of a running bar.
func (p *ProgressbarPrinter) Progress() (current, total int) {
	if p.run == nil {
		return p.Current, p.Total
	}

	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	return p.run.current, p.run.total
}

// Elapsed returns the time since the bar started.
func (p *ProgressbarPrinter) Elapsed() time.Duration {
	if p.run == nil {
		return 0
	}

	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	return time.Since(p.run.started)
}

// ResetTimer restarts the elapsed time of a running bar.
func (p *ProgressbarPrinter) ResetTimer() {
	if p.run == nil {
		return
	}

	p.run.mu.Lock()
	p.run.started = time.Now()
	p.run.mu.Unlock()
	p.draw()
}

// Stop ends the bar as a success, whatever its progress. When the context
// of the bar already ended, the bar ends as canceled instead, so Stop can
// wait for the canceled line before the program exits.
func (p *ProgressbarPrinter) Stop() error {
	st := statusSuccess
	if p.Context != nil && p.Context.Err() != nil {
		st = statusCanceled
	}

	p.finish(st, "")

	return nil
}

// Fail ends the bar with an error glyph and the given text, or its title
// when none is given.
func (p *ProgressbarPrinter) Fail(text ...any) { p.finish(statusError, sprintln(text...)) }

// draw shows the current state of the bar.
func (p *ProgressbarPrinter) draw() {
	if p.run.region == nil || !StylingEnabled() {
		return
	}

	line := p.frame(canvasFor(p.out()))
	p.run.region.Set(line)
}

// frame renders one frame of the running bar.
func (p *ProgressbarPrinter) frame(c canvas) string {
	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	current, total := p.run.current, p.run.total

	var head, tail []string

	if p.ShowTitle && p.run.titleWidth > 0 {
		head = append(head, c.paint(p.TitleTone, text.PadRight(p.run.title, p.run.titleWidth)))
	}

	if p.ShowPercentage {
		tail = append(tail, c.paint(p.PercentageTone, text.PadLeft(strconv.Itoa(percentOf(current, total))+"%", 4)))
	}

	if p.ShowCount {
		tail = append(tail, c.paint(p.DetailTone, p.count(current, total, true)))
	}

	if times := p.times(c, current, total); times != "" {
		tail = append(tail, times)
	}

	const gap = "  "

	fixed := 0
	for _, part := range head {
		fixed += text.Width(part) + len(gap)
	}

	for _, part := range tail {
		fixed += text.Width(part) + len(gap)
	}

	barWidth := p.BarWidth
	if limit := p.lineLimit(c); limit > 0 {
		barWidth = min(barWidth, limit-fixed)
	}

	parts := head

	if barWidth >= 3 {
		ratio := 0.0
		if total > 0 {
			ratio = float64(current) / float64(total)
		}

		parts = append(parts, drawMeter(p.barCanvas(c), ratio, barWidth, p.Tone, p.TrackTone))
	}

	return strings.Join(append(parts, tail...), gap)
}

// lineLimit returns how wide the line may get, or 0 for no limit.
func (p *ProgressbarPrinter) lineLimit(c canvas) int {
	limit := p.MaxWidth
	if c.width > 0 && (limit <= 0 || c.width-1 < limit) {
		limit = c.width - 1
	}

	return max(limit, 0)
}

// count formats "current/total". Padded, it keeps the same width while the
// numbers grow. The caller holds p.run.mu.
func (p *ProgressbarPrinter) count(current, total int, padded bool) string {
	if total <= 0 {
		return strconv.Itoa(current)
	}

	cur := strconv.Itoa(current)
	if padded {
		cur = text.PadLeft(cur, len(strconv.Itoa(total)))
	}

	return cur + "/" + strconv.Itoa(total)
}

// times returns the muted elapsed time and ETA. The caller holds p.run.mu.
func (p *ProgressbarPrinter) times(c canvas, current, total int) string {
	elapsed := time.Since(p.run.started)

	var parts []string

	if p.ShowElapsedTime {
		parts = append(parts, formatDuration(elapsed))
	}

	if p.ShowETA && current > 0 && total > current {
		eta := time.Duration(float64(elapsed) / float64(current) * float64(total-current))
		parts = append(parts, "eta "+formatDuration(eta))
	}

	if len(parts) == 0 {
		return ""
	}

	return c.paint(p.DetailTone, strings.Join(parts, " "+c.g.Dot+" "))
}

// final renders the status line the bar leaves behind.
func (p *ProgressbarPrinter) final(c canvas, st status, title string) string {
	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	if title == "" && p.ShowTitle {
		title = p.run.title
	}

	parts := []string{st.glyph(c)}
	if title != "" {
		parts = append(parts, c.paint(p.TitleTone, title))
	}

	switch {
	case p.ShowCount:
		parts = append(parts, c.paint(p.DetailTone, p.count(p.run.current, p.run.total, false)))
	case p.ShowPercentage || len(parts) == 1:
		parts = append(parts, c.paint(p.DetailTone, strconv.Itoa(percentOf(p.run.current, p.run.total))+"%"))
	}

	if p.ShowElapsedTime {
		parts = append(parts, c.paint(p.DetailTone, formatDuration(time.Since(p.run.started))))
	}

	if st == statusCanceled {
		parts = append(parts, c.paint(ToneMuted, "canceled"))
	}

	return node(strings.Join(parts, " "))
}

// finish ends the bar. When several goroutines end it at once, the first
// one wins and the others return after its final line was printed, so a
// Stop right before the program exits never loses that line.
func (p *ProgressbarPrinter) finish(st status, title string) {
	if p.run == nil {
		return
	}

	p.run.finished.Do(func() {
		p.run.mu.Lock()
		p.run.stopped = true
		close(p.run.done)
		p.run.mu.Unlock()

		if p.run.region == nil {
			return
		}

		final := ""
		if !p.RemoveWhenDone {
			final = p.final(canvasFor(p.out()), st, title)
		}

		p.run.region.Close(final)
	})
}

// percentOf returns current as a whole percentage of total, from 0 to
// 100. It rounds down, so a bar never shows 100% before it is done.
func percentOf(current, total int) int {
	if total <= 0 {
		return 0
	}

	return min(max(current*100/total, 0), 100)
}

// barCanvas returns c with the bar glyphs of p. A custom fill has no
// eighth blocks to go with it, so the bar then fills whole columns.
func (p *ProgressbarPrinter) barCanvas(c canvas) canvas {
	if p.Fill != "" {
		c.g.MeterFill, c.g.Eighths = p.Fill, nil
	}

	if p.Track != "" {
		c.g.MeterTrack = p.Track
	}

	return c
}
