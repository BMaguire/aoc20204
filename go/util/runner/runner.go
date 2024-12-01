package runner

import "fmt"

// A structure that will hold the information needed to run each day

type Solution interface {
	Parse(string)
	Part1() int
	Part2() int
}

type DayRunner struct {
	Day      int
	Solution Solution
}

const SOLUTION_PATH = "../inputs/"

func Run(runner DayRunner) {
	path := SOLUTION_PATH + "day" + fmt.Sprintf("%02d", runner.Day) + "/input"

	fmt.Printf("Parse input %s \n", path)

	runner.Solution.Parse(SOLUTION_PATH + "day" + fmt.Sprintf("%02d", runner.Day) + "/input")

	fmt.Println("Part 1")
	fmt.Printf("result: %d \n", runner.Solution.Part1())

	fmt.Println("Part 2")
	fmt.Printf("result: %d \n", runner.Solution.Part2())

}
