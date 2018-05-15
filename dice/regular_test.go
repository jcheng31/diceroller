package dice

import (
	"testing"

	"github.com/jcheng31/diceroller/roller"
)

func Test_Regular_RollN_CorrectlySumsResults(t *testing.T) {
	seq := []int{2, 2, 2, 2, 2}

	r := roller.WithSequence(seq)

	die := Regular(r, 8)

	expected := 10
	actual := die.RollN(5)

	if expected != actual.Total {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Regular_RollN_GivesAllResults(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5}

	r := roller.WithSequence(seq)

	die := Regular(r, 8)

	results := die.RollN(5)

	if len(results.Rolls) != len(seq) {
		t.Fatalf("RollN expected %v, actual %v", seq, results.Rolls)
	}

	for i, actual := range results.Rolls {
		if seq[i] != actual {
			t.Fatalf("RollN expected %v, actual %v", seq, results.Rolls)
		}
	}
}
