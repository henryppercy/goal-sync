package post

import "fmt"

type Terminals struct {
	Programming []byte
	Fitness     []byte
	Spanish     []byte
	Reading     []byte
}

func Write(terminals Terminals, filePath string) error {
	fmt.Printf("Mock: Would write to %s\n", filePath)
	fmt.Printf("  - Programming: %d bytes\n", len(terminals.Programming))
	fmt.Printf("  - Fitness: %d bytes\n", len(terminals.Fitness))
	fmt.Printf("  - Spanish: %d bytes\n", len(terminals.Spanish))
	fmt.Printf("  - Reading: %d bytes\n", len(terminals.Reading))

	return nil
}
