package dice

import "github.com/jcheng31/diceroller/roller"

// Regular represents a normal die.
type Regular struct {
	Roller  roller.Roller
	Maximum int
}

// RollN returns the total result of rolling n die.
func (r Regular) RollN(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += r.Roller.Roll(r.Maximum)
	}
	return sum
}

// RollNDetailed returns the results of rolling n die.
func (r Regular) RollNDetailed(n int) []int {
	results := make([]int, n)
	for i := 0; i < n; i++ {
		results[i] = r.Roller.Roll(r.Maximum)
	}
	return results
}
