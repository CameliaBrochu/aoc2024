package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readProblemInput(day string, test bool) string {
	var sb strings.Builder

	sb.WriteString("inputs/")
	sb.WriteString(day)
	if test {
		sb.WriteString("_test")
	}
	sb.WriteString(".txt")

	data, err := os.ReadFile(sb.String())
	if err != nil {
		panic(err)
	}

	return string(data)
}

type Directions int

const (
	UP Directions = iota
	DOWN
)

func isLevelSafe(levels []string) bool {

	// Check first change
	firstLevel, err := strconv.ParseFloat(levels[0], 32)
	checkError(err)

	secondLevel, err := strconv.ParseFloat(levels[1], 32)
	checkError(err)

	if math.Abs(firstLevel-secondLevel) < 1 || math.Abs(firstLevel-secondLevel) > 3 {
		return false
	}

	direction := UP
	if firstLevel > secondLevel {
		direction = DOWN
	}

	// Check the rest
	for i, v := range levels {
		// Skip first because we already checked it
		if i == 0 || i == len(levels)-1 {
			continue
		}

		level, err := strconv.ParseFloat(v, 32)
		checkError(err)

		nextLevel, err := strconv.ParseFloat(levels[i+1], 32)
		checkError(err)

		diff := level - nextLevel
		if math.Abs(diff) < 1 || math.Abs(diff) > 3 {
			return false
		}

		if direction == UP && diff > 0 {
			return false
		} else if direction == DOWN && diff < 0 {
			return false
		}
	}

	return true
}

type day struct {
}

func (d day) run1(test bool) {

	data := readProblemInput("02", test)
	reports := strings.Split(data, "\n")
	safeCount := 0
	for _, v := range reports {
		levels := strings.Fields(v)

		if isLevelSafe(levels) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func (d day) run2(test bool) {

}

func main() {
	currentDay := day{}
	currentDay.run1(false)
}
