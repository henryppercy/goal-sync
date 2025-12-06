package terminal

import (
	"bytes"
	"fmt"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/henryppercy/goal-sync/goals"
)

// ```zsh
// henry@2026:~/goals/programming/course $ tail runtime.log
// [INFO] course loaded: command-line-applications-in-go
// [INFO] latest module: 5.0 concurrency and streams
// [INFO] status: progress

// henry@2026:~/goals/programming/projects $ ls -l
// total 1
// drwxr-xr-x  2 henrypercy  staff   64 02 Mar 12:30 progress-cli
// ```
func Programming(course goals.Course, projects []goals.Project) string {
	command1 := "henry@2026:~/goals/programming/course $ tail runtime.log"
	out1 := fmt.Sprintf(
		"[INFO] course loaded: %s\n[INFO] latest module: %s\n[INFO] status: %s",
		course.Name,
		course.Module,
		course.Status,
	)

	command2 := "henry@2026:~/goals/programming/projects $ ls -l"

	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	for _, project := range projects {
		err := createProjectLogLine(w, &project)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}

	w.Flush()

	out2 := fmt.Sprintf(
		"total %d\n%s",
		len(projects),
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

func createProjectLogLine(w *tabwriter.Writer, project *goals.Project) error {
	t, err := time.Parse(time.RFC3339, project.Date)
	if err != nil {
		return err
	}

	fmt.Fprintf(
		w,
		"drwxr-xr-x\t2\thenrypercy\tstaff\t64\t%s\t%s\n",
		strings.Join((strings.Split(t.Format("2 Jan 15:04"), " ")), "\t"),
		project.Name,
	)

	return nil
}
