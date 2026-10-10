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

	_ RenderPrinter = BulletListPrinter{}
	_ RenderPrinter = DefinitionsPrinter{}
	_ RenderPrinter = KeyHintsPrinter{}
	_ RenderPrinter = KeyValuePrinter{}
	_ RenderPrinter = PanelPrinter{}
	_ RenderPrinter = RulePrinter{}
	_ RenderPrinter = TablePrinter{}
	_ RenderPrinter = TreePrinter{}

	_ LivePrinter = (*AreaPrinter)(nil)
	_ LivePrinter = (*FlowPrinter)(nil)
	_ LivePrinter = (*MultiPrinter)(nil)
	_ LivePrinter = (*ProgressbarPrinter)(nil)
	_ LivePrinter = (*SpinnerPrinter)(nil)

	_ InteractivePrinter[bool] = ConfirmPrinter{}
)
