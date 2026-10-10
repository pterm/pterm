package pterm

import (
	"errors"
	"io"
	"sync"
)

// MultiPrinter groups live printers so they can be stopped together.
//
// Live printers that run at the same time already stack below each other
// on their own, so a MultiPrinter is not needed to show several of them.
// It is a convenience for code that starts a batch of bars or spinners and
// wants to end them in one call:
//
//	multi, _ := pterm.Multi.Start()
//	a, _ := pterm.Progressbar.WithWriter(multi.NewWriter()).Start("a.zip")
//	b, _ := pterm.Progressbar.WithWriter(multi.NewWriter()).Start("b.zip")
//	multi.Add(a, b)
//	...
//	multi.Stop() // stops a and b
type MultiPrinter struct {
	// Writer is the writer NewWriter hands out. Nil means the default
	// output.
	Writer io.Writer

	run *multiRun
}

type multiRun struct {
	mu       sync.Mutex
	printers []LivePrinter
	stopped  bool
}

// Multi is the default MultiPrinter.
var Multi = MultiPrinter{}

// WithWriter returns a copy of p whose printers write to w.
func (p MultiPrinter) WithWriter(w io.Writer) *MultiPrinter {
	p.Writer = w
	return &p
}

// Start returns a running group that holds printers.
func (p MultiPrinter) Start(printers ...LivePrinter) (*MultiPrinter, error) {
	p.run = &multiRun{}
	p.Add(printers...)

	return &p, nil
}

// NewWriter returns the writer the printers of the group should use. Every
// call returns the same writer; it exists so code written for older pterm
// versions keeps working.
func (p *MultiPrinter) NewWriter() io.Writer {
	return orDefault(p.Writer)
}

// Add puts printers into a running group. Printers added after Stop are
// stopped right away.
func (p *MultiPrinter) Add(printers ...LivePrinter) *MultiPrinter {
	if p.run == nil {
		return p
	}

	p.run.mu.Lock()
	stopped := p.run.stopped

	if !stopped {
		p.run.printers = append(p.run.printers, printers...)
	}
	p.run.mu.Unlock()

	if stopped {
		for _, lp := range printers {
			_ = lp.Stop()
		}
	}

	return p
}

// Stop stops every printer of the group that is still running, in the
// order they were added, and returns their errors joined.
func (p *MultiPrinter) Stop() error {
	if p.run == nil {
		return nil
	}

	p.run.mu.Lock()
	printers := p.run.printers
	p.run.printers = nil
	p.run.stopped = true
	p.run.mu.Unlock()

	errs := make([]error, 0, len(printers))
	for _, lp := range printers {
		errs = append(errs, lp.Stop())
	}

	return errors.Join(errs...)
}
