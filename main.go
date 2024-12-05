package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

type day struct {
}

func (d day) run1(test bool) {

	data := readProblemInput("03", test)

	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	muls := r.FindAllStringSubmatch(data, -1)

	total := 0
	for _, m := range muls {
		num1, _ := strconv.Atoi(m[1])
		num2, _ := strconv.Atoi(m[2])

		total += num1 * num2
	}

	fmt.Println(total)
}

type Instructions int

const (
	DO Instructions = iota
	DONT
)

func getInstructionLength(length int) Instructions {
	if length == 4 {
		return DO
	} else {
		return DONT
	}
}

func (d day) run2(test bool) {
	data := readProblemInput("03", test)

	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	muls := r.FindAllStringSubmatchIndex(data, -1)

	r2, _ := regexp.Compile(`don't\(\)|do\(\)`)
	instructions := r2.FindAllStringSubmatchIndex(data, -1)

	total := 0
	currentInstruction := DO

	done := false
	i1, i2 := 0, 0
	for !done {
		if i1 == len(muls)-1 {
			done = true
		}

		if instructions[i2][0] < muls[i1][0] {
			currentInstruction = getInstructionLength(instructions[i2][1] - instructions[i2][0])
			if i2 != len(instructions)-1 {
				i2++
			}
		}

		if currentInstruction == DO {
			num1, _ := strconv.Atoi(data[muls[i1][2]:muls[i1][3]])
			num2, _ := strconv.Atoi(data[muls[i1][4]:muls[i1][5]])

			total += num1 * num2
		}

		i1++
	}

	fmt.Println(total)
}

func main() {
	currentDay := day{}
	currentDay.run2(false)
}
