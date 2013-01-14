package seq

import (
	"math/rand"
)

type RandSeq struct {
	*rand.Rand
}

func (self *RandSeq) Create() Sequence {
	ch := make(chan int)
	go func() {
		defer close(ch)

		for {
			ch <- self.Int()
		}
	}()

	return Sequence((<-chan int)(ch))
}

func Rand(rnd *rand.Rand) *RandSeq {
	return &RandSeq{rnd}
}
