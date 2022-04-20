# SpinnerPrinter

<!--
Replace all of the following strings with the current printer.
     spinner Spinner SpinnerPrinter DefaultSpinner
-->

![SpinnerPrinter Example](https://raw.githubusercontent.com/pterm/pterm/master/_examples/spinner/animation.svg)

<p align="center"><a href="https://github.com/pterm/pterm/blob/master/_examples/spinner/main.go" target="_blank">(Show source of demo)</a></p>

## Usage

### Basic usage

```go
pterm.DefaultSpinner.Start()
```

### Functions

| Function                                                                                               | Description                                                                                                                         |
| ------------------------------------------------------------------------------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------- |
| [Fail(message ...interface{})](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.Fail)          | Fail displays the fail printer. If no message is given, the text of the SpinnerPrinter will be reused as the default message.       |
| [Success(message ...interface{})](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.Success)    | Success displays the success printer. If no message is given, the text of the SpinnerPrinter will be reused as the default message. |
| [UpdateText(text string)](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.UpdateText)         | UpdateText updates the message of the active SpinnerPrinter. Can be used live.                                                      |
| [Warning(message ...interface{})](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.UpdateText) | Warning displays the warning printer. If no message is given, the text of the SpinnerPrinter will be reused as the default message. |

### Options

> To make a copy with modified options you can use:
> `pterm.DefaultSpinner.WithOptionName(option)`
>
> To change multiple options at once, you can chain the functions:
> `pterm.DefaultSpinner.WithOptionName(option).WithOptionName2(option2)...`

> [!TIP]
> Click the options and types to show the documentation on _pkg.go.dev_

| Option                                                                                                  | Type                                                       |
| ------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------- |
| [Delay](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithDelay)                             | time.Duration                                              |
| [MessageStyle](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithMessageStyle)               | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [RemoveWhenDone](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithRemoveWhenDone)           | ...bool                                                    |
| [Sequence](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithSequence)                       | ...string                                                  |
| [Style](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithStyle)                             | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [Text](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithText)                               | string                                                     |
| [ShowTimer](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithShowTimer)                     | ...bool                                                    |
| [TimerRoundingFactor](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithTimerRoundingFactor) | time.Duration                                              |
| [TimerStyle](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithTimerStyle)                   | [\*Style](https://pkg.go.dev/github.com/pterm/pterm#Style) |
| [Writer](https://pkg.go.dev/github.com/pterm/pterm#SpinnerPrinter.WithWriter)                           | io.Writer                                                  |

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
