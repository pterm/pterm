package pterm

import (
	"strings"
	"time"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"

	"github.com/pterm/pterm/internal"
)

// DefaultInteractiveConfirm is the default InteractiveConfirm printer.
// Pressing "y" will return true, "n" will return false.
// Pressing enter without typing "y" or "n" will return the configured default value (by default set to "no").
var DefaultInteractiveConfirm = InteractiveConfirmPrinter{
	DefaultValue: false,
	DefaultText:  "Please confirm",
	TextStyle:    &ThemeDefault.PrimaryStyle,
	ConfirmText:  "Yes",
	ConfirmStyle: &ThemeDefault.SuccessMessageStyle,
	RejectText:   "No",
	RejectStyle:  &ThemeDefault.ErrorMessageStyle,
	SuffixStyle:  &ThemeDefault.SecondaryStyle,
	Delimiter:    ": ",
}

// InteractiveConfirmPrinter is a printer for interactive confirm prompts.
type InteractiveConfirmPrinter struct {
	DefaultValue    bool
	DefaultText     string
	Delimiter       string
	TextStyle       *Style
	ConfirmText     string
	ConfirmStyle    *Style
	RejectText      string
	RejectStyle     *Style
	SuffixStyle     *Style
	OnInterruptFunc func()
	Timeout         time.Duration
}

// WithDefaultText sets the default text.
func (p InteractiveConfirmPrinter) WithDefaultText(text string) *InteractiveConfirmPrinter {
	p.DefaultText = text
	return &p
}

// WithDefaultValue sets the default value, which will be returned when the user presses enter without typing "y" or "n".
func (p InteractiveConfirmPrinter) WithDefaultValue(value bool) *InteractiveConfirmPrinter {
	p.DefaultValue = value
	return &p
}

// WithTextStyle sets the text style.
func (p InteractiveConfirmPrinter) WithTextStyle(style *Style) *InteractiveConfirmPrinter {
	p.TextStyle = style
	return &p
}

// WithConfirmText sets the confirmation text.
func (p InteractiveConfirmPrinter) WithConfirmText(text string) *InteractiveConfirmPrinter {
	p.ConfirmText = text
	return &p
}

// WithConfirmStyle sets the confirmation style.
func (p InteractiveConfirmPrinter) WithConfirmStyle(style *Style) *InteractiveConfirmPrinter {
	p.ConfirmStyle = style
	return &p
}

// WithRejectText sets the reject text.
func (p InteractiveConfirmPrinter) WithRejectText(text string) *InteractiveConfirmPrinter {
	p.RejectText = text
	return &p
}

// WithRejectStyle sets the reject style.
func (p InteractiveConfirmPrinter) WithRejectStyle(style *Style) *InteractiveConfirmPrinter {
	p.RejectStyle = style
	return &p
}

// WithSuffixStyle sets the suffix style.
func (p InteractiveConfirmPrinter) WithSuffixStyle(style *Style) *InteractiveConfirmPrinter {
	p.SuffixStyle = style
	return &p
}

// WithOnInterruptFunc sets the function to execute on exit of the input reader.
func (p InteractiveConfirmPrinter) WithOnInterruptFunc(exitFunc func()) *InteractiveConfirmPrinter {
	p.OnInterruptFunc = exitFunc
	return &p
}

// WithDelimiter sets the delimiter between the message and the input.
func (p InteractiveConfirmPrinter) WithDelimiter(delimiter string) *InteractiveConfirmPrinter {
	p.Delimiter = delimiter
	return &p
}

// WithTimeout sets the timeout to wait before confirming with the default value.
func (p InteractiveConfirmPrinter) WithTimeout(timeout time.Duration) *InteractiveConfirmPrinter {
	p.Timeout = timeout
	return &p
}

// Show shows the confirmation prompt.
//
// Example:
//
//	result, _ := pterm.DefaultInteractiveConfirm.Show("Are you sure?")
//	pterm.Println(result)
func (p InteractiveConfirmPrinter) Show(text ...string) (bool, error) {
	// should be the first defer statement to make sure it is executed last
	// and all the needed cleanup can be done before
	cancel, exit := internal.NewCancelationSignal(p.OnInterruptFunc)
	defer exit()

	if len(text) == 0 || text[0] == "" {
		text = []string{p.DefaultText}
	}

	p.TextStyle.Print(text[0] + " " + p.getSuffix() + p.Delimiter)
	y, n := p.getShortHandles()

	resultChan := make(chan bool)

	var interrupted bool
	go func() {
		_ = keyboard.Listen(func(keyInfo keys.Key) (stop bool, err error) {
			key := keyInfo.Code
			char := strings.ToLower(keyInfo.String())

			switch key {
			case keys.RuneKey:
				switch char {
				case y:
					p.ConfirmStyle.Print(p.ConfirmText)
					Println()
					resultChan <- true
					return true, nil
				case n:
					p.RejectStyle.Print(p.RejectText)
					Println()
					resultChan <- false
					return true, nil
				}
			case keys.Enter:
				if p.DefaultValue {
					p.ConfirmStyle.Print(p.ConfirmText)
				} else {
					p.RejectStyle.Print(p.RejectText)
				}
				Println()
				resultChan <- p.DefaultValue
				return true, nil
			case keys.CtrlC:
				cancel()
				interrupted = true
				resultChan <- p.DefaultValue
				return true, nil
			}
			return false, nil
		})
		if !interrupted {
			cursor.StartOfLine()
		}
	}()

	if p.Timeout > 0 {
		select {
		case <-time.After(p.Timeout):
			return p.DefaultValue, nil
		case result := <-resultChan:
			return result, nil
		}
	} else {
		result := <-resultChan
		return result, nil
	}
}

// getShortHandles returns the shorthand answers for the confirmation prompt.
func (p InteractiveConfirmPrinter) getShortHandles() (string, string) {
	y := strings.ToLower(string([]rune(p.ConfirmText)[0]))
	n := strings.ToLower(string([]rune(p.RejectText)[0]))

	return y, n
}

// getSuffix returns the confirmation prompt suffix.
func (p InteractiveConfirmPrinter) getSuffix() string {
	y, n := p.getShortHandles()
	if p.DefaultValue {
		y = strings.ToUpper(y)
	} else {
		n = strings.ToUpper(n)
	}

	return p.SuffixStyle.Sprintf("[%s/%s]", y, n)
}
