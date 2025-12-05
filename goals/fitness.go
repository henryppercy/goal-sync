package goals

import (
	_ "embed"
	"encoding/json"
)

//go:embed data/fitness.json
var fitnessJSON []byte

type FitnessWeeks struct {
	Trained []int `json:"trained"`
	Missed  []int `json:"missed"`
}

func GetWeeks() (FitnessWeeks, error) {
	var fitness FitnessWeeks
	err := json.Unmarshal(fitnessJSON, &fitness)
	if err != nil {
		return FitnessWeeks{}, err
	}

	return fitness, nil
}
