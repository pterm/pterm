package pterm

import "io"

// LivePrinter is a printer which can update it's output live.
type LivePrinter interface {
	// GenericStart runs Start, but returns a LivePrinter.
	// This is used for the interface LivePrinter.
	// You most likely want to use Start instead of this in your program.
	GenericStart() (*LivePrinter, error)

	// GenericStop runs Stop, but returns a LivePrinter.
	// This is used for the interface LivePrinter.
	// You most likely want to use Stop instead of this in your program.
	GenericStop() (*LivePrinter, error)

	SetWriter(writer io.Writer)
}
