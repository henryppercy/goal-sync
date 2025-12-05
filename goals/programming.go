package goals

type Course struct {
	Name   string
	Module string
	Status string
}

type Project struct {
	Name string
	Date string
}

func GetCourse() Course {
	return Course{}
}

func GetProjects() []Project {
	return []Project{}
}
