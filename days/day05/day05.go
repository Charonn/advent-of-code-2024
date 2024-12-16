package day05

import (
	"advent-of-code-2024/types"
	utils "advent-of-code-2024/utils"
	"fmt"
	"slices"
	"strings"
)

const day = "Day 05"

func Run() {
	part1(utils.ReadInput("inputs/day05.txt"))
	part2(utils.ReadInput("inputs/day05.txt"))
}

func check(values []int, rule []types.Vec2) {
	for y := 0; y < len(values); y++ {
		for x := y + 1; x < len(values); x++ {
			if slices.Contains(rule, types.Vec2{X: values[x], Y: values[y]}) {
				tmp := values[y]
				values[y] = values[x]
				values[x] = tmp
			}
		}
	}
}

func part1(input []string) {
	emptyLineId := slices.Index(input, "")
	rules := make([]types.Vec2, 0, emptyLineId)
	sum := 0

	for _, line := range input[:emptyLineId] {
		values := utils.ToIntSlice(strings.Split(line, "|"))
		rules = append(rules, types.Vec2{X: values[0], Y: values[1]})
	}

	for _, line := range input[emptyLineId+1:] {
		values := utils.ToIntSlice(strings.Split(line, ","))
		clone := slices.Clone(values)

		check(clone, rules)

		if slices.Compare(values, clone) == 0 {
			sum += clone[(len(clone)-1)/2]
		}
		
	}

	fmt.Printf("%s: Part 1: %d\n", day, sum)
}

func part2(input []string) {
	emptyLineId := slices.Index(input, "")
	rules := make([]types.Vec2, 0, emptyLineId)
	sum := 0

	for _, line := range input[:emptyLineId] {
		values := utils.ToIntSlice(strings.Split(line, "|"))
		rules = append(rules, types.Vec2{X: values[0], Y: values[1]})
	}

	for _, line := range input[emptyLineId+1:] {
		values := utils.ToIntSlice(strings.Split(line, ","))
		clone := slices.Clone(values)

		check(clone, rules)

		if slices.Compare(values, clone) != 0 {
			sum += clone[(len(clone)-1)/2]
		}
		
	}

	fmt.Printf("%s: Part 2: %d\n", day, sum)
}
