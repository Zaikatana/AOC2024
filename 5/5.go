package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("input_5.txt")
	check(err)
	strDat := string(dat)
	processedDat := s.Split(strDat, "\n\n")
	p1Total := 0
	p2Total := 0

	updates := s.Split(processedDat[1], "\n")
	rules := s.Split(processedDat[0], "\n")

	// create hash map
	m := map[string][]string{}

	for _, rule := range rules {
		rulePages := s.Split(rule, "|")
		m[rulePages[0]] = append(m[rulePages[0]], rulePages[1])
	}

	problemUpdates := []string{}

	for _, update := range updates {
		updatePages := s.Split(update, ",")
		updateValid := true
		for i := 1; i < len(updatePages); i++ {
			for j := 0; j < i; j++ {
				if slices.Contains(m[updatePages[i]], updatePages[j]) {
					updateValid = false
					break
				}
			}
			if !updateValid {
				break
			}
		}
		// If update is valid add middle index
		if updateValid {
			middlePage := updatePages[int(len(updatePages)/2)]
			middlePageInt, err := strconv.Atoi(middlePage)
			check(err)
			p1Total += middlePageInt
		} else {
			problemUpdates = append(problemUpdates, update)
		}
	}

	for _, update := range problemUpdates {
		updatePages := s.Split(update, ",")
		for i := len(updatePages) - 1; i > 0; i-- {
			swapOccurred := false
			for j := i - 1; j >= 0; j-- {
				if slices.Contains(m[updatePages[i]], updatePages[j]) {
					temp := updatePages[i]
					updatePages[i] = updatePages[j]
					updatePages[j] = temp
					swapOccurred = true
					break
				}
			}
			if swapOccurred {
				i++
			}
		}
		middlePage := updatePages[int(len(updatePages)/2)]
		middlePageInt, err := strconv.Atoi(middlePage)
		check(err)
		p2Total += middlePageInt
	}

	fmt.Println("Part 1 Solution", p1Total)
	fmt.Println("Part 2 Solution", p2Total)
}
