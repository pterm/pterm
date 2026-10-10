package putils

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/pterm/pterm"
)

// DownloadWithProgress downloads url into w and shows a progress bar with
// title while it runs. Without a Content-Length header the bar cannot know
// the total and shows a spinner instead. The download stops when ctx ends.
//
//	f, _ := os.Create("go.tar.gz")
//	defer f.Close()
//	err := putils.DownloadWithProgress(ctx, "go1.27.tar.gz", url, f)
func DownloadWithProgress(ctx context.Context, title, url string, w io.Writer) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("putils: download: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("putils: download: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("putils: download %s: %s", url, resp.Status)
	}

	if resp.ContentLength <= 0 {
		return RunWithSpinner(ctx, title, func(context.Context, *pterm.SpinnerPrinter) error {
			_, err := io.Copy(w, resp.Body)
			return err
		})
	}

	bar, _ := pterm.Progressbar.
		WithTotal(int(resp.ContentLength)).
		WithShowCount(false).
		WithContext(ctx).
		Start(title)

	if _, err := io.Copy(w, io.TeeReader(resp.Body, progressWriter{bar})); err != nil {
		bar.Fail()
		return fmt.Errorf("putils: download %s: %w", url, err)
	}

	return bar.Stop()
}

type progressWriter struct{ bar *pterm.ProgressbarPrinter }

func (p progressWriter) Write(b []byte) (int, error) {
	p.bar.Add(len(b))
	return len(b), nil
}
