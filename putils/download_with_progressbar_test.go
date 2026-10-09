package putils

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pterm/pterm"
)

func discardProgressbar() *pterm.ProgressbarPrinter {
	return pterm.DefaultProgressbar.WithWriter(io.Discard)
}

func TestDownloadFileWithProgressbar(t *testing.T) {
	payload := bytes.Repeat([]byte("pterm test payload "), 1000)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Length", strconv.Itoa(len(payload)))
		_, _ = w.Write(payload)
	}))
	t.Cleanup(server.Close)

	outputPath := filepath.Join(t.TempDir(), "download.bin")

	err := DownloadFileWithProgressbar(discardProgressbar(), outputPath, server.URL, 0o600)
	require.NoError(t, err)

	written, err := os.ReadFile(outputPath) //nolint:gosec // Path is test-controlled.
	require.NoError(t, err)
	assert.Equal(t, payload, written)

	info, err := os.Stat(outputPath)
	require.NoError(t, err)
	assert.Equal(t, os.FileMode(0o600), info.Mode().Perm())
}

func TestDownloadFileWithProgressbarInvalidOutputPath(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte("data"))
	}))
	t.Cleanup(server.Close)

	outputPath := filepath.Join(t.TempDir(), "missing-dir", "download.bin")

	err := DownloadFileWithProgressbar(discardProgressbar(), outputPath, server.URL, 0o600)
	assert.ErrorContains(t, err, "could not create download path")
}

func TestDownloadFileWithProgressbarConnectionError(t *testing.T) {
	// A server that is closed immediately guarantees a connection error.
	server := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	server.Close()

	outputPath := filepath.Join(t.TempDir(), "download.bin")

	err := DownloadFileWithProgressbar(discardProgressbar(), outputPath, server.URL, 0o600)
	assert.ErrorContains(t, err, "error while downloading file")
}

func TestDownloadFileWithProgressbarMissingContentLength(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		// Flushing before writing forces chunked transfer encoding, so the
		// response carries no Content-Length header.
		w.(http.Flusher).Flush()
		_, _ = w.Write([]byte("data"))
	}))
	t.Cleanup(server.Close)

	outputPath := filepath.Join(t.TempDir(), "download.bin")

	err := DownloadFileWithProgressbar(discardProgressbar(), outputPath, server.URL, 0o600)
	assert.ErrorContains(t, err, "could not determine file size")
}
