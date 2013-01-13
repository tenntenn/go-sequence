package sequence

// Retuerns count up or down channel.
func CountNtoM(n, m int) <-chan int {
	ch := make(chan int)
	go func() {
		if n < m {
			for i := n; i <= m; i++ {
				ch <- i
			}
		} else {
			for i := n; i >= m; i-- {
				ch <- i
			}
		}
		close(ch)
	}()

	return (<-chan int)(ch)
}

// Count up from 0 to n.
func CountUp(n int) <-chan int {

	if n < 0 {
		panic("n must be more than 0.")
	}

	return CountNtoM(0, n)
}

// Count down from n to 0.
func CountDown(n int) <-chan int {

	if n < 0 {
		panic("n must be more than 0.")
	}

	return CountNtoM(n, 0)
}

// Send by step.
func StepBy(src <-chan int, step int) <-chan int {
	if src == nil {
		panic("src must not be nil")
	} else if step <= 0 {
		panic("step must be more than 0.")
	}

	dest := make(chan int)
	go func() {
		for i := range src {
			dest <- i
			for _ = range CountUp(step-2) {
				i = <-src
			}
		}
		close(dest)
	}()

	return (<-chan int)(dest)
}
