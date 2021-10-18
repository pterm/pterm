package pterm

type Grouper interface {
	// Append the XXX to the Grouper instance
	Append(items ...interface{}) Grouper

	// Render the XXX to the terminal.
	Render() error

	// Srender returns the rendered string of XXX.
	Srender() (string, error)
}
