# ProgressbarPrinter

<!--
Replace all of the following strings with the current printer.
     progressbar Progressbar ProgressbarPrinter DefaultProgressbar
-->

![ProgressbarPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/progressbar/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/progressbar/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
progressbar := pterm.DefaultProgressbar.WithTotal(totalSteps).Start()
// Logic here
progressbar.Increment()
// More logic
```

### Functions

| Function                                                                                              | Description                                                                        |
| ----------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| [Add(count int)](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.Add)                    | Add `count` to current value.                                                      |
| [GetElapsedTime()](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.GetElapsedTime)       | GetElapsedTime returns the elapsed time, since the ProgressbarPrinter was started. |
| [Increment()](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.Increment)                 | Increment current value by one.                                                    |
| [UpdateTitle(title string)](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.UpdateTitle) | Update the progressbar's title.                                                    |

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultProgressbar.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultProgressbar.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                                  | Type                                                       |
| ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| [BarCharacter](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithBarCharacter)                           | string                                                     |
| [BarStyle](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithBarStyle)                                   | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [Current](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithCurrent)                                     | int                                                        |
| [ElapsedTimeRoundingFactor](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithElapsedTimeRoundingFactor) | time.Duration                                              |
| [LastCharacter](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithLastCharacter)                         | string                                                     |
| [RemoveWhenDone](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithRemoveWhenDone)                       | ...bool                                                    |
| [ShowCount](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithShowCount)                                 | ...bool                                                    |
| [ShowElapsedTime](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithShowElapsedTime)                     | ...bool                                                    |
| [ShowPercentage](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithShowPercentage)                       | ...bool                                                    |
| [ShowTitle](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithShowTitle)                                 | ...bool                                                    |
| [Title](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithTitle)                                         | string                                                     |
| [TitleStyle](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithTitleStyle)                               | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [Total](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithTotal)                                         | int                                                        |
| [BarFiller](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithBarFiller)                                 | string                                                     |
| [MaxWidth](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithMaxWidth)                                   | int                                                        |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#ProgressbarPrinter.WithWriter)                                       | io.Writer                                                  |

### Output functions

> This printer implements the interface [`LivePrinter`](https://github.com/pterm/pterm/blob/master/interface_live_printer.go)

| Function       | Description                                         |
| -------------- | --------------------------------------------------- |
| Start()        | Returns itself and possible errors                  |
| Stop()         | Returns itself and possible errors                  |
| GenericStart() | Returns the started LivePrinter and possible errors |
| GenericStop()  | Returns the stopped LivePrinter and possible errors |

> [!NOTE]
> The generic start and stop methods are only used to implement the printer into the interface.
> Use the normal `Start()` and `Stop()` methods if possible.

## Related

- [Override default printers](docs/customizing/override-default-printer.md)
