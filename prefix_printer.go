package pterm

var (
	GrayBoxStyle = NewStyle(BgGray, FgLightWhite)
)

var (
	InfoPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "INFO",
			Style: NewStyle(FgLightWhite, BgCyan),
		},
		MessageStyle: NewStyle(FgLightCyan),
	}
	PrintInfo = InfoPrinter.Println

	WarningPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "WARNING",
			Style: NewStyle(FgLightWhite, BgYellow),
		},
		MessageStyle: NewStyle(FgYellow),
	}
	PrintWarning = WarningPrinter.Println

	SuccessPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "SUCCESS",
			Style: NewStyle(FgLightWhite, BgGreen),
		},
		MessageStyle: NewStyle(FgGreen),
	}
	PrintSuccess = SuccessPrinter.Println

	ErrorPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "ERROR",
			Style: NewStyle(FgLightWhite, BgLightRed),
		},
		MessageStyle: NewStyle(FgLightRed),
	}
	PrintError = ErrorPrinter.Println

	DescriptionPrinter = PrefixPrinter{
		Prefix: Prefix{
			Text:  "Description",
			Style: Style{BgDarkGray, FgLightWhite},
		},
		MessageStyle: Style{BgDarkGray, FgLightWhite},
	}
	PrintDescription = DescriptionPrinter.Println
)

// PrefixPrinter is the printer used to print a Prefix
type PrefixPrinter struct {
	Prefix       Prefix
	Scope        Scope
	MessageStyle Style
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

// WithScope adds a scope to the Prefix
func (p PrefixPrinter) WithScope(scope string, style ...Style) *PrefixPrinter {
	p.Scope.Text = scope
	if len(style) > 0 {
		p.Scope.Style = style[0]
	} else {
		p.Scope.Style = p.Prefix.Style
	}
	return &p
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
