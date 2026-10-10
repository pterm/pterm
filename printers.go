package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter = BadgePrinter{}
	_ TextPrinter = BasicTextPrinter{}
	_ TextPrinter = BoxPrinter{}
	_ TextPrinter = CalloutPrinter{}
	_ TextPrinter = CenterPrinter{}
	_ TextPrinter = HeaderPrinter{}
	_ TextPrinter = ParagraphPrinter{}
	_ TextPrinter = PrefixPrinter{}
	_ TextPrinter = RulePrinter{}
	_ TextPrinter = SectionPrinter{}

	_ RenderPrinter = KeyHintsPrinter{}
	_ RenderPrinter = RulePrinter{}

	_ LivePrinter = (*SpinnerPrinter)(nil)
	_ LivePrinter = (*FlowPrinter)(nil)

	_ InteractivePrinter[bool] = ConfirmPrinter{}
)
