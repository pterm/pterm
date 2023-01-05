package pterm

import (
	"fmt"
	"io"
	"runtime"
	"strings"

	"github.com/pterm/pterm/internal"
)

var (
	// GrayBoxStyle wraps text in a gray box.
	GrayBoxStyle = NewStyle(BgGray, FgLightWhite)
)

var (
	// Info returns a PrefixPrinter, which can be used to print text with an "info" Prefix.
	Info = PrefixPrinter{
		MessageStyle: &ThemeDefault.InfoMessageStyle,
		Prefix: Prefix{
			Style: &ThemeDefault.InfoPrefixStyle,
			Text:  "INFO",
		},
	}

	// Warning returns a PrefixPrinter, which can be used to print text with a "warning" Prefix.
	Warning = PrefixPrinter{
		MessageStyle: &ThemeDefault.WarningMessageStyle,
		Prefix: Prefix{
			Style: &ThemeDefault.WarningPrefixStyle,
			Text:  "WARNING",
		},
	}

	// Success returns a PrefixPrinter, which can be used to print text with a "success" Prefix.
	Success = PrefixPrinter{
		MessageStyle: &ThemeDefault.SuccessMessageStyle,
		Prefix: Prefix{
			Style: &ThemeDefault.SuccessPrefixStyle,
			Text:  "SUCCESS",
		},
	}

	// Error returns a PrefixPrinter, which can be used to print text with an "error" Prefix.
	Error = PrefixPrinter{
		MessageStyle: &ThemeDefault.ErrorMessageStyle,
		Prefix: Prefix{
			Style: &ThemeDefault.ErrorPrefixStyle,
			Text:  " ERROR ",
		},
	}

	// Fatal returns a PrefixPrinter, which can be used to print text with an "fatal" Prefix.
	// NOTICE: Fatal terminates the application immediately!
	Fatal = PrefixPrinter{
		MessageStyle: &ThemeDefault.FatalMessageStyle,
		Prefix: Prefix{
			Style: &ThemeDefault.FatalPrefixStyle,
			Text:  " FATAL ",
		},
		Fatal: true,
	}

	// Debug Prints debug messages. By default it will only print if PrintDebugMessages is true.
	// You can change PrintDebugMessages with EnableDebugMessages and DisableDebugMessages, or by setting the variable itself.
	Debug = PrefixPrinter{
		MessageStyle: &ThemeDefault.DebugMessageStyle,
		Prefix: Prefix{
			Text:  " DEBUG ",
			Style: &ThemeDefault.DebugPrefixStyle,
		},
		Debugger: true,
	}

	// Description returns a PrefixPrinter, which can be used to print text with a "description" Prefix.
	Description = PrefixPrinter{
		MessageStyle: &ThemeDefault.DescriptionMessageStyle,
		Prefix: Prefix{
			Style: &ThemeDefault.DescriptionPrefixStyle,
			Text:  "Description",
		},
	}
)

// PrefixPrinter is the printer used to print a Prefix.
type PrefixPrinter struct {
	Prefix           Prefix
	Scope            Scope
	MessageStyle     *Style
	Fatal            bool
	ShowLineNumber   bool
	LineNumberOffset int
	Writer           io.Writer
	// If Debugger is true, the printer will only print if PrintDebugMessages is set to true.
	// You can change PrintDebugMessages with EnableDebugMessages and DisableDebugMessages, or by setting the variable itself.
	Debugger bool
}

// WithPrefix adds a custom prefix to the printer.
func (p PrefixPrinter) WithPrefix(prefix Prefix) *PrefixPrinter {
	p.Prefix = prefix
	return &p
}

// WithScope adds a scope to the Prefix.
func (p PrefixPrinter) WithScope(scope Scope) *PrefixPrinter {
	p.Scope = scope
	return &p
}

// WithMessageStyle adds a custom prefix to the printer.
func (p PrefixPrinter) WithMessageStyle(style *Style) *PrefixPrinter {
	p.MessageStyle = style
	return &p
}

// WithFatal sets if the printer should panic after printing.
// NOTE:
// The printer will only panic if either PrefixPrinter.Println, PrefixPrinter.Print
// or PrefixPrinter.Printf is called.
func (p PrefixPrinter) WithFatal(b ...bool) *PrefixPrinter {
	p.Fatal = internal.WithBoolean(b)
	return &p
}

// WithShowLineNumber sets if the printer should print the line number from where it's called in a go file.
func (p PrefixPrinter) WithShowLineNumber(b ...bool) *PrefixPrinter {
	p.ShowLineNumber = internal.WithBoolean(b)
	return &p
}

// WithDebugger returns a new Printer with specific Debugger value.
// If Debugger is true, the printer will only print if PrintDebugMessages is set to true.
// You can change PrintDebugMessages with EnableDebugMessages and DisableDebugMessages, or by setting the variable itself.
func (p PrefixPrinter) WithDebugger(b ...bool) *PrefixPrinter {
	p.Debugger = internal.WithBoolean(b)
	return &p
}

// WithLineNumberOffset can be used to exclude a specific amount of calls in the call stack.
// If you make a wrapper function for example, you can set this to one.
// The printed line number will then be the line number where your wrapper function is called.
func (p PrefixPrinter) WithLineNumberOffset(offset int) *PrefixPrinter {
	p.LineNumberOffset = offset
	return &p
}

// WithWriter sets the custom Writer.
func (p PrefixPrinter) WithWriter(writer io.Writer) *PrefixPrinter {
	p.Writer = writer
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p *PrefixPrinter) Sprint(a ...interface{}) string {
	m := Sprint(a...)
	if p.Debugger && !PrintDebugMessages {
		return ""
	}

	if RawOutput {
		if p.Prefix.Text != "" {
			return Sprintf("%s: %s", strings.TrimSpace(p.Prefix.Text), Sprint(a...))
		} else {
			return Sprint(a...)
		}
	}

	if p.Prefix.Style == nil {
		p.Prefix.Style = NewStyle()
	}
	if p.Scope.Style == nil {
		p.Scope.Style = NewStyle()
	}
	if p.MessageStyle == nil {
		p.MessageStyle = NewStyle()
	}

	var ret string
	var newLine bool

	if strings.HasSuffix(m, "\n") {
		m = strings.TrimRight(m, "\n")
		newLine = true
	}

	messageLines := strings.Split(m, "\n")
	for i, m := range messageLines {
		if i == 0 {
			ret += p.GetFormattedPrefix() + " "
			if p.Scope.Text != "" {
				ret += NewStyle(*p.Scope.Style...).Sprint(" (" + p.Scope.Text + ") ")
			}
			ret += p.MessageStyle.Sprint(m)
		} else {
			ret += "\n" + p.Prefix.Style.Sprint(strings.Repeat(" ", len(p.Prefix.Text)+2)) + " " + p.MessageStyle.Sprint(m)
		}
	}

	if p.ShowLineNumber {
		_, fileName, line, _ := runtime.Caller(3 + p.LineNumberOffset)
		ret += FgGray.Sprint("\n└ " + fmt.Sprintf("(%s:%d)\n", fileName, line))
		newLine = false
	}

	if newLine {
		ret += "\n"
	}

	return Sprint(ret)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p PrefixPrinter) Sprintln(a ...interface{}) string {
	if p.Debugger && !PrintDebugMessages {
		return ""
	}
	str := fmt.Sprintln(a...)
	return p.Sprint(str)
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p PrefixPrinter) Sprintf(format string, a ...interface{}) string {
	if p.Debugger && !PrintDebugMessages {
		return ""
	}
	return p.Sprint(Sprintf(format, a...))
}

// Sprintfln formats according to a format specifier and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p PrefixPrinter) Sprintfln(format string, a ...interface{}) string {
	if p.Debugger && !PrintDebugMessages {
		return ""
	}
	return p.Sprintf(format, a...) + "\n"
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *PrefixPrinter) Print(a ...interface{}) *TextPrinter {
	tp := TextPrinter(p)
	if p.Debugger && !PrintDebugMessages {
		return &tp
	}
	p.LineNumberOffset--
	Fprint(p.Writer, p.Sprint(a...))
	p.LineNumberOffset++
	checkFatal(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *PrefixPrinter) Println(a ...interface{}) *TextPrinter {
	tp := TextPrinter(p)
	if p.Debugger && !PrintDebugMessages {
		return &tp
	}
	Fprint(p.Writer, p.Sprintln(a...))
	checkFatal(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *PrefixPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	tp := TextPrinter(p)
	if p.Debugger && !PrintDebugMessages {
		return &tp
	}
	Fprint(p.Writer, p.Sprintf(format, a...))
	checkFatal(p)
	return &tp
}

// Printfln formats according to a format specifier and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *PrefixPrinter) Printfln(format string, a ...interface{}) *TextPrinter {
	tp := TextPrinter(p)
	if p.Debugger && !PrintDebugMessages {
		return &tp
	}
	p.LineNumberOffset++
	Fprint(p.Writer, p.Sprintfln(format, a...))
	p.LineNumberOffset--
	checkFatal(p)
	return &tp
}

// PrintOnError prints every error which is not nil.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
//
// Note: Use WithFatal(true) or Fatal to panic after first non nil error.
func (p *PrefixPrinter) PrintOnError(a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				p.Println(err)
			}
		}
	}

	tp := TextPrinter(p)
	return &tp
}

// PrintOnErrorf wraps every error which is not nil and prints it.
// If every error is nil, nothing will be printed.
// This can be used for simple error checking.
func (p *PrefixPrinter) PrintOnErrorf(format string, a ...interface{}) *TextPrinter {
	for _, arg := range a {
		if err, ok := arg.(error); ok {
			if err != nil {
				p.Println(fmt.Errorf(format, err))
			}
		}
	}

	tp := TextPrinter(p)
	return &tp
}

// GetFormattedPrefix returns the Prefix as a styled text string.
func (p PrefixPrinter) GetFormattedPrefix() string {
	return p.Prefix.Style.Sprint(" " + p.Prefix.Text + " ")
}

// Prefix contains the data used as the beginning of a printed text via a PrefixPrinter.
type Prefix struct {
	Text  string
	Style *Style
}

// Scope contains the data of the optional scope of a prefix.
// If it has a text, it will be printed after the Prefix in brackets.
type Scope struct {
	Text  string
	Style *Style
}

func checkFatal(p *PrefixPrinter) {
	if p.Fatal {
		panic("")
	}
}
