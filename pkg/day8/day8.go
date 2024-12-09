package day8

import (
	"fmt"
	"slices"
)

type coordinate struct {
	x int
	y int
}

type bounds struct {
	xMin int
	yMin int
	xMax int
	yMax int
}

func abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func GetBounds(arr []string) bounds {
	yMax := len(arr) - 1
	xMax := len(arr[0]) - 1
	return bounds{xMin: 0, yMin: 0, xMax: xMax, yMax: yMax}
}

func InBounds(pos coordinate, bounds bounds) bool {
	if pos.x < bounds.xMin || pos.x > bounds.xMax {
		return false
	} else if pos.y < bounds.yMin || pos.y > bounds.yMax {
		return false
	}
	return true
}

func ParseInput(input []string) (map[string][]coordinate, []coordinate) {
	antena_to_coordinates := map[string][]coordinate{}
	occupied_cells := []coordinate{}

	for idxY := range input {
		for idxX := range input {
			cell := string(input[idxY][idxX])
			if cell != "." {
				antena_to_coordinates[cell] = append(antena_to_coordinates[cell], coordinate{x: idxX, y: idxY})
				occupied_cells = append(occupied_cells, coordinate{x: idxX, y: idxY})
			}
		}
	}

	return antena_to_coordinates, occupied_cells
}

func GetAntiNodes(n1 coordinate, n2 coordinate) []coordinate {
	xDistance := abs(n1.x - n2.x)
	yDistance := abs(n1.y - n2.y)
	var x1 int
	var x2 int
	var y1 int
	var y2 int

	if n1.x <= n2.x {
		x1 = n1.x - xDistance
		x2 = n2.x + xDistance
	} else {
		x1 = n1.x + xDistance
		x2 = n2.x - xDistance
	}

	if n1.y <= n2.y {
		y1 = n1.y - yDistance
		y2 = n2.y + yDistance
	} else {
		y1 = n1.y + yDistance
		y2 = n2.y - yDistance
	}

	return []coordinate{
		{x: x1, y: y1},
		{x: x2, y: y2},
	}
}

func GetAntiNodesCoordinates(nodes []coordinate) []coordinate {
	antiNodes := []coordinate{}

	for idx, n1 := range nodes[:len(nodes)-1] {
		for _, n2 := range nodes[idx+1:] {
			newNodes := GetAntiNodes(n1, n2)
			for _, node := range newNodes {
				hasSeen := slices.Contains(antiNodes, node)
				if !hasSeen {
					antiNodes = append(antiNodes, node)
				}
			}
		}
	}
	return antiNodes
}

func GetAllAntiNodes(antena_map map[string][]coordinate) []coordinate {
	antiNodes := []coordinate{}
	for _, coordinates := range antena_map {
		new_coordinates := GetAntiNodesCoordinates(coordinates)
		for _, cord := range new_coordinates {
			exists := slices.Contains(antiNodes, cord)
			if !exists {
				antiNodes = append(antiNodes, cord)
			}
		}
	}
	return antiNodes
}

func GetAntiNodes2(n1 coordinate, n2 coordinate, bounds bounds) []coordinate {
	newNodes := []coordinate{}

	x := n1.x - n2.x
	y := n1.y - n2.y

	newNode := coordinate{x: n1.x + x, y: n1.y + y}
	for InBounds(newNode, bounds) {
		newNodes = append(newNodes, newNode)
		newNode = coordinate{x: newNode.x + x, y: newNode.y + y}
	}

	x = n2.x - n1.x
	y = n2.y - n1.y
	newNode = coordinate{x: n2.x + x, y: n2.y + y}
	for InBounds(newNode, bounds) {
		newNodes = append(newNodes, newNode)
		newNode = coordinate{x: newNode.x + x, y: newNode.y + y}
	}
	return newNodes
}

func GetAntiNodesCoordinates2(nodes []coordinate, bounds bounds) []coordinate {
	antiNodes := []coordinate{}

	for idx, n1 := range nodes[:len(nodes)-1] {
		for _, n2 := range nodes[idx+1:] {
			newNodes := GetAntiNodes2(n1, n2, bounds)
			for _, node := range newNodes {
				hasSeen := slices.Contains(antiNodes, node)
				if !hasSeen {
					antiNodes = append(antiNodes, node)
				}
			}
		}
	}
	return antiNodes
}

func GetAllAntiNodes2(antena_map map[string][]coordinate, bounds bounds) []coordinate {
	antiNodes := []coordinate{}
	for _, coordinates := range antena_map {
		new_coordinates := GetAntiNodesCoordinates2(coordinates, bounds)
		for _, cord := range new_coordinates {
			exists := slices.Contains(antiNodes, cord)
			if !exists {
				antiNodes = append(antiNodes, cord)
			}
		}
	}
	return antiNodes
}

func GetAllowedAntiNodes(bounds bounds, occupiedCells []coordinate, antiNodes []coordinate) int64 {
	var result int64 = 0
	for _, node := range antiNodes {
		inBounds := InBounds(node, bounds)
		if !inBounds {
			continue
		}
		result += 1
	}
	return result
}

func GetAllowedAntiNodes2(bounds bounds, occupiedCells []coordinate, antiNodes []coordinate) int64 {
	var result int64 = int64(len(occupiedCells))

	fmt.Println("Occupoied cells", len(occupiedCells))
	fmt.Println("Antinodes cells", len(antiNodes))
	for _, node := range antiNodes {
		inBounds := InBounds(node, bounds)
		if !inBounds {
			continue
		}
		isOccupied := slices.Contains(occupiedCells, node)
		if !isOccupied {
			result += 1
			occupiedCells = append(occupiedCells, node)
		}
	}
	return result
}

func Part1(input []string) (int64, error) {
	bounds := GetBounds(input)
	antenaMap, occupiedCells := ParseInput(input)
	antiNodes := GetAllAntiNodes(antenaMap)
	result := GetAllowedAntiNodes(bounds, occupiedCells, antiNodes)

	return result, nil
}

func Part2(input []string) (int64, error) {
	bounds := GetBounds(input)
	antenaMap, occupiedCells := ParseInput(input)
	antiNodes := GetAllAntiNodes2(antenaMap, bounds)
	result := GetAllowedAntiNodes2(bounds, occupiedCells, antiNodes)

	return result, nil
}
