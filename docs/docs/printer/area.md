# AreaPrinter

![AreaPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/area/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/area/main.go" target="_blank">(Show source of demo)</a></p>


## Usage

### Basic usage

```go
area, _ := pterm.DefaultArea.Start() // Start the Area printer.
for i := 0; i < 10; i++ {
    str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(time.Now().Format("15:04:05"))).Srender() // Save current time in str.
    str = pterm.DefaultCenter.Sprint(str) // Center str.
    area.Update(str) // Update Area contents.
    time.Sleep(time.Second)
}
area.Stop()
```
### Functions

|Function|Description|
|--------|-----------|
|[Update](https://pkg.go.dev/github.com/pterm/pterm#AreaPrinter.Update)|Update overwrites the content of the AreaPrinter.|
|[GetContent](https://pkg.go.dev/github.com/pterm/pterm#AreaPrinter.GetContent)|GetContent returns the current area content.|
|[Clear](https://pkg.go.dev/github.com/pterm/pterm#AreaPrinter.Clear)|Clear function clears the content of the area.|

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultArea.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultArea.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

|Option|Type|
|------|----|
|[RemoveWhenDone](https://pkg.go.dev/github.com/pterm/pterm#AreaPrinter.RemoveWhenDone)|bool|
|[Fullscreen](https://pkg.go.dev/github.com/pterm/pterm#AreaPrinter.Fullscreen)|bool|
|[Center](https://pkg.go.dev/github.com/pterm/pterm#AreaPrinter.Center)|bool|

### Output functions
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

## Related
- [Override default printers](../customizing/override-default-printer.md)
