package pterm

import (
	"time"

	"github.com/pterm/pterm/internal"
)

var activeSpinnerPrinters []*SpinnerPrinter

// DefaultSpinner is the default SpinnerPrinter.
var DefaultSpinner = SpinnerPrinter{
	Sequence:            []string{"▀ ", " ▀", " ▄", "▄ "},
	Style:               &ThemeDefault.SpinnerStyle,
	Delay:               time.Millisecond * 200,
	ShowTimer:           true,
	TimerRoundingFactor: time.Second,
	TimerStyle:          &ThemeDefault.TimerStyle,
	MessageStyle:        &ThemeDefault.SpinnerTextStyle,
	SuccessPrinter:      &Success,
	FailPrinter:         &Error,
	WarningPrinter:      &Warning,
}

// SpinnerPrinter is a loading animation, which can be used if the progress is unknown.
// It's an animation loop, which can have a text and supports throwing errors or warnings.
// A TextPrinter is used to display all outputs, after the SpinnerPrinter is done.
type SpinnerPrinter struct {
	Text                string
	Sequence            []string
	Style               *Style
	Delay               time.Duration
	MessageStyle        *Style
	SuccessPrinter      TextPrinter
	FailPrinter         TextPrinter
	WarningPrinter      TextPrinter
	RemoveWhenDone      bool
	ShowTimer           bool
	TimerRoundingFactor time.Duration
	TimerStyle          *Style

	IsActive bool

	startedAt       time.Time
	currentSequence string
}

// WithText adds a text to the SpinnerPrinter.
func (s SpinnerPrinter) WithText(text string) *SpinnerPrinter {
	s.Text = text
	return &s
}

// WithSequence adds a sequence to the SpinnerPrinter.
func (s SpinnerPrinter) WithSequence(sequence ...string) *SpinnerPrinter {
	s.Sequence = sequence
	return &s
}

// WithStyle adds a style to the SpinnerPrinter.
func (s SpinnerPrinter) WithStyle(style *Style) *SpinnerPrinter {
	s.Style = style
	return &s
}

// WithDelay adds a delay to the SpinnerPrinter.
func (s SpinnerPrinter) WithDelay(delay time.Duration) *SpinnerPrinter {
	s.Delay = delay
	return &s
}

// WithMessageStyle adds a style to the SpinnerPrinter message.
func (s SpinnerPrinter) WithMessageStyle(style *Style) *SpinnerPrinter {
	s.MessageStyle = style
	return &s
}

// WithRemoveWhenDone removes the SpinnerPrinter after it is done.
func (s SpinnerPrinter) WithRemoveWhenDone(b ...bool) *SpinnerPrinter {
	s.RemoveWhenDone = internal.WithBoolean(b)
	return &s
}

// WithShowTimer shows how long the spinner is running.
func (s SpinnerPrinter) WithShowTimer(b ...bool) *SpinnerPrinter {
	s.ShowTimer = internal.WithBoolean(b)
	return &s
}

// WithTimerRoundingFactor sets the rounding factor for the timer.
func (s SpinnerPrinter) WithTimerRoundingFactor(factor time.Duration) *SpinnerPrinter {
	s.TimerRoundingFactor = factor
	return &s
}

// WithTimerStyle adds a style to the SpinnerPrinter timer.
func (s SpinnerPrinter) WithTimerStyle(style *Style) *SpinnerPrinter {
	s.TimerStyle = style
	return &s
}

// UpdateText updates the message of the active SpinnerPrinter.
// Can be used live.
func (s *SpinnerPrinter) UpdateText(text string) {
	s.Text = text
	if !RawOutput {
		clearLine()
		Printo(s.Style.Sprint(s.currentSequence) + " " + s.MessageStyle.Sprint(s.Text))
	}
	if RawOutput {
		Println(s.Text)
	}
}

// Start the SpinnerPrinter.
func (s SpinnerPrinter) Start(text ...interface{}) (*SpinnerPrinter, error) {
	s.IsActive = true
	s.startedAt = time.Now()
	activeSpinnerPrinters = append(activeSpinnerPrinters, &s)

	if len(text) != 0 {
		s.Text = Sprint(text...)
	}

	if RawOutput {
		Println(s.Text)
	}

	go func() {
		for s.IsActive {
			for _, seq := range s.Sequence {
				if !s.IsActive || RawOutput {
					continue
				}

				var timer string
				if s.ShowTimer {
					timer = " (" + time.Since(s.startedAt).Round(s.TimerRoundingFactor).String() + ")"
				}
				Printo(s.Style.Sprint(seq) + " " + s.MessageStyle.Sprint(s.Text) + s.TimerStyle.Sprint(timer))
				s.currentSequence = seq
				time.Sleep(s.Delay)
			}
		}
	}()
	return &s, nil
}

// Stop terminates the SpinnerPrinter immediately.
// The SpinnerPrinter will not resolve into anything.
func (s *SpinnerPrinter) Stop() error {
	s.IsActive = false
	if s.RemoveWhenDone {
		clearLine()
		Printo()
	} else {
		Println()
	}
	return nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (s *SpinnerPrinter) GenericStart() (*LivePrinter, error) {
	_, _ = s.Start()
	lp := LivePrinter(s)
	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (s *SpinnerPrinter) GenericStop() (*LivePrinter, error) {
	_ = s.Stop()
	lp := LivePrinter(s)
	return &lp, nil
}

// Success displays the success printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Success(message ...interface{}) {
	if s.SuccessPrinter == nil {
		s.SuccessPrinter = &Success
	}

	if len(message) == 0 {
		message = []interface{}{s.Text}
	}
	clearLine()
	Printo(s.SuccessPrinter.Sprint(message...))
	_ = s.Stop()
}

// Fail displays the fail printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Fail(message ...interface{}) {
	if s.FailPrinter == nil {
		s.FailPrinter = &Error
	}

	if len(message) == 0 {
		message = []interface{}{s.Text}
	}
	clearLine()
	Printo(s.FailPrinter.Sprint(message...))
	_ = s.Stop()
}

// Warning displays the warning printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Warning(message ...interface{}) {
	if s.WarningPrinter == nil {
		s.WarningPrinter = &Warning
	}

	if len(message) == 0 {
		message = []interface{}{s.Text}
	}
	clearLine()
	Printo(s.WarningPrinter.Sprint(message...))
	_ = s.Stop()
}
