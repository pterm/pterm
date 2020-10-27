// Package pterm is a modern go module to beautify console output.
// It can be used without configuration, but if desired, everything can be customized down to the smallest detail.
// View the animated examples here: https://github.com/pterm/pterm#-examples
package pterm

var (
	// DisableOutput completely disables output from pterm. Can be used in CLI application quiet mode.
	DisableOutput = false
	// PrintDebugMessages sets if messages printed by the DebugPrinter should be printed.
	PrintDebugMessages = false
)

// EnableDebugMessages enables the output of DebugPrinter.
func EnableDebugMessages() {
	PrintDebugMessages = true
}

// DisableDebugMessages disables the output of DebugPrinter.
func DisableDebugMessages() {
	PrintDebugMessages = false
}
