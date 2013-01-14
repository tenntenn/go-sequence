package seq

type NumberSequence struct {
	Start int
	End int
}

func (self *NumberSequence) Create() Sequence {
	ch := make(chan int)
	go func() {
		if self.Start < self.End {
			for i := self.Start; i <= self.End; i++ {
				ch <- i
			}
		} else {
			for i := self.Start; i >= self.End; i-- {
				ch <- i
			}
		}
		close(ch)
	}()

	return Sequence((<-chan int)(ch))
}

// Retuerns count up or down channel.
func CountNtoM(n, m int) Sequence {
	ns := &NumberSequence{n, m}
	return ns.Create()
}

// Count up from 0 to n.
func CountUp(n int) Sequence {

	if n < 0 {
		panic("n must be more than 0.")
	}

	return CountNtoM(0, n)
}

// Count down from n to 0.
func CountDown(n int) Sequence {

	if n < 0 {
		panic("n must be more than 0.")
	}

	return CountNtoM(n, 0)
}
