package pterm

import (
	"time"

	"github.com/pterm/pterm/internal"
)

// DefaultSpinner is the default spinner.
var DefaultSpinner = Spinner{
	Sequence:       []string{"▀ ", " ▀", " ▄", "▄ "},
	Style:          &ThemeDefault.SpinnerStyle,
	Delay:          time.Millisecond * 200,
	MessageStyle:   &ThemeDefault.SpinnerTextStyle,
	SuccessPrinter: &Success,
	FailPrinter:    &Error,
	WarningPrinter: &Warning,
}

// Spinner is a loading animation, which can be used if the progress is unknown.
// It's an animation loop, which can have a text and supports throwing errors or warnings.
// A TextPrinter is used to display all outputs, after the spinner is done.
type Spinner struct {
	Text           string
	Sequence       []string
	Style          *Style
	Delay          time.Duration
	MessageStyle   *Style
	SuccessPrinter TextPrinter
	FailPrinter    TextPrinter
	WarningPrinter TextPrinter
	RemoveWhenDone bool

	IsActive bool
}

// WithText adds a text to the spinner.
func (s Spinner) WithText(text string) *Spinner {
	s.Text = text
	return &s
}

// WithSequence adds a sequence to the spinner.
func (s Spinner) WithSequence(sequence ...string) *Spinner {
	s.Sequence = sequence
	return &s
}

// WithStyle adds a style to the spinner.
func (s Spinner) WithStyle(style *Style) *Spinner {
	s.Style = style
	return &s
}

// WithDelay adds a delay to the spinner.
func (s Spinner) WithDelay(delay time.Duration) *Spinner {
	s.Delay = delay
	return &s
}

// WithMessageStyle adds a style to the spinner message.
func (s Spinner) WithMessageStyle(style *Style) *Spinner {
	s.MessageStyle = style
	return &s
}

// WithRemoveWhenDone removes the spinner after it is done.
func (s Spinner) WithRemoveWhenDone(b ...bool) *Spinner {
	s.RemoveWhenDone = internal.WithBoolean(b)
	return &s
}

// UpdateText updates the message of the active spinner.
// Can be used live.
func (s *Spinner) UpdateText(text string) {
	clearLine()
	s.Text = text
}

// Start the spinner.
func (s Spinner) Start(text ...interface{}) (*Spinner, error) {
	s.IsActive = true

	if len(text) != 0 {
		s.Text = Sprint(text...)
	}

	go func() {
		for s.IsActive {
			for _, seq := range s.Sequence {
				if s.IsActive {
					Printo(s.Style.Sprint(seq) + " " + s.MessageStyle.Sprint(s.Text))
					time.Sleep(s.Delay)
				}
			}
		}
	}()
	return &s, nil
}

// Stop terminates the Spinner immediately.
// The Spinner will not resolve into anything.
func (s *Spinner) Stop() error {
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
func (s *Spinner) GenericStart() (*LivePrinter, error) {
	_, err := s.Start()
	if err != nil {
		return nil, err
	}
	lp := LivePrinter(s)
	return &lp, nil
}

// GenericStop runs Stop, but returns a LivePrinter.
// This is used for the interface LivePrinter.
// You most likely want to use Stop instead of this in your program.
func (s *Spinner) GenericStop() (*LivePrinter, error) {
	err := s.Stop()
	if err != nil {
		return nil, err
	}
	lp := LivePrinter(s)
	return &lp, nil
}

// Success displays the success printer.
// If no message is given, the text of the spinner will be reused as the default message.
func (s *Spinner) Success(message ...interface{}) {
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
// If no message is given, the text of the spinner will be reused as the default message.
func (s *Spinner) Fail(message ...interface{}) {
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
// If no message is given, the text of the spinner will be reused as the default message.
func (s *Spinner) Warning(message ...interface{}) {
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
