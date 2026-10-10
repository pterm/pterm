package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter = PrefixPrinter{}
	_ TextPrinter = CalloutPrinter{}
	_ LivePrinter = (*SpinnerPrinter)(nil)
	_ LivePrinter = (*ProgressbarPrinter)(nil)
	_ LivePrinter = (*AreaPrinter)(nil)
	_ LivePrinter = (*MultiPrinter)(nil)
)
