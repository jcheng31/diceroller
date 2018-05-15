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

	if expected != actual.Total {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Exploding_RollN_CorrectlySumsExplodingResults(t *testing.T) {
	seq := []int{2, 8, 3}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	expected := 13
	actual := die.RollN(2)

	if expected != actual.Total {
		t.Errorf("RollN expected %v, actual %v", expected, actual)
	}
}

func Test_Exploding_RollN_CorrectlyLimitsExplosions(t *testing.T) {
	seq := []int{8, 8, 8, 8, 8, 8, 8, 8, 8, 8}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 5)

	expected := seq[0:6]
	result := die.RollN(1)

	if len(expected) != len(result.Rolls) {
		t.Fatalf("RollN expected %v, actual %v", expected, result.Rolls)
	}

	for i, actual := range result.Rolls {
		if expected[i] != actual {
			t.Fatalf("RollN expected %v, actual %v", expected, result.Rolls)
		}
	}
}

func Test_Exploding_RollN_CorrectlyReturnsResults(t *testing.T) {
	seq := []int{1, 2, 3, 4}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	results := die.RollN(4)
	if len(results.Rolls) != len(seq) {
		t.Fatalf("RollN expected %v, actual %v", seq, results)
	}

	for i, actual := range results.Rolls {
		expected := seq[i]
		if expected != actual {
			t.Fatalf("RollN expected %v, actual %v", seq, results.Rolls)
		}
	}
}

func Test_Exploding_RollN_CorrectlyReturnsExplosions(t *testing.T) {
	seq := []int{4, 8, 8, 1}
	r := roller.WithSequence(seq)

	die := Exploding(r, 8, 10)

	results := die.RollN(2)
	if len(results.Rolls) != len(seq) {
		t.Fatalf("RollN expected %v, actual %v", seq, results.Rolls)
	}

	for i, actual := range results.Rolls {
		expected := seq[i]
		if expected != actual {
			t.Fatalf("RollN expected %v, actual %v", seq, results.Rolls)
		}
	}
}
