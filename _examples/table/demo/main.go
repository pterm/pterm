package main

import "github.com/pterm/pterm"

func main() {
	// The first row is the header. Sizes and numbers are detected and
	// right-aligned, so the SIZE and FILES columns line up on their units.
	pterm.Table.WithData(pterm.TableData{
		{"Cache", "State", "Size", "Files", "Path"},
		{"npm", "stale", "1.20 GB", "48,211", "~/.npm/_cacache"},
		{"bun", "fresh", "393 MB", "9,874", "~/.bun/install/cache"},
		{"go", "fresh", "2.84 GB", "112,530", "~/go/pkg/mod"},
		{"pip", "stale", "87 MB", "1,302", "~/.cache/pip"},
		{"cargo", "fresh", "640 MB", "20,447", "~/.cargo/registry"},
	}).Render()
}
