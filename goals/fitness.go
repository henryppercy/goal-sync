package goals

import (
	"encoding/json"
	"os"
)

type FitnessWeeks struct {
	Trained []int `json:"trained"`
	Missed  []int `json:"missed"`
}

func GetWeeks() (FitnessWeeks, error) {
	data, err := os.ReadFile("data/fitness.json")
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
