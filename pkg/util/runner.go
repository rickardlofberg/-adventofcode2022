package util

func SolveProblem(day int, part int) (int64, error) {
	inputPath := GetInputPath(day)
	solver := GetSolverString(day, part)

	input, err := ReadInputFromFile(inputPath)
	if err != nil {
		return int64(0), err
	}

	solution, err := GetSolver(solver)(input)
	if err != nil {
		return int64(0), err
	}

	return solution, nil
}