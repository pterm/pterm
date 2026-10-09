package main

import "github.com/pterm/pterm"

func main() {
	tableData := pterm.TableData{
		{"Service", "Region", "Status"},
		{"api", "eu-west-1", "running"},
		{"worker", "eu-west-1", "running"},
		{"billing", "us-east-1", "degraded"},
	}

	// Any BoxPrinter works here, so the table box can get its own title and
	// border color. WithBoxPrinter also turns boxing on.
	box := pterm.DefaultBox.WithTitle("Services").WithBoxStyle(pterm.NewStyle(pterm.FgCyan))

	pterm.DefaultTable.WithHasHeader().WithBoxPrinter(*box).WithData(tableData).Render()
}
