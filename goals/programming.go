package goals

import (
	_ "embed"
	"encoding/json"
)

//go:embed data/course.json
var courseJSON []byte

//go:embed data/projects.json
var projectsJSON []byte

type Course struct {
	Name   string `json:"name"`
	Module string `json:"current_module"`
	Status string `json:"status"`
}

type Project struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

func GetCourse() (Course, error) {
	var course Course
	err := json.Unmarshal(courseJSON, &course)
	if err != nil {
		return Course{}, err
	}

	return course, nil
}

func GetProjects() ([]Project, error) {
	var projects []Project
	err := json.Unmarshal(projectsJSON, &projects)
	if err != nil {
		return []Project{}, err
	}

	return projects, nil
}
