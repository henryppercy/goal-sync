package post

import "fmt"

type Terminals struct {
	Programming string
	Fitness     string
	Spanish     string
	Reading     string
}

func (t Terminals) Write(filePath string) error {
	return nil
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
