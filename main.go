package main

import (
	"fmt"

	day1 "github.com/jalvaydev/advent-of-code/day_1"
)

func main() {
	fmt.Println("--- Day 1 ---")
	fmt.Println("Day 1, Part 1:", day1.FindTopCalorieElf("./day_1/input"))
	fmt.Println("Day 1, Part 2:", day1.FindTopKCalorieElves("./day_1/input", 3))
}
