# PUtils - PTerm Utils

Package putils contains utility functions for PTerm, to make it's usage even easier!
It helps you to simplify various scenarios for which PTerm is frequently used.

You can read the documentation of this package [here](https://pkg.go.dev/github.com/pterm/pterm/putils#section-documentation).

Feel free to contribute your utility functions via pull request!

## Util Functions

```go
func DownloadFileWithDefaultProgressbar(title, outputPath, url string, mode os.FileMode) error
func DownloadFileWithProgressbar(progressbar *pterm.ProgressbarPrinter, outputPath, url string, mode os.FileMode) error
func RunWithDefaultSpinner(initzialSpinnerText string, f func(spinner *pterm.SpinnerPrinter) error) error
func RunWithSpinner(spinner *pterm.SpinnerPrinter, f func(spinner *pterm.SpinnerPrinter) error) error
```
