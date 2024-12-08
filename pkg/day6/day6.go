package day6

import (
	"fmt"
	"slices"
)

type coordinate struct {
	x int
	y int
}

type bounds struct {
	xMin int
	yMin int
	xMax int
	yMax int
}

type turn struct {
	x         int
	y         int
	direction int
}

const (
	Up    int = 0
	Right int = 1
	Down  int = 2
	Left  int = 3
)

const (
	Obstruction string = "#"
	Player      string = "^"
)

func GetBounds(arr []string) bounds {
	yMax := len(arr) - 1
	xMax := len(arr[0]) - 1
	return bounds{xMin: 0, yMin: 0, xMax: xMax, yMax: yMax}
}

func GetBoardState(arr []string) ([]coordinate, coordinate) {
	obstructions := []coordinate{}
	player := coordinate{}

	for yIdx := range arr {
		for xIdx := range arr[yIdx] {
			cell := string(arr[yIdx][xIdx])
			if cell == Obstruction {
				obstruction := coordinate{x: xIdx, y: yIdx}
				obstructions = append(obstructions, obstruction)
			} else if cell == Player {
				player = coordinate{x: xIdx, y: yIdx}
			}

		}
	}
	return obstructions, player
}

func InBounds(pos coordinate, bounds bounds) bool {
	if pos.x < bounds.xMin || pos.x > bounds.xMax {
		return false
	} else if pos.y < bounds.yMin || pos.y > bounds.yMax {
		return false
	}
	return true
}

func ChangeDirection(direction int) int {
	return (direction + 1) % 4
}

func GetNewPosition(player coordinate, direction int) coordinate {
	if direction == 0 {
		return coordinate{x: player.x, y: player.y - 1}
	} else if direction == 1 {
		return coordinate{x: player.x + 1, y: player.y}
	} else if direction == 2 {
		return coordinate{x: player.x, y: player.y + 1}
	} else {
		return coordinate{x: player.x - 1, y: player.y}
	}
}

func GetDirection(player coordinate, direction int, obstructions []coordinate) int {
	isObstructed := true
	for isObstructed {
		newPos := GetNewPosition(player, direction)
		isObstructed = slices.Contains(obstructions, newPos)
		if isObstructed {
			direction = ChangeDirection(direction)
		}
	}
	return direction
}

func GetDistinctPaths(bounds bounds, obstructions []coordinate, player coordinate) int64 {
	visited := []coordinate{}
	direction := Up
	for InBounds(player, bounds) {
		isVisited := slices.Contains(visited, player)
		if !isVisited {
			visited = append(visited, player)
		}
		direction = GetDirection(player, direction, obstructions)
		player = GetNewPosition(player, direction)
	}
	return int64(len(visited))
}

func IsStuck(bounds bounds, obstructions []coordinate, player coordinate) bool {
	visited := []coordinate{}
	direction := Up
	previousTurns := []turn{}

	for InBounds(player, bounds) {
		isVisited := slices.Contains(visited, player)
		if !isVisited {
			visited = append(visited, player)
		}
		newDirection := GetDirection(player, direction, obstructions)
		if newDirection != direction {
			turn := turn{x: player.x, y: player.y, direction: newDirection}
			hasTurnedBefore := slices.Contains(previousTurns, turn)
			if hasTurnedBefore {
				return true
			} else {
				previousTurns = append(previousTurns, turn)
			}
		}
		direction = newDirection
		player = GetNewPosition(player, direction)
	}
	return false
}

func IsStuckOld(bounds bounds, obstructions []coordinate, player coordinate, searchDepth int64) bool {
	visited := []coordinate{}
	direction := Up
	previousTurns := []turn{}

	for InBounds(player, bounds) {
		isVisited := slices.Contains(visited, player)
		if !isVisited {
			visited = append(visited, player)
		}
		newDirection := GetDirection(player, direction, obstructions)
		if newDirection != direction {
			turn := turn{x: player.x, y: player.y, direction: newDirection}
			if slices.Contains(previousTurns, turn) {
				return true
			} else {
				previousTurns = append(previousTurns, turn)
			}
		}
		direction = newDirection
		player = GetNewPosition(player, direction)
		searchDepth -= 1
		if searchDepth < 0 {
			fmt.Println("stuck")
			return true
		}
	}
	return false
}

func GetObstructionPermutations(bounds bounds, obstructions []coordinate, player coordinate) [][]coordinate {
	obstrucitonPermutations := [][]coordinate{}
	for yIdx := range bounds.yMax + 1 {
		for xIdx := range bounds.xMax + 1 {
			newObstruction := coordinate{x: xIdx, y: yIdx}
			isStartingPosition := newObstruction == player
			alredyObstructed := slices.Contains(obstructions, newObstruction)
			if isStartingPosition || alredyObstructed {
				continue
			}
			newObstructions := append(obstructions, newObstruction)
			newObstructionsCopy := make([]coordinate, len(newObstructions))
			_ = copy(newObstructionsCopy, newObstructions)
			obstrucitonPermutations = append(obstrucitonPermutations, newObstructionsCopy)
		}
	}

	return obstrucitonPermutations
}

func GetWaysOfObstructiong(bounds bounds, obstructions []coordinate, player coordinate) int64 {
	obstructionPermutations := GetObstructionPermutations(bounds, obstructions, player)
	result := int64(0)
	for _, permutation := range obstructionPermutations {
		isStuck := IsStuck(bounds, permutation, player)
		if isStuck {
			result += 1
		}
	}
	return result
}

func Part1(input []string) (int64, error) {
	bounds := GetBounds(input)
	obstructions, player := GetBoardState(input)
	result := GetDistinctPaths(bounds, obstructions, player)
	return result, nil
}

func Part2(input []string) (int64, error) {
	bounds := GetBounds(input)
	obstructions, player := GetBoardState(input)
	result := GetWaysOfObstructiong(bounds, obstructions, player)
	return result, nil
}
