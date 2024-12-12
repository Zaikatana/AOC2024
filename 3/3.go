package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processMul(mulSlice []string) int {
	total := 0

	for _, mul_element := range mulSlice {
		num_r, _ := regexp.Compile(`[0-9]+`)
		num := num_r.FindAllString(mul_element, -1)

		int_a, err_a := strconv.Atoi(num[0])
		int_b, err_b := strconv.Atoi(num[1])
		check(err_a)
		check(err_b)

		total += (int_a * int_b)
	}

	return total
}

func main() {
	dat, err := os.ReadFile("input_3.txt")
	check(err)
	strDat := string(dat)
	mulR, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	mulSlice := mulR.FindAllString(strDat, -1)

	p1Total := processMul(mulSlice)
	p2Total := p1Total

	condR, _ := regexp.Compile(`don't\(\).+?do\(\)|don't\(\).+`)
	condSlice := condR.FindAllString(strDat, -1)

	fmt.Println(len(condSlice))

	for _, condElement := range condSlice {
		fmt.Println(condElement, "\n")
		mulSlice := mulR.FindAllString(condElement, -1)
		condTotal := processMul(mulSlice)

		p2Total -= condTotal
	}

	fmt.Println("Part 1 Solution", p1Total)
	fmt.Println("Part 2 Solution", p2Total)
}
