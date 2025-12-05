package terminal

import (
	"fmt"
	"strings"
)

// henry@2026:~/goals/spanish $ ./progress.sh
// [################........................] 354/1000 hrs - 34%
func Spanish(hours int) string {
	const goal = 1000
	const barWidth = 40

	bar := buildProgressBar(hours, goal, barWidth)

	percentage := hours * 100 / goal

	return fmt.Sprintf(
		"henry@2026:~/goals/spanish $ ./progress.sh\n%s %d/%d hrs - %d%%",
		bar,
		hours,
		goal,
		percentage,
	)
}

func buildProgressBar(current, total, width int) string {
	filled := (width * current) / total
	empty := width - filled

	return "[" + strings.Repeat("#", filled) + strings.Repeat(".", empty) + "]"
}
