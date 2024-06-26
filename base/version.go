package base

import (
	"fmt"
)

var (
	Version   = "(untracked)"
	CommitSHA = "(unknown)"
	BuildDate = "(unknown)"
)

func PrintVersion() {
	fmt.Printf("capybara %s (%s %s)\n", Version, CommitSHA, BuildDate)
}
