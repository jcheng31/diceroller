package dice

import "github.com/jcheng31/diceroller/roller"

type exploding struct {
	roller        roller.Roller
	maxVal        int
	maxExplosions int
}

// Exploding returns a Die that explodes - it rolls another die if it
// rolls its max value. This and all subsequent die also explode.
//
// All die explode **after** the `n` rolls have been made, and all subsequent
// explosions occur after all initial ones have completed.
func Exploding(r roller.Roller, maxVal, maxExplosions int) Die {
	return exploding{r, maxVal, maxExplosions}
}

func (e exploding) RollN(n int) RollResults {
	result := RollResults{0, make([]int, 0)}

	toExplode := 0
	for i := 0; i < n; i++ {
		r := e.roller.Roll(e.maxVal)
		result.Total += r
		result.Rolls = append(result.Rolls, r)

		if r == e.maxVal {
			toExplode++
		}
	}

	for i := 0; i < toExplode && i < e.maxExplosions; i++ {
		r := e.roller.Roll(e.maxVal)
		result.Total += r
		result.Rolls = append(result.Rolls, r)

		if r == e.maxVal {
			toExplode++
		}
	}

	return result
}
