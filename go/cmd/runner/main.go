package main

import (
	"flag"
	"fmt"

	"solution/solutions/day02"
	"solution/util/runner"
)

func main() {

	wordPtr := flag.String("day", "*", "a string")

	flag.Parse()

	fmt.Println("day:", *wordPtr)

	runner.Run(day02.Day02)

	/*
		for line := range util.ReadFile(SOLUTION_PATH+"example/test", '\n') {

			fmt.Println(line)

			for token := range util.ReadText(line, ',') {
				fmt.Println(token)
			}
		}
	*/

}
