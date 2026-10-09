package pterm

import (
	"io"
	"sync"
)

// globalMu guards the package-level configuration variables that pterm reads
// from background goroutines (live printers) and that user code mutates via
// the Enable/Disable/Set helpers. All internal reads should go through the
// accessors in this file; all writes happen inside the matching setter
// functions while holding globalMu.
//
// Direct reads or writes of the exported variables (e.g. `pterm.Output = false`)
// remain supported for backward compatibility but are not concurrency-safe.
// Callers that need to toggle these from multiple goroutines should use the
// EnableX/DisableX/SetX helpers instead.
var globalMu sync.RWMutex

// outputEnabled returns the current value of Output under globalMu.
func outputEnabled() bool {
	globalMu.RLock()
	defer globalMu.RUnlock()

	return Output
}

// rawOutput returns the current value of RawOutput under globalMu.
func rawOutput() bool {
	globalMu.RLock()
	defer globalMu.RUnlock()

	return RawOutput
}

// printColorEnabled returns the current value of PrintColor under globalMu.
func printColorEnabled() bool {
	globalMu.RLock()
	defer globalMu.RUnlock()

	return PrintColor
}

// printDebugMessages returns the current value of PrintDebugMessages under globalMu.
func printDebugMessages() bool {
	globalMu.RLock()
	defer globalMu.RUnlock()

	return PrintDebugMessages
}

// getDefaultWriter returns the current default writer under globalMu.
func getDefaultWriter() io.Writer {
	globalMu.RLock()
	defer globalMu.RUnlock()

	return defaultWriter
}

// GetDefaultOutput returns the writer that pterm currently writes to by default.
// This is the concurrency-safe counterpart to reading defaultWriter directly.
func GetDefaultOutput() io.Writer {
	return getDefaultWriter()
}

// activeProgressBarsMu guards ActiveProgressBarPrinters.
var activeProgressBarsMu sync.RWMutex

// registerProgressBar appends p to ActiveProgressBarPrinters under the registry lock.
func registerProgressBar(p *ProgressbarPrinter) {
	activeProgressBarsMu.Lock()

	ActiveProgressBarPrinters = append(ActiveProgressBarPrinters, p)
	activeProgressBarsMu.Unlock()
}

// snapshotProgressBars returns a copy of ActiveProgressBarPrinters that callers
// can iterate without holding the registry lock.
func snapshotProgressBars() []*ProgressbarPrinter {
	activeProgressBarsMu.RLock()
	defer activeProgressBarsMu.RUnlock()

	if len(ActiveProgressBarPrinters) == 0 {
		return nil
	}

	cp := make([]*ProgressbarPrinter, len(ActiveProgressBarPrinters))
	copy(cp, ActiveProgressBarPrinters)

	return cp
}

// activeSpinnersMu guards activeSpinnerPrinters.
var activeSpinnersMu sync.RWMutex

// registerSpinner appends s to activeSpinnerPrinters under the registry lock.
func registerSpinner(s *SpinnerPrinter) {
	activeSpinnersMu.Lock()

	activeSpinnerPrinters = append(activeSpinnerPrinters, s)
	activeSpinnersMu.Unlock()
}

// snapshotSpinners returns a copy of activeSpinnerPrinters that callers can
// iterate without holding the registry lock.
func snapshotSpinners() []*SpinnerPrinter {
	activeSpinnersMu.RLock()
	defer activeSpinnersMu.RUnlock()

	if len(activeSpinnerPrinters) == 0 {
		return nil
	}

	cp := make([]*SpinnerPrinter, len(activeSpinnerPrinters))
	copy(cp, activeSpinnerPrinters)

	return cp
}
