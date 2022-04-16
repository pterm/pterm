package putils

import "github.com/pterm/pterm"

// RunWithSpinner starts a spinner, then runs a function and after the function is done, the spinner will stop again.
func RunWithSpinner(spinner *pterm.SpinnerPrinter, f func(spinner *pterm.SpinnerPrinter) error) error {
	s, err := spinner.Start()
	if err != nil {
		return err
	}

	err = f(s)

	if s.IsActive {
		s.Stop()
	}

	return err
}

// RunWithDefaultSpinner starts a default spinner, then runs a function and after the function is done, the spinner will stop again.
func RunWithDefaultSpinner(initzialSpinnerText string, f func(spinner *pterm.SpinnerPrinter) error) error {
	return RunWithSpinner(pterm.DefaultSpinner.WithText(initzialSpinnerText), f)
}
