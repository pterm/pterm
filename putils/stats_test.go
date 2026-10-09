package putils

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestPrintAverageExecutionTime(t *testing.T) {
	var buf bytes.Buffer

	pterm.SetDefaultOutput(&buf)
	t.Cleanup(func() { pterm.SetDefaultOutput(os.Stdout) })

	var indices []int

	err := PrintAverageExecutionTime(3, func(i int) error {
		indices = append(indices, i)

		return nil
	})

	assert.NoError(t, err)
	assert.Equal(t, []int{0, 1, 2}, indices, "the function must be called once per iteration with its index")
	assert.Contains(t, buf.String(), "Average execution time:")
}

func TestPrintAverageExecutionTimeStopsOnError(t *testing.T) {
	var buf bytes.Buffer

	pterm.SetDefaultOutput(&buf)
	t.Cleanup(func() { pterm.SetDefaultOutput(os.Stdout) })

	wantErr := errors.New("iteration failed")
	calls := 0

	err := PrintAverageExecutionTime(5, func(i int) error {
		calls++

		if i == 1 {
			return wantErr
		}

		return nil
	})

	assert.ErrorIs(t, err, wantErr)
	assert.Equal(t, 2, calls, "iterations must stop at the first error")
	assert.NotContains(t, buf.String(), "Average execution time:", "no result must be printed on error")
}
