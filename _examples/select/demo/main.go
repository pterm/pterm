package main

import "github.com/pterm/pterm"

func main() {
	// Options can carry a description and a right-aligned detail. Typing
	// filters the list with a fuzzy match, so "k8s us" finds kapp-k8s-us.
	// Disabled options stay visible but the cursor skips them.
	clusters := []pterm.Option{
		{Label: "kapp-k8s-at", Description: "Vienna, production", Detail: "3 nodes"},
		{Label: "kapp-k8s-de", Description: "Frankfurt, production", Detail: "5 nodes"},
		{Label: "kapp-k8s-us", Description: "Virginia, production", Detail: "12 nodes"},
		{Label: "staging", Value: "stg", Description: "shared playground", Detail: "2 nodes"},
		{Label: "legacy-eu", Description: "read only until March", Detail: "1 node", Disabled: true},
	}

	// Show returns the Value of the chosen option, or its Label when the
	// option has no Value.
	cluster, _ := pterm.Select.WithItems(clusters).Show("Which cluster?")

	pterm.Info.Printfln("Deploying to %s", cluster)
}
