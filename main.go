package main

import (
	"fmt"
	"os"
	"slices"
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
	data := readProblemInput("01", test)
	lines := strings.Split(data, "\n")
	var list1 []int
	var list2 []int
	for _, v := range lines {
		values := strings.Fields(v)

		first, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, first)

		second, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		list2 = append(list2, second)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	diffTotal := 0
	for i, v := range list1 {
		diff := 0
		if list2[i] > v {
			diff = list2[i] - v
		} else {
			diff = v - list2[i]
		}

		diffTotal += diff
	}

	fmt.Println(diffTotal)
}

func (d day) run2(test bool) {
	data := readProblemInput("01", test)
	lines := strings.Split(data, "\n")

	var list1 []int
	list2 := make(map[int]int)

	for _, v := range lines {
		values := strings.Fields(v)

		first, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		list1 = append(list1, first)

		second, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}

		list2[second] = list2[second] + 1
	}

	total := 0
	for _, v := range list1 {
		total += v * list2[v]
	}

	fmt.Println(total)
}

func main() {
	currentDay := day{}
	currentDay.run2(false)
}
