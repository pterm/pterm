package putils

import (
	"strings"

	"github.com/pterm/pterm"
)

// TableDataFromSeparatedValues converts values, separated by separator, into pterm.TableData.
//
// Usage:
//	pterm.DefaultTable.WithData(putils.TableDataFromCSV(csv)).Render()
func TableDataFromSeparatedValues(text, separator string) (td pterm.TableData) {
	for _, line := range strings.Split(text, "\n") {
		td = append(td, strings.Split(line, separator))
	}

	return
}
