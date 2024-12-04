package day02

import (
	"testing"
)

func Test_day02(t *testing.T) {

	day02Solution := Day02Solution{
		values: [][]int{
			[]int{7, 6, 4, 2, 1},
			[]int{1, 2, 7, 8, 9},
			[]int{9, 7, 6, 2, 1},
			[]int{1, 3, 2, 4, 5},
			[]int{8, 6, 4, 4, 1},
			[]int{1, 3, 6, 7, 9},
		},
	}

	got := day02Solution.Part1()
	want := 2

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

}

func Test_isSafe(t *testing.T) {

	input := []int{1, 3, 6, 7, 9}
	got := isSafe(&input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_isNotSafe(t *testing.T) {

	input := []int{1, 2, 7, 8, 9}
	got := isSafe(&input)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_isSafe_jagged(t *testing.T) {
	input := []int{1, 0, 2, 3}
	got := isSafe(&input)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_isDescending(t *testing.T) {
	input := []int{3, 2, 1, 0}
	got := isSafe(&input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_isAscending(t *testing.T) {
	input := []int{0, 1, 2, 3}
	got := isSafe(&input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_isDescending_long(t *testing.T) {
	input := []int{1, 1, 1, 0}
	got := isSafe(&input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func Test_isAscending_long(t *testing.T) {
	input := []int{0, 0, 0, 1}
	got := isSafe(&input)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
