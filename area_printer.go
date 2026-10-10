package pterm

import (
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/pterm/pterm/internal/screen"
	"github.com/pterm/pterm/internal/text"
)

// AreaPrinter is a block of free-form text that you redraw in place, for
// live output no other printer covers, such as a chart that updates every
// second:
//
//	area, _ := pterm.Area.Start("Waiting for data")
//	area.Update(chart)
//	area.Stop()
//
// On writers that are not terminals only the last content is printed, when
// the area stops.
type AreaPrinter struct {
	// RemoveWhenDone removes the content when the area stops instead of
	// leaving it on screen.
	RemoveWhenDone bool
	// Fullscreen pads the content to the height of the terminal.
	Fullscreen bool
	// Center centers every line horizontally, and the content vertically
	// when Fullscreen is set.
	Center bool
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Context stops the area when it ends.
	Context context.Context
	// OnInterrupt runs when the user presses Ctrl+C while the area is shown,
	// after the terminal was restored. Nil exits the program with code 130.
	OnInterrupt func()

	run *areaRun
}

// areaRun is the state of a started area. Copies of a started AreaPrinter
// share it.
type areaRun struct {
	mu       sync.Mutex
	content  string
	region   *screen.Region
	done     chan struct{}
	stopped  bool
	finished sync.Once
}

// Area is the default area.
var Area = AreaPrinter{}

// WithRemoveWhenDone returns a copy of p that clears its content when it
// stops.
func (p AreaPrinter) WithRemoveWhenDone(remove ...bool) *AreaPrinter {
	p.RemoveWhenDone = flag(remove)
	return &p
}

// WithFullscreen returns a copy of p that fills the terminal height.
func (p AreaPrinter) WithFullscreen(fullscreen ...bool) *AreaPrinter {
	p.Fullscreen = flag(fullscreen)
	return &p
}

// WithCenter returns a copy of p that centers its content.
func (p AreaPrinter) WithCenter(center ...bool) *AreaPrinter {
	p.Center = flag(center)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p AreaPrinter) WithWriter(w io.Writer) *AreaPrinter {
	p.Writer = w
	return &p
}

// WithContext returns a copy of p that stops when ctx ends.
func (p AreaPrinter) WithContext(ctx context.Context) *AreaPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p that runs fn on Ctrl+C instead of
// exiting.
func (p AreaPrinter) WithOnInterrupt(fn func()) *AreaPrinter {
	p.OnInterrupt = fn
	return &p
}

// Start shows the area with the given content and returns the running
// area.
func (p AreaPrinter) Start(content ...any) (*AreaPrinter, error) {
	run := &areaRun{content: fmt.Sprint(content...), done: make(chan struct{})}
	p.run = run

	if OutputEnabled() {
		run.region = screen.For(p.out()).NewRegion(screen.RegionOptions{
			OnInterrupt: p.final,
			Handler:     p.OnInterrupt,
		})
	}

	p.draw()

	if ctx := p.Context; ctx != nil {
		go func() {
			select {
			case <-run.done:
			case <-ctx.Done():
				_ = p.Stop()
			}
		}()
	}

	return &p, nil
}

func (p *AreaPrinter) out() io.Writer { return orDefault(p.Writer) }

// Update replaces the content of a running area.
func (p *AreaPrinter) Update(content ...any) {
	if p.run == nil {
		return
	}

	p.run.mu.Lock()
	if p.run.stopped {
		p.run.mu.Unlock()
		return
	}

	p.run.content = fmt.Sprint(content...)
	p.run.mu.Unlock()
	p.draw()
}

// Clear empties a running area. It stays open for the next Update.
func (p *AreaPrinter) Clear() {
	p.Update()
}

// Content returns the content last given to Start or Update.
func (p *AreaPrinter) Content() string {
	if p.run == nil {
		return ""
	}

	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	return p.run.content
}

// Stop ends the area. Its last content stays on screen unless
// RemoveWhenDone is set.
func (p *AreaPrinter) Stop() error {
	if p.run == nil {
		return nil
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
			final = p.final()
		}

		p.run.region.Close(final)
	})

	return nil
}

// draw shows the current content.
func (p *AreaPrinter) draw() {
	if p.run.region == nil || !StylingEnabled() {
		return
	}

	p.run.mu.Lock()
	content := p.layout(p.run.content, true)
	p.run.mu.Unlock()

	p.run.region.Set(content)
}

// final returns what the area leaves behind. A terminal keeps the last
// frame as it was; other writers get the content without the padding a
// fullscreen area adds.
func (p *AreaPrinter) final() string {
	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	return p.layout(p.run.content, StylingEnabled() && screen.IsTerminal(p.out()))
}

// layout wraps and centers content for the terminal. The caller holds
// p.run.mu.
func (p *AreaPrinter) layout(content string, live bool) string {
	content = strings.TrimRight(content, "\n")
	width, height := screen.Size(p.out())

	// Rows are cut one column before the edge, so wrap there too.
	lines := strings.Split(content, "\n")
	if width > 1 {
		lines = text.Wrap(content, width-1)
	}

	if p.Center && width > 1 {
		widest := 0
		for _, l := range lines {
			widest = max(widest, text.Width(l))
		}

		// Center the block as a whole, so its lines stay aligned with each
		// other.
		pad := strings.Repeat(" ", max(0, (width-1-widest)/2))

		for i, l := range lines {
			if l != "" {
				lines[i] = pad + l
			}
		}
	}

	if live && p.Fullscreen && height > 1 {
		// The screen draws at most height-1 rows of live output.
		free := height - 1 - len(lines)
		if free > 0 {
			top := 0
			if p.Center {
				top = free / 2
			}

			lines = append(make([]string, top), lines...)
			lines = append(lines, make([]string, free-top)...)
		}
	}

	return strings.Join(lines, "\n")
}
