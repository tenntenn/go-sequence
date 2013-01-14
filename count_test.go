package seq

import (
	"testing"
)

func TestCountUpNtoM(t *testing.T) {
	n := -2
	m := 30

	j := n
	for i := range CountNtoM(n, m) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}

		j++
	}

	if j != m + 1 {
		t.Errorf("j must be %d.", m + 1)
	}
}

func TestCountDownNtoM(t *testing.T) {
	n := 10
	m := -5

	j := n
	for i := range CountNtoM(n, m) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}

		j--
	}

	if j != m - 1 {
		t.Errorf("j must be %d.", m - 1)
	}
}

func TestCountUp(t *testing.T) {
	n := 10

	j := 0
	for i := range CountUp(n) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}
		j++
	}

	if j != n + 1 {
		t.Errorf("j must be %d.", n + 1)
	}
}

func TestCountDown(t *testing.T) {
	n := 10

	j := 10
	for i := range CountDown(n) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}
		j--
	}

	if j != -1 {
		t.Errorf("j must be %d.", -1)
	}
}

func TestStepBy(t *testing.T) {
	step := 3
	n := 200
	j := 0
	for i := range StepBy(CountUp(n), step) {
		if i != j {
			t.Errorf("i and j must be equal.")
		}
		j += step
	}

	if j != 201 {
		t.Errorf("j must be %d.", 201)
	}
}
