package goals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"text/tabwriter"
)

type FitnessWeeks struct {
	Trained []int `json:"trained"`
	Missed  []int `json:"missed"`
}

func GetWeeks(path string) (FitnessWeeks, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return FitnessWeeks{}, err
	}

	var fitness FitnessWeeks
	err = json.Unmarshal(data, &fitness)
	if err != nil {
		return FitnessWeeks{}, err
	}

	return fitness, nil
}

func (w FitnessWeeks) ToTerminal() string {
	command := "henry@2026:~/goals/fitness $ ./calendar.sh"
	legend := "• trained | ○ pending | x missed"

	entries := w.generateEntries()

	const R_NUM = 4

	symbolArr := strings.Split(entries, "")

	symbolCount := len(symbolArr)
	rowLength := symbolCount / R_NUM

	var buf bytes.Buffer

	writer := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	i := 0
	for i < R_NUM {
		start := i * rowLength
		end := start + rowLength

		startWeek := fmt.Sprintf("%02d", start)
		endWeek := fmt.Sprintf("%02d", end)

		createFitnessCalendarRow(writer, fitnessRow{
			prefix:  fmt.Sprintf("Wk %s-%s:", startWeek, endWeek),
			entries: strings.Join(symbolArr[start:end], " "),
			suffix:  "Q" + strconv.Itoa(i+1),
		})

		i++
	}

	writer.Flush()

	return fmt.Sprintf(
		"%s\n\n%s\n%s",
		command,
		buf.String(),
		legend,
	)
}

type fitnessRow struct {
	prefix  string
	entries string
	suffix  string
}

func createFitnessCalendarRow(w *tabwriter.Writer, row fitnessRow) {
	fmt.Fprintf(
		w,
		"%s\t%s\t%s\n",
		row.prefix,
		row.entries,
		row.suffix,
	)
}

// TODO: improve performance
func (w FitnessWeeks) generateEntries() string {
	var entries string

	i := 1
	for i <= 52 {
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
