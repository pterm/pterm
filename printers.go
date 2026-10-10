package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter = BadgePrinter{}
	_ TextPrinter = BasicTextPrinter{}
	_ TextPrinter = BigTextPrinter{}
	_ TextPrinter = BoxPrinter{}
	_ TextPrinter = CalloutPrinter{}
	_ TextPrinter = CenterPrinter{}
	_ TextPrinter = HeaderPrinter{}
	_ TextPrinter = ParagraphPrinter{}
	_ TextPrinter = PrefixPrinter{}
	_ TextPrinter = RulePrinter{}
	_ TextPrinter = SectionPrinter{}

	_ RenderPrinter = BarChartPrinter{}
	_ RenderPrinter = BigTextPrinter{}
	_ RenderPrinter = BulletListPrinter{}
	_ RenderPrinter = DefinitionsPrinter{}
	_ RenderPrinter = HeatmapPrinter{}
	_ RenderPrinter = KeyHintsPrinter{}
	_ RenderPrinter = KeyValuePrinter{}
	_ RenderPrinter = MeterPrinter{}
	_ RenderPrinter = PanelPrinter{}
	_ RenderPrinter = RulePrinter{}
	_ RenderPrinter = SparklinePrinter{}
	_ RenderPrinter = TablePrinter{}
	_ RenderPrinter = TreePrinter{}

	_ LivePrinter = (*AreaPrinter)(nil)
	_ LivePrinter = (*FlowPrinter)(nil)
	_ LivePrinter = (*MultiPrinter)(nil)
	_ LivePrinter = (*ProgressbarPrinter)(nil)
	_ LivePrinter = (*SpinnerPrinter)(nil)

	_ InteractivePrinter[bool]     = ConfirmPrinter{}
	_ InteractivePrinter[string]   = SelectPrinter{}
	_ InteractivePrinter[[]string] = MultiselectPrinter{}
	_ InteractivePrinter[string]   = TextInputPrinter{}
	_ InteractivePrinter[string]   = ContinuePrinter{}
)
