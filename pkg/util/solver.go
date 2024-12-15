package util

import (
	"github.com/rickardlofberg/adventofcode2024/pkg/day1"
	"github.com/rickardlofberg/adventofcode2024/pkg/day10"
	"github.com/rickardlofberg/adventofcode2024/pkg/day11"
	"github.com/rickardlofberg/adventofcode2024/pkg/day12"
	"github.com/rickardlofberg/adventofcode2024/pkg/day2"
	"github.com/rickardlofberg/adventofcode2024/pkg/day3"
	"github.com/rickardlofberg/adventofcode2024/pkg/day4"
	"github.com/rickardlofberg/adventofcode2024/pkg/day5"
	"github.com/rickardlofberg/adventofcode2024/pkg/day6"
	"github.com/rickardlofberg/adventofcode2024/pkg/day7"
	"github.com/rickardlofberg/adventofcode2024/pkg/day8"
	"github.com/rickardlofberg/adventofcode2024/pkg/day9"
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
	case day == 5 && part == 1:
		return day5.Part1
	case day == 5 && part == 2:
		return day5.Part2
	case day == 6 && part == 1:
		return day6.Part1
	case day == 6 && part == 2:
		return day6.Part2
	case day == 7 && part == 1:
		return day7.Part1
	case day == 7 && part == 2:
		return day7.Part2
	case day == 8 && part == 1:
		return day8.Part1
	case day == 8 && part == 2:
		return day8.Part2
	case day == 9 && part == 1:
		return day9.Part1
	case day == 9 && part == 2:
		return day9.Part2
	case day == 10 && part == 1:
		return day10.Part1
	case day == 10 && part == 2:
		return day10.Part2
	case day == 11 && part == 1:
		return day11.Part1
	case day == 11 && part == 2:
		return day11.Part2
	case day == 12 && part == 1:
		return day12.Part1
	case day == 12 && part == 2:
		return day12.Part2
	default:
		return nil
	}
}
