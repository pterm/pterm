# table/from-structs

![Animation](https://vhs.charm.sh/vhs-7JLw6Gm8B76cG5m18AysEs.gif)

```go
package main

import (
	"github.com/pterm/pterm/putils"
)

// User is a regular struct; no tags or interfaces are needed.
type User struct {
	Name  string
	Age   int
	Email string
}

func main() {
	users := []User{
		{Name: "Ada Lovelace", Age: 36, Email: "ada@example.com"},
		{Name: "Alan Turing", Age: 41, Email: "alan@example.com"},
		{Name: "Grace Hopper", Age: 85, Email: "grace@example.com"},
	}

	// DefaultTableFromStructSlice fills the default table via reflection: the
	// field names become the first row, so WithHasHeader styles them as the
	// header.
	putils.DefaultTableFromStructSlice(users).WithHasHeader().Render()
}
```
