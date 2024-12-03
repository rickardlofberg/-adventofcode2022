package day3

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1(input []string) (int64, error) {
	var result int64 = 0
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	inData := strings.Join(input, "")
	var res = r.FindAllString(inData, -1)
	for _, m := range res {
		integers := strings.Split(m, ",")
		if len(integers) != 2 {
			return -1, errors.New("bad input for numbers")
		}
		leftString := integers[0][4:]
		rightString := integers[1]
		rightString = rightString[:len(rightString)-1]

		leftNumber, err := strconv.ParseInt(leftString, 10, 64)
		if err != nil {
			fmt.Println(leftString)
			return -1, errors.New("bad integer")
		}
		rightNumber, err := strconv.ParseInt(rightString, 10, 64)
		if err != nil {
			fmt.Println(rightString)
			return -1, errors.New("bad integer")
		}
		result += leftNumber * rightNumber
	}
	return result, nil
}

func Part2(input []string) (int64, error) {
	var result int64 = 0
	r := regexp.MustCompile(`(do\(\)|don\'t\(\)|mul\([0-9]{1,3},[0-9]{1,3}\))`)
	enabled := true
	inData := strings.Join(input, "")
	var res = r.FindAllString(inData, -1)
	for _, m := range res {
		if m == "do()" {
			enabled = true
			continue
		} else if m == "don't()" {
			enabled = false
		}

		if !enabled {
			continue
		}

		integers := strings.Split(m, ",")
		if len(integers) != 2 {
			return -1, errors.New("bad input for numbers")
		}
		leftString := integers[0][4:]
		rightString := integers[1]
		rightString = rightString[:len(rightString)-1]

		leftNumber, err := strconv.ParseInt(leftString, 10, 64)
		if err != nil {
			fmt.Println(leftString)
			return -1, errors.New("bad integer")
		}
		rightNumber, err := strconv.ParseInt(rightString, 10, 64)
		if err != nil {
			fmt.Println(rightString)
			return -1, errors.New("bad integer")
		}
		result += leftNumber * rightNumber
	}
	return result, nil
}
