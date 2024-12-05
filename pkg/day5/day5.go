package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func GetRule(line string) (string, string) {
	arr := strings.Split(line, "|")
	if len(arr) != 2 {
		fmt.Println("Bad rule line. could not parse")
	}
	return arr[0], arr[1]
}

func ReduceRules(input string, rules map[string][]string) map[string][]string {
	newRules := make(map[string][]string)
	arr := strings.Split(input, ",")

	for key, beforeRule := range rules {
		reducedBefore := []string{}
		for _, before := range beforeRule {
			keepBefore := slices.Contains(arr, before)
			if keepBefore {
				reducedBefore = append(reducedBefore, before)
			}
		}
		if len(reducedBefore) > 0 {
			newRules[key] = reducedBefore
		}
	}
	return newRules
}

func CheckPage(rules map[string][]string, pages []string) int64 {
	hasSeen := []string{}

	for _, page := range pages {
		mustHaveSeen := rules[page]
		for _, b := range mustHaveSeen {
			seen := slices.Contains(hasSeen, b)
			if !seen {
				return 0
			}
		}
		hasSeen = append(hasSeen, page)
	}
	middleIdx := int((len(pages) - 1) / 2)
	middle := pages[middleIdx]
	middleNumber, err := strconv.ParseInt(middle, 10, 64)
	if err != nil {
		fmt.Println("Not able to find middle", middle, pages)
	}

	return middleNumber
}

func RemoveFirstInstance(pages []string, toRemove string) []string {
	newPages := []string{}
	for idx, page := range pages {
		if page == toRemove {
			if idx+1 != len(pages) {
				remaningPages := pages[idx+1:]
				newPages = slices.Concat(newPages, remaningPages)
				return newPages
			}
			return newPages
		}
		newPages = append(newPages, page)
	}
	return newPages
}

func FixPage(rules map[string][]string, pages []string) []string {
	result := CheckPage(rules, pages)
	isValid := 0 != result
	if isValid {
		// fmt.Println("Good", pages)
		return pages
	}

	hasSeen := []string{}
	newPages := []string{}
	for idx, page := range pages {
		mustHaveSeen := rules[page]
		for _, b := range mustHaveSeen {
			seen := slices.Contains(hasSeen, b)
			if !seen {
				// Add expected page in front
				newPages = append(newPages, b)
				// Add next page after (according to rule)
				newPages = append(newPages, page)

				// Remove the bad page from remaining part
				removedBadPage := RemoveFirstInstance(pages[idx+1:], b)
				newPages := slices.Concat(newPages, removedBadPage)

				// Recursivly fix
				return FixPage(rules, newPages)
			}
		}
		hasSeen = append(hasSeen, page)
	}
	return newPages
}

func Solve(input []string, part int64) int64 {
	isRules := true
	ruleMap := make(map[string][]string)
	var result int64 = 0
	for _, line := range input {
		if line == "" {
			isRules = false
			continue
		}

		if isRules {
			before, after := GetRule(line)
			ruleMap[after] = append(ruleMap[after], before)
		} else {
			lineRules := ReduceRules(line, ruleMap)
			pages := strings.Split(line, ",")
			if part == 1 {
				result += CheckPage(lineRules, pages)

			} else {
				bad := 0 == CheckPage(lineRules, pages)
				if bad {
					pages = FixPage(lineRules, pages)
					result += CheckPage(lineRules, pages)
				}
			}
		}
	}
	return result
}

func Part1(input []string) (int64, error) {
	return Solve(input, 1), nil
}

func Part2(input []string) (int64, error) {
	return Solve(input, 2), nil
}
