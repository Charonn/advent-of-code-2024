package day01

import (
	utils "advent-of-code-2024/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const day = "Day 01"

func Run() {
	part1(utils.ReadInput("inputs/day01.txt"))
	part2(utils.ReadInput("inputs/day01.txt"))
}

func part1(input []string) {
	left := make([]int, 0, len(input))
	right := make([]int, 0, len(input))

	for _, line := range input {
		raw := strings.Split(line, "   ")
		leftI, _ := strconv.Atoi(raw[0])
		rightI, _ := strconv.Atoi(raw[1])
		left = append(left, leftI)
		right = append(right, rightI)
	}
	sort.Ints(left)
	sort.Ints(right)

	sum := 0
	for i := 0; i < len(left); i++ {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	fmt.Printf("%s: Part 1: %-15d", day, sum)
}

func part2(input []string) {
	left := make([]int, 0, len(input))
	right := make([]int, 0, len(input))

	for _, line := range input {
		raw := strings.Split(line, "   ")
		leftI, _ := strconv.Atoi(raw[0])
		rightI, _ := strconv.Atoi(raw[1])
		left = append(left, leftI)
		right = append(right, rightI)
	}

	sum := 0
	for _, i := range left {
		count := 0

		for _, j := range right {
			if i == j {
				count++
			}
		}

		sum += i * count
	}

	fmt.Printf("Part 2: %-15d", sum)
}
