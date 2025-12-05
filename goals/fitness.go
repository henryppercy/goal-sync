package goals

import (
	"encoding/json"
	"os"
)

type FitnessWeeks struct {
	Trained []int `json:"trained"`
	Missed  []int `json:"missed"`
}

const FITNESS_FILE = "fitness.json"

func GetWeeks() (FitnessWeeks, error) {
	f, err := os.ReadFile(FITNESS_FILE)
	if err != nil {
		return FitnessWeeks{}, err
	}

	var fitness FitnessWeeks
	err = json.Unmarshal(f, &fitness)
	if err != nil {
		return FitnessWeeks{}, err
	}

	return fitness, nil
}
