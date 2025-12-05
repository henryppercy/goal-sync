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

const COURSE_FILE = "data/course.json"
const PROJECTS_FILE = "data/projects.json"

func GetCourse() (Course, error) {
	f, err := os.ReadFile(COURSE_FILE)
	if err != nil {
		return Course{}, err
	}

	var course Course
	err = json.Unmarshal(f, &course)
	if err != nil {
		return Course{}, err
	}

	return course, nil
}

func GetProjects() ([]Project, error) {
	f, err := os.ReadFile(PROJECTS_FILE)
	if err != nil {
		return []Project{}, err
	}

	var projects []Project
	err = json.Unmarshal(f, &projects)
	if err != nil {
		return []Project{}, err
	}

	return projects, nil
}
