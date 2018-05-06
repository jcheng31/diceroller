// Package roller provides functionality to return numbers.
package roller

// Roller provides a function to return numbers.
type Roller interface {
	// Roll returns a number less than or equal to `max`.
	Roll(max int) int
}
