package day14

import (
	"os"
	"testing"
)

func TestMovement(t *testing.T) {
	testInput := []string{
		"p=2,4 v=2,-3",
	}

	robots := parseInput(testInput)
	robots = updateRobotPosition(robots, 1, 7, 11)

	robot := robots[0]
	expectedX := 4
	if robot.x != expectedX {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedX, robot.x)
	}

	expectedY := 1
	if robot.y != expectedY {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedY, robot.y)
	}
}

func TestMovementSteps2(t *testing.T) {
	testInput := []string{
		"p=2,4 v=2,-3",
	}

	robots := parseInput(testInput)
	robots = updateRobotPosition(robots, 2, 7, 11)

	robot := robots[0]
	expectedX := 6
	if robot.x != expectedX {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedX, robot.x)
	}

	expectedY := 5
	if robot.y != expectedY {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedY, robot.y)
	}
}

func TestMovementDoubleNetativeWrap(t *testing.T) {
	testInput := []string{
		"p=1,1 v=-5,-5",
	}

	robots := parseInput(testInput)
	robots = updateRobotPosition(robots, 1, 5, 5)

	robot := robots[0]
	expectedX := 1
	if robot.x != expectedX {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedX, robot.x)
	}

	expectedY := 1
	if robot.y != expectedY {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedY, robot.y)
	}
}

func TestMovementDoubleNetativeWrapTwo(t *testing.T) {
	testInput := []string{
		"p=1,1 v=-2,-2",
	}

	robots := parseInput(testInput)
	robots = updateRobotPosition(robots, 3, 5, 5)

	robot := robots[0]
	expectedX := 0
	if robot.x != expectedX {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedX, robot.x)
	}

	expectedY := 0
	if robot.y != expectedY {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedY, robot.y)
	}
}

func TestMovementDoublePositiveWrap(t *testing.T) {
	testInput := []string{
		"p=4,4 v=2,2",
	}

	robots := parseInput(testInput)
	robots = updateRobotPosition(robots, 3, 5, 5)

	robot := robots[0]
	expectedX := 0
	if robot.x != expectedX {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedX, robot.x)
	}

	expectedY := 0
	if robot.y != expectedY {
		t.Fatalf(`TestMovement: expected: %d got: %d`, expectedY, robot.y)
	}
}

func TestPart1(t *testing.T) {
	// if os.Getenv("RUN_TEMPLATE") != "true" {
	// 	t.Skip("Skipping template test.")
	// }
	testInput := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	var expectedResult int64 = 12

	os.Setenv("test", "true")
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
