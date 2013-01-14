go-sequence
===========

This library provide functions which return channels.
The channels send sequence of numbers. 

Samples
-----------

	// CountUp 0 to 100
	for i := range seq.CountUp(100) {
		...
	}

	// CountDown 100 to 0
	for i := range seq.CountDown(100) {
		...
	}

	// CountUp 10 to 100
	for i := range seq.CountNtoM(10, 100) {
		...
	}

	// CountDown 100 to 10
	for i := range seq.CountNtoM(100, 10) {
		...
	}

	// CountUp 0 to 100 by step 3
	for i := range seq.StepBy(seq.CountUp(100)) {
		...
	}
