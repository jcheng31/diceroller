package roller

import "math/rand"

// Random returns random numbers.
type Random struct {
	r *rand.Rand
}

// WithRandomSource returns a Roller which returns random numbers,
// using a rand.Source as the source.
func WithRandomSource(src rand.Source) Roller {
	return &Random{rand.New(src)}
}

// Roll returns a random number in [1, `max`).
func (r *Random) Roll(max int) int {
	return r.r.Intn(max) + 1
}
