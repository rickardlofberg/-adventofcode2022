package day7

import (
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

func ParseLine(line string) (int64, []int64) {
	result_numbers := strings.Split(line, ":")
	expeceted_result := StringToInt64OrPanic(result_numbers[0])
	str_numbers := strings.Fields(result_numbers[1])
	numbers := ToInts64OrPanic(str_numbers)
	return expeceted_result, numbers
}

func GetCombinations(n1 int64, n2 int64) []int64 {
	return []int64{
		n1 + n2,
		n1 * n2,
	}
}

func GetCombinations2(n1 int64, n2 int64) []int64 {
	s1 := strconv.FormatInt(n1, 10)
	s2 := strconv.FormatInt(n2, 10)
	combination := strings.Join([]string{s1, s2}, "")
	combination_int := StringToInt64OrPanic(combination)

	return []int64{
		n1 + n2,
		n1 * n2,
		combination_int,
	}
}

func FindSolution(toFind int64, numbers []int64) int64 {
	last_result := []int64{numbers[0]}
	for idx := range numbers[1:] {
		new_results := []int64{}
		for _, result := range last_result {
			new_combination := GetCombinations(result, numbers[idx+1])
			new_results = append(new_results, new_combination...)
		}
		last_result = new_results
	}
	isSolved := slices.Contains(last_result, toFind)
	if isSolved {
		return toFind
	}
	return 0
}

func FindSolution2(toFind int64, numbers []int64) int64 {
	last_result := []int64{numbers[0]}
	for idx := range numbers[1:] {
		new_results := []int64{}
		for _, result := range last_result {
			new_combination := GetCombinations2(result, numbers[idx+1])
			new_results = append(new_results, new_combination...)
		}
		last_result = new_results
	}
	isSolved := slices.Contains(last_result, toFind)
	if isSolved {
		return toFind
	}
	return 0
}

func Solve(input []string) int64 {
	var result int64 = 0
	for _, row := range input {
		expected_result, numbers := ParseLine(row)
		result += FindSolution(expected_result, numbers)
	}
	return result
}

func Solve2(input []string) int64 {
	var result int64 = 0
	for _, row := range input {
		expected_result, numbers := ParseLine(row)
		result += FindSolution2(expected_result, numbers)
	}
	return result
}

func Part1(input []string) (int64, error) {
	return Solve(input), nil
}

func Part2(input []string) (int64, error) {
	return Solve2(input), nil

}
