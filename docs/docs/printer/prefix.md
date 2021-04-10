# PrefixPrinter

<!-- 
Replace all of the following strings with the current printer.
     prefix Prefix PrefixPrinter DefaultPrefix
-->

![PrefixPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/prefix/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/prefix/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.Debug.Println("Hello, World!") // Print Debug.
pterm.Info.Println("Hello, World!") // Print Info.
pterm.Success.Println("Hello, World!") // Print Success.
pterm.Warning.Println("Hello, World!") // Print Warning.
pterm.Error.Println("Hello, World!") // Print Error.
pterm.Fatal.Println("Hello, World!") // Print Fatal.
```

### Options

> To make a copy with modified options you can use:
> ```go
pterm.Debug.Println("Hello, World!").WithOptionName(option)
pterm.Info.Println("Hello, World!").WithOptionName(option)
pterm.Success.Println("Hello, World!").WithOptionName(option)
pterm.Warning.Println("Hello, World!").WithOptionName(option)
pterm.Error.Println("Hello, World!").WithOptionName(option)
pterm.Fatal.Println("Hello, World!").WithOptionName(option)
> ```
>
> To change multiple options at once, you can chain the functions:
> ```go
pterm.Debug.Println("Hello, World!").WithOptionName(option).WithOptionName2(option2)...
pterm.Info.Println("Hello, World!").WithOptionName(option).WithOptionName2(option2)...
pterm.Success.Println("Hello, World!").WithOptionName(option).WithOptionName2(option2)...
pterm.Warning.Println("Hello, World!").WithOptionName(option).WithOptionName2(option2)...
pterm.Error.Println("Hello, World!").WithOptionName(option).WithOptionName2(option2)...
pterm.Fatal.Println("Hello, World!").WithOptionName(option).WithOptionName2(option2)...
> ```

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

|Option|Type|
|------|----|
|[FormattedPrefix](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.GetFormattedPrefix)|string|
|[Debugger](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.WithDebugger)|bool|
|[Fatal](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.WithFatal)|bool|
|[Debugger](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.WithDebugger)|[*Style](https://pkg.go.dev/github.com/pterm/pterm#Style)|
|[Prefix](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.WithPrefix)|[Prefix](https://pkg.go.dev/github.com/pterm/pterm#Prefix)|
|[Scope](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.WithPrefix)|[Scope](https://pkg.go.dev/github.com/pterm/pterm#Scope)|
|[ShowLineNumber](https://pkg.go.dev/github.com/pterm/pterm#PrefixPrinter.WithShowLineNumber)|bool|

### Output functions

> This printer implements the interface [`TextPrinter`](https://github.com/pterm/pterm/blob/master/interface_text_printer.go)

|Function|Description|
|------|---------|
|Sprint(a ...interface{})|Returns a string|
|Sprintln(a ...interface{})|Returns a string with a new line at the end|
|Sprintf(format string, a ...interface{})|Returns a string, formatted according to a format specifier|
|Print(a ...interface{})|Prints to the terminal|
|Println(a ...interface{})|Prints to the terminal with a new line at the end|
|Printf(format string, a ...interface{})|Prints to the terminal, formatted according to a format specifier|

## Related

- [Override default printers](docs/customizing/override-default-printer.md)