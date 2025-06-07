package pterm_test

import (
	"io"
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/pterm/pterm"
)

func TestProgressbarPrinter_NoPanicOnZeroRoundingFactor(t *testing.T) {
	p := pterm.DefaultProgressbar
	p.ElapsedTimeRoundingFactor = 0
	p.ShowElapsedTime = true
	p.Writer = io.Discard
	pb, err := p.Start()
	testza.AssertNoError(t, err)
	testza.AssertNotPanics(t, func() {
		pb.Add(1)
	})
	pb.Stop()
}
