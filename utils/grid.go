package utils

import "advent-of-code-2024/types"

func ParseInputToMap(input []string) map[types.Vec2]int32 {
	m := map[types.Vec2]int32{}
	for y, row := range input {
		for x, c := range row {
			m[types.Vec2{X: x, Y: y}] = c
		}
	}
	return m
}
