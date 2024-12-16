package day04

import (
	"advent-of-code-2024/types"
	utils "advent-of-code-2024/utils"
	"fmt"
	"strconv"
)

const day = "Day 04"

func search(m map[types.Vec2]int32, pos types.Vec2, s string, dir types.Vec2) int {
	if s == "" {
		return 1
	}

	head := int32(s[0])
	if m[pos] != head {
		return 0
	}

	return search(m, pos.Add(&dir), s[1:], dir)
}

func Run() {
	part1(utils.ReadInput("inputs/day04.txt"))
	part2(utils.ReadInput("inputs/day04.txt"))
}

func part1(input []string) {
	m := utils.ParseInputToMap(input)
	count := 0

	for y := range input {
		for x := range input[y] {
			head := types.Vec2{X: x, Y: y}
			count += search(m, head, "XMAS", types.Vec2{X: 0, Y: -1})
			count += search(m, head, "XMAS", types.Vec2{X: 1, Y: -1})
			count += search(m, head, "XMAS", types.Vec2{X: 1, Y: 0})
			count += search(m, head, "XMAS", types.Vec2{X: 1, Y: 1})
			count += search(m, head, "XMAS", types.Vec2{X: 0, Y: 1})
			count += search(m, head, "XMAS", types.Vec2{X: -1, Y: 1})
			count += search(m, head, "XMAS", types.Vec2{X: -1, Y: 0})
			count += search(m, head, "XMAS", types.Vec2{X: -1, Y: -1})
		}
	}

	fmt.Printf("%s: Part 1: %s\n", day, strconv.Itoa(count))
}

func part2(input []string) {
	m := utils.ParseInputToMap(input)
	count := 0
	dir := types.Vec2{X: 1, Y: 1}

	for y := range input {
		for x := range input[y] {
			head := types.Vec2{X: x, Y: y}

			if (search(m, head, "AS", dir) != 1 || search(m, head, "AM", dir.Multiply(-1)) != 1) && (search(m, head, "AM", dir) != 1 || search(m, head, "AS", dir.Multiply(-1)) != 1) {
				continue
			}

			if (search(m, head, "AS", dir.RotateRight()) == 1 && search(m, head, "AM", dir.RotateRight().Multiply(-1)) == 1) || search(m, head, "AM", dir.RotateRight()) == 1 && search(m, head, "AS", dir.RotateRight().Multiply(-1)) == 1 {
				count++
			}
		}
	}
	fmt.Printf("%s: Part 2: %s\n", day, strconv.Itoa(count))
}

