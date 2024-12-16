package day13

import (
	"os"
	"testing"
)

// func TestPart1(t *testing.T) {
// 	testInput := []string{
// 		"Button A: X+94, Y+34",
// 		"Button B: X+22, Y+67",
// 		"Prize: X=8400, Y=5400",
// 		"",
// 		"Button A: X+26, Y+66",
// 		"Button B: X+67, Y+21",
// 		"Prize: X=12748, Y=12176",
// 		"",
// 		"Button A: X+17, Y+86",
// 		"Button B: X+84, Y+37",
// 		"Prize: X=7870, Y=6450",
// 		"",
// 		"Button A: X+69, Y+23",
// 		"Button B: X+27, Y+71",
// 		"Prize: X=18641, Y=10279",
// 		"",
// 	}
// 	var expectedResult int64 = 480

// 	result, err := Part1(testInput)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if result != expectedResult {
// 		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
// 	}
// }

func TestPart2(t *testing.T) {
	if os.Getenv("RUN_TEMPLATE") != "true" {
		t.Skip("Skipping template test.")
	}
	testInput := []string{
		"Button A: X+94, Y+34",
		"Button B: X+22, Y+67",
		"Prize: X=10000000008400, Y=10000000005400",
		"",
		"Button A: X+26, Y+66",
		"Button B: X+67, Y+21",
		"Prize: X=10000000012748, Y=10000000012176",
		"",
		"Button A: X+17, Y+86",
		"Button B: X+84, Y+37",
		"Prize: X=10000000007870, Y=10000000006450",
		"",
		"Button A: X+69, Y+23",
		"Button B: X+27, Y+71",
		"Prize: X=10000000018641, Y=10000000010279",
		"",
	}
	var expectedResult int64 = 0

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
