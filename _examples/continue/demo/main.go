package main

import "github.com/pterm/pterm"

func main() {
	files := []string{"config.yaml", "docker-compose.yml", ".env"}
	overwriteAll := false

	for _, file := range files {
		if overwriteAll {
			pterm.Success.Printfln("Overwrote %s", file)
			continue
		}

		// Typing the first letter of an answer picks it, enter takes the
		// default. "all" lets the loop skip the remaining questions.
		answer, _ := pterm.Continue.Show("Overwrite " + file + "?")

		switch answer {
		case "yes":
			pterm.Success.Printfln("Overwrote %s", file)
		case "all":
			overwriteAll = true

			pterm.Success.Printfln("Overwrote %s", file)
		case "cancel":
			return
		}
	}
}
