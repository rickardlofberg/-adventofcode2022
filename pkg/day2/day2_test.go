package day2

import (
	"testing"
)

func TestIsSafeLevelPart1Ok(t *testing.T) {
	reports := map[string]bool{
		"7 6 4 2 1": true,
		"1 2 7 8 9": false,
		"9 7 6 2 1": false,
		"1 3 2 4 5": false,
		"8 6 4 4 1": false,
		"1 3 6 7 9": true,
	}

	for report, expectedResult := range reports {
		var result, err = IsSafeLevel(report, 1, 3, 0)
		if result != expectedResult {
			print(report)
			t.Fatalf(`Is safe: expected: %t got: %t, err: %d`, expectedResult, result, err)
		}
	}

}

func TestIsSafeLevelPart2Ok(t *testing.T) {
	reports := map[string]bool{
		"7 6 4 2 1": true,
		"1 2 7 8 9": false,
		"9 7 6 2 1": false,
		"1 3 2 4 5": true,
		"8 6 4 4 1": true,
		"1 3 6 7 9": true,
	}

	for report, expectedResult := range reports {
		var result, err = IsSafeLevel(report, 1, 3, 1)
		if result != expectedResult {
			print(report)
			t.Fatalf(`Is safe: expected: %t got: %t, err: %d`, expectedResult, result, err)
		}
	}
}

func TestIsSafeEdgeCases(t *testing.T) {
	reports := map[string]bool{
		"2 1 3 5 7":     true,
		"1 3 5 7 5":     true,
		"2 2 2 2 2":     false,
		"1 3 2 2 7":     false,
		"1 3 1 3 1":     false,
		"1 1 2 1 1":     false,
		"1 2 2 1 1":     false,
		"1 2 3 4 5":     true,
		"5 2 3 4 5":     true,
		"20 16 12 10 8": false,
		"6 6 7 8 7 6":   false,
	}

	for report, expectedResult := range reports {
		var result, err = IsSafeLevel(report, 1, 3, 1)
		if result != expectedResult {
			print(report)
			t.Fatalf(`Is safe: expected: %t got: %t, err: %d`, expectedResult, result, err)
		}
	}

}

func TestPart1(t *testing.T) {
	testInput := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	var expectedResult int64 = 2

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
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	var expectedResult int64 = 4

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
	}
}
