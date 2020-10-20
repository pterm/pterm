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
		BarStyle:                  &ThemeDefault.ProgressbarBarStyle,
		TitleStyle:                &ThemeDefault.ProgressbarTitleStyle,
		ShowTitle:                 true,
		ShowCount:                 true,
		ShowPercentage:            true,
		ShowElapsedTime:           true,
		BarFiller:                 " ",
	}
)

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
	RemoveWhenDone  bool

	TitleStyle *Style
	BarStyle   *Style

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
func (p Progressbar) WithTitleStyle(style *Style) *Progressbar {
	p.TitleStyle = style
	return &p
}

// WithBarStyle sets the style of the bar.
func (p Progressbar) WithBarStyle(style *Style) *Progressbar {
	p.BarStyle = style
	return &p
}

// WithRemoveWhenDone sets if the progressbar should be removed when it is done.
func (p Progressbar) WithRemoveWhenDone(b ...bool) *Progressbar {
	p.RemoveWhenDone = internal.WithBoolean(b)
	return &p
}

// Increment current value by one.
func (p *Progressbar) Increment() *Progressbar {
	p.Add(1)
	return p
}

// Add to current value.
func (p *Progressbar) Add(count int) *Progressbar {
	if p.TitleStyle == nil {
		p.TitleStyle = NewStyle()
	}
	if p.BarStyle == nil {
		p.BarStyle = NewStyle()
	}

	if p.Total == 0 {
		return nil
	}

	p.Current += count

	var before string
	var after string

	width := GetTerminalWidth()
	currentPercentage := int(internal.PercentageRound(float64(int64(p.Total)), float64(int64(p.Current))))

	decoratorCount := Gray("[") + LightWhite(p.Current) + Gray("/") + LightWhite(p.Total) + Gray("]")

	decoratorCurrentPercentage := color.RGB(NewRGB(255, 0, 0).Fade(0, float32(p.Total), float32(p.Current), NewRGB(0, 255, 0)).GetValues()).
		Sprint(strconv.Itoa(currentPercentage) + "%")

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
		if p.RemoveWhenDone {
			clearLine()
		} else {
			Println()
		}
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
func (p Progressbar) GenericStart() *LivePrinter {
	lp := LivePrinter(p.Start())
	return &lp
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (p Progressbar) GenericStop() *LivePrinter {
	lp := LivePrinter(p.Stop())
	return &lp
}

// GetElapsedTime returns the elapsed time, since the progressbar was started.
func (p *Progressbar) GetElapsedTime() time.Duration {
	return time.Since(p.startedAt)
}

func (p *Progressbar) parseElapsedTime() string {
	s := p.GetElapsedTime().Round(p.ElapsedTimeRoundingFactor).String()
	return s
}
