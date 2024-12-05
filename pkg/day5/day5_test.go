package day5

import (
	"fmt"
	"slices"
	"testing"
)

func TestRemoveFirstInstnaceFirstChar(t *testing.T) {
	toRemove := "a"
	slice := []string{"a", "b", "c"}
	expected := []string{"b", "c"}

	result := RemoveFirstInstance(slice, toRemove)
	good := slices.Equal(expected, result)
	if !good {
		fmt.Println(result)
		fmt.Println(expected)
		t.Fatalf("Error removal")
	}
}

func TestRemoveFirstInstnaceFirstCharDoulbe(t *testing.T) {
	toRemove := "a"
	slice := []string{"a", "b", "c", "a"}
	expected := []string{"b", "c", "a"}

	result := RemoveFirstInstance(slice, toRemove)
	good := slices.Equal(expected, result)
	if !good {
		fmt.Println(result)
		fmt.Println(expected)
		t.Fatalf("Error removal")
	}
}

func TestFixPageLastCahnge(t *testing.T) {
	slice := []string{"61", "13", "29"}
	expected := []string{"61", "29", "13"}
	rules := map[string][]string{
		"13": {"61", "29"},
		"29": {"61"},
		"47": {},
		"53": {"61"},
		"61": {},
		"75": {},
	}

	result := FixPage(rules, slice)
	good := slices.Equal(expected, result)
	if !good {
		fmt.Println(result)
		fmt.Println(expected)
		t.Fatalf("Error removal")
	}
}

func TestPart1(t *testing.T) {
	testInput := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29", // Good
		"97,61,53,29,13", // Good
		"75,29,13",       // Good
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}
	var expectedResult int64 = 143

	result, err := Part1(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part1: expected: %d got: %d`, expectedResult, result)
	}
}

func TestPart2(t *testing.T) {
	testInput := []string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	}
	var expectedResult int64 = 123

	result, err := Part2(testInput)

	if err != nil {
		t.Fatal(err)
	}

	if result != expectedResult {
		t.Fatalf(`Part2: expected: %d got: %d`, expectedResult, result)
	}
}
