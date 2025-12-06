package cmd

import (
	"fmt"

	"github.com/henryppercy/goal-sync/goals"
	"github.com/henryppercy/goal-sync/post"
)

const BOOK_LIMIT = 4

func Execute() error {
	programming, err := goals.GetProgramming()
	if err != nil {
		return err
	}

	fitness, err := goals.GetWeeks()
	if err != nil {
		return err
	}

	spanish, err := goals.GetSpanish()
	if err != nil {
		return err
	}

	reading, err := goals.GetReading(BOOK_LIMIT)
	if err != nil {
		return err
	}

	t := post.Terminals{
		Programming: programming.ToTerminal(),
		Fitness:     fitness.ToTerminal(),
		Spanish:     spanish.ToTerminal(),
		Reading:     reading.ToTerminal(),
	}

	fmt.Println(t.String())

	filePath := "/path/to/2026-goals.md"
	return t.Write(filePath)
}
