package goals

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"
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

func GetRead(path string, length int) ([]Book, error) {
	files, err := filepath.Glob(path)
	if err != nil {
		return []Book{}, err
	}

	var books []Book

	for _, file := range files {
		book, err := parseBookFile(file)
		if err != nil {
			continue
		}

		if book.Date == "" {
			books = append(books, book)
			continue
		}

		t, err := time.Parse("2006-01-02", book.Date)
		if err != nil {
			return []Book{}, err
		}

		if t.Year() == 2026 {
			books = append(books, book)
		}
	}

	sort.Slice(books, func(i, j int) bool {
		if books[i].Date == "" && books[j].Date == "" {
			return false
		}
		if books[i].Date == "" {
			return false
		}
		if books[j].Date == "" {
			return true
		}
		return books[i].Date < books[j].Date
	})

	return books, nil
}

type ReadingProgress struct {
	Books  []Book
	Length int
}

func GetReading(path string, length int) (ReadingProgress, error) {
	books, err := GetRead(path, length)
	if err != nil {
		return ReadingProgress{}, err
	}

	return ReadingProgress{
		Books:  books,
		Length: length,
	}, nil
}

func (r ReadingProgress) ToTerminal() string {
	var completed int

	for _, book := range r.Books {
		if book.Date != "" {
			completed++
		}
	}

	command1 := "henry@2026:~/goals/reading $ wc -l books_completed.txt"
	out1 := fmt.Sprintf("%d books_completed.txt", completed)
	command2 := fmt.Sprintf("henry@2026:~/goals/reading $ tail -n %d reading_log.txt", r.Length)

	if len(r.Books) < r.Length {
		return fmt.Sprintf("%s\n%s", command1, out1)
	}

	logBooks := r.Books[len(r.Books)-r.Length:]

	var buf bytes.Buffer

	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0)

	for i, book := range logBooks {
		createBookLogLine(w, book, len(logBooks) == i+1)
	}

	w.Flush()

	return fmt.Sprintf("%s\n%s\n\n%s\n%s", command1, out1, command2, buf.String())
}

func createBookLogLine(w *tabwriter.Writer, book Book, last bool) {
	status := "DONE"
	date := book.Date

	if book.Date == "" {
		status = "OPEN"

		now := time.Now()
		date = now.Format("2006-01-02")
	}

	lastChar := "\n"
	if last {
		lastChar = ""
	}

	fmt.Fprintf(
		w,
		"[%s][%s]\t%s\t%s\t%s\t%s%s",
		date,
		status,
		book.Title,
		book.Authors[0],
		book.DaysElapsed,
		book.Rating,
		lastChar,
	)
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
	if fm.DateStarted != "" {
		now := time.Now()
		endDate := fm.DateFinished
		if endDate == "" {
			endDate = now.Format("2006-01-02")
		}

		started, err1 := time.Parse("2006-01-02", fm.DateStarted)
		finished, err2 := time.Parse("2006-01-02", endDate)

		if err1 == nil && err2 == nil {
			days := int(finished.Sub(started).Hours() / 24)
			daysElapsed = fmt.Sprintf("%dd", days)
		}
	}

	date := fm.DateFinished

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
