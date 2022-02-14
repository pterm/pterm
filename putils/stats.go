package putils

import (
	"fmt"
	"time"

	"github.com/pterm/pterm"
)

// PrintAverageExecutionTime times the average execution time of a function.
func PrintAverageExecutionTime(count int, f func(i int) error) error {
	var total time.Duration
	for i := 0; i < count; i++ {
		start := time.Now()
		err := f(i)
		duration := time.Since(start)

		if err != nil {
			return fmt.Errorf("error while calculating average execution time: %w", err)
		}

		total += duration
	}

	averageExecutionTime := total / time.Duration(count)

	pterm.Printfln(pterm.Cyan("Average execution time: %s"), pterm.NewStyle(pterm.Bold, pterm.FgLightCyan).Sprint(averageExecutionTime))

	return nil
}
