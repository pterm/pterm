package pterm

// RenderPrinter is used to display renderable content.
// Example for renderable content is a Table.
type RenderPrinter interface {
	// Render the XXX to the terminal.
	Render() error

	// Srender returns the rendered string of XXX.
	Srender() (string, error)
}
