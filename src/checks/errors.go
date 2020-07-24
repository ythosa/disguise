package checks

import "fmt"

type InvalidInputError struct {
	Name string
}

func (e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Name)
}
