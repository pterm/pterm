package putils

import (
	"context"

	"github.com/pterm/pterm"
)

// RunWithSpinner shows a spinner with text while fn runs. The spinner ends
// as a success when fn returns nil. Otherwise it ends as a failure that
// reads "text: error", and the error is returned as is. fn can change the spinner text
// through s, and should stop when ctx ends.
//
//	err := putils.RunWithSpinner(ctx, "Fetching releases", func(ctx context.Context, s *pterm.SpinnerPrinter) error {
//		releases, err = client.Releases(ctx)
//		return err
//	})
func RunWithSpinner(ctx context.Context, text string, fn func(ctx context.Context, s *pterm.SpinnerPrinter) error) error {
	s, _ := pterm.Spinner.WithContext(ctx).Start(text)

	if err := fn(ctx, s); err != nil {
		s.Fail(text + ": " + err.Error())
		return err
	}

	s.Success()

	return nil
}
