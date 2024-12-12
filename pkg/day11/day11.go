package day11

import (
	"strconv"
	"strings"
)

type cacheKey struct {
	value      string
	iterations int64
}

var cache = map[cacheKey]int64{}

func RunIterations(stones []string, iterations int) []string {
	if iterations == 0 {
		return stones
	}
	newStones := []string{}
	for _, stone := range stones {
		if stone == "0" {
			newStones = append(newStones, "1")
		} else if len(stone)%2 == 0 {
			middle := len(stone) / 2
			firstHalf := stone[:middle]
			secondHalf := stone[middle:]
			secondHalf = strings.TrimLeft(secondHalf, "0")
			if secondHalf == "" {
				secondHalf = "0"
			}
			newStones = append(newStones, firstHalf)
			newStones = append(newStones, secondHalf)
		} else {
			number, _ := strconv.Atoi(stone)
			number = number * 2024
			newStones = append(newStones, strconv.Itoa(number))
		}
	}
	iterations -= 1
	return RunIterations(newStones, iterations)
}

func RunIterationsSingle(stone string, iterations int) int64 {
	if iterations == 0 {
		return 1
	}

	newStones := []string{}
	if stone == "0" {
		newStones = append(newStones, "1")
	} else if len(stone)%2 == 0 {
		middle := len(stone) / 2
		firstHalf := stone[:middle]
		secondHalf := stone[middle:]
		secondHalf = strings.TrimLeft(secondHalf, "0")
		if secondHalf == "" {
			secondHalf = "0"
		}
		newStones = append(newStones, firstHalf)
		newStones = append(newStones, secondHalf)
	} else {
		number, _ := strconv.Atoi(stone)
		number = number * 2024
		newStones = append(newStones, strconv.Itoa(number))
	}
	iterations -= 1

	result := int64(0)
	for _, stone := range newStones {
		result += RunIterationsSingle(stone, iterations)
	}

	return result
}

func RunIterationsSingleCache(stone string, iterations int) int64 {
	if iterations == 0 {
		return 1
	}

	newStones := []string{}
	if stone == "0" {
		newStones = append(newStones, "1")
	} else if len(stone)%2 == 0 {
		middle := len(stone) / 2
		firstHalf := stone[:middle]
		secondHalf := stone[middle:]
		secondHalf = strings.TrimLeft(secondHalf, "0")
		if secondHalf == "" {
			secondHalf = "0"
		}
		newStones = append(newStones, firstHalf)
		newStones = append(newStones, secondHalf)
	} else {
		number, _ := strconv.Atoi(stone)
		number = number * 2024
		newStones = append(newStones, strconv.Itoa(number))
	}
	iterations -= 1

	result := int64(0)
	for _, newStone := range newStones {

		cacheKey := cacheKey{value: newStone, iterations: int64(iterations)}
		cacheVal := cache[cacheKey]
		if cacheVal == 0 {
			numberOfStones := RunIterationsSingleCache(newStone, iterations)
			cache[cacheKey] = numberOfStones
			result += numberOfStones
		} else {
			result += cacheVal
		}
	}

	return result
}

func SolvePart1(stones string, iterations int) int64 {
	stonesArr := strings.Split(stones, " ")
	afterIter := RunIterations(stonesArr, iterations)
	numberOfStones := len(afterIter)
	return int64(numberOfStones)
}

func SolvePart2(stones string, iterations int) int64 {
	size := iterations
	stonesArr := strings.Split(stones, " ")

	result := int64(0)
	for _, stone := range stonesArr {
		result += RunIterationsSingleCache(stone, size)
	}

	return result
}

func Part1(input []string) (int64, error) {
	data := input[0]
	result := SolvePart1(data, 25)
	return result, nil
}

func Part2(input []string) (int64, error) {
	data := input[0]
	result := SolvePart2(data, 75)
	return result, nil
}
