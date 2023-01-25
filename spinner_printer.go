package pterm

import (
	"io"
	"sync"
	"time"

	"github.com/pterm/pterm/internal"
	"go.uber.org/atomic"
)

type atomicActiveSpinnerPrinters struct {
	printers []*SpinnerPrinter
	lock     *sync.Mutex
}

var (
	// DefaultSpinner is the default SpinnerPrinter.
	DefaultSpinner = SpinnerPrinter{
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
	}

	activeSpinnerPrinters = atomicActiveSpinnerPrinters{
		printers: []*SpinnerPrinter{},
		lock:     &sync.Mutex{},
	}
)

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

	startedAt       time.Time
	currentSequence *atomic.String

	// Thread-safe versions of existing variables used internally
	atomicIsActive *atomic.Bool
	atomicText     *atomic.String

	Writer io.Writer
}

// Lazy init used to initialize thread-safe variables
func (s *SpinnerPrinter) lazyInit() {
	if s.atomicIsActive == nil {
		s.atomicIsActive = atomic.NewBool(s.IsActive)
	}
	if s.atomicText == nil {
		s.atomicText = atomic.NewString(s.Text)
	}
	if s.currentSequence == nil {
		s.currentSequence = atomic.NewString("")
	}
}

// WithText adds a text to the SpinnerPrinter.
func (s SpinnerPrinter) WithText(text string) *SpinnerPrinter {
	s.lazyInit()
	s.atomicText.Store(text)
	// We still set Text here so it is available to the users, it is not read anywhere
	s.Text = text
	return &s
}

// WithSequence adds a sequence to the SpinnerPrinter.
func (s SpinnerPrinter) WithSequence(sequence ...string) *SpinnerPrinter {
	s.lazyInit()
	s.Sequence = sequence
	return &s
}

// WithStyle adds a style to the SpinnerPrinter.
func (s SpinnerPrinter) WithStyle(style *Style) *SpinnerPrinter {
	s.lazyInit()
	s.Style = style
	return &s
}

// WithDelay adds a delay to the SpinnerPrinter.
func (s SpinnerPrinter) WithDelay(delay time.Duration) *SpinnerPrinter {
	s.lazyInit()
	s.Delay = delay
	return &s
}

// WithMessageStyle adds a style to the SpinnerPrinter message.
func (s SpinnerPrinter) WithMessageStyle(style *Style) *SpinnerPrinter {
	s.lazyInit()
	s.MessageStyle = style
	return &s
}

// WithRemoveWhenDone removes the SpinnerPrinter after it is done.
func (s SpinnerPrinter) WithRemoveWhenDone(b ...bool) *SpinnerPrinter {
	s.lazyInit()
	s.RemoveWhenDone = internal.WithBoolean(b)
	return &s
}

// WithShowTimer shows how long the spinner is running.
func (s SpinnerPrinter) WithShowTimer(b ...bool) *SpinnerPrinter {
	s.lazyInit()
	s.ShowTimer = internal.WithBoolean(b)
	return &s
}

// WithTimerRoundingFactor sets the rounding factor for the timer.
func (s SpinnerPrinter) WithTimerRoundingFactor(factor time.Duration) *SpinnerPrinter {
	s.lazyInit()
	s.TimerRoundingFactor = factor
	return &s
}

// WithTimerStyle adds a style to the SpinnerPrinter timer.
func (s SpinnerPrinter) WithTimerStyle(style *Style) *SpinnerPrinter {
	s.lazyInit()
	s.TimerStyle = style
	return &s
}

// WithWriter sets the custom Writer.
func (s SpinnerPrinter) WithWriter(writer io.Writer) *SpinnerPrinter {
	s.lazyInit()
	s.Writer = writer
	return &s
}

// UpdateText updates the message of the active SpinnerPrinter.
// Can be used live.
func (s *SpinnerPrinter) UpdateText(text string) {
	s.lazyInit()
	s.atomicText.Store(text)
	// We still set Text here so it is available to the users, it is not read anywhere
	s.Text = text
	if !RawOutput.Load() {
		fClearLine(s.Writer)
		Fprinto(s.Writer, s.Style.Sprint(s.currentSequence.Load())+" "+s.MessageStyle.Sprint(s.atomicText.Load()))
	}
	if RawOutput.Load() {
		Fprintln(s.Writer, s.atomicText.Load())
	}
}

// Start the SpinnerPrinter.
func (s SpinnerPrinter) Start(text ...interface{}) (*SpinnerPrinter, error) {
	s.lazyInit()
	s.atomicIsActive.Store(true)
	s.IsActive = true
	// We still set IsActive here so it is available to the users, it is not read anywhere
	s.startedAt = time.Now()

	activeSpinnerPrinters.lock.Lock()
	activeSpinnerPrinters.printers = append(activeSpinnerPrinters.printers, &s)
	activeSpinnerPrinters.lock.Unlock()

	if len(text) != 0 {
		s.atomicText.Store(Sprint(text...))
	}

	if RawOutput.Load() {
		Fprintln(s.Writer, s.atomicText.Load())
	}

	go func() {
		for s.atomicIsActive.Load() {
			for _, seq := range s.Sequence {
				if !s.atomicIsActive.Load() || RawOutput.Load() {
					continue
				}

				var timer string
				if s.ShowTimer {
					timer = " (" + time.Since(s.startedAt).Round(s.TimerRoundingFactor).String() + ")"
				}
				fClearLine(s.Writer)
				Fprinto(s.Writer, s.Style.Sprint(seq)+" "+s.MessageStyle.Sprint(s.atomicText.Load())+s.TimerStyle.Sprint(timer))
				s.currentSequence.Store(seq)
				time.Sleep(s.Delay)
			}
		}
	}()
	return &s, nil
}

// Stop terminates the SpinnerPrinter immediately.
// The SpinnerPrinter will not resolve into anything.
func (s *SpinnerPrinter) Stop() error {
	s.lazyInit()
	if !s.atomicIsActive.Load() {
		return nil
	}
	s.atomicIsActive.Store(false)
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
	s.lazyInit()
	_, _ = s.Start()
	lp := LivePrinter(s)
	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (s *SpinnerPrinter) GenericStop() (*LivePrinter, error) {
	s.lazyInit()
	_ = s.Stop()
	lp := LivePrinter(s)
	return &lp, nil
}

// Info displays an info message
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Info(message ...interface{}) {
	s.lazyInit()
	if s.InfoPrinter == nil {
		s.InfoPrinter = &Info
	}

	if len(message) == 0 {
		message = []interface{}{s.atomicText.Load()}
	}
	fClearLine(s.Writer)
	Fprinto(s.Writer, s.InfoPrinter.Sprint(message...))
	_ = s.Stop()
}

// Success displays the success printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Success(message ...interface{}) {
	s.lazyInit()
	if s.SuccessPrinter == nil {
		s.SuccessPrinter = &Success
	}

	if len(message) == 0 {
		message = []interface{}{s.atomicText.Load()}
	}
	fClearLine(s.Writer)
	Fprinto(s.Writer, s.SuccessPrinter.Sprint(message...))
	_ = s.Stop()
}

// Fail displays the fail printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Fail(message ...interface{}) {
	s.lazyInit()
	if s.FailPrinter == nil {
		s.FailPrinter = &Error
	}

	if len(message) == 0 {
		message = []interface{}{s.atomicText.Load()}
	}
	fClearLine(s.Writer)
	Fprinto(s.Writer, s.FailPrinter.Sprint(message...))
	_ = s.Stop()
}

// Warning displays the warning printer.
// If no message is given, the text of the SpinnerPrinter will be reused as the default message.
func (s *SpinnerPrinter) Warning(message ...interface{}) {
	s.lazyInit()
	if s.WarningPrinter == nil {
		s.WarningPrinter = &Warning
	}

	if len(message) == 0 {
		message = []interface{}{s.atomicText.Load()}
	}
	fClearLine(s.Writer)
	Fprinto(s.Writer, s.WarningPrinter.Sprint(message...))
	_ = s.Stop()
}
