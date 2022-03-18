# CenterPrinter

<!--
Replace all of the following strings with the current printer.
        center Center CenterPrinter DefaultCenter
-->

![CenterPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/center/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/center/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultCenter.Println("Hello,\nWorld!")
```

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultCenter.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultCenter.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                           | Type      |
| ---------------------------------------------------------------------------------------------------------------- | --------- |
| [CenterEachLineSeparately](https://pkg.go.dev/github.com/pterm/pterm#CenterPrinter.WithCenterEachLineSeparately) | ...bool   |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#CenterPrinter.WithWriter)                                     | io.Writer |

### Output functions

> This printer implements the interface [`TextPrinter`](https://github.com/pterm/pterm/blob/master/interface_text_printer.go)

| Function                                   | Description                                                                                  |
| ------------------------------------------ | -------------------------------------------------------------------------------------------- |
| Sprint(a ...interface{})                   | Returns a string                                                                             |
| Sprintln(a ...interface{})                 | Returns a string with a new line at the end                                                  |
| Sprintf(format string, a ...interface{})   | Returns a string, formatted according to a format specifier                                  |
| Sprintfln(format string, a ...interface{}) | Returns a string, formatted according to a format specifier with a new line at the end       |
| Print(a ...interface{})                    | Prints to the terminal                                                                       |
| Println(a ...interface{})                  | Prints to the terminal with a new line at the end                                            |
| Printf(format string, a ...interface{})    | Prints to the terminal, formatted according to a format specifier                            |
| Printfln(format string, a ...interface{})  | Prints to the terminal, formatted according to a format specifier with a new line at the end |

## Related

- [Override default printers](docs/customizing/override-default-printer.md)
