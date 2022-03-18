# BigTextPrinter

<!--
Replace all of the following strings with the current printer.
        bigtext BigText BigTextPrinter DefaultBigText
-->

![BigTextPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bigtext/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/bigtext/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("Hello")).Render()
```

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultBigText.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultBigText.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                      | Type                                                         |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------ |
| [Letters](https://pkg.go.dev/github.com/pterm/pterm#BigTextPrinter.WithLetters)             | [Letters](https://pkg.go.dev/github.com/pterm/pterm#Letters) |
| [BigCharacters](https://pkg.go.dev/github.com/pterm/pterm#BigTextPrinter.WithBigCharacters) | map[string]string                                            |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#BigTextPrinter.WithWriter)               | io.Writer                                                    |

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
