package dice

import "github.com/jcheng31/diceroller/roller"

type exploding struct {
	roller        roller.Roller
	maxVal        int
	maxExplosions int
}

// Exploding returns a die that "explodes" - it rolls another die if it
// rolls its max value. This and all subsequent die also explode.
//
// All die explode **after** the `n` rolls have been made.
func Exploding(r roller.Roller, maxVal, maxExplosions int) Die {
	return exploding{r, maxVal, maxExplosions}
}

func (e exploding) RollN(n int) int {
	sum := 0
	toExplode := 0
	for i := 0; i < n; i++ {
		result := e.roller.Roll(e.maxVal)
		sum += result

		if result == e.maxVal {
			toExplode++
		}
	}

	for i := 0; i < toExplode && i < e.maxExplosions; i++ {
		result := e.roller.Roll(e.maxVal)
		sum += result

		if result == e.maxVal {
			toExplode++
		}
	}

	return sum
}

func (e exploding) RollNDetailed(n int) []int {
	results := make([]int, 0)

	toExplode := 0
	for i := 0; i < n; i++ {
		result := e.roller.Roll(e.maxVal)
		results = append(results, result)

		if result == e.maxVal {
			toExplode++
		}
	}

	for i := 0; i < toExplode && i < e.maxExplosions; i++ {
		result := e.roller.Roll(e.maxVal)
		results = append(results, result)

		if result == e.maxVal {
			toExplode++
		}
	}

	return results
}
