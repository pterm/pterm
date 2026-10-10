package pterm

import "errors"

var (
	// ErrNotInteractive is returned by interactive printers when standard
	// input is not a terminal and the printer has no default value.
	ErrNotInteractive = errors.New("pterm: standard input is not a terminal")

	// ErrCanceled is returned when the context of a printer ends before the
	// printer finished.
	ErrCanceled = errors.New("pterm: canceled")

	// ErrInterrupted is returned by interactive printers whose interrupt
	// handler returned after the user pressed Ctrl+C.
	ErrInterrupted = errors.New("pterm: interrupted")

	// ErrEmpty is returned by render printers that have nothing to render.
	ErrEmpty = errors.New("pterm: nothing to render")

	// ErrAlreadyActive is returned by Start when a live printer is running.
	ErrAlreadyActive = errors.New("pterm: printer is already running")
)
