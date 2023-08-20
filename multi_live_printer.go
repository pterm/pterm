package pterm

import (
	"bytes"
	"io"
	"os"
	"strings"
	"time"
)

var DefaultMultiPrinter = MultiPrinter{
	Printers: []LivePrinter{},
	Writer:   os.Stdout,

	buffers: []*bytes.Buffer{},
	area:    DefaultArea,
}

type MultiPrinter struct {
	Printers []LivePrinter
	Writer   io.Writer

	buffers []*bytes.Buffer
	area    AreaPrinter
}

// SetWriter sets the writer for the AreaPrinter.
func (p *MultiPrinter) SetWriter(writer io.Writer) {
	p.Writer = writer
}

func (p *MultiPrinter) NewWriter() io.Writer {
	buf := bytes.NewBufferString("")
	p.buffers = append(p.buffers, buf)
	return buf
}

// getString returns all buffers appended and separated by a newline.
func (p *MultiPrinter) getString() string {
	var buffer bytes.Buffer
	for _, b := range p.buffers {
		s := b.String()
		parts := strings.Split(s, "\r") // only get the last override
		s = parts[len(parts)-1]
		s = strings.Trim(s, "\n\r")
		buffer.WriteString(s)
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (p *MultiPrinter) Start() (*MultiPrinter, error) {
	for _, printer := range p.Printers {
		printer.GenericStart()
	}

	go func() {
		ticker := time.NewTicker(time.Millisecond * 100)

		for range ticker.C {
			p.area.Update(p.getString())
		}
	}()

	return p, nil
}

func (p *MultiPrinter) Stop() (*MultiPrinter, error) {
	for _, printer := range p.Printers {
		printer.GenericStop()
	}
	p.area.Stop()
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
