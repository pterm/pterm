package main

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type cache struct {
	Name  string
	State string
	Size  string
	Files int
	Path  string `pterm:"Location"`
	inode uint64
}

func main() {
	caches := []cache{
		{Name: "npm", State: "stale", Size: "1.20 GB", Files: 48211, Path: "~/.npm/_cacache"},
		{Name: "go", State: "fresh", Size: "2.84 GB", Files: 112530, Path: "~/go/pkg/mod"},
		{Name: "pip", State: "stale", Size: "87 MB", Files: 1302, Path: "~/.cache/pip"},
	}

	// Field names become the header. A pterm tag renames a column, and
	// unexported fields are left out.
	_ = pterm.Table.WithData(putils.TableFromStructs(caches)).Render()
}
