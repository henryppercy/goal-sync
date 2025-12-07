package goals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"time"
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

type ProgrammingProgress struct {
	Course   Course
	Projects []Project
}

func GetProgramming() (ProgrammingProgress, error) {
	course, err := GetCourse()
	if err != nil {
		return ProgrammingProgress{}, err
	}

	projects, err := GetProjects()
	if err != nil {
		return ProgrammingProgress{}, err
	}

	return ProgrammingProgress{
		Course:   course,
		Projects: projects,
	}, nil
}

func (p ProgrammingProgress) ToTerminal() string {
	command1 := "henry@2026:~/goals/programming/course $ tail runtime.log"
	out1 := fmt.Sprintf(
		"[INFO] course loaded: %s\n[INFO] latest module: %s\n[INFO] status: %s",
		p.Course.Name,
		p.Course.Module,
		p.Course.Status,
	)

	command2 := "henry@2026:~/goals/programming/projects $ ls -l"

	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	for i, project := range p.Projects {
		err := createProjectLogLine(w, &project, len(p.Projects) == i+1)
		if err != nil {
			return ""
		}
	}

	w.Flush()

	out2 := fmt.Sprintf(
		"total %d\n%s",
		len(p.Projects),
		buf.String(),
	)

	return fmt.Sprintf(
		"%s\n%s\n\n%s\n%s",
		command1,
		out1,
		command2,
		out2,
	)
}

func createProjectLogLine(w *tabwriter.Writer, project *Project, last bool) error {
	t, err := time.Parse(time.RFC3339, project.Date)
	if err != nil {
		return err
	}

	lastChar := "\n"
	if last {
		lastChar = ""
	}

	fmt.Fprintf(
		w,
		"drwxr-xr-x\t2\thenrypercy\tstaff\t64\t%s\t%s%s",
		strings.Join((strings.Split(t.Format("2 Jan 15:04"), " ")), "\t"),
		project.Name,
		lastChar,
	)

	return nil
}
