package day1

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	if os.Getenv("RUN_TEMPLATE") != "false" {
		t.Skip("Skipping template test.")
	}

	testInput := []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
	var expectedResult int64 = 11

	result, err := Part1(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	if os.Getenv("RUN_TEMPLATE") != "false" {
		t.Skip("Skipping template test.")
	}

	testInput := []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
	var expectedResult int64 = 31

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
