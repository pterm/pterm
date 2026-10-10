### progressbar/demo

![Animation](demo/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	packages := []string{
		"react", "react-dom", "typescript", "vite", "eslint", "prettier",
		"vitest", "zod", "date-fns", "clsx", "tailwindcss", "postcss",
	}

	bar, _ := pterm.Progressbar.WithTotal(len(packages)).Start("Installing packages")

	for i, pkg := range packages {
		bar.UpdateTitle("Installing " + pkg)

		// One package takes a while, which the ETA picks up.
		if i == 6 {
			time.Sleep(1500 * time.Millisecond)
		}

		time.Sleep(300 * time.Millisecond)

		// Output printed while the bar runs appears above it.
		pterm.Success.Println("Installed " + pkg)

		// The bar stops on its own at its total and keeps its title in the
		// status line it leaves behind, so name the result before the last
		// step.
		if i == len(packages)-1 {
			bar.UpdateTitle("Installed packages")
		}

		bar.Increment()
	}
}
```

</details>

### progressbar/cancel

![Animation](cancel/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"context"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// The upload takes longer than the timeout allows, so the bar ends as
	// canceled when the context runs out.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	bar, _ := pterm.Progressbar.WithTotal(50).WithContext(ctx).Start("Uploading backup")

	// Stop after the context ended prints the canceled line before main
	// returns.
	defer bar.Stop()

	for range 50 {
		select {
		case <-ctx.Done():
			return
		case <-time.After(100 * time.Millisecond):
			bar.Increment()
		}
	}
}
```

</details>

### progressbar/custom

![Animation](custom/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"time"

	"github.com/pterm/pterm"
)

func main() {
	// Every part of the bar can be changed: here a green fill, a wider bar,
	// and no ETA because the work comes in uneven chunks.
	bar, _ := pterm.Progressbar.
		WithTotal(200).
		WithTone(pterm.ToneSuccess).
		WithBarWidth(40).
		WithShowETA(false).
		Start("Migrating rows")

	for range 20 {
		time.Sleep(100 * time.Millisecond)
		bar.Add(10)
	}

	// Without title, count and timer, only the bar and the percentage are
	// left.
	bar, _ = pterm.Progressbar.
		WithShowTitle(false).
		WithShowCount(false).
		WithShowElapsedTime(false).
		WithShowETA(false).
		Start()

	for range 100 {
		time.Sleep(20 * time.Millisecond)
		bar.Increment()
	}
}
```

</details>

### progressbar/multiple

![Animation](multiple/animation.gif)

<details>

<summary>SHOW SOURCE</summary>

```go
package main

import (
	"sync"
	"time"

	"github.com/pterm/pterm"
)

func main() {
	downloads := []struct {
		name   string
		chunks int
	}{
		{"ubuntu-24.04.iso", 60},
		{"node-v22.tar.gz", 25},
		{"go1.27.zip", 40},
	}

	// Bars that run at the same time stack below each other on their own.
	var wg sync.WaitGroup

	for _, d := range downloads {
		bar, _ := pterm.Progressbar.WithTotal(d.chunks).Start(d.name)

		wg.Go(func() {
			for range d.chunks {
				time.Sleep(50 * time.Millisecond)
				bar.Increment()
			}
		})
	}

	wg.Wait()
}
```

</details>

