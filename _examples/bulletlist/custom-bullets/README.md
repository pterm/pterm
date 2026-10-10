# bulletlist/custom-bullets

![Animation](animation.gif)

```go
package main

import "github.com/pterm/pterm"

func main() {
	glyphs := pterm.CurrentGlyphs()

	// Items can bring their own bullet and tone, which turns a list into a
	// small checklist. Glyphs come from the glyph set, so they fall back to
	// ASCII on terminals without Unicode.
	pterm.BulletList.WithItems([]pterm.BulletListItem{
		{Text: "Config file found", Bullet: glyphs.Success, Tone: pterm.ToneSuccess},
		{Text: "Cache is stale", Bullet: glyphs.Warning, Tone: pterm.ToneWarning},
		{Text: "Token expired", Bullet: glyphs.Error, Tone: pterm.ToneError},
		{Level: 1, Text: "Run `mjw login` to get a new one"},
	}).Render()
}
```
