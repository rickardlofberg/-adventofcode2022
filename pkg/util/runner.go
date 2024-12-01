package util

func SolveProblem(day int, part int) (int64, error) {
	inputPath := GetInputPath(day)

	input, err := ReadInputFromFile(inputPath)
	if err != nil {
		return int64(0), err
	}

	solution, err := GetSolver(day, part)(input)
	if err != nil {
		return int64(0), err
	}

	return solution, nil
}
