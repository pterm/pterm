package main

import (
	"github.com/pterm/pterm"
)

func main() {
	// var tis []pterm.TreeListItem
	//
	// tis = []pterm.TreeListItem{
	// 	{Children: []pterm.TreeListItem{
	// 		{Children: []pterm.TreeListItem{
	// 			{Children: []pterm.TreeListItem{
	// 				{Children: nil, ItemName: "2"},
	// 				{Children: nil, ItemName: "3"},
	// 				{Children: nil, ItemName: "4"},
	// 			}, ItemName: "2"},
	// 			{Children: nil, ItemName: "3"},
	// 			{Children: nil, ItemName: "4"},
	// 		}, ItemName: "2"},
	// 		{Children: nil, ItemName: "3"},
	// 		{Children: nil, ItemName: "4"},
	// 	}, ItemName: "0"},
	// 	{Children: nil, ItemName: "1"},
	// 	{Children: []pterm.TreeListItem{
	// 		{Children: []pterm.TreeListItem{
	// 			{Children: []pterm.TreeListItem{
	// 				{Children: []pterm.TreeListItem{
	// 					{Children: []pterm.TreeListItem{
	// 						{Children: nil, ItemName: "2"},
	// 						{Children: nil, ItemName: "3"},
	// 						{Children: nil, ItemName: "4"},
	// 					}, ItemName: "2"},
	// 					{Children: nil, ItemName: "3"},
	// 					{Children: nil, ItemName: "4"},
	// 				}, ItemName: "2"},
	// 				{Children: nil, ItemName: "3"},
	// 				{Children: nil, ItemName: "4"},
	// 			}, ItemName: "0"},
	// 		}, ItemName: "2"},
	// 		{Children: nil, ItemName: "3"},
	// 		{Children: nil, ItemName: "4"},
	// 	}, ItemName: "2"},
	// 	{Children: nil, ItemName: "3"},
	// }

	tis2 := pterm.LvlTreeListItems{
		pterm.LvlTreeListItem{
			Level: 0,
			Text:  "0",
		},
		pterm.LvlTreeListItem{
			Level: 1,
			Text:  "1",
		},
		pterm.LvlTreeListItem{
			Level: 1,
			Text:  "1.1",
		},
		pterm.LvlTreeListItem{
			Level: 0,
			Text:  "0.1",
		},
		pterm.LvlTreeListItem{
			Level: 0,
			Text:  "0.2",
		},
		pterm.LvlTreeListItem{
			Level: 1,
			Text:  "12",
		},
		pterm.LvlTreeListItem{
			Level: 1,
			Text:  "12.1",
		},
		pterm.LvlTreeListItem{
			Level: 0,
			Text:  "0.3",
		},
		pterm.LvlTreeListItem{
			Level: 0,
			Text:  "0.4",
		},
		pterm.LvlTreeListItem{
			Level: 0,
			Text:  "0.5",
		}}

	tis3 := tis2.ConvertLeveledListToTreeListItems(0, 0)

	pterm.DefaultTreeList.WithItems(tis3).Render()
}
