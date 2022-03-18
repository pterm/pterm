# BarChartPrinter

<!--
Replace all of the following strings with the current printer.
     barchart BarChart BarChartPrinter DefaultBarChart
-->

![BarChartPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/barchart/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/barchart/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
bars := pterm.Bars{
		pterm.Bar{
			Label: "Bar 1",
			Value: 5,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Bar 2",
			Value: 3,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
		pterm.Bar{
			Label: "Longer Label",
			Value: 7,
			Style: pterm.NewStyle(pterm.FgCyan),
		},
	}

	_ = pterm.DefaultBarChart.WithBars(bars).Render()
```

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultBarChart.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultBarChart.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                     | Type                                                   |
| ---------------------------------------------------------------------------------------------------------- | ------------------------------------------------------ |
| [Bars](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.Bars)                                     | [Bars](https://pkg.go.dev/github.com/pterm/pterm#Bars) |
| [Horizontal](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.Horizontal)                         | bool                                                   |
| [ShowValue](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.ShowValue)                           | bool                                                   |
| [Height](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.Height)                                 | int                                                    |
| [Width](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.Width)                                   | int                                                    |
| [VerticalBarCharacter](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.VerticalBarCharacter)     | string                                                 |
| [HorizontalBarCharacter](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.HorizontalBarCharacter) | string                                                 |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#BarChartPrinter.WithWriter)                             | io.Writer                                              |

### Output functions

<!-- Remove comment of the correct interface -->

<!--
> This printer implements the interface [`TextPrinter`](https://github.com/pterm/pterm/blob/master/interface_text_printer.go)

|Function|Description|
|------|---------|
|Sprint(a ...interface{})|Returns a string|
|Sprintln(a ...interface{})|Returns a string with a new line at the end|
|Sprintf(format string, a ...interface{})|Returns a string, formatted according to a format specifier|
|Print(a ...interface{})|Prints to the terminal|
|Println(a ...interface{})|Prints to the terminal with a new line at the end|
|Printf(format string, a ...interface{})|Prints to the terminal, formatted according to a format specifier|
-->

> This printer implements the interface [`RenderablePrinter`](https://github.com/pterm/pterm/blob/master/interface_renderable_printer.go)

| Function  | Description        |
| --------- | ------------------ |
| Render()  | Prints to Terminal |
| Srender() | Returns a string   |

<!--
> This printer implements the interface [`LivePrinter`](https://github.com/pterm/pterm/blob/master/interface_live_printer.go)

|Function|Description|
|------|---------|
|Start()|Returns itself and possible errors|
|Stop()|Returns itself and possible errors|
|GenericStart()|Returns the started LivePrinter and possible errors|
|GenericStop()|Returns the stopped LivePrinter and possible errors|

> [!NOTE]
> The generic start and stop methods are only used to implement the printer into the interface.
> Use the normal `Start()` and `Stop()` methods if possible.
-->

## Related

- [Override default printers](docs/customizing/override-default-printer.md)
