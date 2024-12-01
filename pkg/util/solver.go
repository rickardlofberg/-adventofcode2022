package util

import (
	"github.com/rickardlofberg/adventofcode2024/pkg/day1"
)

type fn func([]string) (int64, error)

func GetSolver(day int, part int) fn {
	switch true {
	case day == 1 && part == 1:
		return day1.Part1
	case day == 1 && part == 2:
		return day1.Part2
	default:
		return nil
	}
}
