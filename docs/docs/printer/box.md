# BoxPrinter

<!--
Replace all of the following strings with the current printer.
     box Box BoxPrinter DefaultBox
-->

![BoxPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/box/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/box/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultBox.Println("test")
```

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultBox.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultBox.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                      | Type                                                       |
| ----------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| [BottomLeftCornerString](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithBottomLeftCornerString)   | string                                                     |
| [BottomPadding](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithBottomPadding)                     | int                                                        |
| [BottomRightCornerString](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithBottomRightCornerString) | string                                                     |
| [BoxStyle](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithBoxStyle)                               | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [HorizontalString](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithHorizontalString)               | string                                                     |
| [LeftPadding](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithLeftPadding)                         | int                                                        |
| [RightPadding](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithRightPadding)                       | int                                                        |
| [Text](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithText)                                       | string                                                     |
| [TextStyle](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTextStyle)                             | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [Title](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitle)                                     | string                                                     |
| [TitleBottomCenter](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitleBottomCenter)             | boolean                                                    |
| [TitleBottomLeft](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitleBottomLeft)                 | boolean                                                    |
| [TitleBottomRight](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitleBottomRight)               | boolean                                                    |
| [TitleTopCenter](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitleTopCenter)                   | boolean                                                    |
| [TitleTopLeft](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitleTopLeft)                       | boolean                                                    |
| [TitleTopRight](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTitleTopRight)                     | boolean                                                    |
| [TopLeftCornerString](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTopLeftCornerString)         | string                                                     |
| [TopPadding](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTopPadding)                           | int                                                        |
| [TopRightCornerString](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithTopRightCornerString)       | string                                                     |
| [VerticalString](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithVerticalString)                   | string                                                     |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#BoxPrinter.WithWriter)                                   | io.Writer                                                  |

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
