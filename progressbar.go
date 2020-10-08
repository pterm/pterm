package pterm

import (
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"

	"github.com/pterm/pterm/internal"
)

// ActiveProgressBars contains all running progressbars.
// Generally, there should only be one active Progressbar at a time.
var ActiveProgressBars []*Progressbar

var (
	// DefaultProgressbar is the default progressbar.
	DefaultProgressbar = Progressbar{
		Total:                     100,
		LineCharacter:             "█",
		LastCharacter:             "█",
		ElapsedTimeRoundingFactor: time.Second,
		BarStyle:                  Style{FgLightCyan},
		TitleStyle:                Style{FgCyan},
		ShowTitle:                 true,
		ShowCount:                 true,
		ShowPercentage:            true,
		ShowElapsedTime:           true,
		BarFiller:                 " ",
	}
)

var fade = []string{
	"FF3D3D", "FC3F3C", "F9423C", "F7453C", "F4483C", "F24B3C", "EF4D3C", "EC503C", "EA533C", "E7563C",
	"E5593C", "E25C3C", "DF5E3C", "DD613C", "DA643C", "D8673C", "D56A3B", "D36D3B", "D06F3B", "CD723B",
	"CB753B", "C8783B", "C67B3B", "C37D3B", "C0803B", "BE833B", "BB863B", "B9893B", "B68C3B", "B38E3B",
	"B1913B", "AE943B", "AC973A", "A99A3A", "A79D3A", "A49F3A", "A1A23A", "9FA53A", "9CA83A", "9AAB3A",
	"97AE3A", "94B03A", "92B33A", "8FB63A", "8DB93A", "8ABC3A", "87BE3A", "85C13A", "82C439", "80C739",
	"7DCA39", "7BCD39", "78CF39", "75D239", "73D539", "70D839", "6EDB39", "6BDE39", "68E039", "66E339",
	"63E639", "61E939", "5EEC39", "5CEF39",
}

// Progressbar shows a progress animation in the terminal.
type Progressbar struct {
	Title                     string
	Total                     int
	Current                   int
	UpdateDelay               time.Duration
	LineCharacter             string
	LastCharacter             string
	ElapsedTimeRoundingFactor time.Duration
	BarFiller                 string

	ShowElapsedTime bool
	ShowCount       bool
	ShowTitle       bool
	ShowPercentage  bool

	TitleStyle Style
	BarStyle   Style

	IsActive bool

	startedAt time.Time
}

// WithTitle sets the name of the progressbar.
func (p Progressbar) WithTitle(name string) *Progressbar {
	p.Title = name
	return &p
}

// WithTotal sets the total value of the progressbar.
func (p Progressbar) WithTotal(total int) *Progressbar {
	p.Total = total
	return &p
}

// WithCurrent sets the current value of the progressbar.
func (p Progressbar) WithCurrent(current int) *Progressbar {
	p.Current = current
	return &p
}

// WithUpdateDelay sets the update delay of the progressbar.
func (p Progressbar) WithUpdateDelay(delay time.Duration) *Progressbar {
	p.UpdateDelay = delay
	return &p
}

// WithLineCharacter sets the line character of the progressbar.
func (p Progressbar) WithLineCharacter(char string) *Progressbar {
	p.LineCharacter = char
	return &p
}

// WithLastCharacter sets the last character of the progressbar.
func (p Progressbar) WithLastCharacter(char string) *Progressbar {
	p.LastCharacter = char
	return &p
}

// WithElapsedTimeRoundingFactor sets the rounding factor of the elapsed time.
func (p Progressbar) WithElapsedTimeRoundingFactor(duration time.Duration) *Progressbar {
	p.ElapsedTimeRoundingFactor = duration
	return &p
}

// WithShowElapsedTime sets if the elapsed time should be displayed in the progressbar.
func (p Progressbar) WithShowElapsedTime(b ...bool) *Progressbar {
	p.ShowElapsedTime = internal.WithBoolean(b)
	return &p
}

// WithShowCount sets if the total and current count should be displayed in the progressbar.
func (p Progressbar) WithShowCount(b ...bool) *Progressbar {
	p.ShowCount = internal.WithBoolean(b)
	return &p
}

// WithShowTitle sets if the title should be displayed in the progressbar.
func (p Progressbar) WithShowTitle(b ...bool) *Progressbar {
	p.ShowTitle = internal.WithBoolean(b)
	return &p
}

// WithShowPercentage sets if the completed percentage should be displayed in the progressbar.
func (p Progressbar) WithShowPercentage(b ...bool) *Progressbar {
	p.ShowPercentage = internal.WithBoolean(b)
	return &p
}

// WithTitleStyle sets the style of the title.
func (p Progressbar) WithTitleStyle(style Style) *Progressbar {
	p.TitleStyle = style
	return &p
}

// WithBarStyle sets the style of the bar.
func (p Progressbar) WithBarStyle(style Style) *Progressbar {
	p.BarStyle = style
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

	var before string
	var after string

	width := GetTerminalWidth()
	currentPercentage := int(internal.PercentageRound(float64(int64(p.Total)), float64(int64(p.Current)), float64(width)))

	decoratorCount := Gray("[") + LightWhite(p.Current) + Gray("/") + LightWhite(p.Total) + Gray("]")
	decoratorCurrentPercentage := color.HEX(fade[int(0.63*float64(currentPercentage))]).Sprint(strconv.Itoa(currentPercentage) + "%")

	decoratorTitle := p.TitleStyle.Sprint(p.Title)

	if p.ShowTitle {
		before += decoratorTitle + " "
	}
	if p.ShowCount {
		before += decoratorCount + " "
	}

	after += " "

	if p.ShowPercentage {
		after += decoratorCurrentPercentage + " "
	}
	if p.ShowElapsedTime {
		after += "| " + p.parseElapsedTime()
	}

	barMaxLength := width - len(RemoveColorFromString(before)) - len(RemoveColorFromString(after)) - 1
	barCurrentLength := (p.Current * barMaxLength) / p.Total
	barFiller := strings.Repeat(p.BarFiller, barMaxLength-barCurrentLength)

	bar := p.BarStyle.Sprint(strings.Repeat(p.LineCharacter, barCurrentLength)+p.LastCharacter) + barFiller
	Printo(before + bar + after)

	if p.Current == p.Total {
		p.Stop()
		Println()
	}
	return p
}

// Start the progressbar.
func (p Progressbar) Start() *Progressbar {
	p.IsActive = true
	ActiveProgressBars = append(ActiveProgressBars, &p)
	p.startedAt = time.Now()

	if p.UpdateDelay == 0 {
		p.UpdateDelay = time.Millisecond * 100
	}

	p.Add(0)

	return &p
}

// Stop the progressbar.
func (p *Progressbar) Stop() *Progressbar {
	p.IsActive = false
	for i, bar := range ActiveProgressBars {
		if p == bar {
			ActiveProgressBars = append(ActiveProgressBars[:i], ActiveProgressBars[i+1:]...)
		}
	}
	return p
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (p Progressbar) GenericStart() LivePrinter {
	return p.Start()
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (p Progressbar) GenericStop() LivePrinter {
	return p.Stop()
}

// GetElapsedTime returns the elapsed time, since the progressbar was started.
func (p *Progressbar) GetElapsedTime() time.Duration {
	return time.Since(p.startedAt)
}

func (p *Progressbar) parseElapsedTime() string {
	s := p.GetElapsedTime().Round(p.ElapsedTimeRoundingFactor).String()
	return s
}
