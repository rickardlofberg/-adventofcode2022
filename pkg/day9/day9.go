package day9

import (
	"errors"
	"strconv"
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

func Part2(input []string) (int64, error) {
	return -1, errors.New("not yet implemented")
}
