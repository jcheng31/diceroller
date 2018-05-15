// Package dice provides implementations of several dice types.
package dice

// Die represents a type of die.
type Die interface {
	// RollN rolls `n` dice and returns the results.
	RollN(n int) RollResults
}

// RollResults provides the result of rolling some dice.
type RollResults struct {
	Total int
	Rolls []int
}
