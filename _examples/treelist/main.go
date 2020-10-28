package main

import (
	"github.com/pterm/pterm"
)

func main() {
	var tis []pterm.TreeListItem

	tis = []pterm.TreeListItem{
		{Children: []pterm.TreeListItem{
			{Children: []pterm.TreeListItem{
				{Children: []pterm.TreeListItem{
					{Children: nil, ItemName: "2"},
					{Children: nil, ItemName: "3"},
					{Children: nil, ItemName: "4"},
				}, ItemName: "2"},
				{Children: nil, ItemName: "3"},
				{Children: nil, ItemName: "4"},
			}, ItemName: "2"},
			{Children: nil, ItemName: "3"},
			{Children: nil, ItemName: "4"},
		}, ItemName: "0"},
		{Children: nil, ItemName: "1"},
		{Children: []pterm.TreeListItem{
			{Children: []pterm.TreeListItem{
				{Children: []pterm.TreeListItem{
					{Children: []pterm.TreeListItem{
						{Children: []pterm.TreeListItem{
							{Children: nil, ItemName: "2"},
							{Children: nil, ItemName: "3"},
							{Children: nil, ItemName: "4"},
						}, ItemName: "2"},
						{Children: nil, ItemName: "3"},
						{Children: nil, ItemName: "4"},
					}, ItemName: "2"},
					{Children: nil, ItemName: "3"},
					{Children: nil, ItemName: "4"},
				}, ItemName: "0"},
			}, ItemName: "2"},
			{Children: nil, ItemName: "3"},
			{Children: nil, ItemName: "4"},
		}, ItemName: "2"},
		{Children: nil, ItemName: "3"},
	}

	pterm.DefaultTreeList.WithItems(tis).Render()
}
