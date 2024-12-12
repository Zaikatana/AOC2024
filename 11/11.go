package main

import (
	"fmt"
	"maps"
	"os"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func blinkAction(m map[string]int) {
	origMap := maps.Clone(m)
	futureMap := map[string]int{}
	for stone := range origMap {
		if m[stone] > 0 {
			m[stone] = 0
			if stone == "0" {
				futureMap["1"] += origMap[stone]
			} else if len(stone)%2 == 0 {
				stoneA, _ := strconv.Atoi(stone[:len(stone)/2])
				stoneB, _ := strconv.Atoi(stone[len(stone)/2:])
				futureMap[strconv.Itoa(stoneA)] += origMap[stone]
				futureMap[strconv.Itoa(stoneB)] += origMap[stone]
			} else {
				stoneToInt, _ := strconv.Atoi(stone)
				futureMap[strconv.Itoa(stoneToInt*2024)] += origMap[stone]
			}
		}
	}
	clear(origMap)
	for stone := range futureMap {
		m[stone] = futureMap[stone]
	}
	clear(futureMap)
}

func main() {
	dat, err := os.ReadFile("input_11.txt")
	check(err)
	strDat := string(dat)
	p1Total := 0
	p2Total := 0

	stones := s.Split(strDat, " ")

	blink := 0
	m := map[string]int{}

	for _, stone := range stones {
		m[stone] += 1
	}

	for blink < 25 {
		blinkAction(m)
		blink += 1
	}

	for stone := range m {
		p1Total += m[stone]
	}

	for blink < 75 {
		blinkAction(m)
		blink += 1
	}

	for stone := range m {
		p2Total += m[stone]
	}

	fmt.Println("Part 1 solution", p1Total)
	fmt.Println("Part 2 solution", p2Total)
}
