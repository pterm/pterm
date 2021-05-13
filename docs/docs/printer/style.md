# StylePrinter

<!-- 
Replace all of the following strings with the current printer.
     style Style StylePrinter DefaultStyle
-->

![StylePrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/style/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/style/main.go" target="_blank">(Show source of demo)</a></p>


## Usage

### Basic usage

```go
pterm.NewStyle(pterm.FgLightCyan, pterm.BgGray, pterm.Bold).Println("Hello, World!")
```

### Functions

|Function|Description|
|--------|-----------|
|[NewStyle](https://pkg.go.dev/github.com/pterm/pterm#Style.NewStyle)|NewStyle returns a new Style. Accepts multiple colors.|
|[Add](https://pkg.go.dev/github.com/pterm/pterm#Style.Add)|Add styles to the current Style.|
|[Code](https://pkg.go.dev/github.com/pterm/pterm#Style.Code)|Code convert to code string. returns like "32;45;3".|
|[String](https://pkg.go.dev/github.com/pterm/pterm#Style.String)|String convert to code string. returns like "32;45;3".|

### Output functions

> This printer implements the interface [`TextPrinter`](https://github.com/pterm/pterm/blob/master/interface_text_printer.go)

|Function|Description|
|------|---------|
|Sprint(a ...interface{})|Returns a string|
|Sprintln(a ...interface{})|Returns a string with a new line at the end|
|Sprintf(format string, a ...interface{})|Returns a string, formatted according to a format specifier|
|Sprintfln(format string, a ...interface{})|Returns a string, formatted according to a format specifier with a new line at the end|
|Print(a ...interface{})|Prints to the terminal|
|Println(a ...interface{})|Prints to the terminal with a new line at the end|
|Printf(format string, a ...interface{})|Prints to the terminal, formatted according to a format specifier|
|Printfln(format string, a ...interface{})|Prints to the terminal, formatted according to a format specifier with a new line at the end|

## Related
- [Override default printers](docs/customizing/override-default-printer.md)