package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter   = PrefixPrinter{}
	_ TextPrinter   = CalloutPrinter{}
	_ LivePrinter   = (*SpinnerPrinter)(nil)
	_ RenderPrinter = TablePrinter{}
	_ RenderPrinter = TreePrinter{}
	_ RenderPrinter = BulletListPrinter{}
	_ RenderPrinter = PanelPrinter{}
	_ RenderPrinter = KeyValuePrinter{}
	_ RenderPrinter = DefinitionsPrinter{}
)
