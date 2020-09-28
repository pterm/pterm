package pterm

var (
	// GrayBoxStyle wraps text in a gray box
	GrayBoxStyle = NewStyle(BgGray, FgLightWhite)
)

var (
	// Info returns a PrefixPrinter, which can be used to print text with an "info" Prefix
	Info = PrefixPrinter{
		Prefix: Prefix{
			Text:  "INFO",
			Style: NewStyle(FgLightWhite, BgCyan),
		},
		MessageStyle: NewStyle(FgLightCyan),
	}
	// PrintInfo is a shortcut to Info.Println
	PrintInfo = Info.Println

	// Warning returns a PrefixPrinter, which can be used to print text with a "warning" Prefix
	Warning = PrefixPrinter{
		Prefix: Prefix{
			Text:  "WARNING",
			Style: NewStyle(FgLightWhite, BgYellow),
		},
		MessageStyle: NewStyle(FgYellow),
	}
	// PrintWarning is a shortcut to Warning.Println
	PrintWarning = Warning.Println

	// Success returns a PrefixPrinter, which can be used to print text with a "success" Prefix
	Success = PrefixPrinter{
		Prefix: Prefix{
			Text:  "SUCCESS",
			Style: NewStyle(FgLightWhite, BgGreen),
		},
		MessageStyle: NewStyle(FgGreen),
	}
	// PrintSuccess is a shortcut to Success.Println
	PrintSuccess = Success.Println

	// Error returns a PrefixPrinter, which can be used to print text with an "error" Prefix
	Error = PrefixPrinter{
		Prefix: Prefix{
			Text:  "ERROR",
			Style: NewStyle(FgLightWhite, BgLightRed),
		},
		MessageStyle: NewStyle(FgLightRed),
	}
	// PrintError is a shortcut to Error.Println
	PrintError = Error.Println

	// Description returns a PrefixPrinter, which can be used to print text with a "description" Prefix
	Description = PrefixPrinter{
		Prefix: Prefix{
			Text:  "Description",
			Style: Style{BgDarkGray, FgLightWhite},
		},
		MessageStyle: Style{BgDarkGray, FgLightWhite},
	}
	// PrintDescription is a shortcut to Description.Println
	PrintDescription = Description.Println
)

// PrefixPrinter is the printer used to print a Prefix
type PrefixPrinter struct {
	Prefix       Prefix
	Scope        Scope
	MessageStyle Style
}

// WithPrefix adds a custom prefix to the printer
func (p PrefixPrinter) WithPrefix(prefix Prefix) *PrefixPrinter {
	p.Prefix = prefix
	return &p
}

// WithScope adds a scope to the Prefix
func (p PrefixPrinter) WithScope(scope string, colors ...Color) *PrefixPrinter {
	p.Scope.Text = scope
	p.Scope.Style = colors
	return &p
}

// WithMessageStyle adds a custom prefix to the printer
func (p PrefixPrinter) WithMessageStyle(prefix Prefix) *PrefixPrinter {
	p.Prefix = prefix
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p PrefixPrinter) Sprint(a ...interface{}) string {
	var args []interface{}
	args = append(args, p.GetFormattedPrefix())
	if p.Scope.Text != "" {
		args = append(args, NewStyle(p.Scope.Style...).Sprint(" ("+p.Scope.Text+") "))
	}
	args = append(args, p.GetFormattedMessage(a...))

	return Sprint(args...)
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p PrefixPrinter) Sprintln(a ...interface{}) string {
	return p.Sprint(a...) + "\n"
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p PrefixPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p PrefixPrinter) Print(a ...interface{}) GenericPrinter {
	Print(p.Sprint(a...))
	return p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p PrefixPrinter) Println(a ...interface{}) GenericPrinter {
	Print(p.Sprintln(a...))
	return p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p PrefixPrinter) Printf(format string, a ...interface{}) GenericPrinter {
	Print(Sprintf(format, a...))
	return p
}

// GetFormattedPrefix returns the Prefix as a styled text string
func (p PrefixPrinter) GetFormattedPrefix() string {
	return NewStyle(p.Prefix.Style...).Sprint(" " + p.Prefix.Text + " ")
}

// GetFormattedMessage returns the message as a styled text string
func (p PrefixPrinter) GetFormattedMessage(a ...interface{}) string {
	var args []interface{}
	args = append(args, " ")
	args = append(args, a...)
	if p.MessageStyle == nil {
		p.MessageStyle = NewStyle()
	}
	return NewStyle(p.MessageStyle...).Sprint(args...)
}

// Prefix contains the data used as the beginning of a printed text via a PrefixPrinter
type Prefix struct {
	Text  string
	Style Style
}

// Scope contains the data of the optional scope of a prefix.
// If it has a text, it will be printed after the Prefix in brackets.
type Scope struct {
	Text  string
	Style Style
}
