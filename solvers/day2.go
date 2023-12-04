package solvers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var gameRegex = regexp.MustCompile(`^Game \d+: (.*?)$`)
var redRegex = regexp.MustCompile(`(\d+) red`)
var greenRegex = regexp.MustCompile(`(\d+) green`)
var blueRegex = regexp.MustCompile(`(\d+) blue`)

const (
	red   int = 12
	green int = 13
	blue  int = 14
)

func getAmount(s string, r *regexp.Regexp) int {
	if match := r.FindStringSubmatch(s); len(match) > 0 {
		amount, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}
		return amount
	}
	return 0
}

func possibleGame(s string) bool {
	gameString := gameRegex.FindStringSubmatch(s)[1]
	games := strings.Split(gameString, "; ")
	for _, game := range games {
		if getAmount(game, redRegex) > red {
			return false
		}
		if getAmount(game, greenRegex) > green {
			return false
		}
		if getAmount(game, blueRegex) > blue {
			return false
		}
	}
	return true
}

func minimumColors(s string) (int, int, int) {
	gameString := gameRegex.FindStringSubmatch(s)[1]
	games := strings.Split(gameString, "; ")
	mr := 0
	mg := 0
	mb := 0
	for _, game := range games {
		mr = max(mr, getAmount(game, redRegex))
		mg = max(mg, getAmount(game, greenRegex))
		mb = max(mb, getAmount(game, blueRegex))
	}
	return mr, mg, mb
}

func Day2(input string) (string, string) {
	lines := strings.Split(input, "\r\n")

	p1 := 0
	p2 := 0
	for i, line := range lines {
		if possibleGame(line) {
			p1 += i + 1
		}
		mr, mg, mb := minimumColors(line)
		p2 += mr * mg * mb
	}

	return strconv.Itoa(p1), strconv.Itoa(p2)
}
