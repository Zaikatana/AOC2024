package main

import (
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type node struct {
	value          int
	addition       *node
	multiplication *node
	concat         *node
}

func (n *node) addToNode(operators []string, operatorIdx int, expectedVal int) bool {
	if operatorIdx == len(operators) {
		if n.value == expectedVal {
			return true
		}

		return false
	}

	val, valErr := strconv.Atoi(operators[operatorIdx])
	check(valErr)

	n.addition = &node{value: n.value + val}
	n.multiplication = &node{value: n.value * val}

	return n.addition.addToNode(operators, operatorIdx+1, expectedVal) ||
		n.multiplication.addToNode(operators, operatorIdx+1, expectedVal)
}

func (n *node) addToNodeConcat(operators []string, operatorIdx int, expectedVal int) bool {
	if operatorIdx == len(operators) {
		if n.value == expectedVal {
			return true
		}

		return false
	}

	val, valErr := strconv.Atoi(operators[operatorIdx])
	check(valErr)

	n.addition = &node{value: n.value + val}
	n.multiplication = &node{value: n.value * val}

	concatRes, concatResErr := strconv.Atoi(strconv.Itoa(n.value) + operators[operatorIdx])
	check(concatResErr)
	n.concat = &node{value: concatRes}

	return n.addition.addToNodeConcat(operators, operatorIdx+1, expectedVal) ||
		n.multiplication.addToNodeConcat(operators, operatorIdx+1, expectedVal) ||
		n.concat.addToNodeConcat(operators, operatorIdx+1, expectedVal)
}

func main() {
	dat, err := os.ReadFile("input_7.txt")
	check(err)
	strDat := string(dat)
	p1Total := 0
	p2Total := 0

	calibrations := s.Split(strDat, "\n")

	for _, calibration := range calibrations {
		calibrationParams := s.Split(calibration, ": ")
		expectedResult, expectedResultErr := strconv.Atoi(calibrationParams[0])
		check(expectedResultErr)
		operators := s.Split(calibrationParams[1], " ")

		// create root node
		rootVal, _ := strconv.Atoi(operators[0])
		root := node{value: rootVal}

		if root.addToNode(operators, 1, expectedResult) {
			p1Total += expectedResult
		}

		root2 := node{value: rootVal}
		if root2.addToNodeConcat(operators, 1, expectedResult) {
			p2Total += expectedResult
		}
	}

	fmt.Println("Part 1 Solution", p1Total)
	fmt.Println("Part 2 Solution", p2Total)
}
