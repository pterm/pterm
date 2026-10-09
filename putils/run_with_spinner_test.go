package putils

import (
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestRunWithSpinner(t *testing.T) {
	var inner *pterm.SpinnerPrinter

	err := RunWithSpinner(pterm.DefaultSpinner.WithWriter(io.Discard), func(spinner *pterm.SpinnerPrinter) error {
		inner = spinner

		assert.True(t, spinner.IsActive, "the spinner must be running while the function runs")

		return nil
	})

	assert.NoError(t, err)
	assert.NotNil(t, inner)
	assert.False(t, inner.IsActive, "the spinner must be stopped after the function returned")
}

func TestRunWithSpinnerPassesErrorThrough(t *testing.T) {
	wantErr := errors.New("something went wrong")

	var inner *pterm.SpinnerPrinter

	err := RunWithSpinner(pterm.DefaultSpinner.WithWriter(io.Discard), func(spinner *pterm.SpinnerPrinter) error {
		inner = spinner

		return wantErr
	})

	assert.Same(t, wantErr, err)
	assert.False(t, inner.IsActive, "the spinner must be stopped even if the function failed")
}

func TestRunWithSpinnerDoesNotStopAnAlreadyStoppedSpinner(t *testing.T) {
	err := RunWithSpinner(pterm.DefaultSpinner.WithWriter(io.Discard), func(spinner *pterm.SpinnerPrinter) error {
		return spinner.Stop()
	})

	assert.NoError(t, err)
}
