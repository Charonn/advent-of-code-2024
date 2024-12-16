package day03

import (
	utils "advent-of-code-2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const day = "Day 03"

// eval receives a row as input and evaluates the valid multiplications
func eval(row string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	sum := 0
	matches := re.FindAllString(row, -1)

	for _, match := range matches {
		sum += mulExec(match)
	}

	return sum
}

// evalAdvanced receives a row as input and evaluates the valid multiplications with additional switch logic
func evalAdvanced(row string) int {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	sum := 0
	enabled := true
	matches := re.FindAllString(row, -1)

	for _, match := range matches {
		if match == "do()" {
			enabled = true
		} else if match == "don't()" {
			enabled = false
		} else if enabled {
			sum += mulExec(match)
		}
	}

	return sum
}

// mul receives a mul instruction and returns the multiplication result
func mulExec(str string) int {
	values := strings.Split(str[4:len(str)-1], ",")
	i, _ := strconv.Atoi(values[0])
	j, _ := strconv.Atoi(values[1])
	return i * j
}

func Run() {
	part1(utils.ReadInput("inputs/day03.txt"))
	part2(utils.ReadInput("inputs/day03.txt"))
}

func part1(input []string) {
	sum := 0
	for _, row := range input {
		sum += eval(row)
	}
	fmt.Printf("%s: Part 1: %-15d", day, sum)
}

func part2(input []string) {
	joined := ""
	for _, row := range input {
		joined += row
	}

	fmt.Printf("Part 2: %-15d", evalAdvanced(joined))
}

