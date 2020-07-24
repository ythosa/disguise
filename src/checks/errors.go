package checks

import "fmt"

// invalidInputError is error type if passed argument is invalid
type invalidInputError struct {
	Name string
}

// Error returns formatted InvalidInputErr
func (e invalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Name)
}
