package day9

import (
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	testInput := []string{"2333133121414131402"}
	var expectedResult int64 = 1928

	result, err := Part1(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	if os.Getenv("RUN_TEMPLATE") != "true" {
		t.Skip("Skipping template test.")
	}

	testInput := []string{}
	var expectedResult int64 = 0

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
