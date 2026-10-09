package pterm

import "fmt"

// This file documents and enforces pterm's printer patterns.
//
// Every printer belongs to exactly one of the following families and must
// implement the family's interface (enforced by the compile-time assertions
// below):
//
//   - TextPrinter: prints formatted text directly (Sprint*/Print* methods).
//     Examples: BasicTextPrinter, PrefixPrinter, HeaderPrinter, Color, RGB.
//
//   - RenderPrinter: renders complex, multi-line content and exposes it via
//     Render() and Srender(). Examples: TablePrinter, TreePrinter,
//     BarChartPrinter.
//
//   - LivePrinter: printers whose output updates in place over time,
//     started with Start (which returns the started instance) and terminated
//     with Stop. Examples: SpinnerPrinter, ProgressbarPrinter, AreaPrinter,
//     MultiPrinter.
//
//   - Interactive printers: read user input and return the result from
//     Show(). Examples: InteractiveConfirmPrinter, InteractiveSelectPrinter.
//     Their return types differ, so there is no shared interface; they follow
//     the common pattern of a Default* value, chainable With* options and a
//     Show method.
//
// All printers follow the builder pattern: With* methods take a value
// receiver, modify the copy, and return a pointer to it, so configuring a
// printer never mutates the printer it was derived from.
var (
	_ TextPrinter = &BasicTextPrinter{}
	_ TextPrinter = BoxPrinter{}
	_ TextPrinter = CenterPrinter{}
	_ TextPrinter = &HeaderPrinter{}
	_ TextPrinter = &ParagraphPrinter{}
	_ TextPrinter = &PrefixPrinter{}
	_ TextPrinter = &SectionPrinter{}
	_ TextPrinter = FgDefault
	_ TextPrinter = RGB{}
	_ TextPrinter = RGBStyle{}

	_ RenderPrinter = BarChartPrinter{}
	_ RenderPrinter = BigTextPrinter{}
	_ RenderPrinter = BulletListPrinter{}
	_ RenderPrinter = HeatmapPrinter{}
	_ RenderPrinter = PanelPrinter{}
	_ RenderPrinter = TablePrinter{}
	_ RenderPrinter = TreePrinter{}

	_ LivePrinter = &AreaPrinter{}
	_ LivePrinter = &MultiPrinter{}
	_ LivePrinter = &ProgressbarPrinter{}
	_ LivePrinter = &SpinnerPrinter{}
)

// printOnError prints every argument that is a non-nil error via p.
func printOnError(p TextPrinter, a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok && err != nil {
			p.Println(err)
		}
	}
}

// printOnErrorf wraps every argument that is a non-nil error with format and
// prints it via p.
func printOnErrorf(p TextPrinter, format string, a ...any) {
	for _, arg := range a {
		if err, ok := arg.(error); ok && err != nil {
			p.Println(fmt.Errorf(format, err))
		}
	}
}
