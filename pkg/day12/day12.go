package day12

import "fmt"

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

type cell struct {
	perimitorCount int
	value          string
}

func ParseInput(input []string) map[coordinate]cell {
	coordinates := map[coordinate]cell{}

	for idxY := range input {
		for idxX := range input {
			cell := cell{value: string(input[idxY][idxX]), perimitorCount: 0}
			corr := coordinate{x: idxX, y: idxY}
			coordinates[corr] = cell
		}
	}
	return coordinates
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

func GetAdjecentCooridinates(corr coordinate) []coordinate {
	south := coordinate{x: corr.x, y: corr.y - 1}
	north := coordinate{x: corr.x, y: corr.y + 1}
	east := coordinate{x: corr.x + 1, y: corr.y}
	west := coordinate{x: corr.x - 1, y: corr.y}
	return []coordinate{south, north, east, west}
}

func GetCellPerimitor(corr coordinate, in map[coordinate]cell, bounds bounds) cell {
	adjecent := GetAdjecentCooridinates(corr)
	validAdjecentCount := 0
	currentCellValue := in[corr].value

	for _, adj := range adjecent {
		isValid := true
		isValid = isValid && InBounds(adj, bounds)
		isValid = isValid && in[adj].value == currentCellValue
		if !isValid {
			validAdjecentCount += 1
		}
	}
	return cell{value: currentCellValue, perimitorCount: validAdjecentCount}
}

func UpdatePerimitor(in map[coordinate]cell, bounds bounds) map[coordinate]cell {
	coordinates := map[coordinate]cell{}

	for idxY := range bounds.yMax + 1 {
		for idxX := range bounds.xMax + 1 {
			corr := coordinate{x: idxX, y: idxY}
			newCell := GetCellPerimitor(corr, in, bounds)
			coordinates[corr] = newCell
		}
	}
	return coordinates
}

func PairWiseComparission(expectedVal string, in map[coordinate]cell, c1 coordinate, c2 coordinate) int {
	cell1 := in[c1]
	cell2 := in[c2]

	// If one of these have the same value it is not a corner
	if cell1.value == expectedVal || cell2.value == expectedVal {
		return 0
	}

	return 1
}

func SimpleCornerCount(corr coordinate, in map[coordinate]cell) int {
	corners := 0
	cellValue := in[corr]

	up := coordinate{x: corr.x, y: corr.y - 1}
	down := coordinate{x: corr.x, y: corr.y + 1}
	left := coordinate{x: corr.x - 1, y: corr.y}
	right := coordinate{x: corr.x + 1, y: corr.y}

	corners += PairWiseComparission(cellValue.value, in, up, right)
	corners += PairWiseComparission(cellValue.value, in, right, down)
	corners += PairWiseComparission(cellValue.value, in, down, left)
	corners += PairWiseComparission(cellValue.value, in, left, up)
	// fmt.Println(corr, corners)
	return corners
}

func InwardsCorner(org cell, c1 cell, c2 cell, c3 cell) bool {
	isInwards := true
	isInwards = isInwards && org.value == c1.value
	isInwards = isInwards && org.value == c2.value
	isInwards = isInwards && org.value != c3.value
	return isInwards
}

func IsInwardsCorner(corr coordinate, in map[coordinate]cell) int {
	up := in[coordinate{x: corr.x, y: corr.y - 1}]
	down := in[coordinate{x: corr.x, y: corr.y + 1}]
	left := in[coordinate{x: corr.x - 1, y: corr.y}]
	right := in[coordinate{x: corr.x + 1, y: corr.y}]
	upRight := in[coordinate{x: corr.x + 1, y: corr.y - 1}]
	downRight := in[coordinate{x: corr.x + 1, y: corr.y + 1}]
	downLeft := in[coordinate{x: corr.x - 1, y: corr.y + 1}]
	upLeft := in[coordinate{x: corr.x - 1, y: corr.y - 1}]

	upRightCorner := InwardsCorner(in[corr], up, right, upRight)
	downRightCorner := InwardsCorner(in[corr], down, right, downRight)
	downLeftCorner := InwardsCorner(in[corr], down, left, downLeft)
	upLeftCorner := InwardsCorner(in[corr], up, left, upLeft)

	corners := 0
	if upRightCorner {
		corners += 1
	}
	if downRightCorner {
		corners += 1
	}
	if downLeftCorner {
		corners += 1
	}
	if upLeftCorner {
		corners += 1
	}
	return corners
}

func IsCorner(corr coordinate, in map[coordinate]cell, bounds bounds) cell {
	cornerCount := SimpleCornerCount(corr, in)
	cornerCount += IsInwardsCorner(corr, in)

	fmt.Println(in[corr].value, corr, cornerCount)

	return cell{value: in[corr].value, perimitorCount: cornerCount}
}

func UpdatePerimitorToCorner(in map[coordinate]cell, bounds bounds) map[coordinate]cell {
	coordinates := map[coordinate]cell{}

	for idxY := range bounds.yMax + 1 {
		for idxX := range bounds.xMax + 1 {
			corr := coordinate{x: idxX, y: idxY}
			newCell := IsCorner(corr, in, bounds)
			coordinates[corr] = newCell
		}
	}
	return coordinates
}

func FindArea(corr coordinate, in map[coordinate]cell, bounds bounds, visited map[coordinate]bool) int64 {
	perimitorCount := int64(in[corr].perimitorCount)
	size := int64(1)

	visited[corr] = true
	potentialCells := GetAdjecentCooridinates(corr)

	for len(potentialCells) > 0 {
		toCheckCorr := potentialCells[len(potentialCells)-1:][0]
		toCheck := in[toCheckCorr]
		potentialCells = potentialCells[:len(potentialCells)-1]

		isValid := !visited[toCheckCorr]
		isValid = isValid && InBounds(toCheckCorr, bounds)
		isValid = isValid && toCheck.value == in[corr].value
		if isValid {
			visited[toCheckCorr] = true
			perimitorCount += int64(toCheck.perimitorCount)
			size++
			potentialCells = append(potentialCells, GetAdjecentCooridinates(toCheckCorr)...)
		}
	}
	fmt.Println(in[corr].value, perimitorCount, size)
	return perimitorCount * size
}

func SolvePart1(in map[coordinate]cell, bounds bounds) int64 {
	result := int64(0)
	visited := make(map[coordinate]bool)

	for idxY := range bounds.yMax + 1 {
		for idxX := range bounds.xMax + 1 {
			corr := coordinate{x: idxX, y: idxY}
			if !visited[corr] {
				result += FindArea(corr, in, bounds, visited)
			}
		}
	}
	return result
}

func Part1(input []string) (int64, error) {
	bounds := GetBounds(input)
	board := ParseInput(input)
	board = UpdatePerimitor(board, bounds)
	result := SolvePart1(board, bounds)
	return result, nil
}

func Part2(input []string) (int64, error) {
	bounds := GetBounds(input)
	board := ParseInput(input)
	board = UpdatePerimitorToCorner(board, bounds)
	result := SolvePart1(board, bounds)
	return result, nil
}
