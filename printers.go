package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter = PrefixPrinter{}
	_ TextPrinter = CalloutPrinter{}

	_ LivePrinter = (*SpinnerPrinter)(nil)
	_ LivePrinter = (*FlowPrinter)(nil)

	_ InteractivePrinter[bool]     = ConfirmPrinter{}
	_ InteractivePrinter[string]   = SelectPrinter{}
	_ InteractivePrinter[[]string] = MultiselectPrinter{}
	_ InteractivePrinter[string]   = TextInputPrinter{}
	_ InteractivePrinter[string]   = ContinuePrinter{}
)
