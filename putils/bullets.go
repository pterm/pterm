package putils

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
)

// BulletsFromSlice turns any slice into top level list items, formatting
// each element with fmt.Sprint.
//
//	pterm.BulletList.WithItems(putils.BulletsFromSlice(files)).Render()
func BulletsFromSlice[T any](items []T) []pterm.BulletListItem {
	list := make([]pterm.BulletListItem, len(items))
	for i, item := range items {
		list[i] = pterm.BulletListItem{Text: fmt.Sprint(item)}
	}

	return list
}

// BulletsFromIndented turns indented lines into list items. Deeper
// indentation nests an item under the line above it, as in
// TreeFromIndented.
func BulletsFromIndented(s string) []pterm.BulletListItem {
	var (
		list    []pterm.BulletListItem
		indents []int
	)

	for line := range strings.SplitSeq(s, "\n") {
		text := strings.TrimSpace(line)
		if text == "" {
			continue
		}

		indent := len(line) - len(strings.TrimLeft(line, " \t"))

		for len(indents) > 0 && indents[len(indents)-1] >= indent {
			indents = indents[:len(indents)-1]
		}

		list = append(list, pterm.BulletListItem{Level: len(indents), Text: text})
		indents = append(indents, indent)
	}

	return list
}
