package underscore

import "testing"

func Test_Range_StepEq0(t *testing.T) {
	defer func() {
		if rv := recover(); rv == nil {
			t.Error("wrong")
		}
	}()

	dst := make([]int, 0)
	Range2(0, 10, 0).Value(&dst)
}

func Test_Range_StartEqStop(t *testing.T) {
	dst := make([]int, 0)
	Range2(0, 0, 1).Value(&dst)
	if len(dst) != 0 {
		t.Error("wrong")
	}
}

func Test_Range_Increment(t *testing.T) {
	size := 10
	dst := make([]int, 0)
	Range2(0, size, 1).Value(&dst)
	if len(dst) != size {
		t.Fatal(dst)
	}

	for i := 0; i < size; i++ {
		if dst[i] != i {
			t.Fatal(dst)
		}
	}
}

func Test_Range_Decrement(t *testing.T) {
	start := 10
	step := -2
	dst := make([]int, 0)
	Range2(start, 0, step).Value(&dst)
	if len(dst) != 5 {
		t.Fatal(dst)
	}

	for i := 0; i < 5; i++ {
		if dst[i] != start {
			t.Fatal(dst)
		}
		start += step
	}
}
