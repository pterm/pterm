//go:build !windows

package color

// consoleSupport reports the color support guaranteed by the platform console
// itself. Outside of Windows the console needs no special setup and gives no
// verdict, so color detection is purely environment based.
func consoleSupport() (Level, bool) {
	return LevelNone, false
}
