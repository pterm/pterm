package pterm_test

import (
	"io"
	"os"
	"testing"
	"time"

	"github.com/pterm/pterm"
	"github.com/stretchr/testify/assert"
)

func TestSpinnerPrinter_NilPrint(_ *testing.T) {
	p := pterm.SpinnerPrinter{}
	p.Info()
	p.Success()
	p.Warning()
	p.Fail()
}

func TestSpinnerPrinter_Fail(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		pterm.DefaultSpinner.WithWriter(w).Fail(a)
	})
}

func TestSpinnerPrinter_GenericStart(t *testing.T) {
	p := pterm.DefaultSpinner
	started, err := p.GenericStart()
	assert.NoError(t, err)
	// GenericStart returns the started instance; stopping the original would
	// be a no-op and leak the animation goroutine.
	_, _ = (*started).GenericStop()
}

func TestSpinnerPrinter_GenericStartRawOutput(t *testing.T) {
	pterm.DisableStyling()

	defer pterm.EnableStyling()

	p := pterm.DefaultSpinner
	started, err := p.GenericStart()
	assert.NoError(t, err)

	_, _ = (*started).GenericStop()
}

func TestSpinnerPrinter_GenericStop(_ *testing.T) {
	p := pterm.DefaultSpinner
	_, _ = p.GenericStop()
}

func TestSpinnerPrinter_Info(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		pterm.DefaultSpinner.WithWriter(w).Info(a)
	})
}

func TestSpinnerPrinter_Success(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		pterm.DefaultSpinner.WithWriter(w).Success(a)
	})
}

func TestSpinnerPrinter_UpdateText(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		p := pterm.DefaultSpinner

		sp, _ := p.Start()
		defer sp.Stop()

		p.UpdateText("test")

		assert.Equal(t, "test", p.Text)
	})

	t.Run("Override", func(t *testing.T) {
		out := captureStdout(func(w io.Writer) {
			// Set a really long delay to make sure text doesn't get updated before function returns.
			p := pterm.DefaultSpinner.WithDelay(1 * time.Hour).WithWriter(w)

			sp, _ := p.Start("An initial long message")
			defer sp.Stop()

			p.UpdateText("A short message")
		})
		assert.Contains(t, out, "A short message")
	})
}

func TestSpinnerPrinter_UpdateTextRawOutput(t *testing.T) {
	pterm.DisableStyling()

	p := pterm.DefaultSpinner

	sp, _ := p.Start()
	defer sp.Stop()

	p.UpdateText("test")

	assert.Equal(t, "test", p.Text)
	pterm.EnableStyling()
}

func TestSpinnerPrinter_StopSetsIsActiveWhenRawOutput(t *testing.T) {
	// Regression test for https://github.com/pterm/pterm/issues/763
	// Stop() must always set IsActive = false, even when RawOutput is true,
	// to prevent the background goroutine from leaking.
	pterm.DisableStyling()

	defer pterm.EnableStyling()

	sp, err := pterm.DefaultSpinner.Start()
	assert.NoError(t, err)
	assert.True(t, sp.IsActive)

	err = sp.Stop()
	assert.NoError(t, err)
	assert.False(t, sp.IsActive)
}

func TestSpinnerPrinter_Warning(t *testing.T) {
	testPrintContains(t, func(w io.Writer, a any) {
		pterm.DefaultSpinner.WithWriter(w).Warning(a)
	})
}

func TestSpinnerPrinter_WithDelay(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithDelay(time.Second)

	assert.Equal(t, time.Second, p2.Delay)
}

func TestSpinnerPrinter_WithMessageStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.SpinnerPrinter{}
	p2 := p.WithMessageStyle(s)

	assert.Equal(t, s, p2.MessageStyle)
}

func TestSpinnerPrinter_WithRemoveWhenDone(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithRemoveWhenDone()

	assert.True(t, p2.RemoveWhenDone)
}

func TestSpinnerPrinter_WithSequence(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithSequence("a", "b", "c")

	assert.Equal(t, []string{"a", "b", "c"}, p2.Sequence)
}

func TestSpinnerPrinter_WithStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.SpinnerPrinter{}
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
}

func TestSpinnerPrinter_WithText(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithText("test")

	assert.Equal(t, "test", p2.Text)
}

func TestSpinnerPrinter_WithShowTimer(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithShowTimer()

	assert.True(t, p2.ShowTimer)
}

func TestSpinnerPrinter_WithTimerStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.SpinnerPrinter{}
	p2 := p.WithTimerStyle(s)

	assert.Equal(t, s, p2.TimerStyle)
}

func TestSpinnerPrinter_WithTimerRoundingFactor(t *testing.T) {
	s := time.Millisecond * 200
	p := pterm.SpinnerPrinter{}
	p2 := p.WithTimerRoundingFactor(s)

	assert.Equal(t, s, p2.TimerRoundingFactor)
}

func TestSpinnerPrinter_DifferentVariations(t *testing.T) {
	type fields struct {
		Text           string
		Sequence       []string
		Style          *pterm.Style
		Delay          time.Duration
		MessageStyle   *pterm.Style
		InfoPrinter    pterm.TextPrinter
		SuccessPrinter pterm.TextPrinter
		FailPrinter    pterm.TextPrinter
		WarningPrinter pterm.TextPrinter
		RemoveWhenDone bool
		IsActive       bool
	}
	type args struct {
		text []any
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "WithText", fields: fields{Text: "test"}, args: args{}},
		{name: "WithText", fields: fields{}, args: args{[]any{"test"}}},
		{name: "WithRemoveWhenDone", fields: fields{RemoveWhenDone: true}, args: args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := pterm.SpinnerPrinter{
				Text:           tt.fields.Text,
				Sequence:       tt.fields.Sequence,
				Style:          tt.fields.Style,
				Delay:          tt.fields.Delay,
				MessageStyle:   tt.fields.MessageStyle,
				InfoPrinter:    tt.fields.InfoPrinter,
				SuccessPrinter: tt.fields.SuccessPrinter,
				FailPrinter:    tt.fields.FailPrinter,
				WarningPrinter: tt.fields.WarningPrinter,
				RemoveWhenDone: tt.fields.RemoveWhenDone,
				IsActive:       tt.fields.IsActive,
			}
			sp, _ := s.Start(tt.args.text)
			_ = sp.Stop()
		})
	}
}

func TestSpinnerPrinter_WithWriter(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	s := os.Stderr
	p2 := p.WithWriter(s)

	assert.Equal(t, s, p2.Writer)
	assert.Zero(t, p.Writer)
}

func TestSpinnerPrinter_OutputToWriters(t *testing.T) {
	testCases := map[string]struct {
		action                func(*pterm.SpinnerPrinter)
		expectOutputToContain string
	}{
		"ExpectWarningMessageToBeWrittenToStderr": {
			action:                func(sp *pterm.SpinnerPrinter) { sp.Warning("A warning") },
			expectOutputToContain: "A warning",
		},
		"ExpectFailMessageToBeWrittenToStderr": {
			action:                func(sp *pterm.SpinnerPrinter) { sp.Fail("An error") },
			expectOutputToContain: "An error",
		},
		"ExpectUpdatedTextToBeWrittenToStderr": {
			action: func(sp *pterm.SpinnerPrinter) {
				sp.UpdateText("Updated text")
			},
			expectOutputToContain: "Updated text",
		},
	}

	for testTitle, testCase := range testCases {
		t.Run(testTitle, func(t *testing.T) {
			buf := &syncBuffer{}
			sp, err := pterm.DefaultSpinner.WithText("Hello world").WithWriter(buf).Start()
			assert.NoError(t, err)

			defer sp.Stop()

			waitForOutput(t, buf, "Hello world")
			testCase.action(sp)
			waitForOutput(t, buf, testCase.expectOutputToContain)
		})
	}
}

// func TestClearActiveSpinners(t *testing.T) {
// 	activeSpinnerPrinters = []*pterm.SpinnerPrinter{}
// }
