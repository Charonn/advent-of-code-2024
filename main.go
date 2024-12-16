package main

import (
	"advent-of-code-2024/days/day01"
	"advent-of-code-2024/days/day02"
	"advent-of-code-2024/days/day03"
	"advent-of-code-2024/days/day04"
	"advent-of-code-2024/days/day05"
	"advent-of-code-2024/days/day06"
	"fmt"
	"time"
)

func measureExecutionTime(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("\t%.3f\n", elapsed.Seconds())
}

func main() {
	measureExecutionTime(day01.Run)
	measureExecutionTime(day02.Run)
	measureExecutionTime(day03.Run)
	measureExecutionTime(day04.Run)
	measureExecutionTime(day05.Run)
	measureExecutionTime(day06.Run)
}
