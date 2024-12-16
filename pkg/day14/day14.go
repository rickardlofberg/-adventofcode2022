package day14

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Robot struct {
	x     int
	y     int
	xVelo int
	yVelo int
}

func parseInput(in []string) []Robot {
	robots := []Robot{}
	for _, line := range in {
		line := strings.Split(line, " ")
		left := line[0][2:]
		right := line[1][2:]
		position := strings.Split(left, ",")
		velocity := strings.Split(right, ",")

		xPos, _ := strconv.Atoi(position[0])
		yPos, _ := strconv.Atoi(position[1])
		xVel, _ := strconv.Atoi(velocity[0])
		yVel, _ := strconv.Atoi(velocity[1])

		robot := Robot{
			x:     xPos,
			y:     yPos,
			xVelo: xVel,
			yVelo: yVel,
		}

		robots = append(robots, robot)
	}
	return robots
}

func updateRobotPosition(robots []Robot, steps, height, width int) []Robot {
	newRobots := []Robot{}
	for _, robot := range robots {
		newX := robot.x
		newY := robot.y
		for range steps {
			newX = (newX + width + robot.xVelo) % width
			newY = (newY + height + robot.yVelo) % height
		}

		newRobot := Robot{x: newX, y: newY, xVelo: robot.xVelo, yVelo: robot.yVelo}
		newRobots = append(newRobots, newRobot)
	}
	return newRobots
}

func getQudrants(robots []Robot, height, width int) int64 {
	var topLeft int64
	var topRight int64
	var bottomLeft int64
	var bottomRight int64
	middleHeight := height / 2
	middleWidth := width / 2

	for _, robot := range robots {
		fmt.Println(robot)
		if robot.x < middleWidth && robot.y < middleHeight {
			topLeft++
		} else if robot.x > middleWidth && robot.y < middleHeight {
			topRight++
		} else if robot.x < middleWidth && robot.y > middleHeight {
			bottomLeft++
		} else if robot.x > middleWidth && robot.y > middleHeight {
			bottomRight++
		}
	}
	fmt.Println(topLeft, topRight, bottomLeft, bottomRight)
	return topLeft * topRight * bottomLeft * bottomRight
}

func Part1(input []string) (int64, error) {
	var height int
	var width int

	if os.Getenv("test") == "true" {
		height = 7
		width = 11
	} else {
		height = 103
		width = 101
	}
	robots := parseInput(input)
	robots = updateRobotPosition(robots, 100, height, width)
	result := getQudrants(robots, height, width)
	return result, nil
}

func Part2(input []string) (int64, error) {
	return -1, errors.New("not yet implemented")
}
