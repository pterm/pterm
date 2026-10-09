package pterm

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"atomicgo.dev/schedule"
)

// DefaultMultiPrinter contains the default configuration for a MultiPrinter.
var DefaultMultiPrinter = MultiPrinter{
	printers:    []LivePrinter{},
	Writer:      os.Stdout,
	UpdateDelay: time.Millisecond * 200,

	buffers: []*multiPrinterBuffer{},
	area:    DefaultArea,
}

// multiPrinterBuffer is the thread-safe buffer handed out by NewWriter.
// Live printers write to it from their own goroutines while the MultiPrinter's
// update task reads it, so all access must be synchronized. bytes.Buffer alone
// is not safe for concurrent use.
type multiPrinterBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

// Write appends p to the buffer under the buffer lock.
func (b *multiPrinterBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.buf.Write(p)
}

// String returns the buffered content under the buffer lock.
func (b *multiPrinterBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.buf.String()
}

// MultiPrinter is able to print the output of multiple live printers
// (e.g. spinners and progressbars) at the same time.
type MultiPrinter struct {
	IsActive    bool
	Writer      io.Writer
	UpdateDelay time.Duration

	printers []LivePrinter
	buffers  []*multiPrinterBuffer
	area     AreaPrinter

	// updateTask repaints the area on every UpdateDelay tick while the
	// MultiPrinter is active. Stop stops it explicitly so the goroutine cannot
	// leak or repaint after Stop.
	updateTask *schedule.Task

	// mu serializes access to the printer's mutable state once Start has run.
	// It is a pointer so the value-receiver With* methods can copy the struct
	// without tripping go vet's copylocks check; it is lazily allocated.
	mu *sync.Mutex
}

// lock locks the printer's runtime mutex, allocating it on first use so the
// value-receiver builders can be used safely before Start.
func (p *MultiPrinter) lock() {
	if p.mu == nil {
		p.mu = &sync.Mutex{}
	}

	p.mu.Lock()
}

// unlock releases the printer's runtime mutex.
func (p *MultiPrinter) unlock() {
	p.mu.Unlock()
}

// SetWriter sets the writer for the AreaPrinter.
func (p *MultiPrinter) SetWriter(writer io.Writer) {
	p.Writer = writer
}

// WithWriter returns a fork of the MultiPrinter with a new writer.
func (p MultiPrinter) WithWriter(writer io.Writer) *MultiPrinter {
	p.Writer = writer
	return &p
}

// WithUpdateDelay returns a fork of the MultiPrinter with a new update delay.
func (p MultiPrinter) WithUpdateDelay(delay time.Duration) *MultiPrinter {
	p.UpdateDelay = delay
	return &p
}

// NewWriter returns a new writer that can be passed to a live printer
// (e.g. via WithWriter) so its output is rendered inside the MultiPrinter.
func (p *MultiPrinter) NewWriter() io.Writer {
	p.lock()
	defer p.unlock()

	buf := &multiPrinterBuffer{}
	p.buffers = append(p.buffers, buf)

	return buf
}

// getStringLocked returns all buffers appended and separated by a newline.
// The caller must hold p.mu.
func (p *MultiPrinter) getStringLocked() string {
	var buffer bytes.Buffer

	for _, b := range p.buffers {
		s := b.String()
		s = strings.Trim(s, "\n")

		parts := strings.Split(s, "\r") // only get the last override

		// check if the last part is empty, if so remove it, repeat until not
		// empty. If there is no part left, don't do anything
		for len(parts) > 0 && parts[len(parts)-1] == "" {
			parts = parts[:len(parts)-1]
		}

		if len(parts) > 0 {
			s = strings.Trim(parts[len(parts)-1], "\n\r")
			buffer.WriteString(s)
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

// Start starts the MultiPrinter and all its registered live printers.
func (p *MultiPrinter) Start() (*MultiPrinter, error) {
	p.lock()
	p.IsActive = true
	// Render the area to the configured writer if the cursor can be moved on
	// it (i.e. it exposes a file descriptor); otherwise it stays on stdout.
	p.area.SetWriter(p.Writer)

	printers := make([]LivePrinter, len(p.printers))
	copy(printers, p.printers)

	delay := p.UpdateDelay
	p.unlock()

	for _, printer := range printers {
		_, _ = printer.GenericStart()
	}

	task := schedule.Every(delay, func() bool {
		p.lock()
		defer p.unlock()

		if !p.IsActive {
			return false
		}

		p.area.Update(p.getStringLocked())

		return true
	})

	p.lock()
	p.updateTask = task
	p.unlock()

	return p, nil
}

// Stop stops the MultiPrinter and all its registered live printers.
func (p *MultiPrinter) Stop() (*MultiPrinter, error) {
	p.lock()
	p.IsActive = false
	task := p.updateTask
	p.updateTask = nil

	printers := make([]LivePrinter, len(p.printers))
	copy(printers, p.printers)
	p.unlock()

	// Stop the repaint task before the final repaint below. An in-flight tick
	// serializes on p.mu and skips its repaint once IsActive is false.
	if task != nil && task.IsActive() {
		task.Stop()
	}

	for _, printer := range printers {
		_, _ = printer.GenericStop()
	}

	// Repaint once more so the printers' final output is visible.
	p.lock()
	p.area.Update(p.getStringLocked())
	_ = p.area.Stop()
	p.unlock()

	return p, nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (p MultiPrinter) GenericStart() (*LivePrinter, error) {
	p2, _ := p.Start()
	lp := LivePrinter(p2)

	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (p MultiPrinter) GenericStop() (*LivePrinter, error) {
	p2, _ := p.Stop()
	lp := LivePrinter(p2)

	return &lp, nil
}
