package internal

import "os"

// RunsInCi returns true if the current build is running on a CI server.
func RunsInCi() bool {
	return os.Getenv("CI") != ""
}
