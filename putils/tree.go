package putils

import (
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"

	"github.com/pterm/pterm"
)

// TreeFromMap turns nested maps into a tree under a root with text root.
// Keys are sorted. A nested map or slice becomes a branch; any other value
// becomes a leaf with the value as its description.
//
//	pterm.Tree.WithRoot(putils.TreeFromMap("config", cfg)).Render()
func TreeFromMap[V any](root string, m map[string]V) pterm.TreeNode {
	node := pterm.TreeNode{Text: root}

	for _, k := range slices.Sorted(maps.Keys(m)) {
		node.Children = append(node.Children, treeValue(k, reflect.ValueOf(m[k])))
	}

	return node
}

func treeValue(text string, v reflect.Value) pterm.TreeNode {
	for v.IsValid() && (v.Kind() == reflect.Pointer || v.Kind() == reflect.Interface) {
		if v.IsNil() {
			return pterm.TreeNode{Text: text}
		}

		v = v.Elem()
	}

	node := pterm.TreeNode{Text: text}

	switch v.Kind() {
	case reflect.Map:
		keys := v.MapKeys()
		slices.SortFunc(keys, func(a, b reflect.Value) int {
			return strings.Compare(fmt.Sprint(a.Interface()), fmt.Sprint(b.Interface()))
		})

		for _, k := range keys {
			node.Children = append(node.Children, treeValue(fmt.Sprint(k.Interface()), v.MapIndex(k)))
		}

	case reflect.Slice, reflect.Array:
		for i := range v.Len() {
			item := v.Index(i)
			for item.Kind() == reflect.Interface && !item.IsNil() {
				item = item.Elem()
			}

			switch item.Kind() {
			case reflect.Map, reflect.Slice, reflect.Array:
				node.Children = append(node.Children, treeValue(fmt.Sprint(i), item))
			default:
				node.Children = append(node.Children, pterm.TreeNode{Text: format(item)})
			}
		}

	case reflect.Invalid:
	default:
		node.Description = format(v)
	}

	return node
}

// TreeFromIndented builds a tree from indented lines, the way a file tree
// or an outline is written by hand. Each level is indented further than
// its parent, with spaces or tabs; blank lines are skipped. The result has
// an empty root, so only the lines are printed.
//
//	putils.TreeFromIndented(`
//	src
//	  main.go
//	  internal
//	    db.go
//	README.md`)
func TreeFromIndented(s string) pterm.TreeNode {
	var list pterm.LeveledList

	var indents []int // the indent of every open level

	for line := range strings.SplitSeq(s, "\n") {
		text := strings.TrimSpace(line)
		if text == "" {
			continue
		}

		indent := len(line) - len(strings.TrimLeft(line, " \t"))

		for len(indents) > 0 && indents[len(indents)-1] >= indent {
			indents = indents[:len(indents)-1]
		}

		list = append(list, pterm.LeveledListItem{Level: len(indents), Text: text})
		indents = append(indents, indent)
	}

	return list.Tree()
}
