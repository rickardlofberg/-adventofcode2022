package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func abs(i int64) int64 {
	if i < 0 {
		return i * -1
	}
	return i
}

type Direction string

const (
	Increase Direction = "increase"
	Decrease Direction = "decrease"
	Same     Direction = "same"
)

func GetDirection(first int64, second int64) Direction {
	if first < second {
		return Increase
	} else if first > second {
		return Decrease
	} else {
		return Same
	}
}

func IsValidDirection(directions map[Direction]int64, direction Direction) bool {
	switch direction {
	case Same:
		return false
	case Increase:
		if directions[Decrease] > 0 {
			return false
		}
	case Decrease:
		if directions[Increase] > 0 {
			return false
		}
	}
	return true
}

func IsAllowedDifference(first int64, second int64, minStep int64, maxStep int64, direction Direction) bool {
	var difference int64 = 0
	switch direction {
	case Same:
		return false
	case Increase:
		difference = second - first
	case Decrease:
		difference = first - second
	}
	if difference < minStep {
		return false
	}
	if difference > maxStep {
		return false
	}
	return true
}

type edge struct {
	start     int64
	end       int64
	direction Direction
	length    int64
}

func (e edge) HasValidLength(min int64, max int64) bool {
	var abs_length = abs(e.length)
	if abs_length < min {
		return false
	}
	if abs_length > max {
		return false
	}
	return true
}

func GetEdgesDirections(edges []edge) map[Direction]int64 {
	directions := make(map[Direction]int64)
	for idx := range edges {
		edge := edges[idx]
		directions[edge.direction] += 1
	}
	return directions
}

func GetEdges(integers []string) ([]edge, error) {
	var previous int64 = -1
	var edges []edge
	for idx := range len(integers) {
		current, err := strconv.ParseInt(integers[idx], 10, 64)
		if err != nil {
			return nil, errors.New("bad integer")
		}
		if previous == -1 {
			previous = current
			continue
		}
		var direction = GetDirection(previous, current)
		var length = current - previous
		var edge = edge{start: previous, end: current, direction: direction, length: length}
		edges = append(edges, edge)
		previous = current
	}
	return edges, nil
}

func MergeEdges(firstEdge edge, secondEdge edge) edge {
	start := firstEdge.end
	end := secondEdge.end
	direction := GetDirection(start, end)
	length := end - start
	return edge{start: start, end: end, direction: direction, length: length}
}

func FilterEdgesWithValue(edges []edge, min int64, max int64) []edge {
	var new_edges []edge
	var skip bool = false
	for idx := range len(edges) {
		var is_start = idx == 0
		var is_end = idx == len(edges)-1
		var edge = edges[idx]

		if skip {
			skip = false
			continue
		}

		// To remove
		if !edge.HasValidLength(min, max) {
			// Start and finish can just be removed
			if is_start || is_end {
				continue
			}
			previous_edge := edges[idx-1]
			next_edge := edges[idx+1]
			edge = MergeEdges(previous_edge, next_edge)
			skip = true
		}
		new_edges = append(new_edges, edge)
	}
	return new_edges

}

func FilterEdgesWithDirection(edges []edge, direction Direction) []edge {
	var new_edges []edge
	var skip bool = false
	for idx := range len(edges) {
		var is_start = idx == 0
		var is_end = idx == len(edges)-1
		var edge = edges[idx]

		if skip {
			skip = false
			continue
		}

		// To remove
		if edge.direction == direction || edge.direction == Same {
			// Start and finish can just be removed
			if is_start || is_end {
				continue
			}
			previous_edge := edges[idx-1]
			next_edge := edges[idx+1]
			edge = MergeEdges(previous_edge, next_edge)
			skip = true
		}
		new_edges = append(new_edges, edge)
	}
	return new_edges
}

func IsSafeLevel(row string, minStep int64, maxStep int64, tolerance int64) (bool, error) {
	integers := strings.Fields(row)
	edges, _ := GetEdges(integers)
	directions := GetEdgesDirections(edges)

	if directions[Same] > tolerance {
		return false, nil
	}

	var minorityDirection Direction
	if directions[Increase] > directions[Decrease] {
		minorityDirection = Decrease
	} else {
		minorityDirection = Increase
	}

	if directions[minorityDirection] > tolerance {
		return false, nil
	}
	edges_before := int64(len(edges))
	edges = FilterEdgesWithDirection(edges, minorityDirection)
	changes := edges_before - int64(len(edges))
	if tolerance > changes {
		edges = FilterEdgesWithValue(edges, minStep, maxStep)
	}

	changes = edges_before - int64(len(edges))
	if changes > tolerance {
		return false, nil
	}

	for idx := range edges {
		edge := edges[idx]
		if !edge.HasValidLength(minStep, maxStep) {
			return false, nil
		}
	}
	return true, nil
}

func IsSafeLevel2(row string, minStep int64, maxStep int64, tolerance int64) (bool, error) {
	integers := strings.Fields(row)
	directions := make(map[Direction]int64)
	var differences int64 = 0
	var first int64 = -1
	var second int64
	for idx := range len(integers) {
		current, err := strconv.ParseInt(integers[idx], 10, 64)
		if err != nil {
			return false, errors.New("bad integer")
		}

		if first == -1 {
			first = current
			continue
		}
		second = current

		var direction = GetDirection(first, second)
		var allowedDifference = IsAllowedDifference(first, second, minStep, maxStep, direction)
		var allowedDirection = IsValidDirection(directions, direction)

		if !allowedDirection {
			fmt.Println(allowedDirection)
		}

		if !allowedDifference || !allowedDirection {
			differences += 1
			if differences > tolerance {
				fmt.Println("Too many differences exit")
				return false, nil
			} else {
				continue
			}
		} else {
			directions[direction] += 1
		}

		first = second
	}

	fmt.Println("All good")
	return true, nil
}

func Part1(input []string) (int64, error) {
	var safeLevels int64 = 0
	var minStep int64 = 1
	var maxStep int64 = 3
	for _, row := range input {
		var isSafe, _ = IsSafeLevel(row, minStep, maxStep, 0)
		if isSafe {
			safeLevels += 1
		}
	}
	return safeLevels, nil
}

func Part2(input []string) (int64, error) {
	var safeLevels int64 = 0
	var minStep int64 = 1
	var maxStep int64 = 3
	for _, row := range input {
		var isSafe, _ = IsSafeLevel(row, minStep, maxStep, 1)
		if isSafe {
			safeLevels += 1
		}
	}
	return safeLevels, nil
}
