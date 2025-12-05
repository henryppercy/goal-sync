package goals

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type Book struct {
	Title       string
	Authors     []string
	Date        string
	DaysElapsed string
	Rating      string
}

type bookFrontmatter struct {
	Title        string   `yaml:"title"`
	Authors      []string `yaml:"authors"`
	DateStarted  string   `yaml:"date_started"`
	DateFinished string   `yaml:"date_finished"`
	Rating       int      `yaml:"rating"`
}

func GetRead() ([]Book, error) {
	files, err := filepath.Glob("book/*.mdx")
	if err != nil {
		return []Book{}, err
	}

	var books []Book

	for _, file := range files {
		book, err := parseBookFile(file)
		if err != nil {
			continue
		}
		books = append(books, book)
	}

	return books, nil
}

func parseBookFile(filePath string) (Book, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return Book{}, err
	}

	parts := strings.SplitN(string(content), "---", 3)
	if len(parts) < 3 {
		return Book{}, fmt.Errorf("invalid frontmatter format")
	}

	var fm bookFrontmatter
	err = yaml.Unmarshal([]byte(parts[1]), &fm)
	if err != nil {
		return Book{}, err
	}

	daysElapsed := ""
	if fm.DateStarted != "" && fm.DateFinished != "" {
		started, err1 := time.Parse("2006-01-02", fm.DateStarted)
		finished, err2 := time.Parse("2006-01-02", fm.DateFinished)
		if err1 == nil && err2 == nil {
			days := int(finished.Sub(started).Hours() / 24)
			daysElapsed = fmt.Sprintf("%dd", days)
		}
	}

	date := fm.DateFinished
	if date == "" {
		date = fm.DateStarted
	}

	rating := ""
	if fm.Rating > 0 {
		rating = fmt.Sprintf("%d", fm.Rating)
	} else {
		rating = "-"
	}

	return Book{
		Title:       fm.Title,
		Authors:     fm.Authors,
		Date:        date,
		DaysElapsed: daysElapsed,
		Rating:      rating,
	}, nil
}
