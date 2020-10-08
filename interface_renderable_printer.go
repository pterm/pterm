package pterm

// RenderPrinter is used to display renderable content.
// Example for renderable content is a Table.
type RenderPrinter interface {
	Render()
}
