package pterm

import (
	"io"
	"os"
	"sync"
	"sync/atomic"
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
	InfoPrinter:         &Info,
	SuccessPrinter:      &Success,
	FailPrinter:         &Error,
	WarningPrinter:      &Warning,
	Writer:              os.Stderr,
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
	InfoPrinter         TextPrinter
	SuccessPrinter      TextPrinter
	FailPrinter         TextPrinter
	WarningPrinter      TextPrinter
	RemoveWhenDone      bool
	ShowTimer           bool
	TimerRoundingFactor time.Duration
	TimerStyle          *Style

	IsActive bool

	// active mirrors IsActive but is safe to read concurrently with Stop.
	// The animation goroutine reads it instead of IsActive to avoid the data
	// race that -race flags between the goroutine and Stop().
	active atomic.Bool

	// mu guards Text and currentSequence so UpdateText can safely run while
	// the animation goroutine is reading them.
	mu              sync.Mutex
	startedAt       time.Time
	currentSequence string

	Writer io.Writer
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

// WithStartedAt sets the time when the SpinnerPrinter started.
func (s SpinnerPrinter) WithStartedAt(t time.Time) *SpinnerPrinter {
	s.startedAt = t
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

// WithWriter sets the custom Writer.
func (p SpinnerPrinter) WithWriter(writer io.Writer) *SpinnerPrinter {
	p.Writer = writer
	return &p
}

// SetWriter sets the custom Writer.
func (p *SpinnerPrinter) SetWriter(writer io.Writer) {
	p.Writer = writer
}

// ResetTimer resets the timer of the SpinnerPrinter.
func (s *SpinnerPrinter) ResetTimer() {
	s.startedAt = time.Now()
}

// SetStartedAt sets the time when the SpinnerPrinter started.
func (s *SpinnerPrinter) SetStartedAt(t time.Time) {
	s.startedAt = t
}

// UpdateText updates the message of the active SpinnerPrinter.
// Can be used live.
func (s *SpinnerPrinter) UpdateText(text string) {
	s.mu.Lock()
	s.Text = text
	currentSequence := s.currentSequence
	s.mu.Unlock()
	if !RawOutput {
		Fprinto(s.Writer, "\033[K"+s.Style.Sprint(currentSequence)+" "+s.MessageStyle.Sprint(text))
	} else {
		Fprintln(s.Writer, text)
	}
}

// Start the SpinnerPrinter.
func (s SpinnerPrinter) Start(text ...any) (*SpinnerPrinter, error) {
	s.IsActive = true
	s.active.Store(true)
	s.startedAt = time.Now()
	activeSpinnerPrinters = append(activeSpinnerPrinters, &s)

	if len(text) != 0 {
		s.Text = Sprint(text...)
	}

	sp := &s
	go func() {
		for sp.active.Load() {
			for _, seq := range sp.Sequence {
				if !sp.active.Load() {
					continue
				}

				if RawOutput {
					time.Sleep(sp.Delay)
					continue
				}

				var timer string
				if sp.ShowTimer {
					timer = " (" + time.Since(sp.startedAt).Round(sp.TimerRoundingFactor).String() + ")"
				}

				sp.mu.Lock()
				text := sp.Text
				sp.currentSequence = seq
				sp.mu.Unlock()

				Fprinto(sp.Writer, sp.Style.Sprint(seq)+" "+sp.MessageStyle.Sprint(text)+sp.TimerStyle.Sprint(timer))
				time.Sleep(sp.Delay)
			}
		}
	}()

	return sp, nil
}

// Stop terminates the SpinnerPrinter immediately.
// The SpinnerPrinter will not resolve into anything.
func (s *SpinnerPrinter) Stop() error {
	if !s.active.Load() {
		return nil
	}

	s.IsActive = false
	s.active.Store(false)

	if RawOutput {
		return nil
	}

	if s.RemoveWhenDone {
		fClearLine(s.Writer)
		Fprinto(s.Writer)
	} else {
		Fprintln(s.Writer)
	}

	return nil
}

// GenericStart runs Start, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Start instead of this in your program.
func (s *SpinnerPrinter) GenericStart() (*LivePrinter, error) {
	p2, _ := s.Start()
	lp := LivePrinter(p2)

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

// Info displays an info message
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Info(message ...any) {
	if s.InfoPrinter == nil {
		s.InfoPrinter = &Info
	}

	if len(message) == 0 {
		message = []any{s.Text}
	}

	fClearLine(s.Writer)

	Fprinto(s.Writer, s.InfoPrinter.Sprint(message...))
	_ = s.Stop()
}

// Success displays the success printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Success(message ...any) {
	if s.SuccessPrinter == nil {
		s.SuccessPrinter = &Success
	}

	if len(message) == 0 {
		message = []any{s.Text}
	}

	fClearLine(s.Writer)

	Fprinto(s.Writer, s.SuccessPrinter.Sprint(message...))
	_ = s.Stop()
}

// Fail displays the fail printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Fail(message ...any) {
	if s.FailPrinter == nil {
		s.FailPrinter = &Error
	}

	if len(message) == 0 {
		message = []any{s.Text}
	}

	fClearLine(s.Writer)

	Fprinto(s.Writer, s.FailPrinter.Sprint(message...))
	_ = s.Stop()
}

// Warning displays the warning printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Warning(message ...any) {
	if s.WarningPrinter == nil {
		s.WarningPrinter = &Warning
	}

	if len(message) == 0 {
		message = []any{s.Text}
	}

	fClearLine(s.Writer)

	Fprinto(s.Writer, s.WarningPrinter.Sprint(message...))
	_ = s.Stop()
}
