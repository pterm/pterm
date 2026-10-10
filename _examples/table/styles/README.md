# table/styles

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Styles use theme tones, so the table follows the active theme. Here the
	// status column is toned and the failed row stands out.
	pterm.Table.
		WithData(pterm.TableData{
			{"Job", "Status", "Duration"},
			{"lint", "passed", "12.4s"},
			{"test", "passed", "1m 3s"},
			{"build", "failed", "48.0s"},
			{"deploy", "skipped", "-"},
		}).
		WithColumnStyle(1, pterm.ToneAccent.Style()).
		WithRowStyle(2, pterm.ToneError.Style()).
		Render()

	pterm.Println()

	// An alternate row style makes long rows easier to follow.
	pterm.Table.
		WithData(pterm.TableData{
			{"Firstname", "Lastname", "Email"},
			{"Paul", "Dean", "paul@example.com"},
			{"Callie", "Mckay", "callie@example.com"},
			{"Libby", "Camacho", "libby@example.com"},
			{"张", "小宝", "zhang@example.com"},
		}).
		WithAlternateRowStyle(pterm.ToneMuted.Style()).
		Render()
}
```
