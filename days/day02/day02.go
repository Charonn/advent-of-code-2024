package day02

import (
	utils "advent-of-code-2024/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const day = "Day 02"

func Run() {
	part1(utils.ReadInput("inputs/day02.txt"))
	part2(utils.ReadInput("inputs/day02.txt"))
}


func isValid(plan []int) bool {
	sign := math.Signbit(float64(plan[1] - plan[0]))
	for i := 1; i < len(plan); i++ {
		diff := plan[i] - plan[i-1]
		absDiff := math.Abs(float64(diff))
		if math.Signbit(float64(diff)) != sign || absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func part1(input []string) {
	count := 0
	
	for _, plan := range input {
		values := utils.ToIntSlice(strings.Split(plan, " "))
		if isValid(values) {
			count++
		}
	}

	fmt.Printf("%s: Part 1: %s\n", day, strconv.Itoa(count))
}

func part2(input []string) {
	count := 0

	for _, report := range input {
		values := utils.ToIntSlice(strings.Split(report, " "))

		if isValid(values) {
			count++
		} else {
			for i := range len(values) {
				adjustedValues := make([]int, 0, len(values)-1)
				adjustedValues = append(adjustedValues, values[:i]...)
				adjustedValues = append(adjustedValues, values[i+1:]...)

				if isValid(adjustedValues) {
					count++
					break
				}
			}
		}
	}
	fmt.Printf("%s: Part 2: %d\n", day, count)
}
