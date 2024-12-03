package day3

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testInput := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
	var expectedResult int64 = 161

	result, err := Part1(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	testInput := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
	var expectedResult int64 = 48

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
