package day6

import (
	"fmt"
	"testing"
)

func TestGetObstructionPermutations(t *testing.T) {
	testInput := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	obstructions, player := GetBoardState(testInput)
	bounds := GetBounds(testInput)

	fmt.Println("Number of obstructions", len(obstructions))
	fmt.Println("Bounds", bounds)

	nonValidPlacements := len(obstructions) + 1
	possiblePlacements := (bounds.xMax + 1) * (bounds.yMax + 1)
	validPlacements := possiblePlacements - nonValidPlacements

	suggestedPlacements := GetObstructionPermutations(bounds, obstructions, player)

	if len(suggestedPlacements) != validPlacements {
		t.Fatalf(`Obstruction permutations: expected: %d got: %d`, validPlacements, len(suggestedPlacements))

	}
}

func TestPart1(t *testing.T) {
	testInput := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	var expectedResult int64 = 41

	result, err := Part1(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	testInput := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}
	var expectedResult int64 = 6

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
