package day1

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

func ToIntegers(row string) (int64, int64, error) {
	integers := strings.Fields(row)
	if len(integers) < 2 {
		return -1, -1, errors.New("bad input row")
	}
	left_number, err := strconv.ParseInt(integers[0], 10, 64)
	if err != nil {
		return -1, -1, errors.New("bad integer")
	}
	right_number, err := strconv.ParseInt(integers[1], 10, 64)
	if err != nil {
		return -1, -1, errors.New("bad integer")
	}

	return left_number, right_number, nil
}

func abs(i int64) int64 {
	if i < 0 {
		return i * -1
	}
	return i
}

func Part1(input []string) (int64, error) {
	var left_numbers []int64
	var right_numbers []int64

	for _, row := range input {
		n1, n2, err := ToIntegers(row)
		if err != nil {
			return -1, errors.New("something went wrong")
		}
		left_numbers = append(left_numbers, n1)
		right_numbers = append(right_numbers, n2)
	}

	sort.Slice(left_numbers, func(i, j int) bool {
		return left_numbers[i] < left_numbers[j]
	})
	sort.Slice(right_numbers, func(i, j int) bool {
		return right_numbers[i] < right_numbers[j]
	})

	if len(left_numbers) != len(right_numbers) {
		return -1, errors.New("length of slices not correct")
	}

	var result int64
	result = 0
	for i := range len(left_numbers) {
		result += abs(left_numbers[i] - right_numbers[i])
	}

	return result, nil
}

func Part2(input []string) (int64, error) {
	var left_numbers []int64
	var right_numbers_count = make(map[int64]int64)

	for _, row := range input {
		n1, n2, err := ToIntegers(row)
		if err != nil {
			return -1, errors.New("something went wrong")
		}
		left_numbers = append(left_numbers, n1)
		right_numbers_count[n2] = right_numbers_count[n2] + 1
	}

	var result int64
	result = 0
	for _, val := range left_numbers {
		result += val * right_numbers_count[val]
	}

	return result, nil
}
