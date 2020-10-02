package pterm

import (
	"strings"
	"time"
)

var (
	DefaultProgressbar = Progressbar{
		Total:         100,
		LineCharacter: '=',
		LastCharacter: '=',
	}
)

// Progressbar shows a progress animation in the terminal.
type Progressbar struct {
	Name          string
	Total         int
	Current       int
	UpdateDelay   time.Duration
	LineCharacter rune
	LastCharacter rune

	IsActive bool
}

// SetName sets the name of the progressbar.
func (p Progressbar) SetName(name string) *Progressbar {
	p.Name = name
	return &p
}

// SetTotal sets the total value of the progressbar.
func (p Progressbar) SetTotal(total int) *Progressbar {
	p.Total = total
	return &p
}

// Increment current value by one.
func (p *Progressbar) Increment() *Progressbar {
	p.Add(1)
	return p
}

// Add to current value.
func (p *Progressbar) Add(count int) *Progressbar {
	p.Current += count
	return p
}

// Start the progressbar.
func (p Progressbar) Start() *Progressbar {
	p.IsActive = true

	if p.UpdateDelay == 0 {
		p.UpdateDelay = time.Millisecond * 100
	}

	if p.LineCharacter == 0 {
		p.LineCharacter = '='
	}

	if p.LastCharacter == 0 {
		p.LastCharacter = '='
	}

	go func() {
		for p.IsActive {
			if p.Current > p.Total {
				p.Total = p.Current
			}

			decoratorCurrentTotal := Sprintf("[%d/%d]", p.Current, p.Total)

			before := p.Name + " " + decoratorCurrentTotal + " "
			after := " After bar"

			barMaxLength := GetTerminalWidth() - len(before) - len(after)
			barCurrentLength := (p.Current * barMaxLength) / p.Total
			barFiller := strings.Repeat(" ", barMaxLength-barCurrentLength)

			bar := strings.Repeat(string(p.LineCharacter), barCurrentLength-1) + string(p.LastCharacter) + barFiller
			Printo(before + bar + after)

			time.Sleep(p.UpdateDelay)
		}
	}()

	return &p
}

// Stop the progressbar.
func (p *Progressbar) Stop() *Progressbar {
	p.IsActive = false
	return p
}
