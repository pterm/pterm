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
	Total uint64
	pb    *pterm.ProgressbarPrinter
}

func (w *progressbarWriter) Write(p []byte) (int, error) {
	n := len(p)
	w.Total += uint64(n)
	w.pb.Add(len(p))
	return n, nil
}

// DownloadFileWithProgressbar downloads a file, by url, and writes it to outputPath.
// The download progress, will be reported via a progressbar.
func DownloadFileWithProgressbar(progressbar *pterm.ProgressbarPrinter, outputPath, url string, mode os.FileMode) error {
	path := filepath.Clean(outputPath)
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create download path: %w", err)
	}

	resp, err := http.Get(url) //nolint:gosec
	if err != nil {
		out.Close()
		return fmt.Errorf("error while downloading file: %w", err)
	}
	defer resp.Body.Close()

	counter := &progressbarWriter{}
	fileSize, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return fmt.Errorf("could not determine file size: %w", err)
	}

	counter.pb, _ = progressbar.WithTotal(fileSize).Start()
	if _, err = io.Copy(out, io.TeeReader(resp.Body, counter)); err != nil {
		out.Close()
		return err
	}

	err = os.Chmod(path, mode)
	if err != nil {
		return fmt.Errorf("could not chmod file: %w", err)
	}

	out.Close()
	return nil
}

// DownloadFileWithDefaultProgressbar downloads a file, by url, and writes it to outputPath.
// The download progress, will be reported via the default progressbar.
func DownloadFileWithDefaultProgressbar(title, outputPath, url string, mode os.FileMode) error {
	return DownloadFileWithProgressbar(pterm.DefaultProgressbar.WithTitle(title), outputPath, url, mode)
}
