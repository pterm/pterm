package pterm

var (
	// DefaultHeaderPrinter returns the printer for a default header text.
	// Defaults to LightWhite, Bold Text and a Gray Header background.
	DefaultHeaderPrinter = HeaderPrinter{Header: Header{
		TextStyle:       Style{FgLightWhite, Bold},
		BackgroundStyle: Style{BgGray},
		Margin:          5,
	}}

	// PrintHeader is the short form of DefaultHeaderPrinter.Println
	PrintHeader = DefaultHeaderPrinter.Println
)

type Header struct {
	TextStyle       Style
	BackgroundStyle Style
	Margin          int
}

type HeaderPrinter struct {
	Header Header
}

func (p HeaderPrinter) Sprint(a ...interface{}) string {
	text := Sprint(a...)
	textLength := len(text) + p.Header.Margin*2
	var marginString string
	for i := 0; i < p.Header.Margin; i++ {
		marginString += " "
	}
	var blankLine string
	for i := 0; i < textLength; i++ {
		blankLine += " "
	}

	var ret string

	ret += p.Header.BackgroundStyle.Sprint(blankLine) + "\n"
	ret += p.Header.BackgroundStyle.Sprint(p.Header.TextStyle.Sprint(marginString+text+marginString)) + "\n"
	ret += p.Header.BackgroundStyle.Sprint(blankLine) + "\n\n"

	return ret
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
