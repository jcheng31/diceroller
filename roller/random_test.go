package roller

import (
	"math/rand"
	"testing"
)

func Test_Random_Roll_ReturnsFromSource(t *testing.T) {
	source := rand.NewSource(42)
	roller := WithRandomSource(source)

	testSource := rand.NewSource(42)
	testRand := rand.New(testSource)
	for i := 0; i < 100; i++ {
		expected := testRand.Intn(8) + 1
		actual := roller.Roll(8)
		if expected != actual {
			t.Errorf("Random.Roll expected %v, got %v", expected, actual)
		}
	}
}
