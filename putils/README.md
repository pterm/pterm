# PUtils - PTerm Utils

This package contains some utility functions, to get you started with PTerm even faster!  

## Util Functions

```go
func RunWithDefaultSpinner(initzialSpinnerText string, f func(spinner *pterm.SpinnerPrinter) error) error
func RunWithSpinner(spinner *pterm.SpinnerPrinter, f func(spinner *pterm.SpinnerPrinter) error) error
```
