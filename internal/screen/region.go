package screen

import (
	"os"
	"os/signal"
	"slices"
	"strings"
	"sync"
	"syscall"
)

// Region is a block of lines that a live printer redraws in place. Create
// one with Screen.NewRegion, update it with Set, and end it with Close.
//
// All methods are safe for concurrent use. After Close, or after an
// interrupt closed the region, Set and Close do nothing.
type Region struct {
	s       *Screen
	content string
	closed  bool

	// onInterrupt returns the final content to print when the program gets
	// interrupted while the region is open.
	onInterrupt func() string
	// handler is the printer's own interrupt handler, or nil.
	handler func()
}

// RegionOptions configure a new region.
type RegionOptions struct {
	// OnInterrupt returns the content the region shows after an interrupt,
	// for example the step marked as canceled. Nil keeps the last frame.
	OnInterrupt func() string

	// Handler replaces the default reaction to an interrupt, which is to
	// exit with code 130. It runs after every region was closed.
	Handler func()
}

// NewRegion adds an empty region below the existing ones.
func (s *Screen) NewRegion(opts RegionOptions) *Region {
	r := &Region{s: s, onInterrupt: opts.OnInterrupt, handler: opts.Handler}

	s.mu.Lock()
	s.regions = append(s.regions, r)
	s.mu.Unlock()

	watchInterrupts()

	return r
}

// Set replaces the content of r and redraws the regions.
func (r *Region) Set(content string) {
	s := r.s

	s.mu.Lock()
	defer s.mu.Unlock()

	if r.closed || r.content == content {
		return
	}

	r.content = content

	s.redraw()
}

// Content returns what r currently shows.
func (r *Region) Content() string {
	r.s.mu.Lock()
	defer r.s.mu.Unlock()

	return r.content
}

// Closed reports whether r was closed, by Close or by an interrupt.
func (r *Region) Closed() bool {
	r.s.mu.Lock()
	defer r.s.mu.Unlock()

	return r.closed
}

// Close removes r and prints final as normal output in its place, ending
// it with a newline if it has none. An empty final removes the region
// without a trace. It reports false when r was
// already closed.
func (r *Region) Close(final string) bool {
	s := r.s

	s.mu.Lock()
	defer s.mu.Unlock()

	if r.closed {
		return false
	}

	s.closeRegion(r, final)

	return true
}

// closeRegion removes r and prints final. The caller holds s.mu.
func (s *Screen) closeRegion(r *Region, final string) {
	r.closed = true
	s.regions = slices.DeleteFunc(s.regions, func(x *Region) bool { return x == r })

	if final != "" && !strings.HasSuffix(final, "\n") {
		final += "\n"
	}

	s.writeItem(final)
}

// Interrupted reacts to Ctrl+C. It closes every open region on every screen
// with its interrupt content, shows the cursor again and then either runs
// the handlers of the interrupted printers or, when none of them has one,
// exits with code 130.
//
// Interactive printers call it themselves, because a terminal in raw mode
// delivers Ctrl+C as a key instead of a signal.
func Interrupted() {
	var handlers []func()

	registryMu.Lock()

	screens := make([]*Screen, 0, len(registry))
	for _, s := range registry {
		screens = append(screens, s)
	}

	registryMu.Unlock()

	for _, s := range screens {
		s.mu.Lock()

		for len(s.regions) > 0 {
			r := s.regions[0]

			final := r.content
			if r.onInterrupt != nil {
				final = r.onInterrupt()
			}

			if r.handler != nil {
				handlers = append(handlers, r.handler)
			}

			s.closeRegion(r, final)
		}

		s.mu.Unlock()
	}

	if len(handlers) == 0 {
		Exit(130)
		return
	}

	for _, h := range handlers {
		h()
	}
}

// Exit ends the program. Tests replace it.
var Exit = os.Exit

var (
	signalOnce sync.Once
	signals    = make(chan os.Signal, 1)
)

// watchInterrupts starts listening for SIGINT and SIGTERM once the first
// region opens. A program that never draws a live printer keeps the default
// signal behavior.
func watchInterrupts() {
	signalOnce.Do(func() {
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

		go func() {
			for range signals {
				if openRegions() {
					Interrupted()
					continue
				}

				// Nothing is drawn, so behave as if pterm never installed a
				// handler.
				Exit(130)
			}
		}()
	})
}

func openRegions() bool {
	registryMu.Lock()
	defer registryMu.Unlock()

	for _, s := range registry {
		s.mu.Lock()
		n := len(s.regions)
		s.mu.Unlock()

		if n > 0 {
			return true
		}
	}

	return false
}
