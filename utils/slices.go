package utils

import "strconv" 

func ToIntSlice(numbers []string) []int {
	s := make([]int, 0, len(numbers))
	for _, number := range numbers {
		i, _ := strconv.Atoi(number)
		s = append(s, i)
	}
	return s
}
