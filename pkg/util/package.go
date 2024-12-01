package util

import "fmt"

func GetInputPath(day int) string {
	return fmt.Sprintf("./pkg/inputs/day%d.txt", day)
}

func GetSolverString(day int, part int) string {
	return fmt.Sprintf("%d.%d", day, part)
}