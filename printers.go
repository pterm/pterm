package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter = PrefixPrinter{}
	_ TextPrinter = CalloutPrinter{}
	_ LivePrinter = (*SpinnerPrinter)(nil)

	_ TextPrinter   = BasicTextPrinter{}
	_ TextPrinter   = HeaderPrinter{}
	_ TextPrinter   = SectionPrinter{}
	_ TextPrinter   = BoxPrinter{}
	_ TextPrinter   = CenterPrinter{}
	_ TextPrinter   = ParagraphPrinter{}
	_ TextPrinter   = BadgePrinter{}
	_ TextPrinter   = RulePrinter{}
	_ RenderPrinter = RulePrinter{}
	_ RenderPrinter = KeyHintsPrinter{}
)
