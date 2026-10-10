package pterm

import (
	"context"
	"io"
	"strings"
	"sync"

	"github.com/pterm/pterm/internal/screen"
)

// FlowPrinter groups the output of a multi-step task on a rail. Steps and
// prompts stand on the rail as status lines, and everything else printed
// while the flow runs, from any printer, is drawn next to it:
//
//	┌ disk clean
//	│
//	✔ Measuring what can be freed 5.4s
//	│
//	│ ┃ ▲ Heads up
//	│ ┃ Nothing here can be undone.
//	│
//	? What should be cleaned?
//	└ Done
//
// Start a flow, then use Step for work that takes time, Section to group
// steps, and any other printer as usual. Flows do not nest.
type FlowPrinter struct {
	// Title is shown at the top of the rail.
	Title string
	// Subtitle is shown muted after the title.
	Subtitle string
	// Spacing draws an empty rail line between items.
	Spacing bool
	// Writer receives the output. Nil means the default output.
	Writer io.Writer
	// Context is passed on to the steps of the flow.
	Context context.Context
	// OnInterrupt runs when the user presses Ctrl+C during a step, instead
	// of exiting.
	OnInterrupt func()

	run *flowRun
}

type flowRun struct {
	mu      sync.Mutex
	pop     func()
	steps   []*SpinnerPrinter
	stopped bool
}

// Flow is the default flow.
var Flow = FlowPrinter{Spacing: true}

// WithTitle returns a copy of p with title t.
func (p FlowPrinter) WithTitle(t string) *FlowPrinter {
	p.Title = t
	return &p
}

// WithSubtitle returns a copy of p with subtitle s.
func (p FlowPrinter) WithSubtitle(s string) *FlowPrinter {
	p.Subtitle = s
	return &p
}

// WithSpacing returns a copy of p that does or does not separate items with
// an empty rail line.
func (p FlowPrinter) WithSpacing(b ...bool) *FlowPrinter {
	p.Spacing = flag(b)
	return &p
}

// WithWriter returns a copy of p that writes to w.
func (p FlowPrinter) WithWriter(w io.Writer) *FlowPrinter {
	p.Writer = w
	return &p
}

// WithContext returns a copy of p whose steps stop as canceled when ctx
// ends.
func (p FlowPrinter) WithContext(ctx context.Context) *FlowPrinter {
	p.Context = ctx
	return &p
}

// WithOnInterrupt returns a copy of p whose steps run fn on Ctrl+C instead
// of exiting.
func (p FlowPrinter) WithOnInterrupt(fn func()) *FlowPrinter {
	p.OnInterrupt = fn
	return &p
}

func (p *FlowPrinter) out() io.Writer { return orDefault(p.Writer) }

// Start draws the top of the rail, with the given title or p.Title, and
// returns the running flow. It returns ErrAlreadyActive when another flow
// runs on the same writer.
func (p FlowPrinter) Start(title ...any) (*FlowPrinter, error) {
	if len(title) > 0 {
		p.Title = sprintln(title...)
	}

	s := screen.For(p.out())
	if s.HasDecorators() {
		return nil, ErrAlreadyActive
	}

	c := canvasFor(p.out())

	head := c.paint(ToneBorder, c.g.RailStart)
	if p.Title != "" {
		head += " " + c.bold(TonePrimary, p.Title)
	}

	if p.Subtitle != "" {
		head += "  " + c.paint(ToneMuted, p.Subtitle)
	}

	emit(p.out(), node(head)+"\n")

	p.run = &flowRun{}
	p.run.pop = s.PushDecorator(p.decorate)

	return &p, nil
}

// decorate draws the rail next to output. Node lines replace the rail with
// their glyph; all other lines get the rail in front of them.
func (p *FlowPrinter) decorate(s string, _ bool) string {
	if s == "" {
		return s
	}

	c := canvasFor(p.out())
	rail := c.paint(ToneBorder, c.g.Rail)

	body, newline := strings.CutSuffix(s, "\n")
	lines := strings.Split(body, "\n")

	for i, line := range lines {
		switch {
		case strings.HasPrefix(line, screen.Node):
			// Node lines stand on the rail as they are.
		case line == "":
			lines[i] = rail
		default:
			lines[i] = rail + " " + line
		}
	}

	if p.Spacing {
		lines = append([]string{rail}, lines...)
	}

	out := strings.Join(lines, "\n")
	if newline {
		out += "\n"
	}

	return out
}

// Step starts a spinner on the rail for work that takes time. End it with
// Success, Fail, Warning, Info or Skip; it then stays on the rail as a status
// line with the time it took.
func (p *FlowPrinter) Step(text ...any) *SpinnerPrinter {
	sp := Spinner.WithWriter(p.out()).WithOnInterrupt(p.OnInterrupt)
	if p.Context != nil {
		sp = sp.WithContext(p.Context)
	}

	s, _ := sp.Start(text...)

	if p.run != nil {
		p.run.mu.Lock()
		p.run.steps = append(p.run.steps, s)
		p.run.mu.Unlock()
	}

	return s
}

// Section prints a bold divider on the rail to group the steps after it.
func (p *FlowPrinter) Section(title string) {
	c := canvasFor(p.out())
	emit(p.out(), c.bold(TonePrimary, title)+"\n")
}

// Log prints a muted line on the rail, for details that do not deserve a
// step of their own.
func (p *FlowPrinter) Log(a ...any) {
	c := canvasFor(p.out())
	emit(p.out(), c.paint(ToneMuted, sprintln(a...))+"\n")
}

// Add renders r on the rail. Other printers join the flow on their own when
// they print while it runs; Add reads better when the flow is built in one
// place.
func (p *FlowPrinter) Add(r RenderPrinter) error {
	return r.Render()
}

// Stop closes the rail. Steps that are still running end as a success.
func (p *FlowPrinter) Stop() error {
	p.end(ToneMuted, "")
	return nil
}

// Success closes the rail with text in the success tone.
func (p *FlowPrinter) Success(text ...any) { p.end(ToneSuccess, sprintln(text...)) }

// Fail closes the rail with text in the error tone.
func (p *FlowPrinter) Fail(text ...any) { p.end(ToneError, sprintln(text...)) }

func (p *FlowPrinter) end(t Tone, text string) {
	if p.run == nil {
		return
	}

	p.run.mu.Lock()
	defer p.run.mu.Unlock()

	if p.run.stopped {
		return
	}

	p.run.stopped = true

	for _, step := range p.run.steps {
		_ = step.Stop()
	}

	c := canvasFor(p.out())

	line := c.paint(ToneBorder, c.g.RailEnd)
	if text != "" {
		line += " " + c.paint(t, text)
	}

	// The end line is written before the decorator goes, so it gets the
	// same spacing as every other item.
	emit(p.out(), node(line)+"\n")
	p.run.pop()
}
