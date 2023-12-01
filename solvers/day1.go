package solvers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var letterRegex = regexp.MustCompile(`[A-Za-z]`)

var stringToNumber = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func extractNumbers(s string) int {
	numbers := letterRegex.ReplaceAllString(s, "")

	if numbers == "" {
		return 0
	}

	n, err := strconv.Atoi(string(numbers[0]) + string(numbers[len(numbers)-1]))
	if err != nil {
		log.Fatal(err)
	}

	return n
}

func extractNumberswithStrings(s string) int {
	for k, v := range stringToNumber {
		s = strings.ReplaceAll(s, k, v)
	}
	return extractNumbers(s)
}

func Day1(input string) (string, string) {
	lines := strings.Split(input, "\r\n")

	p1 := 0
	p2 := 0
	for _, line := range lines {
		p1 += extractNumbers(line)
		p2 += extractNumberswithStrings(line)
	}

	return strconv.Itoa(p1), strconv.Itoa(p2)
}
