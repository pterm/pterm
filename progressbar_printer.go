package pterm

import (
	"strconv"
	"strings"
	"time"

	"github.com/gookit/color"

	"github.com/pterm/pterm/internal"
)

// ActiveSoloProgressBarPrinters contains all solo running ProgressbarPrinters.
var ActiveSoloProgressBarPrinters []*ProgressbarPrinter

// ActiveMultiProgressBarPrinters contains all running ProgressbarPrinters that should be printed together.
var ActiveMultiProgressBarPrinters []*ProgressbarPrinter

var (
	// DefaultProgressbar is the default ProgressbarPrinter.
	DefaultProgressbar = ProgressbarPrinter{
		Total:                     100,
		BarCharacter:              "█",
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

// ProgressbarPrinter shows a progress animation in the terminal.
type ProgressbarPrinter struct {
	Title                     string
	Total                     int
	Current                   int
	BarCharacter              string
	LastCharacter             string
	ElapsedTimeRoundingFactor time.Duration
	BarFiller                 string
	area                      *AreaPrinter
	timePassed                string

	ShowElapsedTime bool
	ShowCount       bool
	ShowTitle       bool
	ShowPercentage  bool
	RemoveWhenDone  bool

	TitleStyle *Style
	BarStyle   *Style

	IsActive bool

	PrintTogether bool

	startedAt time.Time
}

// WithTitle sets the name of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithTitle(name string) *ProgressbarPrinter {
	p.Title = name
	return &p
}

// WithTotal sets the total value of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithTotal(total int) *ProgressbarPrinter {
	p.Total = total
	return &p
}

// WithCurrent sets the current value of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithCurrent(current int) *ProgressbarPrinter {
	p.Current = current
	return &p
}

// WithBarCharacter sets the bar character of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithBarCharacter(char string) *ProgressbarPrinter {
	p.BarCharacter = char
	return &p
}

// WithLastCharacter sets the last character of the ProgressbarPrinter.
func (p ProgressbarPrinter) WithLastCharacter(char string) *ProgressbarPrinter {
	p.LastCharacter = char
	return &p
}

// WithElapsedTimeRoundingFactor sets the rounding factor of the elapsed time.
func (p ProgressbarPrinter) WithElapsedTimeRoundingFactor(duration time.Duration) *ProgressbarPrinter {
	p.ElapsedTimeRoundingFactor = duration
	return &p
}

// WithShowElapsedTime sets if the elapsed time should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowElapsedTime(b ...bool) *ProgressbarPrinter {
	p.ShowElapsedTime = internal.WithBoolean(b)
	return &p
}

// WithShowCount sets if the total and current count should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowCount(b ...bool) *ProgressbarPrinter {
	p.ShowCount = internal.WithBoolean(b)
	return &p
}

// WithShowTitle sets if the title should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowTitle(b ...bool) *ProgressbarPrinter {
	p.ShowTitle = internal.WithBoolean(b)
	return &p
}

// WithShowPercentage sets if the completed percentage should be displayed in the ProgressbarPrinter.
func (p ProgressbarPrinter) WithShowPercentage(b ...bool) *ProgressbarPrinter {
	p.ShowPercentage = internal.WithBoolean(b)
	return &p
}

// WithTitleStyle sets the style of the title.
func (p ProgressbarPrinter) WithTitleStyle(style *Style) *ProgressbarPrinter {
	p.TitleStyle = style
	return &p
}

// WithBarStyle sets the style of the bar.
func (p ProgressbarPrinter) WithBarStyle(style *Style) *ProgressbarPrinter {
	p.BarStyle = style
	return &p
}

// WithRemoveWhenDone sets if the ProgressbarPrinter should be removed when it is done.
func (p ProgressbarPrinter) WithRemoveWhenDone(b ...bool) *ProgressbarPrinter {
	p.RemoveWhenDone = internal.WithBoolean(b)
	return &p
}

// WithPrintTogether sets if the ProgressbarPrinter should be printed with other ProgressbarPrinters.
func (p ProgressbarPrinter) WithPrintTogether(area *AreaPrinter, b ...bool) *ProgressbarPrinter {
	p.PrintTogether = internal.WithBoolean(b)
	p.area = area
	return &p
}

// Increment current value by one.
func (p *ProgressbarPrinter) Increment() *ProgressbarPrinter {
	p.Add(1)
	return p
}

// Add to current value.
func (p *ProgressbarPrinter) Add(count int) *ProgressbarPrinter {
	var progressBars string
	if p.PrintTogether {
		p.area.Update(Sprinto(progressBars))
		for i, printer := range ActiveMultiProgressBarPrinters {
			if !RawOutput {
				progressBars += retProgressbarString(printer, count, printer == p)
				if i < len(ActiveMultiProgressBarPrinters)-1 {
					progressBars += "\n"
				}
			}
		}
		p.area.Update(progressBars)
	} else {
		Printo(retProgressbarString(p, count, true))
	}
	return p
}

func retProgressbarString(printer *ProgressbarPrinter, count int, same bool) string {
	var ret string
	if printer.TitleStyle == nil {
		printer.TitleStyle = NewStyle()
	}
	if printer.BarStyle == nil {
		printer.BarStyle = NewStyle()
	}

	if printer.Total == 0 {
		return ""
	}

	if same {
		printer.Current += count
	}

	var before string
	var after string

	width := GetTerminalWidth()
	currentPercentage := int(internal.PercentageRound(float64(int64(printer.Total)), float64(int64(printer.Current))))

	decoratorCount := Gray("[") + LightWhite(printer.Current) + Gray("/") + LightWhite(printer.Total) + Gray("]")

	decoratorCurrentPercentage := color.RGB(NewRGB(255, 0, 0).Fade(0, float32(printer.Total), float32(printer.Current), NewRGB(0, 255, 0)).GetValues()).
		Sprint(strconv.Itoa(currentPercentage) + "%")

	decoratorTitle := printer.TitleStyle.Sprint(printer.Title)

	if printer.ShowTitle {
		before += decoratorTitle + " "
	}
	if printer.ShowCount {
		before += decoratorCount + " "
	}

	after += " "

	if printer.ShowPercentage {
		after += decoratorCurrentPercentage + " "
	}
	if printer.ShowElapsedTime {
		after += "| " + printer.parseElapsedTime()
	}

	barMaxLength := width - len(RemoveColorFromString(before)) - len(RemoveColorFromString(after)) - 1
	barCurrentLength := (printer.Current * barMaxLength) / printer.Total
	barFiller := strings.Repeat(printer.BarFiller, barMaxLength-barCurrentLength)

	bar := printer.BarStyle.Sprint(strings.Repeat(printer.BarCharacter, barCurrentLength)+printer.LastCharacter) + barFiller

	if same && printer.Current == printer.Total {
		printer.Stop()
	}

	if !RawOutput {
		ret += before + bar + after
	}

	return ret
}

// Start the ProgressbarPrinter.
func (p ProgressbarPrinter) Start() (*ProgressbarPrinter, error) {
	if RawOutput && p.ShowTitle {
		Println(p.Title)
	}
	p.IsActive = true
	if p.PrintTogether {
		ActiveMultiProgressBarPrinters = append(ActiveMultiProgressBarPrinters, &p)
	} else {
		ActiveSoloProgressBarPrinters = append(ActiveSoloProgressBarPrinters, &p)
	}
	p.startedAt = time.Now()

	p.Add(0)

	return &p, nil
}

// Stop the ProgressbarPrinter.
func (p *ProgressbarPrinter) Stop() (*ProgressbarPrinter, error) {
	if !p.IsActive {
		return p, nil
	}
	p.IsActive = false
	if p.RemoveWhenDone {
		clearLine()
		Printo()
	} else {
		Println()
	}
	return p, nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (p ProgressbarPrinter) GenericStart() (*LivePrinter, error) {
	p2, _ := p.Start()
	lp := LivePrinter(p2)
	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (p ProgressbarPrinter) GenericStop() (*LivePrinter, error) {
	p2, _ := p.Stop()
	lp := LivePrinter(p2)
	return &lp, nil
}

// GetElapsedTime returns the elapsed time, since the ProgressbarPrinter was started.
func (p *ProgressbarPrinter) GetElapsedTime() time.Duration {
	return time.Since(p.startedAt)
}

func (p *ProgressbarPrinter) parseElapsedTime() string {
	if p.IsActive {
		p.timePassed = p.GetElapsedTime().Round(p.ElapsedTimeRoundingFactor).String()
	}
	return p.timePassed
}
