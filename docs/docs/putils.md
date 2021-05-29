# PUtils - PTerm Utils

Package putils contains utility functions for PTerm, to make it's usage even easier!
It helps you to simplify various scenarios for which PTerm is frequently used.

## Util Functions

```go
func RunWithDefaultSpinner(initzialSpinnerText string, f func(spinner *pterm.SpinnerPrinter) error) error
func RunWithSpinner(spinner *pterm.SpinnerPrinter, f func(spinner *pterm.SpinnerPrinter) error) error
```
