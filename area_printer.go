package pterm

import (
	"strings"

	"github.com/atomicgo/cursor"

	"github.com/pterm/pterm/internal"
)

// DefaultArea is the default area printer.
var DefaultArea = AreaPrinter{}

// AreaPrinter prints an area which can be updated easily.
// use this printer for live output like charts, algorithm visualizations, simulations and even games.
type AreaPrinter struct {
	RemoveWhenDone bool
	FullScreen     bool

	content  string
	isActive bool

	area *cursor.Area
}

// GetContent returns the current area content.
func (s *AreaPrinter) GetContent() string {
	return s.content
}

// WithRemoveWhenDone removes the AreaPrinter content after it is stopped.
func (s AreaPrinter) WithRemoveWhenDone(b ...bool) *AreaPrinter {
	s.RemoveWhenDone = internal.WithBoolean(b)
	return &s
}

// WithFullscreen starts the AreaPrinter width the same height as the terminal, making it fullscreen.
func (s AreaPrinter) WithFullscreen(b ...bool) *AreaPrinter {
	s.FullScreen = internal.WithBoolean(b)
	return &s
}

// Update overwrites the content of the AreaPrinter.
// Can be used live.
func (s *AreaPrinter) Update(text ...interface{}) {
	if s.area == nil {
		newArea := cursor.NewArea()
		s.area = &newArea
	}
	str := Sprint(text...)
	s.content = str
	if s.FullScreen {
		str = strings.TrimRight(str, "\n")
		height := GetTerminalHeight()
		contentHeight := strings.Count(str, "\n")

		if height > contentHeight {
			str += strings.Repeat("\n", height-contentHeight-2)
		}
	}
	s.area.Update(str)
}

// Start the AreaPrinter.
func (s *AreaPrinter) Start(text ...interface{}) (*AreaPrinter, error) {
	s.isActive = true
	str := Sprint(text...)
	newArea := cursor.NewArea()
	s.area = &newArea

	s.Update(str)

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
