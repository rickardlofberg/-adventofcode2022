package util

import (
	"github.com/rickardlofberg/adventofcode2024/pkg/day1"
)

type fn func([]string) (int64, error)

func GetSolver(problem string) fn {
	var solvers = map[string]fn{
		"1.1":  day1.Part1
	}
	return solvers[problem]
}