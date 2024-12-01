package util

import "fmt"

func GetInputPath(day int) string {
	return fmt.Sprintf("./pkg/inputs/day%d.txt", day)
}
