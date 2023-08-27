package pterm

import (
	"io"
	"strings"

	"atomicgo.dev/cursor"

	"github.com/pterm/pterm/internal"
)

// DefaultArea is the default area printer.
var DefaultArea = AreaPrinter{}

// AreaPrinter prints an area which can be updated easily.
// use this printer for live output like charts, algorithm visualizations, simulations and even games.
type AreaPrinter struct {
	RemoveWhenDone bool
	Fullscreen     bool
	Center         bool

	content  string
	isActive bool

	area *cursor.Area
}

// GetContent returns the current area content.
func (p *AreaPrinter) GetContent() string {
	return p.content
}

// WithRemoveWhenDone removes the AreaPrinter content after it is stopped.
func (p AreaPrinter) WithRemoveWhenDone(b ...bool) *AreaPrinter {
	p.RemoveWhenDone = internal.WithBoolean(b)
	return &p
}

// WithFullscreen sets the AreaPrinter height the same height as the terminal, making it fullscreen.
func (p AreaPrinter) WithFullscreen(b ...bool) *AreaPrinter {
	p.Fullscreen = internal.WithBoolean(b)
	return &p
}

// WithCenter centers the AreaPrinter content to the terminal.
func (p AreaPrinter) WithCenter(b ...bool) *AreaPrinter {
	p.Center = internal.WithBoolean(b)
	return &p
}

// SetWriter sets the writer for the AreaPrinter.
func (p *AreaPrinter) SetWriter(writer io.Writer) {

}

// Update overwrites the content of the AreaPrinter.
// Can be used live.
func (p *AreaPrinter) Update(text ...interface{}) {
	if p.area == nil {
		newArea := cursor.NewArea()
		p.area = &newArea
	}
	str := Sprint(text...)
	p.content = str

	if p.Center {
		str = DefaultCenter.Sprint(str)
	}

	if p.Fullscreen {
		str = strings.TrimRight(str, "\n")
		height := GetTerminalHeight()
		contentHeight := strings.Count(str, "\n")

		topPadding := 0
		bottomPadding := height - contentHeight - 2

		if p.Center {
			topPadding = (bottomPadding / 2) + 1
			bottomPadding /= 2
		}

		if height > contentHeight {
			str = strings.Repeat("\n", topPadding) + str
			str += strings.Repeat("\n", bottomPadding)
		}
	}
	p.area.Update(str)
}

// Start the AreaPrinter.
func (p *AreaPrinter) Start(text ...interface{}) (*AreaPrinter, error) {
	p.isActive = true
	str := Sprint(text...)
	newArea := cursor.NewArea()
	p.area = &newArea

	p.Update(str)

	return p, nil
}

// Stop terminates the AreaPrinter immediately.
// The AreaPrinter will not resolve into anything.
func (p *AreaPrinter) Stop() error {
	if !p.isActive {
		return nil
	}
	p.isActive = false
	if p.RemoveWhenDone {
		p.Clear()
	}
	return nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (p *AreaPrinter) GenericStart() (*LivePrinter, error) {
	_, _ = p.Start()
	lp := LivePrinter(p)
	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (p *AreaPrinter) GenericStop() (*LivePrinter, error) {
	_ = p.Stop()
	lp := LivePrinter(p)
	return &lp, nil
}

// Clear is a Wrapper function that clears the content of the Area
// moves the cursor to the bottom of the terminal, clears n lines upwards from
// the current position and moves the cursor again.
func (p *AreaPrinter) Clear() {
	p.area.Clear()
}
