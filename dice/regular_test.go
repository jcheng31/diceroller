package dice

import (
	"testing"

	"github.com/jcheng31/diceroller/roller"
)

func Test_Regular_RollN_CorrectlySumsResults(t *testing.T) {
	seq := []int{2, 2, 2, 2, 2}

	r := roller.WithSequence(seq)

	die := Regular{r, 8}

	expected := 10
	actual := die.RollN(5)

	if expected != actual {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Regular_RollNDetailed_GivesAllResults(t *testing.T) {
	seq := []int{1, 2, 3, 4, 5}

	r := roller.WithSequence(seq)

	die := Regular{r, 8}

	results := die.RollNDetailed(5)

	for i, actual := range results {
		expected := seq[i]
		if expected != actual {
			t.Errorf("RollNDetailed expected %v, actual %v", expected, actual)
		}
	}
}
