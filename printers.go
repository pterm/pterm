package pterm

// Compile-time checks that every printer implements its family.
var (
	_ TextPrinter = PrefixPrinter{}
	_ TextPrinter = CalloutPrinter{}
	_ LivePrinter = (*SpinnerPrinter)(nil)

	_ TextPrinter   = BigTextPrinter{}
	_ RenderPrinter = BigTextPrinter{}
	_ RenderPrinter = BarChartPrinter{}
	_ RenderPrinter = HeatmapPrinter{}
	_ RenderPrinter = MeterPrinter{}
	_ RenderPrinter = SparklinePrinter{}
)
