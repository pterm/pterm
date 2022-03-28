# BulletListPrinter

<!--
Replace all of the following strings with the current printer.
        bulletlist BulletList BulletListPrinter DefaultBulletList
-->

![BulletListPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/bulletlist/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/bulletlist/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{{Level: 0, Text: "Level 0"}}).Render()
```

### Functions

| Function                                                                                                                                   | Description                                                                                                                            |
| ------------------------------------------------------------------------------------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------- |
| [NewBulletListFromString(s string, padding string)](https://pkg.go.dev/github.com/pterm/pterm#TemplatePrinter.NewBulletListFromString)     | NewBulletListFromString returns a BulletListPrinter with Text using the NewTreeListItemFromString method, splitting after return (\n). |
| [NewBulletListFromStrings(s []string, padding string)](https://pkg.go.dev/github.com/pterm/pterm#TemplatePrinter.NewBulletListFromStrings) | NewBulletListFromStrings returns a BulletListPrinter with Text using the NewTreeListItemFromString method.                             |

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultBulletList.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultBulletList.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                     | Type                                                                         |
| ------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| [Items](https://pkg.go.dev/github.com/pterm/pterm#BulletListPrinter.WithItems)             | [[]BulletListItem](https://pkg.go.dev/github.com/pterm/pterm#BulletListItem) |
| [TextStyle](https://pkg.go.dev/github.com/pterm/pterm#BulletListPrinter.WithTextStyle)     | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style)                   |
| [Bullet](https://pkg.go.dev/github.com/pterm/pterm#BulletListPrinter.WithBullet)           | string                                                                       |
| [BulletStyle](https://pkg.go.dev/github.com/pterm/pterm#BulletListPrinter.WithBulletStyle) | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style)                   |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#BulletListPrinter.WithWriter)           | io.Writer                                                                    |

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
