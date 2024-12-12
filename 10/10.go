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
	dat, err := os.ReadFile("input_10.txt")
	check(err)
	strDat := string(dat)
	p1Total := 0
	p2Total := 0

	rows := s.Split(strDat, "\n")
	trailheads := [][]int{}

	// find all 0s (Trailhead starts)
	for rowIdx, row := range rows {
		for colIdx, col := range row {
			if col == '0' {
				trailheadIdx := []int{rowIdx, colIdx}
				trailheads = append(trailheads, trailheadIdx)
			}
		}
	}

	// find all ways that 0 can get to 9 (score)
	for _, trailheadIdx := range trailheads {
		m := map[string]bool{}
		stack := [][]int{trailheadIdx}
		// pop stack
		// if visited return 0
		// else add to map
		for len(stack) > 0 {
			poppedElement := stack[0]
			stack = slices.Delete(stack, 0, 1)
			mapKey := strconv.Itoa(poppedElement[0]) + "|" + strconv.Itoa(poppedElement[1])
			_, prs := m[mapKey]
			if prs {
				continue
			}
			m[mapKey] = true

			// find valid neighbours, add them to stack
			poppedElementNum, poppedElementNumErr := strconv.Atoi(string(rows[poppedElement[0]][poppedElement[1]]))
			check(poppedElementNumErr)

			if poppedElementNum == 9 {
				p1Total += 1
				continue
			}

			if poppedElement[0]+1 < len(rows) && string(rows[poppedElement[0]+1][poppedElement[1]]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0] + 1, poppedElement[1]})
			}

			if poppedElement[1]+1 < len(rows[0]) && string(rows[poppedElement[0]][poppedElement[1]+1]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0], poppedElement[1] + 1})
			}

			if poppedElement[0]-1 >= 0 && string(rows[poppedElement[0]-1][poppedElement[1]]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0] - 1, poppedElement[1]})
			}

			if poppedElement[1]-1 >= 0 && string(rows[poppedElement[0]][poppedElement[1]-1]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0], poppedElement[1] - 1})
			}
		}

		// For part 2
		stack = [][]int{trailheadIdx}
		for len(stack) > 0 {
			poppedElement := stack[0]
			stack = slices.Delete(stack, 0, 1)

			// find valid neighbours, add them to stack
			poppedElementNum, poppedElementNumErr := strconv.Atoi(string(rows[poppedElement[0]][poppedElement[1]]))
			check(poppedElementNumErr)

			if poppedElementNum == 9 {
				p2Total += 1
				continue
			}

			if poppedElement[0]+1 < len(rows) && string(rows[poppedElement[0]+1][poppedElement[1]]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0] + 1, poppedElement[1]})
			}

			if poppedElement[1]+1 < len(rows[0]) && string(rows[poppedElement[0]][poppedElement[1]+1]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0], poppedElement[1] + 1})
			}

			if poppedElement[0]-1 >= 0 && string(rows[poppedElement[0]-1][poppedElement[1]]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0] - 1, poppedElement[1]})
			}

			if poppedElement[1]-1 >= 0 && string(rows[poppedElement[0]][poppedElement[1]-1]) == strconv.Itoa(poppedElementNum+1) {
				stack = append(stack, []int{poppedElement[0], poppedElement[1] - 1})
			}
		}
	}

	fmt.Println("Part 1 solution", p1Total)
	fmt.Println("Part 2 solution", p2Total)
}
