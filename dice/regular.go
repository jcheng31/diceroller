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
func (r regular) RollN(n int) RollResults {
	result := RollResults{0, make([]int, n)}
	for i := 0; i < n; i++ {
		r := r.roller.Roll(r.max)
		result.Total += r
		result.Rolls[i] = r
	}
	return result
}
