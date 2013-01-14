package seq

import (
	"testing"
	"math/rand"
)

func TestRand(t *testing.T) {
	src1 := rand.NewSource(1000)
	src2 := rand.NewSource(1000)
	rnd1 := rand.New(src1)
	rnd2 := rand.New(src2)

	i := 0
	end := 10
	for actual := range Rand(rnd1) {

		if i == end {
			break
		}

		expect := rnd2.Int()
		if expect != actual {
			t.Errorf("expect and actual must be same.")
		}

		i++
	}
}
