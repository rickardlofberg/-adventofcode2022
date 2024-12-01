package util

import (
	"bufio"
	"os"
)

func ReadInputFromFile(fileName string) ([]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return []string{}, err
	}

	defer f.Close()

	var inputArray []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		bits := scanner.Text()
		inputArray = append(inputArray, bits)
	}

	return inputArray, err
}
