package pterm_test

import (
	"errors"
	"testing"

	"github.com/pterm/pterm"
)

func TestFlow(t *testing.T) {
	term := setup(t)

	flow, err := pterm.Flow.WithSubtitle("high-frequency records").Start("capture")
	if err != nil {
		t.Fatal(err)
	}

	step := flow.Step("Measuring what can be freed")

	pterm.Println("a plain line")
	step.Success()
	pterm.Callout.Warn("Heads up").Println("Nothing here can be undone.")
	flow.Section("Packets")
	pterm.Success.Println("3 packets captured")
	flow.Success("Capture finished")

	pterm.Println("after the flow")

	got := stripTimer(term.Screen())
	want := `┌ capture  high-frequency records
│
│ a plain line
│
✔ Measuring what can be freed
│
│ ┃ ▲ Heads up
│ ┃ Nothing here can be undone.
│
│ Packets
│
✔ 3 packets captured
│
└ Capture finished
after the flow`

	if got != want {
		t.Errorf("screen:\n%s\nwant:\n%s", got, want)
	}
}

func TestFlowDoesNotNest(t *testing.T) {
	setup(t)

	flow, _ := pterm.Flow.Start("outer")
	defer flow.Stop()

	if _, err := pterm.Flow.Start("inner"); !errors.Is(err, pterm.ErrAlreadyActive) {
		t.Errorf("nested Start error = %v, want ErrAlreadyActive", err)
	}
}

func TestFlowPlain(t *testing.T) {
	setup(t)

	out := capture(t, func() {
		flow, _ := pterm.Flow.WithSpacing(false).Start("build")
		flow.Step("compile")
		pterm.Info.Println("cached")
		_ = flow.Stop()
	})

	want := "┌ build\n● cached\n✔ compile\n└\n"
	if got := stripTimer(out); got != want {
		t.Errorf("output = %q, want %q", got, want)
	}
}
