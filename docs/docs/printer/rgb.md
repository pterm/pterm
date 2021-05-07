# RGBPrinter

<!-- 
Replace all of the following strings with the current printer.
     print-color-rgb RGB RGBPrinter DefaultRGB
-->

![RGBPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/print-color-rgb/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/print-color-rgb/main.go" target="_blank">(Show source of demo)</a></p>


## Usage

### Basic usage

```go
pterm.NewRGB(178, 44, 199).Println("This text is printed with a custom RGB!")
```

### Functions

|Function|Description|
|--------|-----------|
|[NewRGB](https://pkg.go.dev/github.com/pterm/pterm#RGB.NewRGB)|NewRGB returns a new RGB.|
|[NewRGBFromHEX](https://pkg.go.dev/github.com/pterm/pterm#RGB.NewRGBFromHEX)|NewRGBFromHEX converts a HEX and returns a new RGB.|
|[Fade](https://pkg.go.dev/github.com/pterm/pterm#RGB.Fade)|Fade fades one RGB value (over other RGB values) to another RGB value, by giving the function a minimum, maximum and current value.|
|[GetValues](https://pkg.go.dev/github.com/pterm/pterm#RGB.GetValues)|GetValues returns the RGB values separately.|

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