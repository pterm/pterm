package putils

import (
	"strings"

	"github.com/pterm/pterm"
)

// TableDataFromSeparatedValues converts values, separated by separator, into pterm.TableData.
//
// Usage:
//
//	pterm.DefaultTable.WithData(putils.TableDataFromCSV(csv)).Render()
func TableDataFromSeparatedValues(text, valueSeparator, rowSeparator string) (td pterm.TableData) {
	for _, line := range strings.Split(text, rowSeparator) {
		td = append(td, strings.Split(line, valueSeparator))
	}

	return
}
