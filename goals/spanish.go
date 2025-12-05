package goals

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
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
