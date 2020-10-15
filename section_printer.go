package pterm

// DefaultSection is the default section printer.
var DefaultSection = SectionPrinter{
	Style:         ThemeDefault.SectionStyle,
	Level:         1,
	TopPadding:    1,
	BottomPadding: 1,
}

// SectionPrinter prints a new section title.
// It can be used to structure longer text, or different chapters of your program.
type SectionPrinter struct {
	Style         Style
	Level         int
	TopPadding    int
	BottomPadding int
}

// WithStyle returns a new SectionPrinter with a specific style.
func (p SectionPrinter) WithStyle(style Style) *SectionPrinter {
	p.Style = style
	return &p
}

// WithLevel returns a new SectionPrinter with a specific level.
func (p SectionPrinter) WithLevel(level int) *SectionPrinter {
	p.Level = level
	return &p
}

// WithTopPadding returns a new SectionPrinter with a specific top padding.
func (p SectionPrinter) WithTopPadding(level int) *SectionPrinter {
	p.TopPadding = level
	return &p
}

// WithBottomPadding returns a new SectionPrinter with a specific top padding.
func (p SectionPrinter) WithBottomPadding(level int) *SectionPrinter {
	p.BottomPadding = level
	return &p
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p SectionPrinter) Sprint(a ...interface{}) string {
	var ret string

	for i := 0; i < p.TopPadding; i++ {
		ret += "\n"
	}

	ret += p.Style.Sprint(a...)

	for i := 0; i < p.BottomPadding; i++ {
		ret += "\n"
	}

	return ret
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p SectionPrinter) Sprintln(a ...interface{}) string {
	return Sprintln(p.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p SectionPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Print(a ...interface{}) *TextPrinter {
	Print(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Println(a ...interface{}) *TextPrinter {
	Println(p.Sprint(a...))
	tp := TextPrinter(p)
	return &tp
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p *SectionPrinter) Printf(format string, a ...interface{}) *TextPrinter {
	Print(p.Sprintf(format, a...))
	tp := TextPrinter(p)
	return &tp
}
