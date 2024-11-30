package main

import (
	"flag"
	"fmt"
	"solution/solutions/day01"
	_ "solution/solutions/day01"
	"solution/util/runner"
)

func main() {

	wordPtr := flag.String("day", "*", "a string")

	flag.Parse()

	fmt.Println("day:", *wordPtr)

	if *wordPtr == "*" {

	} else {
		runner.Run(day01.Day01)

	}
	/*
		for line := range util.ReadFile(SOLUTION_PATH+"example/test", '\n') {

			fmt.Println(line)

			for token := range util.ReadText(line, ',') {
				fmt.Println(token)
			}
		}
	*/

}
