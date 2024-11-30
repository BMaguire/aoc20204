package runner

import "fmt"

// A structure that will hold the information needed to run each day

type Solution interface {
	Parse(string)
	Part1()
	Part2()
}

type DayRunner struct {
	Day      int
	Solution Solution
}

func (c *DayRunner) setSolution(s Solution) {
	c.Solution = s
}

const SOLUTION_PATH = "../assets/"

func Run(runner DayRunner) {

	runner.Solution.Parse(SOLUTION_PATH + "day" + fmt.Sprintf("%02d", runner.Day) + "/input")

	runner.Solution.Part1()

	runner.Solution.Part2()

}
