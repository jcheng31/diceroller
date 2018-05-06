package dice

import "github.com/jcheng31/diceroller/roller"

type regular struct {
	roller roller.Roller
	max    int
}

// Regular returns a normal die.
func Regular(r roller.Roller, max int) Die {
	return regular{r, max}
}

// RollN returns the total result of rolling n die.
func (r regular) RollN(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += r.roller.Roll(r.max)
	}
	return sum
}

// RollNDetailed returns the results of rolling n die.
func (r regular) RollNDetailed(n int) []int {
	results := make([]int, n)
	for i := 0; i < n; i++ {
		results[i] = r.roller.Roll(r.max)
	}
	return results
}
