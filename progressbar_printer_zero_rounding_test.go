package pterm_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pterm/pterm"
)

func TestProgressbarPrinter_NoPanicOnZeroRoundingFactor(t *testing.T) {
	p := pterm.DefaultProgressbar
	p.ElapsedTimeRoundingFactor = 0
	p.ShowElapsedTime = true
	p.Writer = io.Discard
	pb, err := p.Start()
	assert.NoError(t, err)
	assert.NotPanics(t, func() {
		pb.Add(1)
	})

	_, _ = pb.Stop()
}
