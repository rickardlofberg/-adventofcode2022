package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type PrizeCoordinate struct {
	x int
	y int
}

type Distance struct {
	xDist int
	yDist int
}

type Presses struct {
	aPress int
	bPress int
}

type Machine struct {
	aButtonMove Distance
	bButtonMove Distance
	prize       PrizeCoordinate
}

func ParseInput(input []string) []Machine {
	machines := []Machine{}
	distanceA := Distance{}
	distanceB := Distance{}
	prize := PrizeCoordinate{}

	for _, line := range input {
		if len(line) == 0 {
			machine := Machine{aButtonMove: distanceA, bButtonMove: distanceB, prize: prize}
			machines = append(machines, machine)
			fmt.Println(machine)
			continue
		} else if line[:6] == "Button" {
			coord := line[9:]
			coors := strings.Split(coord, ",")
			corr1, corr2 := coors[0][3:], coors[1][3:]
			intDistX, _ := strconv.Atoi(corr1)
			intDistY, _ := strconv.Atoi(corr2)
			if string(line[7]) == "A" {
				distanceA = Distance{xDist: intDistX, yDist: intDistY}
			} else if string(line[7]) == "B" {
				distanceB = Distance{xDist: intDistX, yDist: intDistY}

			}
		} else if line[:5] == "Prize" {
			coord := line[8:]
			coors := strings.Split(coord, ",")
			corr1, corr2 := coors[0][1:], coors[1][3:]
			xCorr, _ := strconv.Atoi(corr1)
			yCorr, _ := strconv.Atoi(corr2)
			prize = PrizeCoordinate{x: xCorr, y: yCorr}
		}
	}
	return machines
}

func FindCombinations(machine Machine) []Presses {
	presses := []Presses{}
	goal := machine.prize

	for aPresses := 0; aPresses <= 100; aPresses++ {
		for bPresses := 0; bPresses <= 100; bPresses++ {
			xPositionA := aPresses * machine.aButtonMove.xDist
			xPositionB := bPresses * machine.bButtonMove.xDist
			yPositionA := aPresses * machine.aButtonMove.yDist
			yPositionB := bPresses * machine.bButtonMove.yDist
			xPos := xPositionA + xPositionB
			yPos := yPositionA + yPositionB

			if xPos == goal.x && yPos == goal.y {
				press := Presses{aPress: aPresses, bPress: bPresses}
				presses = append(presses, press)
			}
		}
	}
	return presses
}

func ReturnCheapestPrice(presses []Presses) int64 {
	result := int64(0)
	aCost := 3
	bCost := 1
	for _, press := range presses {
		a := press.aPress * aCost
		b := press.bPress * bCost
		cost := a + b
		if result == 0 {
			result = int64(cost)
		} else if int64(cost) < result {
			result = int64(cost)

		}
	}
	return result
}

func SolvePart1(machines []Machine) int64 {
	result := int64(0)
	for _, machine := range machines {
		solution := FindCombinations(machine)
		result += ReturnCheapestPrice(solution)
	}
	return result
}

func Part1(input []string) (int64, error) {
	machines := ParseInput(input)
	result := SolvePart1(machines)
	return result, nil
}

func Part2(input []string) (int64, error) {
	machines := ParseInput(input)
	result := SolvePart1(machines)
	result = -1
	return result, nil
}
