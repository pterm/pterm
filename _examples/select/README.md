### select/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
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
```

</details>

### select/long-list

![Animation](long-list/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	zones := []string{
		"Africa/Cairo", "Africa/Lagos", "America/Chicago", "America/Los_Angeles",
		"America/New_York", "America/Sao_Paulo", "Asia/Dubai", "Asia/Kolkata",
		"Asia/Shanghai", "Asia/Singapore", "Asia/Tokyo", "Australia/Sydney",
		"Europe/Berlin", "Europe/London", "Europe/Paris", "Europe/Vienna",
		"Pacific/Auckland", "UTC",
	}

	// Long lists scroll and show how many options are hidden above and
	// below. Typing narrows the list down, which beats scrolling.
	zone, _ := pterm.Select.
		WithOptions(zones).
		WithMaxHeight(6).
		Show("Time zone")

	pterm.Info.Printfln("Using %s", zone)
}
```

</details>

### select/options

![Animation](options/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// A plain list of strings is all a select needs. The default option is
	// where the cursor starts, and what Show returns without asking when
	// the program does not run in a terminal.
	language, _ := pterm.Select.
		WithOptions([]string{"Go", "Rust", "TypeScript", "Python", "Zig"}).
		WithDefaultOption("Go").
		Show("Favorite language?")

	pterm.Info.Printfln("You picked %s", language)
}
```

</details>

