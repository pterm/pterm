package pterm

var (
	// DefaultHeaderPrinter returns the printer for a default header text.
	// Defaults to LightWhite, Bold Text and a Gray Header background.
	DefaultHeaderPrinter = HeaderPrinter{Header: Header{
		TextStyle:       Style{FgLightWhite, Bold},
		BackgroundStyle: Style{BgGray},
	}}

	// PrintHeader is the short form of DefaultHeaderPrinter.Println
	PrintHeader = DefaultHeaderPrinter.Println
)

type Header struct {
	TextStyle       Style
	BackgroundStyle Style
}

type HeaderPrinter struct {
	Header Header
}

func (p HeaderPrinter) Sprint(a ...interface{}) string {
	return p.Header.BackgroundStyle.Sprint("\n", "    "+p.Header.TextStyle.Sprint(Sprint(a...)), p.Header.BackgroundStyle.Sprint("\n\n"))
}

func (p HeaderPrinter) Sprintln(a ...interface{}) string {
	return Sprint(p.Sprint(a...) + "\n")
}

func (p HeaderPrinter) Sprintf(format string, a ...interface{}) string {
	panic("implement me")
}

func (p HeaderPrinter) Print(a ...interface{}) GenericPrinter {
	Print(p.Sprint(a...))
	return p
}

func (p HeaderPrinter) Println(a ...interface{}) GenericPrinter {
	Println(p.Sprint(a...))
	return p
}

func (p HeaderPrinter) Printf(format string, a ...interface{}) GenericPrinter {
	p.Print(Sprintf(format, a...))
	return p
}
