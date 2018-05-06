package roller

// Sequence returns numbers from a known sequence.
// Intended for use in tests.
type Sequence struct {
	i   int
	seq []int
}

// WithSequence returns a Roller which will return numbers from a given sequence.
func WithSequence(seq []int) Roller {
	return &Sequence{-1, seq}
}

// Roll returns the next number in the sequence.
// Note that `max` is ignored for Sequences.
func (s *Sequence) Roll(max int) int {
	s.i = (s.i + 1) % len(s.seq)
	return s.seq[s.i]
}
