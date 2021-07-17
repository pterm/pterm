package pterm

import (
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSpinnerPrinter_NilPrint(t *testing.T) {
	p := SpinnerPrinter{}
	p.Success()
	p.Warning()
	p.Fail()
}

func TestSpinnerPrinter_Fail(t *testing.T) {
	p := DefaultSpinner
	testPrintContains(t, func(w io.Writer, a interface{}) {
		p.Fail(a)
	})
}

func TestSpinnerPrinter_GenericStart(t *testing.T) {
	p := DefaultSpinner
	p.GenericStart()
	p.GenericStop()
}

func TestSpinnerPrinter_GenericStartRawOutput(t *testing.T) {
	DisableStyling()
	p := DefaultSpinner
	p.GenericStart()
	p.GenericStop()
	EnableStyling()
}

func TestSpinnerPrinter_GenericStop(t *testing.T) {
	p := DefaultSpinner
	p.GenericStop()
}

func TestSpinnerPrinter_Success(t *testing.T) {
	p := DefaultSpinner
	testPrintContains(t, func(w io.Writer, a interface{}) {
		p.Success(a)
	})
}

func TestSpinnerPrinter_UpdateText(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		p := DefaultSpinner
		p.Start()
		p.UpdateText("test")

		assert.Equal(t, "test", p.Text)
	})

	t.Run("Override", func(t *testing.T) {
		out := captureStdout(func(io.Writer) {
			// Set a really long delay to make sure text doesn't get updated before function returns.
			p := DefaultSpinner.WithDelay(1 * time.Hour)
			p.Start("An initial long message")
			p.UpdateText("A short message")
		})
		assert.Contains(t, out, "A short message")
	})
}

func TestSpinnerPrinter_UpdateTextRawOutput(t *testing.T) {
	DisableStyling()
	p := DefaultSpinner
	p.Start()
	p.UpdateText("test")

	assert.Equal(t, "test", p.Text)
	p.Stop()
	EnableStyling()
}

func TestSpinnerPrinter_Warning(t *testing.T) {
	p := DefaultSpinner
	testPrintContains(t, func(w io.Writer, a interface{}) {
		p.Warning(a)
	})
}

func TestSpinnerPrinter_WithDelay(t *testing.T) {
	p := SpinnerPrinter{}
	p2 := p.WithDelay(time.Second)

	assert.Equal(t, time.Second, p2.Delay)
}

func TestSpinnerPrinter_WithMessageStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := SpinnerPrinter{}
	p2 := p.WithMessageStyle(s)

	assert.Equal(t, s, p2.MessageStyle)
}

func TestSpinnerPrinter_WithRemoveWhenDone(t *testing.T) {
	p := SpinnerPrinter{}
	p2 := p.WithRemoveWhenDone()

	assert.True(t, p2.RemoveWhenDone)
}

func TestSpinnerPrinter_WithSequence(t *testing.T) {
	p := SpinnerPrinter{}
	p2 := p.WithSequence("a", "b", "c")

	assert.Equal(t, []string{"a", "b", "c"}, p2.Sequence)
}

func TestSpinnerPrinter_WithStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := SpinnerPrinter{}
	p2 := p.WithStyle(s)

	assert.Equal(t, s, p2.Style)
}

func TestSpinnerPrinter_WithText(t *testing.T) {
	p := SpinnerPrinter{}
	p2 := p.WithText("test")

	assert.Equal(t, "test", p2.Text)
}

func TestSpinnerPrinter_WithShowTimer(t *testing.T) {
	p := SpinnerPrinter{}
	p2 := p.WithShowTimer()

	assert.True(t, p2.ShowTimer)
}

func TestSpinnerPrinter_WithTimerStyle(t *testing.T) {
	s := NewStyle(FgRed, BgBlue, Bold)
	p := SpinnerPrinter{}
	p2 := p.WithTimerStyle(s)

	assert.Equal(t, s, p2.TimerStyle)
}

func TestSpinnerPrinter_WithTimerRoundingFactor(t *testing.T) {
	s := time.Millisecond * 200
	p := SpinnerPrinter{}
	p2 := p.WithTimerRoundingFactor(s)

	assert.Equal(t, s, p2.TimerRoundingFactor)
}

func TestSpinnerPrinter_WithRawOutput(t *testing.T) {
	RawOutput = true
	s, _ := DefaultSpinner.Start()
	go func() {
		time.Sleep(time.Millisecond * 50)
		s.Stop()
		RawOutput = false
	}()
}

func TestSpinnerPrinter_DifferentVariations(t *testing.T) {
	type fields struct {
		Text           string
		Sequence       []string
		Style          *Style
		Delay          time.Duration
		MessageStyle   *Style
		SuccessPrinter TextPrinter
		FailPrinter    TextPrinter
		WarningPrinter TextPrinter
		RemoveWhenDone bool
		IsActive       bool
	}
	type args struct {
		text []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "WithText", fields: fields{Text: "test"}, args: args{}},
		{name: "WithText", fields: fields{}, args: args{[]interface{}{"test"}}},
		{name: "WithRemoveWhenDone", fields: fields{RemoveWhenDone: true}, args: args{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SpinnerPrinter{
				Text:           tt.fields.Text,
				Sequence:       tt.fields.Sequence,
				Style:          tt.fields.Style,
				Delay:          tt.fields.Delay,
				MessageStyle:   tt.fields.MessageStyle,
				SuccessPrinter: tt.fields.SuccessPrinter,
				FailPrinter:    tt.fields.FailPrinter,
				WarningPrinter: tt.fields.WarningPrinter,
				RemoveWhenDone: tt.fields.RemoveWhenDone,
				IsActive:       tt.fields.IsActive,
			}
			s.Start(tt.args.text)
			s.Stop()
		})
	}
}

func TestClearActiveSpinners(t *testing.T) {
	activeSpinnerPrinters = []*SpinnerPrinter{}
}
