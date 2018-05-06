package roller

import (
	"testing"
)

func Test_Sequence_Roll_ReturnsSequence(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5}

	roller := WithSequence(seq)

	for _, expected := range seq {
		actual := roller.Roll(5)
		if expected != actual {
			t.Errorf("Sequence.Roll expected %v, got %v", expected, actual)
		}
	}
}
