package putils

import (
	"strings"

	"github.com/pterm/pterm"
)

// TableDataFromCSV converts CSV data into pterm.TableData.
//
// Usage:
//	pterm.DefaultTable.WithData(putils.TableDataFromCSV(csv)).Render()
func TableDataFromCSV(csv string) (td pterm.TableData) {
	for _, line := range strings.Split(csv, "\n") {
		td = append(td, strings.Split(line, ","))
	}

	return
}
