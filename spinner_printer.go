package pterm

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/pterm/pterm/internal"
)

var activeSpinnerPrinters []*SpinnerPrinter

// DefaultSpinner is the default SpinnerPrinter.
var DefaultSpinner = SpinnerPrinter{
	Sequence:            []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	Style:               &ThemeDefault.SpinnerStyle,
	Delay:               time.Millisecond * 100,
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

	startedAt       time.Time
	currentSequence string

	Writer io.Writer

	// mu serializes access to the printer's mutable state once Start has run.
	// It is a pointer so the value-receiver With* methods can copy the struct
	// without tripping go vet's copylocks check; Start lazily allocates it.
	mu *sync.RWMutex
	// wg lets Stop block until the animation goroutine has fully exited,
	// avoiding leaked goroutines that would otherwise accumulate across
	// successive Start/Stop cycles (especially in tests).
	wg *sync.WaitGroup
	// stopCh is closed by Stop to interrupt the animation goroutine's sleep;
	// without it the goroutine would not exit until its next Delay tick fires.
	stopCh   chan struct{}
	stopOnce *sync.Once
}

// lock locks the printer's runtime mutex, allocating it on first use so the
// value-receiver builders can be used safely before Start.
func (s *SpinnerPrinter) lock() {
	if s.mu == nil {
		s.mu = &sync.RWMutex{}
	}

	s.mu.Lock()
}

// unlock releases the printer's runtime mutex.
func (s *SpinnerPrinter) unlock() {
	s.mu.Unlock()
}

// rlock takes a read lock on the printer's runtime mutex, allocating it on
// first use.
func (s *SpinnerPrinter) rlock() {
	if s.mu == nil {
		s.mu = &sync.RWMutex{}
	}

	s.mu.RLock()
}

// runlock releases a read lock taken by rlock.
func (s *SpinnerPrinter) runlock() {
	s.mu.RUnlock()
}

// isActive returns whether the spinner is currently active. Used internally by
// print.go to inspect spinners from another goroutine without racing with
// Start/Stop.
func (s *SpinnerPrinter) isActive() bool {
	s.rlock()
	defer s.runlock()

	return s.IsActive
}

// writer returns the configured writer under the runtime lock.
func (s *SpinnerPrinter) writer() io.Writer {
	s.rlock()
	defer s.runlock()

	return s.Writer
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
func (s SpinnerPrinter) WithWriter(writer io.Writer) *SpinnerPrinter {
	s.Writer = writer
	return &s
}

// SetWriter sets the custom Writer.
func (s *SpinnerPrinter) SetWriter(writer io.Writer) {
	s.Writer = writer
}

// ResetTimer resets the timer of the SpinnerPrinter.
func (s *SpinnerPrinter) ResetTimer() {
	s.lock()
	defer s.unlock()

	s.startedAt = time.Now()
}

// SetStartedAt sets the time when the SpinnerPrinter started.
func (s *SpinnerPrinter) SetStartedAt(t time.Time) {
	s.lock()
	defer s.unlock()

	s.startedAt = t
}

// UpdateText updates the message of the active SpinnerPrinter.
// Can be used live.
func (s *SpinnerPrinter) UpdateText(text string) {
	s.lock()
	s.Text = text
	style := s.Style
	currentSequence := s.currentSequence
	messageStyle := s.MessageStyle
	writer := s.Writer
	s.unlock()

	if !rawOutput() {
		Fprinto(writer, "\033[K"+style.Sprint(currentSequence)+" "+messageStyle.Sprint(text))
	} else {
		Fprintln(writer, text)
	}
}

// Start the SpinnerPrinter.
func (s SpinnerPrinter) Start(text ...any) (*SpinnerPrinter, error) {
	// Allocate the runtime mutex eagerly so the spinner goroutine and any
	// subsequent user calls share the same lock without racing on lazy init.
	s.mu = &sync.RWMutex{}
	s.wg = &sync.WaitGroup{}
	s.stopCh = make(chan struct{})
	s.stopOnce = &sync.Once{}

	s.IsActive = true
	s.startedAt = time.Now()

	if len(text) != 0 {
		s.Text = Sprint(text...)
	}

	registerSpinner(&s)

	sp := &s
	sp.wg.Go(func() {
		sp.runAnimation()
	})

	return sp, nil
}

// sleepOrStop sleeps for d, returning early if stopCh is closed.
// Returns true if the spinner was asked to stop.
func (s *SpinnerPrinter) sleepOrStop(d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()

	select {
	case <-s.stopCh:
		return true
	case <-t.C:
		return false
	}
}

// runAnimation drives the spinner frames. Reads fields under the runtime lock
// so it is safe to call concurrently with UpdateText/Stop.
func (s *SpinnerPrinter) runAnimation() {
	for s.isActive() {
		s.rlock()
		sequence := s.Sequence
		s.runlock()

		for _, seq := range sequence {
			if !s.isActive() {
				return
			}

			if rawOutput() {
				s.rlock()
				delay := s.Delay
				s.runlock()

				if s.sleepOrStop(delay) {
					return
				}

				continue
			}

			s.lock()
			style := s.Style
			messageStyle := s.MessageStyle
			timerStyle := s.TimerStyle
			text := s.Text
			showTimer := s.ShowTimer
			startedAt := s.startedAt
			rounding := s.TimerRoundingFactor
			delay := s.Delay
			writer := s.Writer
			s.currentSequence = seq
			s.unlock()

			var timer string
			if showTimer {
				timer = " (" + time.Since(startedAt).Round(rounding).String() + ")"
			}

			Fprinto(writer, style.Sprint(seq)+" "+messageStyle.Sprint(text)+timerStyle.Sprint(timer))

			if s.sleepOrStop(delay) {
				return
			}
		}
	}
}

// Stop terminates the SpinnerPrinter immediately.
// The SpinnerPrinter will not resolve into anything.
//
// Stop signals the animation goroutine to exit (via stopCh) and waits for it
// to finish before returning. Without this synchronization, leaked goroutines
// would accumulate across Start/Stop cycles in tests and degrade the runtime
// of subsequent tests under -race.
func (s *SpinnerPrinter) Stop() error {
	s.lock()
	wasActive := s.IsActive
	s.IsActive = false
	removeWhenDone := s.RemoveWhenDone
	writer := s.Writer
	stopOnce := s.stopOnce
	stopCh := s.stopCh
	wg := s.wg
	s.unlock()

	// Signal the animation goroutine to stop. Stop may be called multiple
	// times (e.g. by Warning followed by an explicit defer Stop), so the
	// channel is closed via sync.Once.
	if stopOnce != nil {
		stopOnce.Do(func() { close(stopCh) })
	}
	// Wait for the goroutine to exit so any further writes from the spinner
	// cannot race with the caller draining the writer.
	if wg != nil {
		wg.Wait()
	}

	if !wasActive {
		return nil
	}

	if rawOutput() {
		return nil
	}

	if removeWhenDone {
		fClearLine(writer)
		Fprinto(writer)
	} else {
		Fprintln(writer)
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
