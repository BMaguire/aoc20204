package day01

import (
	"slices"
	"solution/util/parser"
	"solution/util/runner"
	"strconv"
	"strings"
)

type Day01Solution struct {
	left  []int
	right []int
}

func (d *Day01Solution) Parse(path string) {
	var lines = parser.ReadFileWithDelim(path, '\n')
	var leftList []int = []int{}
	var rightList []int = []int{}

	for line := range lines {
		var left, right int

		locations := strings.Split(line, "   ")

		left, _ = strconv.Atoi(locations[0])
		right, _ = strconv.Atoi(strings.Trim(locations[1], "\r\n"))

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	d.left = append(make([]int, 0, len(leftList)), leftList...)
	d.right = append(make([]int, 0, len(rightList)), rightList...)
}

func (d *Day01Solution) Part1() int {
	// sort both lists
	sortedLeft := make([]int, len(d.left))
	sortedRight := make([]int, len(d.right))

	copy(sortedLeft, d.left)
	copy(sortedRight, d.right)

	slices.Sort(sortedRight)
	slices.Sort(sortedLeft)

	var total = 0

	for index, leftVal := range sortedLeft {
		var difference int
		if leftVal < sortedRight[index] {
			difference = sortedRight[index] - leftVal
		} else {

			difference = leftVal - sortedRight[index]
		}
		total += difference
	}

	return total
}

func (d *Day01Solution) Part2() int {
	rightAppearances := map[int]int{}

	for _, right := range d.right {
		val, prs := rightAppearances[right]
		if prs {
			rightAppearances[right] = val + 1
		} else {
			rightAppearances[right] = 1
		}
	}

	var total = 0
	for _, left := range d.left {
		val, prs := rightAppearances[left]
		if prs {
			total += left * val
		}
	}

	return total
}

var day01Solution = Day01Solution{}

var Day01 = runner.DayRunner{
	Day:      1,
	Solution: &day01Solution,
}
