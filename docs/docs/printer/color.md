# ColorPrinter

<!-- 
Replace all of the following strings with the current printer.
     color Color ColorPrinter DefaultColor
-->

![ColorPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-with-color/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/print-with-color/main.go" target="_blank">(Show source of demo)</a></p>


## Usage

### Basic usage

```go
pterm.FgBlack.Println("FgBlack")
pterm.FgRed.Println("FgRed")
pterm.FgGreen.Println("FgGreen")
pterm.FgYellow.Println("FgYellow")
pterm.FgBlue.Println("FgBlue")
pterm.FgMagenta.Println("FgMagenta")
pterm.FgCyan.Println("FgCyan")
pterm.FgWhite.Println("FgWhite")
pterm.FgLightRed.Println("FgLightRed")
pterm.FgLightGreen.Println("FgLightGreen")
pterm.FgLightYellow.Println("FgLightYellow")
pterm.FgLightBlue.Println("FgLightBlue")
pterm.FgLightMagenta.Println("FgLightMagenta")
pterm.FgLightCyan.Println("FgLightCyan")
pterm.FgLightWhite.Println("FgLightWhite")

// For quick usage in strings, you can also use the color names as functions:
pterm.Cyan("Cyan")
pterm.LightMagenta("LightMagenta")
// etc...
```
<!-- Delete this section if the printer does not expose functions other than the default output functions -->
### Functions

|Function|Description|
|--------|-----------|
|[String](https://pkg.go.dev/github.com/pterm/pterm#Color.String)|String converts the color to a string.|

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
