package terminal

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/henryppercy/goal-sync/goals"
)

type Row struct {
	prefix string
	entries string
	suffix string
}

// ```zsh
// henry@2026:~/goals/fitness $ ./calendar.sh

// Wk 01-13:  • • x • • ○ ○ ○ ○ ○ ○ ○ ○   (Q1)
// Wk 14-26:  ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○   (Q2)
// Wk 27-39:  ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○   (Q3)
// Wk 40-52:  ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○ ○   (Q4)

// • trained | ○ pending | x missed
// ```
func Fitness(weeks goals.FitnessWeeks) string {
	command := "henry@2026:~/goals/fitness $ ./calendar.sh"
	legend := "• trained | ○ pending | x missed"

	entries := generateEntries(weeks)

	const R_NUM = 4;

	symbolArr := strings.Split(entries, "")
	
	symbolCount := len(symbolArr)
	rowLength := symbolCount/R_NUM

	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	i := 0
	for i < R_NUM {
		start := i * rowLength
		end := start + rowLength

		startWeek := fmt.Sprintf("%02d", start)
		endWeek := fmt.Sprintf("%02d", end)

		createFitnessCalendarRow(w, &Row{
			prefix: fmt.Sprintf("Wk %s-%s:", startWeek, endWeek),
			entries: strings.Join(symbolArr[start:end], " "),
			suffix: "Q" + strconv.Itoa(i+1),
		})

		i++
	}

	w.Flush()
	
	return fmt.Sprintf(
		"%s\n\n%s\n%s",
		command,
		buf.String(),
		legend,
	)
}

func createFitnessCalendarRow(w *tabwriter.Writer, row *Row) {
	fmt.Fprintf(
		w,
		"%s\t%s\t%s\n",
		row.prefix,
		row.entries,
		row.suffix,
	)
}

// TODO: improve performance of below
func generateEntries(w goals.FitnessWeeks) string {
	var entries string

	i := 0
	for i < 52 {
		if slices.Contains(w.Trained, i) {
			entries += "•"
			i++

			continue
		}

		if slices.Contains(w.Missed, i) {
			entries += "x"
			i++

			continue
		}

		entries += "○"
		i++
	}

	return entries
}
