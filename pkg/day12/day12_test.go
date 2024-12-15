package day12

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testInput := []string{
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
	var expectedResult int64 = 140

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
		"AAAA",
		"BBCD",
		"BBCC",
		"EEEC",
	}
	var expectedResult int64 = 80

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2E(t *testing.T) {
	testInput := []string{
		"EEEEE",
		"EXXXX",
		"EEEEE",
		"EXXXX",
		"EEEEE",
	}
	var expectedResult int64 = 236

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2AB(t *testing.T) {
	testInput := []string{
		"AAAAAA",
		"AAABBA",
		"AAABBA",
		"ABBAAA",
		"ABBAAA",
		"AAAAAA",
	}
	var expectedResult int64 = 368

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
