package solvers

// SolverFunc is a function from the input string to the part 1 and part 2 output string.
type SolverFunc func(string) (string, string)

var DayToSolverFunc map[int]SolverFunc = map[int]SolverFunc{
	1: Day1,
	2: Day2,
}
