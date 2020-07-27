// Package checks provides functions check passed into CLI arguments.
package checks

import "fmt"

// InvalidInputError is error type if passed argument is invalid.
type InvalidInputError struct {
	Name string
}

// Error returns formatted InvalidInputErr.
func (e InvalidInputError) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Name)
}
