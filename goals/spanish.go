package goals

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const URL = "https://www.dreamingspanish.com/.netlify/functions/dayWatchedTime?language=es"

type Log struct {
	Date        string  `json:"date"`
	TimeSeconds float64 `json:"timeSeconds"`
	GoalReached bool    `json:"goalReached"`
}

func GetHours() (int, error) {
	logs, err := getWatchLogs()
	if err != nil {
		return 0, err
	}

	var seconds float64

	for _, log := range logs {
		seconds += log.TimeSeconds
	}

	return int(seconds / 3600), err
}

type SpanishProgress struct {
	Hours int
}

func GetSpanish() (SpanishProgress, error) {
	hours, err := GetHours()
	if err != nil {
		return SpanishProgress{}, err
	}

	return SpanishProgress{Hours: hours}, nil
}

func (s SpanishProgress) ToTerminal() string {
	const goal = 1000
	const barWidth = 40

	bar := buildProgressBar(s.Hours, goal, barWidth)

	percentage := s.Hours * 100 / goal

	return fmt.Sprintf(
		"henry@2026:~/goals/spanish $ ./progress.sh\n%s %d/%d hrs - %d%%",
		bar,
		s.Hours,
		goal,
		percentage,
	)
}

func buildProgressBar(current, total, width int) string {
	filled := (width * current) / total
	empty := width - filled

	return "[" + strings.Repeat("#", filled) + strings.Repeat(".", empty) + "]"
}

func getWatchLogs() ([]Log, error) {
	token := os.Getenv("DS_TOKEN")
	if token == "" {
		return nil, errors.New("dreaming Spanish API bearer token not set")
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var logs []Log
	err = json.Unmarshal(body, &logs)
	if err != nil {
		return nil, err
	}

	return logs, nil
}
