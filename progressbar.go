package pterm

import (
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"

	"github.com/pterm/pterm/internal"
)

var (
	DefaultProgressbar = Progressbar{
		Total:         100,
		LineCharacter: '█',
		LastCharacter: '█',
	}
)

var fade = []string{"FF3D3D", "FC3F3C", "F9423C", "F7453C", "F4483C", "F24B3C", "EF4D3C", "EC503C", "EA533C", "E7563C",
	"E5593C", "E25C3C", "DF5E3C", "DD613C", "DA643C", "D8673C", "D56A3B", "D36D3B", "D06F3B", "CD723B", "CB753B",
	"C8783B", "C67B3B", "C37D3B", "C0803B", "BE833B", "BB863B", "B9893B", "B68C3B", "B38E3B", "B1913B", "AE943B",
	"AC973A", "A99A3A", "A79D3A", "A49F3A", "A1A23A", "9FA53A", "9CA83A", "9AAB3A", "97AE3A", "94B03A", "92B33A",
	"8FB63A", "8DB93A", "8ABC3A", "87BE3A", "85C13A", "82C439", "80C739", "7DCA39", "7BCD39", "78CF39", "75D239",
	"73D539", "70D839", "6EDB39", "6BDE39", "68E039", "66E339", "63E639", "61E939", "5EEC39", "5CEF39"}

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
			width := GetTerminalWidth()
			decoratorCurrentTotal := Sprintf("[%s%s%s]", Green(p.Current), Gray("/"), Red(p.Total))
			currentPercent := int(internal.PercentageRound(float64(int64(p.Total)), float64(int64(p.Current)), float64(width)))

			before := Cyan(p.Name) + " " + decoratorCurrentTotal + " "
			after := " " + color.HEX(fade[int(0.63*float64(currentPercent))]).Sprint(strconv.Itoa(currentPercent)+"%")

			barMaxLength := width - len(RemoveColors(before)) - len(RemoveColors(after)) - 1
			barCurrentLength := (p.Current * barMaxLength) / p.Total
			barFiller := strings.Repeat(" ", barMaxLength-barCurrentLength)

			bar := LightCyan(strings.Repeat(string(p.LineCharacter), barCurrentLength)+string(p.LastCharacter)) + barFiller
			Printo(before + bar + after)

			if p.Current == p.Total {
				Println()
				p.Stop()
			}

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
