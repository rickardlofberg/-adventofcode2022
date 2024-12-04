package util

import (
	"github.com/rickardlofberg/adventofcode2024/pkg/day1"
	"github.com/rickardlofberg/adventofcode2024/pkg/day2"
	"github.com/rickardlofberg/adventofcode2024/pkg/day3"
	"github.com/rickardlofberg/adventofcode2024/pkg/day4"
)

type fn func([]string) (int64, error)

func GetSolver(day int, part int) fn {
	switch true {
	case day == 1 && part == 1:
		return day1.Part1
	case day == 1 && part == 2:
		return day1.Part2
	case day == 2 && part == 1:
		return day2.Part1
	case day == 2 && part == 2:
		return day2.Part2
	case day == 3 && part == 1:
		return day3.Part1
	case day == 3 && part == 2:
		return day3.Part2
	case day == 4 && part == 1:
		return day4.Part1
	case day == 4 && part == 2:
		return day4.Part2
	default:
		return nil
	}
}
