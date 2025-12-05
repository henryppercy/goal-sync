package main

import (
	"os"

	"github.com/henryppercy/goal-sync/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}
