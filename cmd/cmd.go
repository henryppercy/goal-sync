package cmd

import (
	"github.com/henryppercy/goal-sync/goals"
	"github.com/henryppercy/goal-sync/post"
	"github.com/henryppercy/goal-sync/terminal"
)

const BOOK_LIMIT = 4

func Execute() error {
	course, err := goals.GetCourse()
	if err != nil {
		return err
	}

	projects, err := goals.GetProjects()
	if err != nil {
		return err
	}

	weeks, err := goals.GetWeeks()
	if err != nil {
		return err
	}

	hours, err := goals.GetHours()
	if err != nil {
		return err
	}

	books, err := goals.GetRead(BOOK_LIMIT)
	if err != nil {
		return err
	}

	p := terminal.Programming(course, projects)
	f := terminal.Fitness(weeks)
	s := terminal.Spanish(hours)
	r := terminal.Reading(books, BOOK_LIMIT)

	terminals := post.Terminals{
		Programming: p,
		Fitness:     f,
		Spanish:     s,
		Reading:     r,
	}

	filePath := "/path/to/2026-goals.md"
	return post.Write(terminals, filePath)
}
