package main

import (
	"day15/solution"
	"fmt"
	"os"
)

func main() {
	text, _ := os.ReadFile("input.txt")
	data := string(text)
	// fmt.Println(solution.Part1(data))
	fmt.Println(solution.Part2(data))
}
