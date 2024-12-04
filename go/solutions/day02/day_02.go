package day02

import (
	"math"
	"solution/util/parser"
	"solution/util/runner"
	"strconv"
	"strings"
)

type Day02Solution struct {
	values [][]int
}

var day02Solution = Day02Solution{}

var Day02 = runner.DayRunner{
	Day:      2,
	Solution: &day02Solution,
}

func (d *Day02Solution) Parse(path string) {

	var lines = parser.ReadFileWithDelim(path, '\n')

	items := [][]int{}
	for line := range lines {

		values := strings.Split(strings.Trim(line, "\r\n"), " ")

		nums := []int{}

		for _, value := range values {
			val, _ := strconv.Atoi(value)

			nums = append(nums, val)

		}

		items = append(items, nums)
	}

	d.values = append(make([][]int, 0, len(items)), items...)

}

func (d *Day02Solution) Part1() int {

	count := 0
	for _, row := range d.values {

		if isSafe(&row) {
			count += 1
		}

	}

	return count

}

func isSafe(row *[]int) bool {
	var prev int = 0
	if len(*row) < 2 {
		return false
	}

	var gradient bool = (*row)[len(*row)-1]-(*row)[0] > 0
	for i, height := range *row {
		if i > 0 {
			diff := height - prev

			if diff == 0 {
				return false
			}

			if math.Abs(float64(diff)) > 3 {
				return false
			}

			if !(diff > 0 == gradient) {
				return false
			}
		}

		prev = height

	}

	return true
}

func (d *Day02Solution) Part2() int {

	return 0
}
