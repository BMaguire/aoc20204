package day01

import (
	"testing"
)

func Test_day01(t *testing.T) {

	day01Solution := Day01Solution{
		left:  []int{3, 4, 2, 1, 3, 3},
		right: []int{4, 3, 5, 3, 9, 3},
	}

	got := day01Solution.Part1()
	want := 11

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func Test_day02(t *testing.T) {

	day01Solution := Day01Solution{
		left:  []int{3, 4, 2, 1, 3, 3},
		right: []int{4, 3, 5, 3, 9, 3},
	}

	got := day01Solution.Part2()
	want := 31

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}
