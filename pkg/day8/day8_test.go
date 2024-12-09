package day8

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetAntiNodes(t *testing.T) {
	n1 := coordinate{x: 1, y: 1}
	n2 := coordinate{x: 3, y: 3}
	expectedResult := []coordinate{
		{x: -1, y: -1},
		{x: 5, y: 5},
	}
	result := GetAntiNodes(n1, n2)

	if !slices.Equal(result, expectedResult) {
		fmt.Println(result, "!=", expectedResult)
		t.Fatalf("Getting nodes wrong")

	}
}

func TestPart1(t *testing.T) {
	testInput := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
	var expectedResult int64 = 14

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
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}
	var expectedResult int64 = 34

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
