package cmd

import (
	"github.com/henryppercy/goal-sync/goals"
	"github.com/henryppercy/goal-sync/post"
	"github.com/henryppercy/goal-sync/terminal"
)

func Execute() error {
	course := goals.GetCourse()
	projects := goals.GetProjects()
	weeks := goals.GetWeeks()
	hours := goals.GetHours()
	books := goals.GetRead()

	p := terminal.Programming(course, projects)
	f := terminal.Fitness(weeks)
	s := terminal.Spanish(hours)
	r := terminal.Reading(books)

	terminals := post.Terminals{
		Programming: p,
		Fitness:     f,
		Spanish:     s,
		Reading:     r,
	}

	filePath := "/path/to/2026-goals.md"
	return post.Write(terminals, filePath)
}
