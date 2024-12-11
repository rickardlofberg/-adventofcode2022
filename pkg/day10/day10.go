package day10

import (
	"slices"
	"strconv"
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

func ParseInput(input []string) map[coordinate]int {
	lavaMap := map[coordinate]int{}

	for idxY := range input {
		for idxX := range input {
			height, _ := strconv.Atoi(string(input[idxY][idxX]))
			corr := coordinate{x: idxX, y: idxY}
			lavaMap[corr] = height
		}
	}
	return lavaMap
}

func GetBounds(arr []string) bounds {
	yMax := len(arr) - 1
	xMax := len(arr[0]) - 1
	return bounds{xMin: 0, yMin: 0, xMax: xMax, yMax: yMax}
}

func InBounds(pos coordinate, bounds bounds) bool {
	if pos.x < bounds.xMin || pos.x > bounds.xMax {
		return false
	} else if pos.y < bounds.yMin || pos.y > bounds.yMax {
		return false
	}
	return true
}

func GetAvailablePaths(corr coordinate, lavaMap map[coordinate]int, bounds bounds) []coordinate {
	validPaths := []coordinate{}
	currentElevation := lavaMap[corr]

	south := coordinate{x: corr.x, y: corr.y - 1}
	north := coordinate{x: corr.x, y: corr.y + 1}
	east := coordinate{x: corr.x + 1, y: corr.y}
	west := coordinate{x: corr.x - 1, y: corr.y}
	paths := []coordinate{south, north, east, west}

	for _, path := range paths {
		validElevation := lavaMap[path] - 1
		isValid := true
		isValid = isValid && InBounds(path, bounds)
		isValid = isValid && currentElevation == validElevation
		if isValid {
			validPaths = append(validPaths, path)
		}
	}

	return validPaths
}

func GetNumberOfPaths(corr coordinate, lavaMap map[coordinate]int, bounds bounds, target int) []coordinate {
	result := []coordinate{}

	isTarget := lavaMap[corr] == target
	if isTarget {
		return []coordinate{corr}
	}

	availablePaths := GetAvailablePaths(corr, lavaMap, bounds)
	for _, path := range availablePaths {
		paths := GetNumberOfPaths(path, lavaMap, bounds, target)
		result = append(result, paths...)
	}

	return result
}

func GetUniquePaths(paths []coordinate) int64 {
	seenPaths := []coordinate{}
	result := int64(0)

	for _, path := range paths {
		isSeen := slices.Contains(seenPaths, path)
		if !isSeen {
			seenPaths = append(seenPaths, path)
			result += 1
		}
	}
	return result
}

func SolvePart1(bounds bounds, lavaMap map[coordinate]int) int64 {
	result := int64(0)
	for idxY := range bounds.yMax + 1 {
		for idxX := range bounds.xMax + 1 {
			corr := coordinate{x: idxX, y: idxY}
			if lavaMap[corr] == 0 {
				numberOfPaths := GetNumberOfPaths(corr, lavaMap, bounds, 9)
				uniquePaths := GetUniquePaths(numberOfPaths)
				result += uniquePaths
			}
		}
	}
	return result
}

func SolvePart2(bounds bounds, lavaMap map[coordinate]int) int64 {
	result := int64(0)
	for idxY := range bounds.yMax + 1 {
		for idxX := range bounds.xMax + 1 {
			corr := coordinate{x: idxX, y: idxY}
			if lavaMap[corr] == 0 {
				result += int64(len(GetNumberOfPaths(corr, lavaMap, bounds, 9)))
			}
		}
	}
	return result
}

func Part1(input []string) (int64, error) {
	bounds := GetBounds(input)
	lavaMap := ParseInput(input)

	result := SolvePart1(bounds, lavaMap)
	return result, nil
}

func Part2(input []string) (int64, error) {
	bounds := GetBounds(input)
	lavaMap := ParseInput(input)

	result := SolvePart2(bounds, lavaMap)
	return result, nil
}
