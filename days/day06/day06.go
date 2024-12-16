package day06

import (
	"advent-of-code-2024/types"
	utils "advent-of-code-2024/utils"
	"fmt"
	"maps"
	"slices"
)

const day = "Day 06"

type pair struct {
	pos types.Vec2
	dir types.Vec2
}

func Run() {
	part1(utils.ReadInput("inputs/day06.txt"))
	part2(utils.ReadInput("inputs/day06.txt"))
}

func getLocations(pos types.Vec2, guardDir types.Vec2, m map[types.Vec2]int32) (int, []types.Vec2) {
	var visitedLocations []types.Vec2
	var history []pair

	for ok := true; ok; _, ok = m[pos] {

		if !slices.Contains(visitedLocations, pos) {
			visitedLocations = append(visitedLocations, pos)
		}

		p := pair{pos: pos, dir: guardDir}
		if !slices.Contains(history, p) {
			history = append(history, p)
		} else {
			return -1, visitedLocations
		}

		nextPosition := pos.Add(&guardDir)

		if m[nextPosition] != '#' {
			pos = nextPosition
		} else {
			guardDir = guardDir.RotateRight()
		}
	}

	return len(visitedLocations), visitedLocations
}

func part1(input []string) {
	m := utils.ParseInputToMap(input)
	var guard types.Vec2
	for pos, val := range m {
		if val == '^' {
			guard = pos
			break
		}
	}
	locations, _ := getLocations(guard, types.Vec2{X: 0, Y: -1}, m)

	fmt.Printf("%s: Part 1: %-15d", day, locations)
}

func part2(input []string) {
	m := utils.ParseInputToMap(input)
	var guard types.Vec2

	for pos, char := range m {
		if char == '^' {
			guard = pos
			break
		}
	}

	loops := 0
	_, path := getLocations(guard, types.Vec2{X: 0, Y: -1}, m)

	for _, toReplace := range path {
		if m[toReplace] != '.' {
			continue
		}

		clone := maps.Clone(m)
		clone[toReplace] = '#'

		if count, _ := getLocations(guard, types.Vec2{X: 0, Y: -1}, clone); count == -1 {
			loops++
		}
	}
	fmt.Printf("Part 2: %-15d", loops)
}

