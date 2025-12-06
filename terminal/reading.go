package terminal

import (
	"bytes"
	"fmt"
	"text/tabwriter"
	"time"

	"github.com/henryppercy/goal-sync/goals"
)

// henry@2026:~/goals/reading $ wc -l books_completed.txt
// 3 books_completed.txt
//
// henry@2026:~/goals/reading $ tail -n 4 reading_log.txt
// [2025-09-11][DONE]  The Cartel      Don Winslow      21d  5
// [2025-10-30][DONE]  Blood Meridian  Cormac McCarthy  45d  2
// [2025-11-11][DONE]  Animal Farm     George Orwell    9d   3
// [2025-12-06][OPEN]  The Border      Don Winslow      25d  -
func Reading(books []goals.Book, length int) string {
	var completed int

	for _, book := range books {
		if book.Date != "" {
			completed++
		}
	}

	command1 := "henry@2026:~/goals/reading $ wc -l books_completed.txt"
	out1 := fmt.Sprintf("%d books_completed.txt", completed)
	command2 := fmt.Sprintf("henry@2026:~/goals/reading $ tail -n %d reading_log.txt", length)

	logBooks := books[len(books)-length:]

	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	for _, book := range logBooks {
		createLogLine(w, &book)
	}

	w.Flush()

	return fmt.Sprintf("%s\n%s\n\n%s\n%s", command1, out1, command2, buf.String())
}

func createLogLine(w *tabwriter.Writer, book *goals.Book) {
	status := "DONE"
	date := book.Date

	if book.Date == "" {
		status = "OPEN"

		now := time.Now()
		date = now.Format("2006-01-02")
	}

	fmt.Fprintf(
		w,
		"[%s][%s]\t%s\t%s\t%s\t%s\n",
		date,
		status,
		book.Title,
		book.Authors[0],
		book.DaysElapsed,
		book.Rating,
	)
}
