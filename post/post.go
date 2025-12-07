package post

import (
	"fmt"
	"os"
	"strings"
)

type Terminals struct {
	Programming string
	Fitness     string
	Spanish     string
	Reading     string
}

func (t Terminals) Write(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	newContent := string(content)
	newContent = replaceSection(newContent, "Programming", t.Programming)
	newContent = replaceSection(newContent, "Fitness", t.Fitness)
	newContent = replaceSection(newContent, "Spanish", t.Spanish)
	newContent = replaceSection(newContent, "Reading", t.Reading)

	return os.WriteFile(fileName, []byte(newContent), 0644)
}

func replaceSection(content, title, terminal string) string {
	titleStr := "### " + title

	titleIdx := strings.Index(content, titleStr)
	if titleIdx == -1 {
		return content
	}

	afterTitle := content[titleIdx:]
	codeStartRelative := strings.Index(afterTitle, "```zsh")
	if codeStartRelative == -1 {
		return content
	}
	codeStartAbsolute := titleIdx + codeStartRelative

	afterCodeStart := content[codeStartAbsolute+6:]
	codeEndRelative := strings.Index(afterCodeStart, "```")
	if codeEndRelative == -1 {
		return content
	}
	codeEndAbsolute := codeStartAbsolute + 6 + codeEndRelative

	before := content[:codeStartAbsolute+6]
	after := content[codeEndAbsolute:]

	return before + "\n" + terminal + "\n" + after
}

func (t Terminals) String() string {
	return fmt.Sprintf(
		"%s\n%s\n\n%s\n\n%s\n",
		t.Programming,
		t.Fitness,
		t.Spanish,
		t.Reading,
	)
}
