package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Moonlington/AOC23/solvers"
)

// SolverFunc is a function from the input string to the part 1 and part 2 output string.
type SolverFunc func(string) (string, string)

var dayToSolverFunc map[int]SolverFunc = map[int]SolverFunc{
	1: solvers.Day1,
}

func loadInput(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("inputs/day%d.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("passed %d arguments, want %d", len(os.Args)-1, 1)
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	solver := dayToSolverFunc[day]
	p1, p2 := solver(loadInput(day))

	fmt.Printf("Solving for Day %d\n", day)
	fmt.Printf("  Part 1: %s\n  Part 2: %s\n", p1, p2)
}
