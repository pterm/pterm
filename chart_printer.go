package pterm

// ChartBar is a bar which can be used to display a HorizontalBarChartPrinter or VerticalBarChartPrinter.
type ChartBar struct {
	Label string
	Value int
}

// HorizontalBarChartPrinter can print errors of the default Error type.
type HorizontalBarChartPrinter struct {
	Bars []ChartBar
}

// Sprint formats using the default formats for its operands and returns the resulting string.
// Spaces are added between operands when neither is a string.
func (p HorizontalBarChartPrinter) Sprint(a ...interface{}) string {
	panic("not implemented")
}

// Sprintln formats using the default formats for its operands and returns the resulting string.
// Spaces are always added between operands and a newline is appended.
func (p HorizontalBarChartPrinter) Sprintln(a ...interface{}) string {
	return Sprintln(p.Sprint(a...))
}

// Sprintf formats according to a format specifier and returns the resulting string.
func (p HorizontalBarChartPrinter) Sprintf(format string, a ...interface{}) string {
	return p.Sprint(Sprintf(format, a...))
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
func (p HorizontalBarChartPrinter) Print(a ...interface{}) GenericPrinter {
	Print(p.Sprint(a...))
	return &p
}

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func (p HorizontalBarChartPrinter) Println(a ...interface{}) GenericPrinter {
	Println(p.Sprint(a...))
	return &p
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
func (p HorizontalBarChartPrinter) Printf(format string, a ...interface{}) GenericPrinter {
	Print(p.Sprintf(format, a...))
	return &p
}
