package seq

// Sequence of number.
type Sequence <-chan int

// Sequence generator.
type SequenceGenerator interface {
	// Create sequence generator.
	Create() Sequence
}

// Send by step.
func StepBy(src Sequence, step int) Sequence {
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

	return Sequence((<-chan int)(dest))
}
