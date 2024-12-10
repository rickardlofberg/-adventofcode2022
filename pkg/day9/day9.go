package day9

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func StringToInt64OrPanic(s string) int64 {
	number, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return number
}

func SolvePart1(data []int64, empty []int64) int64 {
	result := int64(0)
	idxPointer := int64(0)
	valueIdxPointer := int64(0)
	popIndex := int64(len(data) - 1)
	popsLeft := data[popIndex]
	fillEmpty := false
	// var buffer bytes.Buffer
	for idxPointer < popIndex {
		if !fillEmpty {
			for range data[idxPointer] {
				// buffer.WriteString(strconv.FormatInt(idxPointer, 10))
				result += valueIdxPointer * idxPointer
				valueIdxPointer += 1
			}
			fillEmpty = true
		} else {
			for range empty[idxPointer] {
				if popsLeft <= 0 {
					popIndex -= 1
					if idxPointer >= popIndex {
						break
					}
					popsLeft = data[popIndex]
				}
				// buffer.WriteString(strconv.FormatInt(popIndex, 10))
				result += valueIdxPointer * popIndex
				popsLeft -= 1
				valueIdxPointer += 1
			}
			fillEmpty = false
			idxPointer += 1
		}
	}

	for range popsLeft {
		if popsLeft <= 0 {
			popIndex -= 1
			if idxPointer >= popIndex {
				break
			}
			popsLeft = data[popIndex]
		}
		// buffer.WriteString(strconv.FormatInt(popIndex, 10))

		result += valueIdxPointer * popIndex
		popsLeft -= 1
		valueIdxPointer += 1
	}

	// fmt.Println(buffer.String())
	return result
}

func Part1(input []string) (int64, error) {
	row := input[0]
	data := []int64{}
	empty := []int64{}
	for idx := range row {
		numberOf := StringToInt64OrPanic(string(row[idx]))
		if idx%2 == 0 {
			data = append(data, numberOf)
		} else {
			empty = append(empty, numberOf)
		}
	}

	result := SolvePart1(data, empty)

	return result, nil
}

type DataBox struct {
	size  int64
	value int64
}

func GetValue(boxes []DataBox) int64 {
	result := int64(0)
	idx := int64(0)
	for _, box := range boxes {
		for range box.size {
			if box.value != -1 {
				result += box.value * idx
			}
			idx += 1
		}
	}
	return result
}

func GetNextEmptyBox(startIdx int, boxes []DataBox) int {
	for idx, box := range boxes {
		if box.value == 0 {
			return startIdx + idx + 1
		}
	}
	return len(boxes) + startIdx
}

func GetPreviousDataBox(boxes []DataBox) int {
	idxOffest := 0
	maxOffset := len(boxes) - 1
	for idx := range len(boxes) - 1 {
		idxOffest = maxOffset - idx
		if idxOffest >= 0 {
			potentialBox := boxes[idxOffest]
			if potentialBox.value != -1 {
				return idxOffest
			}
		}
	}
	return idxOffest
}

func PrintDataBox(boxes []DataBox) {
	var sb strings.Builder
	for _, box := range boxes[1:] {
		for range box.size {
			if box.value != -1 {
				str := strconv.FormatInt(box.value, 10)
				sb.WriteString(str)
			} else {
				sb.WriteString(".")
			}
		}
	}
	fmt.Println(sb.String())
}

func GetSize(boxes []DataBox) int64 {
	res := int64(0)
	for _, box := range boxes {
		res += box.size
	}
	return res
}

func SovlePart2(boxes []DataBox) int64 {
	toMoveDataIdx := len(boxes) - 1
	for toMoveDataIdx >= 1 {
		for emptyStartIdx, box := range boxes[:toMoveDataIdx] {
			// Skip non-empty boxes
			if box.value != -1 {
				continue
			}
			if boxes[toMoveDataIdx].value == -1 {
				break
			}

			moveToBox := boxes[emptyStartIdx]
			boxToMove := boxes[toMoveDataIdx]

			if moveToBox.size >= boxToMove.size {
				// Swap places
				boxes[emptyStartIdx] = boxToMove
				boxes[toMoveDataIdx] = DataBox{size: boxToMove.size, value: -1}

				// Extra space?
				if moveToBox.size > boxToMove.size {
					extraSpace := moveToBox.size - boxToMove.size
					nextBox := boxes[emptyStartIdx+1]
					if nextBox.value == -1 {
						// Add space to empty box
						boxes[emptyStartIdx+1].size = nextBox.size + extraSpace
					} else {
						// Create new empty box
						emptyBox := DataBox{size: extraSpace, value: -1}
						boxes = slices.Insert(boxes, emptyStartIdx+1, emptyBox)
					}
				}
				// Break inner loop find next box to try to swap
				break
			}
		}
		toMoveDataIdx = GetPreviousDataBox(boxes[:toMoveDataIdx])
	}
	fmt.Println()
	size := GetSize(boxes)
	fmt.Println("after", len(boxes), size)
	// PrintDataBox(boxes)

	return GetValue(boxes)
}

func Part2(input []string) (int64, error) {
	row := input[0]
	boxes := []DataBox{}
	var box DataBox
	idxVal := 0
	for idx := range row {
		size := StringToInt64OrPanic(string(row[idx]))
		if idx%2 == 0 {
			box = DataBox{size: size, value: int64(idxVal)}
			idxVal += 1
		} else {
			// -1 to represent empty box
			box = DataBox{size: size, value: -1}
		}
		boxes = append(boxes, box)
	}
	size := GetSize(boxes)
	fmt.Println("before", len(boxes), size)
	result := SovlePart2(boxes)

	return result, nil
}
