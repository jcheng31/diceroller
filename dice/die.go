// Package dice provides implementations of several dice types.
package dice

// Die represents a type of die.
type Die interface {
	// RollN rolls `n` dice and returns the total.
	RollN(n int) int

	// RollNDetailed rolls `n` dice and returns the results of each.
	RollNDetailed(n int) []int
}
