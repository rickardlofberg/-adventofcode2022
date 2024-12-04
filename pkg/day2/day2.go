package day2

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

func ToInts64OrPanic(row []string) []int64 {
	var integers []int64
	for _, s := range row {
		integer := StringToInt64OrPanic(s)
		integers = append(integers, integer)
	}
	return integers
}

func GetVariations(numbers []int64) [][]int64 {
	result := make([][]int64, len(numbers)+1)
	result[0] = numbers
	for idx := range numbers {
		var newNumbers []int64
		if idx == 0 {
			newNumbers = numbers[idx+1:]
		} else if idx == len(numbers)-1 {
			newNumbers = numbers[:idx]
		} else {
			right := numbers[:idx]
			left := numbers[idx+1:]
			newNumbers = slices.Concat(right, left)
		}
		result[idx+1] = newNumbers
	}
	return result
}

func IsValid(integers []int64, minStep int64, maxStep int64) bool {
	isAscending := integers[0] < integers[1]
	var length int64 = 0

	for idx, number := range integers[1:] {
		prevNumber := integers[idx]

		if isAscending {
			length = number - prevNumber
		} else {
			length = prevNumber - number
		}
		if !(length >= minStep && length <= maxStep) {
			return false
		}
	}
	return true
}

func IsSafeLevel2(row string, minStep int64, maxStep int64, tolerance int64) (bool, error) {
	numbers := strings.Fields(row)
	integers := ToInts64OrPanic(numbers)
	// fmt.Println(integers)
	variations := GetVariations(integers)
	// fmt.Println(variations)
	fmt.Println()
	for _, variation := range variations {
		isValid := IsValid(variation, minStep, maxStep)
		if isValid {
			return true, nil
		}
	}
	return false, nil
}

func IsSafeLevel(row string, minStep int64, maxStep int64, tolerance int64) (bool, error) {
	numbers := strings.Fields(row)
	integers := ToInts64OrPanic(numbers)
	isValid := IsValid(integers, minStep, maxStep)
	return isValid, nil
}

func Part1(input []string) (int64, error) {
	var safeLevels int64 = 0
	var minStep int64 = 1
	var maxStep int64 = 3
	for _, row := range input {
		var isSafe, _ = IsSafeLevel(row, minStep, maxStep, 0)
		if isSafe {
			safeLevels += 1
		}
	}
	return safeLevels, nil
}

func Part2(input []string) (int64, error) {
	var safeLevels int64 = 0
	var minStep int64 = 1
	var maxStep int64 = 3
	for _, row := range input {
		var isSafe, _ = IsSafeLevel2(row, minStep, maxStep, 1)
		if isSafe {
			safeLevels += 1
		}
	}
	return safeLevels, nil
}
