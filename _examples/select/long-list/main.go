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
