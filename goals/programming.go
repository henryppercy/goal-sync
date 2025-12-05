package goals

import (
	"encoding/json"
	"os"
)

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
	data, err := os.ReadFile("data/course.json")
	if err != nil {
		return Course{}, err
	}

	var course Course
	err = json.Unmarshal(data, &course)
	if err != nil {
		return Course{}, err
	}

	return course, nil
}

func GetProjects() ([]Project, error) {
	data, err := os.ReadFile("data/projects.json")
	if err != nil {
		return []Project{}, err
	}

	var projects []Project
	err = json.Unmarshal(data, &projects)
	if err != nil {
		return []Project{}, err
	}

	return projects, nil
}
