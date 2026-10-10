### sparkline/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// One column per value, scaled between the smallest and the largest.
	pterm.Sparkline.
		WithValues(3, 5, 4, 8, 12, 9, 14, 18, 16, 21, 19, 24, 22, 17, 13, 15).
		WithLabel("requests per second").
		Render()
}
```

</details>

### sparkline/thresholds

![Animation](thresholds/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import "github.com/pterm/pterm"

func main() {
	// Load per CPU core. A fixed range of 0 to 100 makes the thresholds
	// mean percent, and each column is toned by its own level.
	pterm.Sparkline.
		WithValues(12, 34, 88, 56, 71, 23, 95, 41, 8, 63, 77, 19, 50, 82, 30, 67).
		WithRange(0, 100).
		WithThresholds(pterm.LoadThresholds()...).
		WithLabel("16 cores").
		Render()
}
```

</details>

