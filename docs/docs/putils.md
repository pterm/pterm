# PUtils - PTerm Utils

Package putils contains utility functions for PTerm, to make it's usage even easier!\
It contains pre-made functions, that utilize PTerm, to print various stuff to the terminal.\
You can use PUtils, to simplify various scenarios for which PTerm is frequently used.\

You can read the documentation of this package [here](https://pkg.go.dev/github.com/pterm/pterm/putils#section-documentation).

Feel free to contribute your utility functions via pull request!

## Util Functions

```go
func DefaultTableFromStructSlice(structSlice interface{}) *pterm.TablePrinter
func DownloadFileWithDefaultProgressbar(title, outputPath, url string, mode os.FileMode) error
func DownloadFileWithProgressbar(progressbar *pterm.ProgressbarPrinter, outputPath, url string, mode os.FileMode) error
func NewBulletListFromString(s string, padding string) pterm.BulletListPrinter
func NewBulletListFromStrings(s []string, padding string) pterm.BulletListPrinter
func NewBulletListItemFromString(text string, padding string) pterm.BulletListItem
func NewLettersFromString(text string) pterm.Letters
func NewLettersFromStringWithRGB(text string, rgb pterm.RGB) pterm.Letters
func NewLettersFromStringWithStyle(text string, style *pterm.Style) pterm.Letters
func NewTreeFromLeveledList(leveledListItems pterm.LeveledList) pterm.TreeNode
func PrintAverageExecutionTime(count int, f func(i int) error) error
func RunWithDefaultSpinner(initzialSpinnerText string, f func(spinner *pterm.SpinnerPrinter) error) error
func RunWithSpinner(spinner *pterm.SpinnerPrinter, f func(spinner *pterm.SpinnerPrinter) error) error
func TableDataFromCSV(csv string) (td pterm.TableData)
func TableDataFromSeparatedValues(text, valueSeparator, rowSeparator string) (td pterm.TableData)
func TableDataFromTSV(csv string) (td pterm.TableData)
func TableFromStructSlice(tablePrinter pterm.TablePrinter, structSlice interface{}) *pterm.TablePrinter
```
