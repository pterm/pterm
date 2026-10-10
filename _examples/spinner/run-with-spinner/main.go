package main

import (
	"context"
	"errors"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	ctx := context.Background()

	// The spinner ends as a success when the function returns nil.
	_ = putils.RunWithSpinner(ctx, "Fetching releases", func(ctx context.Context, s *pterm.SpinnerPrinter) error {
		time.Sleep(time.Second)
		s.UpdateText("Fetched 12 releases")

		return nil
	})

	// An error ends it as a failure that names the task and the error.
	err := putils.RunWithSpinner(ctx, "Uploading report", func(ctx context.Context, _ *pterm.SpinnerPrinter) error {
		time.Sleep(time.Second)
		return errors.New("the bucket does not exist")
	})
	if err != nil {
		pterm.Info.Println("Create the bucket and run the upload again.")
	}
}
