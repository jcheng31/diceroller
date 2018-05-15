package dice

import (
	"testing"

	"github.com/jcheng31/diceroller/roller"
)

func Test_Exploding_RollN_CorrectlySumsRegularResults(t *testing.T) {
	seq := []int{2, 2}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	expected := 4
	actual := die.RollN(2)

	if expected != actual {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Exploding_RollN_CorrectlySumsExplodingResults(t *testing.T) {
	seq := []int{2, 8, 3}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	expected := 13
	actual := die.RollN(2)

	if expected != actual {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Exploding_RollN_CorrectlyLimitsExplosions(t *testing.T) {
	seq := []int{8, 8, 8, 8, 8, 8, 8, 8, 8, 8}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 5)

	expected := 48
	actual := die.RollN(1)

	if expected != actual {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Exploding_RollNDetailed_CorrectlyReturnsResults(t *testing.T) {
	seq := []int{1, 2, 3, 4}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	results := die.RollNDetailed(4)
	if len(results) != len(seq) {
		t.Fatalf("RollNDetailed expected %v, actual %v", seq, results)
	}

	for i, actual := range results {
		expected := seq[i]
		if expected != actual {
			t.Fatalf("RollNDetailed expected %v, actual %v", seq, results)
		}
	}
}

func Test_Exploding_RollNDetailed_CorrectlyReturnsExplosions(t *testing.T) {
	seq := []int{4, 8, 8, 1}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	results := die.RollNDetailed(2)
	if len(results) != len(seq) {
		t.Fatalf("RollNDetailed expected %v, actual %v", seq, results)
	}

	for i, actual := range results {
		expected := seq[i]
		if expected != actual {
			t.Fatalf("RollNDetailed expected %v, actual %v", seq, results)
		}
	}
}

func Test_Exploding_RollNDetailed_CorrectlyLimitsExplosions(t *testing.T) {
	seq := []int{8, 8, 8, 8, 8, 8, 8, 8, 8}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 2)

	expected := seq[0:3]
	results := die.RollNDetailed(1)
	if len(results) != len(expected) {
		t.Fatalf("RollNDetailed expected %v, actual %v", expected, results)
	}

	for i, actual := range results {
		expected := expected[i]
		if expected != actual {
			t.Fatalf("RollNDetailed expected %v, actual %v", expected, results)
		}
	}

}
