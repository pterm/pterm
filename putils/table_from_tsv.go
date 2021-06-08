package putils

import (
	"strings"

	"github.com/pterm/pterm"
)

// TableDataFromTSV converts TSV data into pterm.TableData.
//
// Usage:
//	pterm.DefaultTable.WithData(putils.TableDataFromTSV(tsv)).Render()
func TableDataFromTSV(csv string) (td pterm.TableData) {
	for _, line := range strings.Split(csv, "\n") {
		td = append(td, strings.Split(line, "\t"))
	}

	return
}
