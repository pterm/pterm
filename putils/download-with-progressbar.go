package putils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pterm/pterm"
)

// progressbarWriter counts the number of bytes written to it and adds those to a progressbar.
type progressbarWriter struct {
	pb *pterm.ProgressbarPrinter
}

func (w *progressbarWriter) Write(p []byte) (int, error) {
	w.pb.Add(len(p))
	return len(p), nil
}

// DownloadFileWithProgressbar downloads a file, by url, and writes it to outputPath.
// The download progress, will be reported via a progressbar.
func DownloadFileWithProgressbar(progressbar *pterm.ProgressbarPrinter, outputPath, url string, mode os.FileMode) error {
	path := filepath.Clean(outputPath)

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create download path: %w", err)
	}

	defer func() { _ = out.Close() }()

	resp, err := http.Get(url) //nolint:gosec // the URL is intentionally caller-provided
	if err != nil {
		return fmt.Errorf("error while downloading file: %w", err)
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("error while downloading file: unexpected status %q", resp.Status)
	}

	fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return fmt.Errorf("could not determine file size: %w", err)
	}

	pb, err := progressbar.WithTotal(fileSize).Start()
	if err != nil {
		return fmt.Errorf("could not start progressbar: %w", err)
	}

	counter := &progressbarWriter{pb: pb}
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		return err
	}

	err = os.Chmod(path, mode)
	if err != nil {
		return fmt.Errorf("could not chmod file: %w", err)
	}

	return nil
}

// DownloadFileWithDefaultProgressbar downloads a file, by url, and writes it to outputPath.
// The download progress, will be reported via the default progressbar.
func DownloadFileWithDefaultProgressbar(title, outputPath, url string, mode os.FileMode) error {
	return DownloadFileWithProgressbar(pterm.DefaultProgressbar.WithTitle(title), outputPath, url, mode)
}
