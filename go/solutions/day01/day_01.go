package day01

import (
	"fmt"
	util "solution/util/parser"
	"solution/util/runner"
)

type inputType = []string

var day01Solution = runner.Solution{
	Part1: part1,
	Part2: part2,
	Parse: parse,
}

type Day01 struct {
}

/*
func (l *Day01) evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}
*/

func parse(path string) inputType {
	var lines = util.ReadFileWithDelim(path, '\n')
	var out []string = []string{}

	for line := range lines {
		for token := range util.ReadText(line, ',') {
			fmt.Println(token)
			out = append(out, token)
		}
	}

	return out
}

func part1(lines inputType) int {
	return 0
}

func part2(lines inputType) int {
	return 0
}
