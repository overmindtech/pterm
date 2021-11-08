package pterm_test

import (
	"io"
	"testing"
	"time"

	"github.com/MarvinJWendt/testza"
	"github.com/pterm/pterm"
)

func TestSpinnerPrinter_NilPrint(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p.Success()
	p.Warning()
	p.Fail()
}

func TestSpinnerPrinter_Fail(t *testing.T) {
	p := pterm.DefaultSpinner
	testPrintContains(t, func(w io.Writer, a interface{}) {
		p.Fail(a)
	})
}

func TestSpinnerPrinter_GenericStart(t *testing.T) {
	p := pterm.DefaultSpinner
	p.GenericStart()
	p.GenericStop()
}

func TestSpinnerPrinter_GenericStartRawOutput(t *testing.T) {
	pterm.DisableStyling()
	p := pterm.DefaultSpinner
	p.GenericStart()
	p.GenericStop()
	pterm.EnableStyling()
}

func TestSpinnerPrinter_GenericStop(t *testing.T) {
	p := pterm.DefaultSpinner
	p.GenericStop()
}

func TestSpinnerPrinter_Success(t *testing.T) {
	p := pterm.DefaultSpinner
	testPrintContains(t, func(w io.Writer, a interface{}) {
		p.Success(a)
	})
}

func TestSpinnerPrinter_UpdateText(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		p := pterm.DefaultSpinner
		p.Start()
		p.UpdateText("test")

		testza.AssertEqual(t, "test", p.Text)
	})

	t.Run("Override", func(t *testing.T) {
		out := captureStdout(func(io.Writer) {
			// Set a really long delay to make sure text doesn't get updated before function returns.
			p := pterm.DefaultSpinner.WithDelay(1 * time.Hour)
			p.Start("An initial long message")
			p.UpdateText("A short message")
		})
		testza.AssertContains(t, out, "A short message")
		testza.SnapshotCreateOrValidate(t, t.Name(), out)
	})
}

func TestSpinnerPrinter_UpdateTextRawOutput(t *testing.T) {
	pterm.DisableStyling()
	p := pterm.DefaultSpinner
	p.Start()
	p.UpdateText("test")

	testza.AssertEqual(t, "test", p.Text)
	p.Stop()
	pterm.EnableStyling()
}

func TestSpinnerPrinter_Warning(t *testing.T) {
	p := pterm.DefaultSpinner
	testPrintContains(t, func(w io.Writer, a interface{}) {
		p.Warning(a)
	})
}

func TestSpinnerPrinter_WithDelay(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithDelay(time.Second)

	testza.AssertEqual(t, time.Second, p2.Delay)
}

func TestSpinnerPrinter_WithMessageStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.SpinnerPrinter{}
	p2 := p.WithMessageStyle(s)

	testza.AssertEqual(t, s, p2.MessageStyle)
}

func TestSpinnerPrinter_WithRemoveWhenDone(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithRemoveWhenDone()

	testza.AssertTrue(t, p2.RemoveWhenDone)
}

func TestSpinnerPrinter_WithSequence(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithSequence("a", "b", "c")

	testza.AssertEqual(t, []string{"a", "b", "c"}, p2.Sequence)
}

func TestSpinnerPrinter_WithStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.SpinnerPrinter{}
	p2 := p.WithStyle(s)

	testza.AssertEqual(t, s, p2.Style)
}

func TestSpinnerPrinter_WithText(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithText("test")

	testza.AssertEqual(t, "test", p2.Text)
}

func TestSpinnerPrinter_WithShowTimer(t *testing.T) {
	p := pterm.SpinnerPrinter{}
	p2 := p.WithShowTimer()

	testza.AssertTrue(t, p2.ShowTimer)
}

func TestSpinnerPrinter_WithTimerStyle(t *testing.T) {
	s := pterm.NewStyle(pterm.FgRed, pterm.BgBlue, pterm.Bold)
	p := pterm.SpinnerPrinter{}
	p2 := p.WithTimerStyle(s)

	testza.AssertEqual(t, s, p2.TimerStyle)
}

func TestSpinnerPrinter_WithTimerRoundingFactor(t *testing.T) {
	s := time.Millisecond * 200
	p := pterm.SpinnerPrinter{}
	p2 := p.WithTimerRoundingFactor(s)

	testza.AssertEqual(t, s, p2.TimerRoundingFactor)
}

func TestSpinnerPrinter_WithRawOutput(t *testing.T) {
	pterm.RawOutput = true
	s, _ := pterm.DefaultSpinner.Start()
	go func() {
		time.Sleep(time.Millisecond * 50)
		s.Stop()
		pterm.RawOutput = false
	}()
}

func TestSpinnerPrinter_DifferentVariations(t *testing.T) {
	type fields struct {
		Text           string
		Sequence       []string
		Style          *pterm.Style
		Delay          time.Duration
		MessageStyle   *pterm.Style
		SuccessPrinter pterm.TextPrinter
		FailPrinter    pterm.TextPrinter
		WarningPrinter pterm.TextPrinter
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
			s := pterm.SpinnerPrinter{
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

// func TestClearActiveSpinners(t *testing.T) {
// 	activeSpinnerPrinters = []*pterm.SpinnerPrinter{}
// }
