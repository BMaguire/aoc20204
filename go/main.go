package main

import (
	"fmt"
	"solution/util"
)

const SOLUTION_PATH = "../assets/"

func main() {

	for line := range util.ReadFile(SOLUTION_PATH+"example/test", '\n') {

		fmt.Println(line)

		for token := range util.ReadText(line, ',') {
			fmt.Println(token)
		}
	}
}
