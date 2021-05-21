package pterm

import (
	"github.com/atomicgo/cursor"

	"github.com/pterm/pterm/internal"
)

var DefaultArea = AreaPrinter{}

type AreaPrinter struct {
	Content        string
	RemoveWhenDone bool

	isActive bool

	area *cursor.Area
}

// WithRemoveWhenDone removes the AreaPrinter content after it is stopped.
func (s AreaPrinter) WithRemoveWhenDone(b ...bool) *AreaPrinter {
	s.RemoveWhenDone = internal.WithBoolean(b)
	return &s
}

// Update updates the message of the active AreaPrinter.
// Can be used live.
func (s *AreaPrinter) Update(text ...interface{}) {
	str := Sprint(text...)
	s.Content = str
	s.area.Update(str)
}

// Start the AreaPrinter.
func (s *AreaPrinter) Start(text ...interface{}) (*AreaPrinter, error) {
	s.isActive = true
	str := Sprint(text...)
	newArea := cursor.NewArea()
	s.area = &newArea

	s.area.Update(str)

	return s, nil
}

// Stop terminates the AreaPrinter immediately.
// The AreaPrinter will not resolve into anything.
func (s *AreaPrinter) Stop() error {
	s.isActive = false
	if s.RemoveWhenDone {
		s.area.Clear()
	}
	return nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (s *AreaPrinter) GenericStart() (*LivePrinter, error) {
	_, _ = s.Start()
	lp := LivePrinter(s)
	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (s *AreaPrinter) GenericStop() (*LivePrinter, error) {
	_ = s.Stop()
	lp := LivePrinter(s)
	return &lp, nil
}
