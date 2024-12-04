package day4

type direction struct {
	x int64
	y int64
}

var searchDirections = []direction{
	direction{x: 0, y: -1},
	direction{x: 0, y: 0},
	direction{x: 0, y: 1},
	direction{x: 1, y: -1},
	direction{x: 1, y: 0},
	direction{x: 1, y: 1},
	direction{x: -1, y: -1},
	direction{x: -1, y: 0},
	direction{x: -1, y: 1},
}

func InputToArray(input []string) [][]string {
	height := len(input)
	width := len(input[0])
	array := make([][]string, height)

	for yIdx := range input {
		row := input[yIdx]
		rowArr := make([]string, width)
		for xIdx := range row {
			char := string(row[xIdx])
			rowArr[xIdx] = char
		}
		array[yIdx] = rowArr
	}
	return array
}

func IsWithinBounds(width int64, height int64, x int64, y int64) bool {
	if x < 0 || x >= width {
		return false
	}
	if y < 0 || y >= height {
		return false
	}
	return true
}

func ContinueSearch(width int64, height int64, x int64, y int64, searchDepth int64, maxDepth int64) bool {
	if searchDepth >= maxDepth {
		return false
	}
	return IsWithinBounds(width, height, x, y)
}

func FindXmas(arr [][]string, width int64, height int64, x int64, y int64, find string) int64 {
	maxSearchDepth := int64(len(find)) - 1
	var xmaxCount int64 = 0

	for _, direction := range searchDirections {
		currentX := x
		currentY := y
		searchDepth := int64(0)

		found := arr[currentY][currentX]
		for ContinueSearch(width, height, currentX, currentY, searchDepth, maxSearchDepth) {
			searchDepth += 1
			currentX += direction.x
			currentY += direction.y
			if IsWithinBounds(width, height, currentX, currentY) {
				found += arr[currentY][currentX]
			}
		}
		if found == find {
			xmaxCount += 1
		}
	}
	return xmaxCount
}

func InSearchBounds(width int64, height int64, x int64, y int64) bool {
	inBounds := true
	inBounds = inBounds && IsWithinBounds(width, height, x-1, y-1)
	inBounds = inBounds && IsWithinBounds(width, height, x-1, y+1)
	inBounds = inBounds && IsWithinBounds(width, height, x+1, y-1)
	inBounds = inBounds && IsWithinBounds(width, height, x+1, y+1)
	return inBounds
}

func FindXmasPart2(arr [][]string, width int64, height int64, x int64, y int64, find string) int64 {
	currentX := x
	currentY := y

	current := arr[currentY][currentX]
	if current == "A" && InSearchBounds(width, height, x, y) {
		NW := arr[currentY-1][currentX-1]
		NE := arr[currentY-1][currentX+1]
		SW := arr[currentY+1][currentX-1]
		SE := arr[currentY+1][currentX+1]
		ok := true
		ok = ok && ((NW == "S" && SE == "M") || (NW == "M" && SE == "S"))
		ok = ok && ((SW == "S" && NE == "M") || (SW == "M" && NE == "S"))
		if ok {
			return 1
		}
	}
	return 0
}

func Part1(input []string) (int64, error) {
	array := InputToArray(input)
	height := len(input)
	width := len(input[0])
	var numberOfXmas int64 = 0

	for yIdx := range height {
		for xIdx := range width {
			numberOfXmas += FindXmas(array, int64(width), int64(height), int64(xIdx), int64(yIdx), "XMAS")
		}
	}
	return numberOfXmas, nil
}

func Part2(input []string) (int64, error) {
	array := InputToArray(input)
	height := len(input)
	width := len(input[0])
	var numberOfXmas int64 = 0

	for yIdx := range height {
		for xIdx := range width {
			numberOfXmas += FindXmasPart2(array, int64(width), int64(height), int64(xIdx), int64(yIdx), "XMAS")
		}
	}
	return numberOfXmas, nil
}
