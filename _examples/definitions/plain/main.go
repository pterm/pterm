package main

import "github.com/pterm/pterm"

func main() {
	// Without groups there is no heading and no indent, which suits short
	// reference lists such as environment variables.
	pterm.Definitions.WithDefinitions(
		pterm.Def("PTERM_THEME", "macchiato, latte, ansi or auto"),
		pterm.Def("PTERM_ASCII", "use ASCII glyphs"),
		pterm.Def("PTERM_THEME_<SLOT>", "override one color of the theme"),
		pterm.Def("NO_COLOR", "turn colors off"),
	).Render()
}
