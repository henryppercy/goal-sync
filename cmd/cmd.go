package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/henryppercy/goal-sync/goals"
	"github.com/henryppercy/goal-sync/post"
)

const BOOK_LIMIT = 4

func Execute() error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	programming, err := goals.GetProgramming(config.Paths.Course, config.Paths.Projects)
	if err != nil {
		return err
	}

	fitness, err := goals.GetWeeks(config.Paths.Fitness)
	if err != nil {
		return err
	}

	spanish, err := goals.GetSpanish()
	if err != nil {
		return err
	}

	reading, err := goals.GetReading(config.Paths.Books, config.BookLimit)
	if err != nil {
		return err
	}

	t := post.Terminals{
		Programming: programming.ToTerminal(),
		Fitness:     fitness.ToTerminal(),
		Spanish:     spanish.ToTerminal(),
		Reading:     reading.ToTerminal(),
	}

	return t.Write(config.Paths.Post)
}

type Paths struct {
	Course   string `json:"course"`
	Projects string `json:"projects"`
	Fitness  string `json:"fitness"`
	Books    string `json:"books"`
	Post     string `json:"post"`
}

type Config struct {
	Paths     Paths `json:"paths"`
	BookLimit int   `json:"book_limit"`
}

func loadConfig() (Config, error) {
    ex, err := os.Executable()
    if err != nil {
        return Config{}, err
    }
    exDir := filepath.Dir(ex)

    configPath := filepath.Join(exDir, "config.json")
    
    data, err := os.ReadFile(configPath)
    if err != nil {
        return Config{}, fmt.Errorf("config not found at %s: %w", configPath, err)
    }
    
    var config Config
    err = json.Unmarshal(data, &config)
    return config, err
}
