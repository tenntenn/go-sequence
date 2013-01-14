package seq

type FiboSequence struct {}

func (self *FiboSequence) Create() Sequence {
	ch := make(chan int)
	go func() {
		defer close(ch)

		fibo := []int{0, 1}
		ch <- fibo[0]
		ch <- fibo[1]
		for i := 2;; i++ {
			fibo = append(fibo, fibo[i-2] + fibo[i-1])
			ch <- fibo[i]
		}
	}()

	return Sequence((<-chan int)(ch))
}

func Fibo() Sequence {
	fibo := new(FiboSequence)
	return fibo.Create()
}
