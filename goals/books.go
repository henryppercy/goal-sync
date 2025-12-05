package goals

type Book struct {
	Title       string
	Authors     []string
	Date        string
	DaysElapsed string
	Rating      string
}

func GetRead() []Book {
	return []Book{}
}
